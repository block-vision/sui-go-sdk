package main

import (
	"context"
	"fmt"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
)

var ctx = context.Background()
var cli = sui.NewSuiClient(constant.BvTestnetEndpoint)

func main() {
	SuiXGetBalance()
	SuiXGetAllBalance()
	SuiXGetCoins()
	SuiXGetAllCoins()
	SuiXGetCoinMetadata()
	SuiXGetTotalSupply()
}

func SuiXGetBalance() {
	rsp, err := cli.SuiXGetBalance(ctx, models.SuiXGetBalanceRequest{
		Owner:    "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5",
		CoinType: "0x2::sui::SUI",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetAllBalance() {
	rsp, err := cli.SuiXGetAllBalance(ctx, models.SuiXGetAllBalanceRequest{
		Owner: "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetCoins() {
	rsp, err := cli.SuiXGetCoins(ctx, models.SuiXGetCoinsRequest{
		Owner:    "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5",
		CoinType: "0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI",
		Limit:    5,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetAllCoins() {
	rsp, err := cli.SuiXGetAllCoins(ctx, models.SuiXGetAllCoinsRequest{
		Owner: "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5",
		Limit: 5,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetCoinMetadata() {
	rsp, err := cli.SuiXGetCoinMetadata(ctx, models.SuiXGetCoinMetadataRequest{
		CoinType: "0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetTotalSupply() {
	rsp, err := cli.SuiXGetTotalSupply(ctx, models.SuiXGetTotalSupplyRequest{
		CoinType: "0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}
