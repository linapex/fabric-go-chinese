
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456011238871040>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package decorator

import (
	"github.com/hyperledger/fabric/core/handlers/decoration"
	"github.com/hyperledger/fabric/protos/peer"
)

//new decorator创建新的decorator
func NewDecorator() decoration.Decorator {
	return &decorator{}
}

type decorator struct {
}

//修饰通过更改链码输入来修饰它
func (d *decorator) Decorate(proposal *peer.Proposal, input *peer.ChaincodeInput) *peer.ChaincodeInput {
	return input
}

