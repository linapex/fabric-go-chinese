
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:29</date>
//</624456094609051648>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package msgprocessor

import (
	"testing"

	"github.com/golang/protobuf/proto"
	mockconfig "github.com/hyperledger/fabric/common/mocks/config"
	cb "github.com/hyperledger/fabric/protos/common"
	ab "github.com/hyperledger/fabric/protos/orderer"
	"github.com/stretchr/testify/assert"
)

func TestMaxBytesRule(t *testing.T) {
	dataSize := uint32(100)
	maxBytes := calcMessageBytesForPayloadDataSize(dataSize)
	msf := NewSizeFilter(&mockconfig.Orderer{BatchSizeVal: &ab.BatchSize{AbsoluteMaxBytes: maxBytes}})

	t.Run("LessThan", func(t *testing.T) {
		assert.Nil(t, msf.Apply(makeMessage(make([]byte, dataSize-1))))
	})
	t.Run("Exact", func(t *testing.T) {
		assert.Nil(t, msf.Apply(makeMessage(make([]byte, dataSize))))
	})
	t.Run("TooBig", func(t *testing.T) {
		assert.NotNil(t, msf.Apply(makeMessage(make([]byte, dataSize+1))))
	})
}

func calcMessageBytesForPayloadDataSize(dataSize uint32) uint32 {
	return messageByteSize(makeMessage(make([]byte, dataSize)))
}

func makeMessage(data []byte) *cb.Envelope {
	data, err := proto.Marshal(&cb.Payload{Data: data})
	if err != nil {
		panic(err)
	}
	return &cb.Envelope{Payload: data}
}

