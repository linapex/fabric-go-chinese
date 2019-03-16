
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:04</date>
//</624455987931123712>

/*
版权所有，State Street Corp.保留所有权利。
γ
SPDX许可证标识符：Apache-2.0
**/


package ccmetadata

import (
	"github.com/hyperledger/fabric/common/flogging"
)

//此包使用的记录器
var logger = flogging.MustGetLogger("chaincode.platform.metadata")

//MetadataProvider由每个平台以特定于平台的方式实现。
//它可以处理以不同格式存储在chaincodedeploymentspec中的元数据。
//通用格式是targz。当前用户希望显示元数据
//作为tar文件条目（直接从以targz格式存储的链码中提取）。
//将来，我们希望通过扩展接口来提供更好的抽象
type MetadataProvider interface {
	GetMetadataAsTarEntries() ([]byte, error)
}

