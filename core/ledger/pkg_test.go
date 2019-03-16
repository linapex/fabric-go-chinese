
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:14</date>
//</624456031887429632>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package ledger

import (
	"testing"

	"github.com/hyperledger/fabric/protos/ledger/rwset"
	"github.com/stretchr/testify/assert"
)

func TestTxPvtData(t *testing.T) {
	txPvtData := &TxPvtData{}
	assert.False(t, txPvtData.Has("ns", "coll"))

	txPvtData.WriteSet = &rwset.TxPvtReadWriteSet{
		DataModel: rwset.TxReadWriteSet_KV,
		NsPvtRwset: []*rwset.NsPvtReadWriteSet{
			{
				Namespace: "ns",
				CollectionPvtRwset: []*rwset.CollectionPvtReadWriteSet{
					{
						CollectionName: "coll-1",
						Rwset:          []byte("RandomBytes-PvtRWSet-ns1-coll1"),
					},
					{
						CollectionName: "coll-2",
						Rwset:          []byte("RandomBytes-PvtRWSet-ns1-coll2"),
					},
				},
			},
		},
	}

	assert.True(t, txPvtData.Has("ns", "coll-1"))
	assert.True(t, txPvtData.Has("ns", "coll-2"))
	assert.False(t, txPvtData.Has("ns", "coll-3"))
	assert.False(t, txPvtData.Has("ns1", "coll-1"))
}

func TestPvtNsCollFilter(t *testing.T) {
	filter := NewPvtNsCollFilter()
	filter.Add("ns", "coll-1")
	filter.Add("ns", "coll-2")
	assert.True(t, filter.Has("ns", "coll-1"))
	assert.True(t, filter.Has("ns", "coll-2"))
	assert.False(t, filter.Has("ns", "coll-3"))
	assert.False(t, filter.Has("ns1", "coll-3"))
}

