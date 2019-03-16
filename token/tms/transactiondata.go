
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:38</date>
//</624456133158899712>

/*
版权所有IBM Corp.2017保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package tms

import "github.com/hyperledger/fabric/protos/token"

//TransactionData结构包含令牌事务和结构事务ID。
//令牌事务由验证方对等方创建，但仅创建事务ID
//稍后由客户机执行（使用令牌事务和nonce）。在验证和提交时
//提交对等机需要时间、令牌事务和事务ID。
//将它们存储在一个结构中有助于处理它们。
type TransactionData struct {
	Tx *token.TokenTransaction
//结构事务ID
	TxID string
}

