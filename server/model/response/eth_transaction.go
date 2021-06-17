package response

import "service/server/model"

type TxHash struct {
	Hash string `json:"hash" bson:"hash"`
}

type Hashs struct {
	TxHashs []TxHash `json:"txhashs" bson:"txhashs"`
}

type EthTx struct {
	TxHashs []model.ETHTransaction `json:"txhashs" bson:"txhashs"`
}

type EthTxV2 struct {
	TxHashs []model.ETHTransactionV2 `json:"txhashs" bson:"txhashs"`
}
