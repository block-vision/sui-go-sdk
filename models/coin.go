package models

type GetCoinsRequeset struct {
	Owner    string
	CoinType string
	Cursor   string
	Limit    uint64
}

type GetCoinsResponse struct {
	Data []CoinPage `json:"data"`
}

type CoinPage struct {
	CoinType         string      `json:"coinType"`
	CoinObjectId     string      `json:"coinObjectId"`
	Version          uint64      `json:"version"`
	Digest           string      `json:"digest"`
	Balance          uint64      `json:"balance"`
	LockedUntilEpoch interface{} `json:"lockedUntilEpoch"`
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
	CoinType        string
	CoinObjectCount uint64
	TotalBalance    uint64
	LockedBalance   interface{}
}
