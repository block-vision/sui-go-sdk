package httpconn

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/block-vision/sui-go-sdk/models"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	defaultTimeout    = time.Second * 5
	defaultRetryCount = 3
)

type HttpConn struct {
	c            *http.Client
	rl           *rate.Limiter
	rpcUrl       string
	backupRPCURL []string
	retryCount   int
	timeout      time.Duration
}

func newDefaultRateLimiter() *rate.Limiter {
	rateLimiter := rate.NewLimiter(rate.Every(1*time.Second), 10000) // 10000 request every 1 seconds
	return rateLimiter
}

func NewHttpConn(rpcUrl string, backupRPCs []string) *HttpConn {
	return &HttpConn{
		c:            &http.Client{},
		rpcUrl:       rpcUrl,
		backupRPCURL: backupRPCs,
		timeout:      defaultTimeout,
		retryCount:   defaultRetryCount,
	}
}

func NewCustomHttpConn(rpcUrl string, cli *http.Client) *HttpConn {
	return &HttpConn{
		c:       cli,
		rpcUrl:  rpcUrl,
		timeout: defaultTimeout,
	}
}

func (h *HttpConn) Request(ctx context.Context, op Operation) ([]byte, error) {
	jsonRPCReq := models.JsonRPCRequest{
		JsonRPC: "2.0",
		ID:      time.Now().UnixMilli(),
		Method:  op.Method,
		Params:  op.Params,
	}
	reqBytes, err := json.Marshal(jsonRPCReq)
	if err != nil {
		return []byte{}, err
	}

	// 尝试所有可用的 RPC URL
	rpcURLs := append([]string{h.rpcUrl}, h.backupRPCURL...)
	var lastErr error

	for _, rpc := range rpcURLs {
		request, err := http.NewRequest("POST", rpc, bytes.NewBuffer(reqBytes))
		if err != nil {
			lastErr = fmt.Errorf("new request %s err: %v rpc: %s", op.Method, err, rpc)
			continue
		}
		request = request.WithContext(ctx)
		request.Header.Add("Content-Type", "application/json")

		rsp, err := h.c.Do(request.WithContext(ctx))
		if err != nil {
			lastErr = fmt.Errorf("request %s err: %v rpc: %s", op.Method, err, rpc)
			continue
		}
		defer rsp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			lastErr = err
			continue
		}
		return bodyBytes, nil
	}

	return []byte{}, lastErr
}
