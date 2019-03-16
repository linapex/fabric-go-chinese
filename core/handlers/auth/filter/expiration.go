
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456010387427328>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package filter

import (
	"context"
	"time"

	"github.com/hyperledger/fabric/common/crypto"
	"github.com/hyperledger/fabric/core/handlers/auth"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/pkg/errors"
)

//NewExpirationCheckFilter创建检查标识过期的新筛选器
func NewExpirationCheckFilter() auth.Filter {
	return &expirationCheckFilter{}
}

type expirationCheckFilter struct {
	next peer.EndorserServer
}

//init用下一个背书服务器初始化筛选器
func (f *expirationCheckFilter) Init(next peer.EndorserServer) {
	f.next = next
}

func validateProposal(signedProp *peer.SignedProposal) error {
	prop, err := utils.GetProposal(signedProp.ProposalBytes)
	if err != nil {
		return errors.Wrap(err, "failed parsing proposal")
	}

	hdr, err := utils.GetHeader(prop.Header)
	if err != nil {
		return errors.Wrap(err, "failed parsing header")
	}

	sh, err := utils.GetSignatureHeader(hdr.SignatureHeader)
	if err != nil {
		return errors.Wrap(err, "failed parsing signature header")
	}
	expirationTime := crypto.ExpiresAt(sh.Creator)
	if !expirationTime.IsZero() && time.Now().After(expirationTime) {
		return errors.New("identity expired")
	}
	return nil
}

//处理建议处理已签名的建议
func (f *expirationCheckFilter) ProcessProposal(ctx context.Context, signedProp *peer.SignedProposal) (*peer.ProposalResponse, error) {
	if err := validateProposal(signedProp); err != nil {
		return nil, err
	}
	return f.next.ProcessProposal(ctx, signedProp)
}

