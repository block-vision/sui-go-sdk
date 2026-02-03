// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"
	"github.com/block-vision/sui-go-sdk/common/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
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
	handler *BaseRequestHandler
}

func newSuiWriteTransactionImpl(conn *httpconn.HttpConn) *suiWriteTransactionImpl {
	return &suiWriteTransactionImpl{
		handler: NewBaseRequestHandler(conn),
	}
}

// SuiExecuteTransactionBlock implements the method `sui_executeTransactionBlock`, executes a transaction using the transaction data and signature(s).
func (s *suiWriteTransactionImpl) SuiExecuteTransactionBlock(ctx context.Context, req models.SuiExecuteTransactionBlockRequest) (models.SuiTransactionBlockResponse, error) {
	var rsp models.SuiTransactionBlockResponse
	params := []interface{}{req.TxBytes, req.Signature, req.Options, req.RequestType}
	err := s.handler.ExecuteRequest(ctx, "sui_executeTransactionBlock", params, &rsp)
	return rsp, err
}

// MoveCall implements the method `unsafe_moveCall`, creates an unsigned transaction to execute a Move call on the network, by calling the specified function in the module of a given package.
func (s *suiWriteTransactionImpl) MoveCall(ctx context.Context, req models.MoveCallRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Signer, req.PackageObjectId, req.Module, req.Function, req.TypeArguments, req.Arguments, req.Gas, req.GasBudget}
	err := s.handler.ExecuteRequest(ctx, "unsafe_moveCall", params, &rsp)
	return rsp, err
}

// MergeCoins implements the method `unsafe_mergeCoins`, creates an unsigned transaction to merge multiple coins into one coin.
func (s *suiWriteTransactionImpl) MergeCoins(ctx context.Context, req models.MergeCoinsRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Signer, req.PrimaryCoin, req.CoinToMerge, req.Gas, req.GasBudget}
	err := s.handler.ExecuteRequest(ctx, "unsafe_mergeCoins", params, &rsp)
	return rsp, err
}

// SplitCoin implements the method `unsafe_splitCoin`, creates an unsigned transaction to split a coin object into multiple coins.
func (s *suiWriteTransactionImpl) SplitCoin(ctx context.Context, req models.SplitCoinRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Signer, req.CoinObjectId, req.SplitAmounts, req.Gas, req.GasBudget}
	err := s.handler.ExecuteRequest(ctx, "unsafe_splitCoin", params, &rsp)
	return rsp, err
}

// SplitCoinEqual implements the method `unsafe_splitCoinEqual`, creates an unsigned transaction to split a coin object into multiple equal-size coins.
func (s *suiWriteTransactionImpl) SplitCoinEqual(ctx context.Context, req models.SplitCoinEqualRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Signer, req.CoinObjectId, req.SplitCount, req.Gas, req.GasBudget}
	err := s.handler.ExecuteRequest(ctx, "unsafe_splitCoinEqual", params, &rsp)
	return rsp, err
}

// Publish implements the method `unsafe_publish`, creates an unsigned transaction to publish a Move package.
func (s *suiWriteTransactionImpl) Publish(ctx context.Context, req models.PublishRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Sender, req.CompiledModules, req.Dependencies, req.Gas, req.GasBudget}
	err := s.handler.ExecuteRequest(ctx, "unsafe_publish", params, &rsp)
	return rsp, err
}

// TransferObject implements the method `unsafe_transferObject`, creates an unsigned transaction to transfer an object from one address to another. The object's type must allow public transfers.
func (s *suiWriteTransactionImpl) TransferObject(ctx context.Context, req models.TransferObjectRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Signer, req.ObjectId, req.Gas, req.GasBudget, req.Recipient}
	err := s.handler.ExecuteRequest(ctx, "unsafe_transferObject", params, &rsp)
	return rsp, err
}

// TransferSui implements the method `unsafe_transferSui`, creates an unsigned transaction to send SUI coin object to a Sui address. The SUI object is also used as the gas object.
func (s *suiWriteTransactionImpl) TransferSui(ctx context.Context, req models.TransferSuiRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Signer, req.SuiObjectId, req.GasBudget, req.Recipient, req.Amount}
	err := s.handler.ExecuteRequest(ctx, "unsafe_transferSui", params, &rsp)
	return rsp, err
}

// Pay implements the method `unsafe_pay`, send `Coin<T>` to a list of addresses, where `T` can be any coin type, following a list of amounts.
func (s *suiWriteTransactionImpl) Pay(ctx context.Context, req models.PayRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Signer, req.SuiObjectId, req.Recipient, req.Amount, req.Gas, req.GasBudget}
	err := s.handler.ExecuteRequest(ctx, "unsafe_pay", params, &rsp)
	return rsp, err
}

// PaySui implements the method `unsafe_paySui`, send SUI coins to a list of addresses, following a list of amounts.
func (s *suiWriteTransactionImpl) PaySui(ctx context.Context, req models.PaySuiRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Signer, req.SuiObjectId, req.Recipient, req.Amount, req.GasBudget}
	err := s.handler.ExecuteRequest(ctx, "unsafe_paySui", params, &rsp)
	return rsp, err
}

// PayAllSui implements the method `unsafe_payAllSui`, send all SUI coins to one recipient.
func (s *suiWriteTransactionImpl) PayAllSui(ctx context.Context, req models.PayAllSuiRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Signer, req.SuiObjectId, req.Recipient, req.GasBudget}
	err := s.handler.ExecuteRequest(ctx, "unsafe_payAllSui", params, &rsp)
	return rsp, err
}

// RequestAddStake implements the method `unsafe_requestAddStake`, add stake to a validator's staking pool using multiple coins and amount.
func (s *suiWriteTransactionImpl) RequestAddStake(ctx context.Context, req models.AddStakeRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Signer, req.Coins, req.Amount, req.Validator, req.Gas, req.GasBudget}
	err := s.handler.ExecuteRequest(ctx, "unsafe_requestAddStake", params, &rsp)
	return rsp, err
}

// RequestWithdrawStake implements the method `unsafe_requestWithdrawStake`, withdraw stake from a validator's staking pool.
func (s *suiWriteTransactionImpl) RequestWithdrawStake(ctx context.Context, req models.WithdrawStakeRequest) (models.TxnMetaData, error) {
	var rsp models.TxnMetaData
	params := []interface{}{req.Signer, req.StakedObjectId, req.Gas, req.GasBudget}
	err := s.handler.ExecuteRequest(ctx, "unsafe_requestWithdrawStake", params, &rsp)
	return rsp, err
}

// BatchTransaction implements the method `unsafe_batchTransaction`, creates an unsigned batched transaction.
func (s *suiWriteTransactionImpl) BatchTransaction(ctx context.Context, req models.BatchTransactionRequest) (models.BatchTransactionResponse, error) {
	var rsp models.BatchTransactionResponse
	params := []interface{}{req.Signer, req.RPCTransactionRequestParams, req.Gas, req.GasBudget, req.SuiTransactionBlockBuilderMode}
	err := s.handler.ExecuteRequest(ctx, "unsafe_batchTransaction", params, &rsp)
	return rsp, err
}

// SignAndExecuteTransactionBlock sign a transaction block and submit to the Fullnode for execution.
func (s *suiWriteTransactionImpl) SignAndExecuteTransactionBlock(ctx context.Context, req models.SignAndExecuteTransactionBlockRequest) (models.SuiTransactionBlockResponse, error) {
	var rsp models.SuiTransactionBlockResponse
	signedTxn := req.TxnMetaData.SignSerializedSigWith(req.PriKey)
	params := []interface{}{signedTxn.TxBytes, []string{signedTxn.Signature}, req.Options, req.RequestType}
	err := s.handler.ExecuteRequest(ctx, "sui_executeTransactionBlock", params, &rsp)
	return rsp, err
}