package cps_rebate_discounts

import (
	"context"
	"fmt"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg/entity"
	"howell/howell_rpc/kitex_gen/coder/hao/howell_rpc"
	"howell/howell_rpc/log"
	"howell/howell_rpc/util"
)

func Create(ctx context.Context, entityInfo *entity.CpsRebateDiscounts) (string, error) {
	// 没有传入ID自动生成
	if entityInfo.ID == "" {
		entityID, err := util.SecureRandId(8)
		if err != nil {
			_ = log.NewPrefixLogger("server_cps_rebate_discounts").Error("secure rand id error, err:%v", err)
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

func Update(ctx context.Context, req *howell_rpc.UpdateCpsRebateDiscountsRequest) error {
	// 没有传入ID自动生成
	if req.GetEntityId() == "" {
		return fmt.Errorf("invalid entity_id")
	}
	if req.CRDEntity == nil {
		return fmt.Errorf("invalid crd entity")
	}

	items, err := cps_rebate_discounts_agg.CpsRebateDiscountsAgg().GetCpsRebateDiscountsByID(ctx, []string{req.GetEntityId()})
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return fmt.Errorf("not found crd")
	}

	updateEntity := items[0]
	req.CRDEntity.Id = &updateEntity.ID
	updateEntity = updateEntity.FromRPC(req.CRDEntity)
	if err := updateEntity.Check(); err != nil {
		_ = log.NewPrefixLogger("server_cps_rebate_discounts").Error("params check error, err:%v", err)
		return err
	}

	if err := cps_rebate_discounts_agg.CpsRebateDiscountsAgg().UpdateCpsRebateDiscountsByID(ctx, req.GetEntityId(), updateEntity); err != nil {
		_ = log.NewPrefixLogger("server_cps_rebate_discounts").Error("db.UpdateCpsRebateDiscountsByID error, err:%v", err)
		return err
	}

	return nil
}
