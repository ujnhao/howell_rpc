package repo

import (
	"context"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg/entity"
	"howell/howell_rpc/infra/db"
	"testing"
)

func TestCreateCpsRebateDiscounts(t *testing.T) {
	db.Init()
	ctx := context.Background()
	err := NewCpsRebateDiscountsAggImpl().CreateCpsRebateDiscounts(ctx, &entity.CpsRebateDiscounts{
		ID: "1",
	})
	if err != nil {
		t.Fatalf("sql err:%v", err)
	}
}
