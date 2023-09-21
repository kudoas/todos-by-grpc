package main

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/go-playground/assert"
	todov1 "github.com/kudoas/todos-by-grpc/gen/todo/v1"
)

func TestCreateTask(t *testing.T) {
	taskServer := &TaskServer{}
	reqTask := &connect.Request[todov1.CreateTaskRequest]{
		Msg: &todov1.CreateTaskRequest{
			Title:   "sample title",
			Content: "sample task",
		},
	}

	resTask := &connect.Response[todov1.CreateTaskResponse]{
		Msg: &todov1.CreateTaskResponse{
			Id: 1,
			Title:   "sample title",
			Content: "sample task",
		},
		
	}
	actual, err := taskServer.CreateTask(context.TODO(),reqTask)
	expected := resTask
	assert.Equal(t,actual.Msg,expected.Msg)
	assert.Equal(t,err,nil)
}