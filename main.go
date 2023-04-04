package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {
	cli := sui.NewSuiClient("http://35.185.176.5:9000")

	res, err := cli.GetAllNFT(context.Background(), "0x6207ebfdef685b73be4308645815738caabcedf80866d21419d9b9982d171838")
	if err != nil {
		fmt.Println(err)
	}

	for _, r := range res {
		// metadata, err := cli.GetDynamicField(context.Background(), models.GetCoinMetadataRequest{
		// 	CoinType: r.Data.Type,
		// })
		// if err != nil {
		// 	fmt.Println(err)
		// }

		fmt.Printf("%+v\n", r)
	}
}
