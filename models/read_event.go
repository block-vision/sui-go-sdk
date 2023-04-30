package models

type SuiGetEventsRequest struct {
	Digest string `json:"digest"`
}

type EventId struct {
	TxDigest string `json:"txDigest"`
	EventSeq string `json:"eventSeq"`
}

type ParsedJson struct {
	Amount           string `json:"amount"`
	Epoch            string `json:"epoch"`
	PoolId           string `json:"pool_id"`
	StakerAddress    string `json:"staker_address"`
	ValidatorAddress string `json:"validator_address"`
}

type SuiEventResponse struct {
	Id                EventId    `json:"id"`
	PackageId         string     `json:"packageId"`
	TransactionModule string     `json:"transactionModule"`
	Sender            string     `json:"sender"`
	Type              string     `json:"type"`
	ParsedJson        ParsedJson `json:"parsedJson"`
	Bcs               string     `json:"bcs"`
	TimestampMs       string     `json:"timestampMs"`
}

type GetEventsResponse []*SuiEventResponse

type MoveModule struct {
	Package string `json:"package"`
	Module  string `json:"module"`
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

// the event query by `Package`
type EventFilterByPackage struct {
	Package string `json:"Package"`
}

// the event query by `MoveModule`
type EventFilterByMoveModule struct {
	MoveModule MoveModule `json:"MoveModule"`
}

// the event query by `MoveEventType`
type EventFilterByMoveEventType struct {
	MoveEventType string `json:"MoveEventType"`
}

// the event query by `MoveEventField`
type EventFilterByMoveEventField struct {
	MoveEventField MoveEventField `json:"MoveEventField"`
}

// the event query by `Transaction`
type EventFilterByTransaction struct {
	Transaction string `json:"Transaction"`
}

// the event query by `TimeRange`
type EventFilterByTimeRange struct {
	TimeRange TimeRange `json:"TimeRange"`
}

// the event query by `Sender`
type EventFilterBySuiAddress struct {
	Sender string `json:"Sender"`
}

type SuiXQueryEventsRequest struct {
	// the event query criteria.
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
