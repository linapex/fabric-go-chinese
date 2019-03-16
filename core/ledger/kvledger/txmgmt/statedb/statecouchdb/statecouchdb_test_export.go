
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456025377869824>

/*
版权所有IBM公司。保留所有权利。
SPDX许可证标识符：Apache-2.0
**/


package statecouchdb

import (
	"testing"

	"github.com/hyperledger/fabric/common/metrics/disabled"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/stretchr/testify/assert"
)

//testvdbenv提供了一个基于coach-db的版本数据库用于测试
type TestVDBEnv struct {
	t          testing.TB
	DBProvider statedb.VersionedDBProvider
}

//newtestvdbenv实例化和新的coach db支持的testvdb
func NewTestVDBEnv(t testing.TB) *TestVDBEnv {
	t.Logf("Creating new TestVDBEnv")

	dbProvider, _ := NewVersionedDBProvider(&disabled.Provider{})
	testVDBEnv := &TestVDBEnv{t, dbProvider}
//没有新测试环境的清理。需要为测试中使用的每个数据库清除每个测试。
	return testVDBEnv
}

//Cleanup drops the test couch databases and closes the db provider
func (env *TestVDBEnv) Cleanup() {
	env.t.Logf("Cleaningup TestVDBEnv")
	CleanupDB(env.t, env.DBProvider)

	env.DBProvider.Close()
}

func CleanupDB(t testing.TB, dbProvider statedb.VersionedDBProvider) {
	couchdbProvider, _ := dbProvider.(*VersionedDBProvider)
	for _, v := range couchdbProvider.databases {
		if _, err := v.metadataDB.DropDatabase(); err != nil {
			assert.Failf(t, "DropDatabase %s fails. err: %v", v.metadataDB.DBName, err)
		}

		for _, db := range v.namespaceDBs {
			if _, err := db.DropDatabase(); err != nil {
				assert.Failf(t, "DropDatabase %s fails. err: %v", db.DBName, err)
			}
		}
	}
}

