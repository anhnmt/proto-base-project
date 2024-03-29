// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: auth/v1/auth.proto

package authv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/xdorro/proto-base-project/proto-gen-go/auth/v1"
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
	// AuthServiceName is the fully-qualified name of the AuthService service.
	AuthServiceName = "auth.v1.AuthService"
)

// AuthServiceClient is a client for the auth.v1.AuthService service.
type AuthServiceClient interface {
	Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.TokenResponse], error)
	RevokeToken(context.Context, *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.CommonResponse], error)
	RefreshToken(context.Context, *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.TokenResponse], error)
}

// NewAuthServiceClient constructs a client for the auth.v1.AuthService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authServiceClient{
		login: connect_go.NewClient[v1.LoginRequest, v1.TokenResponse](
			httpClient,
			baseURL+"/auth.v1.AuthService/Login",
			opts...,
		),
		revokeToken: connect_go.NewClient[v1.TokenRequest, v1.CommonResponse](
			httpClient,
			baseURL+"/auth.v1.AuthService/RevokeToken",
			opts...,
		),
		refreshToken: connect_go.NewClient[v1.TokenRequest, v1.TokenResponse](
			httpClient,
			baseURL+"/auth.v1.AuthService/RefreshToken",
			opts...,
		),
	}
}

// authServiceClient implements AuthServiceClient.
type authServiceClient struct {
	login        *connect_go.Client[v1.LoginRequest, v1.TokenResponse]
	revokeToken  *connect_go.Client[v1.TokenRequest, v1.CommonResponse]
	refreshToken *connect_go.Client[v1.TokenRequest, v1.TokenResponse]
}

// Login calls auth.v1.AuthService.Login.
func (c *authServiceClient) Login(ctx context.Context, req *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.TokenResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// RevokeToken calls auth.v1.AuthService.RevokeToken.
func (c *authServiceClient) RevokeToken(ctx context.Context, req *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return c.revokeToken.CallUnary(ctx, req)
}

// RefreshToken calls auth.v1.AuthService.RefreshToken.
func (c *authServiceClient) RefreshToken(ctx context.Context, req *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.TokenResponse], error) {
	return c.refreshToken.CallUnary(ctx, req)
}

// AuthServiceHandler is an implementation of the auth.v1.AuthService service.
type AuthServiceHandler interface {
	Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.TokenResponse], error)
	RevokeToken(context.Context, *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.CommonResponse], error)
	RefreshToken(context.Context, *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.TokenResponse], error)
}

// NewAuthServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthServiceHandler(svc AuthServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/auth.v1.AuthService/Login", connect_go.NewUnaryHandler(
		"/auth.v1.AuthService/Login",
		svc.Login,
		opts...,
	))
	mux.Handle("/auth.v1.AuthService/RevokeToken", connect_go.NewUnaryHandler(
		"/auth.v1.AuthService/RevokeToken",
		svc.RevokeToken,
		opts...,
	))
	mux.Handle("/auth.v1.AuthService/RefreshToken", connect_go.NewUnaryHandler(
		"/auth.v1.AuthService/RefreshToken",
		svc.RefreshToken,
		opts...,
	))
	return "/auth.v1.AuthService/", mux
}

// UnimplementedAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthServiceHandler struct{}

func (UnimplementedAuthServiceHandler) Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.TokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("auth.v1.AuthService.Login is not implemented"))
}

func (UnimplementedAuthServiceHandler) RevokeToken(context.Context, *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("auth.v1.AuthService.RevokeToken is not implemented"))
}

func (UnimplementedAuthServiceHandler) RefreshToken(context.Context, *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.TokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("auth.v1.AuthService.RefreshToken is not implemented"))
}
