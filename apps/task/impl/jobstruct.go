package impl

import "encoding/xml"

type Project struct {
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
				HudsonModelStringParameterDefinition struct {
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
