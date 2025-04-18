package transaction

import (
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/models/sui_types"
)

// TransactionData https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L1625
type TransactionData struct {
	Sender     *models.SuiAddress
	Expiration TransactionExpiration
	GasData    GasData
	Inputs     []CallArg
	Commands   []Command
}

// GasData https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L1600
type GasData struct {
	Payment []sui_types.SuiObjectRef
	Owner   *models.SuiAddress
	Price   *uint64
	Budget  *uint64
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

// CallArg https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L80
// - Pure
// - Object
type CallArg interface {
	callArgKind() string
}

type Pure struct {
	Value []byte
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
	Value sui_types.SuiSharedObject
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

type InputPure struct {
	Value []byte
}

func (i InputPure) argumentKind() string {
	return "Pure"
}

type InputObject struct {
	Value string
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
