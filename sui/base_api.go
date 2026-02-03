// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"
	"errors"
	"github.com/block-vision/sui-go-sdk/common/httpconn"
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
	parsedJson := gjson.ParseBytes(resp)
	if parsedJson.Get("error").Exists() {
		return nil, errors.New(parsedJson.Get("error").String())
	}
	return parsedJson.String(), nil
}
