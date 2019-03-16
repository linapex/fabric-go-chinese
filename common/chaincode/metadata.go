
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455944624934912>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package chaincode

import (
	"sync"

	"github.com/hyperledger/fabric/protos/gossip"
)

//InstalledChaincode定义有关已安装链代码的元数据
type InstalledChaincode struct {
	Name    string
	Version string
	Id      []byte
}

//元数据定义链码的通道范围元数据
type Metadata struct {
	Name              string
	Version           string
	Policy            []byte
	Id                []byte
	CollectionsConfig []byte
}

//metadataset定义元数据的聚合
type MetadataSet []Metadata

//aschaincodes将此元数据集转换为八卦片段。chaincodes
func (ccs MetadataSet) AsChaincodes() []*gossip.Chaincode {
	var res []*gossip.Chaincode
	for _, cc := range ccs {
		res = append(res, &gossip.Chaincode{
			Name:    cc.Name,
			Version: cc.Version,
		})
	}
	return res
}

//metadata mapping定义了从chaincode名称到元数据的映射
type MetadataMapping struct {
	sync.RWMutex
	mdByName map[string]Metadata
}

//NewMetadataMapping创建新的元数据映射
func NewMetadataMapping() *MetadataMapping {
	return &MetadataMapping{
		mdByName: make(map[string]Metadata),
	}
}

//查找返回与给定链码关联的元数据
func (m *MetadataMapping) Lookup(cc string) (Metadata, bool) {
	m.RLock()
	defer m.RUnlock()
	md, exists := m.mdByName[cc]
	return md, exists
}

//更新更新更新映射中的链码元数据
func (m *MetadataMapping) Update(ccMd Metadata) {
	m.Lock()
	defer m.Unlock()
	m.mdByName[ccMd.Name] = ccMd
}

//聚合将所有元数据聚合到元数据集
func (m *MetadataMapping) Aggregate() MetadataSet {
	m.RLock()
	defer m.RUnlock()
	var set MetadataSet
	for _, md := range m.mdByName {
		set = append(set, md)
	}
	return set
}

