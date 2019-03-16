
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455963532857344>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package testdata

import "github.com/hyperledger/fabric/common/metrics"

//金多克：忽略

//因为上面的gendoc:ignore语句，所以文档生成应该忽略这一点。

var (
	Ignored = metrics.CounterOpts{
		Namespace: "ignored",
		Name:      "ignored",
	}
)

