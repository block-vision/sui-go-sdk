// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/utils"
	"testing"
)

var ctx = context.Background()
var cli = NewSuiClient(constant.BvTestnetEndpoint)

func TestOnReadSystemFromSui(t *testing.T) {
	t.Run("test on sui_getCheckpoint", func(t *testing.T) {
		rsp, err := cli.SuiGetCheckpoint(ctx, models.SuiGetCheckpointRequest{
			CheckpointID: "1628214",
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

	t.Run("test on sui_getCheckpoints", func(t *testing.T) {
		rsp, err := cli.SuiGetCheckpoints(ctx, models.SuiGetCheckpointsRequest{
			Cursor:          "13200",
			Limit:           5,
			DescendingOrder: true,
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		for _, checkpoint := range rsp.Data {
			utils.PrettyPrint(checkpoint)
		}
	})

	t.Run("test on sui_getLatestCheckpointSequenceNumber", func(t *testing.T) {
		rsp, err := cli.SuiGetLatestCheckpointSequenceNumber(ctx)

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

	t.Run("test on suix_getReferenceGasPrice", func(t *testing.T) {
		rsp, err := cli.SuiXGetReferenceGasPrice(ctx)

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

}

func TestOnReadCoinFromSui(t *testing.T) {

	t.Run("test on suix_getBalance", func(t *testing.T) {
		rsp, err := cli.SuiXGetBalance(ctx, models.SuiXGetBalanceRequest{
			Owner:    "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5",
			CoinType: "0x2::sui::SUI",
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

	t.Run("test on suix_getAllBalances", func(t *testing.T) {
		rsp, err := cli.SuiXGetAllBalance(ctx, models.SuiXGetAllBalanceRequest{
			Owner: "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5",
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

	t.Run("test on suix_getCoins", func(t *testing.T) {
		rsp, err := cli.SuiXGetCoins(ctx, models.SuiXGetCoinsRequest{
			Owner:    "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5",
			CoinType: "0x2::sui::SUI",
			Limit:    5,
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

	t.Run("test on suix_getAllCoins", func(t *testing.T) {
		rsp, err := cli.SuiXGetAllCoins(ctx, models.SuiXGetAllCoinsRequest{
			Owner: "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5",
			Limit: 5,
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

	t.Run("test on suix_getCoinMetadata", func(t *testing.T) {
		rsp, err := cli.SuiXGetCoinMetadata(ctx, models.SuiXGetCoinMetadataRequest{
			CoinType: "0xf7a0b8cc24808220226301e102dae27464ea46e0d74bb968828318b9e3a921fa::busd::BUSD",
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

	t.Run("test on suix_getCoinMetadata", func(t *testing.T) {
		rsp, err := cli.SuiXGetTotalSupply(ctx, models.SuiXGetTotalSupplyRequest{
			CoinType: "0x3d5ef021274cdc194009ea17e8018ac00ff63843d34dd0fdb57e2696a2020293::polat::POLAT",
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})
}

func TestOnReadTransactionFromSui(t *testing.T) {

	t.Run("test on sui_getTotalTransactionBlocks", func(t *testing.T) {
		rsp, err := cli.SuiGetTotalTransactionBlocks(ctx)

		if err != nil {
			fmt.Println()
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

	t.Run("test on sui_getTransactionBlock", func(t *testing.T) {
		rsp, err := cli.SuiGetTransactionBlock(ctx, models.SuiGetTransactionBlockRequest{
			Digest: "CeVpDXKKU3Gs89efej9pKiYYQyTzifE2BDxWwquUaUht",
			Options: models.SuiTransactionBlockOptions{
				ShowInput:          true,
				ShowRawInput:       true,
				ShowEffects:        true,
				ShowEvents:         true,
			},
		})

		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

	t.Run("test on sui_multiGetTransactionBlocks", func(t *testing.T) {
		rsp, err := cli.SuiMultiGetTransactionBlocks(ctx, models.SuiMultiGetTransactionBlocksRequest{
			Digests: []string{"CeVpDXKKU3Gs89efej9pKiYYQyTzifE2BDxWwquUaUht", "C2zZu2dpX2sLQy2234yt6ecRiNTVgQTXeQpgw9GhxGgo"},
			Options: models.SuiTransactionBlockOptions{
				ShowInput:          true,
				ShowRawInput:       true,
				ShowEffects:        true,
			},
		})

		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}

		for _, transactionBlock := range rsp {
			utils.PrettyPrint(*transactionBlock)
		}
	})

	t.Run("test on suix_queryTransactionBlocks", func(t *testing.T) {
		rsp, err := cli.SuiXQueryTransactionBlocks(ctx, models.SuiXQueryTransactionBlocksRequest{
			SuiTransactionBlockResponseQuery: models.SuiTransactionBlockResponseQuery{
				TransactionFilter: models.TransactionFilter{
					"FromAddress": "0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e",
				},
				Options: models.SuiTransactionBlockOptions{
					ShowInput:          true,
					ShowRawInput:       true,
					ShowEffects:        true,
				},
			},
			Limit:           5,
			DescendingOrder: false,
		})
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}

		for _, transactionBlock := range rsp.Data {
			utils.PrettyPrint(transactionBlock)
		}
	})
}

func TestOnReadObjectFromSui(t *testing.T) {

	t.Run("test on suix_getOwnedObjects", func(t *testing.T) {
		suiObjectResponseQuery := models.SuiObjectResponseQuery{
			Options: models.SuiObjectDataOptions{
				ShowType:                true,
				ShowContent:             true,
				ShowBcs:                 true,
				ShowOwner:               true,
				ShowPreviousTransaction: true,
				ShowStorageRebate:       true,
				ShowDisplay:             true,
			},
		}

		rsp, err := cli.SuiXGetOwnedObjects(ctx, models.SuiXGetOwnedObjectsRequest{
			Address: "0x00432a757bb5d93b2c88936e63692b24baed581090fc0d9a8ffb7e8b633d7ad8",
			Query:   suiObjectResponseQuery,
			Limit:   5,
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

	t.Run("test on sui_multiGetObjects", func(t *testing.T) {
		rsp, err := cli.SuiMultiGetObjects(ctx, models.SuiMultiGetObjectsRequest{
			ObjectIds: []string{"0xeeb964d1e640219c8ddb791cc8548f3242a3392b143ff47484a3753291cad898"},
			Options: models.SuiObjectDataOptions{
				ShowContent:             true,
				ShowDisplay:             true,
				ShowType:                true,
				ShowBcs:                 true,
				ShowOwner:               true,
				ShowPreviousTransaction: true,
				ShowStorageRebate:       true,
			},
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)

	})

	t.Run("test on suix_getDynamicFields", func(t *testing.T) {
		rsp, err := cli.SuiXGetDynamicField(ctx, models.SuiXGetDynamicFieldRequest{
			ObjectId: "0xeeb964d1e640219c8ddb791cc8548f3242a3392b143ff47484a3753291cad898",
			Limit:    5,
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)

	})

}

func TestOnReadEventFromSui(t *testing.T) {

	// t.Run("test on sui_getEvents", func(t *testing.T) {
	// 	rsp, err := cli.SuiGetEvents(ctx, models.SuiGetEventsRequest{
	// 		Digest: "HATq5p7MNynkBL5bLsdVqL3K38PxWHbqs7vndGiz5qrA",
	// 	})
	//
	// 	if err != nil {
	// 		t.Error(err.Error())
	// 		t.FailNow()
	// 	}
	//
	// 	for _, event := range rsp {
	// 		utils.PrettyPrint(*event)
	// 	}
	// })

	t.Run("test on suix_queryEvents", func(t *testing.T) {
		rsp, err := cli.SuiXQueryEvents(ctx, models.SuiXQueryEventsRequest{
			SuiEventFilter: models.EventFilterByMoveEventType{
				MoveEventType: "0x3::validator::StakingRequestEvent",
			},
			Limit:           5,
			DescendingOrder: false,
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})
}

func TestOnWriteTransactionToSui(t *testing.T) {

	t.Run("test on sui_executeTransactionBlock", func(t *testing.T) {
		rsp, err := cli.SuiExecuteTransactionBlock(ctx, models.SuiExecuteTransactionBlockRequest{
			TxBytes:   "AAACAQEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQEAAAAAAAAAAQEAjgDW4hJZlqvw654RGR3SdndKkdjoC0pzXQLxja/NUahLowQAAAAAACBEQGwClI9RQX68dzbN7PN29/Pw/Sc1hbtZwNAny7wZ+wEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMKc3VpX3N5c3RlbRZyZXF1ZXN0X3dpdGhkcmF3X3N0YWtlAAIBAAABAQC3+Y0yfxn2dDR+HkBkFAglMULW5+UJOnyW7ajN/X2btQEqzrI5x8BMQ6LjmCSgAykfjisdYCcyTfW79nyzDB/PvtZBpwAAAAAAIAm+IREDziwoZLm7lc4ZKegZ2J5viEgoss9zgrFkHLh6t/mNMn8Z9nQ0fh5AZBQIJTFC1uflCTp8lu2ozf19m7XoAwAAAAAAAFDhjyoAAAAAAA==",
			Signature: []string{"ALISOTYXKlmBvQ1Uc4UrlWieczU9cGwkyT0Mg65Y2r6VvriElBGB63JDjqg1714Z8B0m84g3S4yrJIIws+leugKOjKY5Wf3dV/la/GVL26whJPWy7WsrWUH2wtmlmcgN6w=="},
			Options: models.SuiTransactionBlockOptions{
				ShowInput:          true,
				ShowRawInput:       true,
				ShowEffects:        true,
				ShowEvents:         true,
				ShowObjectChanges:  true,
				ShowBalanceChanges: true,
			},
			RequestType: "WaitForLocalExecution",
		})

		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		utils.PrettyPrint(rsp)
	})

}
