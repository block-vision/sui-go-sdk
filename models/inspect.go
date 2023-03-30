package models

import "github.com/shoshinsquare/sui-go-sdk/models/sui_json_rpc_types"

type DevInsepctTransactionRequest struct {
	TxBytes string
}

type DevInspectTransactionResponse struct {
	Effects sui_json_rpc_types.SuiTransactionEffects `json:"effects"`
	Results []interface{}                            `json:"results"`
}

type DevInspectMoveCallRequest struct {
	SenderAddress   string
	PackageObjectID string
	Module          string
	Function        string
	TypeArguments   []string
	Arguments       []interface{}
}

type DevInspectMoveCallResponse struct {
	Effects sui_json_rpc_types.SuiTransactionEffects `json:"effects"`
	Results []interface{}                            `json:"results"`
}
