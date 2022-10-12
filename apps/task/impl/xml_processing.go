package impl

import (
	"encoding/xml"
	"fmt"
	"github.com/tqtcloud/workflow-backend/apps/task"
)
// go 模板函数处理
func goXmlProc(ins *task.Task,config string) ([]byte,error) {
	data := Project{}
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config Unmarshal error, %s ", err)
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
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}

// java 模板函数处理
func javaXmlProc(ins *task.Task,config string) ([]byte,error) {
	data := Maven2Moduleset{}
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config Unmarshal error, %s ", err)
	}

	data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL = ins.Data.GitUrl
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.Branch = ins.Data.Branch
	if ins.Data.Buildeshell != "" {
		data.Postbuilders.HudsonTasksShell.Command = ins.Data.Buildeshell
	}
	data.Description = ins.Data.Description

	//通过环境对docker name 添加尾缀
	imageName ,_ := imageTail(ins)
	// CONTAINER_NAME
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[1].DefaultValue = imageName
	// appname
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[2].DefaultValue = imageName
	// CODE_MODULE mvn 打包服务 需要使用原名
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[4].DefaultValue = ins.Data.AppName

	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}
// imageTail docker name 尾缀添加，为了相应后端的镜像构建：bigdata_server_dev
func imageTail(ins *task.Task) (string,error)  {
	switch ins.Data.Env {
	case task.JenkinsEnv_DEV:
		return ins.Data.AppName+"_dev",nil
	case task.JenkinsEnv_TEST:
		return ins.Data.AppName+"_sit",nil
	case task.JenkinsEnv_UAT:
		return ins.Data.AppName+"_uat",nil
	default:
		return "",fmt.Errorf("imageTail  %s  does not exist ",ins.Data.Env)
	}
}