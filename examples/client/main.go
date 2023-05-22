package main

import (
	"context"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
	"net/http"
	"net/url"
	"time"
)

func main() {
	runClientDefault()
	runClientWithProxy()
}

func runClientDefault() {
	var ctx = context.Background()

	// create a sui client with default http client
	var clientWithProxy = sui.NewSuiClient(constant.BvMainnetEndpoint)

	// use default client to get object data by object id
	objectData, err := clientWithProxy.SuiGetObject(ctx, models.SuiGetObjectRequest{
		ObjectId: "0x0000000000000000000000000000000000000000000000000000000000000006",
		Options: models.SuiObjectDataOptions{
			ShowContent: true,
			ShowDisplay: true,
			ShowType:    true,
			ShowBcs:     true,
		},
	})
	if err != nil {
		return
	}

	utils.PrettyPrint(objectData)
}

func runClientWithProxy() {
	var ctx = context.Background()

	// create a proxy url
	proxyUrl, err := url.Parse("http://127.0.0.1:7893")
	if err != nil {
		panic(err)
	}

	// create a sui client with custom http client with proxy
	var clientWithProxy = sui.NewSuiClientWithCustomClient(constant.BvMainnetEndpoint, &http.Client{
		Timeout:   time.Second * 10,
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
	})

	// use proxy client to get object data by object id
	objectData, err := clientWithProxy.SuiGetObject(ctx, models.SuiGetObjectRequest{
		ObjectId: "0x0000000000000000000000000000000000000000000000000000000000000006",
		Options: models.SuiObjectDataOptions{
			ShowContent: true,
			ShowDisplay: true,
			ShowType:    true,
			ShowBcs:     true,
		},
	})
	if err != nil {
		return
	}

	utils.PrettyPrint(objectData)
}
