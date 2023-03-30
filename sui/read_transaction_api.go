package sui

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/shoshinsquare/sui-go-sdk/common/rpc_client"
	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

type IReadTransactionFromSuiAPI interface {
	GetRecentTransactions(ctx context.Context, req models.GetRecentTransactionRequest, opts ...interface{}) (models.GetRecentTransactionResponse, error)
	GetTotalTransactionNumber(ctx context.Context, req models.GetTotalTransactionNumberRequest, opts ...interface{}) (models.GetTotalTransactionNumberResponse, error)
	GetTransaction(ctx context.Context, req models.GetTransactionRequest, opts ...interface{}) (models.GetTransactionResponse, error)
	GetTransactionsByInputObject(ctx context.Context, req models.GetTransactionsByInputObjectRequest, opts ...interface{}) (models.GetTransactionsByInputObjectResponse, error)
	GetTransactionsByMoveFunction(ctx context.Context, req models.GetTransactionsByMoveFunctionRequest, opts ...interface{}) (models.GetTransactionsByMoveFunctionResponse, error)
	GetTransactionsByMutatedObject(ctx context.Context, req models.GetTransactionsByMutatedObjectRequest, opts ...interface{}) (models.GetTransactionsByMutatedObjectResponse, error)
	GetTransactionsFromAddress(ctx context.Context, req models.GetTransactionsFromAddressRequest, opts ...interface{}) (models.GetTransactionsFromAddressResponse, error)
	GetTransactionsInRange(ctx context.Context, req models.GetTransactionsInRangeRequest, opts ...interface{}) (models.GetTransactionsInRangeResponse, error)
	GetTransactionsToAddress(ctx context.Context, req models.GetTransactionsToAddressRequest, opts ...interface{}) (models.GetTransactionsToAddressResponse, error)
	GetTransactionAuthSigners(ctx context.Context, req models.GetTransactionAuthSignersRequest, opts ...interface{}) (models.GetTransactionAuthSignersResponse, error)
}

type suiReadTransactionFromSuiImpl struct {
	cli *rpc_client.RPCClient
}

// GetRecentTransactions implements method `sui_getRecentTransactions`.
// Returns an array of transactions' metadata
func (s *suiReadTransactionFromSuiImpl) GetRecentTransactions(ctx context.Context, req models.GetRecentTransactionRequest, opts ...interface{}) (models.GetRecentTransactionResponse, error) {
	var rsp models.GetRecentTransactionResponse
	reqList := make([]interface{}, 0)
	reqList = append(reqList, req.Count)
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getRecentTransactions",
		Params: reqList,
	})
	if err != nil {
		return models.GetRecentTransactionResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetRecentTransactionResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	results := gjson.ParseBytes(respBytes).Get("result").Array()
	for i := range results {
		if len(results[i].Array()) < 2 {
			continue
		}
		rsp.Result = append(rsp.Result, models.GetTransactionMetaData{
			GatewayTxSeqNumber: results[i].Array()[0].Uint(),
			TransactionDigest:  results[i].Array()[1].String(),
		})
	}
	return rsp, nil
}

// GetTotalTransactionNumber implements method `sui_getTotalTransactionNumber`
// Returns a number of total transactions
func (s *suiReadTransactionFromSuiImpl) GetTotalTransactionNumber(ctx context.Context, req models.GetTotalTransactionNumberRequest, opts ...interface{}) (models.GetTotalTransactionNumberResponse, error) {
	var rsp models.GetTotalTransactionNumberResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getTotalTransactionNumber",
		Params: []interface{}{},
	})
	if err != nil {
		return models.GetTotalTransactionNumberResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetTotalTransactionNumberResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	rsp.TotalNumberOfTransaction = gjson.ParseBytes(respBytes).Get("result").Uint()
	return rsp, nil
}

// GetTransaction implements method `sui_getTransaction`
// Returns detail info of the transaction
func (s *suiReadTransactionFromSuiImpl) GetTransaction(ctx context.Context, req models.GetTransactionRequest, opts ...interface{}) (models.GetTransactionResponse, error) {
	var rsp models.GetTransactionResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getTransaction",
		Params: []interface{}{
			req.Digest,
		},
	})
	if err != nil {
		return models.GetTransactionResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetTransactionResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").Raw), &rsp)
	if err != nil {
		return models.GetTransactionResponse{}, err
	}
	return rsp, nil
}

// GetTransactionsByInputObject implements method `sui_getTransactionsByInputObject`.
// Returns an array of transactions' metadata
func (s *suiReadTransactionFromSuiImpl) GetTransactionsByInputObject(ctx context.Context, req models.GetTransactionsByInputObjectRequest, opts ...interface{}) (models.GetTransactionsByInputObjectResponse, error) {
	var rsp models.GetTransactionsByInputObjectResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getTransactionsByInputObject",
		Params: []interface{}{
			req.ObjectID,
		},
	})
	if err != nil {
		return models.GetTransactionsByInputObjectResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetTransactionsByInputObjectResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	results := gjson.ParseBytes(respBytes).Get("result").Array()
	for i := range results {
		if len(results[i].Array()) < 2 {
			continue
		}
		rsp.Result = append(rsp.Result, models.GetTransactionMetaData{
			GatewayTxSeqNumber: results[i].Array()[0].Uint(),
			TransactionDigest:  results[i].Array()[1].String(),
		})
	}
	return rsp, nil
}

// GetTransactionsByMoveFunction implements method `sui_getTransactionsByInputObject`.
// Returns an array of transactions' metadata.
func (s *suiReadTransactionFromSuiImpl) GetTransactionsByMoveFunction(ctx context.Context, req models.GetTransactionsByMoveFunctionRequest, opts ...interface{}) (models.GetTransactionsByMoveFunctionResponse, error) {
	var rsp models.GetTransactionsByMoveFunctionResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getTransactionsByInputObject",
		Params: []interface{}{
			req.Package,
			req.Module,
			req.Function,
		},
	})
	if err != nil {
		return models.GetTransactionsByMoveFunctionResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetTransactionsByMoveFunctionResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	results := gjson.ParseBytes(respBytes).Get("result").Array()
	for i := range results {
		if len(results[i].Array()) < 2 {
			continue
		}
		rsp.Result = append(rsp.Result, models.GetTransactionMetaData{
			GatewayTxSeqNumber: results[i].Array()[0].Uint(),
			TransactionDigest:  results[i].Array()[1].String(),
		})
	}
	return rsp, nil
}

// GetTransactionsByMutatedObject implements method `sui_getTransactionsByMutatedObject`.
// Returns an array of transactions' metadata.
func (s *suiReadTransactionFromSuiImpl) GetTransactionsByMutatedObject(ctx context.Context, req models.GetTransactionsByMutatedObjectRequest, opts ...interface{}) (models.GetTransactionsByMutatedObjectResponse, error) {
	var rsp models.GetTransactionsByMutatedObjectResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getTransactionsByMutatedObject",
		Params: []interface{}{
			req.ObjectID,
		},
	})
	if err != nil {
		return models.GetTransactionsByMutatedObjectResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetTransactionsByMutatedObjectResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	results := gjson.ParseBytes(respBytes).Get("result").Array()
	for i := range results {
		if len(results[i].Array()) < 2 {
			continue
		}
		rsp.Result = append(rsp.Result, models.GetTransactionMetaData{
			GatewayTxSeqNumber: results[i].Array()[0].Uint(),
			TransactionDigest:  results[i].Array()[1].String(),
		})
	}
	return rsp, nil
}

// GetTransactionsFromAddress implements method `sui_getTransactionsFromAddress`.
// Returns an array of transactions' metadata.
func (s *suiReadTransactionFromSuiImpl) GetTransactionsFromAddress(ctx context.Context, req models.GetTransactionsFromAddressRequest, opts ...interface{}) (models.GetTransactionsFromAddressResponse, error) {
	var rsp models.GetTransactionsFromAddressResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getTransactionsFromAddress",
		Params: []interface{}{
			req.Addr,
		},
	})
	if err != nil {
		return models.GetTransactionsFromAddressResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetTransactionsFromAddressResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	results := gjson.ParseBytes(respBytes).Get("result").Array()
	for i := range results {
		if len(results[i].Array()) < 2 {
			continue
		}
		rsp.Result = append(rsp.Result, models.GetTransactionMetaData{
			GatewayTxSeqNumber: results[i].Array()[0].Uint(),
			TransactionDigest:  results[i].Array()[1].String(),
		})
	}
	return rsp, nil
}

// GetTransactionsInRange implements method `sui_getTransactionsInRange`.
// Returns an array of transactions' metadata.
func (s *suiReadTransactionFromSuiImpl) GetTransactionsInRange(ctx context.Context, req models.GetTransactionsInRangeRequest, opts ...interface{}) (models.GetTransactionsInRangeResponse, error) {
	var rsp models.GetTransactionsInRangeResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getTransactionsInRange",
		Params: []interface{}{
			req.Start,
			req.End,
		},
	})
	if err != nil {
		return models.GetTransactionsInRangeResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetTransactionsInRangeResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	results := gjson.ParseBytes(respBytes).Get("result").Array()
	for i := range results {
		if len(results[i].Array()) < 2 {
			continue
		}
		rsp.Result = append(rsp.Result, models.GetTransactionMetaData{
			GatewayTxSeqNumber: results[i].Array()[0].Uint(),
			TransactionDigest:  results[i].Array()[1].String(),
		})
	}
	return rsp, nil
}

// GetTransactionsToAddress implements method `sui_getTransactionsToAddress`.
// Returns an array of transactions' metadata.
func (s *suiReadTransactionFromSuiImpl) GetTransactionsToAddress(ctx context.Context, req models.GetTransactionsToAddressRequest, opts ...interface{}) (models.GetTransactionsToAddressResponse, error) {
	var rsp models.GetTransactionsToAddressResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getTransactionsToAddress",
		Params: []interface{}{
			req.Addr,
		},
	})

	if err != nil {
		return models.GetTransactionsToAddressResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetTransactionsToAddressResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	results := gjson.ParseBytes(respBytes).Get("result").Array()
	for i := range results {
		if len(results[i].Array()) < 2 {
			continue
		}
		rsp.Result = append(rsp.Result, models.GetTransactionMetaData{
			GatewayTxSeqNumber: results[i].Array()[0].Uint(),
			TransactionDigest:  results[i].Array()[1].String(),
		})
	}
	return rsp, nil
}

func (s *suiReadTransactionFromSuiImpl) GetTransactionAuthSigners(ctx context.Context, req models.GetTransactionAuthSignersRequest, opts ...interface{}) (models.GetTransactionAuthSignersResponse, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getTransactionAuthSigners",
		Params: []interface{}{
			req.Digest,
		},
	})
	if err != nil {
		return models.GetTransactionAuthSignersResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetTransactionAuthSignersResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetTransactionAuthSignersResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetTransactionAuthSignersResponse{}, err
	}
	return rsp, nil
}
