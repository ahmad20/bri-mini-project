package entities

type Customer struct {
	ID           int `gorm:"primaryKey"`
	Email        string
	First_Name   string
	Last_Name    string
	Avatar       string
	RegisteredBy int // admin or superadmin ID
}
