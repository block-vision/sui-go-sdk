// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/block-vision/sui-go-sdk/common/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

type IReadObjectFromSuiAPI interface {
	SuiXGetOwnedObjects(ctx context.Context, req models.SuiXGetOwnedObjectsRequest) (models.PaginatedObjectsResponse, error)
	SuiMultiGetObjects(ctx context.Context, req models.SuiMultiGetObjectsRequest) ([]*models.SuiObjectResponse, error)
	SuiXGetDynamicField(ctx context.Context, req models.SuiXGetDynamicFieldRequest) (models.PaginatedDynamicFieldInfoResponse, error)
	SuiXGetDynamicFieldObject(ctx context.Context, req models.SuiXGetDynamicFieldObjectRequest) (models.SuiObjectResponse, error)
}

type suiReadObjectFromSuiImpl struct {
	conn *httpconn.HttpConn
}

// SuiXGetOwnedObjects implements the method `suix_getOwnedObjects`, gets the list of objects owned by an address.
func (s *suiReadObjectFromSuiImpl) SuiXGetOwnedObjects(ctx context.Context, req models.SuiXGetOwnedObjectsRequest) (models.PaginatedObjectsResponse, error) {
	var rsp models.PaginatedObjectsResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_getOwnedObjects",
		Params: []interface{}{
			req.Address,
			req.Query,
			req.Cursor,
			req.Limit,
		},
	})
	if err != nil {
		return rsp, err
	}
	if !gjson.ValidBytes(respBytes) {
		return rsp, errors.New("not a valid json response")
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiMultiGetObjects implements the method `sui_multiGetObjects`, gets the object data for a list of objects.
func (s *suiReadObjectFromSuiImpl) SuiMultiGetObjects(ctx context.Context, req models.SuiMultiGetObjectsRequest) ([]*models.SuiObjectResponse, error) {
	var rsp []*models.SuiObjectResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_multiGetObjects",
		Params: []interface{}{
			req.ObjectIds,
			req.Options,
		},
	})
	if err != nil {
		return rsp, err
	}
	if !gjson.ValidBytes(respBytes) {
		return rsp, errors.New("not a valid json response")
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXGetDynamicField implements the method `suix_getDynamicFields`, gets the list of dynamic field objects owned by an object.
func (s *suiReadObjectFromSuiImpl) SuiXGetDynamicField(ctx context.Context, req models.SuiXGetDynamicFieldRequest) (models.PaginatedDynamicFieldInfoResponse, error) {
	var rsp models.PaginatedDynamicFieldInfoResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_getDynamicFields",
		Params: []interface{}{
			req.ObjectId,
			req.Cursor,
			req.Limit,
		},
	})
	if err != nil {
		return rsp, err
	}
	if !gjson.ValidBytes(respBytes) {
		return rsp, errors.New("not a valid json response")
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXGetDynamicFieldObject implements the method `suix_getDynamicFieldObject`, gets the dynamic field object information for a specified object.
func (s *suiReadObjectFromSuiImpl) SuiXGetDynamicFieldObject(ctx context.Context, req models.SuiXGetDynamicFieldObjectRequest) (models.SuiObjectResponse, error) {
	var rsp models.SuiObjectResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_getDynamicFieldObject",
		Params: []interface{}{
			req.ObjectId,
			req.DynamicFieldName,
		},
	})
	if err != nil {
		return rsp, err
	}
	if !gjson.ValidBytes(respBytes) {
		return rsp, errors.New("not a valid json response")
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}
