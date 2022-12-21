package models

import (
	"github.com/block-vision/sui-go-sdk/models/sui_json_rpc_types"
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
	GasBudget       uint64        `json:"gasBudget"`
}

type MoveCallResponse struct {
	Gas          sui_types.SuiObjectRef `json:"gas"`
	InputObjects interface{}            `json:"inputObjects"`
	TxBytes      string                 `json:"txBytes"`
}

type MergeCoinsRequest struct {
	Signer      string `json:"signer"`
	PrimaryCoin string `json:"primaryCoin"`
	CoinToMerge string `json:"coinToMerge"`
	Gas         string `json:"gas"`
	GasBudget   uint64 `json:"gasBudget"`
}

type MergeCoinsResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type SplitCoinRequest struct {
	Signer       string   `json:"signer"`
	CoinObjectId string   `json:"coinObjectId"`
	SplitAmounts []uint64 `json:"splitAmounts"`
	Gas          string   `json:"gas"`
	GasBudget    uint64
}

type SplitCoinResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type SplitCoinEqualRequest struct {
	Signer       string `json:"signer"`
	CoinObjectId string `json:"coinObjectId"`
	SplitCount   uint64 `json:"splitCount"`
	Gas          string `json:"gas"`
	GasBudget    uint64
}

type SplitCoinEqualResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type PublishRequest struct {
	Sender          string   `json:"sender"`
	CompiledModules []string `json:"compiledModules"`
	Gas             string   `json:"gas"`
	GasBudget       uint64   `json:"gasBudget"`
}

type PublishResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type TransferObjectRequest struct {
	Signer    string `json:"signer"`
	ObjectId  string `json:"objectId"`
	Gas       string `json:"gas"`
	GasBudget uint64 `json:"gasBudget"`
	Recipient string `json:"recipient"`
}

type TransferObjectResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type TransferSuiRequest struct {
	Signer      string `json:"signer"`
	SuiObjectId string `json:"suiObjectId"`
	GasBudget   uint64 `json:"gasBudget"`
	Recipient   string `json:"recipient"`
	Amount      uint64 `json:"amount"`
}

type TransferSuiResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
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

type ExecuteTransactionRequest struct {
	TxBytes   string `json:"txBytes"`
	SigScheme string `json:"sigScheme"`
	Signature string `json:"signature"`
	PubKey    string `json:"pubKey"`
}

type ExecuteTransactionResponse struct {
	Certificate sui_json_rpc_types.SuiCertifiedTransaction `json:"certificate"`
	Effects     sui_json_rpc_types.SuiTransactionEffects   `json:"effects"`
	TimestampMs uint64                                     `json:"timestamp_ms"`
	ParsedData  interface{}                                `json:"parsed_data"`
}

type DryRunTransactionRequest struct {
	TxBytes   string `json:"txBytes"`
	SigScheme string `json:"sigScheme"`
	Signature string `json:"signature"`
	PubKey    string `json:"pubKey"`
}

type DryRunTransactionResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type PayRequest struct {
	Signer     string   `json:"signer"`
	InputCoins []string `json:"inputCoins"`
	Recipient  []string `json:"recipient"`
	Amounts    []string `json:"amounts"`
	Gas        string   `json:"gas"`
	GasBudget  uint64   `json:"gasBudget"`
}

type PayResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type PayAllSuiRequest struct {
	Signer     string   `json:"suiAddress,omitempty"`
	InputCoins []string `json:"inputCoins,omitempty"`
	Recipient  string   `json:"recipient,omitempty"`
	GasBudget  uint64   `json:"gasBudget,omitempty"`
}

type PayAllSuiResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type PaySuiRequest struct {
	Signer     string   `json:"signer,omitempty"`
	InputCoins []string `json:"inputCoins,omitempty"`
	Recipient  []string `json:"recipient,omitempty"`
	GasBudget  uint64   `json:"gasBudget,omitempty"`
}

type PaySuiResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type MintNFTRequest struct {
	Signer         string
	NFTName        string
	NFTDescription string
	NFTUrl         string
	GasObject      string
	GasBudget      uint64
}
