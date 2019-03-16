
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456011792519168>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package endorsement

import (
	"github.com/hyperledger/fabric/core/handlers/endorsement/api"
	"github.com/hyperledger/fabric/protos/ledger/rwset"
)

//国家定义与世界国家的互动
type State interface {
//getprivatedatamultiplekeys获取单个调用中多个私有数据项的值
	GetPrivateDataMultipleKeys(namespace, collection string, keys []string) ([][]byte, error)

//GetStateMultipleKeys获取单个调用中多个键的值
	GetStateMultipleKeys(namespace string, keys []string) ([][]byte, error)

//getTransientByXid获取与给定txID关联的私有数据值
	GetTransientByTXID(txID string) ([]*rwset.TxPvtReadWriteSet, error)

//完成释放国家占用的资源
	Done()
}

//StateFetcher检索状态的实例
type StateFetcher interface {
	endorsement.Dependency

//FetchState获取状态
	FetchState() (State, error)
}

