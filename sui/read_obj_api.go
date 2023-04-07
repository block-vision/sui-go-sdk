package sui

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/shoshinsquare/sui-go-sdk/common/rpc_client"
	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/models/sui_json_rpc_types"
	"github.com/tidwall/gjson"
)

type IReadObjectFromSuiAPI interface {
	GetObject(ctx context.Context, req models.GetObjectRequest, opts ...interface{}) (models.GetObjectResponse, error)
	GetRawObject(ctx context.Context, req models.GetRawObjectRequest, opts ...interface{}) (models.GetRawObjectResponse, error)
	TryGetPastObject(ctx context.Context, req models.TryGetPastObjectRequest, opt ...interface{}) (models.TryGetPastObjectResponse, error)
	GetCoinMetadata(ctx context.Context, req models.GetCoinMetadataRequest, opt ...interface{}) (models.GetCoinMetadataResponse, error)
	GetOwnedObjects(ctx context.Context, req models.GetOwnedObjectsRequest, opt ...interface{}) (models.GetOwnedObjectsResponse, error)
	GetDynamicField(ctx context.Context, req models.GetDynamicFieldRequest, opt ...interface{}) (models.GetDynamicFieldResponse, error)
	GetAllNFT(ctx context.Context, address string) ([]models.GetObjectResponse, error)
	GetTransactionBlock(ctx context.Context, req models.GetTransactionBlockRequest, opt ...interface{}) (models.GetTransactionBlockResponse, error)
}

type suiReadObjectFromSuiImpl struct {
	cli *rpc_client.RPCClient
}

func (s *suiReadObjectFromSuiImpl) GetTransactionBlock(ctx context.Context, req models.GetTransactionBlockRequest, opt ...interface{}) (models.GetTransactionBlockResponse, error) {
	var rsp models.GetTransactionBlockResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getTransactionBlock",
		Params: []interface{}{
			req.Digest,
			map[string]bool{
				"showInput":          false,
				"showRawInput":       false,
				"showEffects":        false,
				"showEvents":         true,
				"showObjectChanges":  false,
				"showBalanceChanges": false,
			},
		},
	})
	if err != nil {
		return models.GetTransactionBlockResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetTransactionBlockResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetTransactionBlockResponse{}, err
	}
	return rsp, nil
}

// GetObject implements method `sui_getObject`.
// Returns object details
func (s *suiReadObjectFromSuiImpl) GetObject(ctx context.Context, req models.GetObjectRequest, opts ...interface{}) (models.GetObjectResponse, error) {

	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getObject",
		Params: []interface{}{
			req.ObjectID,
			map[string]bool{
				"showType":                true,
				"showOwner":               true,
				"showPreviousTransaction": true,
				"showDisplay":             true,
				"showContent":             true,
				"showBcs":                 true,
				"showStorageRebate":       true,
			},
		},
	})
	if err != nil {
		return models.GetObjectResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetObjectResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetObjectResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetObjectResponse{}, err
	}
	return rsp, nil
}

// GetObjectsOwnedByAddress implements method `sui_getObjectsOwnedByAddress`.
// Returns an array of object information
func (s *suiReadObjectFromSuiImpl) GetObjectsOwnedByAddress(ctx context.Context, req models.GetObjectsOwnedByAddressRequest, opts ...interface{}) (models.GetObjectsOwnedByAddressResponse, error) {
	var rsp models.GetObjectsOwnedByAddressResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getObjectsOwnedByAddress",
		Params: []interface{}{
			req.Address,
		},
	})
	if err != nil {
		return models.GetObjectsOwnedByAddressResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetObjectsOwnedByAddressResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetObjectsOwnedByAddressResponse{}, err
	}
	return rsp, nil
}

// GetRawObject implements method `sui_getRawObject`.
// Returns object details
func (s *suiReadObjectFromSuiImpl) GetRawObject(ctx context.Context, req models.GetRawObjectRequest, opts ...interface{}) (models.GetRawObjectResponse, error) {
	var rsp models.GetRawObjectResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getRawObject",
		Params: []interface{}{
			req.ObjectID,
		},
	})
	if err != nil {
		return models.GetRawObjectResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetRawObjectResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetRawObjectResponse{}, err
	}
	return rsp, nil
}

// TryGetPastObject implements method `sui_tryGetPastObject`
// Note there is no software-level guarantee/SLA that objects with past versions can be retrieved by this API,
// even if the object and version exists/existed.
// The result may vary across nodes depending on their pruning policies.
// Return the object information for a specified version
func (s *suiReadObjectFromSuiImpl) TryGetPastObject(ctx context.Context, req models.TryGetPastObjectRequest, opts ...interface{}) (models.TryGetPastObjectResponse, error) {
	var rsp models.TryGetPastObjectResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_tryGetPastObject",
		Params: []interface{}{
			req.ObjectID,
		},
	})
	if err != nil {
		return models.TryGetPastObjectResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.TryGetPastObjectResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.TryGetPastObjectResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadObjectFromSuiImpl) GetCoinMetadata(ctx context.Context, req models.GetCoinMetadataRequest, opt ...interface{}) (models.GetCoinMetadataResponse, error) {
	var rsp models.GetCoinMetadataResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "suix_getCoinMetadata",
		Params: []interface{}{
			req.CoinType,
		},
	})
	if err != nil {
		return models.GetCoinMetadataResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetCoinMetadataResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetCoinMetadataResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadObjectFromSuiImpl) GetOwnedObjects(ctx context.Context, req models.GetOwnedObjectsRequest, opt ...interface{}) (models.GetOwnedObjectsResponse, error) {
	var rsp models.GetOwnedObjectsResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "suix_getOwnedObjects",
		Params: []interface{}{
			req.Address,
		},
	})
	if err != nil {
		return models.GetOwnedObjectsResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetOwnedObjectsResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetOwnedObjectsResponse{}, err
	}

	return rsp, nil
}

func (s *suiReadObjectFromSuiImpl) GetDynamicField(ctx context.Context, req models.GetDynamicFieldRequest, opt ...interface{}) (models.GetDynamicFieldResponse, error) {
	var rsp models.GetDynamicFieldResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "suix_getDynamicFields",
		Params: []interface{}{
			req.ParentObjectID,
		},
	})
	if err != nil {
		return models.GetDynamicFieldResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetDynamicFieldResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetDynamicFieldResponse{}, err
	}

	return rsp, nil
}

func (s *suiReadObjectFromSuiImpl) GetAllNFT(ctx context.Context, address string) ([]models.GetObjectResponse, error) {
	res, err := s.GetOwnedObjects(context.Background(), models.GetOwnedObjectsRequest{
		Address: address,
	})
	if err != nil {
		return nil, err
	}

	metaData := []models.GetObjectResponse{}
	queue := make(chan models.GetObjectResponse)

	for _, r := range res.Data {
		go func(param sui_json_rpc_types.SuiParsedMoveObject) {
			metaData, err := s.GetObject(context.Background(), models.GetObjectRequest{
				ObjectID: param.Data.ObjectID,
			})
			if err != nil {
				return
			}
			queue <- metaData
		}(r)
	}

	for range res.Data {
		select {
		case data := <-queue:
			if !strings.Contains(data.Data.Type, "0x2::coin::Coin") {
				metaData = append(metaData, data)
			}

		}
	}

	return metaData, nil
}
