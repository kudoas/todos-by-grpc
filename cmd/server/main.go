package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/google/uuid"
	todov1 "github.com/kudoas/todos-by-grpc/gen/todo/v1"        // generated by protoc-gen-go
	"github.com/kudoas/todos-by-grpc/gen/todo/v1/todov1connect" // generated by protoc-gen-connect-go
)

type TaskServer struct {
	store sync.Map
}

type Task struct {
	Id      string
	Title   string
	Content string
	Status  todov1.Status
}

func (s *TaskServer) CreateTask(
	ctx context.Context,
	req *connect.Request[todov1.CreateTaskRequest],
) (*connect.Response[todov1.CreateTaskResponse], error) {
	log.Println("Request headers: ", req.Header())

	uuid, _ := uuid.NewRandom()
	if task, _ := s.store.Load(uuid); task != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("Task.Id %s already exists", uuid.String()))
	}
	task := &Task{
		Id:      uuid.String(),
		Title:   req.Msg.Title,
		Content: req.Msg.Content,
		Status:  todov1.Status_STATUS_DOING,
	}
	s.store.Store(task.Id, task)
	res := connect.NewResponse(&todov1.CreateTaskResponse{
		Id:      task.Id,
		Title:   task.Title,
		Content: task.Content,
	})
	res.Header().Set("Todo-Version", "v1")
	stored, _ := s.store.Load(task.Id)
	log.Println("stored", stored)

	return res, nil
}

func main() {
	task := &TaskServer{}
	mux := http.NewServeMux()
	path, handler := todov1connect.NewTaskServiceHandler(task)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
