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

	// simpleTransaction(ctx, client, signerAccount)
	// moveCallTransaction(ctx, client, signerAccount)
	sponsoredTransaction(ctx, client, signerAccount)
}

func simpleTransaction(ctx context.Context, suiClient *sui.Client, signer *signer.Signer) {
	receiver := ""
	gasCoinObjectId := ""

	gasCoinObj, err := suiClient.SuiGetObject(ctx, models.SuiGetObjectRequest{ObjectId: gasCoinObjectId})
	if err != nil {
		panic(err)
	}
	gasCoin, err := transaction.NewSuiObjectRef(
		models.SuiAddress(gasCoinObjectId),
		gasCoinObj.Data.Version,
		models.ObjectDigest(gasCoinObj.Data.Digest),
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

func moveCallTransaction(ctx context.Context, suiClient *sui.Client, signer *signer.Signer) {
	gasCoinObjectId := ""

	gasCoinObj, err := suiClient.SuiGetObject(ctx, models.SuiGetObjectRequest{ObjectId: gasCoinObjectId})
	if err != nil {
		panic(err)
	}
	gasCoin, err := transaction.NewSuiObjectRef(
		models.SuiAddress(gasCoinObjectId),
		gasCoinObj.Data.Version,
		models.ObjectDigest(gasCoinObj.Data.Digest),
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

	addressBytes, err := transaction.ConvertSuiAddressStringToBytes("0x0000000000000000000000000000000000000000000000000000000000000002")
	if err != nil {
		panic(err)
	}

	tx.MoveCall(
		"0xeffc8ae61f439bb34c9b905ff8f29ec56873dcedf81c7123ff2f1f67c45ec302",
		"utils",
		"check_coin_threshold",
		[]transaction.TypeTag{
			{
				Struct: &transaction.StructTag{
					Address: *addressBytes,
					Module:  "sui",
					Name:    "SUI",
				},
			},
		},
		[]transaction.Argument{
			tx.Gas(),
			tx.Pure(uint64(1000000000 * 0.01)),
		},
	)

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

func sponsoredTransaction(ctx context.Context, suiClient *sui.Client, rawSigner *signer.Signer) {
	sponsoredSigner, err := signer.NewSignertWithMnemonic("")
	if err != nil {
		panic(err)
	}
	fmt.Println("Sponsored: ", sponsoredSigner.Address)

	receiver := ""
	transferCoinObjectId := ""
	sponsoredGasCoinObjectId := ""

	// Raw transaction
	tx := transaction.NewTransaction().SetSuiClient(suiClient)

	obj, err := suiClient.SuiGetObject(ctx, models.SuiGetObjectRequest{ObjectId: transferCoinObjectId})
	if err != nil {
		panic(err)
	}
	ref, err := transaction.NewSuiObjectRef(
		models.SuiAddress(obj.Data.ObjectId),
		obj.Data.Version,
		models.ObjectDigest(obj.Data.Digest),
	)
	if err != nil {
		panic(err)
	}

	tx.TransferObjects(
		[]transaction.Argument{
			tx.Object(
				transaction.CallArg{
					Object: &transaction.ObjectArg{
						ImmOrOwnedObject: ref,
					},
				},
			)},
		tx.Pure(receiver),
	)

	// Sponsored transaction
	newTx, err := tx.NewTransactionFromKind()
	if err != nil {
		panic(err)
	}
	newTx.SetSuiClient(suiClient)

	gasCoinObj, err := suiClient.SuiGetObject(ctx, models.SuiGetObjectRequest{ObjectId: sponsoredGasCoinObjectId})
	if err != nil {
		panic(err)
	}
	gasCoin, err := transaction.NewSuiObjectRef(
		models.SuiAddress(sponsoredGasCoinObjectId),
		gasCoinObj.Data.Version,
		models.ObjectDigest(gasCoinObj.Data.Digest),
	)
	if err != nil {
		panic(err)
	}

	newTx.SetSigner(rawSigner).
		SetSponsoredSigner(sponsoredSigner).
		SetSender(models.SuiAddress(rawSigner.Address)).
		SetGasPrice(1000).
		SetGasBudget(50000000).
		SetGasPayment([]transaction.SuiObjectRef{*gasCoin}).
		SetGasOwner(models.SuiAddress(sponsoredSigner.Address))

	resp, err := newTx.Execute(
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
