package models

import "github.com/block-vision/sui-go-sdk/models/sui_json_rpc_types"

type GetEventsRequest struct {
	EventQuery interface{}
	Cursor     *string
	Limit      uint64
	DescOrder  bool
}

type GetEventsResponse struct {
	Result []sui_json_rpc_types.SuiEventEnvelop
}
