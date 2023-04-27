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

type IWriteTransactionAPI interface {
	SuiExecuteTransactionBlock(ctx context.Context, req models.SuiExecuteTransactionBlockRequest) (models.ExecuteTransactionResponse, error)
	MoveCall(ctx context.Context, req models.MoveCallRequest) (models.MoveCallResponse, error)
}

type suiWriteTransactionImpl struct {
	conn *httpconn.HttpConn
}

// SuiExecuteTransactionBlock implements the method `sui_executeTransactionBlock`, executes a transaction using the transaction data and signature(s).
func (s *suiWriteTransactionImpl) SuiExecuteTransactionBlock(ctx context.Context, req models.SuiExecuteTransactionBlockRequest) (models.ExecuteTransactionResponse, error) {
	var rsp models.ExecuteTransactionResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_executeTransactionBlock",
		Params: []interface{}{
			req.TxBytes,
			req.Signature,
			req.Options,
			req.RequestType,
		},
	})
	if err != nil {
		return rsp, err
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

// MoveCall implements the method `unsafe_moveCall`, creates an unsigned transaction to execute a Move call on the network, by calling the specified function in the module of a given package.
func (s *suiWriteTransactionImpl) MoveCall(ctx context.Context, req models.MoveCallRequest) (models.MoveCallResponse, error) {
	var rsp models.MoveCallResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_moveCall",
		Params: []interface{}{
			req.Signer,
			req.PackageObjectId,
			req.Module,
			req.Function,
			req.TypeArguments,
			req.Arguments,
			req.Gas,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.MoveCallResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.MoveCallResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.MoveCallResponse{}, err
	}
	return rsp, nil
}
