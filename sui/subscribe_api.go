package sui

import (
	"context"
	"encoding/json"
	"log"

	"github.com/block-vision/sui-go-sdk/common/wsconn"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

type ISubscribeAPI interface {
	SubscribeEvent(ctx context.Context, req models.SuiXSubscribeEventsRequest, msgCh chan models.SuiEventResponse) error
	SubscribeTransaction(ctx context.Context, req models.SuiXSubscribeTransactionsRequest, msgCh chan models.SuiEffects) error
}

type suiSubscribeImpl struct {
	conn *wsconn.WsConn
}

// SubscribeEvent implements the method `suix_subscribeEvent`, subscribe to a stream of Sui event.
func (s *suiSubscribeImpl) SubscribeEvent(ctx context.Context, req models.SuiXSubscribeEventsRequest, msgCh chan models.SuiEventResponse) error {
	rsp := make(chan []byte, 10)
	err := s.conn.Call(ctx, wsconn.CallOp{
		Method: "suix_subscribeEvent",
		Params: []interface{}{
			req.SuiEventFilter,
		},
	}, rsp)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case messageData := <-rsp:
				var result models.SuiEventResponse
				if gjson.ParseBytes(messageData).Get("error").Exists() {
					log.Fatal(gjson.ParseBytes(messageData).Get("error").String())
				}

				err := json.Unmarshal([]byte(gjson.ParseBytes(messageData).Get("params.result").String()), &result)
				if err != nil {
					log.Fatal(err)
				}

				msgCh <- result
			}
		}
	}()

	return nil
}

// SubscribeTransaction implements the method `suix_subscribeTransaction`, subscribe to a stream of Sui transaction effects.
func (s *suiSubscribeImpl) SubscribeTransaction(ctx context.Context, req models.SuiXSubscribeTransactionsRequest, msgCh chan models.SuiEffects) error {
	rsp := make(chan []byte, 10)
	err := s.conn.Call(ctx, wsconn.CallOp{
		Method: "suix_subscribeTransaction",
		Params: []interface{}{
			req.TransactionFilter,
		},
	}, rsp)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case messageData := <-rsp:
				var result models.SuiEffects
				if gjson.ParseBytes(messageData).Get("error").Exists() {
					log.Fatal(gjson.ParseBytes(messageData).Get("error").String())
				}

				err := json.Unmarshal([]byte(gjson.ParseBytes(messageData).Get("params.result").String()), &result)
				if err != nil {
					log.Fatal(err)
				}

				msgCh <- result
			}
		}
	}()

	return nil
}
