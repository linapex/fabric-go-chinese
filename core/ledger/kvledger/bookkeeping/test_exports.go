
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:11</date>
//</624456018100752384>

/*
版权所有IBM Corp.2017保留所有权利。

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


package bookkeeping

import (
	"os"
	"testing"
)

//testenv为测试提供簿记提供者env
type TestEnv struct {
	t            testing.TB
	TestProvider Provider
}

//newtestenv构造用于测试的testenv
func NewTestEnv(t testing.TB) *TestEnv {
	removePath(t)
	provider := NewProvider()
	return &TestEnv{t, provider}
}

//清理测试后清理存储环境
func (env *TestEnv) Cleanup() {
	env.TestProvider.Close()
	removePath(env.t)
}

func removePath(t testing.TB) {
	dbPath := getInternalBookkeeperPath()
	if err := os.RemoveAll(dbPath); err != nil {
		t.Fatalf("Err: %s", err)
		t.FailNow()
	}
}

