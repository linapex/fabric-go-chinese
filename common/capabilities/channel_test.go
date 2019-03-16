
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455942594891776>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package capabilities

import (
	"testing"

	"github.com/hyperledger/fabric/msp"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/stretchr/testify/assert"
)

func TestChannelV10(t *testing.T) {
	op := NewChannelProvider(map[string]*cb.Capability{})
	assert.NoError(t, op.Supported())
	assert.True(t, op.MSPVersion() == msp.MSPv1_0)
}

func TestChannelV11(t *testing.T) {
	op := NewChannelProvider(map[string]*cb.Capability{
		ChannelV1_1: {},
	})
	assert.NoError(t, op.Supported())
	assert.True(t, op.MSPVersion() == msp.MSPv1_1)
}

func TestChannelV13(t *testing.T) {
	op := NewChannelProvider(map[string]*cb.Capability{
		ChannelV1_1: {},
		ChannelV1_3: {},
	})
	assert.NoError(t, op.Supported())
	assert.True(t, op.MSPVersion() == msp.MSPv1_3)
}

