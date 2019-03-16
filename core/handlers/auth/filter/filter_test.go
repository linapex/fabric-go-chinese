
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456010735554560>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package filter

import (
	"context"
	"testing"

	"github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/assert"
)

type mockEndorserServer struct {
	invoked bool
}

func (es *mockEndorserServer) ProcessProposal(context.Context, *peer.SignedProposal) (*peer.ProposalResponse, error) {
	es.invoked = true
	return nil, nil
}

func TestFilter(t *testing.T) {
	auth := NewFilter()
	nextEndorser := &mockEndorserServer{}
	auth.Init(nextEndorser)
	auth.ProcessProposal(nil, nil)
	assert.True(t, nextEndorser.invoked)
}

