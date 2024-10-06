package models

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

type GetTransactionMetaData struct {
	GatewayTxSeqNumber uint64 `json:"gatewayTxSeqNumber"`
	TransactionDigest  string `json:"transactionDigest"`
}

type SuiTransactionBlockOptions struct {
	ShowInput          bool `json:"showInput,omitempty"`
	ShowRawInput       bool `json:"showRawInput,omitempty"`
	ShowEffects        bool `json:"showEffects,omitempty"`
	ShowEvents         bool `json:"showEvents,omitempty"`
	ShowObjectChanges  bool `json:"showObjectChanges,omitempty"`
	ShowBalanceChanges bool `json:"showBalanceChanges,omitempty"`
}

type SuiGetTransactionBlockRequest struct {
	Digest  string                     `json:"digest"`
	Options SuiTransactionBlockOptions `json:"options"`
}

type SuiArgument map[string]interface{}
type SuiCallArg map[string]interface{}

type SuiTransactionBlockKind struct {
	Kind         string       `json:"kind"`
	Inputs       []SuiCallArg `json:"inputs"`
	Transactions []any        `json:"transactions"`
}

func MoveCall(data any) *MoveCallSuiTransaction {
	bs, _ := json.Marshal(data)
	res := gjson.GetBytes(bs, "MoveCall").Raw

	if res != "" {
		var data *MoveCallSuiTransaction
		_ = json.Unmarshal([]byte(res), &data)
		return data
	}
	return nil
}

type SuiTransactionEnum struct {
	MakeMoveVec     []interface{}           `json:"MakeMoveVec,omitempty"`
	MergeCoins      []interface{}           `json:"MergeCoins,omitempty"`
	SplitCoins      []interface{}           `json:"SplitCoins,omitempty"`
	TransferObjects []interface{}           `json:"TransferObjects,omitempty"`
	Publish         []interface{}           `json:"Publish,omitempty"`
	Upgrade         []interface{}           `json:"Upgrade,omitempty"`
	MoveCall        *MoveCallSuiTransaction `json:"MoveCall,omitempty"`
}

type ProgrammableTransaction struct {
	Transactions []any        `json:"transactions"`
	Inputs       []SuiCallArg `json:"inputs"`
}

type MoveCallSuiTransaction struct {
	Package       string        `json:"package"`
	Module        string        `json:"module"`
	Function      string        `json:"function"`
	TypeArguments []string      `json:"type_arguments"`
	Arguments     []interface{} `json:"arguments"`
}

type SuiTransactionBlockData struct {
	MessageVersion string                  `json:"messageVersion"`
	Transaction    SuiTransactionBlockKind `json:"transaction"`
	Sender         string                  `json:"sender"`
	GasData        SuiGasData              `json:"gasData"`
}

type SuiTransactionBlock struct {
	Data         SuiTransactionBlockData `json:"data"`
	TxSignatures []string                `json:"txSignatures"`
}

type SuiObjectRef struct {
	ObjectId string `json:"objectId"`
	Version  uint64 `json:"version"`
	Digest   string `json:"digest"`
}

type SuiGasData struct {
	Payment []SuiObjectRef `json:"payment"`
	// the owner's Sui address
	Owner  string `json:"owner"`
	Price  string `json:"price"`
	Budget string `json:"budget"`
}

type SuiObjectChangePublished struct {
	Type      string   `json:"type"`
	PackageId string   `json:"packageId"`
	Version   uint64   `json:"version"`
	Digest    string   `json:"digest"`
	Modules   []string `json:"modules"`
}

type SuiObjectChangeTransferred struct {
	Type       string      `json:"type"`
	Sender     string      `json:"sender"`
	Recipient  ObjectOwner `json:"recipient"`
	ObjectType string      `json:"objectType"`
	ObjectId   string      `json:"objectId"`
	Version    uint64      `json:"version"`
	Digest     string      `json:"digest"`
}

type SuiObjectChangeMutated struct {
	Type            string      `json:"type"`
	Sender          string      `json:"sender"`
	Owner           ObjectOwner `json:"owner"`
	ObjectType      string      `json:"objectType"`
	ObjectId        string      `json:"objectId"`
	Version         uint64      `json:"version"`
	PreviousVersion uint64      `json:"previousVersion"`
	Digest          string      `json:"digest"`
}

type SuiObjectChangeDeleted struct {
	Type       string `json:"type"`
	Sender     string `json:"sender"`
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
	Version    uint64 `json:"version"`
}

type SuiObjectChangeWrapped struct {
	Type       string `json:"type"`
	Sender     string `json:"sender"`
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
	Version    uint64 `json:"version"`
}

type SuiObjectChangeCreated struct {
	Type       string      `json:"type"`
	Sender     string      `json:"sender"`
	Owner      ObjectOwner `json:"owner"`
	ObjectType string      `json:"objectType"`
	ObjectId   string      `json:"objectId"`
	Version    uint64      `json:"version"`
	Digest     string      `json:"digest"`
}

type OwnedObjectRef struct {
	Owner     interface{}  `json:"owner"`
	Reference SuiObjectRef `json:"reference"`
}

type SuiEffects struct {
	MessageVersion     string               `json:"messageVersion"`
	Status             ExecutionStatus      `json:"status"`
	ExecutedEpoch      string               `json:"executedEpoch"`
	GasUsed            GasCostSummary       `json:"gasUsed"`
	ModifiedAtVersions []ModifiedAtVersions `json:"modifiedAtVersions"`
	SharedObjects      []SuiObjectRef       `json:"sharedObjects"`
	TransactionDigest  string               `json:"transactionDigest"`
	Created            []OwnedObjectRef     `json:"created"`
	Mutated            []OwnedObjectRef     `json:"mutated"`
	Deleted            []SuiObjectRef       `json:"deleted"`
	GasObject          OwnedObjectRef       `json:"gasObject"`
	EventsDigest       string               `json:"eventsDigest"`
	Dependencies       []string             `json:"dependencies"`
}

type ExecutionStatus struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type GasCostSummary struct {
	ComputationCost         string `json:"computationCost"`
	StorageCost             string `json:"storageCost"`
	StorageRebate           string `json:"storageRebate"`
	NonRefundableStorageFee string `json:"nonRefundableStorageFee"`
}

type ModifiedAtVersions struct {
	ObjectId       string `json:"objectId"`
	SequenceNumber string `json:"sequenceNumber"`
}

type SuiTransactionBlockResponse struct {
	Digest                  string              `json:"digest"`
	Transaction             SuiTransactionBlock `json:"transaction,omitempty"`
	RawTransaction          string              `json:"rawTransaction,omitempty"`
	Effects                 SuiEffects          `json:"effects,omitempty"`
	Events                  []SuiEventResponse  `json:"events,omitempty"`
	ObjectChanges           []ObjectChange      `json:"objectChanges,omitempty"`
	BalanceChanges          []BalanceChanges    `json:"balanceChanges,omitempty"`
	TimestampMs             string              `json:"timestampMs,omitempty"`
	Checkpoint              string              `json:"checkpoint,omitempty"`
	ConfirmedLocalExecution bool                `json:"confirmedLocalExecution,omitempty"`
}

func (o ObjectChange) GetObjectChangeAddressOwner() string {
	b, err := json.Marshal(o.Owner)
	if err != nil {
		return ""
	}

	if gjson.ParseBytes(b).IsObject() {
		var owner ObjectOwner
		err = json.Unmarshal(b, &owner)
		if err == nil {
			return owner.AddressOwner
		}
	}
	return ""
}
func (o ObjectChange) GetObjectChangeObjectOwner() string {
	b, err := json.Marshal(o.Owner)
	if err != nil {
		return ""
	}

	if gjson.ParseBytes(b).IsObject() {
		var owner ObjectOwner
		err = json.Unmarshal(b, &owner)
		if err == nil {
			return owner.ObjectOwner
		}
	}
	return ""
}
func (o ObjectChange) GetObjectOwnerShare() ObjectShare {
	var share ObjectShare
	b, err := json.Marshal(o.Owner)
	if err != nil {
		return share
	}

	if gjson.ParseBytes(b).IsObject() {
		var owner ObjectOwner
		err = json.Unmarshal(b, &owner)
		if err == nil {
			return owner.Shared
		}
	}
	return share
}

type ObjectChange struct {
	Type            string      `json:"type"`
	Sender          string      `json:"sender"`
	Owner           interface{} `json:"owner"`
	ObjectType      string      `json:"objectType"`
	ObjectId        string      `json:"objectId"`
	PackageId       string      `json:"packageId"`
	Modules         []string    `json:"modules"`
	Version         string      `json:"version"`
	PreviousVersion string      `json:"previousVersion,omitempty"`
	Digest          string      `json:"digest"`
}

type BalanceChanges struct {
	Owner    ObjectOwner `json:"owner"`
	CoinType string      `json:"coinType"`
	Amount   string      `json:"amount"`
}

type SuiMultiGetTransactionBlocksRequest struct {
	Digests []string                   `json:"digests"`
	Options SuiTransactionBlockOptions `json:"options"`
}

type SuiMultiGetTransactionBlocksResponse []*SuiTransactionBlockResponse

type SuiTransactionBlockResponseQuery struct {
	TransactionFilter TransactionFilter          `json:"filter"`
	Options           SuiTransactionBlockOptions `json:"options"`
}

type TransactionFilter map[string]interface{}

// TransactionFilterByFromAddress is a filter for from address
type TransactionFilterByFromAddress struct {
	FromAddress string `json:"FromAddress"`
}

// TransactionFilterByToAddress is a filter for to address
type TransactionFilterByToAddress struct {
	ToAddress string `json:"ToAddress"`
}

// TransactionFilterByInputObject is a filter for input objects
type TransactionFilterByInputObject struct {
	// InputObject is the id of the object
	InputObject string `json:"InputObject"`
}

// TransactionFilterByChangedObjectFilter is a filter for changed objects
type TransactionFilterByChangedObjectFilter struct {
	// ChangedObject is a filter for changed objects
	ChangedObject string `json:"ChangedObject"`
}

// TransactionFilterByMoveFunction is a filter for move functions
type TransactionFilterByMoveFunction struct {
	MoveFunction MoveFunction `json:"MoveFunction"`
}

type MoveFunction struct {
	Package  string  `json:"package"`
	Module   *string `json:"module"`
	Function *string `json:"function"`
}

type SuiXSubscribeTransactionsRequest struct {
	// the transaction query criteria.
	TransactionFilter interface{} `json:"filter"`
}

type SuiXQueryTransactionBlocksRequest struct {
	SuiTransactionBlockResponseQuery SuiTransactionBlockResponseQuery
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
	// query result ordering, default to false (ascending order), oldest record first
	DescendingOrder bool `json:"descendingOrder"`
}

type SuiXQueryTransactionBlocksResponse struct {
	Data        []SuiTransactionBlockResponse `json:"data"`
	NextCursor  string                        `json:"nextCursor"`
	HasNextPage bool                          `json:"hasNextPage"`
}

type SuiDryRunTransactionBlockRequest struct {
	TxBytes string `json:"txBytes"`
}

type SuiDevInspectTransactionBlockRequest struct {
	// the transaction signer's Sui address
	Sender string `json:"sender"`
	// BCS encoded TransactionKind(as opposed to TransactionData, which include gasBudget and gasPrice)
	TxBytes string `json:"txBytes"`
	// Gas is not charged, but gas usage is still calculated. Default to use reference gas price
	GasPrice string `json:"gasPrice"`
	// The epoch to perform the call. Will be set from the system state object if not provided
	Epoch string `json:"epoch"`
}

type SuiXSubscribeEventsRequest struct {
	// the event query criteria.
	SuiEventFilter interface{} `json:"suiEventFilter"`
}
