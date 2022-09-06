package sui

import (
	"context"
	"encoding/json"
	"github.com/block-vision/sui-go-sdk/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui_error"
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
}

type suiReadTransactionFromSuiImpl struct {
	conn *httpconn.HttpConn
}

func (s *suiReadTransactionFromSuiImpl) GetRecentTransactions(ctx context.Context, req models.GetRecentTransactionRequest, opts ...interface{}) (models.GetRecentTransactionResponse, error) {
	var rsp models.GetRecentTransactionResponse
	reqList := make([]interface{}, 0)
	reqList = append(reqList, req.Count)
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getRecentTransactions",
		Params: reqList,
	})
	if err != nil {
		return models.GetRecentTransactionResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetRecentTransactionResponse{}, sui_error.ErrInvalidJson
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

func (s *suiReadTransactionFromSuiImpl) GetTotalTransactionNumber(ctx context.Context, req models.GetTotalTransactionNumberRequest, opts ...interface{}) (models.GetTotalTransactionNumberResponse, error) {
	var rsp models.GetTotalTransactionNumberResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getTotalTransactionNumber",
		Params: []interface{}{},
	})
	if err != nil {
		return models.GetTotalTransactionNumberResponse{}, err
	}
	rsp.TotalNumberOfTransaction = gjson.ParseBytes(respBytes).Get("result").Uint()
	return rsp, nil
}

func (s *suiReadTransactionFromSuiImpl) GetTransaction(ctx context.Context, req models.GetTransactionRequest, opts ...interface{}) (models.GetTransactionResponse, error) {
	var rsp models.GetTransactionResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getTransaction",
		Params: []interface{}{
			req.Digest,
		},
	})
	if err != nil {
		return models.GetTransactionResponse{}, err
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").Raw), &rsp)
	if err != nil {
		return models.GetTransactionResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadTransactionFromSuiImpl) GetTransactionsByInputObject(ctx context.Context, req models.GetTransactionsByInputObjectRequest, opts ...interface{}) (models.GetTransactionsByInputObjectResponse, error) {
	var rsp models.GetTransactionsByInputObjectResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getTransactionsByInputObject",
		Params: []interface{}{
			req.ObjectID,
		},
	})
	if err != nil {
		return models.GetTransactionsByInputObjectResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetTransactionsByInputObjectResponse{}, sui_error.ErrInvalidJson
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

func (s *suiReadTransactionFromSuiImpl) GetTransactionsByMoveFunction(ctx context.Context, req models.GetTransactionsByMoveFunctionRequest, opts ...interface{}) (models.GetTransactionsByMoveFunctionResponse, error) {
	var rsp models.GetTransactionsByMoveFunctionResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
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
	if !gjson.ValidBytes(respBytes) {
		return models.GetTransactionsByMoveFunctionResponse{}, sui_error.ErrInvalidJson
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

func (s *suiReadTransactionFromSuiImpl) GetTransactionsByMutatedObject(ctx context.Context, req models.GetTransactionsByMutatedObjectRequest, opts ...interface{}) (models.GetTransactionsByMutatedObjectResponse, error) {
	var rsp models.GetTransactionsByMutatedObjectResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getTransactionsByMutatedObject",
		Params: []interface{}{
			req.ObjectID,
		},
	})
	if err != nil {
		return models.GetTransactionsByMutatedObjectResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetTransactionsByMutatedObjectResponse{}, sui_error.ErrInvalidJson
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

func (s *suiReadTransactionFromSuiImpl) GetTransactionsFromAddress(ctx context.Context, req models.GetTransactionsFromAddressRequest, opts ...interface{}) (models.GetTransactionsFromAddressResponse, error) {
	var rsp models.GetTransactionsFromAddressResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getTransactionsFromAddress",
		Params: []interface{}{
			req.Addr,
		},
	})
	if err != nil {
		return models.GetTransactionsFromAddressResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetTransactionsFromAddressResponse{}, sui_error.ErrInvalidJson
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

func (s *suiReadTransactionFromSuiImpl) GetTransactionsInRange(ctx context.Context, req models.GetTransactionsInRangeRequest, opts ...interface{}) (models.GetTransactionsInRangeResponse, error) {
	var rsp models.GetTransactionsInRangeResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getTransactionsInRange",
		Params: []interface{}{
			req.Start,
			req.End,
		},
	})
	if err != nil {
		return models.GetTransactionsInRangeResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetTransactionsInRangeResponse{}, sui_error.ErrInvalidJson
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

func (s *suiReadTransactionFromSuiImpl) GetTransactionsToAddress(ctx context.Context, req models.GetTransactionsToAddressRequest, opts ...interface{}) (models.GetTransactionsToAddressResponse, error) {
	var rsp models.GetTransactionsToAddressResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getTransactionsToAddress",
		Params: []interface{}{
			req.Addr,
		},
	})

	if err != nil {
		return models.GetTransactionsToAddressResponse{}, err
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetTransactionsToAddressResponse{}, sui_error.ErrInvalidJson
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
