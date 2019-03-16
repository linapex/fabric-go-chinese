
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:00</date>
//</624455973951508480>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package metadata_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/hyperledger/fabric/common/tools/idemixgen/metadata"
	"github.com/stretchr/testify/assert"
)

func TestGetVersionInfo(t *testing.T) {
	testSHA := "abcdefg"
	metadata.CommitSHA = testSHA

	expected := fmt.Sprintf("%s:\n Version: %s\n Commit SHA: %s\n Go version: %s\n OS/Arch: %s",
		metadata.ProgramName, metadata.Version, testSHA, runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
	assert.Equal(t, expected, metadata.GetVersionInfo())
}

