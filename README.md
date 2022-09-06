# Sui-Go-SDK
[![License Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-red.svg)](LICENSE)
[![GoReport](https://goreportcard.com/badge/github.com/securego/gosec)](https://goreportcard.com/report/github.com/block-vision/sui-go-sdk)

## 

sui-go-sdk is project [Sui]() SDK for Go programming language.  


### Notices
+ You don't need to load your `sui.keystore` file if you just want to send some unsigned transactions.
+ File `sui.keystore` in config folder is test-only. Replace and load your own sui.keystore if your need to sign transactions. 
+ PRs are open to everyone and let's build useful tools for Sui community!


### Features
+ Load your keystore file and sign your messages with specific address.
+ Provide methods `MoveCallAndExecuteTransaction`/`BatchAndExecuteTransaction`.
+ Customized request method `SuiCall`.
+ Unsigned methods can be executed without loading your keystore file.

* [Quick Start](#Quick-Start)
* [Examples](#Examples)

## Quick Start

### Install 
```shell
go get github.com/block-vision/sui-go-sdk
```

### Go Version
| Golang Version |
|----------------|
| \>= 1.18.1     | 

## Examples


```go
package main

import (
	"context"
	"fmt"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
)

func main() {
	// configure your endpoint here
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
```

### Send unsigned transactions



### Send signed transactions


## Contribution  

## Resources
+ [SDK Examples]()
+ [Sui](https://github.com/MystenLabs/sui)





