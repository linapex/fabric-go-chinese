
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456011079487488>

/*
版权所有IBM Corp，SecureKey Technologies Inc.保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package decoration

import (
	"github.com/hyperledger/fabric/protos/peer"
)

//decorator修饰链码输入
type Decorator interface {
//修饰通过更改链码输入来修饰它
	Decorate(proposal *peer.Proposal, input *peer.ChaincodeInput) *peer.ChaincodeInput
}

//按提供的顺序应用装饰
func Apply(proposal *peer.Proposal, input *peer.ChaincodeInput,
	decorators ...Decorator) *peer.ChaincodeInput {
	for _, decorator := range decorators {
		input = decorator.Decorate(proposal, input)
	}

	return input
}

