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

// EtteSubscriptionResponse - Event subscription/ unsubscription response from `ette`
type EtteSubscriptionResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"msg"`
}

// EtteBlock - Newly mined block to be received in this form, from `ette`
type EtteBlock struct {
	Hash                string  `json:"hash"`
	Number              uint64  `json:"number"`
	Time                uint64  `json:"time"`
	ParentHash          string  `json:"parentHash"`
	Difficulty          string  `json:"difficulty"`
	GasUsed             uint64  `json:"gasUsed"`
	GasLimit            uint64  `json:"gasLimit"`
	Nonce               string  `json:"nonce"`
	Miner               string  `json:"miner"`
	Size                float64 `json:"size"`
	StateRootHash       string  `json:"stateRootHash"`
	UncleHash           string  `json:"uncleHash"`
	TransactionRootHash string  `json:"txRootHash"`
	ReceiptRootHash     string  `json:"receiptRootHash"`
	ExtraData           string  `json:"extraData"`
}
