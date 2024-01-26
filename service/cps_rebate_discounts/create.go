package cps_rebate_discounts

import (
	"context"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg/entity"
	"howell/howell_rpc/log"
	"howell/howell_rpc/util"
)

func Create(ctx context.Context, entityInfo *entity.CpsRebateDiscounts) (string, error) {
	// 没有传入ID自动生成
	if entityInfo.ID == "" {
		entityID, err := util.SecureRandId(8)
		if err != nil {
			log.NewPrefixLogger("server_cps_rebate_discounts").Error("secure rand id error, err:%v", err)
			return "", err
		}
		entityInfo.ID = entityID
	}

	if err := entityInfo.Check(); err != nil {
		_ = log.NewPrefixLogger("server_cps_rebate_discounts").Error("params check error, err:%v", err)
		return "", err
	}

	if err := cps_rebate_discounts_agg.CpsRebateDiscountsAgg().CreateCpsRebateDiscounts(ctx, entityInfo); err != nil {
		_ = log.NewPrefixLogger("server_cps_rebate_discounts").Error("db.CreateCpsRebateDiscounts error, err:%v", err)
		return "", err
	}

	return entityInfo.ID, nil
}
