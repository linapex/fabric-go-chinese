
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455942376787968>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package capabilities

import (
	"github.com/hyperledger/fabric/common/flogging"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/pkg/errors"
)

var logger = flogging.MustGetLogger("common.capabilities")

//提供程序是注册表的“plugin”参数。
type provider interface {
//hasCapability应该报告二进制文件是否支持此功能。
	HasCapability(capability string) bool

//类型用于使错误消息更清晰。
	Type() string
}

//注册表是一种公共结构，用于支持功能的特定方面。
//例如订购者、应用程序和通道。
type registry struct {
	provider     provider
	capabilities map[string]*cb.Capability
}

func newRegistry(p provider, capabilities map[string]*cb.Capability) *registry {
	return &registry{
		provider:     p,
		capabilities: capabilities,
	}
}

//支持检查此二进制文件是否支持所有必需的功能。
func (r *registry) Supported() error {
	for capabilityName := range r.capabilities {
		if r.provider.HasCapability(capabilityName) {
			logger.Debugf("%s capability %s is supported and is enabled", r.provider.Type(), capabilityName)
			continue
		}

		return errors.Errorf("%s capability %s is required but not supported", r.provider.Type(), capabilityName)
	}
	return nil
}

