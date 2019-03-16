
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456126552870912>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package client

//go：生成伪造者-o mock/signer_identity.go-fake name signer identity。签名身份

type Signer interface {
//sign对给定的有效负载进行签名并返回签名
	Sign([]byte) ([]byte, error)
}

//SignerIdentity对消息进行签名并将其公共标识序列化为字节
type SignerIdentity interface {
	Signer

//serialize返回用于验证的此标识的字节表示形式
//此SignerIdentity签名的邮件
	Serialize() ([]byte, error)
}

