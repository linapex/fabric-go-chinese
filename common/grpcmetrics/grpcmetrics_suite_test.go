
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:56</date>
//</624455956851331072>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package grpcmetrics_test

import (
	"testing"

	"github.com/hyperledger/fabric/common/grpcmetrics/testpb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate protoc--proto_path=$go path/src/github.com/hyperledger/fabric/common/grpcmetrics/testpb--go_out=plugins=grpc:$go path/src$go path/src/github.com/hyperledger/fabric/common/grpcmetrics/testpb/echo.proto

func TestGrpcmetrics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grpcmetrics Suite")
}

//go：生成仿冒者-o fakes/echo_service.go--echoserviceserver的假名。EchoServiceServer（EchoServiceServer）

type echoServiceServer interface {
	testpb.EchoServiceServer
}

