// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"
	"github.com/block-vision/sui-go-sdk/common/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
)

type IReadNameServiceFromSuiAPI interface {
	SuiXResolveNameServiceAddress(ctx context.Context, req models.SuiXResolveNameServiceAddressRequest) (string, error)
	SuiXResolveNameServiceNames(ctx context.Context, req models.SuiXResolveNameServiceNamesRequest) (models.SuiXResolveNameServiceNamesResponse, error)
}

type suiReadNameServiceFromSuiImpl struct {
	handler *BaseRequestHandler
}

func newSuiReadNameServiceFromSuiImpl(conn *httpconn.HttpConn) *suiReadNameServiceFromSuiImpl {
	return &suiReadNameServiceFromSuiImpl{
		handler: NewBaseRequestHandler(conn),
	}
}

// SuiXResolveNameServiceAddress implements the method `suix_resolveNameServiceAddress`, get the resolved address given resolver and name.
func (s *suiReadNameServiceFromSuiImpl) SuiXResolveNameServiceAddress(ctx context.Context, req models.SuiXResolveNameServiceAddressRequest) (string, error) {
	var rsp string
	params := []interface{}{req.Name}
	err := s.handler.ExecuteRequest(ctx, "suix_resolveNameServiceAddress", params, &rsp)
	return rsp, err
}

// SuiXResolveNameServiceNames implements the method `suix_resolveNameServiceNames`, return the resolved names given address, if multiple names are resolved, the first one is the primary name.
func (s *suiReadNameServiceFromSuiImpl) SuiXResolveNameServiceNames(ctx context.Context, req models.SuiXResolveNameServiceNamesRequest) (models.SuiXResolveNameServiceNamesResponse, error) {
	var rsp models.SuiXResolveNameServiceNamesResponse
	params := []interface{}{req.Address, req.Cursor, req.Limit}
	err := s.handler.ExecuteRequest(ctx, "suix_resolveNameServiceNames", params, &rsp)
	return rsp, err
}