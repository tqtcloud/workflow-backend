package impl

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/infraboard/mcube/exception"
	"github.com/tqtcloud/workflow-backend/apps/task"
	"github.com/tqtcloud/workflow-backend/conf"
	"strings"
	"time"
)

// templateDetermine 后端模板判定，确认是什么语言返回什么语言的xml模板
// 不同模板之间的结构体不一样，所以需要在此处进行更改
func templateDetermine(ins *task.Task, config string) ([]byte, error) {
	switch ins.Data.TemplateName {
	case "jobtemplate/job/deploy-template":
		return deployXmlProc(ins, config)
	case "jobtemplate/job/go-backend-template":
		return goXmlProc(ins, config)
	case "jobtemplate/job/java-backend-template":
		return javaXmlProc(ins, config)
	case "jobtemplate/job/nodejs-backend-template":
		return nil, nil
	default:
		return nil, fmt.Errorf("TemplateName  %s  does not exist ", ins.Data.TemplateName)
	}
}

// createJob 创建job逻辑实现
func createJob(ctx context.Context, ins *task.Task, conf *conf.Config) (*task.Task, error) {
	jenkins, err := envDecision(ctx, ins.Data.Env, conf)
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
	// 处理掉不需要的xml头部信息
	config = strings.TrimPrefix(config, `<?xml version="1.1" encoding="UTF-8" standalone="no"?>`)
	config = strings.TrimPrefix(config, `<?xml version='1.1' encoding='UTF-8'?>`)
	//s.log.Debug(config)

	xmlData, err := templateDetermine(ins, config)
	if err != nil {
		return nil, exception.NewInternalServerError("Job templateDetermine error, %s  JobName: %s", err, ins.Data.JobName)
	}

	job, err = jenkins.CreateJobInFolder(ctx, string(xmlData), ins.Data.JobName, ins.Data.Folder)
	if err != nil {
		//s.log.Errorf("jenkins CreateJobInFolder error：%s,   job名称：%s", err, ins.Data.JobName)
		return nil, exception.NewInternalServerError("Job CreateJobInFolder error, %s  JobName: %s", err, ins.Data.JobName)
	}
	//s.log.Printf("%s 环境 jenkins Job %s 创建成功,目录位与：%s", task.JenkinsEnv_DEV, job.GetName(), ins.Data.Folder)
	return ins, nil
}

// envDecision jenkins 环境变量判定连接
func envDecision(ctx context.Context, env task.JenkinsEnv, conf *conf.Config) (*gojenkins.Jenkins, error) {
	switch env {
	case task.JenkinsEnv_DEV:
		jenkins, err := task.ConnectJenkins(ctx, conf.Jenkins.DevEndpoints, conf.Jenkins.User, conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		return jenkins, nil
	case task.JenkinsEnv_TEST:
		jenkins, err := task.ConnectJenkins(ctx, conf.Jenkins.TestEndpoints, conf.Jenkins.User, conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		return jenkins, nil
	case task.JenkinsEnv_UAT:
		jenkins, err := task.ConnectJenkins(ctx, conf.Jenkins.UatEndpoints, conf.Jenkins.User, conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		return jenkins, nil
	case task.JenkinsEnv_LPT:
		jenkins, err := task.ConnectJenkins(ctx, conf.Jenkins.LptEndpoints, conf.Jenkins.User, conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		return jenkins, nil
	case task.JenkinsEnv_PROD:
		jenkins, err := task.ConnectJenkins(ctx, conf.Jenkins.ProdEndpoints, conf.Jenkins.User, conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		return jenkins, nil
	default:
		return nil, fmt.Errorf("环境错误: %s; 您的：%s", "dev/test/uat/lpt/prod", env)
	}
}

// jobNameJoin 拼接逻辑判断，如果是jenkins 根下的目录直接返回，如果是在文件夹下的拼接后返回
func jobNameJoin(req *task.DescribeTaskRequest) string {
	var jobName string
	if req.Folder == "root" {
		jobName = req.Jobname
	} else {
		jobName = req.Folder + "/" + "job" + "/" + req.Jobname
	}
	return jobName
}

// describeJob 查看job信息
func describeJob(ctx context.Context, req *task.DescribeTaskRequest, jenkins *gojenkins.Jenkins) (*task.Task, error) {
	return getJobConfig(ctx, req, jenkins)
}

func getJobConfig(ctx context.Context, req *task.DescribeTaskRequest, jenkins *gojenkins.Jenkins) (*task.Task, error) {
	//ins := task.NewDescribeTaskRequest(req)
	ins := task.NewDefaultTask()

	// 在此处拼接job名称  例如：test2/job/apijob0927/
	jobName := jobNameJoin(req)

	fmt.Printf("请求job名称：%s \n", jobName)
	job, err := jenkins.GetJob(ctx, jobName)
	if err != nil {
		return nil, exception.NewInternalServerError("validate create func getJobConfig error, %s", err)
	}
	config, err := job.GetConfig(ctx)
	if err != nil {
		return nil, exception.NewInternalServerError("job  GetConfig error, %s", err)
	}
	config = strings.TrimPrefix(config, `<?xml version="1.1" encoding="UTF-8" standalone="no"?>`)
	config = strings.TrimPrefix(config, `<?xml version='1.1' encoding='UTF-8'?>`)
	data := new(Project)
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, exception.NewInternalServerError("Job config Unmarshal error, %s", err)
	}
	ins.Data.GitUrl = data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL
	ins.Data.Branch = data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.Branch
	ins.Data.Buildeshell = data.Builders.HudsonTasksShell.Command
	ins.Data.Description = data.Description
	// appname
	ins.Data.AppName = data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition.DefaultValue
	ins.Data.Folder = req.Folder
	ins.Data.Env = req.Env
	ins.Data.JobName = jobName
	ins.CreateAt = time.Now().UnixMicro()
	ins.Id = data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.Uuid
	return ins, nil
}

func delJob(ctx context.Context, req *task.DeleteTaskRequest, jenkins *gojenkins.Jenkins) error {
	jobName := jobNameJoin(task.NewDescribeTaskRequest(req.Env, req.Folder, req.Jobname))
	yes, err := jenkins.DeleteJob(ctx, jobName)
	if yes == true && err != nil {
		return nil
	}
	return err
}

func updateJob(ctx context.Context, ins *task.Task, conf *conf.Config) (*task.Task, error) {
	jenkins, err := envDecision(ctx, ins.Data.Env, conf)
	if err != nil {
		return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
	}
	// 前端定义好下拉列表在此处进行模板的选择，后端Java? go? nodejs?
	// jobtemplate/job/go-backend-template  go 模板名

	jobName := jobNameJoin(task.NewDescribeTaskRequest(ins.Data.Env.String(), ins.Data.Folder, ins.Data.JobName))
	job, err := jenkins.GetJob(ctx, jobName)
	if err != nil {
		return nil, exception.NewInternalServerError("Jenkins GetJob error, %s", err)
	}
	// job.UpdateConfig()
	// jenkins.UpdateJob()
	config, _ := job.GetConfig(ctx)

	config = strings.TrimPrefix(config, `<?xml version="1.1" encoding="UTF-8" standalone="no"?>`)
	config = strings.TrimPrefix(config, `<?xml version='1.1' encoding='UTF-8'?>`)

	//s.log.Debug(config)
	data := Project{}
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, exception.NewInternalServerError("Job config Unmarshal error, %s", err)
	}

	if ins.Data.GitUrl != "" {
		data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL = ins.Data.GitUrl
	}
	if ins.Data.Branch != "" {
		data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.Branch = ins.Data.Branch
	}
	if ins.Data.Buildeshell != "" {
		data.Builders.HudsonTasksShell.Command = ins.Data.Buildeshell
	}
	if ins.Data.Description != "" {
		data.Description = ins.Data.Description
	}
	if ins.Data.AppName != "" {
		// appname
		data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition.DefaultValue = ins.Data.AppName
	}

	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, exception.NewInternalServerError("Job config MarshalIndent error, %s", err)
	}

	if err = job.UpdateConfig(ctx, string(xmlData)); err != nil {
		//s.log.Errorf("jenkins CreateJobInFolder error：%s,   job名称：%s", err, ins.Data.JobName)
		return nil, exception.NewInternalServerError("Job UpdateConfig error, %s  JobName: %s", err, ins.Data.JobName)
	}

	//s.log.Printf("%s 环境 jenkins Job %s 创建成功,目录位与：%s", task.JenkinsEnv_DEV, job.GetName(), ins.Data.Folder)
	return ins, nil
}
