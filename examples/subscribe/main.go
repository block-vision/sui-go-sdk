package main

import (
	"context"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
)

func main() {
	go SubscribeEvent()
	go SubscribeTransaction()
	select {}
}

func SubscribeEvent() {
	var ctx = context.Background()
	var cli = sui.NewSuiWebsocketClient(constant.WssBvTestnetEndpoint)

	receiveMsgCh := make(chan models.SuiEventResponse, 10)
	err := cli.SubscribeEvent(ctx, models.SuiXSubscribeEventsRequest{
		SuiEventFilter: map[string]interface{}{
			"All": []string{},
		},
	}, receiveMsgCh)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case msg := <-receiveMsgCh:
			utils.PrettyPrint(msg)
		case <-ctx.Done():
			return
		}
	}
}

func SubscribeTransaction() {
	var ctx = context.Background()
	var cli = sui.NewSuiWebsocketClient(constant.WssBvTestnetEndpoint)

	receiveMsgCh := make(chan models.SuiEffects, 10)
	err := cli.SubscribeTransaction(ctx, models.SuiXSubscribeTransactionsRequest{
		TransactionFilter: models.TransactionFilterByFromAddress{
			FromAddress: "0x0000000000000000000000000000000000000000000000000000000000000000",
		},
	}, receiveMsgCh)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case msg := <-receiveMsgCh:
			utils.PrettyPrint(msg)
		case <-ctx.Done():
			return
		}
	}
}
