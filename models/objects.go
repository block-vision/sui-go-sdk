package models

import "github.com/block-vision/sui-go-sdk/models/sui_types"

type SuiObjectInfo struct {
	sui_types.SuiObjectRef
	Type string `json:"type_"`
	sui_types.Owner
	PreviousTransaction string `json:"previousTransaction"`
}

type SuiMoveObject struct {
	Type              string
	Fields            map[string]interface{}
	HasPublicTransfer bool
}

type SuiMovePackage struct {
	Disassembled string `json:"disassembled"`
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

type SuiRawMovePackage struct {
	Id        string `json:"id,omitempty"`
	ModuleMap string `json:"moduleMap,omitempty"`
}

type SuiRawData struct {
	DataType string `json:"dataType"`
	SuiRawMoveObject
	SuiRawMovePackage
}
