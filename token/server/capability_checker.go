
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456128410947584>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package server

import (
	"github.com/hyperledger/fabric/core/peer"
	"github.com/pkg/errors"
)

//go：生成伪造者-o mock/capability\u checker.go-forke name capability checker。能力检查人员

//CapabilityChecker用于检查通道是否支持令牌函数。
type CapabilityChecker interface {
	FabToken(channelId string) (bool, error)
}

//TokenCapabilityChecker实现CapabilityChecker接口
type TokenCapabilityChecker struct {
	PeerOps peer.Operations
}

func (c *TokenCapabilityChecker) FabToken(channelId string) (bool, error) {
	ac, ok := c.PeerOps.GetChannelConfig(channelId).ApplicationConfig()
	if !ok {
		return false, errors.Errorf("no application config found for channel %s", channelId)
	}
	return ac.Capabilities().FabToken(), nil
}

