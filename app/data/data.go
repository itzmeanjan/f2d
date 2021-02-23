package data

import "gorm.io/gorm"

// Resources - File, database, network resources which are to be accessed
// from several go routines, fulfilling different purposes, to be kept/ passed
// along using this struct
type Resources struct {
	DB *gorm.DB
}

// EtteSubscriptionRequest - Subscribe to event(s) of interest, emitted by `ette`
type EtteSubscriptionRequest struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	APIKey string `json:"apiKey"`
}
