package sui

import (
	"context"
	"encoding/json"
	"github.com/block-vision/sui-go-sdk/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui_error"
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
	conn *httpconn.HttpConn
}

func (s *suiReadMoveFromSuiImpl) GetMoveFunctionArgTypes(ctx context.Context, req models.GetMoveFunctionArgTypesRequest, opts ...interface{}) (models.GetMoveFunctionArgTypesResponse, error) {
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
		return models.GetMoveFunctionArgTypesResponse{}, nil
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetMoveFunctionArgTypesResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetMoveFunctionArgTypesResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadMoveFromSuiImpl) GetNormalizedMoveModulesByPackage(ctx context.Context, req models.GetNormalizedMoveModulesByPackageRequest, opts ...interface{}) (models.GetNormalizedMoveModulesByPackageResponse, error) {
	var rsp models.GetNormalizedMoveModulesByPackageResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getNormalizedMoveModulesByPackage",
		Params: []interface{}{
			req.Package,
		},
	})
	if err != nil {
		return models.GetNormalizedMoveModulesByPackageResponse{}, nil
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetNormalizedMoveModulesByPackageResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetNormalizedMoveModulesByPackageResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadMoveFromSuiImpl) GetNormalizedMoveModule(ctx context.Context, req models.GetNormalizedMoveModuleRequest, opts ...interface{}) (models.GetNormalizedMoveModuleResponse, error) {
	var rsp models.GetNormalizedMoveModuleResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "sui_getNormalizedMoveModule",
		Params: []interface{}{
			req.Package,
			req.ModuleName,
		},
	})
	if err != nil {
		return models.GetNormalizedMoveModuleResponse{}, nil
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetNormalizedMoveModuleResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetNormalizedMoveModuleResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadMoveFromSuiImpl) GetNormalizedMoveStruct(ctx context.Context, req models.GetNormalizedMoveStructRequest, opts ...interface{}) (models.GetNormalizedMoveStructResponse, error) {
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
		return models.GetNormalizedMoveStructResponse{}, nil
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetNormalizedMoveStructResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetNormalizedMoveStructResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadMoveFromSuiImpl) GetNormalizedMoveFunction(ctx context.Context, req models.GetNormalizedMoveFunctionRequest, opts ...interface{}) (models.GetNormalizedMoveFunctionResponse, error) {
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
		return models.GetNormalizedMoveFunctionResponse{}, nil
	}
	if !gjson.ValidBytes(respBytes) {
		return models.GetNormalizedMoveFunctionResponse{}, sui_error.ErrInvalidJson
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetNormalizedMoveFunctionResponse{}, err
	}
	return rsp, nil
}
