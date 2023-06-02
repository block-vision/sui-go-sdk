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
	SuiXResolveNameServiceAddress()
	SuiXResolveNameServiceNames()
}

func SuiXResolveNameServiceAddress() {
	rsp, err := cli.SuiXResolveNameServiceAddress(ctx, models.SuiXResolveNameServiceAddressRequest{
		Name: "yoshi.sui",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXResolveNameServiceNames() {
	rsp, err := cli.SuiXResolveNameServiceNames(ctx, models.SuiXResolveNameServiceNamesRequest{
		Address: "0x134c18293d898bb188077cbceb8ababf9c4328d39d8c873285ae7751fb821818",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}
