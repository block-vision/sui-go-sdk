package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {
	cli := sui.NewSuiClient("https://rpc-testnet.suiscan.xyz:443")

	resp, err := cli.QueryEvents(context.Background(), models.QueryEventsRequest{
		EventQuery: models.EventQuery{
			// MoveModule: models.MoveModule{
			// 	Package: "0x0871e7323094a4dd1cfcfa1602f7bb072edb399a740609a0d414ac2dbd4dc8d3",
			// 	Module:  "create_nft",
			// },
			MoveEventType: "0x89f04e53299d54e328f49d6ba6b61f522abb6173f633a7040e8e5e37e260933::launchpad_module::AddNftToProject",
		},
		Cursor:    nil,
		Limit:     1,
		DescOrder: true,
	})
	if err != nil {
		fmt.Println(err)
	}

	for _, event := range resp.Data {
		fmt.Printf("%+v\n", event.ID)
	}
}
