package model

import "time"

// CalibrationRecord 计量器具台账与计量结果记录。
type CalibrationRecord struct {
	ID                     uint       `gorm:"primaryKey" json:"id"`
	InstrumentNo           string     `gorm:"size:64;uniqueIndex;not null" json:"instrument_no"`
	DeviceID               uint       `gorm:"not null;index" json:"device_id"`
	DeviceName             string     `gorm:"size:128;not null" json:"device_name"`
	CalibrationCycleMonths int        `json:"calibration_cycle_months"`
	LastCalibrationDate    *time.Time `json:"last_calibration_date"`
	NextCalibrationDate    *time.Time `gorm:"index" json:"next_calibration_date"`
	Status                 string     `gorm:"size:32;not null;default:normal;index" json:"status"`
	Result                 string     `gorm:"size:16" json:"result"`
	CertificateNo          string     `gorm:"size:128" json:"certificate_no"`
	CalibrationOrg         string     `gorm:"size:128" json:"calibration_org"`
	Remark                 string     `gorm:"size:512" json:"remark"`
	CreatedBy              string     `gorm:"size:64" json:"created_by"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}

// CanRecordResult 判断当前状态是否允许登记新的计量结果。
// scrapped 与 unqualified 为终态：前者设备已报废，后者设备已据此禁用，
// 再次登记会让状态机倒退并掩盖既有不合格结论。
func (c *CalibrationRecord) CanRecordResult() bool {
	return c.Status != "scrapped" && c.Status != "unqualified"
}
