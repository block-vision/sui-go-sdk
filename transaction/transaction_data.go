package transaction

import (
	"bytes"
	"fmt"
	"io"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/mystenbcs"
)

type TransactionData struct {
	V1 *TransactionDataV1
}

func (*TransactionData) IsBcsEnum() {}

func (td *TransactionData) Marshal() ([]byte, error) {
	if td.V1 != nil {
		if err := td.V1.Kind.checkResolved(); err != nil {
			return nil, err
		}
		if td.V1.Expiration == nil {
			v1 := *td.V1
			v1.Expiration = &TransactionExpiration{None: struct{}{}}
			td = &TransactionData{V1: &v1}
		}
	}

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
	Kind    *TransactionKind
	Sender  *models.SuiAddressBytes
	GasData *GasData
	// Not an Option on chain; a nil Expiration is encoded as TransactionExpiration::None.
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
		if input.Object == nil {
			continue
		}

		if input.Object.ImmOrOwnedObject != nil {
			objectId := input.Object.ImmOrOwnedObject.ObjectId
			if objectId.IsEqual(*addressBytes) {
				index := uint16(i)
				return &index
			}
		}
		if input.Object.SharedObject != nil {
			objectId := input.Object.SharedObject.ObjectId
			if objectId.IsEqual(*addressBytes) {
				index := uint16(i)
				return &index
			}
		}
		if input.Object.Receiving != nil {
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
	Payment *[]SuiObjectRef
	Owner   *models.SuiAddressBytes
	Price   *uint64
	Budget  *uint64
}

func (gd *GasData) IsAllSet() bool {
	if gd.Payment == nil || gd.Owner == nil || gd.Price == nil || gd.Budget == nil {
		return false
	}

	return true
}

// TransactionExpiration https://github.com/MystenLabs/sui/blob/mainnet-v1.81.1/crates/sui-types/src/transaction.rs
// - None
// - Epoch
// - ValidDuring
// - Validity
type TransactionExpiration struct {
	None        any
	Epoch       *uint64
	ValidDuring *ValidDuring
	Validity    *Validity
}

func (*TransactionExpiration) IsBcsEnum() {}

// ValidDuring bounds the epochs and timestamps (Unix milliseconds) a transaction may
// execute in. Transactions that pay gas from an address balance (empty gas payment)
// must use it, with a nonce to keep otherwise identical transactions distinct.
type ValidDuring struct {
	MinEpoch     *uint64 `bcs:"optional"`
	MaxEpoch     *uint64 `bcs:"optional"`
	MinTimestamp *uint64 `bcs:"optional"`
	MaxTimestamp *uint64 `bcs:"optional"`
	// Chain is the chain identifier (genesis checkpoint digest), 32 bytes.
	Chain models.ObjectDigestBytes
	Nonce uint32
}

// Validity is ValidDuring plus an optional restriction on which validators may
// propose the transaction in consensus.
type Validity struct {
	MinEpoch         *uint64 `bcs:"optional"`
	MaxEpoch         *uint64 `bcs:"optional"`
	MinTimestamp     *uint64 `bcs:"optional"`
	MaxTimestamp     *uint64 `bcs:"optional"`
	Chain            models.ObjectDigestBytes
	Nonce            uint32
	AllowedProposers *AllowedProposers `bcs:"optional"`
}

// AllowedProposers holds committee indices, strictly increasing and non-empty,
// into Epoch's committee.
type AllowedProposers struct {
	Epoch     uint64
	Proposers []uint32
}

// ProgrammableTransaction https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L702
type ProgrammableTransaction struct {
	Inputs   []*CallArg
	Commands []*Command
}

// TransactionKind https://github.com/MystenLabs/sui/blob/mainnet-v1.81.1/crates/sui-types/src/transaction.rs
// Field order is the variant index. Only the programmable kinds can be encoded or
// decoded; the other system kinds are placeholders that fail with
// ErrUnsupportedTransactionKind instead of decoding into the wrong layout.
type TransactionKind struct {
	ProgrammableTransaction       *ProgrammableTransaction
	ChangeEpoch                   *UnsupportedTransactionKind
	Genesis                       *UnsupportedTransactionKind
	ConsensusCommitPrologue       *UnsupportedTransactionKind
	AuthenticatorStateUpdate      *UnsupportedTransactionKind
	EndOfEpochTransaction         *UnsupportedTransactionKind
	RandomnessStateUpdate         *UnsupportedTransactionKind
	ConsensusCommitPrologueV2     *UnsupportedTransactionKind
	ConsensusCommitPrologueV3     *UnsupportedTransactionKind
	ConsensusCommitPrologueV4     *UnsupportedTransactionKind
	ProgrammableSystemTransaction *ProgrammableTransaction
}

func (*TransactionKind) IsBcsEnum() {}

// checkResolved rejects SDK-only placeholder inputs, which have no on-chain encoding.
func (tk *TransactionKind) checkResolved() error {
	if tk == nil || tk.ProgrammableTransaction == nil {
		return nil
	}
	for i, input := range tk.ProgrammableTransaction.Inputs {
		if input != nil && (input.UnresolvedPure != nil || input.UnresolvedObject != nil) {
			return fmt.Errorf("input %d: %w", i, ErrUnresolvedInput)
		}
	}
	return nil
}

// UnsupportedTransactionKind stands in for system transaction kinds whose payload
// this SDK does not model.
type UnsupportedTransactionKind struct{}

func (*UnsupportedTransactionKind) MarshalBCS() ([]byte, error) {
	return nil, ErrUnsupportedTransactionKind
}

func (*UnsupportedTransactionKind) UnmarshalBCS(io.Reader) (int, error) {
	return 0, ErrUnsupportedTransactionKind
}

func (tk *TransactionKind) Marshal() ([]byte, error) {
	if err := tk.checkResolved(); err != nil {
		return nil, err
	}

	bcsEncodedMsg := bytes.Buffer{}
	bcsEncoder := mystenbcs.NewEncoder(&bcsEncodedMsg)
	err := bcsEncoder.Encode(tk)
	if err != nil {
		return nil, err
	}

	return bcsEncodedMsg.Bytes(), nil
}

// CallArg https://github.com/MystenLabs/sui/blob/mainnet-v1.81.1/crates/sui-types/src/transaction.rs
// - Pure
// - Object
// - FundsWithdrawal
//
// UnresolvedPure and UnresolvedObject are SDK-only placeholders, not on-chain variants:
// they must come after the on-chain ones and are never encoded, so building a
// transaction that still holds one fails with ErrUnresolvedInput.
type CallArg struct {
	Pure             *Pure
	Object           *ObjectArg
	FundsWithdrawal  *FundsWithdrawalArg
	UnresolvedPure   *UnresolvedPure   `bcs:"-"`
	UnresolvedObject *UnresolvedObject `bcs:"-"`
}

func (*CallArg) IsBcsEnum() {}

type Pure struct {
	Bytes []byte
}

type UnresolvedPure struct {
	Value any
}

type UnresolvedObject struct {
	ObjectId models.SuiAddressBytes
	// Version
	// Digest
	// InitialSharedVersion
}

// FundsWithdrawalArg reserves funds from an address balance; the PTB redeems it
// with 0x2::balance::redeem_funds or 0x2::coin::redeem_funds.
type FundsWithdrawalArg struct {
	Reservation  *Reservation
	TypeArg      *WithdrawalTypeArg
	WithdrawFrom *WithdrawFrom
}

// Reservation
// - MaxAmountU64
type Reservation struct {
	MaxAmountU64 *uint64
}

func (*Reservation) IsBcsEnum() {}

// WithdrawalTypeArg
// - Balance: the T of Balance<T>
type WithdrawalTypeArg struct {
	Balance *TypeTag
}

func (*WithdrawalTypeArg) IsBcsEnum() {}

// WithdrawFrom
// - Sender
// - Sponsor
// - SenderAllowance
type WithdrawFrom struct {
	Sender          any
	Sponsor         any
	SenderAllowance *SenderAllowance
}

func (*WithdrawFrom) IsBcsEnum() {}

// SenderAllowance withdraws from Funder's balance under the Allowance object
// granted to the sender.
type SenderAllowance struct {
	Funder    models.SuiAddressBytes
	Allowance models.SuiAddressBytes
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

func (ObjectArg) IsBcsEnum() {}

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
	TypeArguments []*TypeTag
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
	Modules      [][]byte
	Dependencies []models.SuiAddressBytes
}

// MakeMoveVec https://github.com/MystenLabs/sui/blob/fb27c6c7166f5e4279d5fd1b2ebc5580ca0e81b2/crates/sui-types/src/transaction.rs#L731
// Type is an Option<TypeInput> (same layout as TypeTag), it is nil when the vector holds objects.
type MakeMoveVec struct {
	Type     *TypeTag `bcs:"optional"`
	Elements []*Argument
}

type Upgrade struct {
	Modules      [][]byte
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
	GasCoin      any
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

// SharedObjectRef: on chain the last field is SharedObjectMutability
// (Immutable=0, Mutable=1, NonExclusiveWrite=2). Immutable and Mutable encode exactly
// like this bool; NonExclusiveWrite is only used by system settlement transactions
// and is decoded as Mutable=true.
type SharedObjectRef struct {
	ObjectId             models.SuiAddressBytes
	InitialSharedVersion uint64
	Mutable              bool
}

type StructTag struct {
	Address    models.SuiAddressBytes
	Module     string
	Name       string
	TypeParams []*TypeTag
}

// TypeTag https://github.com/MystenLabs/sui/blob/ece197ed5c414eb274f99afc52704664af8d0c38/external-crates/move/crates/move-core-types/src/language_storage.rs#L33
// Do not reorder the fields, it will break the bcs encoding
// Primitive variants carry no payload; set them to &struct{}{} (see ParseTypeTag).
type TypeTag struct {
	Bool    *struct{}
	U8      *struct{}
	U128    *struct{}
	U256    *struct{}
	Address *struct{}
	Signer  *struct{}
	Vector  *TypeTag
	Struct  *StructTag
	U16     *struct{}
	U32     *struct{}
	U64     *struct{}
}

func (*TypeTag) IsBcsEnum() {}
