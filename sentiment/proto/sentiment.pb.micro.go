// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/sentiment.proto

package sentiment

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

// Api Endpoints for Sentiment service

func NewSentimentEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Sentiment service

type SentimentService interface {
	Analyze(ctx context.Context, in *AnalyzeRequest, opts ...client.CallOption) (*AnalyzeResponse, error)
}

type sentimentService struct {
	c    client.Client
	name string
}

func NewSentimentService(name string, c client.Client) SentimentService {
	return &sentimentService{
		c:    c,
		name: name,
	}
}

func (c *sentimentService) Analyze(ctx context.Context, in *AnalyzeRequest, opts ...client.CallOption) (*AnalyzeResponse, error) {
	req := c.c.NewRequest(c.name, "Sentiment.Analyze", in)
	out := new(AnalyzeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sentiment service

type SentimentHandler interface {
	Analyze(context.Context, *AnalyzeRequest, *AnalyzeResponse) error
}

func RegisterSentimentHandler(s server.Server, hdlr SentimentHandler, opts ...server.HandlerOption) error {
	type sentiment interface {
		Analyze(ctx context.Context, in *AnalyzeRequest, out *AnalyzeResponse) error
	}
	type Sentiment struct {
		sentiment
	}
	h := &sentimentHandler{hdlr}
	return s.Handle(s.NewHandler(&Sentiment{h}, opts...))
}

type sentimentHandler struct {
	SentimentHandler
}

func (h *sentimentHandler) Analyze(ctx context.Context, in *AnalyzeRequest, out *AnalyzeResponse) error {
	return h.SentimentHandler.Analyze(ctx, in, out)
}
