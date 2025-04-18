package transaction

import (
	"bytes"
	"errors"
	"math"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/models/sui_types"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
	"github.com/block-vision/sui-go-sdk/utils"
)

type Transaction struct {
	Data TransactionData
}

func NewTransaction() *Transaction {
	data := TransactionData{}

	return &Transaction{
		Data: data,
	}
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
	if inputObject.Value == nil {
		id = inputObject.ObjectId
		if id == "" {
			return nil, errors.New("object id is empty")
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
			id = objArg.(SharedObject).Value.ObjectId
		case Receiving:
			id = objArg.(Receiving).Value.ObjectId
		default:
			return nil, errors.New("object value is not supported")
		}

		callArg = Object{
			Value: objArg,
		}
	}

	findObj := tx.Data.GetObject(id)
	if findObj != nil {
		return findObj, nil
	}

	input := tx.Data.AddInput(callArg)
	arg := Input{
		Input: input.(Input).Input,
		Type:  input.(Input).Type,
	}

	return arg, nil
}

func (tx *Transaction) Pure(inputPure InputPure) (Argument, error) {
	value := inputPure.Value
	bcsEncodedMsg := bytes.Buffer{}
	bcsEncoder := mystenbcs.NewEncoder(&bcsEncodedMsg)
	err := bcsEncoder.Encode(value)
	if err != nil {
		return nil, err
	}

	bcsBase64 := mystenbcs.ToBase64(bcsEncodedMsg.Bytes())

	input := tx.Data.AddInput(Pure{
		bcsBase64,
	})
	arg := Input{
		Input: input.(Input).Input,
		Type:  input.(Input).Type,
	}

	return arg, nil
}

func createTransactionResult(index uint16, length *uint16) Argument {
	// TODO: Support multiple results
	if length == nil {
		length = math.MaxInt
	}

	return NestedResult{
		Value: NestedResultValue{
			Index:       index,
			ResultIndex: 0,
		},
	}
}
