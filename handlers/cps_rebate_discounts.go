package handlers

import (
	"context"
	"howell/howell_rpc/kitex_gen/coder/hao/howell_rpc"
	"howell/howell_rpc/models"
	service_cps_rebate_discounts "howell/howell_rpc/service/cps_rebate_discounts"
)

// CreateCpsRebateDiscounts implements the HowellRpcServiceImpl interface.
func CreateCpsRebateDiscounts(ctx context.Context, req *howell_rpc.CreateCpsRebateDiscountsRequest) (resp *howell_rpc.CreateCpsRebateDiscountsResponse, err error) {
	resp = &howell_rpc.CreateCpsRebateDiscountsResponse{}

	entity := models.NewCpsRebateDiscounts().FromRPC(req.CRDEntity)
	entityID, err := service_cps_rebate_discounts.Create(ctx, entity)
	if err != nil {
		return resp, err
	}

	resp.EntityId = &entityID
	return resp, nil
}
