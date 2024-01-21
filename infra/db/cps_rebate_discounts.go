package db

import (
	"context"
	"howell/howell_rpc/log"
	"howell/howell_rpc/models"
)

func CreateCpsRebateDiscounts(ctx context.Context,
	entity *models.CpsRebateDiscounts) error {
	err := getDB().Table("cps_rebate_discounts").Create(entity).Error
	if err != nil {
		log.NewPrefixLogger("db").Error("CreateCpsRebateDiscounts sql failed, err=%v", err)
		return err
	}
	return nil
}
