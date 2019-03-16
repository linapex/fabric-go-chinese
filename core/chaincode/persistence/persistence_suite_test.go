
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:03</date>
//</624455987100651520>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package persistence_test

import (
	"os"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/persistence"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
type ioReadWriter interface {
	persistence.IOReadWriter
}

//go：生成伪造者-o mock/osfileinfo.go-forke name osfileinfo。OsFielFig
type osFileInfo interface {
	os.FileInfo
}

//go:生成仿冒者-o mock/store_package_provider.go-fake name store package provider。存储包提供程序
type storePackageProvider interface {
	persistence.StorePackageProvider
}

//go：生成仿冒者-o mock/legacy package provider.go-仿冒名称legacy package provider。LegacyPackageProvider（legacyPackageProvider）
type legacyPackageProvider interface {
	persistence.LegacyPackageProvider
}

//go：生成仿冒者-o mock/package_parser.go-forke name package parser。打包分析器
type packageParser interface {
	persistence.PackageParser
}

func TestPersistence(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Persistence Suite")
}

