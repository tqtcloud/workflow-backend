package task

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/infraboard/mcube/exception"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/mitchellh/go-homedir"
	"github.com/rs/xid"
	"golang.org/x/crypto/ssh"
)

const (
	AppName = "task"
)

var (
	validate = validator.New()
)

func NewDefaultCreateTaskRequest() *CreateTaskRequest {
	return &CreateTaskRequest{}
}

func NewCreateTaskRequest(env, folder, jobname string) *CreateTaskRequest {
	Envs, _ := ParseJenkinsEnvFromString(env)
	return &CreateTaskRequest{
		Env:     Envs,
		Folder:  folder,
		JobName: jobname,
	}
}

func NewTask(req *CreateTaskRequest) (*Task, error) {
	if err := req.Validate(); err != nil {
		log.Println(req)
		return nil, err
	}

	return &Task{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMicro(),
		Data:     req,
	}, nil
}

func (req *CreateTaskRequest) Validate() error {
	return validate.Struct(req)
}

func NewTaskSet() *TaskSet {
	return &TaskSet{
		Items: []*Task{},
	}
}

func (s *TaskSet) Add(item *Task) {
	s.Items = append(s.Items, item)
}

func NewDefaultTask() *Task {
	return &Task{
		Data: &CreateTaskRequest{},
	}
}

func (i *Task) Update(req *UpdateTaskRequest) *Task {
	i.UpdateAt = time.Now().UnixMicro()
	i.UpdateBy = req.UpdateBy
	i.Data = req.Data
	return i
}

func (i *Task) Patch(req *UpdateTaskRequest) *Task {
	i.UpdateAt = time.Now().UnixMicro()
	i.UpdateBy = req.UpdateBy
	i.Data = req.Data
	return i
}

func NewDescribeTaskRequest(env, folder, jobname string) *DescribeTaskRequest {
	jenkinsEnv, err := ParseJenkinsEnvFromString(env)
	if err != nil {
		fmt.Printf("ParseJenkinsEnvFromString Error: %s \n", err)
	}
	return &DescribeTaskRequest{
		Jobname: jobname,
		Env:     jenkinsEnv,
		Folder:  folder,
	}
}

func NewQueryTaskRequest() *QueryTaskRequest {
	return &QueryTaskRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewQueryTaskRequestFromHTTP(r *http.Request) *QueryTaskRequest {
	qs := r.URL.Query()

	return &QueryTaskRequest{
		Page:     request.NewPageRequestFromHTTP(r),
		Keywords: qs.Get("keywords"),
	}
}

func NewPutTaskRequest(env, folder, jobname string) *UpdateTaskRequest {
	return &UpdateTaskRequest{
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateTaskRequest(env, folder, jobname),
	}
}

func NewPatchTaskRequest(env, folder, jobname string) *UpdateTaskRequest {
	return &UpdateTaskRequest{
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateTaskRequest(env, folder, jobname),
	}
}

func NewDeleteTaskRequestWithID(env, folder, jobname string) *DeleteTaskRequest {
	return &DeleteTaskRequest{
		Env:     env,
		Folder:  folder,
		Jobname: jobname,
	}
}

// ConnectJenkins 连接jenkins服务器获取客户端
func ConnectJenkins(ctx context.Context, url, user, password string) (*gojenkins.Jenkins, error) {
	jenkins := gojenkins.CreateJenkins(nil, url, user, password)
	_, err := jenkins.Init(ctx)
	if err != nil {
		return nil, fmt.Errorf("连接Jenkins失败, %v\n", err)
	}
	return jenkins, nil
}

// publicKeyAuthFunc 根据目录公钥解析出认证信息
func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		fmt.Printf("查找密钥的主目录失败: %s", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		fmt.Printf("ssh 密钥文件读取失败: %s", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Printf("ssh 关键签名失败", err)
	}
	return ssh.PublicKeys(signer)
}

// NewSshLoginConfig 创建ssh远程登陆配置(通过秘钥登录)  使用后需要关闭ssh通道
// defer sshClient.Close()
func NewSshLoginConfig(sshUser, sshHost, sshPort string, sshKeyPath string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second * 5, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以, 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", sshHost, sshPort), config)
	if err != nil {
		log.Printf("创建ssh client 失败 %s\n", err)
	}
	return sshClient, nil
}

func (req *ExecRequest) Validate() error {
	return validate.Struct(req)
}

func NewExecRequest(req *ExecRequest) (*ExecRequest, error) {
	if err := req.Validate(); err != nil {
		log.Println(req)
		return nil, err
	}

	return &ExecRequest{
		Env:  req.Env,
		Type: req.Type,
		Name: req.Name,
		Port: req.Port,
	}, nil
}

func NewDefaultExecRequest() *ExecRequest {
	return &ExecRequest{}
}

// ExEcShell 根据传入的相关主机信息，远程主机执行命令，然后关闭通道
func ExEcShell(sshUser, sshHost, sshPort, sshKeyPath, Type, Name, Port, sshShell string) (string, error) {
	sshClient, err := NewSshLoginConfig(sshUser, sshHost, sshPort, sshKeyPath)
	if err != nil {
		return "", exception.NewBadRequest("创建ssh client 失败, %s", err)
	}
	defer sshClient.Close()
	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		return "", exception.NewBadRequest("创建ssh session 失败, %s", err)
	}
	defer session.Close()
	//执行远程命令
	cmdShell := fmt.Sprintf("%s %s %s %s", sshShell, Type, Name, Port)
	combo, err := session.CombinedOutput(cmdShell)
	if err != nil {
		return "", exception.NewBadRequest("远程执行cmd 失败", err)
	}

	return string(combo), nil
}

// PasswordConnect 没秘钥使用密码连接的
func PasswordConnect(sshUser, sshHost, sshPort, password, Type, Name, Port, sshShell string) (string, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:            sshUser,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%s", sshHost, sshPort)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return "", exception.NewBadRequest("创建ssh client 失败, %s", err)
	}

	defer sshClient.Close()
	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		return "", exception.NewBadRequest("创建ssh session 失败, %s", err)
	}
	defer session.Close()
	//执行远程命令
	cmdShell := fmt.Sprintf("%s %s %s %s", sshShell, Type, Name, Port)
	combo, err := session.CombinedOutput(cmdShell)
	if err != nil {
		return "", exception.NewBadRequest("远程执行cmd 失败", err)
	}

	return string(combo), nil
}
