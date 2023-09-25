// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: todo/v1/todo.proto

package todov1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/kudoas/todos-by-grpc/gen/todo/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// TaskServiceName is the fully-qualified name of the TaskService service.
	TaskServiceName = "todo.v1.TaskService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TaskServiceReadProcedure is the fully-qualified name of the TaskService's Read RPC.
	TaskServiceReadProcedure = "/todo.v1.TaskService/Read"
	// TaskServiceCreateProcedure is the fully-qualified name of the TaskService's Create RPC.
	TaskServiceCreateProcedure = "/todo.v1.TaskService/Create"
	// TaskServiceDeleteProcedure is the fully-qualified name of the TaskService's Delete RPC.
	TaskServiceDeleteProcedure = "/todo.v1.TaskService/Delete"
)

// TaskServiceClient is a client for the todo.v1.TaskService service.
type TaskServiceClient interface {
	Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error)
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewTaskServiceClient constructs a client for the todo.v1.TaskService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTaskServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TaskServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &taskServiceClient{
		read: connect.NewClient[v1.ReadRequest, v1.ReadResponse](
			httpClient,
			baseURL+TaskServiceReadProcedure,
			opts...,
		),
		create: connect.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+TaskServiceCreateProcedure,
			opts...,
		),
		delete: connect.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+TaskServiceDeleteProcedure,
			opts...,
		),
	}
}

// taskServiceClient implements TaskServiceClient.
type taskServiceClient struct {
	read   *connect.Client[v1.ReadRequest, v1.ReadResponse]
	create *connect.Client[v1.CreateRequest, v1.CreateResponse]
	delete *connect.Client[v1.DeleteRequest, v1.DeleteResponse]
}

// Read calls todo.v1.TaskService.Read.
func (c *taskServiceClient) Read(ctx context.Context, req *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error) {
	return c.read.CallUnary(ctx, req)
}

// Create calls todo.v1.TaskService.Create.
func (c *taskServiceClient) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Delete calls todo.v1.TaskService.Delete.
func (c *taskServiceClient) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// TaskServiceHandler is an implementation of the todo.v1.TaskService service.
type TaskServiceHandler interface {
	Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error)
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewTaskServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTaskServiceHandler(svc TaskServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	taskServiceReadHandler := connect.NewUnaryHandler(
		TaskServiceReadProcedure,
		svc.Read,
		opts...,
	)
	taskServiceCreateHandler := connect.NewUnaryHandler(
		TaskServiceCreateProcedure,
		svc.Create,
		opts...,
	)
	taskServiceDeleteHandler := connect.NewUnaryHandler(
		TaskServiceDeleteProcedure,
		svc.Delete,
		opts...,
	)
	return "/todo.v1.TaskService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TaskServiceReadProcedure:
			taskServiceReadHandler.ServeHTTP(w, r)
		case TaskServiceCreateProcedure:
			taskServiceCreateHandler.ServeHTTP(w, r)
		case TaskServiceDeleteProcedure:
			taskServiceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTaskServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTaskServiceHandler struct{}

func (UnimplementedTaskServiceHandler) Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.TaskService.Read is not implemented"))
}

func (UnimplementedTaskServiceHandler) Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.TaskService.Create is not implemented"))
}

func (UnimplementedTaskServiceHandler) Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.TaskService.Delete is not implemented"))
}
