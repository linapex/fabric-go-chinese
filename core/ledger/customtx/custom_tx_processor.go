
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:11</date>
//</624456017698099200>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package customtx

import (
	"sync"

	"github.com/hyperledger/fabric/protos/common"
)

var processors Processors
var once sync.Once

//处理器维护自定义事务类型与其对应的Tx处理器之间的关联
type Processors map[common.HeaderType]Processor

//初始化设置自定义处理器。此函数只能在ledgermgmt.initialize（）函数期间调用。
func Initialize(customTxProcessors Processors) {
	once.Do(func() {
		initialize(customTxProcessors)
	})
}

func initialize(customTxProcessors Processors) {
	processors = customTxProcessors
}

//GetProcessor返回与txtype关联的处理器
func GetProcessor(txType common.HeaderType) Processor {
	return processors[txType]
}

