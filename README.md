# Sui-Go-SDK

<p align="center">
    <a href="https://github.com/block-vision/sui-go-sdk/blob/main/.github/workflows/ci.yml"><img src="https://github.com/block-vision/sui-go-sdk/actions/workflows/ci.yml/badge.svg"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/License-Apache_2.0-red.svg"></a>
    <a href="https://goreportcard.com/report/github.com/block-vision/sui-go-sdk"><img src="https://goreportcard.com/badge/github.com/securego/gosec"></a>
    <a href="https://pkg.go.dev/github.com/block-vision/sui-go-sdk"> <img src="https://pkg.go.dev/badge/github.com/block-vision/sui-go-sdk.svg"></a>
    <a href="https://discord.gg/Re6prK86Tr"><img src="https://img.shields.io/badge/chat-on%20discord-7289da.svg?sanitize=true"></a>
</p>

## Overview
The Sui-Go-SDK provided by BlockVision aims to offer access to all resources in the BlockVision API and also offers some additional features that make the integration easier.
Sui-Go-SDK is designated for Layer 1 BlockChain [Sui](https://github.com/MystenLabs/sui) in Go programming language.

### Features
+ Support the mainstream methods in the Object, Coin, Event, Read Transaction Blocks, System Data, and Write Transaction Blocks modules.
+ Customized request method `SuiCall`.
+ Unsigned methods can be executed without loading your keystore file.


## Quick Start

### Install 
```shell
go get github.com/block-vision/sui-go-sdk
```

### Go Version
| Golang Version |
|----------------|
| \>= 1.20       | 

## Examples

### Get started
```go
package main

import (
	"context"
	"fmt"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
)

func main() {
	// configure your endpoint here or use blockVision's free Sui RPC endpoint
	cli := sui.NewSuiClient("https://sui-testnet-endpoint.blockvision.org")

	ctx := context.Background()

	rsp, err := cli.SuiXGetAllBalance(ctx, models.SuiXGetAllBalanceRequest{
		Owner: "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(rsp)

	//If you want to request for original json response data, you can use SuiCall().
	callRsp, err := cli.SuiCall(ctx, "suix_getAllBalances", "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(callRsp)
}
```

## Contribution  
+ We welcome your suggestions, comments (including criticisms), comments and contributions.   
+ Please follow the PR/issue template provided to ensure that your contributions are clear and easy to understand.  
+ Thank you to all the people who participate in building better infrastructure! 

## Resources
+ [SDK Examples](https://github.com/block-vision/sui-go-sdk/tree/main/examples)
+ [Sui](https://github.com/MystenLabs/sui)


## License 
[Apache 2.0 license](LICENSE)




