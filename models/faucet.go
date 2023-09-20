package models

type FaucetFixedAmountRequest struct {
	Recipient string `json:"recipient"`
}

type FaucetRequest struct {
	FixedAmountRequest *FaucetFixedAmountRequest `json:"FixedAmountRequest"`
}

type FaucetCoinInfo struct {
	Amount           int    `json:"amount"`
	ID               string `json:"id"`
	TransferTxDigest string `json:"transferTxDigest"`
}
