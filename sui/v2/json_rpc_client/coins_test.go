// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package json_rpc_client

import (
	"testing"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui/v2/types"
)

func TestConvertBalanceSplitsAddressBalance(t *testing.T) {
	sui := normalizeCoinType("0x2::sui::SUI")
	cases := []struct {
		in   models.CoinBalanceResponse
		want types.Balance
	}{
		{ // mainnet 0xc38e… SUI: totalBalance includes fundsInAddressBalance
			models.CoinBalanceResponse{CoinType: "0x2::sui::SUI", TotalBalance: "5858910487314795", FundsInAddressBalance: "5858808487514795"},
			types.Balance{CoinType: sui, Balance_: "5858910487314795", CoinBalance: "101999800000", AddressBalance: "5858808487514795"},
		},
		{ // node without address balance support omits the field
			models.CoinBalanceResponse{CoinType: "0x2::sui::SUI", TotalBalance: "42"},
			types.Balance{CoinType: sui, Balance_: "42", CoinBalance: "42", AddressBalance: "0"},
		},
	}
	for _, c := range cases {
		if got := convertBalance(c.in); got != c.want {
			t.Errorf("got %#v, want %#v", got, c.want)
		}
	}
}
