
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456037428105216>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package validator

import (
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/mock"
)

//mockvalidator实现了一个对测试有用的模拟验证
type MockValidator struct {
	mock.Mock
}

//验证不执行任何操作，返回无错误
func (m *MockValidator) Validate(block *common.Block) error {
	if len(m.ExpectedCalls) == 0 {
		return nil
	}
	return m.Called().Error(0)
}

//mockvsccvalidator是vscc验证接口的模拟实现。
type MockVsccValidator struct {
}

//vsccvalidatetx不做任何操作
func (v *MockVsccValidator) VSCCValidateTx(seq int, payload *common.Payload, envBytes []byte, block *common.Block) (error, peer.TxValidationCode) {
	return nil, peer.TxValidationCode_VALID
}

