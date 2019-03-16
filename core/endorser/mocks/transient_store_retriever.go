
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456009519206400>

//Mokery v1.0.0生成的代码
package mocks

import mock "github.com/stretchr/testify/mock"
import transientstore "github.com/hyperledger/fabric/core/transientstore"

//TransientStoreRetriever是TransientStoreRetriever类型的自动生成的模拟类型
type TransientStoreRetriever struct {
	mock.Mock
}

//storeforchannel为给定字段提供模拟函数：channel
func (_m *TransientStoreRetriever) StoreForChannel(channel string) transientstore.Store {
	ret := _m.Called(channel)

	var r0 transientstore.Store
	if rf, ok := ret.Get(0).(func(string) transientstore.Store); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transientstore.Store)
		}
	}

	return r0
}

