package impl

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/infraboard/mcube/exception"
	"github.com/tqtcloud/workflow-backend/apps/task"
	"strings"
)

// createJob 创建job逻辑实现
func (s *service) createJob(ctx context.Context, ins *task.Task, env task.JenkinsEnv) (*task.Task, error) {
	jenkins, err := s.envDecision(ctx, env)
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
		return nil, exception.NewInternalServerError("Job CreateJobInFolder error, %s  JobName: %s", err, ins.Data.JobName)
	}
	s.log.Printf("%s 环境 jenkins Job %s 创建成功,目录位与：%s", task.JenkinsEnv_DEV, job.GetName(), ins.Data.Folder)
	return ins, nil
}

// envDecision jenkins 环境变量判定连接
func (s *service) envDecision(ctx context.Context, env task.JenkinsEnv) (*gojenkins.Jenkins, error) {
	switch env {
	case task.JenkinsEnv_DEV:
		jenkins, err := task.ConnectJenkins(ctx, s.conf.Jenkins.DevEndpoints, s.conf.Jenkins.User, s.conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		return jenkins, nil
	case task.JenkinsEnv_TEST:
		jenkins, err := task.ConnectJenkins(ctx, s.conf.Jenkins.TestEndpoints, s.conf.Jenkins.User, s.conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		return jenkins, nil
	case task.JenkinsEnv_UAT:
		jenkins, err := task.ConnectJenkins(ctx, s.conf.Jenkins.UatEndpoints, s.conf.Jenkins.User, s.conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		return jenkins, nil
	case task.JenkinsEnv_LPT:
		jenkins, err := task.ConnectJenkins(ctx, s.conf.Jenkins.LptEndpoints, s.conf.Jenkins.User, s.conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		return jenkins, nil
	case task.JenkinsEnv_PROD:
		jenkins, err := task.ConnectJenkins(ctx, s.conf.Jenkins.ProdEndpoints, s.conf.Jenkins.User, s.conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		return jenkins, nil
	default:
		return nil, fmt.Errorf("环境错误: %s", "dev/test/uat/lpt/prod")
	}
}

// describeJob 查看job信息
func (s *service) describeJob(ctx context.Context, req *task.DescribeTaskRequest, jenkins *gojenkins.Jenkins) (*task.Task, error) {
	return s.getJobConfig(ctx, req, jenkins)
}

func (s *service) getJobConfig(ctx context.Context, req *task.DescribeTaskRequest, jenkins *gojenkins.Jenkins) (*task.Task, error) {
	ins := new(task.Task)
	job, err := jenkins.GetJob(ctx, req.Jobname)
	if err != nil {
		return nil, exception.NewInternalServerError("validate create task error, %s", err)
	}
	config, err := job.GetConfig(ctx)
	if err != nil {
		return nil, exception.NewInternalServerError("validate create task error, %s", err)
	}
	config = strings.TrimPrefix(config, `<?xml version="1.1" encoding="UTF-8" standalone="no"?>`)
	config = strings.TrimPrefix(config, `<?xml version='1.1' encoding='UTF-8'?>`)
	data := new(Project)
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, exception.NewInternalServerError("Job config Unmarshal error, %s", err)
	}
	s.log.Debug(data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL)
	s.log.Debug(data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.Branch)
	s.log.Debug(data.Builders.HudsonTasksShell.Command)
	s.log.Debug(data.Description)
	ins.Data.GitUrl = fmt.Sprintln(data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL)
	ins.Data.Branch = data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.Branch
	ins.Data.Buildeshell = data.Builders.HudsonTasksShell.Command
	ins.Data.Description = data.Description
	// appname
	ins.Data.AppName = data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition.DefaultValue
	return ins, nil
}
