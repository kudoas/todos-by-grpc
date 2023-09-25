package main

import (
	"context"
	"sync"
	"testing"

	"connectrpc.com/connect"
	todov1 "github.com/kudoas/todos-by-grpc/gen/todo/v1"
)

func TestDeleteTask(t *testing.T) {
	cases := []struct {
		id      string
		isError bool
	}{
		{
			id:      "exist-id",
			isError: false,
		},
		{
			id:      "not-exist-id",
			isError: true,
		},
	}
	taskServer := &TaskServer{
		store: sync.Map{},
	}
	taskServer.store.Store("exist-id", Task{
		Id:      "exist-id",
		Title:   "sample title",
		Content: "sample task",
	})
	for _, c := range cases {
		request := &connect.Request[todov1.DeleteTaskRequest]{
			Msg: &todov1.DeleteTaskRequest{
				Id: c.id,
			},
		}
		response, err := taskServer.DeleteTask(context.TODO(), request)
		if c.isError && err == nil {
			t.Error("error is expected, but nil")
		}
		if !c.isError && response.Msg.Id != "" {
			t.Error("response.Msg.Id is not empty")
		}
	}
}
