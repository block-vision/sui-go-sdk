package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {
	cli := sui.NewSuiClient("http://35.185.176.5:9000")

	// res, err := cli.GetAllNFT(context.Background(), "0x6207ebfdef685b73be4308645815738caabcedf80866d21419d9b9982d171838")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _, r := range res {
	// 	fmt.Println("================================")
	// 	fmt.Printf("%+v\n", r.Data)
	// 	// realType := strings.Split(r.Data.Type, "<")[0]
	// 	metadata, err := cli.GetDynamicField(context.Background(), models.GetDynamicFieldRequest{
	// 		ParentObjectID: r.Data.ObjectID,
	// 	})
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Printf("%+v\n", metadata)
	// }

	res, err := cli.GetObject(context.Background(), models.GetObjectRequest{
		ObjectID: "0x41085f8d26829e8ee212decb2a5b8d0c9bd4a97f3f96a9e1514bc60e6060a830",
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", res)
}
