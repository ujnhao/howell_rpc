package models

import (
	"fmt"
	"howell/howell_rpc/kitex_gen/coder/hao/howell_rpc"
	"strings"
	"time"
)

type CpsRebateDiscounts struct {
	ID        string    `json:"id"`
	AppID     int64     `json:"app_id"`
	Name      string    `json:"name"`
	CPSType   int32     `json:"cps_type"`
	JumpLink  string    `json:"jump_link"`
	Status    int       `json:"status"`
	Extra     string    `json:"extra"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCpsRebateDiscounts() *CpsRebateDiscounts {
	return &CpsRebateDiscounts{}
}

func (m *CpsRebateDiscounts) FromRPC(r *howell_rpc.CpsRebateDiscounts) *CpsRebateDiscounts {
	return &CpsRebateDiscounts{
		ID:       r.GetId(),
		AppID:    r.GetAppId(),
		Name:     r.GetName(),
		CPSType:  int32(r.GetCpsType()),
		JumpLink: r.GetJumpLink(),
		Status:   0,
		Extra:    r.GetExtra(),
	}
}

func (m *CpsRebateDiscounts) Check() error {
	if m.AppID <= 0 {
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
