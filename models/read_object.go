package models

import (
	"github.com/shoshinsquare/sui-go-sdk/models/sui_json_rpc_types"
	"github.com/shoshinsquare/sui-go-sdk/models/sui_types"
)

type GetObjectRequest struct {
	ObjectID string `json:"objectID"`
}

type GetObjectResponse struct {
	Data struct {
		Digest   string `json:"digest"`
		ObjectID string `json:"objectId"`
		Version  uint64 `json:"version"`
		Type     string `json:"type"`
		Owner    struct {
			AddressOwner string `json:"AddressOwner"`
		} `json:"owner"`
		PreviousTransaction string `json:"previousTransaction"`
		StorageRebate       uint64 `json:"storageRebate"`
		Content             struct {
			DataType          string `json:"dataType"`
			Type              string `json:"type"`
			HasPublicTransfer bool   `json:"hasPublicTransfer"`
		} `json:"content"`
	} `json:"data"`
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
	CoinType string `json:"coin_type"`
}

type GetCoinMetadataResponse struct {
	Decimals    uint8  `json:"decimals"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl,omitempty"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
}

type GetDynamicFieldObejctRequest struct {
	ParentObjectID string
	Name           string
}

type GetDynamicFieldObjectResponse struct {
	Details struct {
		Data sui_json_rpc_types.SuiParsedMoveObject `json:"data"`
		sui_json_rpc_types.OwnedObjectRef
		PreviousTransaction string                 `json:"previousTransaction"`
		StorageRebate       uint64                 `json:"storageRebate"`
		Reference           sui_types.SuiObjectRef `json:"reference"`
	} `json:"details"`
	Status string `json:"status"`
}

type GetOwnedObjectsRequest struct {
	Address string `json:"address"`
}

type GetOwnedObjectsResponse struct {
	Data        []sui_json_rpc_types.SuiParsedMoveObject `json:"data"`
	HasNextPage bool                                     `json:"hasNextPage"`
}

type GetDynamicFieldRequest struct {
	ParentObjectID string `json:"parent_object_id"`
}

type GetDynamicFieldResponse struct {
}
