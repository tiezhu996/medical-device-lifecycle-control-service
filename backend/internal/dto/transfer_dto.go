package dto

import "strings"

// CreateTransferReq 创建调拨申请请求。
type CreateTransferReq struct {
	DeviceID     uint   `json:"device_id" binding:"required,min=1"`
	ToDepartment string `json:"to_department" binding:"required,max=128"`
	ToPerson     string `json:"to_person" binding:"omitempty,max=64"`
	Reason       string `json:"reason" binding:"omitempty,max=512"`
}

// TransferApproveReq 调拨审批请求。
type TransferApproveReq struct {
	Comment  string   `json:"comment" binding:"omitempty,max=512"`
	Evidence []string `json:"evidence" binding:"omitempty,max=20,dive,max=128"`
}

func NormalizeTransferEvidence(values []string) []string {
	out := values[:0]
	for _, value := range values {
		if value = strings.TrimSpace(value); value != "" {
			out = append(out, value)
		}
	}
	return out
}
