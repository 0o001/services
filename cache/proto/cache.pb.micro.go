// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cache.proto

package cache

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

// Api Endpoints for Cache service

func NewCacheEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Cache service

type CacheService interface {
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...client.CallOption) (*SetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Increment(ctx context.Context, in *IncrementRequest, opts ...client.CallOption) (*IncrementResponse, error)
	Decrement(ctx context.Context, in *DecrementRequest, opts ...client.CallOption) (*DecrementResponse, error)
	ListKeys(ctx context.Context, in *ListKeysRequest, opts ...client.CallOption) (*ListKeysResponse, error)
}

type cacheService struct {
	c    client.Client
	name string
}

func NewCacheService(name string, c client.Client) CacheService {
	return &cacheService{
		c:    c,
		name: name,
	}
}

func (c *cacheService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "Cache.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheService) Set(ctx context.Context, in *SetRequest, opts ...client.CallOption) (*SetResponse, error) {
	req := c.c.NewRequest(c.name, "Cache.Set", in)
	out := new(SetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Cache.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheService) Increment(ctx context.Context, in *IncrementRequest, opts ...client.CallOption) (*IncrementResponse, error) {
	req := c.c.NewRequest(c.name, "Cache.Increment", in)
	out := new(IncrementResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheService) Decrement(ctx context.Context, in *DecrementRequest, opts ...client.CallOption) (*DecrementResponse, error) {
	req := c.c.NewRequest(c.name, "Cache.Decrement", in)
	out := new(DecrementResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheService) ListKeys(ctx context.Context, in *ListKeysRequest, opts ...client.CallOption) (*ListKeysResponse, error) {
	req := c.c.NewRequest(c.name, "Cache.ListKeys", in)
	out := new(ListKeysResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cache service

type CacheHandler interface {
	Get(context.Context, *GetRequest, *GetResponse) error
	Set(context.Context, *SetRequest, *SetResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Increment(context.Context, *IncrementRequest, *IncrementResponse) error
	Decrement(context.Context, *DecrementRequest, *DecrementResponse) error
	ListKeys(context.Context, *ListKeysRequest, *ListKeysResponse) error
}

func RegisterCacheHandler(s server.Server, hdlr CacheHandler, opts ...server.HandlerOption) error {
	type cache interface {
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Set(ctx context.Context, in *SetRequest, out *SetResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Increment(ctx context.Context, in *IncrementRequest, out *IncrementResponse) error
		Decrement(ctx context.Context, in *DecrementRequest, out *DecrementResponse) error
		ListKeys(ctx context.Context, in *ListKeysRequest, out *ListKeysResponse) error
	}
	type Cache struct {
		cache
	}
	h := &cacheHandler{hdlr}
	return s.Handle(s.NewHandler(&Cache{h}, opts...))
}

type cacheHandler struct {
	CacheHandler
}

func (h *cacheHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.CacheHandler.Get(ctx, in, out)
}

func (h *cacheHandler) Set(ctx context.Context, in *SetRequest, out *SetResponse) error {
	return h.CacheHandler.Set(ctx, in, out)
}

func (h *cacheHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.CacheHandler.Delete(ctx, in, out)
}

func (h *cacheHandler) Increment(ctx context.Context, in *IncrementRequest, out *IncrementResponse) error {
	return h.CacheHandler.Increment(ctx, in, out)
}

func (h *cacheHandler) Decrement(ctx context.Context, in *DecrementRequest, out *DecrementResponse) error {
	return h.CacheHandler.Decrement(ctx, in, out)
}

func (h *cacheHandler) ListKeys(ctx context.Context, in *ListKeysRequest, out *ListKeysResponse) error {
	return h.CacheHandler.ListKeys(ctx, in, out)
}
