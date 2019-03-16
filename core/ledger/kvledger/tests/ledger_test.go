
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:11</date>
//</624456020244041728>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package tests

import (
	"os"
	"testing"

	"github.com/hyperledger/fabric/common/flogging"
)

func TestMain(m *testing.M) {
	flogging.ActivateSpec("lockbasedtxmgr,statevalidator,statebasedval,statecouchdb,valimpl,pvtstatepurgemgmt,confighistory,kvledger=debug")
	os.Exit(m.Run())
}

func TestLedgerAPIs(t *testing.T) {
	env := newEnv(defaultConfig, t)
	defer env.cleanup()

//创建两个分类帐
	h1 := newTestHelperCreateLgr("ledger1", t)
	h2 := newTestHelperCreateLgr("ledger2", t)

//用示例数据填充分类帐
	dataHelper := newSampleDataHelper(t)
	dataHelper.populateLedger(h1)
	dataHelper.populateLedger(h2)

//验证两个分类帐中的内容
	dataHelper.verifyLedgerContent(h1)
	dataHelper.verifyLedgerContent(h2)
}

