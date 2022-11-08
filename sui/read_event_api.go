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
	GetEventsByEventType(ctx context.Context, req models.GetEventsByEventTypeRequest, opts ...interface{}) (models.GetEventsByEventTypeResponse, error)
	GetEventsByModule(ctx context.Context, req models.GetEventsByModuleRequest, opts ...interface{}) (models.GetEventsByModuleResponse, error)
	GetEventsByObject(ctx context.Context, req models.GetEventsByObjectRequest, opts ...interface{}) (models.GetEventsByObjectResponse, error)
	GetEventsByOwner(ctx context.Context, req models.GetEventsByOwnerRequest, opts ...interface{}) (models.GetEventsByOwnerResponse, error)
	GetEventsBySender(ctx context.Context, req models.GetEventsBySenderRequest, opts ...interface{}) (models.GetEventsBySenderResponse, error)
	GetEventsByTransaction(ctx context.Context, req models.GetEventsByTransactionRequest, opts ...interface{}) (models.GetEventsByTransactionResponse, error)
}

type suiReadEventFromSuiImpl struct {
	cli *rpc_client.RPCClient
}

// GetEventsByEventType implements method `sui_getEventsByEventType`.
// Returns an array of EventEnvelops according to your filter request condition.
func (s *suiReadEventFromSuiImpl) GetEventsByEventType(ctx context.Context, req models.GetEventsByEventTypeRequest, opts ...interface{}) (models.GetEventsByEventTypeResponse, error) {
	var rsp models.GetEventsByEventTypeResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getEventsByEventType",
		Params: []interface{}{
			req.EventType,
			req.Count,
			req.StartTime,
			req.EndTime,
		},
	})
	if err != nil {
		return models.GetEventsByEventTypeResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetEventsByEventTypeResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsByEventTypeResponse{}, err
	}
	return rsp, nil
}

// GetEventsByModule implements method `sui_getEventsByModule`.
// Returns an array of EventEnvelops according to your filter request condition
func (s *suiReadEventFromSuiImpl) GetEventsByModule(ctx context.Context, req models.GetEventsByModuleRequest, opts ...interface{}) (models.GetEventsByModuleResponse, error) {
	var rsp models.GetEventsByModuleResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getEventsByModule",
		Params: []interface{}{
			req.Package,
			req.Module,
			req.Count,
			req.StartTime,
			req.EndTime,
		},
	})
	if err != nil {
		return models.GetEventsByModuleResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetEventsByModuleResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsByModuleResponse{}, err
	}
	return rsp, nil
}

// GetEventsByObject implements method `sui_getEventsByObject`.
// Returns an array of EventEnvelops according to your filter request condition
func (s *suiReadEventFromSuiImpl) GetEventsByObject(ctx context.Context, req models.GetEventsByObjectRequest, opts ...interface{}) (models.GetEventsByObjectResponse, error) {
	var rsp models.GetEventsByObjectResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getEventsByObject",
		Params: []interface{}{
			req.Object,
			req.Count,
			req.StartTime,
			req.EndTime,
		},
	})
	if err != nil {
		return models.GetEventsByObjectResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetEventsByObjectResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsByObjectResponse{}, err
	}
	return rsp, nil
}

// GetEventsByOwner implements method `sui_getEventsByOwner`.
// Returns an array of EventEnvelops according to your filter request condition
func (s *suiReadEventFromSuiImpl) GetEventsByOwner(ctx context.Context, req models.GetEventsByOwnerRequest, opts ...interface{}) (models.GetEventsByOwnerResponse, error) {
	var rsp models.GetEventsByOwnerResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getEventsByOwner",
		Params: []interface{}{
			req.Owner,
			req.Count,
			req.StartTime,
			req.EndTime,
		},
	})
	if err != nil {
		return models.GetEventsByOwnerResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetEventsByOwnerResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsByOwnerResponse{}, err
	}
	return rsp, nil
}

// GetEventsBySender implements method `sui_getEventsBySender`.
// Returns an array of EventEnvelops according to your filter request condition
func (s *suiReadEventFromSuiImpl) GetEventsBySender(ctx context.Context, req models.GetEventsBySenderRequest, opts ...interface{}) (models.GetEventsBySenderResponse, error) {
	var rsp models.GetEventsBySenderResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getEventsBySender",
		Params: []interface{}{
			req.Sender,
			req.Count,
			req.StartTime,
			req.EndTime,
		},
	})
	if err != nil {
		return models.GetEventsBySenderResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetEventsBySenderResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsBySenderResponse{}, err
	}
	return rsp, nil
}

// GetEventsByTransaction implements method `sui_getEventsByTransaction`.
// Returns an array of EventEnvelops according to your filter request condition
func (s *suiReadEventFromSuiImpl) GetEventsByTransaction(ctx context.Context, req models.GetEventsByTransactionRequest, opts ...interface{}) (models.GetEventsByTransactionResponse, error) {
	var rsp models.GetEventsByTransactionResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getEventsByTransaction",
		Params: []interface{}{
			req.Digest,
		},
	})
	if err != nil {
		return models.GetEventsByTransactionResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetEventsByTransactionResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetEventsByTransactionResponse{}, err
	}
	return rsp, nil
}
