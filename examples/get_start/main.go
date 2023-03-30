package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {
	cli := sui.NewSuiClient("https://fullnode.devnet.sui.io:443")
	resp, err := cli.SuiCall(context.Background(), "sui_getTransactionAuthSigners", "Ar9FigfQcR52tAGDPwt2DtKyp4oNDgRCc85yzeCuFU1L")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp result:", resp)

	resp2, err := cli.GetTransactionAuthSigners(context.Background(), models.GetTransactionAuthSignersRequest{
		Digest: "Ar9FigfQcR52tAGDPwt2DtKyp4oNDgRCc85yzeCuFU1L",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v\n", resp2)

	keystoreCli, err := sui.SetAccountKeyStore("../sui.keystore.fortest")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(keystoreCli.Keys())
	fmt.Println(keystoreCli.GetKey("your-address"))
}
