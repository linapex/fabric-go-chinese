
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:10</date>
//</624456014531399680>

//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import common "github.com/hyperledger/fabric/protos/common"
import errors "github.com/hyperledger/fabric/common/errors"
import mock "github.com/stretchr/testify/mock"
import peer "github.com/hyperledger/fabric/protos/peer"

//StateBasedvalidator是StateBasedvalidator类型的自动生成的模拟类型
type StateBasedValidator struct {
	mock.Mock
}

//postvalidate提供了一个具有给定字段的模拟函数：cc、blocknum、txnum、err
func (_m *StateBasedValidator) PostValidate(cc string, blockNum uint64, txNum uint64, err error) {
	_m.Called(cc, blockNum, txNum, err)
}

//prevalidate提供具有给定字段的模拟函数：txnum，block
func (_m *StateBasedValidator) PreValidate(txNum uint64, block *common.Block) {
	_m.Called(txNum, block)
}

//validate提供具有给定字段的模拟函数：cc、blocknum、txnum、rwset、prp、ep、背书
func (_m *StateBasedValidator) Validate(cc string, blockNum uint64, txNum uint64, rwset []byte, prp []byte, ep []byte, endorsements []*peer.Endorsement) errors.TxValidationError {
	ret := _m.Called(cc, blockNum, txNum, rwset, prp, ep, endorsements)

	var r0 errors.TxValidationError
	if rf, ok := ret.Get(0).(func(string, uint64, uint64, []byte, []byte, []byte, []*peer.Endorsement) errors.TxValidationError); ok {
		r0 = rf(cc, blockNum, txNum, rwset, prp, ep, endorsements)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.TxValidationError)
		}
	}

	return r0
}

