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

type IReadSystemFromSuiAPI interface {
	SuiGetCheckpoint(ctx context.Context, req models.SuiGetCheckpointRequest) (models.CheckpointResponse, error)
	SuiGetCheckpoints(ctx context.Context, req models.SuiGetCheckpointsRequest) (models.PaginatedCheckpointsResponse, error)
	SuiGetLatestCheckpointSequenceNumber(ctx context.Context) (uint64, error)
	SuiXGetReferenceGasPrice(ctx context.Context) (uint64, error)
}

type suiReadSystemFromSuiImpl struct {
	conn *httpconn.HttpConn
}

// SuiGetCheckpoint implements the method `sui_getCheckpoint`, gets a checkpoint.
func (s *suiReadSystemFromSuiImpl) SuiGetCheckpoint(ctx context.Context, req models.SuiGetCheckpointRequest) (models.CheckpointResponse, error) {
	var rsp models.CheckpointResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getCheckpoint",
		Params: []interface{}{
			req.CheckpointID,
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

// SuiGetCheckpoints implements the method `sui_getCheckpoints`, gets paginated list of checkpoints.
func (s *suiReadSystemFromSuiImpl) SuiGetCheckpoints(ctx context.Context, req models.SuiGetCheckpointsRequest) (models.PaginatedCheckpointsResponse, error) {
	var rsp models.PaginatedCheckpointsResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getCheckpoints",
		Params: []interface{}{
			req.Cursor,
			req.Limit,
			req.DescendingOrder,
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

// SuiGetLatestCheckpointSequenceNumber implements the method `sui_getLatestCheckpointSequenceNumber`, gets the sequence number of the latest checkpoint that has been executed.
func (s *suiReadSystemFromSuiImpl) SuiGetLatestCheckpointSequenceNumber(ctx context.Context) (uint64, error) {
	var rsp uint64
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getLatestCheckpointSequenceNumber",
		Params: []interface{}{},
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

// SuiXGetReferenceGasPrice implements the method `suix_getReferenceGasPrice`, gets the reference gas price for the network.
func (s *suiReadSystemFromSuiImpl) SuiXGetReferenceGasPrice(ctx context.Context) (uint64, error) {
	var rsp uint64
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_getReferenceGasPrice",
		Params: []interface{}{},
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
