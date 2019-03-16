
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456026875236352>

/*
版权所有IBM公司。保留所有权利。
SPDX许可证标识符：Apache-2.0
**/


package lockbasedtxmgr

import (
	"testing"

	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
	"github.com/stretchr/testify/assert"
)

func TestCollectionValidation(t *testing.T) {
	testEnv := testEnvsMap[levelDBtestEnvName]
	testEnv.init(t, "testLedger", nil)
	defer testEnv.cleanup()
	txMgr := testEnv.getTxMgr()
	populateCollConfigForTest(t, txMgr.(*LockBasedTxMgr),
		[]collConfigkey{
			{"ns1", "coll1"},
			{"ns1", "coll2"},
			{"ns2", "coll1"},
			{"ns2", "coll2"},
		},
		version.NewHeight(1, 1),
	)

	sim, err := txMgr.NewTxSimulator("tx-id1")
	assert.NoError(t, err)

	_, err = sim.GetPrivateData("ns3", "coll1", "key1")
	_, ok := err.(*ledger.CollConfigNotDefinedError)
	assert.True(t, ok)

	err = sim.SetPrivateData("ns3", "coll1", "key1", []byte("val1"))
	_, ok = err.(*ledger.CollConfigNotDefinedError)
	assert.True(t, ok)

	_, err = sim.GetPrivateData("ns1", "coll3", "key1")
	_, ok = err.(*ledger.InvalidCollNameError)
	assert.True(t, ok)

	err = sim.SetPrivateData("ns1", "coll3", "key1", []byte("val1"))
	_, ok = err.(*ledger.InvalidCollNameError)
	assert.True(t, ok)

	err = sim.SetPrivateData("ns1", "coll1", "key1", []byte("val1"))
	assert.NoError(t, err)
}

func TestPvtGetNoCollection(t *testing.T) {
	testEnv := testEnvs[0]
	testEnv.init(t, "test-pvtdata-get-no-collection", nil)
	defer testEnv.cleanup()
	txMgr := testEnv.getTxMgr().(*LockBasedTxMgr)
	queryHelper := newQueryHelper(txMgr, nil)
	valueHash, metadataBytes, err := queryHelper.getPrivateDataValueHash("cc", "coll", "key")
	assert.Nil(t, valueHash)
	assert.Nil(t, metadataBytes)
	assert.Error(t, err)
	assert.IsType(t, &ledger.CollConfigNotDefinedError{}, err)
}

func TestPvtPutNoCollection(t *testing.T) {
	testEnv := testEnvs[0]
	testEnv.init(t, "test-pvtdata-put-no-collection", nil)
	defer testEnv.cleanup()
	txMgr := testEnv.getTxMgr().(*LockBasedTxMgr)
	txsim, err := txMgr.NewTxSimulator("txid")
	assert.NoError(t, err)
	err = txsim.SetPrivateDataMetadata("cc", "coll", "key", map[string][]byte{})
	assert.Error(t, err)
	assert.IsType(t, &ledger.CollConfigNotDefinedError{}, err)
}

