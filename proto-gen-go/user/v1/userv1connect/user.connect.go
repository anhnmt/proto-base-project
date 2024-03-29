// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: user/v1/user.proto

package userv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/xdorro/proto-base-project/proto-gen-go/user/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "user.v1.UserService"
)

// UserServiceClient is a client for the user.v1.UserService service.
type UserServiceClient interface {
	FindAllUsers(context.Context, *connect_go.Request[v1.FindAllUsersRequest]) (*connect_go.Response[v1.FindAllUsersResponse], error)
	// Find User by ID
	FindUserByID(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.User], error)
	// Create new User
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CommonResponse], error)
	// Update User by ID
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.CommonResponse], error)
	// Delete User
	DeleteUser(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.CommonResponse], error)
}

// NewUserServiceClient constructs a client for the user.v1.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		findAllUsers: connect_go.NewClient[v1.FindAllUsersRequest, v1.FindAllUsersResponse](
			httpClient,
			baseURL+"/user.v1.UserService/FindAllUsers",
			opts...,
		),
		findUserByID: connect_go.NewClient[v1.CommonUUIDRequest, v1.User](
			httpClient,
			baseURL+"/user.v1.UserService/FindUserByID",
			opts...,
		),
		createUser: connect_go.NewClient[v1.CreateUserRequest, v1.CommonResponse](
			httpClient,
			baseURL+"/user.v1.UserService/CreateUser",
			opts...,
		),
		updateUser: connect_go.NewClient[v1.UpdateUserRequest, v1.CommonResponse](
			httpClient,
			baseURL+"/user.v1.UserService/UpdateUser",
			opts...,
		),
		deleteUser: connect_go.NewClient[v1.CommonUUIDRequest, v1.CommonResponse](
			httpClient,
			baseURL+"/user.v1.UserService/DeleteUser",
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	findAllUsers *connect_go.Client[v1.FindAllUsersRequest, v1.FindAllUsersResponse]
	findUserByID *connect_go.Client[v1.CommonUUIDRequest, v1.User]
	createUser   *connect_go.Client[v1.CreateUserRequest, v1.CommonResponse]
	updateUser   *connect_go.Client[v1.UpdateUserRequest, v1.CommonResponse]
	deleteUser   *connect_go.Client[v1.CommonUUIDRequest, v1.CommonResponse]
}

// FindAllUsers calls user.v1.UserService.FindAllUsers.
func (c *userServiceClient) FindAllUsers(ctx context.Context, req *connect_go.Request[v1.FindAllUsersRequest]) (*connect_go.Response[v1.FindAllUsersResponse], error) {
	return c.findAllUsers.CallUnary(ctx, req)
}

// FindUserByID calls user.v1.UserService.FindUserByID.
func (c *userServiceClient) FindUserByID(ctx context.Context, req *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.User], error) {
	return c.findUserByID.CallUnary(ctx, req)
}

// CreateUser calls user.v1.UserService.CreateUser.
func (c *userServiceClient) CreateUser(ctx context.Context, req *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// UpdateUser calls user.v1.UserService.UpdateUser.
func (c *userServiceClient) UpdateUser(ctx context.Context, req *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return c.updateUser.CallUnary(ctx, req)
}

// DeleteUser calls user.v1.UserService.DeleteUser.
func (c *userServiceClient) DeleteUser(ctx context.Context, req *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the user.v1.UserService service.
type UserServiceHandler interface {
	FindAllUsers(context.Context, *connect_go.Request[v1.FindAllUsersRequest]) (*connect_go.Response[v1.FindAllUsersResponse], error)
	// Find User by ID
	FindUserByID(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.User], error)
	// Create new User
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CommonResponse], error)
	// Update User by ID
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.CommonResponse], error)
	// Delete User
	DeleteUser(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.CommonResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/user.v1.UserService/FindAllUsers", connect_go.NewUnaryHandler(
		"/user.v1.UserService/FindAllUsers",
		svc.FindAllUsers,
		opts...,
	))
	mux.Handle("/user.v1.UserService/FindUserByID", connect_go.NewUnaryHandler(
		"/user.v1.UserService/FindUserByID",
		svc.FindUserByID,
		opts...,
	))
	mux.Handle("/user.v1.UserService/CreateUser", connect_go.NewUnaryHandler(
		"/user.v1.UserService/CreateUser",
		svc.CreateUser,
		opts...,
	))
	mux.Handle("/user.v1.UserService/UpdateUser", connect_go.NewUnaryHandler(
		"/user.v1.UserService/UpdateUser",
		svc.UpdateUser,
		opts...,
	))
	mux.Handle("/user.v1.UserService/DeleteUser", connect_go.NewUnaryHandler(
		"/user.v1.UserService/DeleteUser",
		svc.DeleteUser,
		opts...,
	))
	return "/user.v1.UserService/", mux
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) FindAllUsers(context.Context, *connect_go.Request[v1.FindAllUsersRequest]) (*connect_go.Response[v1.FindAllUsersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.FindAllUsers is not implemented"))
}

func (UnimplementedUserServiceHandler) FindUserByID(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.User], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.FindUserByID is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.CreateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.UpdateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteUser(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.DeleteUser is not implemented"))
}
