package models

import "github.com/shoshinsquare/sui-go-sdk/models/sui_types"

type SuiObjectInfo struct {
	sui_types.SuiObjectRef
	Type string `json:"type"`
	sui_types.Owner
	PreviousTransaction string `json:"previousTransaction"`
}

type SuiMoveObject struct {
	Type              string
	Fields            map[string]interface{}
	HasPublicTransfer bool
}

type SuiMoveModuleId struct {
	Address string
	Name    string
}

type SuiMoveNormalizedModule struct {
	FileFormatVersion uint64
	Address           string
	Name              string
	Friends           []SuiMoveModuleId
}
