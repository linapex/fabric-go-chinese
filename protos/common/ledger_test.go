
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:34</date>
//</624456114775265280>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLedger(t *testing.T) {
	var info *BlockchainInfo
	info = nil
	assert.Equal(t, uint64(0), info.GetHeight())
	assert.Nil(t, info.GetCurrentBlockHash())
	assert.Nil(t, info.GetPreviousBlockHash())
	info = &BlockchainInfo{
		Height:            uint64(1),
		CurrentBlockHash:  []byte("blockhash"),
		PreviousBlockHash: []byte("previoushash"),
	}
	assert.Equal(t, uint64(1), info.GetHeight())
	assert.NotNil(t, info.GetCurrentBlockHash())
	assert.NotNil(t, info.GetPreviousBlockHash())
	info.Reset()
	assert.Equal(t, uint64(0), info.GetHeight())
	_ = info.String()
	_, _ = info.Descriptor()
	info.ProtoMessage()
}

