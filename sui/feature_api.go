package sui

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/block-vision/sui-go-sdk/httpconn"
	"github.com/block-vision/sui-go-sdk/keypair"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui_error"
	"github.com/tidwall/gjson"
)

type IFeatureSuiAPI interface {
	MoveCallAndExecuteTransaction(ctx context.Context, req models.MoveCallAndExecuteTransactionRequest, opts ...interface{}) (models.MoveCallAndExecuteTransactionResponse, error)
	BatchAndExecuteTransaction(ctx context.Context, req models.BatchAndExecuteTransactionRequest, opts ...interface{}) (models.BatchAndExecuteTransactionResponse, error)
	SignWithAddress(ctx context.Context, address string, msg []byte) ([]byte, error)
	Sign(ctx context.Context, msg []byte) ([]byte, error)
}

type suiFeatureImpl struct {
	conn *httpconn.HttpConn
}

// MoveCallAndExecuteTransaction is a combination of `sui_moveCall` and `sui_executeTransaction`.
// This function free you from setting parameters when you want to execute the transaction of previous `sui_moveCall`
// but you need to `SetAccountKeyStore` first otherwise you cannot sign your transaction
func (s *suiFeatureImpl) MoveCallAndExecuteTransaction(ctx context.Context, req models.MoveCallAndExecuteTransactionRequest, opts ...interface{}) (models.MoveCallAndExecuteTransactionResponse, error) {
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
		return models.MoveCallAndExecuteTransactionResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.MoveCallAndExecuteTransactionResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var moveCallRsp models.MoveCallResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &moveCallRsp)
	if err != nil {
		return models.MoveCallAndExecuteTransactionResponse{}, err
	}
	if req.Signer == "" {
		req.Signer = accountStore.defaultAddress
	}
	_keypair, err := accountStore.GetKey(req.Signer)
	if err != nil {
		return models.MoveCallAndExecuteTransactionResponse{}, err
	}
	msg, err := base64.StdEncoding.DecodeString(moveCallRsp.TxBytes)
	if err != nil {
		return models.MoveCallAndExecuteTransactionResponse{}, err
	}
	signature, err := accountStore.Sign(req.Signer, []byte(msg))
	if err != nil {
		return models.MoveCallAndExecuteTransactionResponse{}, err
	}
	var scheme string
	if _keypair.Flag == keypair.Ed25519Flag {
		scheme = "ED25519"
	} else if _keypair.Flag == keypair.Secp256k1Flag {
		scheme = "Secp256k1"
	} else {
		return models.MoveCallAndExecuteTransactionResponse{}, sui_error.ErrUnknownSignatureScheme
	}
	respBytes, err = s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_executeTransaction",
		Params: []interface{}{
			moveCallRsp.TxBytes,
			scheme,
			signature,
			_keypair.PublicKeyBase64,
		},
	})
	var rsp models.MoveCallAndExecuteTransactionResponse
	if err != nil {
		return models.MoveCallAndExecuteTransactionResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.MoveCallAndExecuteTransactionResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.MoveCallAndExecuteTransactionResponse{}, err
	}
	return rsp, nil
}

// BatchAndExecuteTransaction is a combination of `sui_batchTransaction` and `sui_executeTransaction`.
// This function free you from setting parameters when you want to execute the transactions of previous `sui_batchTransaction`
// but you need to `SetAccountKeyStore` first otherwise you cannot sign your transactions
func (s *suiFeatureImpl) BatchAndExecuteTransaction(ctx context.Context, req models.BatchAndExecuteTransactionRequest, opts ...interface{}) (models.BatchAndExecuteTransactionResponse, error) {
	var batchTxRsp models.BatchTransactionResponse
	batchTxBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_batchTransaction",
		Params: []interface{}{
			req.Signer,
			req.SingleTransactionParams,
			req.Gas,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.BatchAndExecuteTransactionResponse{}, err
	}
	if gjson.ParseBytes(batchTxBytes).Get("error").Exists() {
		return models.BatchAndExecuteTransactionResponse{}, errors.New(gjson.ParseBytes(batchTxBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(batchTxBytes).Get("result").String()), &batchTxRsp)
	if err != nil {
		return models.BatchAndExecuteTransactionResponse{}, err
	}
	if req.Signer == "" {
		req.Signer = accountStore.defaultAddress
	}
	_keypair, err := accountStore.GetKey(req.Signer)
	if err != nil {
		return models.BatchAndExecuteTransactionResponse{}, err
	}
	msg, err := base64.StdEncoding.DecodeString(batchTxRsp.TxBytes)
	if err != nil {
		return models.BatchAndExecuteTransactionResponse{}, err
	}
	signature, err := accountStore.Sign(req.Signer, msg)
	if err != nil {
		return models.BatchAndExecuteTransactionResponse{}, err
	}
	var scheme string
	if _keypair.Flag == keypair.Ed25519Flag {
		scheme = "ED25519"
	} else if _keypair.Flag == keypair.Secp256k1Flag {
		scheme = "SECP256K1"
	} else {
		return models.BatchAndExecuteTransactionResponse{}, sui_error.ErrUnknownSignatureScheme
	}
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_executeTransaction",
		Params: []interface{}{
			batchTxRsp.TxBytes,
			scheme,
			signature,
			_keypair.PublicKeyBase64,
		},
	})
	var rsp models.BatchAndExecuteTransactionResponse
	if err != nil {
		return models.BatchAndExecuteTransactionResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.BatchAndExecuteTransactionResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.BatchAndExecuteTransactionResponse{}, err
	}
	return rsp, nil
}

// SignWithAddress implements Sign method.
// `address` you want to sign with MUST be one of addresses in your `sui.keystore` file you load before.
// Use it after you load the `sui.keystore` file
func (s *suiFeatureImpl) SignWithAddress(ctx context.Context, address string, msg []byte) ([]byte, error) {
	return accountStore.Sign(address, msg)
}

// Sign implements Sign Method with default address.
// It will use your first key in `sui.keystore` file as your default address.
// Use it after you load the `sui.keystore` file
func (s *suiFeatureImpl) Sign(ctx context.Context, msg []byte) ([]byte, error) {
	return accountStore.Sign(accountStore.defaultAddress, msg)
}
