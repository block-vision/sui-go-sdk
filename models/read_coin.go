package models

type SuiXGetBalanceRequest struct {
	Owner    string `json:"owner"`
	CoinType string `json:"coinType"`
}

type SuiXGetAllBalanceRequest struct {
	Owner string `json:"owner"`
}

type CoinLockedBalance struct {
	EpochId int `json:"epochId"`
	Number  int `json:"number"`
}

type CoinBalanceResponse struct {
	CoinType        string            `json:"coinType"`
	CoinObjectCount int               `json:"coinObjectCount"`
	TotalBalance    string            `json:"totalBalance"`
	LockedBalance   CoinLockedBalance `json:"lockedBalance"`
}

type CoinAllBalanceResponse []CoinBalanceResponse

type SuiXGetCoinsRequest struct {
	Owner    string      `json:"owner"`
	CoinType string      `json:"coin_type"`
	Cursor   interface{} `json:"cursor"`
	Limit    uint64      `json:"limit" validate:"lte=50"`
}

type PaginatedCoinsResponse struct {
	Data        []CoinData `json:"data"`
	NextCursor  string     `json:"nextCursor"`
	HasNextPage bool       `json:"hasNextPage"`
}

type CoinData struct {
	CoinType            string `json:"coinType"`
	CoinObjectId        string `json:"coinObjectId"`
	Version             string `json:"version"`
	Digest              string `json:"digest"`
	Balance             string `json:"balance"`
	LockedUntilEpoch    uint64 `json:"lockedUntilEpoch"`
	PreviousTransaction string `json:"previousTransaction"`
}

type SuiXGetAllCoinsRequest struct {
	Owner  string      `json:"owner"`
	Cursor interface{} `json:"cursor"`
	Limit  uint64      `json:"limit" validate:"lte=50"`
}

type SuiXGetCoinMetadataRequest struct {
	CoinType string `json:"coinType"`
}

type CoinMetadataResponse struct {
	Id          string `json:"id"`
	Decimals    int    `json:"decimals"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
}

type SuiXGetTotalSupplyRequest struct {
	CoinType string `json:"coinType"`
}

type TotalSupplyResponse struct {
	Value string `json:"value"`
}
