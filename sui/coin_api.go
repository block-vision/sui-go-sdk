package sui

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/shoshinsquare/sui-go-sdk/common/rpc_client"
	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

var _ ICoinAPI = (*SuiCoinImpl)(nil)

type ICoinAPI interface {
	GetCoins(ctx context.Context, req models.GetCoinsRequeset, opts ...interface{}) (models.GetCoinsResponse, error)
	GetAllBalances(ctx context.Context, req models.GetAllBalancesRequest, opts ...interface{}) (models.GetAllBalancesResponse, error)
	GetAllCoins(ctx context.Context, req models.GetAllCoinsRequest, opts ...interface{}) (models.GetAllCoinsResponse, error)
	GetBalance(ctx context.Context, req models.GetBalanceRequest, opts ...interface{}) (models.GetBalanceResponse, error)
}

type SuiCoinImpl struct {
	cli *rpc_client.RPCClient
}

func (s *SuiCoinImpl) GetCoins(ctx context.Context, req models.GetCoinsRequeset, opts ...interface{}) (models.GetCoinsResponse, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "suix_getCoins",
		Params: []interface{}{
			req.Owner,
			req.CoinType,
			req.Cursor,
			req.Limit,
		},
	})
	if err != nil {
		return models.GetCoinsResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetCoinsResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetCoinsResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetCoinsResponse{}, err
	}
	return rsp, nil
}

func (s *SuiCoinImpl) GetAllCoins(ctx context.Context, req models.GetAllCoinsRequest, opts ...interface{}) (models.GetAllCoinsResponse, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "suix_getAllCoins",
		Params: []interface{}{
			req.Owner,
			req.Cursor,
			req.Limit,
		},
	})
	if err != nil {
		return models.GetAllCoinsResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetAllCoinsResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetAllCoinsResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetAllCoinsResponse{}, err
	}
	return rsp, nil
}

func (s *SuiCoinImpl) GetAllBalances(ctx context.Context, req models.GetAllBalancesRequest, opts ...interface{}) (models.GetAllBalancesResponse, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getAllBalances",
		Params: []interface{}{
			req.Owner,
		},
	})
	if err != nil {
		return models.GetAllBalancesResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetAllBalancesResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetAllBalancesResponse
	var arr []models.Balance
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &arr)
	if err != nil {
		return models.GetAllBalancesResponse{}, err
	}
	rsp.Balance = arr
	return rsp, nil
}

func (s *SuiCoinImpl) GetBalance(ctx context.Context, req models.GetBalanceRequest, opts ...interface{}) (models.GetBalanceResponse, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "suix_getBalance",
		Params: []interface{}{
			req.Owner,
			req.CoinType,
		},
	})
	if err != nil {
		return models.GetBalanceResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetBalanceResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetBalanceResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetBalanceResponse{}, err
	}
	//rsp.Balance = arr
	return rsp, nil
}
