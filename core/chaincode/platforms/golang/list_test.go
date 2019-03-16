
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:04</date>
//</624455988568657920>

/*
版权所有2017-greg haskins<gregory.haskins@gmail.com>

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


package golang

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_listDeps(t *testing.T) {
	_, err := listDeps(nil, "github.com/hyperledger/fabric/peer")
	if err != nil {
		t.Errorf("list failed: %s", err)
	}
}

func Test_runProgram(t *testing.T) {
	_, err := runProgram(
		getEnv(),
		10*time.Millisecond,
		"go",
		"build",
		"github.com/hyperledger/fabric/peer",
	)
	assert.Contains(t, err.Error(), "timed out")

	_, err = runProgram(
		getEnv(),
		1*time.Second,
		"go",
		"cmddoesnotexist",
	)
	assert.Contains(t, err.Error(), "unknown command")
}

