
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:30</date>
//</624456097142411264>

//+构建窗口

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package server

import (
	"os"
)

func addPlatformSignals(sigs map[os.Signal]func()) map[os.Signal]func() {
	return sigs
}

