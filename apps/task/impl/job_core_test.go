package impl

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/tqtcloud/workflow-backend/apps/task"
	"github.com/tqtcloud/workflow-backend/conf"
	"log"
	"testing"
)

var jenkins *gojenkins.Jenkins

func Test_getJobConfig(t *testing.T) {
	req := task.DescribeTaskRequest{
		Env:     task.JenkinsEnv_ENV,
		Jobname: "apitest0927",
	}
	inse, err := getJobConfig(context.Background(), &req, jenkins)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(inse)
}

func init() {
	err := conf.LoadConfigFromToml("D:\\project-git\\workflow-backend\\etc\\config.toml")
	if err != nil {
		log.Println(err)
	}
	c := conf.C()
	jenk, err := envDecision(context.Background(), task.JenkinsEnv_ENV, c)
	if err != nil {
		log.Println(err)
	}
	jenkins = jenk
}
