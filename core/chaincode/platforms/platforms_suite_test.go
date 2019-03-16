
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:04</date>
//</624455990510620672>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package platforms_test

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/platforms"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go：生成仿冒者-o mock/platform.go——仿冒名称平台。平台
type platform interface {
	platforms.Platform
}

//go：生成伪造者-o mock/package-writer.go——伪造名称package writer。包装作家
type packageWriter interface {
	platforms.PackageWriter
}

func TestPlatforms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Platforms Suite")
}

