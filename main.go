package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {
	cli := sui.NewSuiClient("https://fullnode.testnet.sui.io:443")

	res, err := cli.GetCoins(context.Background(), models.GetCoinsRequeset{
		Owner: "0x68af121be266605d657285f08c51378c6d46ff14bac05029bf54b0976d4bf016",
		// CoinType: "0x2::sui::SUI",
		// Cursor:   "",
		Limit: 100,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
