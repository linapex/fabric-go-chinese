
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:38</date>
//</624456131133050880>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package plain

import (
	"sync"

	"github.com/hyperledger/fabric/token/identity"
	"github.com/hyperledger/fabric/token/transaction"
	"github.com/pkg/errors"
)

//管理器用于访问TMS组件。
type Manager struct {
	mutex            sync.RWMutex
	policyValidators map[string]identity.IssuingValidator
}

//gettxprocessor返回用于处理令牌事务的tmstxprocessor。
func (m *Manager) GetTxProcessor(channel string) (transaction.TMSTxProcessor, error) {
	m.mutex.RLock()
	policyValidator := m.policyValidators[channel]
	m.mutex.RUnlock()
	if policyValidator == nil {
		return nil, errors.Errorf("no policy validator found for channel '%s'", channel)
	}
	return &Verifier{IssuingValidator: policyValidator}, nil
}

//setpolicyvalidator为指定通道设置策略验证程序
func (m *Manager) SetPolicyValidator(channel string, validator identity.IssuingValidator) {
	m.mutex.Lock()
	m.policyValidators[channel] = validator
	m.mutex.Unlock()
}

