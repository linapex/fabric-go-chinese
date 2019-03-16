
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:35</date>
//</624456118273314816>

/*
版权所有IBM公司。保留所有权利。
SPDX许可证标识符：Apache-2.0
**/

package rwset

import (
	fmt "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTxPvtRwsetTrim(t *testing.T) {
	txpvtrwset := testutilConstructSampleTxPvtRwset(
		[]*testNsColls{
			{ns: "ns-1", colls: []string{"coll-1", "coll-2"}},
			{ns: "ns-2", colls: []string{"coll-3", "coll-4"}},
		},
	)

	txpvtrwset.Remove("ns-1", "coll-1")
	assert.Equal(
		t,
		testutilConstructSampleTxPvtRwset(
			[]*testNsColls{
				{ns: "ns-1", colls: []string{"coll-2"}},
				{ns: "ns-2", colls: []string{"coll-3", "coll-4"}},
			},
		),
		txpvtrwset,
	)

	txpvtrwset.Remove("ns-1", "coll-2")
	assert.Equal(
		t,
		testutilConstructSampleTxPvtRwset(
			[]*testNsColls{
				{ns: "ns-2", colls: []string{"coll-3", "coll-4"}},
			},
		),
		txpvtrwset,
	)
}

func testutilConstructSampleTxPvtRwset(nsCollsList []*testNsColls) *TxPvtReadWriteSet {
	txPvtRwset := &TxPvtReadWriteSet{}
	for _, nsColls := range nsCollsList {
		ns := nsColls.ns
		nsdata := &NsPvtReadWriteSet{
			Namespace:          ns,
			CollectionPvtRwset: []*CollectionPvtReadWriteSet{},
		}
		txPvtRwset.NsPvtRwset = append(txPvtRwset.NsPvtRwset, nsdata)
		for _, coll := range nsColls.colls {
			nsdata.CollectionPvtRwset = append(nsdata.CollectionPvtRwset,
				&CollectionPvtReadWriteSet{
					CollectionName: coll,
					Rwset:          []byte(fmt.Sprintf("pvtrwset-for-%s-%s", ns, coll)),
				},
			)
		}
	}
	return txPvtRwset
}

type testNsColls struct {
	ns    string
	colls []string
}

