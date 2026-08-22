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

func NextMaintenanceDate(now time.Time, maintenanceType string) (*time.Time, error) {
	offsets := map[string][3]int{
		"daily": {0, 0, 1}, "weekly": {0, 0, 7},
		"monthly": {0, 0, 30}, "yearly": {0, 0, 365}, "repair": {0, 0, 30},
	}
	offset, ok := offsets[maintenanceType]
	if !ok {
		return nil, fmt.Errorf("unsupported maintenance type %q", maintenanceType)
	}
	next := now.AddDate(offset[0], offset[1], offset[2])
	return &next, nil
}
