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

type IReadTransactionFromSuiAPI interface {
	SuiGetTotalTransactionBlocks(ctx context.Context) (uint64, error)
	SuiGetTransactionBlock(ctx context.Context, req models.SuiGetTransactionBlockRequest) (models.SuiTransactionBlockResponse, error)
	SuiMultiGetTransactionBlocks(ctx context.Context, req models.SuiMultiGetTransactionBlocksRequest) (models.SuiMultiGetTransactionBlocksResponse, error)
	SuiXQueryTransactionBlocks(ctx context.Context, req models.SuiXQueryTransactionBlocksRequest) (models.SuiXQueryTransactionBlocksResponse, error)
	SuiDryRunTransactionBlock(ctx context.Context, req models.SuiDryRunTransactionBlockRequest) (models.SuiTransactionBlockResponse, error)
	SuiDevInspectTransactionBlock(ctx context.Context, req models.SuiDevInspectTransactionBlockRequest) (models.SuiTransactionBlockResponse, error)
}

type suiReadTransactionFromSuiImpl struct {
	conn *httpconn.HttpConn
}

// SuiGetTotalTransactionBlocks implements the method `sui_getTotalTransactionBlocks`, gets the total number of transactions known to the node.
func (s *suiReadTransactionFromSuiImpl) SuiGetTotalTransactionBlocks(ctx context.Context) (uint64, error) {
	var rsp uint64
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getTotalTransactionBlocks",
		Params: []interface{}{},
	})
	if err != nil {
		return rsp, err
	}
	rsp = gjson.ParseBytes(respBytes).Get("result").Uint()
	return rsp, nil
}

// SuiGetTransactionBlock implements the method `sui_getTransactionBlock`, gets the transaction response object for a specified transaction digest.
func (s *suiReadTransactionFromSuiImpl) SuiGetTransactionBlock(ctx context.Context, req models.SuiGetTransactionBlockRequest) (models.SuiTransactionBlockResponse, error) {
	var rsp models.SuiTransactionBlockResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getTransactionBlock",
		Params: []interface{}{
			req.Digest,
			req.Options,
		},
	})
	if err != nil {
		return rsp, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").Raw), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiMultiGetTransactionBlocks implements the method `sui_multiGetTransactionBlocks`, gets an ordered list of transaction responses.
func (s *suiReadTransactionFromSuiImpl) SuiMultiGetTransactionBlocks(ctx context.Context, req models.SuiMultiGetTransactionBlocksRequest) (models.SuiMultiGetTransactionBlocksResponse, error) {
	var rsp models.SuiMultiGetTransactionBlocksResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_multiGetTransactionBlocks",
		Params: []interface{}{
			req.Digests,
			req.Options,
		},
	})
	if err != nil {
		return rsp, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").Raw), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXQueryTransactionBlocks implements the method `suix_queryTransactionBlocks`, gets list of transactions for a specified query criteria.
func (s *suiReadTransactionFromSuiImpl) SuiXQueryTransactionBlocks(ctx context.Context, req models.SuiXQueryTransactionBlocksRequest) (models.SuiXQueryTransactionBlocksResponse, error) {
	var rsp models.SuiXQueryTransactionBlocksResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_queryTransactionBlocks",
		Params: []interface{}{
			req.SuiTransactionBlockResponseQuery,
			req.Cursor,
			req.Limit,
			req.DescendingOrder,
		},
	})
	if err != nil {
		return rsp, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").Raw), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiDryRunTransactionBlock implements the method `sui_dryRunTransactionBlock`, gets transaction execution effects including the gas cost summary, while the effects are not committed to the chain.
func (s *suiReadTransactionFromSuiImpl) SuiDryRunTransactionBlock(ctx context.Context, req models.SuiDryRunTransactionBlockRequest) (models.SuiTransactionBlockResponse, error) {
	var rsp models.SuiTransactionBlockResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_dryRunTransactionBlock",
		Params: []interface{}{
			req.TxBytes,
		},
	})
	if err != nil {
		return rsp, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").Raw), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiDevInspectTransactionBlock implements the method `sui_devInspectTransactionBlock`, runs the transaction in dev-inspect mode.
// Which allows for nearly any transaction (or Move call) with any arguments.
// Detailed results are provided, including both the transaction effects and any return values.
func (s *suiReadTransactionFromSuiImpl) SuiDevInspectTransactionBlock(ctx context.Context, req models.SuiDevInspectTransactionBlockRequest) (models.SuiTransactionBlockResponse, error) {
	var rsp models.SuiTransactionBlockResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_devInspectTransactionBlock",
		Params: []interface{}{
			req.Sender,
			req.TxBytes,
			req.GasPrice,
			req.Epoch,
		},
	})
	if err != nil {
		return rsp, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").Raw), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}
