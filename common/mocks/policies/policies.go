
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:59</date>
//</624455966615670784>

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
	"github.com/hyperledger/fabric/common/policies"
	cb "github.com/hyperledger/fabric/protos/common"
)

//策略是策略的模拟实现。策略接口
type Policy struct {
//err是由evaluate返回的错误
	Err error
}

//
func (p *Policy) Evaluate(signatureSet []*cb.SignedData) error {
	return p.Err
}

//
type Manager struct {
//如果策略
//
	Policy *Policy

//返回的policymap用于查找
	PolicyMap map[string]policies.Policy

//
	SubManagersMap map[string]*Manager
}

//
func (m *Manager) Manager(path []string) (policies.Manager, bool) {
	if len(path) == 0 {
		return m, true
	}
	manager, ok := m.SubManagersMap[path[len(path)-1]]
	return manager, ok
}

//getpolicy返回manager.policy的值以及它是否为nil
func (m *Manager) GetPolicy(id string) (policies.Policy, bool) {
	if m.PolicyMap != nil {
		policy, ok := m.PolicyMap[id]
		if ok {
			return policy, true
		}
	}
	return m.Policy, m.Policy != nil
}

