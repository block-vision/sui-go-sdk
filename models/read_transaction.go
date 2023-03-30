package models

import (
	"github.com/shoshinsquare/sui-go-sdk/models/sui_json_rpc_types"
)

type GetRecentTransactionRequest struct {
	Count uint64
}

type GetRecentTransactionResponse struct {
	Result []GetTransactionMetaData `json:"result"`
}

type GetTransactionMetaData struct {
	GatewayTxSeqNumber uint64 `json:"gatewayTxSeqNumber"`
	TransactionDigest  string `json:"transactionDigest"`
}

type GetTotalTransactionNumberRequest struct{}
type GetTotalTransactionNumberResponse struct {
	TotalNumberOfTransaction uint64 `json:"totalNumberOfTransaction"`
}

type GetTransactionRequest struct {
	Digest string `json:"digest"`
}

type GetTransactionResponse struct {
	Certificate sui_json_rpc_types.SuiCertifiedTransaction `json:"certificate"`
	Effects     sui_json_rpc_types.SuiTransactionEffects   `json:"effects"`
	TimestampMs uint64                                     `json:"timestamp_ms,omitempty"`

	// ParserData with Schema
	ParsedData interface{} `json:"parsed_data,omitempty"`
}

type GetTransactionsByInputObjectRequest struct {
	ObjectID string `json:"objectID"`
}

type GetTransactionsByInputObjectResponse struct {
	Result []GetTransactionMetaData `json:"result"`
}

type GetTransactionsByMoveFunctionRequest struct {
	Package  string `json:"package"`
	Module   string `json:"module"`
	Function string `json:"function"`
}
type GetTransactionsByMoveFunctionResponse struct {
	Result []GetTransactionMetaData `json:"result"`
}

type GetTransactionsByMutatedObjectRequest struct {
	ObjectID string
}
type GetTransactionsByMutatedObjectResponse struct {
	Result []GetTransactionMetaData `json:"result"`
}

type GetTransactionsFromAddressRequest struct {
	Addr string `json:"addr"`
}

type GetTransactionsFromAddressResponse struct {
	Result []GetTransactionMetaData `json:"result"`
}

type GetTransactionsInRangeRequest struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}

type GetTransactionsInRangeResponse struct {
	Result []GetTransactionMetaData `json:"result"`
}

type GetTransactionsToAddressRequest struct {
	Addr string
}
type GetTransactionsToAddressResponse struct {
	Result []GetTransactionMetaData
}

type GetTransactionAuthSignersRequest struct {
	Digest string
}

type GetTransactionAuthSignersResponse struct {
	Signers []string `json:"signers"`
}
