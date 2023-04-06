package models

type Operation struct {
	JsonRPC string        `json:"jsonrpc"`
	ID      uint64        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}
