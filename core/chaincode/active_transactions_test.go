
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455978384887808>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package chaincode_test

import (
	"github.com/hyperledger/fabric/core/chaincode"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("ActiveTransactions", func() {
	var activeTx *chaincode.ActiveTransactions

	BeforeEach(func() {
		activeTx = chaincode.NewActiveTransactions()
	})

	It("tracks active transactions", func() {
//添加唯一交易记录
		ok := activeTx.Add("channel-id", "tx-id")
		Expect(ok).To(BeTrue(), "a new transaction should return true")
		ok = activeTx.Add("channel-id", "tx-id-2")
		Expect(ok).To(BeTrue(), "adding a different transaction id should return true")
		ok = activeTx.Add("channel-id-2", "tx-id")
		Expect(ok).To(BeTrue(), "adding a different channel-id should return true")

//尝试添加已存在的事务
		ok = activeTx.Add("channel-id", "tx-id")
		Expect(ok).To(BeFalse(), "attempting to an existing transaction should return false")

//删除现有的并确保ID可以重用
		activeTx.Remove("channel-id", "tx-id")
		ok = activeTx.Add("channel-id", "tx-id")
		Expect(ok).To(BeTrue(), "using a an id that has been removed should return true")
	})

	DescribeTable("NewTxKey",
		func(channelID, txID, expected string) {
			result := chaincode.NewTxKey(channelID, txID)
			Expect(result).To(Equal(expected))
		},
		Entry("empty channel and tx", "", "", ""),
		Entry("empty channel", "", "tx-1", "tx-1"),
		Entry("empty tx", "chan-1", "tx-1", "chan-1tx-1"),
		Entry("channel and tx", "chan-1", "tx-1", "chan-1tx-1"),
	)
})

