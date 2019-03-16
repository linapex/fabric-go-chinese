
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455976501645312>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package mocks

import (
	"testing"

	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/stretchr/testify/mock"
)

type MockACLProvider struct {
//创建一个模拟对象，可以在
//在aclmgmt中注册mockaclprovider
	mock *mock.Mock
}

//把模型弄清楚，我们可以重新开始
func (m *MockACLProvider) Reset() {
	m.mock = &mock.Mock{}
}

func (m *MockACLProvider) CheckACL(resName string, channelID string, idinfo interface{}) error {
	args := m.mock.Called(resName, channelID, idinfo)
	return args.Error(0)
}

func (m *MockACLProvider) GenerateSimulationResults(txEnvelop *common.Envelope, simulator ledger.TxSimulator, initializingLedger bool) error {
	return nil
}

//为方便起见重写模拟方法
func (m *MockACLProvider) On(methodName string, arguments ...interface{}) *mock.Call {
	return m.mock.On(methodName, arguments...)
}

//断言期望覆盖模拟方法以方便
func (m *MockACLProvider) AssertExpectations(t *testing.T) {
	m.mock.AssertExpectations(t)
}

