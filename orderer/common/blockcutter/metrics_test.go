
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:27</date>
//</624456088019800064>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package blockcutter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hyperledger/fabric/orderer/common/blockcutter"
	"github.com/hyperledger/fabric/orderer/common/blockcutter/mock"
)

var _ = Describe("Metrics", func() {
	Describe("NewMetrics", func() {
		var (
			fakeProvider *mock.MetricsProvider
		)

		BeforeEach(func() {
			fakeProvider = &mock.MetricsProvider{}
			fakeProvider.NewHistogramReturns(&mock.MetricsHistogram{})
		})

		It("uses the provider to initialize its field", func() {
			metrics := blockcutter.NewMetrics(fakeProvider)
			Expect(metrics).NotTo(BeNil())
			Expect(metrics.BlockFillDuration).To(Equal(&mock.MetricsHistogram{}))

			Expect(fakeProvider.NewHistogramCallCount()).To(Equal(1))
		})
	})
})

