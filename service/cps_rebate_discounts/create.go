package cps_rebate_discounts

import (
	"context"
	"howell/howell_rpc/infra/db"
	"howell/howell_rpc/log"
	"howell/howell_rpc/models"
	"howell/howell_rpc/util"
)

func Create(ctx context.Context, entity *models.CpsRebateDiscounts) (string, error) {
	// 没有传入ID自动生成
	if entity.ID == "" {
		entityID, err := util.SecureRandId(8)
		if err != nil {
			log.NewPrefixLogger("server_cps_rebate_discounts").Error("secure rand id error, err:%v", err)
			return "", err
		}
		entity.ID = entityID
	}

	if err := entity.Check(); err != nil {
		log.NewPrefixLogger("server_cps_rebate_discounts").Error("params check error, err:%v", err)
		return "", err
	}

	if err := db.CreateCpsRebateDiscounts(ctx, entity); err != nil {
		log.NewPrefixLogger("server_cps_rebate_discounts").Error("db.CreateCpsRebateDiscounts error, err:%v", err)
		return "", err
	}

	return entity.ID, nil
}
