package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func TestGetObject(t *testing.T) {
	cli := sui.NewSuiClient("https://sui-testnet-endpoint.blockvision.org/")

	resp, err := cli.GetObject(context.Background(), models.GetObjectRequest{
		ObjectID: "0xf8e8a03e4112b68f449292d82008d922602fab0ab18a4f1a73c61aa523050100",
	})
	if err != nil {
		fmt.Println(err)
	}

	t.Logf("%+v\n", resp)
}
