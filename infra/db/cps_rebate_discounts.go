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

func GetCpsRebateDiscountsByID(ctx context.Context, ids []string) ([]*models.CpsRebateDiscounts, error) {
	res := make([]*models.CpsRebateDiscounts, 0)
	if len(ids) == 0 {
		return res, nil
	}

	err := getDB().Table("cps_rebate_discounts").Where("id in (?)", ids).Find(&res).Error
	if err != nil {
		log.NewPrefixLogger("db").Error("CreateCpsRebateDiscounts sql failed, err=%v", err)
		return nil, err
	}
	return res, nil
}
