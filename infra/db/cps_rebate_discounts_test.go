package db

import (
	"context"
	"howell/howell_rpc/models"
	"testing"
)

func TestCreateCpsRebateDiscounts(t *testing.T) {
	Init()
	ctx := context.Background()
	err := CreateCpsRebateDiscounts(ctx, &models.CpsRebateDiscounts{
		ID: 1,
	})
	if err != nil {
		t.Fatalf("sql err:%v", err)
	}
}
