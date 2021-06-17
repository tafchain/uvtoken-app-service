package request

type Hashearch struct {
	CoinType      string `json:"coinType" form:"coinType"`
	TokenAddress  string `json:"tokenAddress" form:"tokenAddress"`
	Address       string `json:"address" form:"address"`
	FromBlock     uint64 `json:"fromBlock" form:"fromBlock"`
	ToBlock       uint64 `json:"toBlock" form:"toBlock"`
	Confirmations uint64 `json:"confirmations" form:"confirmations"`
}

type TxSearch struct {
	CoinType string   `json:"coinType" form:"coinType"`
	Hash     []string `json:"hash" form:"hash"`
}
