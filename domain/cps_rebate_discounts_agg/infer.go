package cps_rebate_discounts_agg

import (
	"context"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg/entity"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg/filter"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg/repo"
)

type ICpsRebateDiscountsAgg interface {
	CreateCpsRebateDiscounts(ctx context.Context, entity *entity.CpsRebateDiscounts) error
	GetCpsRebateDiscountsByID(ctx context.Context, ids []string) ([]*entity.CpsRebateDiscounts, error)
	QueryCpsRebateDiscountsByID(ctx context.Context, queryFilter *filter.QueryCRTFilter, offset, limit int32) (res []*entity.CpsRebateDiscounts, total int64, err error)
}

func CpsRebateDiscountsAgg() ICpsRebateDiscountsAgg {
	return repo.NewCpsRebateDiscountsAggImpl()
}
