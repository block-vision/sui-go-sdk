package models

import "github.com/block-vision/sui-go-sdk/models/sui_json_rpc_types"

type GetEventsByEventTypeRequest struct {
	EventType string `json:"eventType"`
	Count     uint64 `json:"count"`
	StartTime uint64 `json:"startTime"`
	EndTime   uint64 `json:"endTime"`
}

type GetEventsByEventTypeResponse struct {
	Result []sui_json_rpc_types.SuiEventEnvelop
}

type GetEventsByModuleRequest struct {
	Package   string `json:"package"`
	Module    string `json:"module"`
	Count     uint64 `json:"count"`
	StartTime uint64 `json:"startTime"`
	EndTime   uint64 `json:"endTime"`
}

type GetEventsByModuleResponse struct {
	Result []sui_json_rpc_types.SuiEventEnvelop
}

type GetEventsByObjectRequest struct {
	Object    string `json:"object"`
	Count     uint64 `json:"count"`
	StartTime uint64 `json:"startTime"`
	EndTime   uint64 `json:"endTime"`
}

type GetEventsByObjectResponse struct {
	Result []sui_json_rpc_types.SuiEventEnvelop
}

type GetEventsByOwnerRequest struct {
	Owner     string `json:"owner"`
	Count     uint64 `json:"count"`
	StartTime uint64 `json:"startTime"`
	EndTime   uint64 `json:"endTime"`
}

type GetEventsByOwnerResponse struct {
	Result []sui_json_rpc_types.SuiEventEnvelop
}

type GetEventsBySenderRequest struct {
	Sender    string `json:"sender"`
	Count     uint64 `json:"count"`
	StartTime uint64 `json:"startTime"`
	EndTime   uint64 `json:"endTime"`
}

type GetEventsBySenderResponse struct {
	Result []sui_json_rpc_types.SuiEventEnvelop
}

type GetEventsByTransactionRequest struct {
	Digest string `json:"digest"`
}

type GetEventsByTransactionResponse struct {
	Result []sui_json_rpc_types.SuiEventEnvelop
}
