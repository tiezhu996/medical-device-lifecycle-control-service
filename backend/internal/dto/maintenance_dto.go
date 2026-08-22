package dto

import (
	"fmt"
	"time"
)

// CreateMaintenanceReq 创建保养/维修工单请求。
type CreateMaintenanceReq struct {
	DeviceID         uint       `json:"device_id" binding:"required,min=1"`
	Type             string     `json:"type" binding:"required,oneof=daily weekly monthly yearly repair"`
	PlannedDate      *time.Time `json:"planned_date"`
	Content          string     `json:"content" binding:"omitempty,max=1024"`
	FaultDescription string     `json:"fault_description" binding:"omitempty,max=1024"`
	Engineer         string     `json:"engineer" binding:"omitempty,max=64"`
}

// StartMaintenanceReq 开始执行请求。
type StartMaintenanceReq struct {
	Engineer string `json:"engineer" binding:"required,max=64"`
}

// CompleteMaintenanceReq 完成工单请求。
type CompleteMaintenanceReq struct {
	Content       string  `json:"content" binding:"required,max=1024"`
	ReplacedParts string  `json:"replaced_parts" binding:"omitempty,max=1024"`
	WorkHours     float64 `json:"work_hours" binding:"omitempty,min=0"`
	Cost          float64 `json:"cost" binding:"omitempty,min=0"`
	RepairResult  string  `json:"repair_result" binding:"omitempty,max=1024"`
}

// CancelMaintenanceReq 取消工单请求。
type CancelMaintenanceReq struct {
	Reason string `json:"reason" binding:"omitempty,max=512"`
}

func DefaultMaintenancePlanTypes() ([]string, error) {
	raw := []string{"daily", "weekly", "monthly", "yearly"}
	seen := make(map[string]struct{}, len(raw))
	out := make([]string, 0, len(raw))
	for _, value := range raw {
		switch value {
		case "daily", "weekly", "monthly", "yearly":
		default:
			return nil, fmt.Errorf("unsupported maintenance type %q", value)
		}
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		out = append(out, value)
	}
	return out, nil
}
