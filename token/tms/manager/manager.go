
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:38</date>
//</624456130273218560>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package manager

import (
	"github.com/hyperledger/fabric/msp/mgmt"
	"github.com/hyperledger/fabric/token/identity"
	"github.com/hyperledger/fabric/token/tms/plain"
	"github.com/hyperledger/fabric/token/transaction"
	"github.com/pkg/errors"
)

//go：生成伪造者-o mock/identity _deserializer _manager.go-fake name deserializer manager。反序列化管理器

//FabricidentityDeserializerManager实现一个DeserializerManager
//通过将呼叫路由到MSP/MGMT包
type FabricIdentityDeserializerManager struct {
}

func (*FabricIdentityDeserializerManager) Deserializer(channel string) (identity.Deserializer, error) {
	id, ok := mgmt.GetDeserializers()[channel]
	if !ok {
		return nil, errors.New("channel not found")
	}
	return id, nil
}

//管理器用于访问TMS组件。
type Manager struct {
	IdentityDeserializerManager identity.DeserializerManager
}

//gettxprocessor返回用于处理令牌事务的tmstxprocessor。
func (m *Manager) GetTxProcessor(channel string) (transaction.TMSTxProcessor, error) {
	identityDeserializerManager, err := m.IdentityDeserializerManager.Deserializer(channel)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting identity deserialiser manager for channel '%s'", channel)
	}

	return &plain.Verifier{IssuingValidator: &AllIssuingValidator{Deserializer: identityDeserializerManager}}, nil
}

