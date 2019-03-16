
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:17</date>
//</624456043614703616>

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


package transientstore

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//store env提供用于测试的store env
type StoreEnv struct {
	t                 testing.TB
	TestStoreProvider StoreProvider
	TestStore         Store
}

//newteststorenv构造用于测试的storeenv
func NewTestStoreEnv(t *testing.T) *StoreEnv {
	removeStorePath(t)
	assert := assert.New(t)
	testStoreProvider := NewStoreProvider()
	testStore, err := testStoreProvider.OpenStore("TestStore")
	assert.NoError(err)
	return &StoreEnv{t, testStoreProvider, testStore}
}

//清理测试后清理存储环境
func (env *StoreEnv) Cleanup() {
	env.TestStoreProvider.Close()
	removeStorePath(env.t)
}

func removeStorePath(t testing.TB) {
	dbPath := GetTransientStorePath()
	if err := os.RemoveAll(dbPath); err != nil {
		t.Fatalf("Err: %s", err)
		t.FailNow()
	}
}

