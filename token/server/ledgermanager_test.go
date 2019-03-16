
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456128595496960>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package server_test

import (
	"github.com/hyperledger/fabric/token/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LedgerManager", func() {
	var (
		ledgerManager *server.PeerLedgerManager
	)

	BeforeEach(func() {
		ledgerManager = &server.PeerLedgerManager{}
	})

	Context("when asking a LedgerReader for a channel that does not exists", func() {
		It("returns the error", func() {
			_, err := ledgerManager.GetLedgerReader("non-existing-channel")
			Expect(err).To(MatchError("ledger not found for channel non-existing-channel"))
		})
	})

})

