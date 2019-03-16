
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:59</date>
//</624455969451020288>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package integration_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var configtxlatorPath string

var _ = SynchronizedBeforeSuite(func() []byte {
	configtxlatorPath, err := gexec.Build("github.com/hyperledger/fabric/common/tools/configtxlator")
	Expect(err).NotTo(HaveOccurred())

	return []byte(configtxlatorPath)
}, func(payload []byte) {
	configtxlatorPath = string(payload)
})

var _ = SynchronizedAfterSuite(func() {
}, func() {
	gexec.CleanupBuildArtifacts()
})

