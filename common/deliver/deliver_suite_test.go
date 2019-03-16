
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:55</date>
//</624455949939118080>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package deliver_test

import (
	"testing"

	"github.com/hyperledger/fabric/common/deliver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go：生成仿冒者-o模拟/过滤的响应\u sender.go-fake name filtered response sender。筛选响应发送器
type filteredResponseSender interface {
	deliver.ResponseSender
	deliver.Filtered
}

func TestDeliver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deliver Suite")
}

