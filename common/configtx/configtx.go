
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455948047486976>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package configtx

import (
	cb "github.com/hyperledger/fabric/protos/common"
)

//验证程序提供了一种机制来建议配置更新，请参见配置更新结果
//并验证配置更新的结果。
type Validator interface {
//验证应用configtx成为新配置的尝试
	Validate(configEnv *cb.ConfigEnvelope) error

//验证针对当前配置状态验证新configtx的尝试
	ProposeConfigUpdate(configtx *cb.Envelope) (*cb.ConfigEnvelope, error)

//chainID检索与此管理器关联的链ID
	ChainID() string

//config proto以proto的形式返回当前配置
	ConfigProto() *cb.Config

//Sequence返回配置的当前序列号
	Sequence() uint64
}

