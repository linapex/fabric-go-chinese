
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456026690686976>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package storageutil

import (
	"testing"

	"github.com/hyperledger/fabric/protos/ledger/rwset/kvrwset"
	"github.com/stretchr/testify/assert"
)

func TestSerializeDeSerialize(t *testing.T) {
	sampleMetadata := []*kvrwset.KVMetadataEntry{
		{Name: "metadata_1", Value: []byte("metadata_value_1")},
		{Name: "metadata_2", Value: []byte("metadata_value_2")},
		{Name: "metadata_3", Value: []byte("metadata_value_3")},
	}

	serializedMetadata, err := SerializeMetadata(sampleMetadata)
	assert.NoError(t, err)
	metadataMap, err := DeserializeMetadata(serializedMetadata)
	assert.NoError(t, err)
	assert.Len(t, metadataMap, 3)
	assert.Equal(t, []byte("metadata_value_1"), metadataMap["metadata_1"])
	assert.Equal(t, []byte("metadata_value_2"), metadataMap["metadata_2"])
	assert.Equal(t, []byte("metadata_value_3"), metadataMap["metadata_3"])
}

