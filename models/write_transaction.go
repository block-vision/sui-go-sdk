package models

import (
	"crypto/ed25519"
	"github.com/block-vision/sui-go-sdk/models/sui_types"
)

type MoveCallRequest struct {
	// the transaction signer's Sui address
	Signer string `json:"signer"`
	// the package containing the module and function
	PackageObjectId string `json:"packageObjectId"`
	// the specific module in the package containing the function
	Module string `json:"module"`
	// the function to be called
	Function string `json:"function"`
	// the type arguments to the function
	TypeArguments []interface{} `json:"typeArguments"`
	// the arguments to the function
	Arguments []interface{} `json:"arguments"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
}

type MoveCallResponse struct {
	Gas          []sui_types.SuiObjectRef `json:"gas"`
	InputObjects interface{}              `json:"inputObjects"`
	TxBytes      string                   `json:"txBytes"`
}

type MergeCoinsRequest struct {
	// the transaction signer's Sui address
	Signer      string `json:"signer"`
	PrimaryCoin string `json:"primaryCoin"`
	CoinToMerge string `json:"coinToMerge"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
}

type SplitCoinRequest struct {
	// the transaction signer's Sui address
	Signer       string   `json:"signer"`
	CoinObjectId string   `json:"coinObjectId"`
	SplitAmounts []string `json:"splitAmounts"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
}

type SplitCoinEqualRequest struct {
	// the transaction signer's Sui address
	Signer       string `json:"signer"`
	CoinObjectId string `json:"coinObjectId"`
	SplitCount   string `json:"splitCount"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
}

type PublishRequest struct {
	// the transaction signer's Sui address
	Sender          string   `json:"sender"`
	CompiledModules []string `json:"compiled_modules"`
	Dependencies    []string `json:"dependencies"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
}

type TransferObjectRequest struct {
	// the transaction signer's Sui address
	Signer   string `json:"signer"`
	ObjectId string `json:"objectId"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
	Recipient string `json:"recipient"`
}

type TransferSuiRequest struct {
	// the transaction signer's Sui address
	Signer      string `json:"signer"`
	SuiObjectId string `json:"suiObjectId"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

type PayRequest struct {
	// the transaction signer's Sui address
	Signer      string   `json:"signer"`
	SuiObjectId []string `json:"suiObjectId"`
	Recipient   []string `json:"recipient"`
	Amount      []string `json:"amount"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
}

type PaySuiRequest struct {
	// the transaction signer's Sui address
	Signer      string   `json:"signer"`
	SuiObjectId []string `json:"suiObjectId"`
	Recipient   []string `json:"recipient"`
	Amount      []string `json:"amount"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
}

type PayAllSuiRequest struct {
	// the transaction signer's Sui address
	Signer      string   `json:"signer"`
	SuiObjectId []string `json:"suiObjectId"`
	Recipient   string   `json:"recipient"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
}

type AddStakeRequest struct {
	// the transaction signer's Sui address
	Signer string `json:"signer"`
	// Coin<SUI> object to stake
	Coins []string `json:"coins"`
	// stake amount
	Amount string `json:"amount"`
	// the validator's Sui address
	Validator string `json:"validator"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
}

type WithdrawStakeRequest struct {
	// the transaction signer's Sui address
	Signer string `json:"signer"`
	// StakedSui object ID
	StakedObjectId string `json:"stakedObjectId"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
}

type TxnMetaData struct {
	Gas          []sui_types.SuiObjectRef `json:"gas"`
	InputObjects []interface{}            `json:"inputObjects"`
	TxBytes      string                   `json:"txBytes"`
}

type BatchTransactionRequest struct {
	// the transaction signer's Sui address
	Signer string `json:"signer"`
	// list of transaction request parameters
	RPCTransactionRequestParams []RPCTransactionRequestParams `json:"RPCTransactionRequestParams"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
	// Whether this is a regular transaction or a Dev Inspect Transaction
	// The optional enumeration values are: `DevInspect`, or `Commit`
	SuiTransactionBlockBuilderMode string `json:"suiTransactionBlockBuilderMode"`
}

type BatchTransactionResponse struct {
	Gas          []sui_types.SuiObjectRef `json:"gas"`
	InputObjects []interface{}            `json:"inputObjects"`
	TxBytes      string                   `json:"txBytes"`
}

type RPCTransactionRequestParams struct {
	MoveCallRequestParams       *MoveCallRequest       `json:"moveCallRequestParams,omitempty"`
	TransferObjectRequestParams *TransferObjectRequest `json:"transferObjectRequestParams,omitempty"`
}

type SuiExecuteTransactionBlockRequest struct {
	// BCS serialized transaction data bytes without its type tag, as base-64 encoded string.
	TxBytes string `json:"txBytes"`
	// A list of signatures (`flag || signature || pubkey` bytes, as base-64 encoded string).
	// Signature is committed to the intent message of the transaction data, as base-64 encoded string.
	Signature []string `json:"signature"`
	// Options for specifying the content to be returned
	Options SuiTransactionBlockOptions `json:"options"`
	// The request type, derived from `SuiTransactionBlockResponseOptions` if None.
	// The optional enumeration values are: `WaitForEffectsCert`, or `WaitForLocalExecution`
	RequestType string `json:"requestType"`
}

type SignAndExecuteTransactionBlockRequest struct {
	TxnMetaData TxnMetaData
	// the address private key to sign the transaction
	PriKey  ed25519.PrivateKey
	Options SuiTransactionBlockOptions `json:"options"`
	// The optional enumeration values are: `WaitForEffectsCert`, or `WaitForLocalExecution`
	RequestType string `json:"requestType"`
}
