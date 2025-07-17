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

type IReadCoinFromSuiAPI interface {
	SuiXGetBalance(ctx context.Context, req models.SuiXGetBalanceRequest) (models.CoinBalanceResponse, error)
	SuiXGetAllBalance(ctx context.Context, req models.SuiXGetAllBalanceRequest) (models.CoinAllBalanceResponse, error)
	SuiXGetCoins(ctx context.Context, req models.SuiXGetCoinsRequest) (models.PaginatedCoinsResponse, error)
	SuiXGetAllCoins(ctx context.Context, req models.SuiXGetAllCoinsRequest) (models.PaginatedCoinsResponse, error)
	SuiXGetCoinMetadata(ctx context.Context, req models.SuiXGetCoinMetadataRequest) (models.CoinMetadataResponse, error)
	SuiXGetTotalSupply(ctx context.Context, req models.SuiXGetTotalSupplyRequest) (models.TotalSupplyResponse, error)
}

type suiReadCoinFromSuiImpl struct {
	conn *httpconn.HttpConn
}

// SuiXGetBalance implements the method `suix_getBalance`, gets the total Coin balance for each coin type owned by an address.
func (s *suiReadCoinFromSuiImpl) SuiXGetBalance(ctx context.Context, req models.SuiXGetBalanceRequest) (models.CoinBalanceResponse, error) {
	var rsp models.CoinBalanceResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_getBalance",
		Params: []interface{}{
			req.Owner,
			req.CoinType,
		},
	})
	if err != nil {
		return rsp, err
	}
	parsedJson := gjson.ParseBytes(respBytes)
	if parsedJson.Get("error").Exists() {
		return rsp, errors.New(parsedJson.Get("error").String())
	}
	err = json.Unmarshal([]byte(parsedJson.Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXGetAllBalance implements the method `suix_getAllBalances`, gets all Coin balances owned by an address.
func (s *suiReadCoinFromSuiImpl) SuiXGetAllBalance(ctx context.Context, req models.SuiXGetAllBalanceRequest) (models.CoinAllBalanceResponse, error) {
	var rsp models.CoinAllBalanceResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_getAllBalances",
		Params: []interface{}{
			req.Owner,
		},
	})
	if err != nil {
		return rsp, err
	}
	parsedJson := gjson.ParseBytes(respBytes)
	if parsedJson.Get("error").Exists() {
		return rsp, errors.New(parsedJson.Get("error").String())
	}
	err = json.Unmarshal([]byte(parsedJson.Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXGetCoins implements the method `suix_getCoins`, gets a list of Coin objects by type owned by an address.
func (s *suiReadCoinFromSuiImpl) SuiXGetCoins(ctx context.Context, req models.SuiXGetCoinsRequest) (models.PaginatedCoinsResponse, error) {
	var rsp models.PaginatedCoinsResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_getCoins",
		Params: []interface{}{
			req.Owner,
			req.CoinType,
			req.Cursor,
			req.Limit,
		},
	})
	if err != nil {
		return rsp, err
	}
	parsedJson := gjson.ParseBytes(respBytes)
	if parsedJson.Get("error").Exists() {
		return rsp, errors.New(parsedJson.Get("error").String())
	}
	err = json.Unmarshal([]byte(parsedJson.Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXGetAllCoins implements the method `suix_getAllCoins`, gets all Coin objects owned by an address.
func (s *suiReadCoinFromSuiImpl) SuiXGetAllCoins(ctx context.Context, req models.SuiXGetAllCoinsRequest) (models.PaginatedCoinsResponse, error) {
	var rsp models.PaginatedCoinsResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_getAllCoins",
		Params: []interface{}{
			req.Owner,
			req.Cursor,
			req.Limit,
		},
	})
	if err != nil {
		return rsp, err
	}
	parsedJson := gjson.ParseBytes(respBytes)
	if parsedJson.Get("error").Exists() {
		return rsp, errors.New(parsedJson.Get("error").String())
	}
	err = json.Unmarshal([]byte(parsedJson.Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXGetCoinMetadata implements the method `suix_getCoinMetadata`, gets metadata(e.g., symbol, decimals) for a coin.
func (s *suiReadCoinFromSuiImpl) SuiXGetCoinMetadata(ctx context.Context, req models.SuiXGetCoinMetadataRequest) (models.CoinMetadataResponse, error) {
	var rsp models.CoinMetadataResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_getCoinMetadata",
		Params: []interface{}{
			req.CoinType,
		},
	})
	if err != nil {
		return rsp, err
	}
	parsedJson := gjson.ParseBytes(respBytes)
	if parsedJson.Get("error").Exists() {
		return rsp, errors.New(parsedJson.Get("error").String())
	}
	err = json.Unmarshal([]byte(parsedJson.Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXGetTotalSupply implements the method `suix_getTotalSupply`, gets total supply for a coin
func (s *suiReadCoinFromSuiImpl) SuiXGetTotalSupply(ctx context.Context, req models.SuiXGetTotalSupplyRequest) (models.TotalSupplyResponse, error) {
	var rsp models.TotalSupplyResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_getTotalSupply",
		Params: []interface{}{
			req.CoinType,
		},
	})
	if err != nil {
		return rsp, err
	}
	parsedJson := gjson.ParseBytes(respBytes)
	if parsedJson.Get("error").Exists() {
		return rsp, errors.New(parsedJson.Get("error").String())
	}
	err = json.Unmarshal([]byte(parsedJson.Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}
