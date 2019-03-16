
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:38</date>
//</624456130130612224>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package server_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

func ProtoMarshal(m proto.Message) []byte {
	bytes, err := proto.Marshal(m)
	Expect(err).NotTo(HaveOccurred())

	return bytes
}

