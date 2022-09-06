package main

import (
	"context"
	"fmt"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
)

func main() {
	cli := sui.NewSuiClient("https://gateway.devnet.sui.io:443")
	resp, err := cli.GetRecentTransactions(context.Background(), models.GetRecentTransactionRequest{
		Count: 5,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)

	//If you want to request for original json response, you can use SuiCall().
	rsp, err := cli.SuiCall(context.Background(), "sui_getRecentTransactions", 5)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(rsp)

	keystoreCli, err := sui.SetAccountKeyStore("../sui.keystore.fortest")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(keystoreCli.Keys())
	fmt.Println(keystoreCli.GetKey("your-address"))
}
