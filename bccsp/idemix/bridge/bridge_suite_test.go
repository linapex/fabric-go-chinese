
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:49</date>
//</624455928455892992>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package bridge_test

import (
	"testing"

	"github.com/hyperledger/fabric-amcl/amcl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPlain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Plain Suite")
}

//newrandanpanic是一个实用程序测试函数，调用时总是死机
func NewRandPanic() *amcl.RAND {
	panic("new rand panic")
}

