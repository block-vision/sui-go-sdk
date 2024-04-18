package sui

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/block-vision/sui-go-sdk/common/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

type IReadMoveFromSuiAPI interface {
	SuiGetMoveFunctionArgTypes(ctx context.Context, req models.GetMoveFunctionArgTypesRequest) (models.GetMoveFunctionArgTypesResponse, error)
	SuiGetNormalizedMoveModulesByPackage(ctx context.Context, req models.GetNormalizedMoveModulesByPackageRequest) (models.GetNormalizedMoveModulesByPackageResponse, error)
	SuiGetNormalizedMoveModule(ctx context.Context, req models.GetNormalizedMoveModuleRequest) (models.GetNormalizedMoveModuleResponse, error)
	SuiGetNormalizedMoveStruct(ctx context.Context, req models.GetNormalizedMoveStructRequest) (models.GetNormalizedMoveStructResponse, error)
	SuiGetNormalizedMoveFunction(ctx context.Context, req models.GetNormalizedMoveFunctionRequest) (models.GetNormalizedMoveFunctionResponse, error)
}

type suiReadMoveFromSuiImpl struct {
	conn *httpconn.HttpConn
}

// SuiGetMoveFunctionArgTypes implements method `sui_getMoveFunctionArgTypes`, return the argument types of a Move function based on normalized type.
func (s *suiReadMoveFromSuiImpl) SuiGetMoveFunctionArgTypes(ctx context.Context, req models.GetMoveFunctionArgTypesRequest) (models.GetMoveFunctionArgTypesResponse, error) {
	var rsp models.GetMoveFunctionArgTypesResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getMoveFunctionArgTypes",
		Params: []interface{}{
			req.Package,
			req.Module,
			req.Function,
		},
	})
	if err != nil {
		return rsp, nil
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

// SuiGetNormalizedMoveModulesByPackage implements method `sui_getNormalizedMoveModulesByPackage`, return the structured representations of all modules in the given package.
func (s *suiReadMoveFromSuiImpl) SuiGetNormalizedMoveModulesByPackage(ctx context.Context, req models.GetNormalizedMoveModulesByPackageRequest) (models.GetNormalizedMoveModulesByPackageResponse, error) {
	var rsp models.GetNormalizedMoveModulesByPackageResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getNormalizedMoveModulesByPackage",
		Params: []interface{}{
			req.Package,
		},
	})
	if err != nil {
		return rsp, nil
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

// SuiGetNormalizedMoveModule implements method `sui_getNormalizedMoveModule`, return a structured representation of a Move module.
func (s *suiReadMoveFromSuiImpl) SuiGetNormalizedMoveModule(ctx context.Context, req models.GetNormalizedMoveModuleRequest) (models.GetNormalizedMoveModuleResponse, error) {
	var rsp models.GetNormalizedMoveModuleResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getNormalizedMoveModule",
		Params: []interface{}{
			req.Package,
			req.ModuleName,
		},
	})
	if err != nil {
		return rsp, nil
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

// SuiGetNormalizedMoveStruct implements method `sui_getNormalizedMoveStruct`, return a structured representation of a Move struct.
func (s *suiReadMoveFromSuiImpl) SuiGetNormalizedMoveStruct(ctx context.Context, req models.GetNormalizedMoveStructRequest) (models.GetNormalizedMoveStructResponse, error) {
	var rsp models.GetNormalizedMoveStructResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getNormalizedMoveStruct",
		Params: []interface{}{
			req.Package,
			req.ModuleName,
			req.StructName,
		},
	})
	if err != nil {
		return rsp, nil
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

// SuiGetNormalizedMoveFunction implements method `sui_getNormalizedMoveFunction`, return a structured representation of a Move function.
func (s *suiReadMoveFromSuiImpl) SuiGetNormalizedMoveFunction(ctx context.Context, req models.GetNormalizedMoveFunctionRequest) (models.GetNormalizedMoveFunctionResponse, error) {
	var rsp models.GetNormalizedMoveFunctionResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getNormalizedMoveFunction",
		Params: []interface{}{
			req.Package,
			req.ModuleName,
			req.FunctionName,
		},
	})
	if err != nil {
		return rsp, nil
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
