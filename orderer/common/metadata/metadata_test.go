
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:29</date>
//</624456093828911104>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package metadata_test

import (
	"fmt"
	"runtime"
	"testing"

	common "github.com/hyperledger/fabric/common/metadata"
	"github.com/hyperledger/fabric/orderer/common/metadata"
	"github.com/stretchr/testify/assert"
)

func TestGetVersionInfo(t *testing.T) {
//对于开发版本，此测试总是失败，因为
//common.version未设置，返回的字符串为“开发版本”
//在此设置此测试以避免此情况。
	if common.Version == "" {
		common.Version = "testVersion"
	}

	expected := fmt.Sprintf(
		"%s:\n Version: %s\n Commit SHA: %s\n Go version: %s\n OS/Arch: %s\n",
		metadata.ProgramName, common.Version,
		common.CommitSHA,
		runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	)
	assert.Equal(t, expected, metadata.GetVersionInfo())
}

