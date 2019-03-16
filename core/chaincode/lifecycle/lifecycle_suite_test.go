
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:02</date>
//</624455981597724672>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package lifecycle_test

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/lifecycle"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go：生成伪造者-o mock/chaincode_stub.go——伪造名称chaincode stub。链状短截线
type chaincodeStub interface {
	shim.ChaincodeStubInterface
}

//go：生成仿冒者-o mock/chaincode\store.go——仿冒名称chaincode store。链式存储器
type chaincodeStore interface {
	lifecycle.ChaincodeStore
}

//go：生成伪造者-o mock/package_parser.go--forke name package parser。打包分析器
type packageParser interface {
	lifecycle.PackageParser
}

//go：生成仿冒者-o mock/scc_functions.go--fake name scc functions。SCC函数
type sccFunctions interface {
	lifecycle.SCCFunctions
}

func TestLifecycle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lifecycle Suite")
}

