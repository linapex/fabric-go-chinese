
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455965458042880>

/*
版权所有IBM Corp.2016保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

                 http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/


package config

import (
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/msp"
)

func nearIdentityHash(input []byte) []byte {
	return util.ConcatenateBytes([]byte("FakeHash("), input, []byte(""))
}

//
type Channel struct {
//
	HashingAlgorithmVal func([]byte) []byte
//
	BlockDataHashingStructureWidthVal uint32
//
	OrdererAddressesVal []string
//
	CapabilitiesVal channelconfig.ChannelCapabilities
}

//
func (scm *Channel) HashingAlgorithm() func([]byte) []byte {
	if scm.HashingAlgorithmVal == nil {
		return nearIdentityHash
	}
	return scm.HashingAlgorithmVal
}

//
func (scm *Channel) BlockDataHashingStructureWidth() uint32 {
	return scm.BlockDataHashingStructureWidthVal
}

//orderAddresses返回orderAddressesVal
func (scm *Channel) OrdererAddresses() []string {
	return scm.OrdererAddressesVal
}

//
func (scm *Channel) Capabilities() channelconfig.ChannelCapabilities {
	return scm.CapabilitiesVal
}

//
type ChannelCapabilities struct {
//SUPPORTEDER由SUPPORTED（）返回
	SupportedErr error

//
	MSPVersionVal msp.MSPVersion
}

//
func (cc *ChannelCapabilities) Supported() error {
	return cc.SupportedErr
}

//
func (cc *ChannelCapabilities) MSPVersion() msp.MSPVersion {
	return cc.MSPVersionVal
}

