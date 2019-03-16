
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:04</date>
//</624455989059391488>

/*
 *版权所有Greg Haskins保留所有权利
 *
 *SPDX许可证标识符：Apache-2.0
 *
 *有关详细信息，请参阅github.com/hyperledger/fabric/test/chaincodes/autovendor/chaincode/main.go。
 **/

package directdep

import (
	"chaincodes/AutoVendor/indirectdep"
)

func PointlessFunction() {
//授权我们间接依赖
	indirectdep.PointlessFunction()
}

