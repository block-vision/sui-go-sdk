package main

import (
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/sui"
)

func main() {
	RequestDevNetSuiFromFaucet()
	RequestTestNetSuiFromFaucet()
}

func RequestDevNetSuiFromFaucet() {
	faucetHost, err := sui.GetFaucetHost(constant.SuiDevnet)
	if err != nil {
		fmt.Println("GetFaucetHost err:", err)
		return
	}

	fmt.Println("faucetHost:", faucetHost)

	recipient := "0xaf9f4d20c205f26051a7e1758601c4c47a9f99df3f9823f70926c17c80882d36"

	header := map[string]string{}
	err = sui.RequestSuiFromFaucet(faucetHost, recipient, header)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// the successful transaction block url: https://suiexplorer.com/txblock/91moaxbXsQnJYScLP2LpbMXV43ZfngS2xnRgj1CT7jLQ?network=devnet
	fmt.Println("Request DevNet Sui From Faucet success")
}

func RequestTestNetSuiFromFaucet() {
	faucetHost, err := sui.GetFaucetHost(constant.SuiTestnet)
	if err != nil {
		fmt.Println("GetFaucetHost err:", err)
		return
	}

	fmt.Println("faucetHost:", faucetHost)

	recipient := "0xaf9f4d20c205f26051a7e1758601c4c47a9f99df3f9823f70926c17c80882d36"

	header := map[string]string{
		"Authority":          "faucet.testnet.sui.io",
		"Accept":             "*/*",
		"Accept-Language":    "zh-CN,zh;q=0.9",
		"Content-Type":       "application/json",
		"Sec-Ch-Ua":          "Not.A/Brand;v=8, Chromium;v=114, Google Chrome;v=114",
		"Sec-Ch-Ua-Mobile":   "?0",
		"Sec-Ch-Ua-Platform": "macOS",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "none",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	}

	err = sui.RequestSuiFromFaucet(faucetHost, recipient, header)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Request TestNet Sui From Faucet success")
}
