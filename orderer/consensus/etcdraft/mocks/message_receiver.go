
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:30</date>
//</624456098723663872>

//Code generated by mockery v1.0.0. 不要编辑。

package mocks

import mock "github.com/stretchr/testify/mock"
import orderer "github.com/hyperledger/fabric/protos/orderer"

//messagereceiver是为messagereceiver类型自动生成的模拟类型
type MessageReceiver struct {
	mock.Mock
}

//步骤为给定字段提供模拟函数：req，sender
func (_m *MessageReceiver) Step(req *orderer.StepRequest, sender uint64) error {
	ret := _m.Called(req, sender)

	var r0 error
	if rf, ok := ret.Get(0).(func(*orderer.StepRequest, uint64) error); ok {
		r0 = rf(req, sender)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//Submit提供了一个具有给定字段的模拟函数：req，sender
func (_m *MessageReceiver) Submit(req *orderer.SubmitRequest, sender uint64) error {
	ret := _m.Called(req, sender)

	var r0 error
	if rf, ok := ret.Get(0).(func(*orderer.SubmitRequest, uint64) error); ok {
		r0 = rf(req, sender)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
