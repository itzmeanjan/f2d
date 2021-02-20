package db

import "time"

// Users - Users table definition
type Users struct {
	APIKey    string    `gorm:"column:apiKey;type:char(66);primaryKey"`
	Address   string    `gorm:"column:address;type:char(42);not null;index"`
	TimeStamp time.Time `gorm:"column:ts;type:timestamp;not null"`
	Enabled   bool      `gorm:"column:enabled;type:boolean;default:true"`
}

// TableName - Overriding default table name
func (Users) TableName() string {
	return "users"
}
