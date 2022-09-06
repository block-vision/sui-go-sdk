package sui

import (
	"context"
	"fmt"
	"github.com/block-vision/sui-go-sdk/models"
	"gotest.tools/assert"
	"testing"
)

var cli = NewSuiClient("https://gateway.devnet.sui.io:443")
var ctx = context.Background()

func TestOnGetTransactionFromSui(t *testing.T) {
	t.Run("test on sui_call", func(t *testing.T) {
		resp, err := cli.SuiCall(ctx, "sui_getTransactionsByInputObject", []interface{}{
			"0xaea08f870ce66d89bc4e1cd1d27445c575195c2c",
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getRecentTransaction", func(t *testing.T) {
		resp, err := cli.GetRecentTransactions(ctx, models.GetRecentTransactionRequest{
			Count: 5,
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getTotalTransactionNumber", func(t *testing.T) {
		resp, err := cli.GetTotalTransactionNumber(ctx, models.GetTotalTransactionNumberRequest{})
		if err != nil {
			fmt.Println()
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getTransaction", func(t *testing.T) {
		resp, err := cli.GetTransaction(ctx, models.GetTransactionRequest{Digest: "dhdxO9moFHP+JI+UnT/RNylhzOXtDT0jjITApWEDVIg="})
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}
		fmt.Printf("%+v", resp)
	})

	t.Run("test on sui_getTransactionByInputObject", func(t *testing.T) {
		resp, err := cli.GetTransactionsByInputObject(ctx, models.GetTransactionsByInputObjectRequest{
			ObjectID: "0x39377e86154771712b3b8377dafeb0d4e5a0bd96",
		})
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getTransactionsByMoveFunction", func(t *testing.T) {
		resp, err := cli.GetTransactionsByMoveFunction(ctx, models.GetTransactionsByMoveFunctionRequest{
			Package:  "0x0000000000000000000000000000000000000002",
			Module:   "devnet_nft",
			Function: "mint",
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getTransactionsByMutatedObject", func(t *testing.T) {
		resp, err := cli.GetTransactionsByMutatedObject(ctx, models.GetTransactionsByMutatedObjectRequest{
			ObjectID: "0x9fc486e101f0e9e2703ac8666d06aecc4ddf7e79",
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getTransactionsFromAddress", func(t *testing.T) {
		resp, err := cli.GetTransactionsFromAddress(ctx, models.GetTransactionsFromAddressRequest{
			Addr: "0x6ce471116a8b96f81c47414bae8375c42c23e0fc",
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getTransactionsInRange", func(t *testing.T) {
		resp, err := cli.GetTransactionsInRange(ctx, models.GetTransactionsInRangeRequest{
			Start: 5,
			End:   8,
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getTransactionsToAddress", func(t *testing.T) {
		resp, err := cli.GetTransactionsToAddress(ctx, models.GetTransactionsToAddressRequest{Addr: "0xb6eb9669d0c206d28ed50358dfd45c7f5f8d2669"})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})
}

func TestOnReadObjectFromSui(t *testing.T) {
	t.Run("test on sui_getObject", func(t *testing.T) {
		resp, err := cli.GetObject(ctx, models.GetObjectRequest{ObjectID: "0x869afb00e643e1f09af0e12d4732ccd71dfca6e2"})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)

		resp2, err := cli.SuiCall(ctx, "sui_getObject", []interface{}{
			"0x869afb00e643e1f09af0e12d4732ccd71dfca6e2",
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp2)
	})

	t.Run("test on sui_getObjectOwnedByAddress", func(t *testing.T) {
		resp, err := cli.GetObjectsOwnedByAddress(ctx, models.GetObjectsOwnedByAddressRequest{
			Address: "0x425c9e7182bff2b2aea5d31ffc1043e73c9a999d",
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getObjectsOwnedByObject", func(t *testing.T) {
		resp, err := cli.GetObjectsOwnedByObject(ctx, models.GetObjectsOwnedByObjectRequest{
			ObjectID: "0x67b855b694a6a69ae248876e885bb38174b50cf3",
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getRawObject", func(t *testing.T) {
		resp, err := cli.GetRawObject(ctx, models.GetRawObjectRequest{
			ObjectID: "0x67b855b694a6a69ae248876e885bb38174b50cf3",
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp.Details.Data.BcsBytes)
	})

}

func TestOnReadEventFromSui(t *testing.T) {
	t.Run("test on sui_getEventsByEventType", func(t *testing.T) {
		resp, err := cli.GetEventsByEventType(ctx, models.GetEventsByEventTypeRequest{
			EventType: "0x2::devnet_nft::MintNFTEvent",
			Count:     10,
			StartTime: 0,
			EndTime:   0,
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getEventsByModule", func(t *testing.T) {
		resp, err := cli.GetEventsByModule(ctx, models.GetEventsByModuleRequest{
			Package:   "",
			Module:    "",
			Count:     0,
			StartTime: 0,
			EndTime:   0,
		})

		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getEventsByObject", func(t *testing.T) {
		resp, err := cli.GetEventsByObject(ctx, models.GetEventsByObjectRequest{
			Object:    "",
			Count:     0,
			StartTime: 0,
			EndTime:   0,
		})

		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getEventsByOwner", func(t *testing.T) {
		resp, err := cli.GetEventsByOwner(ctx, models.GetEventsByOwnerRequest{
			Owner:     "",
			Count:     0,
			StartTime: 0,
			EndTime:   0,
		})

		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getEventsBySender", func(t *testing.T) {
		resp, err := cli.GetEventsBySender(ctx, models.GetEventsBySenderRequest{
			Sender:    "",
			Count:     0,
			StartTime: 0,
			EndTime:   0,
		})

		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_getEventsByTransaction", func(t *testing.T) {
		resp, err := cli.GetEventsByTransaction(ctx, models.GetEventsByTransactionRequest{
			Digest: "",
		})

		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

}

//func TestOnReadMoveDataFromSui(t *testing.T) {
//	t.Run("test on sui_getMoveFunctionArgTypes", func(t *testing.T) {
//		resp, err := cli.GetMoveFunctionArgTypes(ctx, models.GetMoveFunctionArgTypesRequest{
//			Package:  "0x2",
//			Module:   "devnet",
//			Function: "",
//		})
//
//		if err != nil {
//			fmt.Println(err.Error())
//			t.FailNow()
//		}
//		fmt.Println(resp)
//	})
//
//	t.Run("test on sui_getNormalizedMoveModulesByPackage", func(t *testing.T) {
//		resp, err := cli.GetNormalizedMoveModulesByPackage(ctx, models.GetNormalizedMoveModulesByPackageRequest{
//			Package: "",
//		})
//		if err != nil {
//			fmt.Println(err.Error())
//			t.FailNow()
//		}
//		fmt.Println(resp)
//	})
//
//	t.Run("test on sui_getNormalizedMoveModule", func(t *testing.T) {
//		resp, err := cli.GetNormalizedMoveModule(ctx, models.GetNormalizedMoveModuleRequest{
//			Package:    "",
//			ModuleName: "",
//		})
//		if err != nil {
//			fmt.Println(err.Error())
//			t.FailNow()
//		}
//		fmt.Println(resp)
//	})
//
//	t.Run("test on sui_getNormalizedMoveStruct", func(t *testing.T) {
//		resp, err := cli.GetNormalizedMoveStruct(ctx, models.GetNormalizedMoveStructRequest{
//			Package:    "",
//			ModuleName: "",
//			StructName: "",
//		})
//		if err != nil {
//			fmt.Println(err.Error())
//			t.FailNow()
//		}
//		fmt.Println(resp)
//	})
//
//	t.Run("test on sui_getNormalizedMoveFunction", func(t *testing.T) {
//		resp, err := cli.GetNormalizedMoveFunction(ctx, models.GetNormalizedMoveFunctionRequest{
//			Package:      "",
//			ModuleName:   "",
//			FunctionName: "",
//		})
//		if err != nil {
//			fmt.Println(err.Error())
//			t.FailNow()
//		}
//		fmt.Println(resp)
//	})
//
//}

func TestOnWriteTransactionToSui(t *testing.T) {
	t.Run("test on sui_moveCall", func(t *testing.T) {
		resp, err := cli.MoveCall(ctx, models.MoveCallRequest{
			Signer:          "0x0f1584ebdf54c91b8572793b2a79e085514ea6c7",
			PackageObjectId: "0x0000000000000000000000000000000000000002",
			Module:          "devnet_nft",
			Function:        "mint",
			TypeArguments:   []interface{}{},
			Arguments:       []interface{}{"blockvision", "blockvision", "testurl"},
			Gas:             "0xd88db3b517b218503ed90df77f25dec257918903",
			GasBudget:       1000,
		})

		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp.TxBytes)
	})

	t.Run("test on sui_mergeCoins", func(t *testing.T) {
		resp, err := cli.MergeCoins(ctx, models.MergeCoinsRequest{
			Signer:      "0x0f1584ebdf54c91b8572793b2a79e085514ea6c7",
			PrimaryCoin: "0x0407064f11682317c1e0220ce1f5b23246c91d23",
			CoinToMerge: "0x35a8b00aa176337db1dfa9d5681ce18e45183058",
			Gas:         "0xd88db3b517b218503ed90df77f25dec257918903",
			GasBudget:   1000,
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_publish", func(t *testing.T) {
		resp, err := cli.Publish(ctx, models.PublishRequest{
			Sender: "0x0f1584ebdf54c91b8572793b2a79e085514ea6c7",
			CompiledModules: []string{
				"oRzrCwUAAAALAQAOAg4kAzJZBIsBFgWhAZoBB7sC5wEIogQoBsoECgrUBBoM7gSjAQ2RBgQAAAEBAQIBAwEEAQUBBgAHCAAACAgAAAkIAAIKDAEAAQQLAgAGDAIAAxIEAAEUBAEAAQANAAEAAA4CAQAADwMBAAAQBAEAAhUHCAEAAhYJCgEAARcLDAEAARgNCAEAAxkEDgAGGg8QAAUFEQEBCAEVEwgBAAIbFBUBAAMcDgEAAR0BDAEABR4ZAQEIBAYFBgYGBwYKAwsGDAYKFgoXDgYPGAMHCAEHCwMBCAQHCAUAAwYIAgcIAQcIBQEIAAEHCAUCBwsHAQgECwcBCAQBCAQBBgsDAQkAAQMBBwsDAQkAAQcLBwEJAAIHCwcBCQADAQsHAQkAAgcLBwEJAAsHAQkAAQgGAQYIBQEFAgkABQIDCwMBCAQBBgsHAQkAAwcLBwEJAAMHCAUBCwMBCQABCwMBCAQBCAIBCAEBCQAGZG9udXRzB2JhbGFuY2UEY29pbgZvYmplY3QDc3VpCHRyYW5zZmVyCnR4X2NvbnRleHQFRG9udXQJRG9udXRTaG9wDFNob3BPd25lckNhcARDb2luA1NVSQlUeENvbnRleHQJYnV5X2RvbnV0D2NvbGxlY3RfcHJvZml0cwllYXRfZG9udXQEaW5pdAJpZANVSUQFcHJpY2UHQmFsYW5jZQV2YWx1ZQtiYWxhbmNlX211dAVzcGxpdARqb2luA25ldwZzZW5kZXIEdGFrZQZkZWxldGUEemVybwxzaGFyZV9vYmplY3QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAwgAAAAAAAAAAAACAREIBgECAxEIBhMDAQsHAQgEAgIBEQgGAAEEAAUmCgEuOAAKABAAFCYDEAsAAQsBAQsCAQcAJwsBOAEMAwsDCgAQABQ4AgwECwAPAQsEOAMBCgIRCBIACwIuEQk4BAIBAQQAEhAKARABOAUMAwsBDwELAwoCOAYMBAsECwIuEQk4BwICAQQADgYLABMADAELARENAgMAAAABDgoAEQgSAgoALhEJOAgLABEIBugDAAAAAAAAOAkSATgKAgEBAQIA",
			},
			Gas:       "0xd88db3b517b218503ed90df77f25dec257918903",
			GasBudget: 1000,
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_splitCoin", func(t *testing.T) {
		resp, err := cli.SplitCoin(ctx, models.SplitCoinRequest{
			Signer:       "0x0f1584ebdf54c91b8572793b2a79e085514ea6c7",
			CoinObjectId: "0x0407064f11682317c1e0220ce1f5b23246c91d23",
			SplitAmounts: []uint64{1000, 1000},
			Gas:          "0xd88db3b517b218503ed90df77f25dec257918903",
			GasBudget:    1000,
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_splitCoinEqual", func(t *testing.T) {
		resp, err := cli.SplitCoinEqual(ctx, models.SplitCoinEqualRequest{
			Signer:       "0x0f1584ebdf54c91b8572793b2a79e085514ea6c7",
			CoinObjectId: "0x0407064f11682317c1e0220ce1f5b23246c91d23",
			SplitCount:   2,
			Gas:          "0xd88db3b517b218503ed90df77f25dec257918903",
			GasBudget:    1000,
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_transferObject", func(t *testing.T) {
		resp, err := cli.TransferObject(ctx, models.TransferObjectRequest{
			Signer:    "0x0f1584ebdf54c91b8572793b2a79e085514ea6c7",
			ObjectId:  "0xf8da4c3c9d1477e8d2e005ac4d390032b5f81977",
			Gas:       "0xd88db3b517b218503ed90df77f25dec257918903",
			GasBudget: 1000,
			Recipient: "0xb6eb9669d0c206d28ed50358dfd45c7f5f8d2669",
		})

		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_transferSui", func(t *testing.T) {
		resp, err := cli.TransferSui(ctx, models.TransferSuiRequest{
			Signer:      "0x0f1584ebdf54c91b8572793b2a79e085514ea6c7",
			SuiObjectId: "0xf8da4c3c9d1477e8d2e005ac4d390032b5f81977",
			GasBudget:   1000,
			Recipient:   "0xb6eb9669d0c206d28ed50358dfd45c7f5f8d2669",
			Amount:      1,
		})

		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_batchTransaction", func(t *testing.T) {
		resp, err := cli.BatchTransaction(ctx, models.BatchTransactionRequest{
			Signer: "0x0f1584ebdf54c91b8572793b2a79e085514ea6c7",
			SingleTransactionParams: []models.SingleTransactionParams{
				{
					MoveCallRequestParams: &models.MoveCallRequest{
						Signer:          "0x0f1584ebdf54c91b8572793b2a79e085514ea6c7",
						PackageObjectId: "0x0000000000000000000000000000000000000002",
						Module:          "devnet_nft",
						Function:        "mint",
						TypeArguments:   []interface{}{},
						Arguments:       []interface{}{"blockvision", "blockvision", "testurl"},
						Gas:             "0x0407064f11682317c1e0220ce1f5b23246c91d23",
						GasBudget:       1000,
					},
				},
				{
					TransferObjectRequestParams: &models.TransferObjectRequest{
						Signer:    "0x0f1584ebdf54c91b8572793b2a79e085514ea6c7",
						ObjectId:  "0xf8da4c3c9d1477e8d2e005ac4d390032b5f81977",
						Gas:       "0x0407064f11682317c1e0220ce1f5b23246c91d23",
						GasBudget: 1000,
						Recipient: "0xb6eb9669d0c206d28ed50358dfd45c7f5f8d2669",
					},
				},
			},
			Gas:       "0x0407064f11682317c1e0220ce1f5b23246c91d23",
			GasBudget: 1000,
		})

		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})

	t.Run("test on sui_executeTransaction", func(t *testing.T) {
		resp, err := cli.ExecuteTransaction(ctx, models.ExecuteTransactionRequest{
			TxBytes:   "VHJhbnNhY3Rpb25EYXRhOjoAAgAAAAAAAAAAAAAAAAAAAAAAAAACAQAAAAAAAAAgf4wCMzbSQGAtJy5c2FShsm5eDefCLIODnSU2sC07IXMKZGV2bmV0X25mdARtaW50AAMADAtibG9ja3Zpc2lvbgAMC2Jsb2NrdmlzaW9uAAgHdGVzdHVybAgHEPIa1agDuTeHZVW0qooyQdm42I2ztReyGFA+2Q33fyXewleRiQMDAAAAAAAAACDURxg9juaOxAd9LLAJYPApQ2HE2zvHWehpj/PrXIge0QEAAAAAAAAA6AMAAAAAAAA=",
			SigScheme: "ED25519",
			Signature: "d3fLEw4f6JB+4NbRmr+Wh5EeyL5xGLnTYLt6a5Wtsh5Kf7ADx7V1WyWaqDehnulrrR4WS+ybSOkGRiRv7mfnCw==",
			PubKey:    "kslpD84v6KJ8G3J+iiHH+s7bDoHhucTIUAg+Um5sfAo=",
		})

		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Println(resp)
	})
}

func TestOnFeatureAPI(t *testing.T) {
	kps, err := SetAccountKeyStore("../config/sui.keystorefortest")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	actualKeys := kps.Keys()
	assert.Equal(t, "0x72cffd05deb71fa9b30584cb0f512d680cb08eab", actualKeys[0])
	assert.Equal(t, "0x8f3cf7d8ebb187bd655cea775802d0d9c1c5b145", actualKeys[1])
	assert.Equal(t, "0xc697e5fdd38d5f63ebeb14c2b49a864d473849db", actualKeys[2])
	assert.Equal(t, "0xf354bb3497c5879d68b49582d3a8887dbd26e3f0", actualKeys[3])
	assert.Equal(t, "0xfde3698d3e7da3f359e1036078da9cfbfb31f203", actualKeys[4])

	t.Run("test on feature API MoveCallAndExecuteTransaction", func(t *testing.T) {

	})

	t.Run("test on feature API BatchAndExecuteTransaction", func(t *testing.T) {

	})
}
