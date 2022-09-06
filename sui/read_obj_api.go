package sui

import (
	"context"
	"encoding/json"
	"github.com/block-vision/sui-go-sdk/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui_error"
	"github.com/tidwall/gjson"
)

type IReadObjectFromSuiAPI interface {
	GetObject(ctx context.Context, req models.GetObjectRequest, opts ...interface{}) (models.GetObjectResponse, error)
	GetObjectsOwnedByAddress(ctx context.Context, req models.GetObjectsOwnedByAddressRequest, opts ...interface{}) (models.GetObjectsOwnedByAddressResponse, error)
	GetObjectsOwnedByObject(ctx context.Context, req models.GetObjectsOwnedByObjectRequest, opts ...interface{}) (models.GetObjectsOwnedByObjectResponse, error)
	GetRawObject(ctx context.Context, req models.GetRawObjectRequest, opts ...interface{}) (models.GetRawObjectResponse, error)
}

type suiReadObjectFromSuiImpl struct {
	conn *httpconn.HttpConn
}

// GetObject implements method `sui_getObject`.
// Returns object details
func (s *suiReadObjectFromSuiImpl) GetObject(ctx context.Context, req models.GetObjectRequest, opts ...interface{}) (models.GetObjectResponse, error) {
	var rsp models.GetObjectResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getObject",
		Params: []interface{}{
			req.ObjectID,
		},
	})
	if err != nil {
		return models.GetObjectResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetObjectResponse{}, sui_error.ErrInvalidJson
	}
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
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getObjectsOwnedByAddress",
		Params: []interface{}{
			req.Address,
		},
	})
	if err != nil {
		return models.GetObjectsOwnedByAddressResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetObjectsOwnedByAddressResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetObjectsOwnedByAddressResponse{}, err
	}
	return rsp, nil
}

// GetObjectsOwnedByObject implements method `sui_getObjectsOwnedByObject`.
// Returns an array of object information
func (s *suiReadObjectFromSuiImpl) GetObjectsOwnedByObject(ctx context.Context, req models.GetObjectsOwnedByObjectRequest, opts ...interface{}) (models.GetObjectsOwnedByObjectResponse, error) {
	var rsp models.GetObjectsOwnedByObjectResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getObjectsOwnedByObject",
		Params: []interface{}{
			req.ObjectID,
		},
	})
	if err != nil {
		return models.GetObjectsOwnedByObjectResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetObjectsOwnedByObjectResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetObjectsOwnedByObjectResponse{}, err
	}
	return rsp, nil
}

// GetRawObject implements method `sui_getRawObject`.
// Returns object details
func (s *suiReadObjectFromSuiImpl) GetRawObject(ctx context.Context, req models.GetRawObjectRequest, opts ...interface{}) (models.GetRawObjectResponse, error) {
	var rsp models.GetRawObjectResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getRawObject",
		Params: []interface{}{
			req.ObjectID,
		},
	})
	if err != nil {
		return models.GetRawObjectResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetRawObjectResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetRawObjectResponse{}, err
	}
	return rsp, nil
}
