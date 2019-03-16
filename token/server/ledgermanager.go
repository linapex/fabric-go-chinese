
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456128494833664>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package server

import (
	"github.com/hyperledger/fabric/core/peer"
	"github.com/hyperledger/fabric/token/ledger"
	"github.com/pkg/errors"
)

//PeerledgerManager实现了ledgerManager接口
//通过使用对等基础设施
type PeerLedgerManager struct {
}

func (*PeerLedgerManager) GetLedgerReader(channel string) (ledger.LedgerReader, error) {
	l := peer.Default.GetLedger(channel)
	if l == nil {
		return nil, errors.Errorf("ledger not found for channel %s", channel)
	}

	return l.NewQueryExecutor()
}

