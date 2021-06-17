package response

import "service/server/model"

type BtcTx struct {
	TxHashs []BtcTxResp `json:"txhashs" bson:"txhashs"`
}

type BtcTxResp struct {
	Hex           string          `json:"hex" bson:"hex"`
	Txid          string          `json:"txid" bson:"txid"`
	Hash          string          `json:"hash,omitempty" bson:"hash"`
	Size          int32           `json:"size,omitempty" bson:"size"`
	Vsize         int32           `json:"vsize,omitempty" bson:"vsize"`
	Version       int32           `json:"version" bson:"version"`
	LockTime      uint32          `json:"locktime" bson:"locktime"`
	BlockHash     string          `json:"blockhash,omitempty" bson:"blockhash"`
	Time          int64           `json:"time,omitempty" bson:"time"`
	Blocktime     int64           `json:"blocktime,omitempty" bson:"blocktime"`
	BlockHeight   uint64          `json:"blockheight" bson:"blockheight"`
	FromAddresses []AddressAmount `json:"from_addresses"`
	ToAddresses   []AddressAmount `json:"to_addresses"`
	*OmniResp
}

type OmniResp struct {
	Fee        string `json:"fee"`
	PropertyId uint32 `json:"property_id"`
	Valid      bool   `json:"valid"`
	Divisible  bool   `json:"divisible"`
}

type AddressAmount struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

type BtcRawTxs struct {
	Txs []model.BtcRawTransaction `json:"txs" bson:"txs"`
}

type BtcOmniTxs struct {
	Txs []model.OmniTransaction `json:"txs" bson:"txs"`
}
