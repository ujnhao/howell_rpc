package main

import (
	"context"
	"howell/howell_rpc/handlers"
	howell_rpc "howell/howell_rpc/kitex_gen/coder/hao/howell_rpc"
)

// HowellRpcServiceImpl implements the last service interface defined in the IDL.
type HowellRpcServiceImpl struct{}

// CreateCpsRebateDiscounts create cps rebate discounts
func (s *HowellRpcServiceImpl) CreateCpsRebateDiscounts(ctx context.Context, req *howell_rpc.CreateCpsRebateDiscountsRequest) (resp *howell_rpc.CreateCpsRebateDiscountsResponse, err error) {
	return handlers.CreateCpsRebateDiscounts(ctx, req)
}

// MGetCpsRebateDiscounts get cps rebate discounts
func (s *HowellRpcServiceImpl) MGetCpsRebateDiscounts(ctx context.Context, req *howell_rpc.MGetCpsRebateDiscountsRequest) (resp *howell_rpc.MGetCpsRebateDiscountsResponse, err error) {
	return handlers.MGetCpsRebateDiscounts(ctx, req)
}
