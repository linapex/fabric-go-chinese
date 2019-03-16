
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:14</date>
//</624456032386551808>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package testutil

import (
	"github.com/hyperledger/fabric/core/ledger/pvtdatapolicy"
	"github.com/hyperledger/fabric/core/ledger/pvtdatapolicy/mock"
	"github.com/hyperledger/fabric/protos/common"
)

//sample btlpolicy帮助测试创建示例btlpolicy
//示例输入项是[2]字符串ns，coll：btl
func SampleBTLPolicy(m map[[2]string]uint64) pvtdatapolicy.BTLPolicy {
	ccInfoRetriever := &mock.CollectionInfoProvider{}
	ccInfoRetriever.CollectionInfoStub = func(ccName, collName string) (*common.StaticCollectionConfig, error) {
		btl := m[[2]string{ccName, collName}]
		return &common.StaticCollectionConfig{BlockToLive: btl}, nil
	}
	return pvtdatapolicy.ConstructBTLPolicy(ccInfoRetriever)
}

