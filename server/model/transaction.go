package model

import "time"

type ERCToken struct {
	From  string `json:"_from" bson:"_from"`
	To    string `json:"_to" bson:"_to"`
	Value string `json:"_value" bson:"_value"`
}

type ETHTransaction struct {
	Hash             string    `json:"hash" bson:"hash"`
	BlockHash        string    `json:"blockHash" bson:"blockHash"`
	BlockNumber      uint64    `json:"blockNumber" bson:"blockNumber"`
	ContractAddress  string    `json:"contractAddress" bson:"contractAddress"`
	From_            string    `json:"from_" bson:"from_"`
	Gas              string    `json:"gas" bson:"gas"`
	GasUsed          string    `json:"gasUsed" bson:"gasUsed"`
	GasPrice         string    `json:"gasPrice" bson:"gasPrice"`
	Input            string    `json:"input" bson:"input"`
	InputDecode      ERCToken  `json:"inputDecode" bson:"inputDecode"`
	Nonce            string    `json:"nonce" bson:"nonce"`
	To_              string    `json:"to_" bson:"to_"`
	Status           uint      `json:"status" bson:"status"`
	TransactionIndex uint      `json:"transactionIndex" bson:"transactionIndex"`
	Value_           string    `json:"value_" bson:"value_"`
	Timestamp        uint64    `json:"timestamp" bson:"timestamp"`
	Confirmations    uint64    `json:"confirmations" bson:"confirmations"`
	CreationDate     time.Time `json:"creation_date" bson:"creation_date"`
	ModifiedDate     time.Time `json:"modified_date" bson:"modified_date"`
}

type ERCTokenV2 struct {
	Address  string `json:"_address" bson:"_address"`
	From     string `json:"_from" bson:"_from"`
	To       string `json:"_to" bson:"_to"`
	Value    string `json:"_value" bson:"_value"`
	LogIndex uint   `json:"_logIndex" bson:"_logIndex"`
}

type ETHTransactionV2 struct {
	Hash             string       `json:"hash" bson:"hash"`
	BlockHash        string       `json:"blockHash" bson:"blockHash"`
	BlockNumber      uint64       `json:"blockNumber" bson:"blockNumber"`
	ContractAddress  string       `json:"contractAddress" bson:"contractAddress"`
	From_            string       `json:"from_" bson:"from_"`
	Gas              string       `json:"gas" bson:"gas"`
	GasUsed          string       `json:"gasUsed" bson:"gasUsed"`
	GasPrice         string       `json:"gasPrice" bson:"gasPrice"`
	Input            string       `json:"input" bson:"input"`
	Logs             []ERCTokenV2 `json:"logs" bson:"logs"`
	Nonce            string       `json:"nonce" bson:"nonce"`
	To_              string       `json:"to_" bson:"to_"`
	Status           uint         `json:"status" bson:"status"`
	TransactionIndex uint         `json:"transactionIndex" bson:"transactionIndex"`
	Value_           string       `json:"value_" bson:"value_"`
	Timestamp        uint64       `json:"timestamp" bson:"timestamp"`
	Confirmations    uint64       `json:"confirmations" bson:"confirmations"`
	CreationDate     time.Time    `json:"creation_date" bson:"creation_date"`
	ModifiedDate     time.Time    `json:"modified_date" bson:"modified_date"`
}

type BtcRawTransaction struct {
	Hex           string `json:"hex" bson:"hex"`
	Txid          string `json:"txid" bson:"txid"`
	Hash          string `json:"hash,omitempty" bson:"hash"`
	Size          int32  `json:"size,omitempty" bson:"size"`
	Vsize         int32  `json:"vsize,omitempty" bson:"vsize"`
	Version       int32  `json:"version" bson:"version"`
	LockTime      uint32 `json:"locktime" bson:"locktime"`
	Vin           []Vin  `json:"vin" bson:"vin"`
	Vout          []Vout `json:"vout" bson:"vout"`
	BlockHash     string `json:"blockhash,omitempty" bson:"blockhash"`
	Confirmations uint64 `json:"confirmations,omitempty" bson:"confirmations"`
	Time          int64  `json:"time,omitempty" bson:"time"`
	Blocktime     int64  `json:"blocktime,omitempty" bson:"blocktime"`
	BlockHeight   uint64 `json:"blockheight" bson:"blockheight"`
}

type Vin struct {
	Coinbase  string     `json:"coinbase" bson:"coinbase"`
	Txid      string     `json:"txid" bson:"txid"`
	Vout      uint32     `json:"vout" bson:"vout"`
	ScriptSig *ScriptSig `json:"scriptSig" bson:"scriptSig"`
	Sequence  uint32     `json:"sequence" bson:"sequence"`
	Witness   []string   `json:"txinwitness" bson:"txinwitness"`
}

type Vout struct {
	Value        float64            `json:"value" bson:"value"`
	N            uint32             `json:"n" bson:"n"`
	ScriptPubKey ScriptPubKeyResult `json:"scriptPubKey" bson:"scriptPubKey"`
}

type ScriptSig struct {
	Asm string `json:"asm" bson:"asm"`
	Hex string `json:"hex" bson:"hex"`
}

type ScriptPubKeyResult struct {
	Asm       string   `json:"asm" bson:"asm"`
	Hex       string   `json:"hex,omitempty" bson:"hex"`
	ReqSigs   int32    `json:"reqSigs,omitempty" bson:"reqSigs"`
	Type      string   `json:"type" bson:"type"`
	Addresses []string `json:"addresses,omitempty" bson:"addresses"`
}

type OmniTransaction struct {
	Txid             string `json:"txid" bson:"txid"`
	Fee              string `json:"fee" bson:"fee"`
	Sendingaddress   string `json:"sendingaddress" bson:"sendingaddress"`
	Referenceaddress string `json:"referenceaddress" bson:"referenceaddress"`
	Ismine           bool   `json:"ismine" bson:"ismine"`
	Version          int32  `json:"version" bson:"version"`
	Type_int         uint32 `json:"type_int" bson:"type_int"`
	Type             string `json:"type" bson:"type"`
	Propertyid       uint32 `json:"propertyid" bson:"propertyid"`
	Divisible        bool   `json:"divisible" bson:"divisible"`
	Amount           string `json:"amount" bson:"amount"`
	Valid            bool   `json:"valid" bson:"valid"`
	Blockhash        string `json:"blockhash" bson:"blockhash"`
	Blocktime        uint32 `json:"blocktime" bson:"blocktime"`
	Positioninblock  uint64 `json:"positioninblock" bson:"positioninblock"`
	Block            uint64 `json:"block" bson:"block"`
	Confirmations    uint64 `json:"confirmations" bson:"confirmations"`
}
