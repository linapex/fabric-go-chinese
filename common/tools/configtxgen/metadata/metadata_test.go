
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:59</date>
//</624455969232916480>

/*
版权所有2017 Hitachi America

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/


package metadata_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/hyperledger/fabric/common/tools/configtxgen/metadata"
	"github.com/stretchr/testify/assert"
)

func TestGetVersionInfo(t *testing.T) {
	testSHAs := []string{"", "abcdefg"}

	for _, sha := range testSHAs {
		metadata.CommitSHA = sha
		if sha == "" {
			sha = "development build"
		}

		expected := fmt.Sprintf("%s:\n Version: %s\n Commit SHA: %s\n Go version: %s\n OS/Arch: %s",
			metadata.ProgramName, metadata.Version, sha, runtime.Version(),
			fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
		assert.Equal(t, expected, metadata.GetVersionInfo())
	}
}

