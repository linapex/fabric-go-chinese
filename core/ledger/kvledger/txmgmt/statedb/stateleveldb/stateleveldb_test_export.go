
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456026329976832>

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


package stateleveldb

import (
	"os"
	"testing"

	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/hyperledger/fabric/core/ledger/ledgerconfig"
)

//testvdbenv为测试提供了一个支持级别db的版本数据库
type TestVDBEnv struct {
	t          testing.TB
	DBProvider statedb.VersionedDBProvider
}

//newtestvdbenv实例化和新的级别db支持的testvdb
func NewTestVDBEnv(t testing.TB) *TestVDBEnv {
	t.Logf("Creating new TestVDBEnv")
	removeDBPath(t, "NewTestVDBEnv")
	dbProvider := NewVersionedDBProvider()
	return &TestVDBEnv{t, dbProvider}
}

//Cleanup closes the db and removes the db folder
func (env *TestVDBEnv) Cleanup() {
	env.t.Logf("Cleaningup TestVDBEnv")
	env.DBProvider.Close()
	removeDBPath(env.t, "Cleanup")
}

func removeDBPath(t testing.TB, caller string) {
	dbPath := ledgerconfig.GetStateLevelDBPath()
	if err := os.RemoveAll(dbPath); err != nil {
		t.Fatalf("Err: %s", err)
		t.FailNow()
	}
	logger.Debugf("Removed folder [%s] for test environment for %s", dbPath, caller)
}

