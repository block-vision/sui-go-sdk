package models

type GetCoinsRequeset struct {
	Owner    string `json:"owner"`
	CoinType string `json:"coin_type"`
	Cursor   string `json:"cursor"`
	Limit    uint64 `json:"limit"`
}

type GetCoinsResponse struct {
	Data []CoinPage `json:"data"`
}

type CoinPage struct {
	CoinType     string `json:"coinType"`
	CoinObjectId string `json:"coinObjectId"`
	Version      uint64 `json:"version"`
	Digest       string `json:"digest"`
	Balance      uint64 `json:"balance"`
}

type GetAllCoinsRequest struct {
	Owner  string
	Cursor *string
	Limit  uint64
}

type GetAllCoinsResponse struct {
	Data       []CoinPage `json:"data"`
	NextCursor string     `json:"nextCursor"`
}

type GetBalanceRequest struct {
	Owner    string
	CoinType string
}

type GetBalanceResponse struct {
	Balance
}

type GetAllBalancesRequest struct {
	Owner string
}

type GetAllBalancesResponse struct {
	Balance []Balance `json:"balance"`
}

type GetTotalSupplyRequest struct {
	CoinType string
}

type GetTotalSupplyResponse struct{}

type Balance struct {
	CoinType        string      `json:"coinType"`
	CoinObjectCount uint64      `json:"coinObjectCount"`
	TotalBalance    uint64      `json:"totalBalance"`
	LockedBalance   interface{} `json:"lockedBalance"`
}
