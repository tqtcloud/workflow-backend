package impl

import "encoding/xml"

// 通用 数据模板
type GeneralStruct struct {
	XMLName          xml.Name `xml:"project"`
	Text             string   `xml:",chardata"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		OrgJenkinsciPluginsGitlablogoGitlabLogoProperty struct {
			Text           string `xml:",chardata"`
			Plugin         string `xml:"plugin,attr"`
			RepositoryName string `xml:"repositoryName"`
		} `xml:"org.jenkinsci.plugins.gitlablogo.GitlabLogoProperty"`
		ComSynopsysArcJenkinsciPluginsJobrestrictionsJobsJobRestrictionProperty struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"com.synopsys.arc.jenkinsci.plugins.jobrestrictions.jobs.JobRestrictionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                                            string `xml:",chardata"`
				NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
					Text               string `xml:",chardata"`
					Plugin             string `xml:"plugin,attr"`
					Name               string `xml:"name"`
					Description        string `xml:"description"`
					Uuid               string `xml:"uuid"`
					Type               string `xml:"type"`
					Branch             string `xml:"branch"`
					TagFilter          string `xml:"tagFilter"`
					BranchFilter       string `xml:"branchFilter"`
					SortMode           string `xml:"sortMode"`
					DefaultValue       string `xml:"defaultValue"`
					SelectedValue      string `xml:"selectedValue"`
					QuickFilterEnabled string `xml:"quickFilterEnabled"`
					ListSize           string `xml:"listSize"`
				} `xml:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
				//HudsonModelStringParameterDefinition struct {
				//	Text         string `xml:",chardata"`
				//	Name         string `xml:"name"`
				//	Description  string `xml:"description"`
				//	DefaultValue string `xml:"defaultValue"`
				//	Trim         string `xml:"trim"`
				//} `xml:"hudson.model.StringParameterDefinition"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text              string `xml:",chardata"`
		Class             string `xml:"class,attr"`
		Plugin            string `xml:"plugin,attr"`
		ConfigVersion     string `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text          string `xml:",chardata"`
				URL           string `xml:"url"`
				CredentialsId string `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg                      struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions string `xml:"extensions"`
	} `xml:"scm"`
	AssignedNode                     string `xml:"assignedNode"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         struct {
		Text                                        string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsGitLabPushTrigger struct {
			Text                           string `xml:",chardata"`
			Plugin                         string `xml:"plugin,attr"`
			Spec                           string `xml:"spec"`
			TriggerOnPush                  string `xml:"triggerOnPush"`
			TriggerOnMergeRequest          string `xml:"triggerOnMergeRequest"`
			TriggerOnPipelineEvent         string `xml:"triggerOnPipelineEvent"`
			TriggerOnAcceptedMergeRequest  string `xml:"triggerOnAcceptedMergeRequest"`
			TriggerOnClosedMergeRequest    string `xml:"triggerOnClosedMergeRequest"`
			TriggerOnApprovedMergeRequest  string `xml:"triggerOnApprovedMergeRequest"`
			TriggerOpenMergeRequestOnPush  string `xml:"triggerOpenMergeRequestOnPush"`
			TriggerOnNoteRequest           string `xml:"triggerOnNoteRequest"`
			NoteRegex                      string `xml:"noteRegex"`
			CiSkip                         string `xml:"ciSkip"`
			SkipWorkInProgressMergeRequest string `xml:"skipWorkInProgressMergeRequest"`
			SetBuildDescription            string `xml:"setBuildDescription"`
			BranchFilterType               string `xml:"branchFilterType"`
			IncludeBranchesSpec            string `xml:"includeBranchesSpec"`
			ExcludeBranchesSpec            string `xml:"excludeBranchesSpec"`
			SourceBranchRegex              string `xml:"sourceBranchRegex"`
			TargetBranchRegex              string `xml:"targetBranchRegex"`
			SecretToken                    string `xml:"secretToken"`
			PendingBuildName               string `xml:"pendingBuildName"`
			CancelPendingBuildsOnUpdate    string `xml:"cancelPendingBuildsOnUpdate"`
		} `xml:"com.dabsquared.gitlabjenkins.GitLabPushTrigger"`
	} `xml:"triggers"`
	ConcurrentBuild string `xml:"concurrentBuild"`
	Builders        struct {
		Text             string `xml:",chardata"`
		HudsonTasksShell struct {
			Text    string `xml:",chardata"`
			Command string `xml:"command"`
		} `xml:"hudson.tasks.Shell"`
	} `xml:"builders"`
	//Publishers    string `xml:"publishers"`
	Publishers struct {
		Text                                              string `xml:",chardata"`
		JenkinsPluginsPublishOverSshBapSshPublisherPlugin struct {
			Text          string `xml:",chardata"`
			Plugin        string `xml:"plugin,attr"`
			ConsolePrefix string `xml:"consolePrefix"`
			Delegate      struct {
				Text       string `xml:",chardata"`
				Plugin     string `xml:"plugin,attr"`
				Publishers struct {
					Text                                        string `xml:",chardata"`
					JenkinsPluginsPublishOverSshBapSshPublisher struct {
						Text       string `xml:",chardata"`
						Plugin     string `xml:"plugin,attr"`
						ConfigName string `xml:"configName"`
						Verbose    string `xml:"verbose"`
						Transfers  struct {
							Text                                       string `xml:",chardata"`
							JenkinsPluginsPublishOverSshBapSshTransfer struct {
								Text               string `xml:",chardata"`
								RemoteDirectory    string `xml:"remoteDirectory"`
								SourceFiles        string `xml:"sourceFiles"`
								Excludes           string `xml:"excludes"`
								RemovePrefix       string `xml:"removePrefix"`
								RemoteDirectorySDF string `xml:"remoteDirectorySDF"`
								Flatten            string `xml:"flatten"`
								CleanRemote        string `xml:"cleanRemote"`
								NoDefaultExcludes  string `xml:"noDefaultExcludes"`
								MakeEmptyDirs      string `xml:"makeEmptyDirs"`
								PatternSeparator   string `xml:"patternSeparator"`
								ExecCommand        string `xml:"execCommand"`
								ExecTimeout        string `xml:"execTimeout"`
								UsePty             string `xml:"usePty"`
								UseAgentForwarding string `xml:"useAgentForwarding"`
							} `xml:"jenkins.plugins.publish__over__ssh.BapSshTransfer"`
						} `xml:"transfers"`
						UseWorkspaceInPromotion string `xml:"useWorkspaceInPromotion"`
						UsePromotionTimestamp   string `xml:"usePromotionTimestamp"`
					} `xml:"jenkins.plugins.publish__over__ssh.BapSshPublisher"`
				} `xml:"publishers"`
				ContinueOnError         string `xml:"continueOnError"`
				FailOnError             string `xml:"failOnError"`
				AlwaysPublishFromMaster string `xml:"alwaysPublishFromMaster"`
				HostConfigurationAccess struct {
					Text      string `xml:",chardata"`
					Class     string `xml:"class,attr"`
					Reference string `xml:"reference,attr"`
				} `xml:"hostConfigurationAccess"`
			} `xml:"delegate"`
		} `xml:"jenkins.plugins.publish__over__ssh.BapSshPublisherPlugin"`
	} `xml:"publishers"`

	BuildWrappers struct {
		Text                                            string `xml:",chardata"`
		HudsonPluginsTimestamperTimestamperBuildWrapper struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"hudson.plugins.timestamper.TimestamperBuildWrapper"`
		HudsonPluginsAnsicolorAnsiColorBuildWrapper struct {
			Text         string `xml:",chardata"`
			Plugin       string `xml:"plugin,attr"`
			ColorMapName string `xml:"colorMapName"`
		} `xml:"hudson.plugins.ansicolor.AnsiColorBuildWrapper"`
		JenkinsPluginsNodejsNodeJSBuildWrapper struct {
			Text                   string `xml:",chardata"`
			Plugin                 string `xml:"plugin,attr"`
			NodeJSInstallationName string `xml:"nodeJSInstallationName"`
			CacheLocationStrategy  struct {
				Text  string `xml:",chardata"`
				Class string `xml:"class,attr"`
			} `xml:"cacheLocationStrategy"`
		} `xml:"jenkins.plugins.nodejs.NodeJSBuildWrapper"`
	} `xml:"buildWrappers"`
}

// GO 语言项目模板
type GoStruct struct {
	Project struct {
		KeepDependencies string `json:"keepDependencies"`
		Properties       struct {
			ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
				Plugin string `json:"-plugin"`
			} `json:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
			OrgJenkinsciPluginsGitlablogoGitlabLogoProperty struct {
				Plugin string `json:"-plugin"`
			} `json:"org.jenkinsci.plugins.gitlablogo.GitlabLogoProperty"`
			ComSynopsysArcJenkinsciPluginsJobrestrictionsJobsJobRestrictionProperty struct {
				Plugin string `json:"-plugin"`
			} `json:"com.synopsys.arc.jenkinsci.plugins.jobrestrictions.jobs.JobRestrictionProperty"`
			JenkinsModelBuildDiscarderProperty struct {
				Strategy struct {
					Class              string `json:"-class"`
					DaysToKeep         string `json:"daysToKeep"`
					NumToKeep          string `json:"numToKeep"`
					ArtifactDaysToKeep string `json:"artifactDaysToKeep"`
					ArtifactNumToKeep  string `json:"artifactNumToKeep"`
				} `json:"strategy"`
			} `json:"jenkins.model.BuildDiscarderProperty"`
			HudsonModelParametersDefinitionProperty struct {
				ParameterDefinitions struct {
					NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
						Plugin             string `json:"-plugin"`
						Name               string `json:"name"`
						UUID               string `json:"uuid"`
						Type               string `json:"type"`
						Branch             string `json:"branch"`
						TagFilter          string `json:"tagFilter"`
						BranchFilter       string `json:"branchFilter"`
						SortMode           string `json:"sortMode"`
						DefaultValue       string `json:"defaultValue"`
						SelectedValue      string `json:"selectedValue"`
						QuickFilterEnabled string `json:"quickFilterEnabled"`
						ListSize           string `json:"listSize"`
					} `json:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
					HudsonModelStringParameterDefinition struct {
						Name         string `json:"name"`
						DefaultValue string `json:"defaultValue"`
						Trim         string `json:"trim"`
					} `json:"hudson.model.StringParameterDefinition"`
				} `json:"parameterDefinitions"`
			} `json:"hudson.model.ParametersDefinitionProperty"`
		} `json:"properties"`
		Scm struct {
			Class             string `json:"-class"`
			Plugin            string `json:"-plugin"`
			ConfigVersion     string `json:"configVersion"`
			UserRemoteConfigs struct {
				HudsonPluginsGitUserRemoteConfig struct {
					URL           string `json:"url"`
					CredentialsID string `json:"credentialsId"`
				} `json:"hudson.plugins.git.UserRemoteConfig"`
			} `json:"userRemoteConfigs"`
			Branches struct {
				HudsonPluginsGitBranchSpec struct {
					Name string `json:"name"`
				} `json:"hudson.plugins.git.BranchSpec"`
			} `json:"branches"`
			DoGenerateSubmoduleConfigurations string `json:"doGenerateSubmoduleConfigurations"`
			SubmoduleCfg                      struct {
				Class string `json:"-class"`
			} `json:"submoduleCfg"`
		} `json:"scm"`
		AssignedNode                     string `json:"assignedNode"`
		CanRoam                          string `json:"canRoam"`
		Disabled                         string `json:"disabled"`
		BlockBuildWhenDownstreamBuilding string `json:"blockBuildWhenDownstreamBuilding"`
		BlockBuildWhenUpstreamBuilding   string `json:"blockBuildWhenUpstreamBuilding"`
		Jdk                              string `json:"jdk"`
		ConcurrentBuild                  string `json:"concurrentBuild"`
		Builders                         struct {
			HudsonTasksShell []struct {
				Command string `json:"command"`
			} `json:"hudson.tasks.Shell"`
		} `json:"builders"`
		Publishers struct {
			JenkinsPluginsPublishOverSSHBapSSHPublisherPlugin struct {
				Plugin        string `json:"-plugin"`
				ConsolePrefix string `json:"consolePrefix"`
				Delegate      struct {
					Plugin     string `json:"-plugin"`
					Publishers struct {
						JenkinsPluginsPublishOverSSHBapSSHPublisher struct {
							Plugin     string `json:"-plugin"`
							ConfigName string `json:"configName"`
							Verbose    string `json:"verbose"`
							Transfers  struct {
								JenkinsPluginsPublishOverSSHBapSSHTransfer struct {
									SourceFiles        string `json:"sourceFiles"`
									RemoteDirectorySDF string `json:"remoteDirectorySDF"`
									Flatten            string `json:"flatten"`
									CleanRemote        string `json:"cleanRemote"`
									NoDefaultExcludes  string `json:"noDefaultExcludes"`
									MakeEmptyDirs      string `json:"makeEmptyDirs"`
									PatternSeparator   string `json:"patternSeparator"`
									ExecCommand        string `json:"execCommand"`
									ExecTimeout        string `json:"execTimeout"`
									UsePty             string `json:"usePty"`
									UseAgentForwarding string `json:"useAgentForwarding"`
								} `json:"jenkins.plugins.publish__over__ssh.BapSshTransfer"`
							} `json:"transfers"`
							UseWorkspaceInPromotion string `json:"useWorkspaceInPromotion"`
							UsePromotionTimestamp   string `json:"usePromotionTimestamp"`
						} `json:"jenkins.plugins.publish__over__ssh.BapSshPublisher"`
					} `json:"publishers"`
					ContinueOnError         string `json:"continueOnError"`
					FailOnError             string `json:"failOnError"`
					AlwaysPublishFromMaster string `json:"alwaysPublishFromMaster"`
					HostConfigurationAccess struct {
						Class     string `json:"-class"`
						Reference string `json:"-reference"`
					} `json:"hostConfigurationAccess"`
				} `json:"delegate"`
			} `json:"jenkins.plugins.publish__over__ssh.BapSshPublisherPlugin"`
		} `json:"publishers"`
		BuildWrappers struct {
			HudsonPluginsTimestamperTimestamperBuildWrapper struct {
				Plugin string `json:"-plugin"`
			} `json:"hudson.plugins.timestamper.TimestamperBuildWrapper"`
			JenkinsPluginsNodejsNodeJSBuildWrapper struct {
				Plugin                 string `json:"-plugin"`
				NodeJSInstallationName string `json:"nodeJSInstallationName"`
				CacheLocationStrategy  struct {
					Class string `json:"-class"`
				} `json:"cacheLocationStrategy"`
			} `json:"jenkins.plugins.nodejs.NodeJSBuildWrapper"`
		} `json:"buildWrappers"`
	} `json:"project"`
}

type GoBuildDeployStruct struct {
	XMLName          xml.Name `xml:"project"`
	Text             string   `xml:",chardata"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		OrgJenkinsciPluginsGitlablogoGitlabLogoProperty struct {
			Text           string `xml:",chardata"`
			Plugin         string `xml:"plugin,attr"`
			RepositoryName string `xml:"repositoryName"`
		} `xml:"org.jenkinsci.plugins.gitlablogo.GitlabLogoProperty"`
		ComSynopsysArcJenkinsciPluginsJobrestrictionsJobsJobRestrictionProperty struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"com.synopsys.arc.jenkinsci.plugins.jobrestrictions.jobs.JobRestrictionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                                            string `xml:",chardata"`
				NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
					Text               string `xml:",chardata"`
					Plugin             string `xml:"plugin,attr"`
					Name               string `xml:"name"`
					Description        string `xml:"description"`
					Uuid               string `xml:"uuid"`
					Type               string `xml:"type"`
					Branch             string `xml:"branch"`
					TagFilter          string `xml:"tagFilter"`
					BranchFilter       string `xml:"branchFilter"`
					SortMode           string `xml:"sortMode"`
					DefaultValue       string `xml:"defaultValue"`
					SelectedValue      string `xml:"selectedValue"`
					QuickFilterEnabled string `xml:"quickFilterEnabled"`
					ListSize           string `xml:"listSize"`
				} `xml:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text              string `xml:",chardata"`
		Class             string `xml:"class,attr"`
		Plugin            string `xml:"plugin,attr"`
		ConfigVersion     string `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text          string `xml:",chardata"`
				URL           string `xml:"url"`
				CredentialsId string `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg                      struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions string `xml:"extensions"`
	} `xml:"scm"`
	AssignedNode                     string `xml:"assignedNode"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         struct {
		Text                                        string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsGitLabPushTrigger struct {
			Text                           string `xml:",chardata"`
			Plugin                         string `xml:"plugin,attr"`
			Spec                           string `xml:"spec"`
			TriggerOnPush                  string `xml:"triggerOnPush"`
			TriggerOnMergeRequest          string `xml:"triggerOnMergeRequest"`
			TriggerOnPipelineEvent         string `xml:"triggerOnPipelineEvent"`
			TriggerOnAcceptedMergeRequest  string `xml:"triggerOnAcceptedMergeRequest"`
			TriggerOnClosedMergeRequest    string `xml:"triggerOnClosedMergeRequest"`
			TriggerOnApprovedMergeRequest  string `xml:"triggerOnApprovedMergeRequest"`
			TriggerOpenMergeRequestOnPush  string `xml:"triggerOpenMergeRequestOnPush"`
			TriggerOnNoteRequest           string `xml:"triggerOnNoteRequest"`
			NoteRegex                      string `xml:"noteRegex"`
			CiSkip                         string `xml:"ciSkip"`
			SkipWorkInProgressMergeRequest string `xml:"skipWorkInProgressMergeRequest"`
			SetBuildDescription            string `xml:"setBuildDescription"`
			BranchFilterType               string `xml:"branchFilterType"`
			IncludeBranchesSpec            string `xml:"includeBranchesSpec"`
			ExcludeBranchesSpec            string `xml:"excludeBranchesSpec"`
			SourceBranchRegex              string `xml:"sourceBranchRegex"`
			TargetBranchRegex              string `xml:"targetBranchRegex"`
			SecretToken                    string `xml:"secretToken"`
			PendingBuildName               string `xml:"pendingBuildName"`
			CancelPendingBuildsOnUpdate    string `xml:"cancelPendingBuildsOnUpdate"`
		} `xml:"com.dabsquared.gitlabjenkins.GitLabPushTrigger"`
	} `xml:"triggers"`
	ConcurrentBuild string `xml:"concurrentBuild"`
	Builders        struct {
		Text             string `xml:",chardata"`
		HudsonTasksShell struct {
			Text    string `xml:",chardata"`
			Command string `xml:"command"`
		} `xml:"hudson.tasks.Shell"`
	} `xml:"builders"`
	Publishers    string `xml:"publishers"`
	BuildWrappers struct {
		Text                                            string `xml:",chardata"`
		HudsonPluginsTimestamperTimestamperBuildWrapper struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"hudson.plugins.timestamper.TimestamperBuildWrapper"`
		HudsonPluginsAnsicolorAnsiColorBuildWrapper struct {
			Text         string `xml:",chardata"`
			Plugin       string `xml:"plugin,attr"`
			ColorMapName string `xml:"colorMapName"`
		} `xml:"hudson.plugins.ansicolor.AnsiColorBuildWrapper"`
	} `xml:"buildWrappers"`
}

// GoBuildSSHStruct GO 构建 ssh部署 项目模板
type GoBuildSSHStruct struct {
	XMLName          xml.Name `xml:"project"`
	Text             string   `xml:",chardata"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		OrgJenkinsciPluginsGitlablogoGitlabLogoProperty struct {
			Text           string `xml:",chardata"`
			Plugin         string `xml:"plugin,attr"`
			RepositoryName string `xml:"repositoryName"`
		} `xml:"org.jenkinsci.plugins.gitlablogo.GitlabLogoProperty"`
		ComSynopsysArcJenkinsciPluginsJobrestrictionsJobsJobRestrictionProperty struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"com.synopsys.arc.jenkinsci.plugins.jobrestrictions.jobs.JobRestrictionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                                            string `xml:",chardata"`
				NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
					Text               string `xml:",chardata"`
					Plugin             string `xml:"plugin,attr"`
					Name               string `xml:"name"`
					Description        string `xml:"description"`
					Uuid               string `xml:"uuid"`
					Type               string `xml:"type"`
					Branch             string `xml:"branch"`
					TagFilter          string `xml:"tagFilter"`
					BranchFilter       string `xml:"branchFilter"`
					SortMode           string `xml:"sortMode"`
					DefaultValue       string `xml:"defaultValue"`
					SelectedValue      string `xml:"selectedValue"`
					QuickFilterEnabled string `xml:"quickFilterEnabled"`
					ListSize           string `xml:"listSize"`
				} `xml:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text              string `xml:",chardata"`
		Class             string `xml:"class,attr"`
		Plugin            string `xml:"plugin,attr"`
		ConfigVersion     string `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text          string `xml:",chardata"`
				URL           string `xml:"url"`
				CredentialsId string `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg                      struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions string `xml:"extensions"`
	} `xml:"scm"`
	AssignedNode                     string `xml:"assignedNode"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         struct {
		Text                                        string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsGitLabPushTrigger struct {
			Text                           string `xml:",chardata"`
			Plugin                         string `xml:"plugin,attr"`
			Spec                           string `xml:"spec"`
			TriggerOnPush                  string `xml:"triggerOnPush"`
			TriggerOnMergeRequest          string `xml:"triggerOnMergeRequest"`
			TriggerOnPipelineEvent         string `xml:"triggerOnPipelineEvent"`
			TriggerOnAcceptedMergeRequest  string `xml:"triggerOnAcceptedMergeRequest"`
			TriggerOnClosedMergeRequest    string `xml:"triggerOnClosedMergeRequest"`
			TriggerOnApprovedMergeRequest  string `xml:"triggerOnApprovedMergeRequest"`
			TriggerOpenMergeRequestOnPush  string `xml:"triggerOpenMergeRequestOnPush"`
			TriggerOnNoteRequest           string `xml:"triggerOnNoteRequest"`
			NoteRegex                      string `xml:"noteRegex"`
			CiSkip                         string `xml:"ciSkip"`
			SkipWorkInProgressMergeRequest string `xml:"skipWorkInProgressMergeRequest"`
			SetBuildDescription            string `xml:"setBuildDescription"`
			BranchFilterType               string `xml:"branchFilterType"`
			IncludeBranchesSpec            string `xml:"includeBranchesSpec"`
			ExcludeBranchesSpec            string `xml:"excludeBranchesSpec"`
			SourceBranchRegex              string `xml:"sourceBranchRegex"`
			TargetBranchRegex              string `xml:"targetBranchRegex"`
			SecretToken                    string `xml:"secretToken"`
			PendingBuildName               string `xml:"pendingBuildName"`
			CancelPendingBuildsOnUpdate    string `xml:"cancelPendingBuildsOnUpdate"`
		} `xml:"com.dabsquared.gitlabjenkins.GitLabPushTrigger"`
	} `xml:"triggers"`
	ConcurrentBuild string `xml:"concurrentBuild"`
	Builders        struct {
		Text             string `xml:",chardata"`
		HudsonTasksShell struct {
			Text    string `xml:",chardata"`
			Command string `xml:"command"`
		} `xml:"hudson.tasks.Shell"`
	} `xml:"builders"`
	Publishers struct {
		Text                                              string `xml:",chardata"`
		JenkinsPluginsPublishOverSshBapSshPublisherPlugin struct {
			Text          string `xml:",chardata"`
			Plugin        string `xml:"plugin,attr"`
			ConsolePrefix string `xml:"consolePrefix"`
			Delegate      struct {
				Text       string `xml:",chardata"`
				Plugin     string `xml:"plugin,attr"`
				Publishers struct {
					Text                                        string `xml:",chardata"`
					JenkinsPluginsPublishOverSshBapSshPublisher struct {
						Text       string `xml:",chardata"`
						Plugin     string `xml:"plugin,attr"`
						ConfigName string `xml:"configName"`
						Verbose    string `xml:"verbose"`
						Transfers  struct {
							Text                                       string `xml:",chardata"`
							JenkinsPluginsPublishOverSshBapSshTransfer struct {
								Text               string `xml:",chardata"`
								RemoteDirectory    string `xml:"remoteDirectory"`
								SourceFiles        string `xml:"sourceFiles"`
								Excludes           string `xml:"excludes"`
								RemovePrefix       string `xml:"removePrefix"`
								RemoteDirectorySDF string `xml:"remoteDirectorySDF"`
								Flatten            string `xml:"flatten"`
								CleanRemote        string `xml:"cleanRemote"`
								NoDefaultExcludes  string `xml:"noDefaultExcludes"`
								MakeEmptyDirs      string `xml:"makeEmptyDirs"`
								PatternSeparator   string `xml:"patternSeparator"`
								ExecCommand        string `xml:"execCommand"`
							} `xml:"jenkins.plugins.publish__over__ssh.BapSshTransfer"`
						} `xml:"transfers"`
						UseWorkspaceInPromotion string `xml:"useWorkspaceInPromotion"`
						UsePromotionTimestamp   string `xml:"usePromotionTimestamp"`
					} `xml:"jenkins.plugins.publish__over__ssh.BapSshPublisher"`
				} `xml:"publishers"`
				ContinueOnError         string `xml:"continueOnError"`
				FailOnError             string `xml:"failOnError"`
				AlwaysPublishFromMaster string `xml:"alwaysPublishFromMaster"`
				HostConfigurationAccess struct {
					Text      string `xml:",chardata"`
					Class     string `xml:"class,attr"`
					Reference string `xml:"reference,attr"`
				} `xml:"hostConfigurationAccess"`
			} `xml:"delegate"`
		} `xml:"jenkins.plugins.publish__over__ssh.BapSshPublisherPlugin"`
	} `xml:"publishers"`
	BuildWrappers struct {
		Text                                            string `xml:",chardata"`
		HudsonPluginsTimestamperTimestamperBuildWrapper struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"hudson.plugins.timestamper.TimestamperBuildWrapper"`
		HudsonPluginsAnsicolorAnsiColorBuildWrapper struct {
			Text         string `xml:",chardata"`
			Plugin       string `xml:"plugin,attr"`
			ColorMapName string `xml:"colorMapName"`
		} `xml:"hudson.plugins.ansicolor.AnsiColorBuildWrapper"`
	} `xml:"buildWrappers"`
}

// java 项目模板
type Maven2Moduleset struct {
	XMLName          xml.Name `xml:"maven2-moduleset"`
	Text             string   `xml:",chardata"`
	Plugin           string   `xml:"plugin,attr"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		OrgJenkinsciPluginsGitlablogoGitlabLogoProperty struct {
			Text           string `xml:",chardata"`
			Plugin         string `xml:"plugin,attr"`
			RepositoryName string `xml:"repositoryName"`
		} `xml:"org.jenkinsci.plugins.gitlablogo.GitlabLogoProperty"`
		ComSynopsysArcJenkinsciPluginsJobrestrictionsJobsJobRestrictionProperty struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"com.synopsys.arc.jenkinsci.plugins.jobrestrictions.jobs.JobRestrictionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                                            string `xml:",chardata"`
				NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
					Text               string `xml:",chardata"`
					Plugin             string `xml:"plugin,attr"`
					Name               string `xml:"name"`
					Description        string `xml:"description"`
					Uuid               string `xml:"uuid"`
					Type               string `xml:"type"`
					Branch             string `xml:"branch"`
					TagFilter          string `xml:"tagFilter"`
					BranchFilter       string `xml:"branchFilter"`
					SortMode           string `xml:"sortMode"`
					DefaultValue       string `xml:"defaultValue"`
					SelectedValue      string `xml:"selectedValue"`
					QuickFilterEnabled string `xml:"quickFilterEnabled"`
					ListSize           string `xml:"listSize"`
				} `xml:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text              string `xml:",chardata"`
		Class             string `xml:"class,attr"`
		Plugin            string `xml:"plugin,attr"`
		ConfigVersion     string `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text          string `xml:",chardata"`
				URL           string `xml:"url"`
				CredentialsId string `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg                      struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions string `xml:"extensions"`
	} `xml:"scm"`
	AssignedNode                     string `xml:"assignedNode"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         string `xml:"triggers"`
	ConcurrentBuild                  string `xml:"concurrentBuild"`
	RootModule                       struct {
		Text       string `xml:",chardata"`
		GroupId    string `xml:"groupId"`
		ArtifactId string `xml:"artifactId"`
	} `xml:"rootModule"`
	Goals                            string `xml:"goals"`
	MavenName                        string `xml:"mavenName"`
	AggregatorStyleBuild             string `xml:"aggregatorStyleBuild"`
	IncrementalBuild                 string `xml:"incrementalBuild"`
	IgnoreUpstremChanges             string `xml:"ignoreUpstremChanges"`
	IgnoreUnsuccessfulUpstreams      string `xml:"ignoreUnsuccessfulUpstreams"`
	ArchivingDisabled                string `xml:"archivingDisabled"`
	SiteArchivingDisabled            string `xml:"siteArchivingDisabled"`
	FingerprintingDisabled           string `xml:"fingerprintingDisabled"`
	ResolveDependencies              string `xml:"resolveDependencies"`
	ProcessPlugins                   string `xml:"processPlugins"`
	MavenValidationLevel             string `xml:"mavenValidationLevel"`
	RunHeadless                      string `xml:"runHeadless"`
	DisableTriggerDownstreamProjects string `xml:"disableTriggerDownstreamProjects"`
	BlockTriggerWhenBuilding         string `xml:"blockTriggerWhenBuilding"`
	Settings                         struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
	} `xml:"settings"`
	GlobalSettings struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
	} `xml:"globalSettings"`
	Reporters     string `xml:"reporters"`
	Publishers    string `xml:"publishers"`
	BuildWrappers struct {
		Text                                            string `xml:",chardata"`
		HudsonPluginsTimestamperTimestamperBuildWrapper struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"hudson.plugins.timestamper.TimestamperBuildWrapper"`
	} `xml:"buildWrappers"`
	Prebuilders  string `xml:"prebuilders"`
	Postbuilders struct {
		Text             string `xml:",chardata"`
		HudsonTasksShell struct {
			Text    string `xml:",chardata"`
			Command string `xml:"command"`
		} `xml:"hudson.tasks.Shell"`
	} `xml:"postbuilders"`
	RunPostStepsIfResult struct {
		Text          string `xml:",chardata"`
		Name          string `xml:"name"`
		Ordinal       string `xml:"ordinal"`
		Color         string `xml:"color"`
		CompleteBuild string `xml:"completeBuild"`
	} `xml:"runPostStepsIfResult"`
}

// java 项目模板 有wx通知
type Maven2Struct struct {
	XMLName          xml.Name `xml:"maven2-moduleset"`
	Text             string   `xml:",chardata"`
	Plugin           string   `xml:"plugin,attr"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                                            string `xml:",chardata"`
				NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
					Text               string `xml:",chardata"`
					Plugin             string `xml:"plugin,attr"`
					Name               string `xml:"name"`
					Description        string `xml:"description"`
					Uuid               string `xml:"uuid"`
					Type               string `xml:"type"`
					Branch             string `xml:"branch"`
					TagFilter          string `xml:"tagFilter"`
					BranchFilter       string `xml:"branchFilter"`
					SortMode           string `xml:"sortMode"`
					DefaultValue       string `xml:"defaultValue"`
					SelectedValue      string `xml:"selectedValue"`
					QuickFilterEnabled string `xml:"quickFilterEnabled"`
					ListSize           string `xml:"listSize"`
				} `xml:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text              string `xml:",chardata"`
		Class             string `xml:"class,attr"`
		Plugin            string `xml:"plugin,attr"`
		ConfigVersion     string `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text          string `xml:",chardata"`
				URL           string `xml:"url"`
				CredentialsId string `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg                      struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions string `xml:"extensions"`
	} `xml:"scm"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         string `xml:"triggers"`
	ConcurrentBuild                  string `xml:"concurrentBuild"`
	Goals                            string `xml:"goals"`
	AggregatorStyleBuild             string `xml:"aggregatorStyleBuild"`
	IncrementalBuild                 string `xml:"incrementalBuild"`
	IgnoreUpstremChanges             string `xml:"ignoreUpstremChanges"`
	IgnoreUnsuccessfulUpstreams      string `xml:"ignoreUnsuccessfulUpstreams"`
	ArchivingDisabled                string `xml:"archivingDisabled"`
	SiteArchivingDisabled            string `xml:"siteArchivingDisabled"`
	FingerprintingDisabled           string `xml:"fingerprintingDisabled"`
	ResolveDependencies              string `xml:"resolveDependencies"`
	ProcessPlugins                   string `xml:"processPlugins"`
	MavenValidationLevel             string `xml:"mavenValidationLevel"`
	RunHeadless                      string `xml:"runHeadless"`
	DisableTriggerDownstreamProjects string `xml:"disableTriggerDownstreamProjects"`
	BlockTriggerWhenBuilding         string `xml:"blockTriggerWhenBuilding"`
	Settings                         struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
	} `xml:"settings"`
	GlobalSettings struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
	} `xml:"globalSettings"`
	Reporters  string `xml:"reporters"`
	Publishers struct {
		Text                                            string `xml:",chardata"`
		OrgJenkinsciPluginsQywechatQyWechatNotification struct {
			Text            string `xml:",chardata"`
			Plugin          string `xml:"plugin,attr"`
			WebhookUrl      string `xml:"webhookUrl"`
			MentionedId     string `xml:"mentionedId"`
			MentionedMobile string `xml:"mentionedMobile"`
			FailNotify      string `xml:"failNotify"`
			FailSend        string `xml:"failSend"`
			SuccessSend     string `xml:"successSend"`
			AboutSend       string `xml:"aboutSend"`
			UnstableSend    string `xml:"unstableSend"`
			StartBuild      string `xml:"startBuild"`
		} `xml:"org.jenkinsci.plugins.qywechat.QyWechatNotification"`
	} `xml:"publishers"`
	BuildWrappers struct {
		Text                                            string `xml:",chardata"`
		HudsonPluginsTimestamperTimestamperBuildWrapper struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"hudson.plugins.timestamper.TimestamperBuildWrapper"`
	} `xml:"buildWrappers"`
	Prebuilders  string `xml:"prebuilders"`
	Postbuilders struct {
		Text             string `xml:",chardata"`
		HudsonTasksShell struct {
			Text    string `xml:",chardata"`
			Command string `xml:"command"`
		} `xml:"hudson.tasks.Shell"`
	} `xml:"postbuilders"`
	RunPostStepsIfResult struct {
		Text          string `xml:",chardata"`
		Name          string `xml:"name"`
		Ordinal       string `xml:"ordinal"`
		Color         string `xml:"color"`
		CompleteBuild string `xml:"completeBuild"`
	} `xml:"runPostStepsIfResult"`
}

// deploy argocd 项目模板
type DeployStructlod struct {
	XMLName          xml.Name `xml:"project"`
	Text             string   `xml:",chardata"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		OrgJenkinsciPluginsGitlablogoGitlabLogoProperty struct {
			Text           string `xml:",chardata"`
			Plugin         string `xml:"plugin,attr"`
			RepositoryName string `xml:"repositoryName"`
		} `xml:"org.jenkinsci.plugins.gitlablogo.GitlabLogoProperty"`
		ComSynopsysArcJenkinsciPluginsJobrestrictionsJobsJobRestrictionProperty struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"com.synopsys.arc.jenkinsci.plugins.jobrestrictions.jobs.JobRestrictionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                 string `xml:",chardata"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
	} `xml:"scm"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         string `xml:"triggers"`
	ConcurrentBuild                  string `xml:"concurrentBuild"`
	Builders                         struct {
		Text             string `xml:",chardata"`
		HudsonTasksShell struct {
			Text    string `xml:",chardata"`
			Command string `xml:"command"`
		} `xml:"hudson.tasks.Shell"`
	} `xml:"builders"`
	Publishers    string `xml:"publishers"`
	BuildWrappers string `xml:"buildWrappers"`
}

// deploy argocd 项目模板
type DeployStruct struct {
	XMLName          xml.Name `xml:"project"`
	Text             string   `xml:",chardata"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		OrgJenkinsciPluginsGitlablogoGitlabLogoProperty struct {
			Text           string `xml:",chardata"`
			Plugin         string `xml:"plugin,attr"`
			RepositoryName string `xml:"repositoryName"`
		} `xml:"org.jenkinsci.plugins.gitlablogo.GitlabLogoProperty"`
		ComSynopsysArcJenkinsciPluginsJobrestrictionsJobsJobRestrictionProperty struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"com.synopsys.arc.jenkinsci.plugins.jobrestrictions.jobs.JobRestrictionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                 string `xml:",chardata"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
	} `xml:"scm"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         string `xml:"triggers"`
	ConcurrentBuild                  string `xml:"concurrentBuild"`
	Builders                         struct {
		Text             string `xml:",chardata"`
		HudsonTasksShell struct {
			Text    string `xml:",chardata"`
			Command string `xml:"command"`
		} `xml:"hudson.tasks.Shell"`
	} `xml:"builders"`
	Publishers struct {
		Text                                            string `xml:",chardata"`
		OrgJenkinsciPluginsQywechatQyWechatNotification struct {
			Text            string `xml:",chardata"`
			Plugin          string `xml:"plugin,attr"`
			WebhookUrl      string `xml:"webhookUrl"`
			MentionedId     string `xml:"mentionedId"`
			MentionedMobile string `xml:"mentionedMobile"`
			FailNotify      string `xml:"failNotify"`
			FailSend        string `xml:"failSend"`
			SuccessSend     string `xml:"successSend"`
			AboutSend       string `xml:"aboutSend"`
			UnstableSend    string `xml:"unstableSend"`
			StartBuild      string `xml:"startBuild"`
		} `xml:"org.jenkinsci.plugins.qywechat.QyWechatNotification"`
	} `xml:"publishers"`
	BuildWrappers string `xml:"buildWrappers"`
}

//nodejs-template-build-deploy 前端构建部署
type NodejsStruct struct {
	XMLName          xml.Name `xml:"project"`
	Text             string   `xml:",chardata"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                                            string `xml:",chardata"`
				NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
					Text               string `xml:",chardata"`
					Plugin             string `xml:"plugin,attr"`
					Name               string `xml:"name"`
					Description        string `xml:"description"`
					Uuid               string `xml:"uuid"`
					Type               string `xml:"type"`
					Branch             string `xml:"branch"`
					TagFilter          string `xml:"tagFilter"`
					BranchFilter       string `xml:"branchFilter"`
					SortMode           string `xml:"sortMode"`
					DefaultValue       string `xml:"defaultValue"`
					SelectedValue      string `xml:"selectedValue"`
					QuickFilterEnabled string `xml:"quickFilterEnabled"`
					ListSize           string `xml:"listSize"`
				} `xml:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text              string `xml:",chardata"`
		Class             string `xml:"class,attr"`
		Plugin            string `xml:"plugin,attr"`
		ConfigVersion     string `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text          string `xml:",chardata"`
				URL           string `xml:"url"`
				CredentialsId string `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg                      struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions string `xml:"extensions"`
	} `xml:"scm"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         string `xml:"triggers"`
	ConcurrentBuild                  string `xml:"concurrentBuild"`
	Builders                         struct {
		Text             string `xml:",chardata"`
		HudsonTasksShell struct {
			Text    string `xml:",chardata"`
			Command string `xml:"command"`
		} `xml:"hudson.tasks.Shell"`
	} `xml:"builders"`
	Publishers struct {
		Text                                              string `xml:",chardata"`
		JenkinsPluginsPublishOverSshBapSshPublisherPlugin struct {
			Text          string `xml:",chardata"`
			Plugin        string `xml:"plugin,attr"`
			ConsolePrefix string `xml:"consolePrefix"`
			Delegate      struct {
				Text       string `xml:",chardata"`
				Plugin     string `xml:"plugin,attr"`
				Publishers struct {
					Text                                        string `xml:",chardata"`
					JenkinsPluginsPublishOverSshBapSshPublisher []struct {
						Text       string `xml:",chardata"`
						Plugin     string `xml:"plugin,attr"`
						ConfigName string `xml:"configName"`
						Verbose    string `xml:"verbose"`
						Transfers  struct {
							Text                                       string `xml:",chardata"`
							JenkinsPluginsPublishOverSshBapSshTransfer struct {
								Text               string `xml:",chardata"`
								RemoteDirectory    string `xml:"remoteDirectory"`
								SourceFiles        string `xml:"sourceFiles"`
								Excludes           string `xml:"excludes"`
								RemovePrefix       string `xml:"removePrefix"`
								RemoteDirectorySDF string `xml:"remoteDirectorySDF"`
								Flatten            string `xml:"flatten"`
								CleanRemote        string `xml:"cleanRemote"`
								NoDefaultExcludes  string `xml:"noDefaultExcludes"`
								MakeEmptyDirs      string `xml:"makeEmptyDirs"`
								PatternSeparator   string `xml:"patternSeparator"`
								ExecCommand        string `xml:"execCommand"`
								ExecTimeout        string `xml:"execTimeout"`
								UsePty             string `xml:"usePty"`
								UseAgentForwarding string `xml:"useAgentForwarding"`
							} `xml:"jenkins.plugins.publish__over__ssh.BapSshTransfer"`
						} `xml:"transfers"`
						UseWorkspaceInPromotion string `xml:"useWorkspaceInPromotion"`
						UsePromotionTimestamp   string `xml:"usePromotionTimestamp"`
					} `xml:"jenkins.plugins.publish__over__ssh.BapSshPublisher"`
				} `xml:"publishers"`
				ContinueOnError         string `xml:"continueOnError"`
				FailOnError             string `xml:"failOnError"`
				AlwaysPublishFromMaster string `xml:"alwaysPublishFromMaster"`
				HostConfigurationAccess struct {
					Text      string `xml:",chardata"`
					Class     string `xml:"class,attr"`
					Reference string `xml:"reference,attr"`
				} `xml:"hostConfigurationAccess"`
			} `xml:"delegate"`
		} `xml:"jenkins.plugins.publish__over__ssh.BapSshPublisherPlugin"`
		OrgJenkinsciPluginsQywechatQyWechatNotification struct {
			Text            string `xml:",chardata"`
			Plugin          string `xml:"plugin,attr"`
			WebhookUrl      string `xml:"webhookUrl"`
			MentionedId     string `xml:"mentionedId"`
			MentionedMobile string `xml:"mentionedMobile"`
			FailNotify      string `xml:"failNotify"`
			FailSend        string `xml:"failSend"`
			SuccessSend     string `xml:"successSend"`
			AboutSend       string `xml:"aboutSend"`
			UnstableSend    string `xml:"unstableSend"`
			StartBuild      string `xml:"startBuild"`
		} `xml:"org.jenkinsci.plugins.qywechat.QyWechatNotification"`
	} `xml:"publishers"`
	BuildWrappers struct {
		Text                                            string `xml:",chardata"`
		HudsonPluginsTimestamperTimestamperBuildWrapper struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"hudson.plugins.timestamper.TimestamperBuildWrapper"`
		JenkinsPluginsNodejsNodeJSBuildWrapper struct {
			Text                   string `xml:",chardata"`
			Plugin                 string `xml:"plugin,attr"`
			NodeJSInstallationName string `xml:"nodeJSInstallationName"`
			CacheLocationStrategy  struct {
				Text  string `xml:",chardata"`
				Class string `xml:"class,attr"`
			} `xml:"cacheLocationStrategy"`
		} `xml:"jenkins.plugins.nodejs.NodeJSBuildWrapper"`
	} `xml:"buildWrappers"`
}

// 前端初始构建模板 nodejs-template-build
type NodejsBuildStruct struct {
	XMLName          xml.Name `xml:"project"`
	Text             string   `xml:",chardata"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                                            string `xml:",chardata"`
				NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
					Text               string `xml:",chardata"`
					Plugin             string `xml:"plugin,attr"`
					Name               string `xml:"name"`
					Description        string `xml:"description"`
					Uuid               string `xml:"uuid"`
					Type               string `xml:"type"`
					Branch             string `xml:"branch"`
					TagFilter          string `xml:"tagFilter"`
					BranchFilter       string `xml:"branchFilter"`
					SortMode           string `xml:"sortMode"`
					DefaultValue       string `xml:"defaultValue"`
					SelectedValue      string `xml:"selectedValue"`
					QuickFilterEnabled string `xml:"quickFilterEnabled"`
					ListSize           string `xml:"listSize"`
				} `xml:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text              string `xml:",chardata"`
		Class             string `xml:"class,attr"`
		Plugin            string `xml:"plugin,attr"`
		ConfigVersion     string `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text          string `xml:",chardata"`
				URL           string `xml:"url"`
				CredentialsId string `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg                      struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions string `xml:"extensions"`
	} `xml:"scm"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         string `xml:"triggers"`
	ConcurrentBuild                  string `xml:"concurrentBuild"`
	Builders                         struct {
		Text             string `xml:",chardata"`
		HudsonTasksShell struct {
			Text    string `xml:",chardata"`
			Command string `xml:"command"`
		} `xml:"hudson.tasks.Shell"`
	} `xml:"builders"`
	Publishers struct {
		Text                                            string `xml:",chardata"`
		OrgJenkinsciPluginsQywechatQyWechatNotification struct {
			Text            string `xml:",chardata"`
			Plugin          string `xml:"plugin,attr"`
			WebhookUrl      string `xml:"webhookUrl"`
			MentionedId     string `xml:"mentionedId"`
			MentionedMobile string `xml:"mentionedMobile"`
			FailNotify      string `xml:"failNotify"`
			FailSend        string `xml:"failSend"`
			SuccessSend     string `xml:"successSend"`
			AboutSend       string `xml:"aboutSend"`
			UnstableSend    string `xml:"unstableSend"`
			StartBuild      string `xml:"startBuild"`
		} `xml:"org.jenkinsci.plugins.qywechat.QyWechatNotification"`
	} `xml:"publishers"`
	BuildWrappers struct {
		Text                                            string `xml:",chardata"`
		HudsonPluginsTimestamperTimestamperBuildWrapper struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"hudson.plugins.timestamper.TimestamperBuildWrapper"`
		JenkinsPluginsNodejsNodeJSBuildWrapper struct {
			Text                   string `xml:",chardata"`
			Plugin                 string `xml:"plugin,attr"`
			NodeJSInstallationName string `xml:"nodeJSInstallationName"`
			CacheLocationStrategy  struct {
				Text  string `xml:",chardata"`
				Class string `xml:"class,attr"`
			} `xml:"cacheLocationStrategy"`
		} `xml:"jenkins.plugins.nodejs.NodeJSBuildWrapper"`
	} `xml:"buildWrappers"`
}

//nodejs-template-build-nginx-deploy 模板函数处理
type NodeBuildNginxDeployStruct struct {
	XMLName          xml.Name `xml:"project"`
	Text             string   `xml:",chardata"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		OrgJenkinsciPluginsGitlablogoGitlabLogoProperty struct {
			Text           string `xml:",chardata"`
			Plugin         string `xml:"plugin,attr"`
			RepositoryName string `xml:"repositoryName"`
		} `xml:"org.jenkinsci.plugins.gitlablogo.GitlabLogoProperty"`
		ComSynopsysArcJenkinsciPluginsJobrestrictionsJobsJobRestrictionProperty struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"com.synopsys.arc.jenkinsci.plugins.jobrestrictions.jobs.JobRestrictionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                                            string `xml:",chardata"`
				NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
					Text               string `xml:",chardata"`
					Plugin             string `xml:"plugin,attr"`
					Name               string `xml:"name"`
					Description        string `xml:"description"`
					Uuid               string `xml:"uuid"`
					Type               string `xml:"type"`
					Branch             string `xml:"branch"`
					TagFilter          string `xml:"tagFilter"`
					BranchFilter       string `xml:"branchFilter"`
					SortMode           string `xml:"sortMode"`
					DefaultValue       string `xml:"defaultValue"`
					SelectedValue      string `xml:"selectedValue"`
					QuickFilterEnabled string `xml:"quickFilterEnabled"`
					ListSize           string `xml:"listSize"`
				} `xml:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text              string `xml:",chardata"`
		Class             string `xml:"class,attr"`
		Plugin            string `xml:"plugin,attr"`
		ConfigVersion     string `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text          string `xml:",chardata"`
				URL           string `xml:"url"`
				CredentialsId string `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg                      struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions string `xml:"extensions"`
	} `xml:"scm"`
	AssignedNode                     string `xml:"assignedNode"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         string `xml:"triggers"`
	ConcurrentBuild                  string `xml:"concurrentBuild"`
	Builders                         struct {
		Text             string `xml:",chardata"`
		HudsonTasksShell struct {
			Text    string `xml:",chardata"`
			Command string `xml:"command"`
		} `xml:"hudson.tasks.Shell"`
	} `xml:"builders"`
	Publishers struct {
		Text                                              string `xml:",chardata"`
		JenkinsPluginsPublishOverSshBapSshPublisherPlugin struct {
			Text          string `xml:",chardata"`
			Plugin        string `xml:"plugin,attr"`
			ConsolePrefix string `xml:"consolePrefix"`
			Delegate      struct {
				Text       string `xml:",chardata"`
				Plugin     string `xml:"plugin,attr"`
				Publishers struct {
					Text                                        string `xml:",chardata"`
					JenkinsPluginsPublishOverSshBapSshPublisher struct {
						Text       string `xml:",chardata"`
						Plugin     string `xml:"plugin,attr"`
						ConfigName string `xml:"configName"`
						Verbose    string `xml:"verbose"`
						Transfers  struct {
							Text                                       string `xml:",chardata"`
							JenkinsPluginsPublishOverSshBapSshTransfer struct {
								Text               string `xml:",chardata"`
								RemoteDirectory    string `xml:"remoteDirectory"`
								SourceFiles        string `xml:"sourceFiles"`
								Excludes           string `xml:"excludes"`
								RemovePrefix       string `xml:"removePrefix"`
								RemoteDirectorySDF string `xml:"remoteDirectorySDF"`
								Flatten            string `xml:"flatten"`
								CleanRemote        string `xml:"cleanRemote"`
								NoDefaultExcludes  string `xml:"noDefaultExcludes"`
								MakeEmptyDirs      string `xml:"makeEmptyDirs"`
								PatternSeparator   string `xml:"patternSeparator"`
								ExecCommand        string `xml:"execCommand"`
								ExecTimeout        string `xml:"execTimeout"`
								UsePty             string `xml:"usePty"`
								UseAgentForwarding string `xml:"useAgentForwarding"`
							} `xml:"jenkins.plugins.publish__over__ssh.BapSshTransfer"`
						} `xml:"transfers"`
						UseWorkspaceInPromotion string `xml:"useWorkspaceInPromotion"`
						UsePromotionTimestamp   string `xml:"usePromotionTimestamp"`
					} `xml:"jenkins.plugins.publish__over__ssh.BapSshPublisher"`
				} `xml:"publishers"`
				ContinueOnError         string `xml:"continueOnError"`
				FailOnError             string `xml:"failOnError"`
				AlwaysPublishFromMaster string `xml:"alwaysPublishFromMaster"`
				HostConfigurationAccess struct {
					Text      string `xml:",chardata"`
					Class     string `xml:"class,attr"`
					Reference string `xml:"reference,attr"`
				} `xml:"hostConfigurationAccess"`
			} `xml:"delegate"`
		} `xml:"jenkins.plugins.publish__over__ssh.BapSshPublisherPlugin"`
	} `xml:"publishers"`
	BuildWrappers struct {
		Text                                   string `xml:",chardata"`
		JenkinsPluginsNodejsNodeJSBuildWrapper struct {
			Text                   string `xml:",chardata"`
			Plugin                 string `xml:"plugin,attr"`
			NodeJSInstallationName string `xml:"nodeJSInstallationName"`
			CacheLocationStrategy  struct {
				Text  string `xml:",chardata"`
				Class string `xml:"class,attr"`
			} `xml:"cacheLocationStrategy"`
		} `xml:"jenkins.plugins.nodejs.NodeJSBuildWrapper"`
	} `xml:"buildWrappers"`
}

// ssh nginx 分发模板前端
type NodeNginxDeploySshStruct struct {
	XMLName          xml.Name `xml:"project"`
	Text             string   `xml:",chardata"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                                            string `xml:",chardata"`
				NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
					Text               string `xml:",chardata"`
					Plugin             string `xml:"plugin,attr"`
					Name               string `xml:"name"`
					Description        string `xml:"description"`
					Uuid               string `xml:"uuid"`
					Type               string `xml:"type"`
					Branch             string `xml:"branch"`
					TagFilter          string `xml:"tagFilter"`
					BranchFilter       string `xml:"branchFilter"`
					SortMode           string `xml:"sortMode"`
					DefaultValue       string `xml:"defaultValue"`
					SelectedValue      string `xml:"selectedValue"`
					QuickFilterEnabled string `xml:"quickFilterEnabled"`
					ListSize           string `xml:"listSize"`
				} `xml:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text              string `xml:",chardata"`
		Class             string `xml:"class,attr"`
		Plugin            string `xml:"plugin,attr"`
		ConfigVersion     string `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text          string `xml:",chardata"`
				URL           string `xml:"url"`
				CredentialsId string `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg                      struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions string `xml:"extensions"`
	} `xml:"scm"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         string `xml:"triggers"`
	ConcurrentBuild                  string `xml:"concurrentBuild"`
	Builders                         struct {
		Text             string `xml:",chardata"`
		HudsonTasksShell struct {
			Text    string `xml:",chardata"`
			Command string `xml:"command"`
		} `xml:"hudson.tasks.Shell"`
	} `xml:"builders"`
	Publishers struct {
		Text                                              string `xml:",chardata"`
		JenkinsPluginsPublishOverSshBapSshPublisherPlugin struct {
			Text          string `xml:",chardata"`
			Plugin        string `xml:"plugin,attr"`
			ConsolePrefix string `xml:"consolePrefix"`
			Delegate      struct {
				Text       string `xml:",chardata"`
				Plugin     string `xml:"plugin,attr"`
				Publishers struct {
					Text                                        string `xml:",chardata"`
					JenkinsPluginsPublishOverSshBapSshPublisher []struct {
						Text       string `xml:",chardata"`
						Plugin     string `xml:"plugin,attr"`
						ConfigName string `xml:"configName"`
						Verbose    string `xml:"verbose"`
						Transfers  struct {
							Text                                       string `xml:",chardata"`
							JenkinsPluginsPublishOverSshBapSshTransfer struct {
								Text               string `xml:",chardata"`
								RemoteDirectory    string `xml:"remoteDirectory"`
								SourceFiles        string `xml:"sourceFiles"`
								Excludes           string `xml:"excludes"`
								RemovePrefix       string `xml:"removePrefix"`
								RemoteDirectorySDF string `xml:"remoteDirectorySDF"`
								Flatten            string `xml:"flatten"`
								CleanRemote        string `xml:"cleanRemote"`
								NoDefaultExcludes  string `xml:"noDefaultExcludes"`
								MakeEmptyDirs      string `xml:"makeEmptyDirs"`
								PatternSeparator   string `xml:"patternSeparator"`
								ExecCommand        string `xml:"execCommand"`
								ExecTimeout        string `xml:"execTimeout"`
								UsePty             string `xml:"usePty"`
								UseAgentForwarding string `xml:"useAgentForwarding"`
							} `xml:"jenkins.plugins.publish__over__ssh.BapSshTransfer"`
						} `xml:"transfers"`
						UseWorkspaceInPromotion string `xml:"useWorkspaceInPromotion"`
						UsePromotionTimestamp   string `xml:"usePromotionTimestamp"`
					} `xml:"jenkins.plugins.publish__over__ssh.BapSshPublisher"`
				} `xml:"publishers"`
				ContinueOnError         string `xml:"continueOnError"`
				FailOnError             string `xml:"failOnError"`
				AlwaysPublishFromMaster string `xml:"alwaysPublishFromMaster"`
				HostConfigurationAccess struct {
					Text      string `xml:",chardata"`
					Class     string `xml:"class,attr"`
					Reference string `xml:"reference,attr"`
				} `xml:"hostConfigurationAccess"`
			} `xml:"delegate"`
		} `xml:"jenkins.plugins.publish__over__ssh.BapSshPublisherPlugin"`
		OrgJenkinsciPluginsQywechatQyWechatNotification struct {
			Text            string `xml:",chardata"`
			Plugin          string `xml:"plugin,attr"`
			WebhookUrl      string `xml:"webhookUrl"`
			MentionedId     string `xml:"mentionedId"`
			MentionedMobile string `xml:"mentionedMobile"`
			FailNotify      string `xml:"failNotify"`
			FailSend        string `xml:"failSend"`
			SuccessSend     string `xml:"successSend"`
			AboutSend       string `xml:"aboutSend"`
			UnstableSend    string `xml:"unstableSend"`
			StartBuild      string `xml:"startBuild"`
		} `xml:"org.jenkinsci.plugins.qywechat.QyWechatNotification"`
	} `xml:"publishers"`
	BuildWrappers struct {
		Text                                            string `xml:",chardata"`
		HudsonPluginsTimestamperTimestamperBuildWrapper struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"hudson.plugins.timestamper.TimestamperBuildWrapper"`
		JenkinsPluginsNodejsNodeJSBuildWrapper struct {
			Text                   string `xml:",chardata"`
			Plugin                 string `xml:"plugin,attr"`
			NodeJSInstallationName string `xml:"nodeJSInstallationName"`
			CacheLocationStrategy  struct {
				Text  string `xml:",chardata"`
				Class string `xml:"class,attr"`
			} `xml:"cacheLocationStrategy"`
		} `xml:"jenkins.plugins.nodejs.NodeJSBuildWrapper"`
	} `xml:"buildWrappers"`
}

// java-backend-template-build-base 用于构建java后端项目的基础模板 无wx通知
type JavaBackEndTemplateBuildBaseNOprod struct {
	XMLName          xml.Name `xml:"maven2-moduleset"`
	Text             string   `xml:",chardata"`
	Plugin           string   `xml:"plugin,attr"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		OrgJenkinsciPluginsGitlablogoGitlabLogoProperty struct {
			Text           string `xml:",chardata"`
			Plugin         string `xml:"plugin,attr"`
			RepositoryName string `xml:"repositoryName"`
		} `xml:"org.jenkinsci.plugins.gitlablogo.GitlabLogoProperty"`
		ComSynopsysArcJenkinsciPluginsJobrestrictionsJobsJobRestrictionProperty struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"com.synopsys.arc.jenkinsci.plugins.jobrestrictions.jobs.JobRestrictionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                                            string `xml:",chardata"`
				NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
					Text               string `xml:",chardata"`
					Plugin             string `xml:"plugin,attr"`
					Name               string `xml:"name"`
					Description        string `xml:"description"`
					Uuid               string `xml:"uuid"`
					Type               string `xml:"type"`
					Branch             string `xml:"branch"`
					TagFilter          string `xml:"tagFilter"`
					BranchFilter       string `xml:"branchFilter"`
					SortMode           string `xml:"sortMode"`
					DefaultValue       string `xml:"defaultValue"`
					SelectedValue      string `xml:"selectedValue"`
					QuickFilterEnabled string `xml:"quickFilterEnabled"`
					ListSize           string `xml:"listSize"`
				} `xml:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text              string `xml:",chardata"`
		Class             string `xml:"class,attr"`
		Plugin            string `xml:"plugin,attr"`
		ConfigVersion     string `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text          string `xml:",chardata"`
				URL           string `xml:"url"`
				CredentialsId string `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg                      struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions string `xml:"extensions"`
	} `xml:"scm"`
	AssignedNode                     string `xml:"assignedNode"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         string `xml:"triggers"`
	ConcurrentBuild                  string `xml:"concurrentBuild"`
	Goals                            string `xml:"goals"`
	MavenName                        string `xml:"mavenName"`
	AggregatorStyleBuild             string `xml:"aggregatorStyleBuild"`
	IncrementalBuild                 string `xml:"incrementalBuild"`
	IgnoreUpstremChanges             string `xml:"ignoreUpstremChanges"`
	IgnoreUnsuccessfulUpstreams      string `xml:"ignoreUnsuccessfulUpstreams"`
	ArchivingDisabled                string `xml:"archivingDisabled"`
	SiteArchivingDisabled            string `xml:"siteArchivingDisabled"`
	FingerprintingDisabled           string `xml:"fingerprintingDisabled"`
	ResolveDependencies              string `xml:"resolveDependencies"`
	ProcessPlugins                   string `xml:"processPlugins"`
	MavenValidationLevel             string `xml:"mavenValidationLevel"`
	RunHeadless                      string `xml:"runHeadless"`
	DisableTriggerDownstreamProjects string `xml:"disableTriggerDownstreamProjects"`
	BlockTriggerWhenBuilding         string `xml:"blockTriggerWhenBuilding"`
	Settings                         struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
	} `xml:"settings"`
	GlobalSettings struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
	} `xml:"globalSettings"`
	Reporters     string `xml:"reporters"`
	Publishers    string `xml:"publishers"`
	BuildWrappers struct {
		Text                                            string `xml:",chardata"`
		HudsonPluginsTimestamperTimestamperBuildWrapper struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"hudson.plugins.timestamper.TimestamperBuildWrapper"`
	} `xml:"buildWrappers"`
	Prebuilders          string `xml:"prebuilders"`
	Postbuilders         string `xml:"postbuilders"`
	RunPostStepsIfResult struct {
		Text          string `xml:",chardata"`
		Name          string `xml:"name"`
		Ordinal       string `xml:"ordinal"`
		Color         string `xml:"color"`
		CompleteBuild string `xml:"completeBuild"`
	} `xml:"runPostStepsIfResult"`
}

// java-backend-template-build-base 用于构建java后端项目的基础模板 有wx通知
type JavaBackEndTemplateBuildBase struct {
	XMLName          xml.Name `xml:"maven2-moduleset"`
	Text             string   `xml:",chardata"`
	Plugin           string   `xml:"plugin,attr"`
	Actions          string   `xml:"actions"`
	Description      string   `xml:"description"`
	KeepDependencies string   `xml:"keepDependencies"`
	Properties       struct {
		Text                                                         string `xml:",chardata"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text               string `xml:",chardata"`
				Class              string `xml:"class,attr"`
				DaysToKeep         string `xml:"daysToKeep"`
				NumToKeep          string `xml:"numToKeep"`
				ArtifactDaysToKeep string `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep  string `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
		HudsonModelParametersDefinitionProperty struct {
			Text                 string `xml:",chardata"`
			ParameterDefinitions struct {
				Text                                                            string `xml:",chardata"`
				NetUazniaLukanusHudsonPluginsGitparameterGitParameterDefinition struct {
					Text               string `xml:",chardata"`
					Plugin             string `xml:"plugin,attr"`
					Name               string `xml:"name"`
					Description        string `xml:"description"`
					Uuid               string `xml:"uuid"`
					Type               string `xml:"type"`
					Branch             string `xml:"branch"`
					TagFilter          string `xml:"tagFilter"`
					BranchFilter       string `xml:"branchFilter"`
					SortMode           string `xml:"sortMode"`
					DefaultValue       string `xml:"defaultValue"`
					SelectedValue      string `xml:"selectedValue"`
					QuickFilterEnabled string `xml:"quickFilterEnabled"`
					ListSize           string `xml:"listSize"`
				} `xml:"net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition"`
				HudsonModelStringParameterDefinition []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name"`
					Description  string `xml:"description"`
					DefaultValue string `xml:"defaultValue"`
					Trim         string `xml:"trim"`
				} `xml:"hudson.model.StringParameterDefinition"`
			} `xml:"parameterDefinitions"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
	Scm struct {
		Text              string `xml:",chardata"`
		Class             string `xml:"class,attr"`
		Plugin            string `xml:"plugin,attr"`
		ConfigVersion     string `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text          string `xml:",chardata"`
				URL           string `xml:"url"`
				CredentialsId string `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name string `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations string `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg                      struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions string `xml:"extensions"`
	} `xml:"scm"`
	CanRoam                          string `xml:"canRoam"`
	Disabled                         string `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string `xml:"blockBuildWhenUpstreamBuilding"`
	Jdk                              string `xml:"jdk"`
	Triggers                         string `xml:"triggers"`
	ConcurrentBuild                  string `xml:"concurrentBuild"`
	Goals                            string `xml:"goals"`
	AggregatorStyleBuild             string `xml:"aggregatorStyleBuild"`
	IncrementalBuild                 string `xml:"incrementalBuild"`
	IgnoreUpstremChanges             string `xml:"ignoreUpstremChanges"`
	IgnoreUnsuccessfulUpstreams      string `xml:"ignoreUnsuccessfulUpstreams"`
	ArchivingDisabled                string `xml:"archivingDisabled"`
	SiteArchivingDisabled            string `xml:"siteArchivingDisabled"`
	FingerprintingDisabled           string `xml:"fingerprintingDisabled"`
	ResolveDependencies              string `xml:"resolveDependencies"`
	ProcessPlugins                   string `xml:"processPlugins"`
	MavenValidationLevel             string `xml:"mavenValidationLevel"`
	RunHeadless                      string `xml:"runHeadless"`
	DisableTriggerDownstreamProjects string `xml:"disableTriggerDownstreamProjects"`
	BlockTriggerWhenBuilding         string `xml:"blockTriggerWhenBuilding"`
	Settings                         struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
	} `xml:"settings"`
	GlobalSettings struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
	} `xml:"globalSettings"`
	Reporters  string `xml:"reporters"`
	Publishers struct {
		Text                                            string `xml:",chardata"`
		OrgJenkinsciPluginsQywechatQyWechatNotification struct {
			Text            string `xml:",chardata"`
			Plugin          string `xml:"plugin,attr"`
			WebhookUrl      string `xml:"webhookUrl"`
			MentionedId     string `xml:"mentionedId"`
			MentionedMobile string `xml:"mentionedMobile"`
			FailNotify      string `xml:"failNotify"`
			FailSend        string `xml:"failSend"`
			SuccessSend     string `xml:"successSend"`
			AboutSend       string `xml:"aboutSend"`
			UnstableSend    string `xml:"unstableSend"`
			StartBuild      string `xml:"startBuild"`
		} `xml:"org.jenkinsci.plugins.qywechat.QyWechatNotification"`
	} `xml:"publishers"`
	BuildWrappers struct {
		Text                                            string `xml:",chardata"`
		HudsonPluginsTimestamperTimestamperBuildWrapper struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"hudson.plugins.timestamper.TimestamperBuildWrapper"`
	} `xml:"buildWrappers"`
	Prebuilders          string `xml:"prebuilders"`
	Postbuilders         string `xml:"postbuilders"`
	RunPostStepsIfResult struct {
		Text          string `xml:",chardata"`
		Name          string `xml:"name"`
		Ordinal       string `xml:"ordinal"`
		Color         string `xml:"color"`
		CompleteBuild string `xml:"completeBuild"`
	} `xml:"runPostStepsIfResult"`
}
