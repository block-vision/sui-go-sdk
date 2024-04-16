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
	SuiGetMoveFunctionArgTypes()
	SuiGetNormalizedMoveModulesByPackage()
	SuiGetNormalizedMoveModule()
	SuiGetNormalizedMoveStruct()
	SuiGetNormalizedMoveFunction()
}

func SuiGetMoveFunctionArgTypes() {
	rsp, err := cli.SuiGetMoveFunctionArgTypes(ctx, models.GetMoveFunctionArgTypesRequest{
		Package:  "0x7d584c9a27ca4a546e8203b005b0e9ae746c9bec6c8c3c0bc84611bcf4ceab5f",
		Module:   "auction",
		Function: "start_an_auction",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiGetNormalizedMoveModulesByPackage() {
	rsp, err := cli.SuiGetNormalizedMoveModulesByPackage(ctx, models.GetNormalizedMoveModulesByPackageRequest{
		Package: "0x7d584c9a27ca4a546e8203b005b0e9ae746c9bec6c8c3c0bc84611bcf4ceab5f",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiGetNormalizedMoveModule() {
	rsp, err := cli.SuiGetNormalizedMoveModule(ctx, models.GetNormalizedMoveModuleRequest{
		Package:    "0x7d584c9a27ca4a546e8203b005b0e9ae746c9bec6c8c3c0bc84611bcf4ceab5f",
		ModuleName: "auction",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiGetNormalizedMoveStruct() {
	rsp, err := cli.SuiGetNormalizedMoveStruct(ctx, models.GetNormalizedMoveStructRequest{
		Package:    "0x7d584c9a27ca4a546e8203b005b0e9ae746c9bec6c8c3c0bc84611bcf4ceab5f",
		ModuleName: "auction",
		StructName: "BidDetail",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiGetNormalizedMoveFunction() {
	rsp, err := cli.SuiGetNormalizedMoveFunction(ctx, models.GetNormalizedMoveFunctionRequest{
		Package:      "0x7d584c9a27ca4a546e8203b005b0e9ae746c9bec6c8c3c0bc84611bcf4ceab5f",
		ModuleName:   "auction",
		FunctionName: "configure_auction",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}
