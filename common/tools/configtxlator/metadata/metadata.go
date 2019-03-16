
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:59</date>
//</624455969690095616>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package metadata

import (
	"fmt"
	"runtime"
)

//包范围变量

//包装版本
const Version = "1.4.1"

var CommitSHA string

//包范围常量

//程序名
const ProgramName = "configtxlator"

func GetVersionInfo() string {
	if CommitSHA == "" {
		CommitSHA = "development build"
	}

	return fmt.Sprintf("%s:\n Version: %s\n Commit SHA: %s\n Go version: %s\n OS/Arch: %s",
		ProgramName, Version, CommitSHA, runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
}

