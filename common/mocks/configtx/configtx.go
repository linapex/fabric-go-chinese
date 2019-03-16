
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455965927804928>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package configtx

import (
	cb "github.com/hyperledger/fabric/protos/common"
)

//
type Validator struct {
//
	ChainIDVal string

//
	SequenceVal uint64

//ApplyVal由Apply返回
	ApplyVal error

//
	AppliedConfigUpdateEnvelope *cb.ConfigEnvelope

//validateval由validate返回
	ValidateVal error

//
	ProposeConfigUpdateError error

//
	ProposeConfigUpdateVal *cb.ConfigEnvelope

//configProtoval作为configProtoval（）的值返回
	ConfigProtoVal *cb.Config
}

//
func (cm *Validator) ConfigProto() *cb.Config {
	return cm.ConfigProtoVal
}

//
func (cm *Validator) ChainID() string {
	return cm.ChainIDVal
}

//batchsize返回batchsizeval
func (cm *Validator) Sequence() uint64 {
	return cm.SequenceVal
}

//建议配置更新
func (cm *Validator) ProposeConfigUpdate(update *cb.Envelope) (*cb.ConfigEnvelope, error) {
	return cm.ProposeConfigUpdateVal, cm.ProposeConfigUpdateError
}

//
func (cm *Validator) Apply(configEnv *cb.ConfigEnvelope) error {
	cm.AppliedConfigUpdateEnvelope = configEnv
	return cm.ApplyVal
}

//validate返回validateval
func (cm *Validator) Validate(configEnv *cb.ConfigEnvelope) error {
	return cm.ValidateVal
}

