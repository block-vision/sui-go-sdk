// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"
	"github.com/block-vision/sui-go-sdk/common/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
)

type IReadObjectFromSuiAPI interface {
	SuiGetObject(ctx context.Context, req models.SuiGetObjectRequest) (models.SuiObjectResponse, error)
	SuiXGetOwnedObjects(ctx context.Context, req models.SuiXGetOwnedObjectsRequest) (models.PaginatedObjectsResponse, error)
	SuiMultiGetObjects(ctx context.Context, req models.SuiMultiGetObjectsRequest) ([]*models.SuiObjectResponse, error)
	SuiXGetDynamicField(ctx context.Context, req models.SuiXGetDynamicFieldRequest) (models.PaginatedDynamicFieldInfoResponse, error)
	SuiXGetDynamicFieldObject(ctx context.Context, req models.SuiXGetDynamicFieldObjectRequest) (models.SuiObjectResponse, error)
	SuiTryGetPastObject(ctx context.Context, req models.SuiTryGetPastObjectRequest) (models.PastObjectResponse, error)
	SuiGetLoadedChildObjects(ctx context.Context, req models.SuiGetLoadedChildObjectsRequest) (models.ChildObjectsResponse, error)
	SuiTryMultiGetPastObjects(ctx context.Context, req models.SuiTryMultiGetPastObjectsRequest) ([]*models.PastObjectResponse, error)
}

type suiReadObjectFromSuiImpl struct {
	handler *BaseRequestHandler
}

func newSuiReadObjectFromSuiImpl(conn *httpconn.HttpConn) *suiReadObjectFromSuiImpl {
	return &suiReadObjectFromSuiImpl{
		handler: NewBaseRequestHandler(conn),
	}
}

// SuiGetObject implements the method `sui_getObject`, gets the object information for a specified object.
func (s *suiReadObjectFromSuiImpl) SuiGetObject(ctx context.Context, req models.SuiGetObjectRequest) (models.SuiObjectResponse, error) {
	var rsp models.SuiObjectResponse
	params := []interface{}{req.ObjectId, req.Options}
	err := s.handler.ExecuteRequest(ctx, "sui_getObject", params, &rsp)
	return rsp, err
}

// SuiXGetOwnedObjects implements the method `suix_getOwnedObjects`, gets the list of objects owned by an address.
func (s *suiReadObjectFromSuiImpl) SuiXGetOwnedObjects(ctx context.Context, req models.SuiXGetOwnedObjectsRequest) (models.PaginatedObjectsResponse, error) {
	var rsp models.PaginatedObjectsResponse
	params := []interface{}{req.Address, req.Query, req.Cursor, req.Limit}
	err := s.handler.ExecuteRequestWithValidation(ctx, "suix_getOwnedObjects", params, req, &rsp)
	return rsp, err
}

// SuiMultiGetObjects implements the method `sui_multiGetObjects`, gets the object data for a list of objects.
func (s *suiReadObjectFromSuiImpl) SuiMultiGetObjects(ctx context.Context, req models.SuiMultiGetObjectsRequest) ([]*models.SuiObjectResponse, error) {
	var rsp []*models.SuiObjectResponse
	params := []interface{}{req.ObjectIds, req.Options}
	err := s.handler.ExecuteRequest(ctx, "sui_multiGetObjects", params, &rsp)
	return rsp, err
}

// SuiXGetDynamicField implements the method `suix_getDynamicFields`, gets the list of dynamic field objects owned by an object.
func (s *suiReadObjectFromSuiImpl) SuiXGetDynamicField(ctx context.Context, req models.SuiXGetDynamicFieldRequest) (models.PaginatedDynamicFieldInfoResponse, error) {
	var rsp models.PaginatedDynamicFieldInfoResponse
	params := []interface{}{req.ObjectId, req.Cursor, req.Limit}
	err := s.handler.ExecuteRequestWithValidation(ctx, "suix_getDynamicFields", params, req, &rsp)
	return rsp, err
}

// SuiXGetDynamicFieldObject implements the method `suix_getDynamicFieldObject`, gets the dynamic field object information for a specified object.
func (s *suiReadObjectFromSuiImpl) SuiXGetDynamicFieldObject(ctx context.Context, req models.SuiXGetDynamicFieldObjectRequest) (models.SuiObjectResponse, error) {
	var rsp models.SuiObjectResponse
	params := []interface{}{req.ObjectId, req.DynamicFieldName}
	err := s.handler.ExecuteRequest(ctx, "suix_getDynamicFieldObject", params, &rsp)
	return rsp, err
}

// SuiTryGetPastObject implements the method `sui_tryGetPastObject`, gets the object information for a specified version.
// There is no guarantee that objects with past versions can be retrieved by this API. The result may vary across nodes depending on their pruning policies.
func (s *suiReadObjectFromSuiImpl) SuiTryGetPastObject(ctx context.Context, req models.SuiTryGetPastObjectRequest) (models.PastObjectResponse, error) {
	var rsp models.PastObjectResponse
	params := []interface{}{req.ObjectId, req.Version, req.Options}
	err := s.handler.ExecuteRequest(ctx, "sui_tryGetPastObject", params, &rsp)
	return rsp, err
}

// SuiGetLoadedChildObjects implements the method `sui_getLoadedChildObjects`, return the object information for a specified digest.
func (s *suiReadObjectFromSuiImpl) SuiGetLoadedChildObjects(ctx context.Context, req models.SuiGetLoadedChildObjectsRequest) (models.ChildObjectsResponse, error) {
	var rsp models.ChildObjectsResponse
	params := []interface{}{req.Digest}
	err := s.handler.ExecuteRequest(ctx, "sui_getLoadedChildObjects", params, &rsp)
	return rsp, err
}

// SuiTryMultiGetPastObjects implements the method `sui_tryMultiGetPastObjects`,
// note there is no software-level guarantee/SLA that objects with past versions can be retrieved by this API,
// even if the object and version exists/existed. The result may vary across nodes depending on their pruning policies.
// Return the object information for a specified version.
func (s *suiReadObjectFromSuiImpl) SuiTryMultiGetPastObjects(ctx context.Context, req models.SuiTryMultiGetPastObjectsRequest) ([]*models.PastObjectResponse, error) {
	var rsp []*models.PastObjectResponse
	params := []interface{}{req.MultiGetPastObjects, req.Options}
	err := s.handler.ExecuteRequest(ctx, "sui_tryMultiGetPastObjects", params, &rsp)
	return rsp, err
}