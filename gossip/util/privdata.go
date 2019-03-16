
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:24</date>
//</624456073037746176>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package util

import (
	"encoding/hex"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/gossip"
	"github.com/hyperledger/fabric/protos/ledger/rwset"
	"github.com/pkg/errors"
)

//用于封装集合的pvtDataCollections数据类型
//私人数据
type PvtDataCollections []*ledger.TxPvtData

//封送将私有集合编码为字节数组
func (pvt *PvtDataCollections) Marshal() ([][]byte, error) {
	pvtDataBytes := make([][]byte, 0)
	for index, each := range *pvt {
		if each == nil {
			errMsg := fmt.Sprintf("Mallformed private data payload, rwset index %d is nil", index)
			return nil, errors.New(errMsg)
		}
		pvtBytes, err := proto.Marshal(each.WriteSet)
		if err != nil {
			errMsg := fmt.Sprintf("Could not marshal private rwset index %d, due to %s", index, err)
			return nil, errors.New(errMsg)
		}
//使用块中的private rwset+事务索引编写八卦协议消息
		txSeqInBlock := each.SeqInBlock
		pvtDataPayload := &gossip.PvtDataPayload{TxSeqInBlock: txSeqInBlock, Payload: pvtBytes}
		payloadBytes, err := proto.Marshal(pvtDataPayload)
		if err != nil {
			errMsg := fmt.Sprintf("Could not marshal private payload with transaction index %d, due to %s", txSeqInBlock, err)
			return nil, errors.New(errMsg)
		}

		pvtDataBytes = append(pvtDataBytes, payloadBytes)
	}
	return pvtDataBytes, nil
}

//解组读取和解组收集私人数据
//从给定的字节数组
func (pvt *PvtDataCollections) Unmarshal(data [][]byte) error {
	for _, each := range data {
		payload := &gossip.PvtDataPayload{}
		if err := proto.Unmarshal(each, payload); err != nil {
			return err
		}
		pvtRWSet := &rwset.TxPvtReadWriteSet{}
		if err := proto.Unmarshal(payload.Payload, pvtRWSet); err != nil {
			return err
		}
		*pvt = append(*pvt, &ledger.TxPvtData{
			SeqInBlock: payload.TxSeqInBlock,
			WriteSet:   pvtRWSet,
		})
	}

	return nil
}

//privaterwset创建rwset的聚合切片
func PrivateRWSets(rwsets ...PrivateRWSet) [][]byte {
	var res [][]byte
	for _, rws := range rwsets {
		res = append(res, []byte(rws))
	}
	return res
}

//privaterwset包含collectionpvtreadwriteset的字节
type PrivateRWSet []byte

//Digest返回privaterwset的确定性和无冲突表示
func (rws PrivateRWSet) Digest() string {
	return hex.EncodeToString(util.ComputeSHA256(rws))
}

//privaterwsetwithconfig封装了私有读写集
//其中与集合相关的配置信息
type PrivateRWSetWithConfig struct {
	RWSet            []PrivateRWSet
	CollectionConfig *common.CollectionConfig
}

