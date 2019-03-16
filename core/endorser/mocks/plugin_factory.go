
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456009003307008>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


//Mokery v1.0.0生成的代码
package mocks

import endorsement "github.com/hyperledger/fabric/core/handlers/endorsement/api"
import mock "github.com/stretchr/testify/mock"

//PlugInfectory是PlugInfectory类型的自动生成的模拟类型
type PluginFactory struct {
	mock.Mock
}

//new为给定字段提供模拟函数：
func (_m *PluginFactory) New() endorsement.Plugin {
	ret := _m.Called()

	var r0 endorsement.Plugin
	if rf, ok := ret.Get(0).(func() endorsement.Plugin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(endorsement.Plugin)
		}
	}

	return r0
}

