package models

import "github.com/block-vision/sui-go-sdk/models/sui_json_rpc_types"

type GetMoveFunctionArgTypesRequest struct {
	Package  string
	Module   string
	Function string
}

type GetMoveFunctionArgTypesResponse []interface{}

type GetNormalizedMoveModulesByPackageRequest struct {
	Package string `json:"package"`
}

type GetNormalizedMoveModulesByPackageResponse map[string]sui_json_rpc_types.SuiMoveNormalizedModule

type GetNormalizedMoveModuleRequest struct {
	Package    string `json:"package"`
	ModuleName string `json:"moduleName"`
}

type GetNormalizedMoveModuleResponse sui_json_rpc_types.SuiMoveNormalizedModule

type GetNormalizedMoveStructRequest struct {
	Package    string `json:"package"`
	ModuleName string `json:"moduleName"`
	StructName string `json:"structName"`
}

type GetNormalizedMoveStructResponse sui_json_rpc_types.SuiMoveNormalizedStruct

type GetNormalizedMoveFunctionRequest struct {
	Package      string `json:"package"`
	ModuleName   string `json:"moduleName"`
	FunctionName string `json:"functionName"`
}

type GetNormalizedMoveFunctionResponse sui_json_rpc_types.SuiMoveNormalizedFunction
