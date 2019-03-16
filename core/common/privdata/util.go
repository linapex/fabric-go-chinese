
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:07</date>
//</624456002451804160>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package privdata

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/pkg/errors"
)

func getPolicy(collectionPolicyConfig *common.CollectionPolicyConfig, deserializer msp.IdentityDeserializer) (policies.Policy, error) {
	if collectionPolicyConfig == nil {
		return nil, errors.New("Collection policy config is nil")
	}
	accessPolicyEnvelope := collectionPolicyConfig.GetSignaturePolicy()
	if accessPolicyEnvelope == nil {
		return nil, errors.New("Collection config access policy is nil")
	}
//从信封创建访问策略
	npp := cauthdsl.NewPolicyProvider(deserializer)
	polBytes, err := proto.Marshal(accessPolicyEnvelope)
	if err != nil {
		return nil, err
	}
	accessPolicy, _, err := npp.NewPolicy(polBytes)
	if err != nil {
		return nil, err
	}
	return accessPolicy, nil
}

