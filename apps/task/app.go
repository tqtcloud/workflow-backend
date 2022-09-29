package task

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/rs/xid"
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
