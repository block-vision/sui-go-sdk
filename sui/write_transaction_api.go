package sui

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/shoshinsquare/sui-go-sdk/common/rpc_client"
	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

var _ IWriteTransactionAPI = (*suiWriteTransactionImpl)(nil)

type IWriteTransactionAPI interface {
	BatchTransaction(ctx context.Context, req models.BatchTransactionRequest, opts ...interface{}) (models.BatchTransactionResponse, error)
	DryRunTransaction(ctx context.Context, req models.DryRunTransactionRequest, opts ...interface{}) (models.DryRunTransactionResponse, error)
	ExecuteTransaction(ctx context.Context, req models.ExecuteTransactionRequest, opts ...interface{}) (models.ExecuteTransactionResponse, error)

	Pay(ctx context.Context, req models.PayRequest, opts ...interface{}) (models.PayResponse, error)
	PayAllSui(ctx context.Context, req models.PayAllSuiRequest, opts ...interface{}) (models.PayAllSuiResponse, error)
	PaySui(ctx context.Context, req models.PaySuiRequest, opts ...interface{}) (models.PaySuiResponse, error)

	MoveCall(ctx context.Context, req models.MoveCallRequest, opts ...interface{}) (models.MoveCallResponse, error)
	MergeCoins(ctx context.Context, req models.MergeCoinsRequest, opts ...interface{}) (models.MergeCoinsResponse, error)
	SplitCoin(ctx context.Context, req models.SplitCoinRequest, opts ...interface{}) (models.SplitCoinResponse, error)
	SplitCoinEqual(ctx context.Context, req models.SplitCoinEqualRequest, opt ...interface{}) (models.SplitCoinEqualResponse, error)

	Publish(ctx context.Context, req models.PublishRequest, opts ...interface{}) (models.PublishResponse, error)
	TransferObject(ctx context.Context, req models.TransferObjectRequest, opts ...interface{}) (models.TransferObjectResponse, error)
	TransferSui(ctx context.Context, req models.TransferSuiRequest, opts ...interface{}) (models.TransferSuiResponse, error)

	MintNFT(ctx context.Context, req models.MintNFTRequest, opt ...interface{}) (models.MoveCallResponse, error)

	RequestAddDelegation(ctx context.Context, req models.RequestAddDelegationRequest, opts ...interface{}) (models.RequestAddDelegationResponse, error)
	RequestSwitchDelegation(ctx context.Context, req models.RequestSwitchDelegationRequest, opts ...interface{}) (models.RequestSwitchDelegationResponse, error)
	RequestWithdrawDelegation(ctx context.Context, req models.RequestWithdrawDelegationRequest, opts ...interface{}) (models.ReuqestWithdrawDelegationResponse, error)
}

type suiWriteTransactionImpl struct {
	cli *rpc_client.RPCClient
}

func (s *suiWriteTransactionImpl) MoveCall(ctx context.Context, req models.MoveCallRequest, opts ...interface{}) (models.MoveCallResponse, error) {
	var rsp models.MoveCallResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
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
	respBytes, err := s.cli.Request(ctx, models.Operation{
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
	respBytes, err := s.cli.Request(ctx, models.Operation{
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
	respBytes, err := s.cli.Request(ctx, models.Operation{
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
	respBytes, err := s.cli.Request(ctx, models.Operation{
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
	respBytes, err := s.cli.Request(ctx, models.Operation{
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
	respBytes, err := s.cli.Request(ctx, models.Operation{
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
	respBytes, err := s.cli.Request(ctx, models.Operation{
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
	respBytes, err := s.cli.Request(ctx, models.Operation{
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

func (s *suiWriteTransactionImpl) DryRunTransaction(ctx context.Context, req models.DryRunTransactionRequest, opts ...interface{}) (models.DryRunTransactionResponse, error) {
	var rsp models.DryRunTransactionResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_dryRunTransaction",
		Params: []interface{}{
			req.TxBytes,
			req.SigScheme,
			req.Signature,
			req.PubKey,
		},
	})
	if err != nil {
		return models.DryRunTransactionResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.DryRunTransactionResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.DryRunTransactionResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) Pay(ctx context.Context, req models.PayRequest, opts ...interface{}) (models.PayResponse, error) {
	var rsp models.PayResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_pay",
		Params: []interface{}{
			req.Signer,
			req.InputCoins,
			req.Recipient,
			req.Amounts,
			req.Gas,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.PayResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.PayResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.PayResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) PayAllSui(ctx context.Context, req models.PayAllSuiRequest, opts ...interface{}) (models.PayAllSuiResponse, error) {
	var rsp models.PayAllSuiResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_payAllSui",
		Params: []interface{}{
			req.Signer,
			req.InputCoins,
			req.Recipient,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.PayAllSuiResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.PayAllSuiResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.PayAllSuiResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) PaySui(ctx context.Context, req models.PaySuiRequest, opts ...interface{}) (models.PaySuiResponse, error) {
	var rsp models.PaySuiResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_paySui",
		Params: []interface{}{
			req.Signer,
			req.InputCoins,
			req.Recipient,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.PaySuiResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.PaySuiResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.PaySuiResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) MintNFT(ctx context.Context, req models.MintNFTRequest, opt ...interface{}) (models.MoveCallResponse, error) {
	moveCallReq := models.MoveCallRequest{
		Signer:          req.Signer,
		PackageObjectId: "0x0000000000000000000000000000000000000002",
		Module:          "devnet_nft",
		Function:        "mint",
		TypeArguments:   []interface{}{},
		Arguments:       []interface{}{req.NFTName, req.NFTDescription, req.NFTUrl},
		Gas:             req.GasObject,
		GasBudget:       req.GasBudget,
	}
	var rsp models.MoveCallResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_moveCall",
		Params: []interface{}{
			moveCallReq.Signer,
			moveCallReq.PackageObjectId,
			moveCallReq.Module,
			moveCallReq.Function,
			moveCallReq.TypeArguments,
			moveCallReq.Arguments,
			moveCallReq.Gas,
			moveCallReq.GasBudget,
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

func (s *suiWriteTransactionImpl) RequestAddDelegation(ctx context.Context, req models.RequestAddDelegationRequest, opts ...interface{}) (models.RequestAddDelegationResponse, error) {
	var rsp models.RequestAddDelegationResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_requestAddDelegation",
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
		return models.RequestAddDelegationResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.RequestAddDelegationResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.RequestAddDelegationResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) RequestSwitchDelegation(ctx context.Context, req models.RequestSwitchDelegationRequest, opts ...interface{}) (models.RequestSwitchDelegationResponse, error) {
	var rsp models.RequestSwitchDelegationResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_requestSwitchDelegation",
		Params: []interface{}{
			req.Signer,
			req.Delegation,
			req.StakedSui,
			req.NewValidatorAddress,
			req.Gas,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.RequestSwitchDelegationResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.RequestSwitchDelegationResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.RequestSwitchDelegationResponse{}, err
	}
	return rsp, nil
}

func (s *suiWriteTransactionImpl) RequestWithdrawDelegation(ctx context.Context, req models.RequestWithdrawDelegationRequest, opts ...interface{}) (models.ReuqestWithdrawDelegationResponse, error) {
	var rsp models.ReuqestWithdrawDelegationResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_requestWithdrawDelegation",
		Params: []interface{}{
			req.Signer,
			req.Delegation,
			req.StakedSui,
			req.PrincipalWithdrawAmount,
			req.Gas,
			req.GasBudget,
		},
	})
	if err != nil {
		return models.ReuqestWithdrawDelegationResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.ReuqestWithdrawDelegationResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.ReuqestWithdrawDelegationResponse{}, err
	}
	return rsp, nil
}
