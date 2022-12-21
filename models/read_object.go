package models

import (
	"github.com/block-vision/sui-go-sdk/models/sui_json_rpc_types"
	"github.com/block-vision/sui-go-sdk/models/sui_types"
)

type GetObjectRequest struct {
	ObjectID string `json:"objectID"`
}

type GetObjectResponse struct {
	Details struct {
		Data sui_json_rpc_types.SuiParsedMoveObject `json:"data"`
		sui_json_rpc_types.OwnedObjectRef
		PreviousTransaction string                 `json:"previousTransaction"`
		StorageRebate       uint64                 `json:"storageRebate"`
		Reference           sui_types.SuiObjectRef `json:"reference"`
	} `json:"details"`
	Status string `json:"status"`
}

type GetObjectsOwnedByAddressRequest struct {
	Address string `json:"address"`
}

type GetObjectsOwnedByAddressResponse struct {
	Result []SuiObjectInfo `json:"result"`
}

type GetObjectsOwnedByObjectRequest struct {
	ObjectID string `json:"objectID"`
}
type GetObjectsOwnedByObjectResponse struct {
	Result []SuiObjectInfo `json:"result"`
}

type GetRawObjectRequest struct {
	ObjectID string `json:"objectID"`
}
type GetRawObjectResponse struct {
	Details struct {
		Data sui_json_rpc_types.SuiParsedMoveObject `json:"data"`
		sui_json_rpc_types.OwnedObjectRef
		PreviousTransaction string                 `json:"previousTransaction"`
		StorageRebate       uint64                 `json:"storageRebate"`
		Reference           sui_types.SuiObjectRef `json:"reference"`
	} `json:"details"`
	Status string `json:"status"`
}

type TryGetPastObjectRequest struct {
	ObjectID string `json:"objectID"`
	Version  uint64 `json:"version"`
}

type TryGetPastObjectResponse struct {
	Status  string `json:"status"`
	Details struct {
		Data sui_json_rpc_types.SuiParsedMoveObject `json:"data"`
		sui_json_rpc_types.OwnedObjectRef
		PreviousTransaction string                 `json:"previousTransaction"`
		StorageRebate       uint64                 `json:"storageRebate"`
		Reference           sui_types.SuiObjectRef `json:"reference"`
	} `json:"details"`
}

type GetCoinMetadataRequest struct {
	CoinType string
}

type GetCoinMetadataResponse struct {
	Decimals    uint8  `json:"decimals"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl,omitempty"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
}
