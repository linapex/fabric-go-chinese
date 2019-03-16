
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456033829392384>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package pvtdatastorage

import (
	"bytes"
	math "math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataKeyEncoding(t *testing.T) {
	dataKey1 := &dataKey{nsCollBlk: nsCollBlk{ns: "ns1", coll: "coll1", blkNum: 2}, txNum: 5}
	datakey2 := decodeDatakey(encodeDataKey(dataKey1))
	assert.Equal(t, dataKey1, datakey2)
}

func TestDatakeyRange(t *testing.T) {
	blockNum := uint64(20)
	startKey, endKey := datakeyRange(blockNum)
	var txNum uint64
	for txNum = 0; txNum < 100; txNum++ {
		keyOfBlock := encodeDataKey(
			&dataKey{
				nsCollBlk: nsCollBlk{ns: "ns", coll: "coll", blkNum: blockNum},
				txNum:     txNum,
			},
		)
		keyOfPreviousBlock := encodeDataKey(
			&dataKey{
				nsCollBlk: nsCollBlk{ns: "ns", coll: "coll", blkNum: blockNum - 1},
				txNum:     txNum,
			},
		)
		keyOfNextBlock := encodeDataKey(
			&dataKey{
				nsCollBlk: nsCollBlk{ns: "ns", coll: "coll", blkNum: blockNum + 1},
				txNum:     txNum,
			},
		)
		assert.Equal(t, bytes.Compare(keyOfPreviousBlock, startKey), -1)
		assert.Equal(t, bytes.Compare(keyOfBlock, startKey), 1)
		assert.Equal(t, bytes.Compare(keyOfBlock, endKey), -1)
		assert.Equal(t, bytes.Compare(keyOfNextBlock, endKey), 1)
	}
}

func TestEligibleMissingdataRange(t *testing.T) {
	blockNum := uint64(20)
	startKey, endKey := eligibleMissingdatakeyRange(blockNum)
	var txNum uint64
	for txNum = 0; txNum < 100; txNum++ {
		keyOfBlock := encodeMissingDataKey(
			&missingDataKey{
				nsCollBlk:  nsCollBlk{ns: "ns", coll: "coll", blkNum: blockNum},
				isEligible: true,
			},
		)
		keyOfPreviousBlock := encodeMissingDataKey(
			&missingDataKey{
				nsCollBlk:  nsCollBlk{ns: "ns", coll: "coll", blkNum: blockNum - 1},
				isEligible: true,
			},
		)
		keyOfNextBlock := encodeMissingDataKey(
			&missingDataKey{
				nsCollBlk:  nsCollBlk{ns: "ns", coll: "coll", blkNum: blockNum + 1},
				isEligible: true,
			},
		)
		assert.Equal(t, bytes.Compare(keyOfNextBlock, startKey), -1)
		assert.Equal(t, bytes.Compare(keyOfBlock, startKey), 1)
		assert.Equal(t, bytes.Compare(keyOfBlock, endKey), -1)
		assert.Equal(t, bytes.Compare(keyOfPreviousBlock, endKey), 1)
	}
}

func TestEncodeDecodeMissingdataKey(t *testing.T) {
	for i := 0; i < 1000; i++ {
		testEncodeDecodeMissingdataKey(t, uint64(i))
	}
testEncodeDecodeMissingdataKey(t, math.MaxUint64) //角箱
}

func testEncodeDecodeMissingdataKey(t *testing.T, blkNum uint64) {
	key := &missingDataKey{
		nsCollBlk: nsCollBlk{
			ns:     "ns",
			coll:   "coll",
			blkNum: blkNum,
		},
	}

	t.Run("ineligibileKey",
		func(t *testing.T) {
			key.isEligible = false
			decodedKey := decodeMissingDataKey(
				encodeMissingDataKey(key),
			)
			assert.Equal(t, key, decodedKey)
		},
	)

	t.Run("ineligibileKey",
		func(t *testing.T) {
			key.isEligible = true
			decodedKey := decodeMissingDataKey(
				encodeMissingDataKey(key),
			)
			assert.Equal(t, key, decodedKey)
		},
	)
}

