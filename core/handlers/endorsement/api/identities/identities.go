
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456011683467264>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package endorsement

import (
	"github.com/hyperledger/fabric/core/handlers/endorsement/api"
	"github.com/hyperledger/fabric/protos/peer"
)

//SigningIdentity对消息进行签名并将其公共标识序列化为字节
type SigningIdentity interface {
//serialize返回用于验证的此标识的字节表示形式
//由此签名身份签名的邮件
	Serialize() ([]byte, error)

//sign对给定的有效负载进行签名并返回签名
	Sign([]byte) ([]byte, error)
}

//SigningIdentityFetcher根据建议获取签名标识
type SigningIdentityFetcher interface {
	endorsement.Dependency
//SigningIdentityForRequest返回给定建议的签名标识
	SigningIdentityForRequest(*peer.SignedProposal) (SigningIdentity, error)
}

