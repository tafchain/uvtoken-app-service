package service

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/big"
	"service/server/global"
	"service/server/model"
	"service/server/model/request"
	"service/server/model/response"
)

func GetBtcHash(reqInfo request.Hashearch) (error, interface{}) {
	var err error
	var hashs response.Hashs

	return err, hashs
}

func getBtcRawTx(txid string, btcCollection *mongo.Collection) (error, *model.BtcRawTransaction) {
	var rawTx model.BtcRawTransaction

	filter := bson.M{
		"txid": txid,
	}
	one := btcCollection.FindOne(context.TODO(), filter)
	err := one.Decode(&rawTx)
	return err, &rawTx
}

func GetBtcTransaction(reqInfo request.TxSearch) (error, interface{}) {
	var err error
	var btcTx response.BtcTx
	var raws response.BtcRawTxs
	log.Println(reqInfo)
	if reqInfo.CoinType == "BTC" {
		btcCollection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.DataBase).Collection("btc_raw_transaction")

		findOpts := options.Find().SetSort(bson.D{{"blockheight", 1}})
		filter := bson.M{
			"txid": bson.M{"$in": reqInfo.Hash},
		}
		findCursor, err := btcCollection.Find(context.TODO(), filter, findOpts)
		if err != nil {
			return err, nil
		}
		defer findCursor.Close(context.TODO())

		if err = findCursor.All(context.TODO(), &raws.Txs); err != nil {
			return err, nil
		}
		log.Println("raws.Txs ", raws.Txs)

	next_tx:
		for _, tx := range raws.Txs {
			r := response.BtcTxResp{Hex: tx.Hex, Txid: tx.Txid, Hash: tx.Hash, Size: tx.Size, Vsize: tx.Vsize,
				Version: tx.Version, LockTime: tx.LockTime, BlockHash: tx.BlockHash,
				Time: tx.Time, Blocktime: tx.Blocktime, BlockHeight: tx.BlockHeight}
			mAddr := make(map[string]*big.Rat)
			for _, v := range tx.Vin {
				if v.Coinbase != "" {
					continue
				}
				err, rawTx := getBtcRawTx(v.Txid, btcCollection)
				if err != nil {
					log.Println(err, ",txid=", v.Txid)
					continue next_tx
					//return err, nil
				}
				for _, out := range rawTx.Vout {
					fAddresses := out.ScriptPubKey.Addresses
					if len(fAddresses) > 0 && out.N == v.Vout {
						fAddr := fAddresses[0]
						amountRat := new(big.Rat).SetFloat64(out.Value)
						if _, ok := mAddr[fAddr]; !ok {
							mAddr[fAddr] = amountRat
						} else {
							mAddr[fAddr] = mAddr[fAddr].Add(mAddr[fAddr], amountRat)
						}

						break
					}
				}
			}
			for k, v := range mAddr {
				r.FromAddresses = append(r.FromAddresses, response.AddressAmount{Address: k, Amount: v.FloatString(8)})
			}
			mToAddr := make(map[string]*big.Rat)
			for _, v := range tx.Vout {
				amountRat := new(big.Rat).SetFloat64(v.Value)
				tAddresses := v.ScriptPubKey.Addresses
				if len(tAddresses) > 0 {
					tAddr := tAddresses[0]
					if _, ok := mToAddr[tAddr]; !ok {
						mToAddr[tAddr] = amountRat
					} else {
						mToAddr[tAddr] = mToAddr[tAddr].Add(mToAddr[tAddr], amountRat)
					}
				}
			}
			for k, v := range mToAddr {
				r.ToAddresses = append(r.ToAddresses, response.AddressAmount{Address: k, Amount: v.FloatString(8)})
			}
			btcTx.TxHashs = append(btcTx.TxHashs, r)
		}
		return nil, btcTx
	} else if reqInfo.CoinType == "USDT_OMNI" {
		var raws response.BtcOmniTxs
		omniCollection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.DataBase).Collection("btc_omni_transaction")

		findOpts := options.Find().SetSort(bson.D{{"block", 1}})
		filter := bson.M{
			"txid": bson.M{"$in": reqInfo.Hash},
		}
		findCursor, err := omniCollection.Find(context.TODO(), filter, findOpts)
		if err != nil {
			return err, nil
		}
		defer findCursor.Close(context.TODO())

		if err = findCursor.All(context.TODO(), &raws.Txs); err != nil {
			return err, nil
		}

		for _, tx := range raws.Txs {
			r := response.BtcTxResp{Txid: tx.Txid, Version: tx.Version, BlockHash: tx.Blockhash, Blocktime: int64(tx.Blocktime), BlockHeight: tx.Block,
				OmniResp: &response.OmniResp{Fee: tx.Fee, PropertyId: tx.Propertyid, Valid: tx.Valid, Divisible: tx.Divisible}}
			r.FromAddresses = []response.AddressAmount{{Address: tx.Sendingaddress, Amount: tx.Amount}}
			r.ToAddresses = []response.AddressAmount{{Address: tx.Referenceaddress, Amount: tx.Amount}}
			btcTx.TxHashs = append(btcTx.TxHashs, r)
		}
		return nil, btcTx
	}
	return err, ""
}

func GetEthHash(reqInfo request.Hashearch) (error, interface{}) {
	var err error
	var hashs response.Hashs

	ethCollection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.DataBase).Collection("eth_transaction")
	var results []response.TxHash
	var findCursor *mongo.Cursor

	findOpts := options.Find().SetSort(bson.D{{"blockNumber", 1}})
	filter := bson.M{}

	var confirmations uint64 = 6
	if reqInfo.Confirmations != 0 {
		confirmations = reqInfo.Confirmations
	}

	if reqInfo.CoinType == "ETH" {
		if reqInfo.ToBlock == 0 {
			filter = bson.M{
				"blockNumber": bson.M{
					"$gte": reqInfo.FromBlock,
				},
				//"value_": bson.M{"$ne": "0"},
				"$or": bson.A{
					bson.M{"from_": reqInfo.Address},
					bson.M{"to_": reqInfo.Address},
				},
				"confirmations": bson.M{"$gte": confirmations},
			}
		} else {
			filter = bson.M{
				"blockNumber": bson.M{
					"$lte": reqInfo.ToBlock,
					"$gte": reqInfo.FromBlock,
				},
				//"value_": bson.M{"$ne": "0"},
				"$or": bson.A{
					bson.M{"from_": reqInfo.Address},
					bson.M{"to_": reqInfo.Address},
				},
				"confirmations": bson.M{"$gte": confirmations},
			}
		}
	} else {
		if reqInfo.ToBlock == 0 {
			filter = bson.M{
				"blockNumber": bson.M{
					"$gte": reqInfo.FromBlock,
				},
				"to_": reqInfo.TokenAddress,
				"$or": bson.A{
					bson.M{
						"inputDecode._from": reqInfo.Address,
					},
					bson.M{
						"inputDecode._to": reqInfo.Address,
					},
				},
				"confirmations": bson.M{"$gte": confirmations},
			}
		} else {
			filter = bson.M{
				"blockNumber": bson.M{
					"$lte": reqInfo.ToBlock,
					"$gte": reqInfo.FromBlock,
				},
				"to_": reqInfo.TokenAddress,
				"$or": bson.A{
					bson.M{
						"inputDecode._from": reqInfo.Address,
					},
					bson.M{
						"inputDecode._to": reqInfo.Address,
					},
				},
				"confirmations": bson.M{"$gte": confirmations},
			}
		}
	}
	findCursor, err = ethCollection.Find(context.TODO(), filter, findOpts)
	if err != nil {
		return err, nil
	}
	defer findCursor.Close(context.TODO())

	if err = findCursor.All(context.TODO(), &results); err != nil {
		return err, nil
	}

	// removeDuplicate
	resultMap := map[string]bool{}
	for _, v := range results {
		data, _ := json.Marshal(v)
		resultMap[string(data)] = true
	}
	var result []response.TxHash
	for k := range resultMap {
		var t response.TxHash
		_ = json.Unmarshal([]byte(k), &t)
		result = append(result, t)
	}

	hashs.TxHashs = result

	return nil, hashs

}

func GetEthTransaction(reqInfo request.TxSearch) (error, interface{}) {
	var err error
	var ethTx response.EthTx
	ethCollection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.DataBase).Collection("eth_transaction")

	findOpts := options.Find().SetSort(bson.D{{"blockNumber", 1}})
	filter := bson.M{
		"hash": bson.M{"$in": reqInfo.Hash},
	}
	findCursor, err := ethCollection.Find(context.TODO(), filter, findOpts)
	if err != nil {
		return err, nil
	}
	defer findCursor.Close(context.TODO())

	if err = findCursor.All(context.TODO(), &ethTx.TxHashs); err != nil {
		return err, nil
	}

	return nil, ethTx
}

func GetEthHashV2(reqInfo request.Hashearch) (error, interface{}) {
	var err error
	var hashs response.Hashs

	ethCollection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.DataBase).Collection("eth_transaction_v2")
	var results []response.TxHash
	var findCursor *mongo.Cursor

	findOpts := options.Find().SetSort(bson.D{{"blockNumber", 1}})
	filter := bson.M{}

	var confirmations uint64 = 6
	if reqInfo.Confirmations != 0 {
		confirmations = reqInfo.Confirmations
	}
	var toBlock uint64 = 2147483647
	if reqInfo.ToBlock != 0 {
		toBlock = reqInfo.ToBlock
	}

	if reqInfo.CoinType == "ETH" {
		if reqInfo.ToBlock == 0 {
			filter = bson.M{
				"blockNumber": bson.M{
					"$lte": toBlock,
					"$gte": reqInfo.FromBlock,
				},
				"$or": bson.A{
					bson.M{"from_": reqInfo.Address},
					bson.M{"to_": reqInfo.Address},
				},
				"confirmations": bson.M{"$gte": confirmations},
			}
		}
	} else {
		filter = bson.M{
			"blockNumber": bson.M{
				"$lte": toBlock,
				"$gte": reqInfo.FromBlock,
			},
			"logs._address": reqInfo.TokenAddress,
			"$or": bson.A{
				bson.M{
					"logs._from": reqInfo.Address,
				},
				bson.M{
					"logs._to": reqInfo.Address,
				},
			},
			"confirmations": bson.M{"$gte": confirmations},
		}
	}
	findCursor, err = ethCollection.Find(context.TODO(), filter, findOpts)
	if err != nil {
		return err, nil
	}
	defer findCursor.Close(context.TODO())

	if err = findCursor.All(context.TODO(), &results); err != nil {
		return err, nil
	}

	hashs.TxHashs = results

	return nil, hashs

}

func GetEthTransactionV2(reqInfo request.TxSearch) (error, interface{}) {
	var err error
	var ethTx response.EthTxV2
	ethCollection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.DataBase).Collection("eth_transaction_v2")

	findOpts := options.Find().SetSort(bson.D{{"blockNumber", 1}})
	filter := bson.M{
		"hash": bson.M{"$in": reqInfo.Hash},
	}
	findCursor, err := ethCollection.Find(context.TODO(), filter, findOpts)
	if err != nil {
		return err, nil
	}
	defer findCursor.Close(context.TODO())

	if err = findCursor.All(context.TODO(), &ethTx.TxHashs); err != nil {
		return err, nil
	}

	return nil, ethTx
}
