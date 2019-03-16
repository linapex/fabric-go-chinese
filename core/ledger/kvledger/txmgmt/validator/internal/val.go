
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456028209025024>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package internal

import (
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/privacyenabledstate"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = flogging.MustGetLogger("valinternal")

//验证程序应该基于块中存在的公共数据和哈希来验证事务。
//并返回一个应用于更新状态的批
type Validator interface {
	ValidateAndPrepareBatch(block *Block, doMVCCValidation bool) (*PubAndHashUpdates, error)
}

//块用于将信息从其原始格式保存到结构中。
//更适合/更友好地进行验证
type Block struct {
	Num uint64
	Txs []*Transaction
}

//事务用于将信息从其协议格式保存到结构中。
//更适合/更友好地进行验证
type Transaction struct {
	IndexInBlock   int
	ID             string
	RWSet          *rwsetutil.TxRwSet
	ValidationCode peer.TxValidationCode
}

//pubandhashupdates封装公共和哈希更新。用于保存更新的预期用途
//将作为块提交的结果应用于statedb
type PubAndHashUpdates struct {
	PubUpdates  *privacyenabledstate.PubUpdateBatch
	HashUpdates *privacyenabledstate.HashedUpdateBatch
}

//NewPubandHashUpdates构造空的PubandHashUpdates
func NewPubAndHashUpdates() *PubAndHashUpdates {
	return &PubAndHashUpdates{
		privacyenabledstate.NewPubUpdateBatch(),
		privacyenabledstate.NewHashedUpdateBatch(),
	}
}

//如果此事务不限于仅影响公共数据，则containsPvTwrites返回true
func (t *Transaction) ContainsPvtWrites() bool {
	for _, ns := range t.RWSet.NsRwSets {
		for _, coll := range ns.CollHashedRwSets {
			if coll.PvtRwSetHash != nil {
				return true
			}
		}
	}
	return false
}

//RetrieveHash返回存在的私有写入集的哈希值
//在给定命名空间集合的公共数据中
func (t *Transaction) RetrieveHash(ns string, coll string) []byte {
	if t.RWSet == nil {
		return nil
	}
	for _, nsData := range t.RWSet.NsRwSets {
		if nsData.NameSpace != ns {
			continue
		}

		for _, collData := range nsData.CollHashedRwSets {
			if collData.CollectionName == coll {
				return collData.PvtRwSetHash
			}
		}
	}
	return nil
}

//applyWriteset向pubandhashupdates添加（或删除）写入集中存在的键/值
func (u *PubAndHashUpdates) ApplyWriteSet(txRWSet *rwsetutil.TxRwSet, txHeight *version.Height, db privacyenabledstate.DB) error {
	txops, err := prepareTxOps(txRWSet, txHeight, u, db)
	logger.Debugf("txops=%#v", txops)
	if err != nil {
		return err
	}
	for compositeKey, keyops := range txops {
		if compositeKey.coll == "" {
			ns, key := compositeKey.ns, compositeKey.key
			if keyops.isDelete() {
				u.PubUpdates.Delete(ns, key, txHeight)
			} else {
				u.PubUpdates.PutValAndMetadata(ns, key, keyops.value, keyops.metadata, txHeight)
			}
		} else {
			ns, coll, keyHash := compositeKey.ns, compositeKey.coll, []byte(compositeKey.key)
			if keyops.isDelete() {
				u.HashUpdates.Delete(ns, coll, keyHash, txHeight)
			} else {
				u.HashUpdates.PutValHashAndMetadata(ns, coll, keyHash, keyops.value, keyops.metadata, txHeight)
			}
		}
	}
	return nil
}

