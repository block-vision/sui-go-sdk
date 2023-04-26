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
	SuiGetTotalTransactionBlocks()
	SuiGetTransactionBlock()
	SuiMultiGetTransactionBlocks()
	SuiXQueryTransactionBlocks()
}

func SuiGetTotalTransactionBlocks() {
	rsp, err := cli.SuiGetTotalTransactionBlocks(ctx)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiGetTransactionBlock() {
	rsp, err := cli.SuiGetTransactionBlock(ctx, models.SuiGetTransactionBlockRequest{
		Digest: "CeVpDXKKU3Gs89efej9pKiYYQyTzifE2BDxWwquUaUht",
		Options: models.SuiTransactionBlockOptions{
			ShowInput:          true,
			ShowRawInput:       true,
			ShowEffects:        true,
			ShowEvents:         true,
			ShowBalanceChanges: true,
			ShowObjectChanges:  true,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiMultiGetTransactionBlocks() {
	rsp, err := cli.SuiMultiGetTransactionBlocks(ctx, models.SuiMultiGetTransactionBlocksRequest{
		Digests: []string{"CeVpDXKKU3Gs89efej9pKiYYQyTzifE2BDxWwquUaUht", "C2zZu2dpX2sLQy2234yt6ecRiNTVgQTXeQpgw9GhxGgo"},
		Options: models.SuiTransactionBlockOptions{
			ShowInput:          true,
			ShowRawInput:       true,
			ShowEffects:        true,
			ShowEvents:         true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, transactionBlock := range rsp {
		utils.PrettyPrint(*transactionBlock)
	}
}

func SuiXQueryTransactionBlocks() {
	rsp, err := cli.SuiXQueryTransactionBlocks(ctx, models.SuiXQueryTransactionBlocksRequest{
		SuiTransactionBlockResponseQuery: models.SuiTransactionBlockResponseQuery{
			TransactionFilter: models.TransactionFilter{
				"FromAddress": "0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e",
			},
			Options: models.SuiTransactionBlockOptions{
				ShowInput:          true,
				ShowRawInput:       true,
				ShowEffects:        true,
				ShowEvents:         true,
				ShowObjectChanges:  true,
				ShowBalanceChanges: true,
			},
		},
		Limit:           5,
		DescendingOrder: false,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, transactionBlock := range rsp.Data {
		utils.PrettyPrint(transactionBlock)
	}
}
