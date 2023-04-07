package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {
	cli := sui.NewSuiClient("https://sui-testnet-endpoint.blockvision.org")

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
		ObjectID: "0xcf21b5cf7f7ddd6d301713264bd3d477610f9521cbf33951f8c98d67dc094370",
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", res)
}

//  - ID: 0x0465f2d60eae4d51fb2f6bf333c6249d009e8193435e58a6a71516754c768d67 , Owner: Account Address ( 0xa6c86c562e55e8eefd4745928c49e9eb4df764a91bd38c60e08fd05f705b393f )
//   - ID: 0x4191040442a542afa68dcaf0c22cdb89ef8942513a9570dec60768bc8cbeb932 , Owner: Shared
//   - ID: 0x9593e602ee962bc9925dc37ac2e162c4f4b703dd5cd8390609de0e3ad7eecb9c , Owner: Object ID: ( 0xed9f35cf5e2edd4b5f9b572265b4455fc8c1897341f2172701577b698e03c1bf )
