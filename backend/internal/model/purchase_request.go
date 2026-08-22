package model

import "time"

// PurchaseRequest 设备采购申请，状态机：科室申请→设备科审核→院长审批→到货→验收入台账。
type PurchaseRequest struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`
	RequestNo          string     `gorm:"size:64;uniqueIndex;not null" json:"request_no"`
	Department         string     `gorm:"size:128;not null" json:"department"`
	ApplicantName      string     `gorm:"size:64;not null" json:"applicant_name"`
	DeviceName         string     `gorm:"size:128;not null" json:"device_name"`
	Model              string     `gorm:"size:128" json:"model"`
	Manufacturer       string     `gorm:"size:128" json:"manufacturer"`
	Quantity           int        `gorm:"not null;default:1" json:"quantity"`
	BudgetAmount       float64    `gorm:"type:decimal(14,2)" json:"budget_amount"`
	Reason             string     `gorm:"size:512" json:"reason"`
	Status             string     `gorm:"size:32;not null;default:pending_device_admin;index" json:"status"`
	SubmittedBy        string     `gorm:"size:64" json:"submitted_by"`
	SubmittedAt        *time.Time `json:"submitted_at"`
	DeviceAdmin        string     `gorm:"size:64" json:"device_admin"`
	DeviceAdminComment string     `gorm:"size:512" json:"device_admin_comment"`
	DeviceAdminAt      *time.Time `json:"device_admin_at"`
	Dean               string     `gorm:"size:64" json:"dean"`
	DeanComment        string     `gorm:"size:512" json:"dean_comment"`
	DeanAt             *time.Time `json:"dean_at"`
	DeliveredAt        *time.Time `json:"delivered_at"`
	AcceptancePerson   string     `gorm:"size:64" json:"acceptance_person"`
	AcceptanceDate     *time.Time `json:"acceptance_date"`
	PartsList          string     `gorm:"size:1024" json:"parts_list"`
	CertificateNo      string     `gorm:"size:128" json:"certificate_no"`
	RegistrationNo     string     `gorm:"size:128" json:"registration_no"`
	DeviceID           uint       `gorm:"index" json:"device_id"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

// CanAccept 验收前置校验：仅「已到货待验收」且尚未生成设备的申请可验收。
// 必须先到货登记（delivered），且禁止重复验收（DeviceID 已绑定即代表已验收过）。
func (p *PurchaseRequest) CanAccept() bool {
	if p.DeviceID != 0 {
		return false
	}
	return p.Status == "delivered" && p.DeliveredAt != nil
}

func (p *PurchaseRequest) ApplyAcceptance(person string, date *time.Time, parts, certificate, registration string, deviceID uint) {
	p.Status = "accepted"
	p.AcceptancePerson = person
	p.AcceptanceDate = date
	p.PartsList = parts
	p.CertificateNo = certificate
	p.RegistrationNo = registration
	p.DeviceID = deviceID
}
