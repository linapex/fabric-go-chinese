
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:11</date>
//</624456018029449216>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package bookkeeping

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	viper.Set("peer.fileSystemPath", "/tmp/fabric/ledgertests/kvledger/bookkeeping")
	os.Exit(m.Run())
}

func TestProvider(t *testing.T) {
	testEnv := NewTestEnv(t)
	defer testEnv.Cleanup()
	p := testEnv.TestProvider
	db := p.GetDBHandle("TestLedger", PvtdataExpiry)
	assert.NoError(t, db.Put([]byte("key"), []byte("value"), true))
	val, err := db.Get([]byte("key"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("value"), val)
}

