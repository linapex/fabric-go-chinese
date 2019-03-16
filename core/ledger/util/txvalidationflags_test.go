
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456036031401984>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package util

import (
	"testing"

	"github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/assert"
)

func TestTransactionValidationFlags(t *testing.T) {
	txFlags := NewTxValidationFlagsSetValue(10, peer.TxValidationCode_VALID)
	assert.Equal(t, 10, len(txFlags))

	txFlags.SetFlag(0, peer.TxValidationCode_VALID)
	assert.Equal(t, peer.TxValidationCode_VALID, txFlags.Flag(0))
	assert.Equal(t, true, txFlags.IsValid(0))

	txFlags.SetFlag(1, peer.TxValidationCode_MVCC_READ_CONFLICT)
	assert.Equal(t, true, txFlags.IsInvalid(1))
}

