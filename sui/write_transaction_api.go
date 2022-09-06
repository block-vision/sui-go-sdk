package sui

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/block-vision/sui-go-sdk/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

type IWriteTransactionAPI interface {
	MoveCall(ctx context.Context, req models.MoveCallRequest, opts ...interface{}) (models.MoveCallResponse, error)
	MergeCoins(ctx context.Context, req models.MergeCoinsRequest, opts ...interface{}) (models.MergeCoinsResponse, error)
	SplitCoin(ctx context.Context, req models.SplitCoinRequest, opts ...interface{}) (models.SplitCoinResponse, error)
	SplitCoinEqual(ctx context.Context, req models.SplitCoinEqualRequest, opt ...interface{}) (models.SplitCoinEqualResponse, error)
	Publish(ctx context.Context, req models.PublishRequest, opts ...interface{}) (models.PublishResponse, error)
	TransferObject(ctx context.Context, req models.TransferObjectRequest, opts ...interface{}) (models.TransferObjectResponse, error)
	TransferSui(ctx context.Context, req models.TransferSuiRequest, opts ...interface{}) (models.TransferSuiResponse, error)
	BatchTransaction(ctx context.Context, req models.BatchTransactionRequest, opts ...interface{}) (models.BatchTransactionResponse, error)
	ExecuteTransaction(ctx context.Context, req models.ExecuteTransactionRequest, opts ...interface{}) (models.ExecuteTransactionResponse, error)
}

type suiWriteTransactionImpl struct {
	conn *httpconn.HttpConn
}

func (s *suiWriteTransactionImpl) MoveCall(ctx context.Context, req models.MoveCallRequest, opts ...interface{}) (models.MoveCallResponse, error) {
	var rsp models.MoveCallResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_moveCall",
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

func (s *suiWriteTransactionImpl) MergeCoins(ctx context.Context, req models.MergeCoinsRequest, opts ...interface{}) (models.MergeCoinsResponse, error) {
	var rsp models.MergeCoinsResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_mergeCoins",
		Params: []interface{}{
			req.Signer,
			req.PrimaryCoin,
			req.CoinToMerge,
			req.Gas,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.MergeCoinsResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.MergeCoinsResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.MergeCoinsResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) SplitCoin(ctx context.Context, req models.SplitCoinRequest, opts ...interface{}) (models.SplitCoinResponse, error) {
	var rsp models.SplitCoinResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_splitCoin",
		Params: []interface{}{
			req.Signer,
			req.CoinObjectId,
			req.SplitAmounts,
			req.Gas,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.SplitCoinResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.SplitCoinResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.SplitCoinResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) SplitCoinEqual(ctx context.Context, req models.SplitCoinEqualRequest, opts ...interface{}) (models.SplitCoinEqualResponse, error) {
	var rsp models.SplitCoinEqualResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_splitCoinEqual",
		Params: []interface{}{
			req.Signer,
			req.CoinObjectId,
			req.SplitCount,
			req.Gas,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.SplitCoinEqualResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.SplitCoinEqualResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.SplitCoinEqualResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) Publish(ctx context.Context, req models.PublishRequest, opts ...interface{}) (models.PublishResponse, error) {
	var rsp models.PublishResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_publish",
		Params: []interface{}{
			req.Sender,
			req.CompiledModules,
			req.Gas,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.PublishResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.PublishResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.PublishResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) TransferObject(ctx context.Context, req models.TransferObjectRequest, opts ...interface{}) (models.TransferObjectResponse, error) {
	var rsp models.TransferObjectResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_transferObject",
		Params: []interface{}{
			req.Signer,
			req.ObjectId,
			req.Gas,
			req.GasBudget,
			req.Recipient,
		},
	})
	if err != nil {
		return models.TransferObjectResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.TransferObjectResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.TransferObjectResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) TransferSui(ctx context.Context, req models.TransferSuiRequest, opts ...interface{}) (models.TransferSuiResponse, error) {
	var rsp models.TransferSuiResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_transferSui",
		Params: []interface{}{
			req.Signer,
			req.SuiObjectId,
			req.GasBudget,
			req.Recipient,
			req.Amount,
		},
	})
	if err != nil {
		return models.TransferSuiResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.TransferSuiResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.TransferSuiResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) BatchTransaction(ctx context.Context, req models.BatchTransactionRequest, opts ...interface{}) (models.BatchTransactionResponse, error) {
	var rsp models.BatchTransactionResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_batchTransaction",
		Params: []interface{}{
			req.Signer,
			req.SingleTransactionParams,
			req.Gas,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.BatchTransactionResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.BatchTransactionResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.BatchTransactionResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) ExecuteTransaction(ctx context.Context, req models.ExecuteTransactionRequest, opts ...interface{}) (models.ExecuteTransactionResponse, error) {
	var rsp models.ExecuteTransactionResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_executeTransaction",
		Params: []interface{}{
			req.TxBytes,
			req.SigScheme,
			req.Signature,
			req.PubKey,
		},
	})
	if err != nil {
		return models.ExecuteTransactionResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.ExecuteTransactionResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.ExecuteTransactionResponse{}, err
	}
	return rsp, nil
}
