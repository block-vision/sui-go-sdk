package wsconn

type SubscriptionResp struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  int64  `json:"result"`
	Id      int64  `json:"id"`
}
