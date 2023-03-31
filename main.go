package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {
	cli := sui.NewSuiClient("https://fullnode.devnet.sui.io:443")

	res, err := cli.GetOwnedObjects(context.Background(), models.GetOwnedObjectsRequest{
		Address: "0x78dc765e2cb0d0b6d4f7b172213b5a554880ef237fc280d81e410e3af737c62f",
	})
	if err != nil {
		fmt.Println(err)
	}

	for _, r := range res.Data {
		object, err := cli.GetObject(context.Background(), models.GetObjectRequest{
			ObjectID: r.Data.ObjectID,
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", object)
	}
}
