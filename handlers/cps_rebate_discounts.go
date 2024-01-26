package handlers

import (
	"context"
	"howell/howell_rpc/kitex_gen/coder/hao/howell_rpc"
	models2 "howell/howell_rpc/kitex_gen/models"
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

// MGetCpsRebateDiscounts get cps rebate discounts
func MGetCpsRebateDiscounts(ctx context.Context, req *howell_rpc.MGetCpsRebateDiscountsRequest) (resp *howell_rpc.MGetCpsRebateDiscountsResponse, err error) {
	resp = &howell_rpc.MGetCpsRebateDiscountsResponse{
		EntityMap: make(map[string]*models2.CpsRebateDiscounts),
	}

	entities, err := service_cps_rebate_discounts.MGetCpsRebateDiscountsByID(ctx, req.GetEntityIdList())
	if err != nil {
		return resp, err
	}

	for _, entity := range entities {
		resp.EntityMap[entity.ID] = entity.ToRPC()
	}

	return resp, nil
}
