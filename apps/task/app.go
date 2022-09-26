package task

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
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

func NewCreateTaskRequest() *CreateTaskRequest {
	return &CreateTaskRequest{}
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

func (i *Task) Update(req *UpdateTaskRequest) {
	i.UpdateAt = time.Now().UnixMicro()
	i.UpdateBy = req.UpdateBy
	i.Data = req.Data
}

func (i *Task) Patch(req *UpdateTaskRequest) error {
	i.UpdateAt = time.Now().UnixMicro()
	i.UpdateBy = req.UpdateBy
	return mergo.MergeWithOverwrite(i.Data, req.Data)
}

func NewDescribeTaskRequest(jobname, env string) *DescribeTaskRequest {
	jenkinsEnv, err := ParseJenkinsEnvFromString(env)
	if err != nil {
		fmt.Printf("ParseJenkinsEnvFromString Error: %s\n", err)
	}
	return &DescribeTaskRequest{
		Jobname: jobname,
		Env:     jenkinsEnv,
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

func NewPutTaskRequest(id string) *UpdateTaskRequest {
	return &UpdateTaskRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateTaskRequest(),
	}
}

func NewPatchTaskRequest(id string) *UpdateTaskRequest {
	return &UpdateTaskRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateTaskRequest(),
	}
}

func NewDeleteTaskRequestWithID(id string) *DeleteTaskRequest {
	return &DeleteTaskRequest{
		Id: id,
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
