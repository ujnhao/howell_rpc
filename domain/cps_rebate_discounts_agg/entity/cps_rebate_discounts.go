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
	res := &CpsRebateDiscounts{
		ID: r.GetId(),
	}
	if r.AppId != nil {
		res.AppID = r.GetAppId()
	}
	if r.Name != nil {
		res.Name = r.GetName()
	}
	if r.CpsType != nil {
		res.CpsType = r.GetCpsType()
	}
	if r.ActType != nil {
		res.ActType = r.GetActType()
	}
	if r.ActUrl != nil {
		res.ActUrl = r.GetActUrl()
	}
	if r.Images != nil {
		res.Images = r.GetImages()
	}
	if r.Resource != nil {
		res.Resource = r.GetResource()
	}
	if r.Extra != nil {
		res.Extra = r.GetExtra()
	}
	return res
}

func (m *CpsRebateDiscounts) ToRPC() *models.CpsRebateDiscounts {
	return &models.CpsRebateDiscounts{
		Id:       &m.ID,
		AppId:    &m.AppID,
		Name:     &m.Name,
		CpsType:  (*common.CpsType)(&m.CpsType),
		ActType:  (*common.ActType)(&m.ActType),
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
