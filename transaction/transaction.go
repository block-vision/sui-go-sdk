package transaction

import (
	"bytes"
	"encoding/hex"
	"math"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/models/sui_types"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/block-vision/sui-go-sdk/utils"
)

type Transaction struct {
	Data   TransactionDataV1
	Signer *signer.Signer
}

func NewTransaction() *Transaction {
	data := TransactionDataV1{}

	return &Transaction{
		Data: data,
	}
}

func (tx *Transaction) SetSigner(signer *signer.Signer) *Transaction {
	tx.Signer = signer
	return tx
}

func (tx *Transaction) SetSender(sender models.SuiAddress) *Transaction {
	tx.Data.Sender = &sender
	return tx
}

func (tx *Transaction) SetSenderIfNotSet(sender models.SuiAddress) *Transaction {
	if tx.Data.Sender == nil {
		tx.Data.Sender = &sender
	}
	return tx
}

func (tx *Transaction) SetExpiration(expiration TransactionExpiration) *Transaction {
	tx.Data.Expiration = expiration
	return tx
}

func (tx *Transaction) SetGasPayment(payment []sui_types.SuiObjectRef) *Transaction {
	tx.Data.GasData.Payment = payment
	return tx
}

func (tx *Transaction) SetGasOwner(owner models.SuiAddress) *Transaction {
	tx.Data.GasData.Owner = &owner
	return tx
}

func (tx *Transaction) SetGasPrice(price uint64) *Transaction {
	tx.Data.GasData.Price = &price
	return tx
}

func (tx *Transaction) SetGasBudget(budget uint64) *Transaction {
	tx.Data.GasData.Budget = &budget
	return tx
}

func (tx *Transaction) SetGasBudgetIfNotSet(budget uint64) *Transaction {
	if tx.Data.GasData.Budget == nil {
		tx.Data.GasData.Budget = &budget
	}
	return tx
}

func (tx *Transaction) Gas() Argument {
	return GasCoin{
		Value: true,
	}
}

func (tx *Transaction) Add(command Command) Argument {
	index := tx.Data.AddCommand(command)
	return createTransactionResult(index, nil)
}

func (tx *Transaction) SplitCoins(coin Argument, amount []Argument) Argument {
	return tx.Add(splitCoins(SplitCoinsValue{
		Coin:   coin,
		Amount: amount,
	}))
}

func (tx *Transaction) MergeCoins(destination Argument, sources []Argument) Argument {
	return tx.Add(mergeCoins(MergeCoinsValue{
		Destination: destination,
		Sources:     sources,
	}))
}

func (tx *Transaction) Publish(modules []string, dependencies []string) Argument {
	return tx.Add(publish(PublishValue{
		Modules:      modules,
		Dependencies: dependencies,
	}))
}

func (tx *Transaction) Upgrade(
	modules []string,
	dependencies []string,
	packageId string,
	ticket Argument,
) Argument {
	return tx.Add(upgrade(UpgradeValue{
		Modules:      modules,
		Dependencies: dependencies,
		Package:      packageId,
		Ticket:       ticket,
	}))
}

func (tx *Transaction) MoveCall(
	packageId string,
	module string,
	function string,
	typeArguments []string,
	arguments []Argument,
) Argument {
	return tx.Add(moveCall(ProgrammableMoveCall{
		Package:       packageId,
		Module:        module,
		Function:      function,
		TypeArguments: typeArguments,
		Arguments:     arguments,
	}))
}

func (tx *Transaction) transferObjects(objects []Argument, address Argument) Argument {
	return tx.Add(transferObjects(TransferObjectsValue{
		Objects: objects,
		Address: address,
	}))
}

func (tx *Transaction) makeMoveVec(typeValue *string, elements []Argument) Argument {
	return tx.Add(makeMoveVec(MakeMoveVecValue{
		Type:     typeValue,
		Elements: elements,
	}))
}

func (tx *Transaction) Object(inputObject InputObject) (Argument, error) {
	var callArg CallArg

	var id string
	isSharedObjectAndSetMutable := false
	if inputObject.Value == nil {
		id = inputObject.ObjectId
		if id == "" {
			return nil, ErrObjectIdNotSet
		}

		callArg = UnresolvedObject{
			ObjectId: utils.NormalizeSuiAddress(id),
		}
	} else {
		objArg := inputObject.Value.Value
		switch objArg.(type) {
		case ImmOrOwnedObject:
			id = objArg.(ImmOrOwnedObject).Value.ObjectId
		case SharedObject:
			val := objArg.(SharedObject).Value
			isSharedObjectAndSetMutable = val.Mutable
			id = val.ObjectId
		case Receiving:
			id = objArg.(Receiving).Value.ObjectId
		default:
			return nil, ErrObjectTypeNotSupported
		}

		callArg = Object{
			Value: objArg,
		}
	}

	findObj := tx.Data.GetInputObject(id)
	if findObj != nil {
		if isSharedObjectAndSetMutable {
			index := findObj.(Input).Input
			existedInput := tx.Data.TransactionKind.ProgrammableTransaction.Inputs[index]
			if obj, ok := existedInput.(Object); ok {
				if objArg, ok := obj.Value.(SharedObject); ok {
					newObjArg := objArg
					newObjArg.Value.Mutable = true
					tx.Data.TransactionKind.ProgrammableTransaction.Inputs[index] = Object{
						Value: newObjArg,
					}
				}
			}
		}

		return findObj, nil
	}

	input := tx.Data.AddInput(callArg, "Object")
	arg := Input{
		Input: input.(Input).Input,
		Type:  input.(Input).Type,
	}

	return arg, nil
}

func (tx *Transaction) Pure(inputPure InputPure) (Argument, error) {
	val := inputPure.Value

	if s, ok := val.(string); ok && utils.IsValidSuiAddress(s) {
		normalized := utils.NormalizeSuiAddress(s)
		vBytes, err := hex.DecodeString(normalized[2:])
		if err != nil {
			return nil, err
		}
		if len(vBytes) != 32 {
			return nil, ErrInvalidSuiAddress
		}

		var fixedBytes [32]byte
		copy(fixedBytes[:], vBytes)
		val = fixedBytes
	}

	bcsEncodedMsg := bytes.Buffer{}
	bcsEncoder := mystenbcs.NewEncoder(&bcsEncodedMsg)
	err := bcsEncoder.Encode(val)
	if err != nil {
		return nil, err
	}

	bcsBase64 := mystenbcs.ToBase64(bcsEncodedMsg.Bytes())

	input := tx.Data.AddInput(Pure{bcsBase64}, "Pure")
	arg := Input{
		Input: input.(Input).Input,
		Type:  input.(Input).Type,
	}

	return arg, nil
}

func (tx *Transaction) ToSuiExecuteTransactionBlockRequest(
	options models.SuiTransactionBlockOptions,
	requestType string,
) (*models.SuiExecuteTransactionBlockRequest, error) {
	if tx.Signer == nil {
		return nil, ErrSignerNotSet
	}

	txBytes, err := tx.buildTransaction()
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

func (tx *Transaction) buildTransaction() (string, error) {
	if tx.Signer == nil {
		return "", ErrSignerNotSet
	}

	tx.SetGasBudgetIfNotSet(defaultGasBudget)
	tx.SetSenderIfNotSet(models.SuiAddress(tx.Signer.Address))

	return tx.build(false)
}

func (tx *Transaction) build(onlyTransactionKind bool) (string, error) {
	if onlyTransactionKind {
		bcsEncodedMsg, err := tx.Data.TransactionKind.MarshalBCS()
		if err != nil {
			return "", err
		}
		bcsBase64 := mystenbcs.ToBase64(bcsEncodedMsg)
		return bcsBase64, nil
	}

	if tx.Data.Sender == nil {
		return "", ErrSenderNotSet
	}
	// TODO: Support get latest gas data online
	if !tx.Data.GasData.IsFullySet() {
		return "", ErrGasDataNotFullySet
	}

	transactionData := TransactionData{
		V1: tx.Data,
	}
	bcsEncodedMsg, err := transactionData.MarshalBCS()
	if err != nil {
		return "", err
	}
	bcsBase64 := mystenbcs.ToBase64(bcsEncodedMsg)

	return bcsBase64, nil
}

func createTransactionResult(index uint16, length *uint16) Argument {
	if length == nil {
		m := uint16(math.MaxUint16)
		length = &m
	}

	// TODO: Support NestedResult
	return Result{
		Value: index,
	}
}
