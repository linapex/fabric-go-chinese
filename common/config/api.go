
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455947544170496>

/*
版权所有IBM Corp.2017保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package config

import (
	cb "github.com/hyperledger/fabric/protos/common"
)

//config封装config（通道或资源）树
type Config interface {
//configproto返回当前配置
	ConfigProto() *cb.Config

//ProposeConfigUpdate尝试根据当前配置状态验证新的configtx
	ProposeConfigUpdate(configtx *cb.Envelope) (*cb.ConfigEnvelope, error)
}

//管理器提供对资源配置的访问
type Manager interface {
//getchannelconfig定义与通道配置相关的方法
	GetChannelConfig(channel string) Config
}

