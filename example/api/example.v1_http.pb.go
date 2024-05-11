// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.25.2
// source: example/api/example.v1.proto

package exampleapi

import (
	context "context"
	middleware "github.com/go-kratos/kratos/v2/middleware"
	selector "github.com/go-kratos/kratos/v2/middleware/selector"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type _ = middleware.Middleware
type _ = selector.Builder

const OperationExampleHelloWorld = "/github.ccheers.pggh.example.Example/HelloWorld"

type ExampleHTTPServer interface {
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldReply, error)
}

type ExampleHTTPServerMiddlewareConfig struct {
	Middleware1 middleware.Middleware
	Middleware2 middleware.Middleware
}

func NewExampleHTTPServerMiddleware(mc ExampleHTTPServerMiddlewareConfig) middleware.Middleware {
	return selector.Server(
		selector.Server(
			mc.Middleware1,
			mc.Middleware2,
		).Path(OperationExampleHelloWorld).Build(),
	).Path(
		OperationExampleHelloWorld,
	).Build()
}

func RegisterExampleHTTPServer(s *http.Server, srv ExampleHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/example/hello", _Example_HelloWorld0_HTTP_Handler(srv))
}

func _Example_HelloWorld0_HTTP_Handler(srv ExampleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloWorldRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExampleHelloWorld)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.HelloWorld(ctx, req.(*HelloWorldRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloWorldReply)
		return ctx.Result(200, reply)
	}
}

type ExampleHTTPClient interface {
	HelloWorld(ctx context.Context, req *HelloWorldRequest, opts ...http.CallOption) (rsp *HelloWorldReply, err error)
}

type ExampleHTTPClientImpl struct {
	cc *http.Client
}

func NewExampleHTTPClient(client *http.Client) ExampleHTTPClient {
	return &ExampleHTTPClientImpl{client}
}

func (c *ExampleHTTPClientImpl) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...http.CallOption) (*HelloWorldReply, error) {
	var out HelloWorldReply
	pattern := "/v1/example/hello"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationExampleHelloWorld))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
