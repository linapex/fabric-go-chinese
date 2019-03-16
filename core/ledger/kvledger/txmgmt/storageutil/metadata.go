
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456026606800896>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package storageutil

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/ledger/rwset/kvrwset"
)

//serializemetadata序列化用于在statedb中进行stroking的元数据项
func SerializeMetadata(metadataEntries []*kvrwset.KVMetadataEntry) ([]byte, error) {
	metadata := &kvrwset.KVMetadataWrite{Entries: metadataEntries}
	return proto.Marshal(metadata)
}

//反序列化元数据从StateDB反序列化元数据字节
func DeserializeMetadata(metadataBytes []byte) (map[string][]byte, error) {
	if metadataBytes == nil {
		return nil, nil
	}
	metadata := &kvrwset.KVMetadataWrite{}
	if err := proto.Unmarshal(metadataBytes, metadata); err != nil {
		return nil, err
	}
	m := make(map[string][]byte, len(metadata.Entries))
	for _, metadataEntry := range metadata.Entries {
		m[metadataEntry.Name] = metadataEntry.Value
	}
	return m, nil
}

