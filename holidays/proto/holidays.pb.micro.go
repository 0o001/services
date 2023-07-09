// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/holidays.proto

package holidays

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

// Api Endpoints for Holidays service

func NewHolidaysEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Holidays service

type HolidaysService interface {
	Countries(ctx context.Context, in *CountriesRequest, opts ...client.CallOption) (*CountriesResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type holidaysService struct {
	c    client.Client
	name string
}

func NewHolidaysService(name string, c client.Client) HolidaysService {
	return &holidaysService{
		c:    c,
		name: name,
	}
}

func (c *holidaysService) Countries(ctx context.Context, in *CountriesRequest, opts ...client.CallOption) (*CountriesResponse, error) {
	req := c.c.NewRequest(c.name, "Holidays.Countries", in)
	out := new(CountriesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidaysService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Holidays.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Holidays service

type HolidaysHandler interface {
	Countries(context.Context, *CountriesRequest, *CountriesResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterHolidaysHandler(s server.Server, hdlr HolidaysHandler, opts ...server.HandlerOption) error {
	type holidays interface {
		Countries(ctx context.Context, in *CountriesRequest, out *CountriesResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
	}
	type Holidays struct {
		holidays
	}
	h := &holidaysHandler{hdlr}
	return s.Handle(s.NewHandler(&Holidays{h}, opts...))
}

type holidaysHandler struct {
	HolidaysHandler
}

func (h *holidaysHandler) Countries(ctx context.Context, in *CountriesRequest, out *CountriesResponse) error {
	return h.HolidaysHandler.Countries(ctx, in, out)
}

func (h *holidaysHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.HolidaysHandler.List(ctx, in, out)
}
