// Code generated by Kitex v0.8.0. DO NOT EDIT.

package howellrpcservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	howell_rpc "howell/howell_rpc/kitex_gen/coder/hao/howell_rpc"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateCpsRebateDiscounts(ctx context.Context, req *howell_rpc.CreateCpsRebateDiscountsRequest, callOptions ...callopt.Option) (r *howell_rpc.CreateCpsRebateDiscountsResponse, err error)
	MGetCpsRebateDiscounts(ctx context.Context, req *howell_rpc.MGetCpsRebateDiscountsRequest, callOptions ...callopt.Option) (r *howell_rpc.MGetCpsRebateDiscountsResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kHowellRpcServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kHowellRpcServiceClient struct {
	*kClient
}

func (p *kHowellRpcServiceClient) CreateCpsRebateDiscounts(ctx context.Context, req *howell_rpc.CreateCpsRebateDiscountsRequest, callOptions ...callopt.Option) (r *howell_rpc.CreateCpsRebateDiscountsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateCpsRebateDiscounts(ctx, req)
}

func (p *kHowellRpcServiceClient) MGetCpsRebateDiscounts(ctx context.Context, req *howell_rpc.MGetCpsRebateDiscountsRequest, callOptions ...callopt.Option) (r *howell_rpc.MGetCpsRebateDiscountsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetCpsRebateDiscounts(ctx, req)
}
