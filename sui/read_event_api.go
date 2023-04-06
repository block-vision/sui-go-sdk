package sui

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/common/rpc_client"
	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

type IReadEventFromSuiAPI interface {
	GetEvents(ctx context.Context, req models.GetEventsRequest, opts ...interface{}) (interface{}, error)
	QueryEvents(ctx context.Context, req models.QueryEventsRequest, opts ...interface{}) (models.QueryEventsResponse, error)
}

type suiReadEventFromSuiImpl struct {
	cli *rpc_client.RPCClient
}

// GetEventsBy implements method `sui_getEvents`.
// Returns an array of EventEnvelops according to your filter request condition.
func (s *suiReadEventFromSuiImpl) GetEvents(ctx context.Context, req models.GetEventsRequest, opts ...interface{}) (interface{}, error) {

	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getEvents",
		Params: []interface{}{
			req.EventQuery,
			req.Cursor,
			req.Limit,
			req.DescOrder,
		},
	})
	fmt.Println(respBytes, err)
	if err != nil {
		return models.GetEventsResponse{}, err
	}

	var rsp interface{}
	err = json.Unmarshal(respBytes, &rsp)
	if err != nil {
		return models.GetEventsResponse{}, err
	}
	return rsp, nil
}

// GetEventsBy implements method `sui_getEvents`.
// Returns an array of EventEnvelops according to your filter request condition.
func (s *suiReadEventFromSuiImpl) QueryEvents(ctx context.Context, req models.QueryEventsRequest, opts ...interface{}) (models.QueryEventsResponse, error) {

	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "suix_queryEvents",
		Params: []interface{}{
			req.EventQuery,
			req.Cursor,
			req.Limit,
			req.DescOrder,
		},
	})
	if err != nil {
		return models.QueryEventsResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.QueryEventsResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}

	var rsp models.QueryEventsResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.QueryEventsResponse{}, err
	}
	return rsp, nil
}
