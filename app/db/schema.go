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

// Tasks - Submitted job tracker table schema
type Tasks struct {
	ID         string    `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Client     string    `gorm:"column:client;type:char(66);not null;index"`
	StartBlock uint64    `gorm:"column:startBlock;type:bigint;not null;index"`
	Contract   string    `gorm:"column:contract;type:char(66);index"`
	Topic0     string    `gorm:"column:topic0;type:char(66);index"`
	Topic1     string    `gorm:"column:topic1;type:char(66);index"`
	Topic2     string    `gorm:"column:topic2;type:char(66);index"`
	Topic3     string    `gorm:"column:topic3;type:char(66);index"`
	TimeStamp  time.Time `gorm:"column:ts;type:timestamp;not null"`
	Users      Users     `gorm:"references:apiKey"`
}

// TableName - Overriding default table name
func (Tasks) TableName() string {
	return "tasks"
}
