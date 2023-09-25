package main

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	todov1 "github.com/kudoas/todos-by-grpc/gen/todo/v1"
)

func TestCreateTask(t *testing.T) {
	taskServer := &TaskServer{}
	request := &connect.Request[todov1.CreateRequest]{
		Msg: &todov1.CreateRequest{
			Title:   "sample title",
			Content: "sample task",
		},
	}
	response, err := taskServer.Create(context.TODO(), request)
	if err != nil {
		t.Fatal(err)
	}
	if response.Msg.Task.Id == "" {
		t.Error("response.Msg.Id is empty")
	}
	if response.Msg.Task.Title != "sample title" {
		t.Errorf("response.Msg.Title is expected %s, but %s", "sample title", response.Msg.Task.Title)
	}
	if response.Msg.Task.Content != "sample task" {
		t.Errorf("response.Msg.Content is expected %s, but %s", "sample task", response.Msg.Task.Content)
	}
}
