
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455942737498112>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package capabilities

import (
	"testing"

	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/stretchr/testify/assert"
)

func TestOrdererV10(t *testing.T) {
	op := NewOrdererProvider(map[string]*cb.Capability{})
	assert.NoError(t, op.Supported())
	assert.False(t, op.PredictableChannelTemplate())
	assert.False(t, op.Resubmission())
	assert.False(t, op.ExpirationCheck())
}

func TestOrdererV11(t *testing.T) {
	op := NewOrdererProvider(map[string]*cb.Capability{
		OrdererV1_1: {},
	})
	assert.NoError(t, op.Supported())
	assert.True(t, op.PredictableChannelTemplate())
	assert.True(t, op.Resubmission())
	assert.True(t, op.ExpirationCheck())
}

