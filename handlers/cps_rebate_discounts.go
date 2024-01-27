package handlers

import (
	"context"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg/entity"
	"howell/howell_rpc/kitex_gen/base"
	"howell/howell_rpc/kitex_gen/coder/hao/howell_rpc"
	"howell/howell_rpc/kitex_gen/models"
	service_cps_rebate_discounts "howell/howell_rpc/service/cps_rebate_discounts"
)

// CreateCpsRebateDiscounts create crd
func CreateCpsRebateDiscounts(ctx context.Context, req *howell_rpc.CreateCpsRebateDiscountsRequest) (resp *howell_rpc.CreateCpsRebateDiscountsResponse, err error) {
	resp = &howell_rpc.CreateCpsRebateDiscountsResponse{}

	entityInfo := entity.NewCpsRebateDiscounts().FromRPC(req.CRDEntity)
	entityID, err := service_cps_rebate_discounts.Create(ctx, entityInfo)
	if err != nil {
		return resp, err
	}

	resp.EntityId = &entityID
	return resp, nil
}

// UpdateCpsRebateDiscounts update cps_crd
func UpdateCpsRebateDiscounts(ctx context.Context, req *howell_rpc.UpdateCpsRebateDiscountsRequest) (resp *howell_rpc.UpdateCpsRebateDiscountsResponse, err error) {
	resp = &howell_rpc.UpdateCpsRebateDiscountsResponse{}

	err = service_cps_rebate_discounts.Update(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// MGetCpsRebateDiscounts get cps rebate discounts
func MGetCpsRebateDiscounts(ctx context.Context, req *howell_rpc.MGetCpsRebateDiscountsRequest) (resp *howell_rpc.MGetCpsRebateDiscountsResponse, err error) {
	resp = &howell_rpc.MGetCpsRebateDiscountsResponse{
		EntityMap: make(map[string]*models.CpsRebateDiscounts),
		BaseResp:  base.NewBaseResp(),
	}

	entities, err := service_cps_rebate_discounts.MGetCpsRebateDiscountsByID(ctx, req.GetEntityIdList())
	if err != nil {
		return resp, err
	}

	for _, item := range entities {
		resp.EntityMap[item.ID] = item.ToRPC()
	}

	return resp, nil
}

// QueryCpsRebateDiscounts query cps_crd
func QueryCpsRebateDiscounts(ctx context.Context, req *howell_rpc.QueryCpsRebateDiscountsRequest) (resp *howell_rpc.QueryCpsRebateDiscountsResponse, err error) {
	resp = &howell_rpc.QueryCpsRebateDiscountsResponse{
		BaseResp: base.NewBaseResp(),
	}

	entities, pagination, err := service_cps_rebate_discounts.QueryCpsRebateDiscounts(ctx, req)
	if err != nil {
		return resp, err
	}

	resp.ItemList = entities
	resp.Pagination = pagination
	return resp, nil
}
