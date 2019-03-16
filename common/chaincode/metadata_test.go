
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455944687849472>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package chaincode

import (
	"testing"

	"github.com/hyperledger/fabric/protos/gossip"
	"github.com/stretchr/testify/assert"
)

func TestToChaincodes(t *testing.T) {
	ccs := MetadataSet{
		{
			Name:    "foo",
			Version: "1.0",
		},
	}
	assert.Equal(t, []*gossip.Chaincode{
		{Name: "foo", Version: "1.0"},
	}, ccs.AsChaincodes())
}

func TestMetadataMapping(t *testing.T) {
	mm := NewMetadataMapping()
	md1 := Metadata{
		Name:    "cc1",
		Id:      []byte{1},
		Version: "1.0",
		Policy:  []byte{1, 2, 3},
	}
	mm.Update(md1)
	res, found := mm.Lookup("cc1")
	assert.Equal(t, md1, res)
	assert.True(t, found)
	res, found = mm.Lookup("cc2")
	assert.Zero(t, res)
	assert.False(t, found)
	md2 := Metadata{
		Name:    "cc1",
		Id:      []byte{1},
		Version: "1.1",
		Policy:  []byte{2, 2, 2},
	}
	mm.Update(md2)
	res, found = mm.Lookup("cc1")
	assert.Equal(t, md2, res)

	assert.Equal(t, MetadataSet{md2}, mm.Aggregate())
}

