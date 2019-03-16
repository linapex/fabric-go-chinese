
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456010882355200>

/*
版权所有IBM Corp，SecureKey Technologies Inc.保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package main

import (
	"context"

	"github.com/hyperledger/fabric/core/handlers/auth"
	"github.com/hyperledger/fabric/protos/peer"
)

//新建筛选器创建新筛选器
func NewFilter() auth.Filter {
	return &filter{}
}

type filter struct {
	next peer.EndorserServer
}

//init用下一个背书服务器初始化筛选器
func (f *filter) Init(next peer.EndorserServer) {
	f.next = next
}

//处理建议处理已签名的建议
func (f *filter) ProcessProposal(ctx context.Context, signedProp *peer.SignedProposal) (*peer.ProposalResponse, error) {
	return f.next.ProcessProposal(ctx, signedProp)
}

func main() {
}

