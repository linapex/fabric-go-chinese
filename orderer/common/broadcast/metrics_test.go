
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:28</date>
//</624456089085153280>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package broadcast_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hyperledger/fabric/orderer/common/broadcast"
	"github.com/hyperledger/fabric/orderer/common/broadcast/mock"
)

var _ = Describe("Metrics", func() {
	var (
		fakeProvider *mock.MetricsProvider
	)

	BeforeEach(func() {
		fakeProvider = &mock.MetricsProvider{}
		fakeProvider.NewHistogramReturns(&mock.MetricsHistogram{})
		fakeProvider.NewCounterReturns(&mock.MetricsCounter{})
	})

	It("uses the provider to initialize all fields", func() {
		metrics := broadcast.NewMetrics(fakeProvider)
		Expect(metrics).NotTo(BeNil())
		Expect(metrics.ValidateDuration).To(Equal(&mock.MetricsHistogram{}))
		Expect(metrics.EnqueueDuration).To(Equal(&mock.MetricsHistogram{}))
		Expect(metrics.ProcessedCount).To(Equal(&mock.MetricsCounter{}))

		Expect(fakeProvider.NewHistogramCallCount()).To(Equal(2))
		Expect(fakeProvider.NewCounterCallCount()).To(Equal(1))
	})
})

