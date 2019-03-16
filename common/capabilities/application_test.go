
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455942292901888>

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

func TestApplicationV10(t *testing.T) {
	ap := NewApplicationProvider(map[string]*cb.Capability{})
	assert.NoError(t, ap.Supported())
}

func TestApplicationV11(t *testing.T) {
	ap := NewApplicationProvider(map[string]*cb.Capability{
		ApplicationV1_1: {},
	})
	assert.NoError(t, ap.Supported())
	assert.True(t, ap.ForbidDuplicateTXIdInBlock())
	assert.True(t, ap.V1_1Validation())
}

func TestApplicationV12(t *testing.T) {
	ap := NewApplicationProvider(map[string]*cb.Capability{
		ApplicationV1_2: {},
	})
	assert.NoError(t, ap.Supported())
	assert.True(t, ap.ForbidDuplicateTXIdInBlock())
	assert.True(t, ap.V1_1Validation())
	assert.True(t, ap.V1_2Validation())
	assert.True(t, ap.ACLs())
	assert.True(t, ap.CollectionUpgrade())
	assert.True(t, ap.PrivateChannelData())
}

func TestApplicationV13(t *testing.T) {
	ap := NewApplicationProvider(map[string]*cb.Capability{
		ApplicationV1_3: {},
	})
	assert.NoError(t, ap.Supported())
	assert.True(t, ap.ForbidDuplicateTXIdInBlock())
	assert.True(t, ap.V1_1Validation())
	assert.True(t, ap.V1_2Validation())
	assert.True(t, ap.V1_3Validation())
	assert.True(t, ap.KeyLevelEndorsement())
	assert.True(t, ap.ACLs())
	assert.True(t, ap.CollectionUpgrade())
	assert.True(t, ap.PrivateChannelData())
}

func TestApplicationPvtDataExperimental(t *testing.T) {
	ap := NewApplicationProvider(map[string]*cb.Capability{
		ApplicationPvtDataExperimental: {},
	})
	assert.True(t, ap.PrivateChannelData())
}

func TestFabTokenExperimental(t *testing.T) {
	ap := NewApplicationProvider(map[string]*cb.Capability{
		ApplicationFabTokenExperimental: {},
	})
	assert.True(t, ap.FabToken())
}

func TestHasCapability(t *testing.T) {
	ap := NewApplicationProvider(map[string]*cb.Capability{})
	assert.True(t, ap.HasCapability(ApplicationV1_1))
	assert.True(t, ap.HasCapability(ApplicationV1_2))
	assert.True(t, ap.HasCapability(ApplicationV1_3))
	assert.True(t, ap.HasCapability(ApplicationPvtDataExperimental))
	assert.True(t, ap.HasCapability(ApplicationResourcesTreeExperimental))
	assert.False(t, ap.HasCapability("default"))
}

