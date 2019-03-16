
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:10</date>
//</624456013625430016>

//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import common "github.com/hyperledger/fabric/protos/common"
import mock "github.com/stretchr/testify/mock"

//policyEvaluator是policyEvaluator类型的自动生成的模拟类型
type PolicyEvaluator struct {
	mock.Mock
}

//Evaluate提供了一个具有给定字段的模拟函数：policyBytes、signatureSet
func (_m *PolicyEvaluator) Evaluate(policyBytes []byte, signatureSet []*common.SignedData) error {
	ret := _m.Called(policyBytes, signatureSet)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []*common.SignedData) error); ok {
		r0 = rf(policyBytes, signatureSet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
