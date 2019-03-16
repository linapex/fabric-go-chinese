
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456128255758336>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package server

import (
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/token"
	"github.com/pkg/errors"
)

//go:生成伪造者-o mock/acl_provider.go-fake name acl provider。ACL提供者

type ACLProvider interface {
//checkacl检查给定通道的资源访问控制。
//IDinfo是一个对象，如[]*common.signedData，从中
//可以提取一个ID以根据策略进行测试
	CheckACL(resName string, channelID string, idinfo interface{}) error
}

type ACLResources struct {
	IssueTokens    string
	TransferTokens string
	ListTokens     string
}

//PolicyBasedAccessControl实现令牌命令访问控制功能。
type PolicyBasedAccessControl struct {
	ACLProvider  ACLProvider
	ACLResources *ACLResources
}

func (ac *PolicyBasedAccessControl) Check(sc *token.SignedCommand, c *token.Command) error {
	signedData := []*common.SignedData{{
		Identity:  c.Header.Creator,
		Data:      sc.Command,
		Signature: sc.Signature,
	}}

	switch t := c.GetPayload().(type) {

	case *token.Command_ImportRequest:
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.IssueTokens,
			c.Header.ChannelId,
			signedData,
		)
	case *token.Command_ListRequest:
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.ListTokens,
			c.Header.ChannelId,
			signedData,
		)
	case *token.Command_TransferRequest:
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.TransferTokens,
			c.Header.ChannelId,
			signedData,
		)
	case *token.Command_RedeemRequest:
//兑换与转账具有相同的政策
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.TransferTokens,
			c.Header.ChannelId,
			signedData,
		)

	case *token.Command_ApproveRequest:
//批准与转移具有相同的策略
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.TransferTokens,
			c.Header.ChannelId,
			signedData,
		)

	case *token.Command_TransferFromRequest:
//TransferFrom与Transfer具有相同的策略
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.TransferTokens,
			c.Header.ChannelId,
			signedData,
		)

	case *token.Command_ExpectationRequest:
		if c.GetExpectationRequest().GetExpectation() == nil {
			return errors.New("ExpectationRequest has nil Expectation")
		}
		plainExpectation := c.GetExpectationRequest().GetExpectation().GetPlainExpectation()
		if plainExpectation == nil {
			return errors.New("ExpectationRequest has nil PlainExpectation")
		}
		return ac.checkExpectation(plainExpectation, signedData, c)
	default:
		return errors.Errorf("command type not recognized: %T", t)
	}
}

//根据期望中的有效负载类型，检查期望检查问题策略或传输策略
func (ac *PolicyBasedAccessControl) checkExpectation(plainExpectation *token.PlainExpectation, signedData []*common.SignedData, c *token.Command) error {
	switch t := plainExpectation.GetPayload().(type) {
	case *token.PlainExpectation_ImportExpectation:
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.IssueTokens,
			c.Header.ChannelId,
			signedData,
		)
	case *token.PlainExpectation_TransferExpectation:
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.TransferTokens,
			c.Header.ChannelId,
			signedData,
		)
	default:
		return errors.Errorf("expectation payload type not recognized: %T", t)
	}
}

