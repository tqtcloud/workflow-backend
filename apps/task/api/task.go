package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"

	"github.com/tqtcloud/workflow-backend/apps/task"
)

func (h *handler) CreateTask(r *restful.Request, w *restful.Response) {
	req := task.NewDefaultCreateTaskRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	h.log.Debugf("url 入参: %s", req)
	set, err := h.service.CreateTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}

func (h *handler) CopyTask(r *restful.Request, w *restful.Response) {
	req := task.NewDefaultCreateTaskRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	h.log.Debugf("url 入参: %s", req)
	set, err := h.service.CopyTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}

//func (h *handler) QueryTask(r *restful.Request, w *restful.Response) {
//	req := task.NewQueryTaskRequestFromHTTP(r.Request)
//	set, err := h.service.QueryTask(r.Request.Context(), req)
//	if err != nil {
//		response.Failed(w.ResponseWriter, err)
//		return
//	}
//	response.Success(w.ResponseWriter, set)
//}
//
func (h *handler) DescribeTask(r *restful.Request, w *restful.Response) {
	req := task.NewDescribeTaskRequest(r.PathParameter("env"), r.PathParameter("folder"), r.PathParameter("jobname"))
	ins, err := h.service.DescribeTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, ins)
}

func (h *handler) UpdateTask(r *restful.Request, w *restful.Response) {
	req := task.NewPutTaskRequest(r.PathParameter("env"), r.PathParameter("folder"), r.PathParameter("jobname"))
	h.log.Debug("NewDescribeTaskRequest  %s", req)

	if err := r.ReadEntity(req.Data); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.UpdateTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (h *handler) PatchTask(r *restful.Request, w *restful.Response) {
	req := task.NewPatchTaskRequest(r.PathParameter("env"), r.PathParameter("folder"), r.PathParameter("jobname"))

	if err := r.ReadEntity(req.Data); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	h.log.Debug("PatchTask NewDescribeTaskRequest  %s", req.Data)
	set, err := h.service.UpdateTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (h *handler) DeleteTask(r *restful.Request, w *restful.Response) {
	req := task.NewDeleteTaskRequestWithID(r.PathParameter("env"), r.PathParameter("folder"), r.PathParameter("jobname"))
	set, err := h.service.DeleteTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (h *handler) SshExec(r *restful.Request, w *restful.Response) {
	req := task.NewDefaultExecRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	h.log.Debugf("url 入参: %s", req)
	set, err := h.service.SshExec(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}
