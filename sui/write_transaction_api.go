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
	SuiExecuteTransactionBlock(ctx context.Context, req models.SuiExecuteTransactionBlockRequest) (models.SuiTransactionBlockResponse, error)
	MoveCall(ctx context.Context, req models.MoveCallRequest) (models.TxnMetaData, error)
	MergeCoins(ctx context.Context, req models.MergeCoinsRequest) (models.TxnMetaData, error)
	SplitCoin(ctx context.Context, req models.SplitCoinRequest) (models.TxnMetaData, error)
	SplitCoinEqual(ctx context.Context, req models.SplitCoinEqualRequest) (models.TxnMetaData, error)
	Publish(ctx context.Context, req models.PublishRequest) (models.TxnMetaData, error)
	TransferObject(ctx context.Context, req models.TransferObjectRequest) (models.TxnMetaData, error)
	TransferSui(ctx context.Context, req models.TransferSuiRequest) (models.TxnMetaData, error)
	Pay(ctx context.Context, req models.PayRequest) (models.TxnMetaData, error)
	PaySui(ctx context.Context, req models.PaySuiRequest) (models.TxnMetaData, error)
	PayAllSui(ctx context.Context, req models.PayAllSuiRequest) (models.TxnMetaData, error)
	RequestAddStake(ctx context.Context, req models.AddStakeRequest) (models.TxnMetaData, error)
	RequestWithdrawStake(ctx context.Context, req models.WithdrawStakeRequest) (models.TxnMetaData, error)
	BatchTransaction(ctx context.Context, req models.BatchTransactionRequest) (models.BatchTransactionResponse, error)
	SignAndExecuteTransactionBlock(ctx context.Context, req models.SignAndExecuteTransactionBlockRequest) (models.SuiTransactionBlockResponse, error)
}

type suiWriteTransactionImpl struct {
	conn *httpconn.HttpConn
}

// SuiExecuteTransactionBlock implements the method `sui_executeTransactionBlock`, executes a transaction using the transaction data and signature(s).
func (s *suiWriteTransactionImpl) SuiExecuteTransactionBlock(ctx context.Context, req models.SuiExecuteTransactionBlockRequest) (models.SuiTransactionBlockResponse, error) {
	var rsp models.SuiTransactionBlockResponse
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
func (s *suiWriteTransactionImpl) MoveCall(ctx context.Context, req models.MoveCallRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
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

// MergeCoins implements the method `unsafe_mergeCoins`, creates an unsigned transaction to merge multiple coins into one coin.
func (s *suiWriteTransactionImpl) MergeCoins(ctx context.Context, req models.MergeCoinsRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_mergeCoins",
		Params: []interface{}{
			req.Signer,
			req.PrimaryCoin,
			req.CoinToMerge,
			req.Gas,
			req.GasBudget,
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

// SplitCoin implements the method `unsafe_splitCoin`, creates an unsigned transaction to split a coin object into multiple coins.
func (s *suiWriteTransactionImpl) SplitCoin(ctx context.Context, req models.SplitCoinRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_splitCoin",
		Params: []interface{}{
			req.Signer,
			req.CoinObjectId,
			req.SplitAmounts,
			req.Gas,
			req.GasBudget,
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

// SplitCoinEqual implements the method `unsafe_splitCoinEqual`, creates an unsigned transaction to split a coin object into multiple equal-size coins.
func (s *suiWriteTransactionImpl) SplitCoinEqual(ctx context.Context, req models.SplitCoinEqualRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_splitCoinEqual",
		Params: []interface{}{
			req.Signer,
			req.CoinObjectId,
			req.SplitCount,
			req.Gas,
			req.GasBudget,
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

// Publish implements the method `unsafe_publish`, creates an unsigned transaction to publish a Move package.
func (s *suiWriteTransactionImpl) Publish(ctx context.Context, req models.PublishRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_publish",
		Params: []interface{}{
			req.Sender,
			req.CompiledModules,
			req.Dependencies,
			req.Gas,
			req.GasBudget,
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

// TransferObject implements the method `unsafe_transferObject`, creates an unsigned transaction to transfer an object from one address to another. The object's type must allow public transfers.
func (s *suiWriteTransactionImpl) TransferObject(ctx context.Context, req models.TransferObjectRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_transferObject",
		Params: []interface{}{
			req.Signer,
			req.ObjectId,
			req.Gas,
			req.GasBudget,
			req.Recipient,
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

// TransferSui implements the method `unsafe_transferSui`, creates an unsigned transaction to send SUI coin object to a Sui address. The SUI object is also used as the gas object.
func (s *suiWriteTransactionImpl) TransferSui(ctx context.Context, req models.TransferSuiRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_transferSui",
		Params: []interface{}{
			req.Signer,
			req.SuiObjectId,
			req.GasBudget,
			req.Recipient,
			req.Amount,
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

// Pay implements the method `unsafe_pay`, send `Coin<T>` to a list of addresses, where `T` can be any coin type, following a list of amounts.
// The object specified in the `gas` field will be used to pay the gas fee for the transaction.
// The gas object can not appear in `input_coins`. If the gas object is not specified, the RPC server will auto-select one.
func (s *suiWriteTransactionImpl) Pay(ctx context.Context, req models.PayRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_pay",
		Params: []interface{}{
			req.Signer,
			req.SuiObjectId,
			req.Recipient,
			req.Amount,
			req.Gas,
			req.GasBudget,
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

// PaySui implements the method `unsafe_paySui`, send SUI coins to a list of addresses, following a list of amounts.
// This is for SUI coin only and does not require a separate gas coin object.
// Specifically, what pay_sui does are:
// 1. debit each input_coin to create new coin following the order of amounts and assign it to the corresponding recipient.
// 2. accumulate all residual SUI from input coins left and deposit all SUI to the first input coin, then use the first input coin as the gas coin object.
// 3. the balance of the first input coin after tx is sum(input_coins) - sum(amounts) - actual_gas_cost
// 4. all other input coints other than the first one are deleted.
func (s *suiWriteTransactionImpl) PaySui(ctx context.Context, req models.PaySuiRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_paySui",
		Params: []interface{}{
			req.Signer,
			req.SuiObjectId,
			req.Recipient,
			req.Amount,
			req.GasBudget,
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

// PayAllSui implements the method `unsafe_payAllSui`, send all SUI coins to one recipient.
// This is for SUI coin only and does not require a separate gas coin object.
// Specifically, what pay_all_sui does are:
// 1. accumulate all SUI from input coins and deposit all SUI to the first input coin.
// 2. transfer the updated first coin to the recipient and also use this first coin as gas coin object.
// 3. the balance of the first input coin after tx is sum(input_coins) - actual_gas_cost.
// 4. all other input coins other than the first are deleted.
func (s *suiWriteTransactionImpl) PayAllSui(ctx context.Context, req models.PayAllSuiRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_payAllSui",
		Params: []interface{}{
			req.Signer,
			req.SuiObjectId,
			req.Recipient,
			req.GasBudget,
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

// RequestAddStake implements the method `unsafe_requestAddStake`, add stake to a validator's staking pool using multiple coins and amount.
func (s *suiWriteTransactionImpl) RequestAddStake(ctx context.Context, req models.AddStakeRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_requestAddStake",
		Params: []interface{}{
			req.Signer,
			req.Coins,
			req.Amount,
			req.Validator,
			req.Gas,
			req.GasBudget,
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

// RequestWithdrawStake implements the method `unsafe_requestWithdrawStake`, withdraw stake from a validator's staking pool.
func (s *suiWriteTransactionImpl) RequestWithdrawStake(ctx context.Context, req models.WithdrawStakeRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_requestWithdrawStake",
		Params: []interface{}{
			req.Signer,
			req.StakedObjectId,
			req.Gas,
			req.GasBudget,
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

// BatchTransaction implements the method `unsafe_batchTransaction`, creates an unsigned batched transaction.
func (s *suiWriteTransactionImpl) BatchTransaction(ctx context.Context, req models.BatchTransactionRequest) (models.BatchTransactionResponse, error) {
	var rsp models.BatchTransactionResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_batchTransaction",
		Params: []interface{}{
			req.Signer,
			req.RPCTransactionRequestParams,
			req.Gas,
			req.GasBudget,
			req.SuiTransactionBlockBuilderMode,
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

// SignAndExecuteTransactionBlock sign a transaction block and submit to the Fullnode for execution.
func (s *suiWriteTransactionImpl) SignAndExecuteTransactionBlock(ctx context.Context, req models.SignAndExecuteTransactionBlockRequest) (models.SuiTransactionBlockResponse, error) {
	var rsp models.SuiTransactionBlockResponse

	signedTxn := req.TxnMetaData.SignSerializedSigWith(req.PriKey)

	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_executeTransactionBlock",
		Params: []interface{}{
			signedTxn.TxBytes,
			[]string{signedTxn.Signature},
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
