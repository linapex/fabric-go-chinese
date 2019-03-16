
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456025507893248>

/*
版权所有IBM公司。保留所有权利。
SPDX许可证标识符：Apache-2.0
**/


package statecouchdb

import (
	"testing"

	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
	"github.com/stretchr/testify/assert"
)

func TestVersionCache(t *testing.T) {
	verCache := newVersionCache()
	ver1 := version.NewHeight(1, 1)
	ver2 := version.NewHeight(2, 2)
	verCache.setVerAndRev("ns1", "key1", version.NewHeight(1, 1), "rev1")
	verCache.setVerAndRev("ns2", "key2", version.NewHeight(2, 2), "rev2")

	ver, found := verCache.getVersion("ns1", "key1")
	assert.True(t, found)
	assert.Equal(t, ver1, ver)

	ver, found = verCache.getVersion("ns2", "key2")
	assert.True(t, found)
	assert.Equal(t, ver2, ver)

	ver, found = verCache.getVersion("ns1", "key3")
	assert.False(t, found)
	assert.Nil(t, ver)

	ver, found = verCache.getVersion("ns3", "key4")
	assert.False(t, found)
	assert.Nil(t, ver)
}

