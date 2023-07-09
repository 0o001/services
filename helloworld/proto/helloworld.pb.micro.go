// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/helloworld.proto

package helloworld

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "micro.dev/v4/service/api"
	client "micro.dev/v4/service/client"
	server "micro.dev/v4/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Helloworld service

func NewHelloworldEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Helloworld service

type HelloworldService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (Helloworld_StreamService, error)
}

type helloworldService struct {
	c    client.Client
	name string
}

func NewHelloworldService(name string, c client.Client) HelloworldService {
	return &helloworldService{
		c:    c,
		name: name,
	}
}

func (c *helloworldService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Helloworld.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldService) Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (Helloworld_StreamService, error) {
	req := c.c.NewRequest(c.name, "Helloworld.Stream", &StreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &helloworldServiceStream{stream}, nil
}

type Helloworld_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamResponse, error)
}

type helloworldServiceStream struct {
	stream client.Stream
}

func (x *helloworldServiceStream) Close() error {
	return x.stream.Close()
}

func (x *helloworldServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloworldServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloworldServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloworldServiceStream) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Helloworld service

type HelloworldHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	Stream(context.Context, *StreamRequest, Helloworld_StreamStream) error
}

func RegisterHelloworldHandler(s server.Server, hdlr HelloworldHandler, opts ...server.HandlerOption) error {
	type helloworld interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		Stream(ctx context.Context, stream server.Stream) error
	}
	type Helloworld struct {
		helloworld
	}
	h := &helloworldHandler{hdlr}
	return s.Handle(s.NewHandler(&Helloworld{h}, opts...))
}

type helloworldHandler struct {
	HelloworldHandler
}

func (h *helloworldHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.HelloworldHandler.Call(ctx, in, out)
}

func (h *helloworldHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.HelloworldHandler.Stream(ctx, m, &helloworldStreamStream{stream})
}

type Helloworld_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamResponse) error
}

type helloworldStreamStream struct {
	stream server.Stream
}

func (x *helloworldStreamStream) Close() error {
	return x.stream.Close()
}

func (x *helloworldStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloworldStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloworldStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloworldStreamStream) Send(m *StreamResponse) error {
	return x.stream.Send(m)
}
