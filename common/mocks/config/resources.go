
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455965781004288>

/*
版权所有IBM Corp.2016保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package config

import (
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/configtx"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/msp"
)

type Resources struct {
//
	ConfigtxValidatorVal configtx.Validator

//
	PolicyManagerVal policies.Manager

//
	ChannelConfigVal channelconfig.Channel

//
	OrdererConfigVal channelconfig.Orderer

//
	ApplicationConfigVal channelconfig.Application

//
	ConsortiumsConfigVal channelconfig.Consortiums

//mspmanagerval作为mspmanager（）的结果返回
	MSPManagerVal msp.MSPManager

//validatenewr作为validatenew的结果返回
	ValidateNewErr error
}

//
func (r *Resources) ConfigtxValidator() configtx.Validator {
	return r.ConfigtxValidatorVal
}

//
func (r *Resources) PolicyManager() policies.Manager {
	return r.PolicyManagerVal
}

//
func (r *Resources) ChannelConfig() channelconfig.Channel {
	return r.ChannelConfigVal
}

//
func (r *Resources) OrdererConfig() (channelconfig.Orderer, bool) {
	return r.OrdererConfigVal, r.OrdererConfigVal != nil
}

//
func (r *Resources) ApplicationConfig() (channelconfig.Application, bool) {
	return r.ApplicationConfigVal, r.ApplicationConfigVal != nil
}

func (r *Resources) ConsortiumsConfig() (channelconfig.Consortiums, bool) {
	return r.ConsortiumsConfigVal, r.ConsortiumsConfigVal != nil
}

//
func (r *Resources) MSPManager() msp.MSPManager {
	return r.MSPManagerVal
}

//validateNew返回validatener
func (r *Resources) ValidateNew(res channelconfig.Resources) error {
	return r.ValidateNewErr
}

