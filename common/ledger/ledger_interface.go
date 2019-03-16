
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455961293099008>

/*
版权所有IBM Corp.2016保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。


   


根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
*/


package ledger

import (
	"github.com/hyperledger/fabric/protos/common"
)

//
type Ledger interface {
//
	GetBlockchainInfo() (*common.BlockchainInfo, error)
//
//
	GetBlockByNumber(blockNumber uint64) (*common.Block, error)
//
//迭代器是一个阻塞迭代器，也就是说，它阻塞直到下一个块在分类帐中可用为止。
//resultsiterator包含类型blockholder
	GetBlocksIterator(startBlockNumber uint64) (ResultsIterator, error)
//关闭关闭分类帐
	Close()
}

//resultsiterator-查询结果集的迭代器
type ResultsIterator interface {
//Next返回结果集中的下一项。当
//迭代器耗尽
	Next() (QueryResult, error)
//close释放迭代器占用的资源
	Close()
}

//queryresultsiterator-查询结果集的迭代器
type QueryResultsIterator interface {
	ResultsIterator
	GetBookmarkAndClose() string
}

//queryresult-支持不同类型查询结果的通用接口。不同查询的实际类型不同
type QueryResult interface{}

//prunepolicy-支持不同修剪策略的通用接口
type PrunePolicy interface{}

