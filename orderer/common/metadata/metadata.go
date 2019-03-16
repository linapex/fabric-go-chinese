
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:29</date>
//</624456093765996544>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package metadata

import (
	"fmt"
	"runtime"

	common "github.com/hyperledger/fabric/common/metadata"
)

//包范围变量

//包装版本
var Version string

//包范围常量

//程序名
const ProgramName = "orderer"

func GetVersionInfo() string {
	Version = common.Version
	if Version == "" {
		Version = "development build"
	}

	return fmt.Sprintf(
		"%s:\n Version: %s\n Commit SHA: %s\n Go version: %s\n OS/Arch: %s\n",
		ProgramName,
		Version,
		common.CommitSHA,
		runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	)
}

