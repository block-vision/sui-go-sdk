// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/block-vision/sui-go-sdk/common/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

type IReadEventFromSuiAPI interface {
	SuiGetEvents(ctx context.Context, req models.SuiGetEventsRequest) (models.GetEventsResponse, error)
	SuiXQueryEvents(ctx context.Context, req models.SuiXQueryEventsRequest) (models.PaginatedEventsResponse, error)
}

type suiReadEventFromSuiImpl struct {
	conn *httpconn.HttpConn
}

// SuiGetEvents implements the method `sui_getEvents`, gets transaction events.
func (s *suiReadEventFromSuiImpl) SuiGetEvents(ctx context.Context, req models.SuiGetEventsRequest) (models.GetEventsResponse, error) {
	var rsp models.GetEventsResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getEvents",
		Params: []interface{}{
			req.Digest,
		},
	})
	if err != nil {
		return rsp, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXQueryEvents implements the method `suix_queryEvents`, gets list of events for a specified query criteria.
func (s *suiReadEventFromSuiImpl) SuiXQueryEvents(ctx context.Context, req models.SuiXQueryEventsRequest) (models.PaginatedEventsResponse, error) {
	var rsp models.PaginatedEventsResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_queryEvents",
		Params: []interface{}{
			req.SuiEventFilter,
			req.Cursor,
			req.Limit,
			req.DescendingOrder,
		},
	})
	if err != nil {
		return rsp, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}
