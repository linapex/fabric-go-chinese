
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:59</date>
//</624455967177707520>

/*
版权所有IBM Corp.2016保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

                 http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/


package policies

import (
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/utils"
)

//ImplicitMetapolicyWithSubpolicy创建一个ImplicitMetapolicy
func ImplicitMetaPolicyWithSubPolicy(subPolicyName string, rule cb.ImplicitMetaPolicy_Rule) *cb.ConfigPolicy {
	return &cb.ConfigPolicy{
		Policy: &cb.Policy{
			Type: int32(cb.Policy_IMPLICIT_META),
			Value: utils.MarshalOrPanic(&cb.ImplicitMetaPolicy{
				Rule:      rule,
				SubPolicy: subPolicyName,
			}),
		},
	}
}

//
func TemplateImplicitMetaPolicyWithSubPolicy(path []string, policyName string, subPolicyName string, rule cb.ImplicitMetaPolicy_Rule) *cb.ConfigGroup {
	root := cb.NewConfigGroup()
	group := root
	for _, element := range path {
		group.Groups[element] = cb.NewConfigGroup()
		group = group.Groups[element]
	}

	group.Policies[policyName] = ImplicitMetaPolicyWithSubPolicy(subPolicyName, rule)
	return root
}

//
//
func TemplateImplicitMetaPolicy(path []string, policyName string, rule cb.ImplicitMetaPolicy_Rule) *cb.ConfigGroup {
	return TemplateImplicitMetaPolicyWithSubPolicy(path, policyName, policyName, rule)
}

//
func TemplateImplicitMetaAnyPolicy(path []string, policyName string) *cb.ConfigGroup {
	return TemplateImplicitMetaPolicy(path, policyName, cb.ImplicitMetaPolicy_ANY)
}

//
func TemplateImplicitMetaAllPolicy(path []string, policyName string) *cb.ConfigGroup {
	return TemplateImplicitMetaPolicy(path, policyName, cb.ImplicitMetaPolicy_ALL)
}

//tempateImplicitMetaanyPolicy返回templateImplicitMetapolicy，规则为cb.implicitMetapolicy_多数
func TemplateImplicitMetaMajorityPolicy(path []string, policyName string) *cb.ConfigGroup {
	return TemplateImplicitMetaPolicy(path, policyName, cb.ImplicitMetaPolicy_MAJORITY)
}

