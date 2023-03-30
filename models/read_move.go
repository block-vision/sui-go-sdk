package models

import "github.com/shoshinsquare/sui-go-sdk/models/sui_json_rpc_types"

type GetMoveFunctionArgTypesRequest struct {
	Package  string
	Module   string
	Function string
}

type GetMoveFunctionArgTypesResponse struct {
	Result []interface{} `json:"result"`
}

type GetNormalizedMoveModulesByPackageRequest struct {
	Package string `json:"package"`
}

type GetNormalizedMoveModulesByPackageResponse struct {
	Result map[string]sui_json_rpc_types.SuiMoveNormalizedModule
}

type GetNormalizedMoveModuleRequest struct {
	Package    string `json:"package"`
	ModuleName string `json:"moduleName"`
}

type GetNormalizedMoveModuleResponse struct {
	Result sui_json_rpc_types.SuiMoveNormalizedModule
}

type GetNormalizedMoveStructRequest struct {
	Package    string `json:"package"`
	ModuleName string `json:"moduleName"`
	StructName string `json:"structName"`
}

type GetNormalizedMoveStructResponse struct {
	Result sui_json_rpc_types.SuiMoveNormalizedStruct
}

type GetNormalizedMoveFunctionRequest struct {
	Package      string `json:"package"`
	ModuleName   string `json:"moduleName"`
	FunctionName string `json:"functionName"`
}

type GetNormalizedMoveFunctionResponse struct {
	Result sui_json_rpc_types.SuiMoveNormalizedFunction
}
