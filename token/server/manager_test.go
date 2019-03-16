
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456128729714688>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package server_test

import (
	"errors"

	"github.com/hyperledger/fabric/token/ledger/mock"
	"github.com/hyperledger/fabric/token/server"
	"github.com/hyperledger/fabric/token/tms/plain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager", func() {
	Describe("GetIssuer", func() {
		It("returns a plain issuer", func() {
			Manager := &server.Manager{}
			issuer, err := Manager.GetIssuer("test-channel", []byte("private-credential"), []byte("public-credential"))
			Expect(err).NotTo(HaveOccurred())
			Expect(issuer).To(Equal(&plain.Issuer{}))
		})
	})

	Describe("GetTransactor", func() {
		var (
			fakeLedgerReader  *mock.LedgerReader
			fakeLedgerManager *mock.LedgerManager
		)

		BeforeEach(func() {
			fakeLedgerReader = &mock.LedgerReader{}
			fakeLedgerManager = &mock.LedgerManager{}
		})

		It("returns a plain transactor", func() {
			manager := &server.Manager{LedgerManager: fakeLedgerManager}
			fakeLedgerManager.GetLedgerReaderReturns(fakeLedgerReader, nil)
			transactor, err := manager.GetTransactor("test-channel", []byte("private-credential"), []byte("public-credential"))
			Expect(err).NotTo(HaveOccurred())
			Expect(transactor).To(Equal(&plain.Transactor{Ledger: fakeLedgerReader, PublicCredential: []byte("public-credential")}))
		})
		It("returns an error", func() {
			manager := &server.Manager{LedgerManager: fakeLedgerManager}
			fakeLedgerManager.GetLedgerReaderReturns(nil, errors.New("banana ledger"))
			transactor, err := manager.GetTransactor("test-channel", []byte("private-credential"), []byte("public-credential"))
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed getting ledger for channel: test-channel: banana ledger"))
			Expect(transactor).To(BeNil())
		})
	})
})

