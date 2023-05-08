package wsconn

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"log"
	"time"
)

type WsConn struct {
	Conn  *websocket.Conn
	wsUrl string
}

type CallOp struct {
	Method string
	Params []interface{}
}

func NewWsConn(wsUrl string) *WsConn {
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(wsUrl, nil)

	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err, wsUrl)
	}

	return &WsConn{
		Conn:  conn,
		wsUrl: wsUrl,
	}
}

func (w *WsConn) Call(ctx context.Context, op CallOp, receiveMsgCh chan []byte) error {
	jsonRPCCall := models.JsonRPCRequest{
		JsonRPC: "2.0",
		ID:      time.Now().UnixMilli(),
		Method:  op.Method,
		Params:  op.Params,
	}

	callBytes, err := json.Marshal(jsonRPCCall)
	if err != nil {
		return err
	}

	err = w.Conn.WriteMessage(websocket.TextMessage, callBytes)
	if nil != err {
		return err
	}

	_, messageData, err := w.Conn.ReadMessage()
	if nil != err {
		return err
	}

	var rsp SubscriptionResp
	if gjson.ParseBytes(messageData).Get("error").Exists() {
		return fmt.Errorf(gjson.ParseBytes(messageData).Get("error").String())
	}

	err = json.Unmarshal([]byte(gjson.ParseBytes(messageData).String()), &rsp)
	if err != nil {
		return err
	}

	fmt.Printf("establish successfully, subscriptionID: %d, Waiting to accept data...\n", rsp.Result)

	go func(conn *websocket.Conn) {
		for {
			messageType, messageData, err := conn.ReadMessage()
			if nil != err {
				log.Println(err)
				break
			}
			switch messageType {
			case websocket.TextMessage:
				receiveMsgCh <- messageData

			default:
				continue
			}
		}
	}(w.Conn)

	return nil
}
