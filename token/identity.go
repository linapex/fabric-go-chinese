
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456127744053248>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package token

//go：生成伪造者-o客户端/mock/identity.go-伪造名称标识。身份

//身份是指TX的创建者；
type Identity interface {
	Serialize() ([]byte, error)
}

//go：生成伪造者-o客户端/mock/signing_identity.go-伪造名称signingIdentity。签名身份

//SigningIdentity定义签名
//字节数组；需要对传输到的命令进行签名
//提供程序对等服务。
type SigningIdentity interface {
Identity //扩展标识

	Sign(msg []byte) ([]byte, error)

	GetPublicVersion() Identity
}

