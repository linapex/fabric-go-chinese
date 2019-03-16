
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456008873283584>

//Mokery v1.0.0生成的代码
package mocks

import endorser "github.com/hyperledger/fabric/core/endorser"
import mock "github.com/stretchr/testify/mock"

//ChannelStateRetriever是ChannelStateRetriever类型的自动生成的模拟类型
type ChannelStateRetriever struct {
	mock.Mock
}

//newQueryCreator提供了一个具有给定字段的模拟函数：channel
func (_m *ChannelStateRetriever) NewQueryCreator(channel string) (endorser.QueryCreator, error) {
	ret := _m.Called(channel)

	var r0 endorser.QueryCreator
	if rf, ok := ret.Get(0).(func(string) endorser.QueryCreator); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(endorser.QueryCreator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

