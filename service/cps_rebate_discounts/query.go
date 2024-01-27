package cps_rebate_discounts

import (
	"context"
	"fmt"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg/entity"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg/filter"
	"howell/howell_rpc/kitex_gen/coder/hao/howell_rpc"
	"howell/howell_rpc/kitex_gen/common"
	"howell/howell_rpc/kitex_gen/models"
)

// MGetCpsRebateDiscountsByID get cps rebate discounts
func MGetCpsRebateDiscountsByID(ctx context.Context, ids []string) ([]*entity.CpsRebateDiscounts, error) {
	res, err := cps_rebate_discounts_agg.CpsRebateDiscountsAgg().GetCpsRebateDiscountsByID(ctx, ids)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// QueryCpsRebateDiscounts query cps_crd
func QueryCpsRebateDiscounts(ctx context.Context, req *howell_rpc.QueryCpsRebateDiscountsRequest) ([]*models.CpsRebateDiscounts, *common.Pagination, error) {
	if req.GetPageIndex() < 1 {
		return nil, nil, fmt.Errorf("页数下标不能小于1")
	}
	if req.GetPageSize() > 100 {
		return nil, nil, fmt.Errorf("一页不能超过100个")
	}

	queryFilter := buildQueryFilter(ctx, req)
	var limit int32 = req.GetPageSize()
	var offset int32 = (req.GetPageIndex() - 1) * limit
	entities, total, err := cps_rebate_discounts_agg.CpsRebateDiscountsAgg().QueryCpsRebateDiscountsByID(ctx, queryFilter, offset, limit)
	if err != nil {
		return nil, nil, err
	}

	rpcEntityList := make([]*models.CpsRebateDiscounts, 0)
	for _, item := range entities {
		rpcEntityList = append(rpcEntityList, item.ToRPC())
	}
	pagination := &common.Pagination{
		PageIndex:  req.PageIndex,
		PageSize:   req.PageSize,
		TotalCount: int32(total),
	}
	return rpcEntityList, pagination, nil
}

func buildQueryFilter(ctx context.Context, req *howell_rpc.QueryCpsRebateDiscountsRequest) *filter.QueryCRTFilter {
	queryFilter := filter.NewQueryCRTFilter()
	queryFilter.SetIDList(req.GetEntityIdList())
	return queryFilter
}
