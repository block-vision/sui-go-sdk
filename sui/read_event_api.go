// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"

	"github.com/block-vision/sui-go-sdk/common/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
)

type IReadEventFromSuiAPI interface {
	SuiGetEvents(ctx context.Context, req models.SuiGetEventsRequest) (models.GetEventsResponse, error)
	SuiXQueryEvents(ctx context.Context, req models.SuiXQueryEventsRequest) (models.PaginatedEventsResponse, error)
}

type suiReadEventFromSuiImpl struct {
	handler *BaseRequestHandler
}

func newSuiReadEventFromSuiImpl(conn *httpconn.HttpConn) *suiReadEventFromSuiImpl {
	return &suiReadEventFromSuiImpl{
		handler: NewBaseRequestHandler(conn),
	}
}

// SuiXQueryEvents implements the method `suix_queryEvents`, gets list of events for a specified query criteria.
func (s *suiReadEventFromSuiImpl) SuiXQueryEvents(ctx context.Context, req models.SuiXQueryEventsRequest) (models.PaginatedEventsResponse, error) {
	var rsp models.PaginatedEventsResponse
	params := []interface{}{req.SuiEventFilter, req.Cursor, req.Limit, req.DescendingOrder}
	err := s.handler.ExecuteRequestWithValidation(ctx, "suix_queryEvents", params, req, &rsp)
	return rsp, err
}

// SuiGetEvents implements the method `sui_getEvents`, gets transaction events.
func (s *suiReadEventFromSuiImpl) SuiGetEvents(ctx context.Context, req models.SuiGetEventsRequest) (models.GetEventsResponse, error) {
	var rsp models.GetEventsResponse
	params := []interface{}{req.Digest}
	err := s.handler.ExecuteRequestWithValidation(ctx, "sui_getEvents", params, req, &rsp)
	return rsp, err
}
