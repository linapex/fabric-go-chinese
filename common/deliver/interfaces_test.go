
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:55</date>
//</624455950136250368>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package deliver_test

import (
	"github.com/hyperledger/fabric/common/ledger/blockledger"
)

//go：生成仿冒者-o mock/block\u reader.go-fake name block reader。拦截器
type blockledgerReader interface {
	blockledger.Reader
}

//go：生成仿冒者-o mock/block_iterator.go-fake name block iterator。拦截器
type blockledgerIterator interface {
	blockledger.Iterator
}

