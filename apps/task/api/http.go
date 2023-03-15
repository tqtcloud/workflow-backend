package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/tqtcloud/workflow-backend/apps/task"
)

var (
	h = &handler{}
)

type handler struct {
	service task.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(task.AppName)
	h.service = app.GetGrpcApp(task.AppName).(task.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return task.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"tasks"}

	ws.Route(ws.POST("").To(h.CreateTask).
		Doc("create a jenkins job").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(task.CreateTaskRequest{}).
		Writes(response.NewData(task.Task{})))

	ws.Route(ws.POST("exec").To(h.SshExec).
		Doc("ssh exec shell").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(task.ExecRequest{}).
		Writes(response.NewData(task.ExecResp{})))

	ws.Route(ws.PUT("").To(h.CopyTask).
		Doc("CopyTask a  jenkins job").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(task.CreateTaskRequest{}).
		Writes(response.NewData(task.Task{})))
	//ws.Route(ws.GET("/").To(h.QueryTask).
	//	Doc("get all tasks").
	//	Metadata(restfulspec.KeyOpenAPITags, tags).
	//	Metadata("action", "list").
	//	Reads(task.CreateTaskRequest{}).
	//	Writes(response.NewData(task.TaskSet{})).
	//	Returns(200, "OK", task.TaskSet{}))
	//
	ws.Route(ws.GET("/{env}/{folder}/{jobname}").To(h.DescribeTask).
		Doc("get jenkins job config ").
		Param(ws.PathParameter("env", "identifier of the jenkins env").DataType("string").DefaultValue("dev")).
		Param(ws.PathParameter("folder", "identifier of the folder").DataType("string").DefaultValue("")).
		Param(ws.PathParameter("jobname", "identifier of the jobname").DataType("string").DefaultValue("jobtemplate/job/go-backend-template")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(response.NewData(task.Task{})).
		Returns(200, "OK", response.NewData(task.Task{})).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{env}/{folder}/{jobname}").To(h.UpdateTask).
		Doc("update jenkins job").
		Param(ws.PathParameter("env", "identifier of the jenkins env").DataType("string")).
		Param(ws.PathParameter("folder", "identifier of the folder").DataType("string")).
		Param(ws.PathParameter("jobname", "identifier of the jobname").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(task.CreateTaskRequest{}))

	ws.Route(ws.PATCH("/{env}/{folder}/{jobname}").To(h.PatchTask).
		Doc("patch jenkins job").
		Param(ws.PathParameter("env", "identifier of the jenkins env").DataType("string")).
		Param(ws.PathParameter("folder", "identifier of the folder").DataType("string")).
		Param(ws.PathParameter("jobname", "identifier of the jobname").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(task.CreateTaskRequest{}))

	ws.Route(ws.DELETE("/{env}/{folder}/{jobname}").To(h.DeleteTask).
		Doc("delete jenkins job").
		Param(ws.PathParameter("env", "identifier of the jenkins env").DataType("string")).
		Param(ws.PathParameter("folder", "identifier of the folder").DataType("string")).
		Param(ws.PathParameter("jobname", "identifier of the jobname").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

func init() {
	app.RegistryRESTfulApp(h)
}
