package model

import (
	"fmt"
	"time"
)

// MaintenanceRecord 维护保养与故障维修记录。
type MaintenanceRecord struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	RecordNo         string     `gorm:"size:64;uniqueIndex;not null" json:"record_no"`
	DeviceID         uint       `gorm:"not null;index" json:"device_id"`
	DeviceName       string     `gorm:"size:128;not null" json:"device_name"`
	Type             string     `gorm:"size:32;not null;index" json:"type"`
	Status           string     `gorm:"size:32;not null;default:pending;index" json:"status"`
	PlannedDate      *time.Time `json:"planned_date"`
	ExecutedDate     *time.Time `json:"executed_date"`
	Engineer         string     `gorm:"size:64" json:"engineer"`
	Content          string     `gorm:"size:1024" json:"content"`
	ReplacedParts    string     `gorm:"size:1024" json:"replaced_parts"`
	WorkHours        float64    `gorm:"type:decimal(6,1)" json:"work_hours"`
	Cost             float64    `gorm:"type:decimal(12,2)" json:"cost"`
	FaultDescription string     `gorm:"size:1024" json:"fault_description"`
	RepairResult     string     `gorm:"size:1024" json:"repair_result"`
	CreatedBy        string     `gorm:"size:64" json:"created_by"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// NextMaintenanceDate 计算下次保养计划日期。
// 仅日检/周检/月检/年检可自动排期；故障维修（repair）由人工报修触发，不可自动排期。
// 月检/年检按自然月/年进位，落到日历边界（如 1 月 1 日 → 2 月 1 日）。
func NextMaintenanceDate(now time.Time, maintenanceType string) (*time.Time, error) {
	var next time.Time
	switch maintenanceType {
	case "daily":
		next = now.AddDate(0, 0, 1)
	case "weekly":
		next = now.AddDate(0, 0, 7)
	case "monthly":
		next = now.AddDate(0, 1, 0)
	case "yearly":
		next = now.AddDate(1, 0, 0)
	default:
		return nil, fmt.Errorf("unsupported maintenance type %q", maintenanceType)
	}
	return &next, nil
}
