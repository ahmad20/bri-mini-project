package entities

type approvalStatus string

const (
	ApprovalStatusPending  approvalStatus = "waiting"
	ApprovalStatusApproved approvalStatus = "approved"
	ApprovalStatusRejected approvalStatus = "rejected"
)

type Account struct {
	ID             string         `gorm:"primaryKey"`
	Username       string         `gorm:"not null; unique"`
	Password       string         `json:"-"`
	Role           string         `gorm:"not null; default:'admin'"`
	ApprovalStatus approvalStatus `gorm:"not null; default:'waiting'"` //approved or rejected by super admin
	ApprovedBy     int            `gorm:"not null; default:0"`
	Status         string         `gorm:"not null; default:'inactive'"` //active or deactive by super admin
}
