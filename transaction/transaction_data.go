package transaction

import (
	"bytes"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
)

type TransactionData struct {
	V1 TransactionDataV1
}

func (td *TransactionData) Marshal() ([]byte, error) {
	bcsEncodedMsg := bytes.Buffer{}
	bcsEncoder := mystenbcs.NewEncoder(&bcsEncodedMsg)
	err := bcsEncoder.Encode(td)
	if err != nil {
		return nil, err
	}

	return bcsEncodedMsg.Bytes(), nil
}

// TransactionDataV1 https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L1625
type TransactionDataV1 struct {
	Kind       *TransactionKind
	Sender     models.SuiAddressBytes
	GasData    GasData
	Expiration *TransactionExpiration
}

func (td *TransactionDataV1) AddCommand(command Command) (index uint16) {
	index = uint16(len(td.Kind.ProgrammableTransaction.Commands))
	td.Kind.ProgrammableTransaction.Commands = append(td.Kind.ProgrammableTransaction.Commands, &command)

	return index
}

func (td *TransactionDataV1) AddInput(input CallArg) Argument {
	index := uint16(len(td.Kind.ProgrammableTransaction.Inputs))
	td.Kind.ProgrammableTransaction.Inputs = append(td.Kind.ProgrammableTransaction.Inputs, &input)

	return Argument{
		Input: &index,
	}
}

func (td *TransactionDataV1) GetInputObjectIndex(address models.SuiAddress) *uint16 {
	addressBytes, err := ConvertSuiAddressStringToBytes(address)
	if err != nil {
		return nil
	}

	for i, input := range td.Kind.ProgrammableTransaction.Inputs {
		if !input.Object.ImmOrOwnedObject.ObjectId.IsZero() {
			objectId := input.Object.ImmOrOwnedObject.ObjectId
			if objectId.IsEqual(*addressBytes) {
				index := uint16(i)
				return &index
			}
		}
		if !input.Object.SharedObject.ObjectId.IsZero() {
			objectId := input.Object.SharedObject.ObjectId
			if objectId.IsEqual(*addressBytes) {
				index := uint16(i)
				return &index
			}
		}
		if !input.Object.Receiving.ObjectId.IsZero() {
			objectId := input.Object.Receiving.ObjectId
			if objectId.IsEqual(*addressBytes) {
				index := uint16(i)
				return &index
			}
		}
	}

	return nil
}

// GasData https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L1600
type GasData struct {
	Payment []SuiObjectRef
	Owner   *models.SuiAddressBytes
	Price   *uint64
	Budget  *uint64
}

func (gd *GasData) IsFullySet() bool {
	if len(gd.Payment) == 0 {
		return false
	}
	if gd.Owner == nil || gd.Price == nil || gd.Budget == nil {
		return false
	}

	return true
}

// TransactionExpiration https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L1608
// - None
// - Epoch
type TransactionExpiration struct {
	None  *bool
	Epoch *uint64
}

func (*TransactionExpiration) IsBcsEnum() {}

// ProgrammableTransaction https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L702
type ProgrammableTransaction struct {
	Inputs   []*CallArg
	Commands []*Command
}

// TransactionKind https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L303
// - ProgrammableTransaction
// - ChangeEpoch
// - Genesis
// - ConsensusCommitPrologue
type TransactionKind struct {
	ProgrammableTransaction *ProgrammableTransaction
	ChangeEpoch             *bool
	Genesis                 *bool
	ConsensusCommitPrologue *bool
}

func (*TransactionKind) IsBcsEnum() {}

func (tk *TransactionKind) Marshal() ([]byte, error) {
	bcsEncodedMsg := bytes.Buffer{}
	bcsEncoder := mystenbcs.NewEncoder(&bcsEncodedMsg)
	err := bcsEncoder.Encode(tk)
	if err != nil {
		return nil, err
	}

	return bcsEncodedMsg.Bytes(), nil
}

// CallArg https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L80
// - Pure
// - Object
// - UnresolvedPure
// - UnresolvedObject
type CallArg struct {
	Pure             *Pure
	Object           *ObjectArg
	UnresolvedPure   any
	UnresolvedObject *UnresolvedObject
}

type Pure struct {
	Bytes []byte
}

func (*CallArg) IsBcsEnum() {}

type UnresolvedObject struct {
	ObjectId models.SuiAddressBytes
	// Version
	// Digest
	// InitialSharedVersion
}

// ObjectArg
// - ImmOrOwnedObject
// - SharedObject
// - Receiving
type ObjectArg struct {
	ImmOrOwnedObject *SuiObjectRef
	SharedObject     *SharedObjectRef
	Receiving        *SuiObjectRef
}

func (*ObjectArg) IsBcsEnum() {}

// Command https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L712
// - MoveCall
// - TransferObjects
// - SplitCoins
// - MergeCoins
// - Publish
// - MakeMoveVec
// - Upgrade
type Command struct {
	MoveCall        *ProgrammableMoveCall
	TransferObjects *TransferObjects
	SplitCoins      *SplitCoins
	MergeCoins      *MergeCoins
	Publish         *Publish
	MakeMoveVec     *MakeMoveVec
	Upgrade         *Upgrade
}

func (*Command) IsBcsEnum() {}

// ProgrammableMoveCall https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L762
type ProgrammableMoveCall struct {
	Package       models.SuiAddressBytes
	Module        string
	Function      string
	TypeArguments []string
	Arguments     []*Argument
}

type TransferObjects struct {
	Objects []*Argument
	Address *Argument
}

type SplitCoins struct {
	Coin   *Argument
	Amount []*Argument
}

type MergeCoins struct {
	Destination *Argument
	Sources     []*Argument
}

type Publish struct {
	Modules      []models.SuiAddressBytes
	Dependencies []models.SuiAddressBytes
}

type MakeMoveVec struct {
	Type     *string
	Elements []*Argument
}

type Upgrade struct {
	Modules      []models.SuiAddressBytes
	Dependencies []models.SuiAddressBytes
	Package      models.SuiAddressBytes
	Ticket       *Argument
}

// Argument https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L745
// - GasCoin
// - Input
// - Result
// - NestedResult
type Argument struct {
	GasCoin      *bool
	Input        *uint16
	Result       *uint16
	NestedResult *NestedResult
}

func (*Argument) IsBcsEnum() {}

type NestedResult struct {
	Index       uint16
	ResultIndex uint16
}

type SuiObjectRef struct {
	ObjectId models.SuiAddressBytes
	Version  uint64
	Digest   models.ObjectDigestBytes
}

type SharedObjectRef struct {
	ObjectId             models.SuiAddressBytes
	InitialSharedVersion uint64
	Mutable              bool
}
