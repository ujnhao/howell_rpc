package repo

import (
	"context"
	"fmt"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg/entity"
	"howell/howell_rpc/domain/cps_rebate_discounts_agg/filter"
	"howell/howell_rpc/infra/db"
	"howell/howell_rpc/log"
)

func NewCpsRebateDiscountsAggImpl() *CpsRebateDiscountsAggImpl {
	return &CpsRebateDiscountsAggImpl{
		Logger: log.NewPrefixLogger("CpsRebateDiscountsAggImpl"),
	}
}

type CpsRebateDiscountsAggImpl struct {
	Logger log.Logger
}

func (l *CpsRebateDiscountsAggImpl) CreateCpsRebateDiscounts(ctx context.Context, entityInfo *entity.CpsRebateDiscounts) error {
	err := db.GetDB().Debug().Model(&entity.CpsRebateDiscounts{}).Create(entityInfo).Error
	if err != nil {
		_ = l.Logger.Error("sql CreateCpsRebateDiscounts failed, err=%v", err)
		return err
	}
	return nil
}

func (l *CpsRebateDiscountsAggImpl) GetCpsRebateDiscountsByID(ctx context.Context, ids []string) ([]*entity.CpsRebateDiscounts, error) {
	res := make([]*entity.CpsRebateDiscounts, 0)
	if len(ids) == 0 {
		_ = l.Logger.Warn("empty ids..., return")
		return res, nil
	}

	err := db.GetDB().Debug().Model(&entity.CpsRebateDiscounts{}).Where("id in (?)", ids).Find(&res).Error
	if err != nil {
		_ = l.Logger.Error("sql GetCpsRebateDiscounts failed, err=%v", err)
		return nil, err
	}
	return res, nil
}

func (l *CpsRebateDiscountsAggImpl) QueryCpsRebateDiscountsByID(ctx context.Context, queryFilter *filter.QueryCRTFilter, offset, limit int32) (res []*entity.CpsRebateDiscounts, total int64, err error) {
	res = make([]*entity.CpsRebateDiscounts, 0)
	if offset < 0 {
		_ = l.Logger.Warn("index不能小于0")
		return res, total, fmt.Errorf("index不能小于0")
	}

	tx := db.GetDB().Debug().Model(&entity.CpsRebateDiscounts{})
	if queryFilter != nil {
		tx = queryFilter.GenQuery(tx)
	}

	err = tx.Count(&total).Error
	if err != nil {
		_ = l.Logger.Error("sql tx.count error, err=%v", err)
		return res, total, err
	}

	err = tx.Offset(int(offset)).Limit(int(limit)).Find(&res).Error
	if err != nil {
		_ = l.Logger.Error("sql offset limit find failed, err=%v", err)
		return res, total, err
	}
	return res, total, nil
}
