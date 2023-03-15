package impl

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	"github.com/tqtcloud/workflow-backend/apps/task"
	"strings"
)

func (s *service) CreateTask(ctx context.Context, req *task.CreateTaskRequest) (*task.Task, error) {
	ins, err := task.NewTask(req)
	if err != nil {
		s.log.Errorf("validate create task error, %s", err)
		return nil, exception.NewBadRequest("validate create task error, %s", err)
	}

	switch ins.Data.Env {
	case task.JenkinsEnv_DEV:
		ins, err = createJob(ctx, ins, s.conf)
		if err != nil {
			s.log.Errorf("validate create task error, %s", err)
			return nil, exception.NewInternalServerError("validate create task error, %s", err)
		}
	case task.JenkinsEnv_TEST:
		ins, err = createJob(ctx, ins, s.conf)
		if err != nil {
			s.log.Errorf("validate create task error, %s", err)
			return nil, exception.NewInternalServerError("validate create task error, %s", err)
		}
	case task.JenkinsEnv_UAT:
		ins, err = createJob(ctx, ins, s.conf)
		if err != nil {
			s.log.Errorf("validate create task error, %s", err)
			return nil, exception.NewInternalServerError("validate create task error, %s", err)
		}
	case task.JenkinsEnv_LPT:
		ins, err = createJob(ctx, ins, s.conf)
		if err != nil {
			s.log.Errorf("validate create task error, %s", err)
			return nil, exception.NewInternalServerError("validate create task error, %s", err)
		}
	case task.JenkinsEnv_PROD:
		ins, err = createJob(ctx, ins, s.conf)
		if err != nil {
			s.log.Errorf("validate create task error, %s", err)
			return nil, exception.NewInternalServerError("validate create task error, %s", err)
		}
	default:
		s.log.Errorf("Request JenkinsEnv  error, %s", ins.Data.Env)
		return nil, exception.NewBadRequest("Request JenkinsEnv  error,%s  ", ins.Data.Env)
	}
	//if err := s.save(ctx, ins); err != nil {
	//	return nil, err
	//}
	s.log.Infof("%s 环境 jenkins Job %s 创建成功,目录位与：%s", task.JenkinsEnv_DEV.String(), ins.Data.JobName, ins.Data.Folder)
	return ins, nil
}

func (s *service) CopyTask(ctx context.Context, req *task.CreateTaskRequest) (*task.Task, error) {
	ins, err := task.NewTask(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create task error, %s", err)
	}

	switch ins.Data.Env {
	case task.JenkinsEnv_DEV:
		jenkins := gojenkins.CreateJenkins(nil, s.conf.Jenkins.DevEndpoints, s.conf.Jenkins.User, s.conf.Jenkins.Password)
		_, err = jenkins.Init(ctx)
		if err != nil {
			s.log.Printf("连接Jenkins失败, %v\n", err)
			return nil, err
		}
		s.log.Debug("Jenkins连接成功")
		// 注意创建逻辑会在指定的 oldjob 目录创建相同不同名称的job
		// 例如：sit_lebei/job/OldjobName
		// 创建后： sit_lebei/job/NewjobName
		_, err = jenkins.CopyJob(ctx, ins.Data.OldjobName, ins.Data.NewjobName)

		if err != nil {
			s.log.Printf("CreateJob error %s \n", err)
			return nil, err
		}
		s.log.Debug("连接 JenkinsEnv_DEV 创建 job %s ", ins.Data.Env)
		s.log.Debugf("job复制成功. %s 位于文件夹 %s", ins.Data.JobName, ins.Data.Folder)

	case task.JenkinsEnv_TEST:
		s.log.Debug(s.conf.Jenkins.DevEndpoints, s.conf.Jenkins.User, s.conf.Jenkins.Password)

		s.log.Debug("连接 JenkinsEnv_TEST 创建 job %s ", ins.Data.Env)

	case task.JenkinsEnv_UAT:
		s.log.Debug("连接 JenkinsEnv_UAT 创建 job %s ", ins.Data.Env)

	case task.JenkinsEnv_LPT:
		s.log.Debug("连接 JenkinsEnv_LPT 创建 job %s  ", ins.Data.Env)

	case task.JenkinsEnv_PROD:
		s.log.Debugf("连接 JenkinsEnv_PROD 创建 job %s ", ins.Data.Env)

	default:
		return nil, exception.NewBadRequest("Request JenkinsEnv  error,%s  ", ins.Data.Env)
	}

	//if err := s.save(ctx, ins); err != nil {
	//	return nil, err
	//}

	return ins, nil
}

func (s *service) DescribeTask(ctx context.Context, req *task.DescribeTaskRequest) (*task.Task, error) {
	//return s.get(ctx, req.Id)
	s.log.Debugf("DescribeTask req :%s ", req)
	jenkins, err := envDecision(ctx, req.Env, s.conf)
	if err != nil {
		s.log.Errorf("envDecision error, %s", err)
		return nil, exception.NewBadRequest("envDecision error, %s", err)
	}
	ins, err := describeJob(ctx, req, jenkins)
	if err != nil {
		return nil, exception.NewBadRequest("describe  Job error, %s", err)
	}
	//if err := s.save(ctx, ins); err != nil {
	//	return nil, err
	//}
	s.log.Infof("DescribeTask %s ", ins)
	return ins, nil
}

func (s *service) QueryTask(ctx context.Context, req *task.QueryTaskRequest) (*task.TaskSet, error) {
	//query := newQueryTaskRequest(req)
	//return s.query(ctx, query)
	return nil, nil
}

func (s *service) UpdateTask(ctx context.Context, req *task.UpdateTaskRequest) (*task.Task, error) {
	s.log.Debug("UpdateTask Req : ", req.Data)
	ins, err := s.DescribeTask(ctx, task.NewDescribeTaskRequest(req.Data.Env.String(), req.Data.Folder, req.Data.JobName))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	// 全量更新
	case request.UpdateMode_PUT:
		ins = ins.Update(req)
	// 局部更新资源
	case request.UpdateMode_PATCH:
		ins = ins.Patch(req)
	}

	// 校验更新后数据合法性
	if err := ins.Data.Validate(); err != nil {
		return nil, err
	}
	updateIns, err := updateJob(ctx, ins, s.conf)
	//if err := s.update(ctx, ins); err != nil {
	//	return nil, err
	//}

	//return ins, nil
	return updateIns, nil

}

func (s *service) DeleteTask(ctx context.Context, req *task.DeleteTaskRequest) (*task.Task, error) {
	ins, err := s.DescribeTask(ctx, task.NewDescribeTaskRequest(req.Env, req.Folder, req.Jobname))
	if err != nil {
		return nil, err
	}
	//if err := s.deleteTask(ctx, ins); err != nil {
	//	return nil, err
	//}
	env, _ := task.ParseJenkinsEnvFromString(req.Env)
	jenins, err := envDecision(ctx, env, s.conf)
	if err != nil {
		return nil, fmt.Errorf("删除job时连接 jenkins 失败: %s", err)
	}
	if err = delJob(ctx, req, jenins); err != nil {
		return nil, fmt.Errorf("删除job 失败: %s", err)
	}
	s.log.Infof("job 删除成功：%s", ins)
	return ins, nil
}

func (s *service) SshExec(ctx context.Context, req *task.ExecRequest) (*task.ExecResp, error) {
	defer func() {
		if r := recover(); r != nil {
			s.log.Errorf("recover...  远程主机失败: %s", r)
		}
	}()

	ins, err := task.NewExecRequest(req)
	if err != nil {
		s.log.Errorf("validate create task SshExec  error, %s", err)
		return nil, exception.NewBadRequest("validate create SshExec error, %s", err)
	}

	switch ins.Env {
	case task.JenkinsEnv_DEV:
		message, err := task.PasswordConnect(s.conf.SshExec.User, s.conf.SshExec.DevNode, s.conf.SshExec.DevPort, "Password1", ins.Type, ins.Name, ins.Port, s.conf.SshExec.SshShell)
		if err != nil {
			s.log.Errorf("远程执行主机命令错误  error, %s", err)
			return nil, exception.NewBadRequest("远程执行主机命令错误 error, %s", err)
		}
		s.log.Infof("开发环境主机命令执行完毕输出: %s", message)
		return &task.ExecResp{
			Data:    ins,
			Message: strings.TrimRight(message, "\n"),
		}, nil

	case task.JenkinsEnv_TEST:
		message, err := task.PasswordConnect(s.conf.SshExec.User, s.conf.SshExec.TestNode, s.conf.SshExec.TestPort, "Password1", ins.Type, ins.Name, ins.Port, s.conf.SshExec.SshShell)
		if err != nil {
			s.log.Errorf("远程执行主机命令错误  error, %s", err)
			return nil, exception.NewBadRequest("远程执行主机命令错误 error, %s", err)
		}
		s.log.Infof("测试环境主机命令执行完毕输出: %s", message)
		return &task.ExecResp{
			Data:    ins,
			Message: strings.TrimRight(message, "\n"),
		}, nil
	case task.JenkinsEnv_UAT:
		message, err := task.ExEcShell(s.conf.SshExec.User, s.conf.SshExec.UatNode, s.conf.SshExec.UatPort, s.conf.SshExec.UatSshKeyPath, ins.Type, ins.Name, ins.Port, s.conf.SshExec.SshShell)
		if err != nil {
			s.log.Errorf("远程执行主机命令错误  error, %s", err)
			return nil, exception.NewBadRequest("远程执行主机命令错误 error, %s", err)
		}
		s.log.Infof("预发环境主机命令执行完毕输出: %s", message)
		return &task.ExecResp{
			Data:    ins,
			Message: strings.TrimRight(message, "\n"),
		}, nil
	case task.JenkinsEnv_LPT:
		s.log.Error("压测环境已经消失了")
		return nil, exception.NewBadRequest("Request JenkinsEnv  error,%s  ", ins.Env)
	case task.JenkinsEnv_PROD:
		message, err := task.ExEcShell(s.conf.SshExec.User, s.conf.SshExec.UatNode, s.conf.SshExec.UatPort, s.conf.SshExec.ProdSshKeyPath, ins.Type, ins.Name, ins.Port, s.conf.SshExec.SshShell)
		if err != nil {
			s.log.Errorf("远程执行主机命令错误  error, %s", err)
			return nil, exception.NewBadRequest("远程执行主机命令错误 error, %s", err)
		}
		s.log.Infof("生产环境主机命令执行完毕输出: %s", message)
		return &task.ExecResp{
			Data:    ins,
			Message: strings.TrimRight(message, "\n"),
		}, nil
	default:
		s.log.Errorf("Request JenkinsEnv  error, %s", ins.Env)
		return nil, exception.NewBadRequest("Request JenkinsEnv  error,%s  ", ins.Env)
	}
}
