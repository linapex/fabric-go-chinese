
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455942456479744>

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

func TestSatisfied(t *testing.T) {
	var capsMap map[string]*cb.Capability
	for _, provider := range []*registry{
		NewChannelProvider(capsMap).registry,
		NewOrdererProvider(capsMap).registry,
		NewApplicationProvider(capsMap).registry,
	} {
		assert.Nil(t, provider.Supported())
	}
}

func TestNotSatisfied(t *testing.T) {
	capsMap := map[string]*cb.Capability{
		"FakeCapability": {},
	}
	for _, provider := range []*registry{
		NewChannelProvider(capsMap).registry,
		NewOrdererProvider(capsMap).registry,
		NewApplicationProvider(capsMap).registry,
	} {
		assert.Error(t, provider.Supported())
	}
}

