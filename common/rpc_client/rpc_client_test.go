package rpc_client

import (
	"context"
	"fmt"
	"testing"

	"github.com/shoshinsquare/sui-go-sdk/models"
)

func TestOnNewRPCClient(t *testing.T) {
	cli := NewRPCClient("https://fullnode.devnet.sui.io:443")
	content, err := cli.httpRequest(context.Background(), models.Operation{
		Method: "sui_getTransaction",
		Params: []interface{}{
			"3TpIBOxyEbKFB6Z7WuzGHYqYTkAVqlY/Gyl+Xcz9mu8=",
		},
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(content))
}
