package sui

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/shoshinsquare/sui-go-sdk/common/rpc_client"
	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

type IReadMoveFromSuiAPI interface {
	GetMoveFunctionArgTypes(ctx context.Context, req models.GetMoveFunctionArgTypesRequest, opts ...interface{}) (models.GetMoveFunctionArgTypesResponse, error)
	GetNormalizedMoveModulesByPackage(ctx context.Context, req models.GetNormalizedMoveModulesByPackageRequest, opts ...interface{}) (models.GetNormalizedMoveModulesByPackageResponse, error)
	GetNormalizedMoveModule(ctx context.Context, req models.GetNormalizedMoveModuleRequest, opts ...interface{}) (models.GetNormalizedMoveModuleResponse, error)
	GetNormalizedMoveStruct(ctx context.Context, req models.GetNormalizedMoveStructRequest, opts ...interface{}) (models.GetNormalizedMoveStructResponse, error)
	GetNormalizedMoveFunction(ctx context.Context, req models.GetNormalizedMoveFunctionRequest, opts ...interface{}) (models.GetNormalizedMoveFunctionResponse, error)
}

type suiReadMoveFromSuiImpl struct {
	cli *rpc_client.RPCClient
}

// GetMoveFunctionArgTypes implements method `sui_getMoveFunctionArgTypes`
func (s *suiReadMoveFromSuiImpl) GetMoveFunctionArgTypes(ctx context.Context, req models.GetMoveFunctionArgTypesRequest, opts ...interface{}) (models.GetMoveFunctionArgTypesResponse, error) {
	var rsp models.GetMoveFunctionArgTypesResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getMoveFunctionArgTypes",
		Params: []interface{}{
			req.Package,
			req.Module,
			req.Function,
		},
	})
	if err != nil {
		return models.GetMoveFunctionArgTypesResponse{}, nil
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetMoveFunctionArgTypesResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetMoveFunctionArgTypesResponse{}, err
	}
	return rsp, nil
}

// GetNormalizedMoveModulesByPackage implements method `sui_getNormalizedMoveModulesByPackage`
func (s *suiReadMoveFromSuiImpl) GetNormalizedMoveModulesByPackage(ctx context.Context, req models.GetNormalizedMoveModulesByPackageRequest, opts ...interface{}) (models.GetNormalizedMoveModulesByPackageResponse, error) {
	var rsp models.GetNormalizedMoveModulesByPackageResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getNormalizedMoveModulesByPackage",
		Params: []interface{}{
			req.Package,
		},
	})
	if err != nil {
		return models.GetNormalizedMoveModulesByPackageResponse{}, nil
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetNormalizedMoveModulesByPackageResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetNormalizedMoveModulesByPackageResponse{}, err
	}
	return rsp, nil
}

// GetNormalizedMoveModule implements method `sui_getNormalizedMoveModule`.
func (s *suiReadMoveFromSuiImpl) GetNormalizedMoveModule(ctx context.Context, req models.GetNormalizedMoveModuleRequest, opts ...interface{}) (models.GetNormalizedMoveModuleResponse, error) {
	var rsp models.GetNormalizedMoveModuleResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getNormalizedMoveModule",
		Params: []interface{}{
			req.Package,
			req.ModuleName,
		},
	})
	if err != nil {
		return models.GetNormalizedMoveModuleResponse{}, nil
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetNormalizedMoveModuleResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetNormalizedMoveModuleResponse{}, err
	}
	return rsp, nil
}

// GetNormalizedMoveStruct implements method `sui_getNormalizedMoveStruct`.
func (s *suiReadMoveFromSuiImpl) GetNormalizedMoveStruct(ctx context.Context, req models.GetNormalizedMoveStructRequest, opts ...interface{}) (models.GetNormalizedMoveStructResponse, error) {
	var rsp models.GetNormalizedMoveStructResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getNormalizedMoveStruct",
		Params: []interface{}{
			req.Package,
			req.ModuleName,
			req.StructName,
		},
	})
	if err != nil {
		return models.GetNormalizedMoveStructResponse{}, nil
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetNormalizedMoveStructResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetNormalizedMoveStructResponse{}, err
	}
	return rsp, nil
}

// GetNormalizedMoveFunction implements method `sui_getNormalizedMoveFunction`.
func (s *suiReadMoveFromSuiImpl) GetNormalizedMoveFunction(ctx context.Context, req models.GetNormalizedMoveFunctionRequest, opts ...interface{}) (models.GetNormalizedMoveFunctionResponse, error) {
	var rsp models.GetNormalizedMoveFunctionResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getNormalizedMoveFunction",
		Params: []interface{}{
			req.Package,
			req.ModuleName,
			req.FunctionName,
		},
	})
	if err != nil {
		return models.GetNormalizedMoveFunctionResponse{}, nil
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetNormalizedMoveFunctionResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetNormalizedMoveFunctionResponse{}, err
	}
	return rsp, nil
}
