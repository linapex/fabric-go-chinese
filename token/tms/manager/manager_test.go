
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:38</date>
//</624456130466156544>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package manager_test

import (
	"github.com/hyperledger/fabric/token/identity/mock"
	"github.com/hyperledger/fabric/token/tms/manager"
	"github.com/hyperledger/fabric/token/tms/plain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("Manager", func() {
	var (
		mgm                             *manager.Manager
		fakeIdentityDeserializerManager *mock.DeserializerManager
	)

	BeforeEach(func() {
		fakeIdentityDeserializerManager = &mock.DeserializerManager{}
		mgm = &manager.Manager{IdentityDeserializerManager: fakeIdentityDeserializerManager}
	})

	Describe("Get a TxProcessor for a non-existent channel", func() {
		BeforeEach(func() {
			fakeIdentityDeserializerManager.DeserializerReturns(nil, errors.New("GetDeserializerReturns no-way-man"))
		})
		It("returns an error", func() {
			_, err := mgm.GetTxProcessor("boguschannel")
			Expect(err.Error()).To(Equal("failed getting identity deserialiser manager for channel 'boguschannel': GetDeserializerReturns no-way-man"))
		})
	})

	Context("When a channel exists", func() {
		var (
			fakeIdentityDeserializer *mock.Deserializer
			channel                  string
		)
		BeforeEach(func() {
			channel = "ch0"
			fakeIdentityDeserializer = &mock.Deserializer{}
			fakeIdentityDeserializerManager.DeserializerReturns(fakeIdentityDeserializer, nil)
		})

		Describe("Get a TxProcessor for an existing channel", func() {
			It("returns a Verifier that implements the TxProcessor interface", func() {
				txProcessor, err := mgm.GetTxProcessor(channel)
				Expect(err).NotTo(HaveOccurred())
				Expect(txProcessor).NotTo(BeNil())
				Expect(txProcessor).To(Equal(&plain.Verifier{IssuingValidator: &manager.AllIssuingValidator{Deserializer: fakeIdentityDeserializer}}))
			})
		})
	})
})

var _ = Describe("FabricIdentityDeserializerManager", func() {
	Describe("Get an IdentityDeserializer for a non-existent channel", func() {
		var (
			fabricIdentityDeserializerManager *manager.FabricIdentityDeserializerManager
		)
		BeforeEach(func() {
			fabricIdentityDeserializerManager = &manager.FabricIdentityDeserializerManager{}
		})
		It("returns an error", func() {
			_, err := fabricIdentityDeserializerManager.Deserializer("boguschannel")
			Expect(err).To(MatchError("channel not found"))
		})
	})
})

