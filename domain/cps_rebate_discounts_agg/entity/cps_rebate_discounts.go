package entity

import (
	"fmt"
	"howell/howell_rpc/kitex_gen/models"
	"strings"
	"time"
)

type CpsRebateDiscounts struct {
	ID        string    `json:"id"`
	AppID     string    `json:"app_id"`
	Name      string    `json:"name"`
	CpsType   int32     `json:"cps_type"`
	JumpLink  string    `json:"jump_link"`
	Status    int       `json:"status"`
	Extra     string    `json:"extra"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCpsRebateDiscounts() *CpsRebateDiscounts {
	return &CpsRebateDiscounts{}
}

func (m *CpsRebateDiscounts) FromRPC(r *models.CpsRebateDiscounts) *CpsRebateDiscounts {
	return &CpsRebateDiscounts{
		ID:       r.GetId(),
		AppID:    r.GetAppId(),
		Name:     r.GetName(),
		CpsType:  int32(r.GetCpsType()),
		JumpLink: r.GetJumpLink(),
		Status:   0,
		Extra:    r.GetExtra(),
	}
}

func (m *CpsRebateDiscounts) ToRPC() *models.CpsRebateDiscounts {
	status := int32(m.Status)
	return &models.CpsRebateDiscounts{
		Id:    &m.ID,
		AppId: &m.AppID,
		Name:  &m.Name,
		CpsType: models.CpsTypePtr(
			models.CpsType(m.CpsType)),
		JumpLink: &m.JumpLink,
		Status:   &status,
		Extra:    &m.Extra,
	}
}

func (m *CpsRebateDiscounts) Check() error {
	if m.AppID == "" {
		return fmt.Errorf("invalid app_id")
	}
	if m.Name == "" {
		return fmt.Errorf("invalid name")
	}
	if m.JumpLink == "" || len(strings.TrimSpace(m.JumpLink)) != len(m.JumpLink) {
		return fmt.Errorf("invalid jump_link")
	}
	return nil
}
