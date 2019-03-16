
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:07</date>
//</624456001742966784>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package privdata

import (
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protos/common"
)

//MembershipProvider可用于检查对等方是否符合集合的条件
type MembershipProvider struct {
	selfSignedData              common.SignedData
	IdentityDeserializerFactory func(chainID string) msp.IdentityDeserializer
}

//NewMembershipInfoProvider返回MembershipProvider
func NewMembershipInfoProvider(selfSignedData common.SignedData, identityDeserializerFunc func(chainID string) msp.IdentityDeserializer) *MembershipProvider {
	return &MembershipProvider{selfSignedData: selfSignedData, IdentityDeserializerFactory: identityDeserializerFunc}
}

//ammemberof检查当前对等方是否为给定集合配置的成员
func (m *MembershipProvider) AmMemberOf(channelName string, collectionPolicyConfig *common.CollectionPolicyConfig) (bool, error) {
	deserializer := m.IdentityDeserializerFactory(channelName)
	accessPolicy, err := getPolicy(collectionPolicyConfig, deserializer)
	if err != nil {
		return false, err
	}
	if err := accessPolicy.Evaluate([]*common.SignedData{&m.selfSignedData}); err != nil {
		return false, nil
	}
	return true, nil
}

