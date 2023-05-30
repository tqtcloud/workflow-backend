package impl

import (
	"encoding/xml"
	"fmt"
	"github.com/tqtcloud/workflow-backend/apps/task"
	"strings"
)

// argocd  部署模板
func deployXmlProc(ins *task.Task, config string) ([]byte, error) {
	data := DeployStruct{}
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config Unmarshal error, %s ", err)
	}

	if ins.Data.Buildeshell != "" {
		data.Builders.HudsonTasksShell.Command = ins.Data.Buildeshell
	}
	data.Description = ins.Data.Description
	// APP_NAME
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[0].DefaultValue = ins.Data.AppName
	// APP_TYPE
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[1].DefaultValue = ins.Data.AppType
	// APP_ENV
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[2].DefaultValue = appENV(ins.Data.Env)

	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}

// go 构建 argocd 部署模板函数处理
func goBuildDeployXmlProc(ins *task.Task, config string) ([]byte, error) {
	data := GoBuildDeployStruct{}
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

	imageName, _ := imageTail(ins)
	// APP_NAME
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[0].DefaultValue = ins.Data.AppName + imageName
	// APP_TYPE
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[1].DefaultValue = ins.Data.AppType
	// APP_ENV
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[2].DefaultValue = appENV(ins.Data.Env)

	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}

// go 构建 ssh 部署模板函数处理
func goBuildSShXmlProc(ins *task.Task, config string) ([]byte, error) {
	data := GeneralStruct{}
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
	imageName, _ := imageTail(ins)
	// APP_NAME
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[0].DefaultValue = ins.Data.AppName + imageName
	// SERVER_NAME
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[1].DefaultValue = ins.Data.AppName

	// ssh 到那台服务器使用
	if ins.Data.Sshnode != "" {
		data.Publishers.JenkinsPluginsPublishOverSshBapSshPublisherPlugin.Delegate.Publishers.JenkinsPluginsPublishOverSshBapSshPublisher.ConfigName = ins.Data.Sshnode
	}
	// 远程主机执行命令
	if ins.Data.Sshshell != "" {
		data.Publishers.JenkinsPluginsPublishOverSshBapSshPublisherPlugin.Delegate.Publishers.JenkinsPluginsPublishOverSshBapSshPublisher.Transfers.JenkinsPluginsPublishOverSshBapSshTransfer.ExecCommand = ins.Data.Sshshell
	}

	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}

// goBuildXmlProc go 构建模板
func goBuildXmlProc(ins *task.Task, config string) ([]byte, error) {
	// 这里也使用相同模板是因为 build job 没有argocd的触发命令其他相同故使用
	data := GoBuildDeployStruct{}

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

	imageName, _ := imageTail(ins)
	// APP_NAME
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[0].DefaultValue = ins.Data.AppName + imageName
	// SERVER_NAME
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[1].DefaultValue = ins.Data.AppName

	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}

// java 打包 部署 模板函数处理
func javaBuildDeployXmlProc(ins *task.Task, config string) ([]byte, error) {
	data := Maven2Struct{}
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config Unmarshal error, %s ", err)
	}

	data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL = ins.Data.GitUrl
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.DefaultValue = ins.Data.Branch
	if ins.Data.Buildeshell != "" {
		data.Postbuilders.HudsonTasksShell.Command = ins.Data.Buildeshell
	}
	data.Description = ins.Data.Description

	//通过环境对docker name 添加尾缀
	imageName, _ := imageTail(ins)
	// APP_NAME
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[0].DefaultValue = ins.Data.AppName + imageName
	// CODE_MODULE mvn 打包服务 argocd 服务名 需要使用原名
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[1].DefaultValue = ins.Data.AppName
	// APP_TYPE namespace名
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[2].DefaultValue = ins.Data.AppType
	// APP_ENV
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[3].DefaultValue = appENV(ins.Data.Env)

	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}

// java 打包  模板函数处理
func javaBuildXmlProc(ins *task.Task, config string) ([]byte, error) {
	data := Maven2Struct{}
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config Unmarshal error, %s ", err)
	}

	data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL = ins.Data.GitUrl
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.DefaultValue = ins.Data.Branch
	if ins.Data.Buildeshell != "" {
		data.Postbuilders.HudsonTasksShell.Command = ins.Data.Buildeshell
	}
	data.Description = ins.Data.Description

	//通过环境对docker name 添加尾缀
	imageName, _ := imageTail(ins)
	// CONTAINER_NAME
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[0].DefaultValue = strings.ReplaceAll(ins.Data.AppName, "-", "_") + imageName
	// APP_NAME
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[1].DefaultValue = ins.Data.AppName + imageName
	// CODE_MODULE mvn 打包服务 需要使用原名
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[2].DefaultValue = ins.Data.AppName

	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}

// node 打包部署模板函数处理
func nodeBuildDeployXmlProc(ins *task.Task, config string) ([]byte, error) {
	data := NodejsStruct{}
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config nodeBuildDeployXmlProc error, %s ", err)
	}

	data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL = ins.Data.GitUrl
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition.Branch = ins.Data.Branch
	if ins.Data.Buildeshell != "" {
		data.Builders.HudsonTasksShell.Command = ins.Data.Buildeshell
	}
	data.Description = ins.Data.Description

	// ssh 到那台服务器使用
	if ins.Data.Sshnode != "" {
		data.Publishers.JenkinsPluginsPublishOverSshBapSshPublisherPlugin.Delegate.Publishers.JenkinsPluginsPublishOverSshBapSshPublisher.ConfigName = ins.Data.Sshnode
	}
	// 远程主机执行命令
	if ins.Data.Sshshell != "" {
		data.Publishers.JenkinsPluginsPublishOverSshBapSshPublisherPlugin.Delegate.Publishers.JenkinsPluginsPublishOverSshBapSshPublisher.Transfers.JenkinsPluginsPublishOverSshBapSshTransfer.ExecCommand = ins.Data.Sshshell
	}
	// 更换nodejs的相关版本信息
	if ins.Data.Buildenv != "" {
		data.BuildWrappers.JenkinsPluginsNodejsNodeJSBuildWrapper.NodeJSInstallationName = ins.Data.Buildenv
	}
	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}

// node 前端单独构建
func nodeBuildXmlProc(ins *task.Task, config string) ([]byte, error) {
	data := NodejsBuildStruct{}
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config nodeBuildDeployXmlProc error, %s ", err)
	}

	data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL = ins.Data.GitUrl
	data.Scm.Branches.HudsonPluginsGitBranchSpec.Name = ins.Data.Branch
	if ins.Data.Buildeshell != "" {
		data.Builders.HudsonTasksShell.Command = ins.Data.Buildeshell
	}
	data.Description = ins.Data.Description
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition.DefaultValue = ins.Data.AppName
	//// ssh 到那台服务器使用
	//if ins.Data.Sshnode != "" {
	//	data.Publishers.JenkinsPluginsPublishOverSshBapSshPublisherPlugin.Delegate.Publishers.JenkinsPluginsPublishOverSshBapSshPublisher.ConfigName = ins.Data.Sshnode
	//}
	//// 远程主机执行命令
	//if ins.Data.Sshshell != "" {
	//	data.Publishers.JenkinsPluginsPublishOverSshBapSshPublisherPlugin.Delegate.Publishers.JenkinsPluginsPublishOverSshBapSshPublisher.Transfers.JenkinsPluginsPublishOverSshBapSshTransfer.ExecCommand = ins.Data.Sshshell
	//}
	// 更换nodejs的相关版本信息
	if ins.Data.Buildenv != "" {
		data.BuildWrappers.JenkinsPluginsNodejsNodeJSBuildWrapper.NodeJSInstallationName = ins.Data.Buildenv
	}
	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}

//nodejs-template-build-nginx-deploy 模板函数处理
func nodeBuildNginxDeployXmlProc(ins *task.Task, config string) ([]byte, error) {
	data := NodeBuildNginxDeployStruct{}
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config nodeBuildDeployXmlProc error, %s ", err)
	}
	data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL = ins.Data.GitUrl
	data.Scm.Branches.HudsonPluginsGitBranchSpec.Name = ins.Data.Branch
	if ins.Data.Buildeshell != "" {
		data.Builders.HudsonTasksShell.Command = ins.Data.Buildeshell
	}
	data.Description = ins.Data.Description
	// ssh 到那台服务器使用
	if ins.Data.Sshnode != "" {
		data.Publishers.JenkinsPluginsPublishOverSshBapSshPublisherPlugin.Delegate.Publishers.JenkinsPluginsPublishOverSshBapSshPublisher.ConfigName = ins.Data.Sshnode
	}
	//// 远程主机执行命令
	//if ins.Data.Sshshell != "" {
	//	data.Publishers.JenkinsPluginsPublishOverSshBapSshPublisherPlugin.Delegate.Publishers.JenkinsPluginsPublishOverSshBapSshPublisher.Transfers.JenkinsPluginsPublishOverSshBapSshTransfer.ExecCommand = ins.Data.Sshshell
	//}
	//

	// 更换nodejs的相关版本信息
	if ins.Data.Buildenv != "" {
		data.BuildWrappers.JenkinsPluginsNodejsNodeJSBuildWrapper.NodeJSInstallationName = ins.Data.Buildenv
	}
	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}

//nodejs-template-nginx-deploy 前端ssh分发模板处理函数
func nodeNginxDeployXmlProc(ins *task.Task, config string) ([]byte, error) {
	data := NodeNginxDeploySshStruct{}
	if err := xml.Unmarshal([]byte(config), &data); err != nil {
		//s.log.Errorf("jenkins xml 反序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config nodeBuildDeployXmlProc error, %s ", err)
	}
	data.Scm.UserRemoteConfigs.HudsonPluginsGitUserRemoteConfig.URL = ins.Data.GitUrl
	data.Scm.Branches.HudsonPluginsGitBranchSpec.Name = ins.Data.Branch
	if ins.Data.Buildeshell != "" {
		data.Builders.HudsonTasksShell.Command = ins.Data.Buildeshell
	}
	data.Description = ins.Data.Description
	//CODE_MODULE
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[0].DefaultValue = ins.Data.AppName
	//AppName
	data.Properties.HudsonModelParametersDefinitionProperty.ParameterDefinitions.HudsonModelStringParameterDefinition[1].DefaultValue = ins.Data.AppName
	if ins.Data.Buildenv != "" {
		data.BuildWrappers.JenkinsPluginsNodejsNodeJSBuildWrapper.NodeJSInstallationName = ins.Data.Buildenv
	}
	xmlData, err := xml.MarshalIndent(&data, " ", " ")
	if err != nil {
		//s.log.Errorf("jenkins xml 序列化错误：%s,job名称：%s", err, ins.Data.JobName)
		return nil, fmt.Errorf("Job config MarshalIndent error, %s ", err)
	}
	return xmlData, nil
}
