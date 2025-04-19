package transaction

import (
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/models/sui_types"
)

// TransactionDataV1 https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L1625
type TransactionDataV1 struct {
	Sender     *models.SuiAddress
	Expiration TransactionExpiration
	GasData    GasData
	TransactionKind
}

func (td *TransactionDataV1) AddCommand(command Command) (index uint16) {
	index = uint16(len(td.TransactionKind.ProgrammableTransaction.Commands))
	td.TransactionKind.ProgrammableTransaction.Commands = append(td.TransactionKind.ProgrammableTransaction.Commands, command)

	return index
}

func (td *TransactionDataV1) AddInput(input CallArg, inputType string) Argument {
	index := len(td.TransactionKind.ProgrammableTransaction.Inputs)
	td.TransactionKind.ProgrammableTransaction.Inputs = append(td.TransactionKind.ProgrammableTransaction.Inputs, input)

	return Input{
		Input: uint16(index),
		Type:  inputType,
	}
}

func (td *TransactionDataV1) GetInputObject(objectId string) Argument {
	for i, input := range td.TransactionKind.ProgrammableTransaction.Inputs {
		var inputId string

		switch input.(type) {
		case Object:
			obj := input.(Object).Value
			switch obj.(type) {
			case ImmOrOwnedObject:
				inputId = obj.(ImmOrOwnedObject).Value.ObjectId
			case SharedObject:
				inputId = obj.(SharedObject).Value.ObjectId
			case Receiving:
				inputId = obj.(Receiving).Value.ObjectId
			default:
				panic("object value is not supported")
			}
		case UnresolvedObject:
			inputId = input.(UnresolvedObject).ObjectId
		default:
			continue
		}

		if inputId == objectId {
			return Input{
				Input: uint16(i),
				Type:  input.callArgKind(),
			}
		}
	}

	return nil
}

type TransactionData struct {
	V1 TransactionDataV1
}

func (td *TransactionData) MarshalBCS() ([]byte, error) {
	// TODO
	return []byte{}, nil
}

// GasData https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L1600
type GasData struct {
	Payment []sui_types.SuiObjectRef
	Owner   *models.SuiAddress
	Price   *uint64
	Budget  *uint64
}

func (g *GasData) IsFullySet() bool {
	return len(g.Payment) > 0 && g.Owner != nil && g.Price != nil && g.Budget != nil
}

// TransactionExpiration https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L1608
type TransactionExpiration struct {
	Epoch *uint64
}

// ProgrammableTransaction https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L702
type ProgrammableTransaction struct {
	Inputs   []CallArg
	Commands []Command
}

type TransactionKind struct {
	ProgrammableTransaction ProgrammableTransaction
	// ChangeEpoch
	// Genesis
	// ConsensusCommitPrologue
}

func (tk *TransactionKind) MarshalBCS() ([]byte, error) {
	// TODO
	return []byte{}, nil
}

// CallArg https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L80
// - Pure
// - Object
// - UnresolvedPure
// - UnresolvedObject
type CallArg interface {
	callArgKind() string
}

type Pure struct {
	// BCSBates's Base64
	Bytes string
}

func (p Pure) callArgKind() string {
	return "Pure"
}

type Object struct {
	Value ObjectArg
}

func (o Object) callArgKind() string {
	return "Object"
}

type UnresolvedPure struct {
	Value string
}

func (u UnresolvedPure) callArgKind() string {
	return "UnresolvedPure"
}

type UnresolvedObject struct {
	ObjectId string
}

func (u UnresolvedObject) callArgKind() string {
	return "UnresolvedObject"
}

// ObjectArg
// - ImmOrOwnedObject
// - SharedObject
// - Receiving
type ObjectArg interface {
	objectArgKind() string
}

type ImmOrOwnedObject struct {
	Value sui_types.SuiObjectRef
}

func (i ImmOrOwnedObject) objectArgKind() string {
	return "ImmOrOwnedObject"
}

type SharedObject struct {
	Value sui_types.SharedObject
}

func (s SharedObject) objectArgKind() string {
	return "SharedObject"
}

type Receiving struct {
	Value sui_types.SuiObjectRef
}

func (r Receiving) objectArgKind() string {
	return "Receiving"
}

// Command https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L712
// - MoveCall
// - TransferObjects
// - SplitCoins
// - MergeCoins
// - Publish
// - MakeMoveVec
// - Upgrade
type Command interface {
	commandKind() string
}

type MoveCall struct {
	Value ProgrammableMoveCall
}

func (m MoveCall) commandKind() string {
	return "MoveCall"
}

type TransferObjects struct {
	Value TransferObjectsValue
}

func (t TransferObjects) commandKind() string {
	return "TransferObjects"
}

type SplitCoins struct {
	Value SplitCoinsValue
}

func (s SplitCoins) commandKind() string {
	return "SplitCoins"
}

type MergeCoins struct {
	Value MergeCoinsValue
}

func (m MergeCoins) commandKind() string {
	return "MergeCoins"
}

type Publish struct {
	Value PublishValue
}

func (p Publish) commandKind() string {
	return "Publish"
}

type MakeMoveVec struct {
	Value MakeMoveVecValue
}

func (m MakeMoveVec) commandKind() string {
	return "MakeMoveVec"
}

type Upgrade struct {
	Value UpgradeValue
}

func (u Upgrade) commandKind() string {
	return "Upgrade"
}

// ProgrammableMoveCall https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L762
type ProgrammableMoveCall struct {
	Package       string
	Module        string
	Function      string
	TypeArguments []string
	Arguments     []Argument
}

type TransferObjectsValue struct {
	Objects []Argument
	Address Argument
}

type SplitCoinsValue struct {
	Coin   Argument
	Amount []Argument
}

type MergeCoinsValue struct {
	Destination Argument
	Sources     []Argument
}

type PublishValue struct {
	Modules      []string
	Dependencies []string
}

type MakeMoveVecValue struct {
	Type     *string
	Elements []Argument
}

type UpgradeValue struct {
	Modules      []string
	Dependencies []string
	Package      string
	Ticket       Argument
}

// Argument https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L745
// - GasCoin
// - Input
// - InputPure
// - InputObject
// - Result
// - NestedResult
type Argument interface {
	argumentKind() string
}

type GasCoin struct {
	Value bool
}

func (g GasCoin) argumentKind() string {
	return "GasCoin"
}

type Input struct {
	// Index
	Input uint16
	Type  string
}

func (i Input) argumentKind() string {
	return "Input"
}

type InputPure struct {
	Value any
}

func (i InputPure) argumentKind() string {
	return "Pure"
}

type InputObject struct {
	ObjectId string
	Value    *Object
}

func (i InputObject) argumentKind() string {
	return "Object"
}

type Result struct {
	Value uint16
}

func (r Result) argumentKind() string {
	return "Result"
}

type NestedResult struct {
	Value NestedResultValue
}

func (n NestedResult) argumentKind() string {
	return "NestedResult"
}

type NestedResultValue struct {
	Index       uint16
	ResultIndex uint16
}
