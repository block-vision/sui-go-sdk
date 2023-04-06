package models

import "github.com/shoshinsquare/sui-go-sdk/models/sui_json_rpc_types"

type GetEventsRequest struct {
	EventQuery EventQuery
	Cursor     *string
	Limit      uint64
	DescOrder  bool
}

type GetEventsResponse struct {
	Result []sui_json_rpc_types.SuiEventEnvelop
}

type MoveModule struct {
	Package string `json:"package"`
	Module  string `json:"module"`
}
type EventQuery struct {
	MoveModule    MoveModule `json:"MoveModule,omitempty"`
	MoveEventType string     `json:"MoveEventType,omitempty"`
}

type CursorPage struct {
}
type QueryEventsRequest struct {
	EventQuery EventQuery
	Cursor     *CursorPage
	Limit      uint64
	DescOrder  bool
}

type QueryEventsResponse struct {
	Data        []sui_json_rpc_types.SuiEventEnvelop `json:"data"`
	HasNextPage bool                                 `json:"hasNextPage"`
	NextCursor  struct {
		TxDigest string `json:"txDigest"`
		EventSeq uint64 `json:"eventSeq"`
	} `json:"nextCursor"`
}
