package task

import (
	"fmt"
	"testing"
)

func TestNewTask(t *testing.T) {
	req := CreateTaskRequest{
		Env:      0,
		JobName:  "xx",
		CreateBy: "tqt",
		Xml:      "xxx",
	}

	task, err := NewTask(&req)
	fmt.Println(task, err)
}
