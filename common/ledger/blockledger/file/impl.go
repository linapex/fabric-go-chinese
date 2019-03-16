
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455959657320448>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package fileledger

import (
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/ledger"
	"github.com/hyperledger/fabric/common/ledger/blockledger"
	cb "github.com/hyperledger/fabric/protos/common"
	ab "github.com/hyperledger/fabric/protos/orderer"
)

var logger = flogging.MustGetLogger("common.ledger.blockledger.file")

//
type FileLedger struct {
	blockStore FileLedgerBlockStore
	signal     chan struct{}
}

//
//
type FileLedgerBlockStore interface {
	AddBlock(block *cb.Block) error
	GetBlockchainInfo() (*cb.BlockchainInfo, error)
	RetrieveBlocks(startBlockNumber uint64) (ledger.ResultsIterator, error)
}

//
func NewFileLedger(blockStore FileLedgerBlockStore) *FileLedger {
	return &FileLedger{blockStore: blockStore, signal: make(chan struct{})}
}

type fileLedgerIterator struct {
	ledger         *FileLedger
	blockNumber    uint64
	commonIterator ledger.ResultsIterator
}

//
//
func (i *fileLedgerIterator) Next() (*cb.Block, cb.Status) {
	result, err := i.commonIterator.Next()
	if err != nil {
		logger.Error(err)
		return nil, cb.Status_SERVICE_UNAVAILABLE
	}
//覆盖另一个线程在迭代器上调用close的情况。
	if result == nil {
		return nil, cb.Status_SERVICE_UNAVAILABLE
	}
	return result.(*cb.Block), cb.Status_SUCCESS
}

//Close releases resources acquired by the Iterator
func (i *fileLedgerIterator) Close() {
	i.commonIterator.Close()
}

//Iterator returns an Iterator, as specified by an ab.SeekInfo message, and its
//起始块编号
func (fl *FileLedger) Iterator(startPosition *ab.SeekPosition) (blockledger.Iterator, uint64) {
	var startingBlockNumber uint64
	switch start := startPosition.Type.(type) {
	case *ab.SeekPosition_Oldest:
		startingBlockNumber = 0
	case *ab.SeekPosition_Newest:
		info, err := fl.blockStore.GetBlockchainInfo()
		if err != nil {
			logger.Panic(err)
		}
		newestBlockNumber := info.Height - 1
		startingBlockNumber = newestBlockNumber
	case *ab.SeekPosition_Specified:
		startingBlockNumber = start.Specified.Number
		height := fl.Height()
		if startingBlockNumber > height {
			return &blockledger.NotFoundErrorIterator{}, 0
		}
	default:
		return &blockledger.NotFoundErrorIterator{}, 0
	}

	iterator, err := fl.blockStore.RetrieveBlocks(startingBlockNumber)
	if err != nil {
		return &blockledger.NotFoundErrorIterator{}, 0
	}

	return &fileLedgerIterator{ledger: fl, blockNumber: startingBlockNumber, commonIterator: iterator}, startingBlockNumber
}

//height返回分类帐上的块数
func (fl *FileLedger) Height() uint64 {
	info, err := fl.blockStore.GetBlockchainInfo()
	if err != nil {
		logger.Panic(err)
	}
	return info.Height
}

//将新块追加到分类帐
func (fl *FileLedger) Append(block *cb.Block) error {
	err := fl.blockStore.AddBlock(block)
	if err == nil {
		close(fl.signal)
		fl.signal = make(chan struct{})
	}
	return err
}

