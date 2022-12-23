package sui

import (
	"context"
	"fmt"
	"testing"

	"github.com/block-vision/sui-go-sdk/models"
)

func TestOnAPIs(t *testing.T) {
	t.Run("test on sui_getEvents", func(t *testing.T) {
		cli := NewSuiClient("https://fullnode.devnet.sui.io/")
		resp, err := cli.GetEvents(context.Background(), models.GetEventsRequest{
			EventQuery: "All",
			Cursor:     nil,
			Limit:      1,
			DescOrder:  true,
		})
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Printf("%+v", resp)
	})
}
