
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:08</date>
//</624456005048078336>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package container_test

import (
	"testing"

	"github.com/hyperledger/fabric/core/container"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go：生成伪造者-o mock/vm_provider.go--forke name vm provider。VM提供者
type vmProvider interface {
	container.VMProvider
}

//go：生成仿冒者-o mock/vm.go——仿冒名称vm。虚拟机
type vm interface {
	container.VM
}

//go：生成伪造者-o mock/vm req.go--forke name vmcreq。VMCREQ
type vmcReq interface {
	container.VMCReq
}

//go：生成伪造者-o mock/builder.go——伪造名称builder。建设者
type builder interface {
	container.Builder
}

func TestContainer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Container Suite")
}

