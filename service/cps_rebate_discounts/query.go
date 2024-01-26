package cps_rebate_discounts

import (
	"context"
	"howell/howell_rpc/infra/db"
	"howell/howell_rpc/models"
)

// MGetCpsRebateDiscounts get cps rebate discounts
func MGetCpsRebateDiscountsByID(ctx context.Context, ids []string) ([]*models.CpsRebateDiscounts, error) {
	res, err := db.GetCpsRebateDiscountsByID(ctx, ids)
	if err != nil {
		return nil, err
	}
	return res, nil
}
