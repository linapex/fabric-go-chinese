
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455960605233152>

/*
版权所有IBM Corp.2017保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

                 http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
*/


package ramledger

import (
	"sync"

	"github.com/hyperledger/fabric/common/ledger/blockledger"
	cb "github.com/hyperledger/fabric/protos/common"
)

type ramLedgerFactory struct {
	maxSize int
	ledgers map[string]blockledger.ReadWriter
	mutex   sync.Mutex
}

//getorcreate获取现有分类帐（如果存在），或者如果不存在则创建该分类帐。
func (rlf *ramLedgerFactory) GetOrCreate(chainID string) (blockledger.ReadWriter, error) {
	rlf.mutex.Lock()
	defer rlf.mutex.Unlock()

	key := chainID

	l, ok := rlf.ledgers[key]
	if ok {
		return l, nil
	}

	ch := newChain(rlf.maxSize)
	rlf.ledgers[key] = ch
	return ch, nil
}

//
func newChain(maxSize int) blockledger.ReadWriter {
	preGenesis := &cb.Block{
		Header: &cb.BlockHeader{
			Number: ^uint64(0),
		},
	}

	rl := &ramLedger{
		maxSize: maxSize,
		size:    1,
		oldest: &simpleList{
			signal: make(chan struct{}),
			block:  preGenesis,
		},
	}
	rl.newest = rl.oldest
	return rl
}

//chainIds返回工厂知道的链ID
func (rlf *ramLedgerFactory) ChainIDs() []string {
	rlf.mutex.Lock()
	defer rlf.mutex.Unlock()
	ids := make([]string, len(rlf.ledgers))

	i := 0
	for key := range rlf.ledgers {
		ids[i] = key
		i++
	}

	return ids
}

//
func (rlf *ramLedgerFactory) Close() {
return //无事可做
}

//新建创建新的分类帐工厂
func New(maxSize int) blockledger.Factory {
	rlf := &ramLedgerFactory{
		maxSize: maxSize,
		ledgers: make(map[string]blockledger.ReadWriter),
	}

	return rlf
}

