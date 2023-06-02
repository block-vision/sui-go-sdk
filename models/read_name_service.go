package models

type SuiXResolveNameServiceAddressRequest struct {
	Name string `json:"name"`
}

type SuiXResolveNameServiceNamesRequest struct {
	Address string `json:"address"`
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
}

type SuiXResolveNameServiceNamesResponse struct {
	Data        []string `json:"data"`
	NextCursor  string   `json:"nextCursor"`
	HasNextPage bool     `json:"hasNextPage"`
}
