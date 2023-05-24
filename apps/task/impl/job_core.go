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
	case "jobtemplate/job/argocd-deploy-template":
		return deployXmlProc(ins, config)
	case "jobtemplate/job/go-backend-template-build": // go 构建制品
		return goBuildXmlProc(ins, config)
	case "jobtemplate/job/go-backend-template-build-ssh": // go 构建 ssh部署
		return goBuildSShXmlProc(ins, config)
	case "jobtemplate/job/go-backend-template-build-argocd":
		return goBuildDeployXmlProc(ins, config)
	case "jobtemplate/job/java-backend-template-build-argocd":
		return javaBuildDeployXmlProc(ins, config)
	case "jobtemplate/job/java-backend-template-build":
		return javaBuildXmlProc(ins, config)
	case "jobtemplate/job/nodejs-template-build-deploy": // 前端打包构建一起
		return nodeBuildDeployXmlProc(ins, config)
	case "jobtemplate/job/nodejs-template-build": // 前端打包
		return nodeBuildXmlProc(ins, config)
	//case "jobtemplate/job/nodejs-template-deploy": // 前端部署
	//	return nil, nil
	default:
		return nil, fmt.Errorf("TemplateName  %s  does not exist ", ins.Data.TemplateName)
	}
}

//func templateDescribe(ins *task.Task, config string) ([]byte, error) {
//	switch ins.Data.TemplateName {
//	case "jobtemplate/job/deploy-template":
//		return deployXmlUnmarshal(ins, config)
//	case "jobtemplate/job/go-backend-template":
//		return goXmlUnmarshal(ins, config)
//	case "jobtemplate/job/java-backend-template":
//		return javaXmlUnmarshal(ins, config)
//	case "jobtemplate/job/nodejs-backend-template":
//		return nil, nil
//	default:
//		return nil, fmt.Errorf("TemplateName  %s  does not exist ", ins.Data.TemplateName)
//	}
//}

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
	//case task.JenkinsEnv_LPT:
	//	jenkins, err := task.ConnectJenkins(ctx, conf.Jenkins.LptEndpoints, conf.Jenkins.User, conf.Jenkins.Password)
	//	if err != nil {
	//		return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
	//	}
	//	return jenkins, nil
	case task.JenkinsEnv_PROD:
		jenkins, err := task.ConnectJenkins(ctx, conf.Jenkins.ProdEndpoints, conf.Jenkins.User, conf.Jenkins.Password)
		if err != nil {
			return nil, exception.NewInternalServerError("Connect Jenkins error, %s", err)
		}
		return jenkins, nil
	default:
		return nil, fmt.Errorf("环境错误: %s; 您的：%s", "dev/test/uat/prod", env)
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
	// 进行config 反序列化
	ins, err = classUnmarshal(req, ins, config)
	if err != nil {
		return nil, exception.NewInternalServerError("job  GetConfig error, %s", err)
	}

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
	data := GeneralStruct{}
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
		data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[0].DefaultValue = ins.Data.AppName
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

// classUnmarshal get config 不同job类型的分离转换
func classUnmarshal(req *task.DescribeTaskRequest, ins *task.Task, config string) (*task.Task, error) {
	generalData := new(GeneralStruct)
	goData := new(GoStruct)
	javaData := new(Maven2Struct)
	deployData := new(DeployStruct)
	if err := xml.Unmarshal([]byte(config), generalData); err == nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		data := generalUnmarshal(req, ins, generalData)
		//fmt.Printf("jenkins xml 反序列化成功：%s,job名称：%s 属于类型：%T \n", err, ins.Data.JobName,generalData)
		return data, nil
	} else if err := xml.Unmarshal([]byte(config), javaData); err == nil {
		data := javaUnmarshal(req, ins, javaData)
		return data, nil
	} else if err := xml.Unmarshal([]byte(config), goData); err == nil {
		data := generalUnmarshal(req, ins, goData)
		return data, nil

	} else if err := xml.Unmarshal([]byte(config), deployData); err == nil {
		//data = data.(DeployStruct)
		data := generalUnmarshal(req, ins, deployData)
		return data, nil
	} else {
		return nil, exception.NewInternalServerError("Job config classUnmarshal error, %s", err)
	}
}

// generalUnmarshal 通用xml 数据进行处理转换为结构体进行http返回
func generalUnmarshal(req *task.DescribeTaskRequest, ins *task.Task, data any) *task.Task {
	fmt.Println(data.(*GeneralStruct).Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL)
	ins.Data.GitUrl = data.(*GeneralStruct).Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL
	ins.Data.Branch = data.(*GeneralStruct).Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.Branch
	ins.Data.Buildeshell = data.(*GeneralStruct).Builders.HudsonTasksShell.Command
	ins.Data.Description = data.(*GeneralStruct).Description
	ins.Data.AppName = data.(*GeneralStruct).Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[0].DefaultValue
	ins.Data.Folder = req.Folder
	ins.Data.Env = req.Env
	ins.Data.JobName = jobNameJoin(req)
	ins.CreateAt = time.Now().UnixMicro()
	// 增加2个nodejs字段
	ins.Data.Sshshell = data.(*GeneralStruct).Publishers.JenkinsPluginsPublishOverSshBapSshPublisherPlugin.Delegate.Publishers.JenkinsPluginsPublishOverSshBapSshPublisher.Transfers.JenkinsPluginsPublishOverSshBapSshTransfer.ExecCommand
	ins.Data.Sshnode = data.(*GeneralStruct).Publishers.JenkinsPluginsPublishOverSshBapSshPublisherPlugin.Delegate.Publishers.JenkinsPluginsPublishOverSshBapSshPublisher.ConfigName
	// jdk node版本信息
	ins.Data.Buildenv = data.(*GeneralStruct).BuildWrappers.JenkinsPluginsNodejsNodeJSBuildWrapper.NodeJSInstallationName
	ins.Id = data.(*GeneralStruct).Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.Uuid
	return ins
}

// javaUnmarshal Java xml 数据进行处理转换为结构体进行http返回
func javaUnmarshal(req *task.DescribeTaskRequest, ins *task.Task, data any) *task.Task {
	ins.Data.GitUrl = data.(*Maven2Struct).Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL
	ins.Data.Branch = data.(*Maven2Struct).Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.Branch
	ins.Data.Buildeshell = data.(*Maven2Struct).Postbuilders.HudsonTasksShell.Command
	ins.Data.Description = data.(*Maven2Struct).Description
	ins.Data.AppName = data.(*Maven2Struct).Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[4].DefaultValue
	ins.Data.Buildenv = data.(*Maven2Struct).Jdk
	ins.Data.Folder = req.Folder
	ins.Data.Env = req.Env
	ins.Data.JobName = jobNameJoin(req)
	ins.CreateAt = time.Now().UnixMicro()
	return ins
}

// appENV  根据jenkins环境确定argocd环境变量值
func appENV(env task.JenkinsEnv) string {
	switch env {
	case task.JenkinsEnv_DEV:
		return "dev"
	case task.JenkinsEnv_TEST:
		return "sit"
	case task.JenkinsEnv_UAT:
		return "uat"
	case task.JenkinsEnv_PROD:
		return "prod"
	default:
		return "请输入：dev,sit,uat"
	}
}

// imageTail docker name 尾缀添加，为了相应后端的镜像构建：bigdata_server_dev
func imageTail(ins *task.Task) (string, error) {
	switch ins.Data.Env {
	case task.JenkinsEnv_DEV:
		return "_dev", nil
	case task.JenkinsEnv_TEST:
		return "_sit", nil
	default:
		return "", fmt.Errorf("imageTail  %s  does not exist ", ins.Data.Env)
	}
}
