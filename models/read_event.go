package models

type SuiGetEventsRequest struct {
	Digest string `json:"digest"`
}

type EventId struct {
	TxDigest string `json:"txDigest"`
	EventSeq string `json:"eventSeq"`
}

type SuiEventResponse struct {
	Id                EventId                `json:"id"`
	PackageId         string                 `json:"packageId"`
	TransactionModule string                 `json:"transactionModule"`
	Sender            string                 `json:"sender"`
	Type              string                 `json:"type"`
	ParsedJson        map[string]interface{} `json:"parsedJson"`
	Bcs               string                 `json:"bcs"`
	TimestampMs       string                 `json:"timestampMs"`
}

type GetEventsResponse []*SuiEventResponse

type MoveModule struct {
	Package string `json:"package"`
	Module  string `json:"module"`
}

type MoveEventModule struct {
	Package string `json:"package"`
	Module  string `json:"module"`
	Type    string `json:"type"`
}

type MoveEventField struct {
	Field string `json:"field"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type TimeRange struct {
	StartTime uint64 `json:"start_time"`
	EndTime   uint64 `json:"end_time"`
}

type SuiEventFilter map[string]interface{}

// the event query by `Package`: Move package ID
// JSON-RPC Parameter Example: {"Package":"<PACKAGE-ID>"}
type EventFilterByPackage struct {
	Package string `json:"Package"`
}

// the event query by `MoveModule`: Move module where the event was emitted
// JSON-RPC Parameter Example: {"MoveModule": {"package": "<PACKAGE-ID>", "module": "nft"}}
type EventFilterByMoveModule struct {
	MoveModule MoveModule `json:"MoveModule"`
}

// the event query by `MoveEventType`: Move event type defined in the move code
// JSON-RPC Parameter Example:{"MoveEventType":"<PACKAGE-ID>::nft::MintNFTEvent"}
type EventFilterByMoveEventType struct {
	MoveEventType string `json:"MoveEventType"`
}

// the event query by `MoveEventModule`: Move event module defined in the move code
// JSON-RPC Parameter Example: {"MoveEventModule": {"package": "<PACKAGE-ID>", "module": "nft", "event": "MintNFTEvent"}}
type EventFilterByMoveEventModule struct {
	MoveEventModule MoveEventModule `json:"MoveEventModule"`
}

// the event query by `MoveEventField`: Filter using the data fields in the move event object
// JSON-RPC Parameter Example: {"MoveEventField":{ "path":"/name", "value":"NFT"}}
type EventFilterByMoveEventField struct {
	MoveEventField MoveEventField `json:"MoveEventField"`
}

// the event query by `Transaction`: Filter Transaction hash
// JSON-RPC Parameter Example: {"Transaction":"ENmjG42TE4GyqYb1fGNwJe7oxBbbXWCdNfRiQhCNLBJQ"}
type EventFilterByTransaction struct {
	Transaction string `json:"Transaction"`
}

// the event query by `TimeRange`: Time range in millisecond
// JSON-RPC Parameter Example: {"TimeRange": {"start_time": "1685959791871", "end_time": "1685959791871"}}
type EventFilterByTimeRange struct {
	TimeRange TimeRange `json:"TimeRange"`
}

// the event query by `Sender`: Filter Sender address
// JSON-RPC Parameter Example: {"Sender":"0x008e9c621f4fdb210b873aab59a1e5bf32ddb1d33ee85eb069b348c234465106"}
type EventFilterBySuiAddress struct {
	Sender string `json:"Sender"`
}

// the event query by `SenderAddress`: Address that started the transaction
// JSON-RPC Parameter Example: {"SenderAddress": "0x008e9c621f4fdb210b873aab59a1e5bf32ddb1d33ee85eb069b348c234465106"}
type EventFilterBySenderAddress struct {
	SenderAddress string `json:"SenderAddress"`
}

type SuiXQueryEventsRequest struct {
	// the event query criteria. See Event filter documentation[https://docs.sui.io/sui-jsonrpc#suix_queryEvents] for examples.
	SuiEventFilter interface{} `json:"suiEventFilter"`
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
	// query result ordering, default to false (ascending order), oldest record first
	DescendingOrder bool `json:"descendingOrder"`
}

type PaginatedEventsResponse struct {
	Data        []SuiEventResponse `json:"data"`
	NextCursor  EventId            `json:"nextCursor"`
	HasNextPage bool               `json:"hasNextPage"`
}
