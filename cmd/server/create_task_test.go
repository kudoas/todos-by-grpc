package main

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	todov1 "github.com/kudoas/todos-by-grpc/gen/todo/v1"
)

func TestCreateTask(t *testing.T) {
	taskServer := &TaskServer{}
	request := &connect.Request[todov1.CreateTaskRequest]{
		Msg: &todov1.CreateTaskRequest{
			Title:   "sample title",
			Content: "sample task",
		},
	}

	response, err := taskServer.CreateTask(context.TODO(), request)
	if err != nil {
		t.Fatal(err)
	}
	if response.Msg.Id == "" {
		t.Error("response.Msg.Id is empty")
	}
	if response.Msg.Title != "sample title" {
		t.Errorf("response.Msg.Title is expected %s, but %s", "sample title", response.Msg.Title)
	}
	if response.Msg.Content != "sample task" {
		t.Errorf("response.Msg.Content is expected %s, but %s", "sample task", response.Msg.Content)
	}
}
