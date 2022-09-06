package sui

import (
	"context"
	"github.com/block-vision/sui-go-sdk/httpconn"
	"github.com/tidwall/gjson"
)

type IBaseAPI interface {
	SuiCall(ctx context.Context, method string, params ...interface{}) (interface{}, error)
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
	return gjson.ParseBytes(resp).String(), nil
}
