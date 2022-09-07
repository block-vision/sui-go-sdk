package httpconn

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/block-vision/sui-go-sdk/models"
	"io/ioutil"
	"net/http"
	"time"
)

const defaultTimeout = time.Second * 5

type HttpConn struct {
	c       *http.Client
	baseUrl string
	timeout time.Duration
}

func NewHttpConn(baseUrl string) *HttpConn {
	return &HttpConn{
		c:       &http.Client{},
		baseUrl: baseUrl,
		timeout: defaultTimeout,
	}
}

func (h *HttpConn) Request(ctx context.Context, op Operation) ([]byte, error) {
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
	request, err := http.NewRequest("POST", h.baseUrl, bytes.NewBuffer(reqBytes))
	if err != nil {
		return []byte{}, err
	}
	request = request.WithContext(ctx)
	request.Header.Add("Content-Type", "application/json")
	rsp, err := h.c.Do(request)
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
