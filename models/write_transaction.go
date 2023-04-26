package models

import (
	"github.com/block-vision/sui-go-sdk/models/sui_types"
)

type MoveCallRequest struct {
	Signer          string        `json:"signer"`
	PackageObjectId string        `json:"packageObjectId"`
	Module          string        `json:"module"`
	Function        string        `json:"function"`
	TypeArguments   interface{}   `json:"typeArguments"`
	Arguments       []interface{} `json:"arguments"`
	Gas             string        `json:"gas"`
	GasBudget       string        `json:"gasBudget"`
}

type MoveCallResponse struct {
	Gas          []sui_types.SuiObjectRef `json:"gas"`
	InputObjects interface{}              `json:"inputObjects"`
	TxBytes      string                   `json:"txBytes"`
}

type BatchTransactionRequest struct {
	Signer                  string                    `json:"signer"`
	SingleTransactionParams []SingleTransactionParams `json:"singleTransactionParams"`
	Gas                     string                    `json:"gas"`
	GasBudget               uint64                    `json:"gasBudget"`
}

type BatchTransactionResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type SingleTransactionParams struct {
	MoveCallRequestParams       *MoveCallRequest       `json:"moveCallRequestParams,omitempty"`
	TransferObjectRequestParams *TransferObjectRequest `json:"transferObjectRequestParams,omitempty"`
}

type TransferObjectRequest struct {
	Signer    string `json:"signer"`
	ObjectId  string `json:"objectId"`
	Gas       string `json:"gas"`
	GasBudget uint64 `json:"gasBudget"`
	Recipient string `json:"recipient"`
}

type SuiExecuteTransactionBlockRequest struct {
	TxBytes     string                     `json:"txBytes"`
	Signature   []string                   `json:"signature"`
	Options     SuiTransactionBlockOptions `json:"options"`
	RequestType string                     `json:"request_type"`
}

type ExecuteTransactionResponse struct {
	TransactionBytes string                     `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef     `json:"gas"`
	Options          SuiTransactionBlockOptions `json:"options"`
	RequestType      string                     `json:"request_type"`
}
