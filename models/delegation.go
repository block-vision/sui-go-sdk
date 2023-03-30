package models

import "github.com/shoshinsquare/sui-go-sdk/models/sui_types"

type RequestAddDelegationRequest struct {
	Signer    string
	Coins     []string
	Amount    uint64
	Validator string
	Gas       string
	GasBudget uint64
}

type RequestAddDelegationResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type RequestWithdrawDelegationRequest struct {
	Signer                  string
	Delegation              string
	StakedSui               string
	PrincipalWithdrawAmount uint64
	Gas                     string
	GasBudget               uint64
}

type ReuqestWithdrawDelegationResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type RequestSwitchDelegationRequest struct {
	Signer                string
	Delegation            string
	StakedSui             string
	NewValidatorAddress   string
	SwitchPoolTokenAmount uint64
	Gas                   string
	GasBudget             uint64
}

type RequestSwitchDelegationResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}
