package main

import (
	"context"
	"fmt"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/transaction"
)

func main() {
	var ctx = context.Background()
	var cli = sui.NewSuiClient(constant.SuiMainnetEndpoint)
	client, ok := cli.(*sui.Client)
	if !ok {
		panic("not sui client")
	}
	signerAccount, err := signer.NewSignertWithMnemonic("")
	if err != nil {
		panic(err)
	}
	fmt.Println(signerAccount.Address)

	simpleTransfer(ctx, client, signerAccount)
}

func simpleTransfer(ctx context.Context, suiClient *sui.Client, signer *signer.Signer) {
	receiver := ""

	gasCoinObjectId := ""
	gasCoinVersion := uint64(0)
	gasCoinDigest := ""
	gasCoin, err := transaction.NewSuiObjectRef(
		models.SuiAddress(gasCoinObjectId),
		gasCoinVersion,
		models.ObjectDigest(gasCoinDigest),
	)
	if err != nil {
		panic(err)
	}

	tx := transaction.NewTransaction()

	tx.SetSuiClient(suiClient).
		SetSigner(signer).
		SetSender(models.SuiAddress(signer.Address)).
		SetGasPrice(1000).
		SetGasBudget(50000000).
		SetGasPayment([]transaction.SuiObjectRef{*gasCoin}).
		SetGasOwner(models.SuiAddress(signer.Address))

	splitCoin := tx.SplitCoins(tx.Gas(), []transaction.Argument{
		tx.Pure(uint64(1000000000 * 0.01)),
	})
	tx.TransferObjects([]transaction.Argument{splitCoin}, tx.Pure(receiver))

	resp, err := tx.Execute(
		ctx,
		models.SuiTransactionBlockOptions{
			ShowInput:    true,
			ShowRawInput: true,
			ShowEffects:  true,
			ShowEvents:   true,
		},
		"WaitForLocalExecution",
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Digest, resp.Effects, resp.Results)
}
