package impl

import (
	"context"
	"encoding/xml"
	"github.com/bndr/gojenkins"
	"github.com/infraboard/mcube/exception"
	"github.com/tqtcloud/workflow-backend/apps/task"
	"os"
	"strings"
)

func (s *service) CreateTask(ctx context.Context, req *task.CreateTaskRequest) (*task.Task, error) {
	ins, err := task.NewTask(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create task error, %s", err)
	}

	switch ins.Data.Env {
	case task.JenkinsEnv_DEV:
		jenkins, err := task.ConnectJenkins(ctx, s.conf.Jenkins.DevEndpoints, s.conf.Jenkins.User, s.conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		// 前端定义好下拉列表在此处进行模板的选择，后端Java? go? nodejs?
		// jobtemplate/job/go-backend-template  go 模板名
		templateName := ins.Data.TemplateName
		job, err := jenkins.GetJob(ctx, templateName)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		// job.UpdateConfig()
		// jenkins.UpdateJob()
		config, _ := job.GetConfig(ctx)

		config = strings.TrimPrefix(config, `<?xml version="1.1" encoding="UTF-8" standalone="no"?>`)
		config = strings.TrimPrefix(config, `<?xml version='1.1' encoding='UTF-8'?>`)

		s.log.Debug(config)
		data := Project{}
		if err := xml.Unmarshal([]byte(config), &data); err != nil {
			s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
			return nil, exception.NewInternalServerError("Job config Unmarshal error, %s", err)
		}

		data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL = ins.Data.GitUrl
		data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.Branch = ins.Data.Branch
		if ins.Data.Buildeshell != "" {
			data.Builders.HudsonTasksShell.Command = ins.Data.Buildeshell
		}
		data.Description = ins.Data.Description
		// appname
		data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition.DefaultValue = ins.Data.AppName

		xmlData, err := xml.MarshalIndent(&data, " ", " ")
		if err != nil {
			s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
			return nil, exception.NewInternalServerError("Job config MarshalIndent error, %s", err)
		}

		job, err = jenkins.CreateJobInFolder(ctx, string(xmlData), ins.Data.JobName, ins.Data.Folder)
		if err != nil {
			s.log.Errorf("jenkins CreateJobInFolder error：%s,   job名称：%s", err, ins.Data.JobName)
			os.Exit(-1)
		}
		s.log.Printf("%s 环境 jenkins Job %s 创建成功,目录位与：%s", task.JenkinsEnv_DEV, job.GetName(), ins.Data.Folder)
		// 创建在 根 目录的 job
		// _, err = jenkins.CreateJob(ctx, jobXml, "apijob1")
	case task.JenkinsEnv_TEST:
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

func (s *service) DescribeTask(ctx context.Context, req *task.DescribeTaskRequest) (
	*task.Task, error) {
	//return s.get(ctx, req.Id)
	return nil, nil

}

func (s *service) QueryTask(ctx context.Context, req *task.QueryTaskRequest) (*task.TaskSet, error) {
	//query := newQueryTaskRequest(req)
	//return s.query(ctx, query)
	return nil, nil

}

func (s *service) UpdateTask(ctx context.Context, req *task.UpdateTaskRequest) (*task.Task, error) {
	//ins, err := s.DescribeTask(ctx, task.NewDescribeTaskRequest(req.Id))
	//if err != nil {
	//	return nil, err
	//}
	//
	//switch req.UpdateMode {
	//case request.UpdateMode_PUT:
	//	ins.Update(req)
	//case request.UpdateMode_PATCH:
	//	err := ins.Patch(req)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//
	//// 校验更新后数据合法性
	//if err := ins.Data.Validate(); err != nil {
	//	return nil, err
	//}
	//
	//if err := s.update(ctx, ins); err != nil {
	//	return nil, err
	//}

	//return ins, nil
	return nil, nil

}

func (s *service) DeleteTask(ctx context.Context, req *task.DeleteTaskRequest) (*task.Task, error) {
	//ins, err := s.DescribeTask(ctx, task.NewDescribeTaskRequest(req.Id))
	//if err != nil {
	//	return nil, err
	//}
	//
	//if err := s.deleteTask(ctx, ins); err != nil {
	//	return nil, err
	//}
	//
	//return ins, nil
	return nil, nil
}
