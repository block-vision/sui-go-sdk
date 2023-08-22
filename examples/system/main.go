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
var cli = sui.NewSuiClient(constant.BvMainnetEndpoint)

func main() {
	SuiGetCheckpoint()
	SuiGetCheckpoints()
	SuiGetLatestCheckpointSequenceNumber()
	SuiXGetReferenceGasPrice()
	SuiXGetCommitteeInfo()
	SuiXGetStakes()
	SuiXGetStakesByIds()
	SuiXGetLatestSuiSystemState()
	SuiGetChainIdentifier()
	SuiXGetValidatorsApy()
	SuiGetProtocolConfig()
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

func SuiXGetCommitteeInfo() {
	rsp, err := cli.SuiXGetCommitteeInfo(ctx, models.SuiXGetCommitteeInfoRequest{
		Epoch: "754",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetStakes() {
	rsp, err := cli.SuiXGetStakes(ctx, models.SuiXGetStakesRequest{
		Owner: "0xe335d84c489084474aac4322fb9ac5173369d27487c404558e99c7c5ec608075",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetStakesByIds() {
	rsp, err := cli.SuiXGetStakesByIds(ctx, models.SuiXGetStakesByIdsRequest{
		StakedSuiIds: []string{"0x9898fae07add84f032eb109ffc548d4afae7c78cb9b0836aed674e7aec55df19"},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetEpochs() {
	rsp, err := cli.SuiXGetEpochs(ctx, models.SuiXGetEpochsRequest{
		Limit:           5,
		DescendingOrder: true,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetCurrentEpoch() {
	rsp, err := cli.SuiXGetCurrentEpoch(ctx)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetLatestSuiSystemState() {
	rsp, err := cli.SuiXGetLatestSuiSystemState(ctx)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiGetChainIdentifier() {
	rsp, err := cli.SuiGetChainIdentifier(ctx)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXGetValidatorsApy() {
	rsp, err := cli.SuiXGetValidatorsApy(ctx)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiGetProtocolConfig() {
	rsp, err := cli.SuiGetProtocolConfig(ctx, models.SuiGetProtocolConfigRequest{})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}
