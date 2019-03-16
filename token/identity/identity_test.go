
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456127232348160>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package identity_test

import "github.com/hyperledger/fabric/msp"

//go：生成伪造者-o mock/issuing_validator.go-fake name issuing validator。颁发验证程序
//go：生成伪造者-o mock/public_info.go-forke name public info。公共信息
//go:生成伪造者-o mock/deserializer_manager.go-fake name deserializer manager。反序列化管理器
//go：生成伪造者-o mock/deserializer.go-fake name deserializer。解串器
//go:生成伪造者-o mock/identity.go-伪造名称标识。/../../msp/identity

type Identity interface {
	msp.Identity
}

