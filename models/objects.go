package models

import (
	"encoding/json"

	"github.com/block-vision/sui-go-sdk/models/sui_types"
	"github.com/tidwall/gjson"
)

type SuiObjectInfo struct {
	sui_types.SuiObjectRef
	Type string `json:"type_"`
	sui_types.Owner
	PreviousTransaction string `json:"previousTransaction"`
}

type SuiMoveObject struct {
	Type              string                 `json:"type"`
	Fields            map[string]interface{} `json:"fields"`
	HasPublicTransfer bool                   `json:"hasPublicTransfer"`
}

type SuiMovePackage struct {
	Disassembled string `json:"disassembled"`
}

type SuiMoveModuleId struct {
	Address string `json:"address"`
	Name    string `json:"name"`
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

type DynamicFieldName struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

func (v DynamicFieldName) Field(field string) gjson.Result {
	return gjson.GetBytes(v.Value, field)
}
