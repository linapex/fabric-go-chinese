
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:02</date>
//</624455982356893696>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package lifecycle_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/lifecycle"
	lc "github.com/hyperledger/fabric/protos/peer/lifecycle"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProtobufImpl", func() {
	var (
		pi        *lifecycle.ProtobufImpl
		sampleMsg *lc.InstallChaincodeArgs
	)

	BeforeEach(func() {
		pi = &lifecycle.ProtobufImpl{}
		sampleMsg = &lc.InstallChaincodeArgs{
			Name:                    "name",
			Version:                 "version",
			ChaincodeInstallPackage: []byte("install-package"),
		}
	})

	Describe("Marshal", func() {
		It("passes through to the proto implementation", func() {
			res, err := pi.Marshal(sampleMsg)
			Expect(err).NotTo(HaveOccurred())

			msg := &lc.InstallChaincodeArgs{}
			err = proto.Unmarshal(res, msg)
			Expect(err).NotTo(HaveOccurred())
			Expect(proto.Equal(msg, sampleMsg)).To(BeTrue())
		})
	})

	Describe("Unmarshal", func() {
		It("passes through to the proto implementation", func() {
			res, err := proto.Marshal(sampleMsg)
			Expect(err).NotTo(HaveOccurred())

			msg := &lc.InstallChaincodeArgs{}
			err = pi.Unmarshal(res, msg)
			Expect(err).NotTo(HaveOccurred())
			Expect(proto.Equal(msg, sampleMsg)).To(BeTrue())
		})
	})
})

