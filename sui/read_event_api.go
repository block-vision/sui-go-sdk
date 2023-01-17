package sui

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/block-vision/sui-go-sdk/common/rpc_client"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

type IReadEventFromSuiAPI interface {
	GetEvents(ctx context.Context, req models.GetEventsRequest, opts ...interface{}) (models.GetEventsResponse, error)
}

type suiReadEventFromSuiImpl struct {
	cli *rpc_client.RPCClient
}

// GetEventsBy implements method `sui_getEvents`.
// Returns an array of EventEnvelops according to your filter request condition.
func (s *suiReadEventFromSuiImpl) GetEvents(ctx context.Context, req models.GetEventsRequest, opts ...interface{}) (models.GetEventsResponse, error) {
	var rsp models.GetEventsResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getEvents",
		Params: []interface{}{
			req.EventQuery,
			req.Cursor,
			req.Limit,
			req.DescOrder,
		},
	})
	if err != nil {
		return models.GetEventsResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetEventsResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result.data").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsResponse{}, err
	}
	return rsp, nil
}
