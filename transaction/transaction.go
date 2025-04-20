package transaction

import (
	"bytes"
	"context"
	"math"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
	"github.com/samber/lo"
)

type Transaction struct {
	Data      TransactionData
	Signer    *signer.Signer
	SuiClient *sui.Client
}

func NewTransaction() *Transaction {
	data := TransactionData{}
	data.V1.Kind = &TransactionKind{
		ProgrammableTransaction: &ProgrammableTransaction{},
	}

	return &Transaction{
		Data: data,
	}
}

func (tx *Transaction) SetSigner(signer *signer.Signer) *Transaction {
	tx.Signer = signer

	return tx
}

func (tx *Transaction) SetSuiClient(client *sui.Client) *Transaction {
	tx.SuiClient = client

	return tx
}

func (tx *Transaction) SetSender(sender models.SuiAddress) *Transaction {
	address := utils.NormalizeSuiAddress(string(sender))
	addressBytes, err := ConvertSuiAddressStringToBytes(address)
	if err != nil {
		panic(err)
	}
	tx.Data.V1.Sender = addressBytes

	return tx
}

func (tx *Transaction) SetSenderIfNotSet(sender models.SuiAddress) *Transaction {
	if tx.Data.V1.Sender.IsZero() {
		tx.SetSender(sender)
	}

	return tx
}

func (tx *Transaction) SetExpiration(expiration TransactionExpiration) *Transaction {
	tx.Data.V1.Expiration = &expiration

	return tx
}

func (tx *Transaction) SetGasPayment(payment []SuiObjectRef) *Transaction {
	tx.Data.V1.GasData.Payment = payment

	return tx
}

func (tx *Transaction) SetGasOwner(owner models.SuiAddress) *Transaction {
	addressBytes, err := ConvertSuiAddressStringToBytes(owner)
	if err != nil {
		panic(err)
	}
	tx.Data.V1.GasData.Owner = addressBytes

	return tx
}

func (tx *Transaction) SetGasPrice(price uint64) *Transaction {
	tx.Data.V1.GasData.Price = &price

	return tx
}

func (tx *Transaction) SetGasBudget(budget uint64) *Transaction {
	tx.Data.V1.GasData.Budget = &budget

	return tx
}

func (tx *Transaction) SetGasBudgetIfNotSet(budget uint64) *Transaction {
	if tx.Data.V1.GasData.Budget == nil {
		tx.Data.V1.GasData.Budget = &budget
	}

	return tx
}

func (tx *Transaction) Gas() Argument {
	return Argument{
		GasCoin: lo.ToPtr(true),
	}
}

func (tx *Transaction) Add(command Command) Argument {
	index := tx.Data.V1.AddCommand(command)
	return createTransactionResult(index, nil)
}

func (tx *Transaction) SplitCoins(coin Argument, amount []Argument) Argument {
	return tx.Add(splitCoins(SplitCoins{
		Coin:   &coin,
		Amount: convertArgumentsToArgumentPtrs(amount),
	}))
}

func (tx *Transaction) MergeCoins(destination Argument, sources []Argument) Argument {
	return tx.Add(mergeCoins(MergeCoins{
		Destination: &destination,
		Sources:     convertArgumentsToArgumentPtrs(sources),
	}))
}

func (tx *Transaction) Publish(modules []models.SuiAddress, dependencies []models.SuiAddress) Argument {
	moduleAddress := make([]models.SuiAddressBytes, len(modules))
	for i, module := range modules {
		v, err := ConvertSuiAddressStringToBytes(module)
		if err != nil {
			panic(err)
		}
		moduleAddress[i] = *v
	}

	dependenciesAddress := make([]models.SuiAddressBytes, len(dependencies))
	for i, dependency := range dependencies {
		v, err := ConvertSuiAddressStringToBytes(dependency)
		if err != nil {
			panic(err)
		}
		dependenciesAddress[i] = *v
	}

	return tx.Add(publish(Publish{
		Modules:      moduleAddress,
		Dependencies: dependenciesAddress,
	}))
}

func (tx *Transaction) Upgrade(
	modules []models.SuiAddress,
	dependencies []models.SuiAddress,
	packageId models.SuiAddress,
	ticket Argument,
) Argument {
	moduleAddress := make([]models.SuiAddressBytes, len(modules))
	for i, module := range modules {
		v, err := ConvertSuiAddressStringToBytes(module)
		if err != nil {
			panic(err)
		}
		moduleAddress[i] = *v
	}

	dependenciesAddress := make([]models.SuiAddressBytes, len(dependencies))
	for i, dependency := range dependencies {
		v, err := ConvertSuiAddressStringToBytes(dependency)
		if err != nil {
			panic(err)
		}
		dependenciesAddress[i] = *v
	}

	packageIdBytes, err := ConvertSuiAddressStringToBytes(packageId)
	if err != nil {
		panic(err)
	}

	return tx.Add(upgrade(Upgrade{
		Modules:      moduleAddress,
		Dependencies: dependenciesAddress,
		Package:      *packageIdBytes,
		Ticket:       &ticket,
	}))
}

func (tx *Transaction) MoveCall(
	packageId models.SuiAddress,
	module string,
	function string,
	typeArguments []string,
	arguments []Argument,
) Argument {
	packageIdBytes, err := ConvertSuiAddressStringToBytes(packageId)
	if err != nil {
		panic(err)
	}

	return tx.Add(moveCall(ProgrammableMoveCall{
		Package:       *packageIdBytes,
		Module:        module,
		Function:      function,
		TypeArguments: typeArguments,
		Arguments:     convertArgumentsToArgumentPtrs(arguments),
	}))
}

func (tx *Transaction) transferObjects(objects []Argument, address Argument) Argument {
	return tx.Add(transferObjects(TransferObjects{
		Objects: convertArgumentsToArgumentPtrs(objects),
		Address: &address,
	}))
}

func (tx *Transaction) makeMoveVec(typeValue *string, elements []Argument) Argument {
	return tx.Add(makeMoveVec(MakeMoveVec{
		Type:     typeValue,
		Elements: convertArgumentsToArgumentPtrs(elements),
	}))
}

// Object
// - input: string | CallArg | Argument
func (tx *Transaction) Object(input any) *Argument {
	// string
	if s, ok := input.(string); ok {
		if utils.IsValidSuiAddress(models.SuiAddress(s)) {
			address := utils.NormalizeSuiAddress(s)
			addressBytes, err := ConvertSuiAddressStringToBytes(address)
			if err != nil {
				panic(err)
			}

			arg := tx.Data.V1.AddInput(CallArg{
				UnresolvedObject: &UnresolvedObject{
					ObjectId: *addressBytes,
				},
			})

			return &arg
		} else {
			return nil
		}
	}

	// Argument
	if arg, ok := input.(Argument); ok {
		return &arg
	}

	// CallArg
	if v, ok := input.(CallArg); ok {
		if id := v.Object.SharedObject.ObjectId; !id.IsZero() {
			// SharedObject
			address := ConvertSuiAddressBytesToString(id)
			if index := tx.Data.V1.GetInputObjectIndex(address); index != nil {
				// Already exists
				if v.Object.SharedObject.Mutable {
					newExistObject := tx.Data.V1.Kind.ProgrammableTransaction.Inputs[*index]
					if !newExistObject.Object.SharedObject.ObjectId.IsZero() {
						newExistObject.Object.SharedObject.Mutable = true
					}
					tx.Data.V1.Kind.ProgrammableTransaction.Inputs[*index] = newExistObject
				}
			}
		} else {
			if id := v.Object.ImmOrOwnedObject.ObjectId; !id.IsZero() {
				// ImmOrOwnedObject
				arg := tx.Data.V1.AddInput(CallArg{
					Object: v.Object,
				})
				return &arg
			} else if id := v.Object.Receiving.ObjectId; !id.IsZero() {
				// Receiving
				arg := tx.Data.V1.AddInput(CallArg{
					Object: v.Object,
				})
				return &arg
			} else {
				// Not supported
				return nil
			}
		}
	}

	return nil
}

func (tx *Transaction) Pure(input any) *Argument {
	val := input
	if s, ok := input.(string); ok && utils.IsValidSuiAddress(models.SuiAddress(s)) {
		bcsAddress, err := ConvertSuiAddressStringToBytes(models.SuiAddress(s))
		if err != nil {
			panic(err)
		}
		val = *bcsAddress
	}

	bcsEncodedMsg := bytes.Buffer{}
	bcsEncoder := mystenbcs.NewEncoder(&bcsEncodedMsg)
	err := bcsEncoder.Encode(val)
	if err != nil {
		tx.Data.V1.AddInput(CallArg{UnresolvedPure: lo.ToPtr(bcsEncodedMsg.Bytes())})
	}

	arg := tx.Data.V1.AddInput(CallArg{Pure: &Pure{
		Bytes: bcsEncodedMsg.Bytes(),
	}})

	return &arg
}

func (tx *Transaction) ToSuiExecuteTransactionBlockRequest(
	ctx context.Context,
	options models.SuiTransactionBlockOptions,
	requestType string,
) (*models.SuiExecuteTransactionBlockRequest, error) {
	if tx.Signer == nil {
		return nil, ErrSignerNotSet
	}

	txBytes, err := tx.buildTransaction(ctx)
	if err != nil {
		return nil, err
	}

	signedTransaction, err := tx.Signer.SignTransaction(txBytes)
	if err != nil {
		return nil, err
	}

	return &models.SuiExecuteTransactionBlockRequest{
		TxBytes:     signedTransaction.TxBytes,
		Signature:   []string{signedTransaction.Signature},
		Options:     options,
		RequestType: requestType,
	}, nil
}

func (tx *Transaction) buildTransaction(ctx context.Context) (string, error) {
	if tx.Signer == nil {
		return "", ErrSignerNotSet
	}

	if tx.Data.V1.GasData.Price == nil {
		if tx.SuiClient != nil {
			rsp, err := tx.SuiClient.SuiXGetReferenceGasPrice(ctx)
			if err != nil {
				return "", err
			}
			tx.SetGasPrice(rsp)
		}
	}
	tx.SetGasBudgetIfNotSet(defaultGasBudget)
	tx.SetSenderIfNotSet(models.SuiAddress(tx.Signer.Address))

	return tx.build(false)
}

func (tx *Transaction) build(onlyTransactionKind bool) (string, error) {
	if onlyTransactionKind {
		bcsEncodedMsg, err := tx.Data.V1.Kind.Marshal()
		if err != nil {
			return "", err
		}
		bcsBase64 := mystenbcs.ToBase64(bcsEncodedMsg)
		return bcsBase64, nil
	}

	if tx.Data.V1.Sender.IsZero() {
		return "", ErrSenderNotSet
	}
	if !tx.Data.V1.GasData.IsFullySet() {
		return "", ErrGasDataNotFullySet
	}

	bcsEncodedMsg, err := tx.Data.Marshal()
	if err != nil {
		return "", err
	}
	bcsBase64 := mystenbcs.ToBase64(bcsEncodedMsg)

	return bcsBase64, nil
}

func createTransactionResult(index uint16, length *uint16) Argument {
	if length == nil {
		length = lo.ToPtr(uint16(math.MaxUint16))
	}

	// TODO: Support NestedResult
	return Argument{
		Result: lo.ToPtr(index),
	}
}

func convertArgumentsToArgumentPtrs(args []Argument) []*Argument {
	argPtrs := make([]*Argument, len(args))
	for i, arg := range args {
		argPtrs[i] = &arg
	}

	return argPtrs
}
