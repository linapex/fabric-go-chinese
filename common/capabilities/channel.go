
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455942519394304>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package capabilities

import (
	"github.com/hyperledger/fabric/msp"
	cb "github.com/hyperledger/fabric/protos/common"
)

const (
	channelTypeName = "Channel"

//channel v1_1是标准新的非向后兼容结构v1.1通道功能的功能字符串。
	ChannelV1_1 = "V1_1"

//channel v1_3是标准新的非向后兼容结构v1.3通道功能的功能字符串。
	ChannelV1_3 = "V1_3"
)

//ChannelProvider为通道级配置提供功能信息。
type ChannelProvider struct {
	*registry
	v11 bool
	v13 bool
}

//NewChannelProvider创建通道功能提供程序。
func NewChannelProvider(capabilities map[string]*cb.Capability) *ChannelProvider {
	cp := &ChannelProvider{}
	cp.registry = newRegistry(cp, capabilities)
	_, cp.v11 = capabilities[ChannelV1_1]
	_, cp.v13 = capabilities[ChannelV1_3]
	return cp
}

//类型返回用于日志记录的描述性字符串。
func (cp *ChannelProvider) Type() string {
	return channelTypeName
}

//如果此二进制文件支持此功能，则HasCapability返回true。
func (cp *ChannelProvider) HasCapability(capability string) bool {
	switch capability {
//在此处添加新功能名称
	case ChannelV1_3:
		return true
	case ChannelV1_1:
		return true
	default:
		return false
	}
}

//mspversion返回此通道所需的MSP支持级别。
func (cp *ChannelProvider) MSPVersion() msp.MSPVersion {
	switch {
	case cp.v13:
		return msp.MSPv1_3
	case cp.v11:
		return msp.MSPv1_1
	default:
		return msp.MSPv1_0
	}
}

