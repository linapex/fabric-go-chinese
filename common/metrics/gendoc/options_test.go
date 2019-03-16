
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455963155369984>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package gendoc_test

import (
	"github.com/hyperledger/fabric/common/metrics/gendoc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Options", func() {
	It("finds standard options", func() {
		f, err := ParseFile("testdata/basic.go")
		Expect(err).NotTo(HaveOccurred())
		Expect(f).NotTo(BeNil())

		options, err := gendoc.FileOptions(f)
		Expect(err).NotTo(HaveOccurred())
		Expect(options).To(HaveLen(3))
	})

	It("finds options that use named imports", func() {
		f, err := ParseFile("testdata/named_import.go")
		Expect(err).NotTo(HaveOccurred())
		Expect(f).NotTo(BeNil())

		options, err := gendoc.FileOptions(f)
		Expect(err).NotTo(HaveOccurred())
		Expect(options).To(HaveLen(3))
	})

	It("ignores variables that are tagged", func() {
		f, err := ParseFile("testdata/ignored.go")
		Expect(err).NotTo(HaveOccurred())
		Expect(f).NotTo(BeNil())

		options, err := gendoc.FileOptions(f)
		Expect(err).NotTo(HaveOccurred())
		Expect(options).To(BeEmpty())
	})
})

