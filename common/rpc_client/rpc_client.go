package rpc_client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/block-vision/sui-go-sdk/common/sui_error"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/gorilla/websocket"
)

type RPCClient struct {
	ws         *websocket.Conn
	httpClient *http.Client
	baseUrl    string
	timeout    time.Duration
}

const defaultTimeout = time.Second * 5

func NewRPCClient(baseUrl string) *RPCClient {
	var ws *websocket.Conn
	var err error
	if strings.HasPrefix(baseUrl, "ws") {
		ws, _, err = websocket.DefaultDialer.Dial(baseUrl, nil)
		if err != nil {
			log.Fatal("new websocket client failed, err: ", err)
			return nil
		}
	}
	return &RPCClient{
		ws:         ws,
		httpClient: &http.Client{},
		baseUrl:    baseUrl,
		timeout:    defaultTimeout,
	}
}

func (r *RPCClient) WithTimeout(timeout time.Duration) {
	r.timeout = timeout
}

func (r *RPCClient) Request(ctx context.Context, op models.Operation) ([]byte, error) {
	if strings.HasPrefix(r.baseUrl, "ws") {
		return r.websocketRequest(ctx, op)
	}
	return r.httpRequest(ctx, op)
}

func (r *RPCClient) websocketRequest(ctx context.Context, op models.Operation) ([]byte, error) {
	if r.ws == nil {
		return nil, sui_error.ErrInvalidWebsocketClient
	}
	jsonRPCReq := models.JsonRPCRequest{
		JsonRPC: "2.0",
		ID:      1,
		Method:  op.Method,
		Params:  op.Params,
	}
	err := r.ws.WriteJSON(jsonRPCReq)
	if err != nil {
		return nil, err
	}
	_, rsp, err := r.ws.ReadMessage()
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (r *RPCClient) httpRequest(ctx context.Context, op models.Operation) ([]byte, error) {
	jsonRPCReq := models.JsonRPCRequest{
		JsonRPC: "2.0",
		ID:      1,
		Method:  op.Method,
		Params:  op.Params,
	}
	reqBytes, err := json.Marshal(jsonRPCReq)
	if err != nil {
		return []byte{}, err
	}
	request, err := http.NewRequest("POST", r.baseUrl, bytes.NewBuffer(reqBytes))
	if err != nil {
		return []byte{}, err
	}
	request = request.WithContext(ctx)
	request.Header.Add("Content-Type", "application/json")
	rsp, err := r.httpClient.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer rsp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bodyBytes, nil
}

func (r *RPCClient) PostMultiple(ctx context.Context, method string, params [][]interface{}) ([]byte, error) {
	defaultRequest := []map[string]interface{}{}
	for i := range params {
		defaultRequest = append(defaultRequest, map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"method":  method,
			"params":  params[i],
		})
	}
	reqBody, err := json.Marshal(defaultRequest)
	if err != nil {
		return nil, err
	}
	resp, err := r.httpClient.Post(r.baseUrl, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
