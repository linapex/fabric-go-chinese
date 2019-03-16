
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:10</date>
//</624456012866260992>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package validation

import (
	"github.com/hyperledger/fabric/core/handlers/validation/api"
	"github.com/hyperledger/fabric/protos/msp"
)

//IdentityDeserializer转换序列化的标识
//身份认同。
type IdentityDeserializer interface {
	validation.Dependency
//反序列化IDentity反序列化标识。
//如果标识与关联，则反序列化将失败
//与正在执行的MSP不同的MSP
//反序列化。
	DeserializeIdentity(serializedIdentity []byte) (Identity, error)
}

//定义与“证书”关联的操作的标识接口。
//也就是说，身份的公共部分可以被认为是证书，
//并且只提供签名验证功能。这是要用的
//在对等端验证已签名事务的证书时
//并验证与这些证书相对应的签名。
type Identity interface {
//验证使用控制此标识的规则来验证它。
	Validate() error

//satisfiesprincipal检查此实例是否匹配
//mspprincipal中提供的说明。支票可以
//涉及逐字节比较（如果主体是
//或可能需要MSP验证
	SatisfiesPrincipal(principal *msp.MSPPrincipal) error

//使用此标识作为引用验证某些消息上的签名
	Verify(msg []byte, sig []byte) error

//GetIdentityIdentifier返回该标识的标识符
	GetIdentityIdentifier() *IdentityIdentifier

//GetMSPIdentifier返回此实例的MSP ID
	GetMSPIdentifier() string
}

//IdentityIdentifier是特定
//通过提供程序标识符自然命名的标识。
type IdentityIdentifier struct {

//关联的成员身份服务提供程序的标识符
	Mspid string

//提供程序内标识的标识符
	Id string
}

