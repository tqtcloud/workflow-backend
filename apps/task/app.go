package task

import (
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

func NewDescribeTaskRequest(id string) *DescribeTaskRequest {
	return &DescribeTaskRequest{
		Id: id,
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
