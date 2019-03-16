
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:23</date>
//</624456069443227648>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package common

import (
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/gossip"
)

//privdata_common保存privdata和mocks包中使用的类型。
//为了避免循环依赖性而需要

//Digkey定义了一个摘要
//指定特定的哈希rwset
type DigKey struct {
	TxId       string
	Namespace  string
	Collection string
	BlockSeq   uint64
	SeqInBlock uint64
}

type Dig2CollectionConfig map[DigKey]*common.StaticCollectionConfig

//获取pvt数据元素的dpvttatacontainer容器
//回卷人退回
type FetchedPvtDataContainer struct {
	AvailableElements []*gossip.PvtDataElement
	PurgedElements    []*gossip.PvtDataDigest
}

