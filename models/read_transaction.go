package models

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
	Kind         string           `json:"kind"`
	Inputs       []SuiCallArg     `json:"inputs"`
	Transactions []SuiTransaction `json:"transactions"`
}

type SuiTransaction struct {
	MakeMoveVec     []interface{}          `json:"MakeMoveVec,omitempty"`
	MergeCoins      []interface{}          `json:"MergeCoins,omitempty"`
	SplitCoins      []interface{}          `json:"SplitCoins,omitempty"`
	TransferObjects []interface{}          `json:"TransferObjects,omitempty"`
	Publish         []interface{}          `json:"Publish,omitempty"`
	Upgrade         []interface{}          `json:"Upgrade,omitempty"`
	MoveCall        MoveCallSuiTransaction `json:"MoveCall,omitempty"`
}

type ProgrammableTransaction struct {
	Transactions []SuiTransaction `json:"transactions"`
	Inputs       []SuiCallArg     `json:"inputs"`
}

type MoveCallSuiTransaction struct {
	Package       string   `json:"package"`
	Module        string   `json:"module"`
	Function      string   `json:"function"`
	TypeArguments []string `json:"type_arguments"`
	Arguments     []struct {
		Input  int `json:"Input,omitempty"`
		Result int `json:"Result,omitempty"`
	} `json:"arguments"`
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
	Version  int    `json:"version"`
	Digest   string `json:"digest"`
}

type SuiGasData struct {
	Payment []SuiObjectRef `json:"payment"`
	Owner   string         `json:"owner"`
	Price   string         `json:"price"`
	Budget  string         `json:"budget"`
}

type SuiObjectChangePublished struct {
	Type      string   `json:"type"`
	PackageId string   `json:"packageId"`
	Version   int      `json:"version"`
	Digest    string   `json:"digest"`
	Modules   []string `json:"modules"`
}

type SuiObjectChangeTransferred struct {
	Type       string      `json:"type"`
	Sender     string      `json:"sender"`
	Recipient  ObjectOwner `json:"recipient"`
	ObjectType string      `json:"objectType"`
	ObjectId   string      `json:"objectId"`
	Version    int         `json:"version"`
	Digest     string      `json:"digest"`
}

type SuiObjectChangeMutated struct {
	Type            string      `json:"type"`
	Sender          string      `json:"sender"`
	Owner           ObjectOwner `json:"owner"`
	ObjectType      string      `json:"objectType"`
	ObjectId        string      `json:"objectId"`
	Version         int         `json:"version"`
	PreviousVersion int         `json:"previousVersion"`
	Digest          string      `json:"digest"`
}

type SuiObjectChangeDeleted struct {
	Type       string `json:"type"`
	Sender     string `json:"sender"`
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
	Version    int    `json:"version"`
}

type SuiObjectChangeWrapped struct {
	Type       string `json:"type"`
	Sender     string `json:"sender"`
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
	Version    int    `json:"version"`
}

type SuiObjectChangeCreated struct {
	Type       string      `json:"type"`
	Sender     string      `json:"sender"`
	Owner      ObjectOwner `json:"owner"`
	ObjectType string      `json:"objectType"`
	ObjectId   string      `json:"objectId"`
	Version    int         `json:"version"`
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

type ObjectChange struct {
	Type            string      `json:"type"`
	Sender          string      `json:"sender"`
	Owner           ObjectOwner `json:"owner"`
	ObjectType      string      `json:"objectType"`
	ObjectId        string      `json:"objectId"`
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

type SuiXQueryTransactionBlocksRequest struct {
	SuiTransactionBlockResponseQuery SuiTransactionBlockResponseQuery
	Cursor                           interface{} `json:"cursor"`
	Limit                            uint64      `json:"limit" validate:"lte=50"`
	DescendingOrder                  bool        `json:"descendingOrder"`
}

type SuiXQueryTransactionBlocksResponse struct {
	Data        []SuiTransactionBlockResponse `json:"data"`
	NextCursor  string                        `json:"nextCursor"`
	HasNextPage bool                          `json:"hasNextPage"`
}
