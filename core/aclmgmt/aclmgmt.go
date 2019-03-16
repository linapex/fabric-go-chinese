
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455976229015552>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package aclmgmt

import (
	"github.com/hyperledger/fabric/common/flogging"
)

var aclLogger = flogging.MustGetLogger("aclmgmt")

type ACLProvider interface {
//checkacl使用
//IDFIN。IDinfo是一个对象，如SignedProposal，其中
//可以提取ID以根据策略进行测试
	CheckACL(resName string, channelID string, idinfo interface{}) error
}

