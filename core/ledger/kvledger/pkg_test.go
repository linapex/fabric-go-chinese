
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:11</date>
//</624456019799445504>

/*
版权所有IBM Corp.2016保留所有权利。

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


package kvledger

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/spf13/viper"
)

type testEnv struct {
	t    testing.TB
	path string
}

func newTestEnv(t testing.TB) *testEnv {
	path := filepath.Join(
		os.TempDir(),
		"fabric",
		"ledgertests",
		"kvledger",
		strconv.Itoa(rand.Int()))
	return createTestEnv(t, path)
}

func createTestEnv(t testing.TB, path string) *testEnv {
	env := &testEnv{
		t:    t,
		path: path}
	env.cleanup()
	viper.Set("peer.fileSystemPath", env.path)
	return env
}

func (env *testEnv) cleanup() {
	os.RemoveAll(env.path)
}

