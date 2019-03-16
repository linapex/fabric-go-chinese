
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:56</date>
//</624455954796122112>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package grpclogging_test

import (
	"context"
	"time"

	"github.com/hyperledger/fabric/common/grpclogging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ = Describe("Context", func() {
	var inputFields []zapcore.Field

	BeforeEach(func() {
		inputFields = []zapcore.Field{
			zap.String("string-key", "string-value"),
			zap.Duration("duration-key", time.Second),
			zap.Int("int-key", 42),
		}
	})

	It("decorates a context with fields", func() {
		ctx := grpclogging.WithFields(context.Background(), inputFields)
		Expect(ctx).NotTo(Equal(context.Background()))

		fields := grpclogging.Fields(ctx)
		Expect(fields).NotTo(BeEmpty())
	})

	It("extracts fields from a decorated context as a slice of zapcore.Field", func() {
		ctx := grpclogging.WithFields(context.Background(), inputFields)

		fields := grpclogging.Fields(ctx)
		Expect(fields).To(ConsistOf(inputFields))
	})

	It("extracts fields from a decorated context as a slice of interface{}", func() {
		ctx := grpclogging.WithFields(context.Background(), inputFields)

		zapFields := grpclogging.ZapFields(ctx)
		Expect(zapFields).To(Equal(inputFields))
	})

	It("returns the same fields regardless of type", func() {
		ctx := grpclogging.WithFields(context.Background(), inputFields)

		fields := grpclogging.Fields(ctx)
		zapFields := grpclogging.ZapFields(ctx)
		Expect(zapFields).To(ConsistOf(fields))
	})

	Context("when the context isn't decorated", func() {
		It("returns no fields", func() {
			fields := grpclogging.Fields(context.Background())
			Expect(fields).To(BeNil())

			zapFields := grpclogging.ZapFields(context.Background())
			Expect(zapFields).To(BeNil())
		})
	})
})

