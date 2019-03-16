
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456025663082496>

/*
版权所有IBM公司。保留所有权利。
SPDX许可证标识符：Apache-2.0
**/


package statecouchdb

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecodeOldAndNewFormat(t *testing.T) {
	testdata := []*statedb.VersionedValue{
		{
			Version: version.NewHeight(1, 2),
		},
		{
			Version: version.NewHeight(50, 50),
		},
		{
			Version:  version.NewHeight(50, 50),
			Metadata: []byte("sample-metadata"),
		},
	}

	for i, testdatum := range testdata {
		t.Run(fmt.Sprintf("testcase-newfmt-%d", i),
			func(t *testing.T) { testEncodeDecodeNewFormat(t, testdatum) },
		)
	}

	for i, testdatum := range testdata {
		t.Run(fmt.Sprintf("testcase-oldfmt-%d", i),
			func(t *testing.T) { testEncodeDecodeOldFormat(t, testdatum) },
		)
	}
}

func testEncodeDecodeNewFormat(t *testing.T, v *statedb.VersionedValue) {
	encodedVerField, err := encodeVersionAndMetadata(v.Version, v.Metadata)
	assert.NoError(t, err)

	ver, metadata, err := decodeVersionAndMetadata(encodedVerField)
	assert.NoError(t, err)
	assert.Equal(t, v.Version, ver)
	assert.Equal(t, v.Metadata, metadata)
}

func testEncodeDecodeOldFormat(t *testing.T, v *statedb.VersionedValue) {
	encodedVerField := encodeVersionOldFormat(v.Version)
//函数“decodeversionandmetadata”应该能够处理旧格式
	ver, metadata, err := decodeVersionAndMetadata(encodedVerField)
	assert.NoError(t, err)
	assert.Equal(t, v.Version, ver)
	assert.Nil(t, metadata)
}

