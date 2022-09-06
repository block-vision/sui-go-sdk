package sui

import (
	"context"
	"errors"
	"github.com/block-vision/sui-go-sdk/httpconn"
	"github.com/tidwall/gjson"
)

type IBaseAPI interface {
	SuiCall(ctx context.Context, method string, params ...interface{}) (interface{}, error)
	SyncAccountState(ctx context.Context, address string) error
}

type suiBaseImpl struct {
	conn *httpconn.HttpConn
}

// SuiCall send customized request to Sui Node endpoint.
func (s *suiBaseImpl) SuiCall(ctx context.Context, method string, params ...interface{}) (interface{}, error) {
	resp, err := s.conn.Request(ctx, httpconn.Operation{
		Method: method,
		Params: params,
	})
	if err != nil {
		return nil, err
	}
	if gjson.ParseBytes(resp).Get("error").Exists() {
		return nil, errors.New(gjson.ParseBytes(resp).Get("error").String())
	}
	return gjson.ParseBytes(resp).String(), nil
}

// SyncAccountState implements method `sui_syncAccountState`.
// synchronize client state with validators.
func (s *suiBaseImpl) SyncAccountState(ctx context.Context, address string) error {
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_syncAccountState",
		Params: []interface{}{
			address,
		},
	})
	if err != nil {
		return err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	return nil
}
