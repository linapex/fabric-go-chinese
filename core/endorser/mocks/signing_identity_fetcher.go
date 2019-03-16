
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456009183662080>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


//Mokery v1.0.0生成的代码
package mocks

import endorsement "github.com/hyperledger/fabric/core/handlers/endorsement/api/identities"
import mock "github.com/stretchr/testify/mock"
import peer "github.com/hyperledger/fabric/protos/peer"

//SigningIdentityFetcher是为SigningIdentityFetcher类型自动生成的模拟类型
type SigningIdentityFetcher struct {
	mock.Mock
}

//SigningIdentityForRequest提供了一个具有给定字段的模拟函数：a0
func (_m *SigningIdentityFetcher) SigningIdentityForRequest(_a0 *peer.SignedProposal) (endorsement.SigningIdentity, error) {
	ret := _m.Called(_a0)

	var r0 endorsement.SigningIdentity
	if rf, ok := ret.Get(0).(func(*peer.SignedProposal) endorsement.SigningIdentity); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(endorsement.SigningIdentity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*peer.SignedProposal) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

