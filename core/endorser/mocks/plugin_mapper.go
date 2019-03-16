
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456009062027264>

//Mokery v1.0.0生成的代码
package mocks

import endorsement "github.com/hyperledger/fabric/core/handlers/endorsement/api"
import endorser "github.com/hyperledger/fabric/core/endorser"
import mock "github.com/stretchr/testify/mock"

//pluginmapper是为pluginmapper类型自动生成的模拟类型
type PluginMapper struct {
	mock.Mock
}

//PlugInfactoryByName提供了一个具有给定字段的模拟函数：name
func (_m *PluginMapper) PluginFactoryByName(name endorser.PluginName) endorsement.PluginFactory {
	ret := _m.Called(name)

	var r0 endorsement.PluginFactory
	if rf, ok := ret.Get(0).(func(endorser.PluginName) endorsement.PluginFactory); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(endorsement.PluginFactory)
		}
	}

	return r0
}

