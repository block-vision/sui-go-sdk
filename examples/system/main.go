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
	SuiGetCheckpoint()
	SuiGetCheckpoints()
	SuiGetLatestCheckpointSequenceNumber()
	SuiXGetReferenceGasPrice()
}

func SuiGetCheckpoint() {
	rsp, err := cli.SuiGetCheckpoint(ctx, models.SuiGetCheckpointRequest{
		CheckpointID: "1628214",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiGetCheckpoints() {
	rsp, err := cli.SuiGetCheckpoints(ctx, models.SuiGetCheckpointsRequest{
		Cursor:          "13200",
		Limit:           5,
		DescendingOrder: true,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiGetLatestCheckpointSequenceNumber() {
	rsp, err := cli.SuiGetLatestCheckpointSequenceNumber(ctx)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetReferenceGasPrice() {
	rsp, err := cli.SuiXGetReferenceGasPrice(ctx)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}
