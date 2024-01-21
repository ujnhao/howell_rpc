package main

import (
	"context"
	howell_rpc "howell/howell_rpc/kitex_gen/coder/hao/howell_rpc"
)

// HowellRpcServiceImpl implements the last service interface defined in the IDL.
type HowellRpcServiceImpl struct{}

// CreateCpsRebateDiscounts implements the HowellRpcServiceImpl interface.
func (s *HowellRpcServiceImpl) CreateCpsRebateDiscounts(ctx context.Context, req *howell_rpc.CreateCpsRebateDiscountsRequest) (resp *howell_rpc.CreateCpsRebateDiscountsResponse, err error) {
	// TODO: Your code here...
	return
}
