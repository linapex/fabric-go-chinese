
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456011570221056>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package endorsement

import (
	"github.com/hyperledger/fabric/protos/peer"
)

//参数定义了用于背书的参数
type Argument interface {
	Dependency
//arg返回参数的字节数
	Arg() []byte
}

//依赖项标记传递给init（）方法的依赖项
type Dependency interface {
}

//插件认可建议响应
type Plugin interface {
//认可对给定的有效负载（proposalResponsePayLoad字节）进行签名，并可选地对其进行变异。
//返回：
//背书：有效载荷上的签名，以及用于验证签名的标识。
//作为输入给出的有效负载（可以在此函数中修改）
//或失败时出错
	Endorse(payload []byte, sp *peer.SignedProposal) (*peer.Endorsement, []byte, error)

//init将依赖项插入插件的实例中
	Init(dependencies ...Dependency) error
}

//PluginFactory创建插件的新实例
type PluginFactory interface {
	New() Plugin
}

