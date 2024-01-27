package entity

import (
	"fmt"
	"howell/howell_rpc/kitex_gen/common"
	"howell/howell_rpc/kitex_gen/models"
	"strings"
	"time"
)

type CpsRebateDiscounts struct {
	ID        string    `json:"id"`
	AppID     string    `json:"app_id"`
	Name      string    `json:"name"`
	CpsType   string    `json:"cps_type"`
	ActType   string    `json:"act_type"`
	ActUrl    string    `json:"act_url"`
	Images    string    `json:"images"`
	Resource  string    `json:"resource"`
	Status    int32     `json:"status"`
	Extra     string    `json:"extra"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCpsRebateDiscounts() *CpsRebateDiscounts {
	return &CpsRebateDiscounts{}
}

func (m *CpsRebateDiscounts) FromRPC(r *models.CpsRebateDiscounts) *CpsRebateDiscounts {
	return &CpsRebateDiscounts{
		ID:        r.GetId(),
		AppID:     r.GetAppId(),
		Name:      r.GetName(),
		CpsType:   r.GetCpsType(),
		ActType:   r.GetActTpye(),
		ActUrl:    r.GetActUrl(),
		Images:    r.GetImages(),
		Resource:  r.GetResource(),
		Status:    int32(r.GetStatus()),
		Extra:     r.GetExtra(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func (m *CpsRebateDiscounts) ToRPC() *models.CpsRebateDiscounts {
	return &models.CpsRebateDiscounts{
		Id:       &m.ID,
		AppId:    &m.AppID,
		Name:     &m.Name,
		CpsType:  (*common.CpsType)(&m.CpsType),
		ActTpye:  (*common.ActType)(&m.ActType),
		ActUrl:   &m.ActUrl,
		Images:   &m.Images,
		Resource: &m.Resource,
		Status:   common.StatusPtr(common.Status(m.Status)),
		Extra:    &m.Extra,
	}
}

func (m *CpsRebateDiscounts) Check() error {
	if m.Name == "" {
		return fmt.Errorf("invalid name")
	}
	if m.Resource == "" {
		return fmt.Errorf("invalid resource")
	}
	if m.CpsType != common.MeiTuan {
		return fmt.Errorf("invalid cps_type")
	}
	if m.ActType != common.MiniProgram &&
		m.ActType != common.H5 {
		return fmt.Errorf("act_type")
	}
	if m.ActUrl == "" ||
		len(strings.TrimSpace(m.ActUrl)) != len(m.ActUrl) {
		return fmt.Errorf("invalid act_url")
	}
	return nil
}
