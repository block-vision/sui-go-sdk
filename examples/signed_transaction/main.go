package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/sui"
)

// send signed transactions
func main() {
	cli := sui.NewSuiClient("https://fullnode.devnet.sui.io:443")

	keystoreCli, err := sui.SetAccountKeyStore("../../config/sui.keystore.fortest")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(keystoreCli.Keys())
	fmt.Println(keystoreCli.GetKey("your-address"))

	resp, err := cli.MoveCallAndExecuteTransaction(context.Background(), models.MoveCallAndExecuteTransactionRequest{
		MoveCallRequest: models.MoveCallRequest{
			Signer:          "0x4d6f1a54e805038f44ecd3112927af147e9b9ecb",
			PackageObjectId: "0x0000000000000000000000000000000000000002",
			Module:          "devnet_nft",
			Function:        "mint",
			TypeArguments:   []interface{}{},
			Arguments:       []interface{}{"blockvision", "blockvision", "testurl"},
			Gas:             "0x14802aeff2f444c888303f8fbba6d4e8451c38f8",
			GasBudget:       1000,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)

	resp2, err := cli.MoveCall(context.Background(), models.MoveCallRequest{
		Signer:          "0x4d6f1a54e805038f44ecd3112927af147e9b9ecb",
		PackageObjectId: "0x0000000000000000000000000000000000000002",
		Module:          "devnet_nft",
		Function:        "mint",
		TypeArguments:   []interface{}{},
		Arguments:       []interface{}{"blockvision", "blockvision", "testurl"},
		Gas:             "0x14802aeff2f444c888303f8fbba6d4e8451c38f8",
		GasBudget:       1000,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("%+v", resp2))

}
