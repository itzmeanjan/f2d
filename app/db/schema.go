package db

import (
	"time"

	"github.com/lib/pq"
)

// Users - Users table definition
type Users struct {
	APIKey    string    `gorm:"column:apikey;type:char(66);primaryKey"`
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
	StartBlock uint64    `gorm:"column:startblock;type:bigint;not null;index"`
	Contract   string    `gorm:"column:contract;type:char(66);index"`
	Topic0     string    `gorm:"column:topic0;type:char(66);index"`
	Topic1     string    `gorm:"column:topic1;type:char(66);index"`
	Topic2     string    `gorm:"column:topic2;type:char(66);index"`
	Topic3     string    `gorm:"column:topic3;type:char(66);index"`
	TimeStamp  time.Time `gorm:"column:ts;type:timestamp;not null"`
	Enabled    bool      `gorm:"column:enabled;type:boolean;default:true;index"`
	Users      Users     `gorm:"foreignKey:client;references:apikey"`
}

// TableName - Overriding default table name
func (Tasks) TableName() string {
	return "tasks"
}

// EventLogs - Received/ fetched event logs, emitted by smart contract interaction(s)
type EventLogs struct {
	BlockHash       string         `gorm:"column:blockhash;type:char(66);not null;primaryKey"`
	Index           uint           `gorm:"column:index;type:integer;not null;primaryKey"`
	Origin          string         `gorm:"column:origin;type:char(42);not null;index"`
	Topics          pq.StringArray `gorm:"column:topics;type:text[];not null;index:,type:gin"`
	Data            []byte         `gorm:"column:data;type:bytea"`
	TransactionHash string         `gorm:"column:txhash;type:char(66);not null;index"`
	BlockNumber     uint64         `gorm:"column:blocknumber;type:bigint;not null;index"`
}

// TableName - Overriding default table name
func (EventLogs) TableName() string {
	return "event_logs"
}

// TaskResults - Stored job results
type TaskResults struct {
	BlockHash string    `gorm:"column:blockhash;type:char(66);not null;primaryKey"`
	Index     uint      `gorm:"column:index;type:integer;not null;primaryKey"`
	ID        string    `gorm:"column:id;type:uuid;not null;primaryKey"`
	EventLogs EventLogs `gorm:"foreignKey:blockhash,index;references:blockhash,index"`
	Tasks     Tasks     `gorm:"foreignKey:id;references:id"`
}

// TableName - Overriding default table name
func (TaskResults) TableName() string {
	return "task_results"
}
