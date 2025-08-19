// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"
	"github.com/block-vision/sui-go-sdk/common/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
)

type IReadMoveFromSuiAPI interface {
	SuiGetMoveFunctionArgTypes(ctx context.Context, req models.GetMoveFunctionArgTypesRequest) (models.GetMoveFunctionArgTypesResponse, error)
	SuiGetNormalizedMoveFunction(ctx context.Context, req models.GetNormalizedMoveFunctionRequest) (models.GetNormalizedMoveFunctionResponse, error)
	SuiGetNormalizedMoveModule(ctx context.Context, req models.GetNormalizedMoveModuleRequest) (models.GetNormalizedMoveModuleResponse, error)
	SuiGetNormalizedMoveStruct(ctx context.Context, req models.GetNormalizedMoveStructRequest) (models.GetNormalizedMoveStructResponse, error)
	SuiGetNormalizedMoveModulesByPackage(ctx context.Context, req models.GetNormalizedMoveModulesByPackageRequest) (models.GetNormalizedMoveModulesByPackageResponse, error)
}

type suiReadMoveFromSuiImpl struct {
	handler *BaseRequestHandler
}

func newSuiReadMoveFromSuiImpl(conn *httpconn.HttpConn) *suiReadMoveFromSuiImpl {
	return &suiReadMoveFromSuiImpl{
		handler: NewBaseRequestHandler(conn),
	}
}

// SuiGetMoveFunctionArgTypes implements the method `sui_getMoveFunctionArgTypes`, gets the argument types of a Move function, based on normalized Type.
func (s *suiReadMoveFromSuiImpl) SuiGetMoveFunctionArgTypes(ctx context.Context, req models.GetMoveFunctionArgTypesRequest) (models.GetMoveFunctionArgTypesResponse, error) {
	var rsp models.GetMoveFunctionArgTypesResponse
	params := []interface{}{req.Package, req.Module, req.Function}
	err := s.handler.ExecuteRequest(ctx, "sui_getMoveFunctionArgTypes", params, &rsp)
	return rsp, err
}

// SuiGetNormalizedMoveFunction implements the method `sui_getNormalizedMoveFunction`, gets a Move function's normalized representation.
func (s *suiReadMoveFromSuiImpl) SuiGetNormalizedMoveFunction(ctx context.Context, req models.GetNormalizedMoveFunctionRequest) (models.GetNormalizedMoveFunctionResponse, error) {
	var rsp models.GetNormalizedMoveFunctionResponse
	params := []interface{}{req.Package, req.ModuleName, req.FunctionName}
	err := s.handler.ExecuteRequest(ctx, "sui_getNormalizedMoveFunction", params, &rsp)
	return rsp, err
}

// SuiGetNormalizedMoveModule implements the method `sui_getNormalizedMoveModule`, gets a Move module's normalized representation.
func (s *suiReadMoveFromSuiImpl) SuiGetNormalizedMoveModule(ctx context.Context, req models.GetNormalizedMoveModuleRequest) (models.GetNormalizedMoveModuleResponse, error) {
	var rsp models.GetNormalizedMoveModuleResponse
	params := []interface{}{req.Package, req.ModuleName}
	err := s.handler.ExecuteRequest(ctx, "sui_getNormalizedMoveModule", params, &rsp)
	return rsp, err
}

// SuiGetNormalizedMoveStruct implements the method `sui_getNormalizedMoveStruct`, gets a Move struct's normalized representation.
func (s *suiReadMoveFromSuiImpl) SuiGetNormalizedMoveStruct(ctx context.Context, req models.GetNormalizedMoveStructRequest) (models.GetNormalizedMoveStructResponse, error) {
	var rsp models.GetNormalizedMoveStructResponse
	params := []interface{}{req.Package, req.ModuleName, req.StructName}
	err := s.handler.ExecuteRequest(ctx, "sui_getNormalizedMoveStruct", params, &rsp)
	return rsp, err
}

// SuiGetNormalizedMoveModulesByPackage implements the method `suix_getNormalizedMoveModulesByPackage`, gets normalized Move modules by package.
func (s *suiReadMoveFromSuiImpl) SuiGetNormalizedMoveModulesByPackage(ctx context.Context, req models.GetNormalizedMoveModulesByPackageRequest) (models.GetNormalizedMoveModulesByPackageResponse, error) {
	var rsp models.GetNormalizedMoveModulesByPackageResponse
	params := []interface{}{req.Package}
	err := s.handler.ExecuteRequest(ctx, "suix_getNormalizedMoveModulesByPackage", params, &rsp)
	return rsp, err
}