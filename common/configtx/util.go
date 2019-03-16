
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455948408197120>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package configtx

import (
	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/utils"
)

//unmashalconfig尝试将字节解封为*cb.config
func UnmarshalConfig(data []byte) (*cb.Config, error) {
	config := &cb.Config{}
	err := proto.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

//unmashalconfigorpanic尝试将字节解封为*cb.config或出错时出现恐慌。
func UnmarshalConfigOrPanic(data []byte) *cb.Config {
	result, err := UnmarshalConfig(data)
	if err != nil {
		panic(err)
	}
	return result
}

//unmashalconfigupdate尝试将字节取消标记为*cb.configupdate
func UnmarshalConfigUpdate(data []byte) (*cb.ConfigUpdate, error) {
	configUpdate := &cb.ConfigUpdate{}
	err := proto.Unmarshal(data, configUpdate)
	if err != nil {
		return nil, err
	}
	return configUpdate, nil
}

//unmashalconfigupdate或panic尝试将字节解封为*cb.configupdate或出错时暂停
func UnmarshalConfigUpdateOrPanic(data []byte) *cb.ConfigUpdate {
	result, err := UnmarshalConfigUpdate(data)
	if err != nil {
		panic(err)
	}
	return result
}

//unmashalconfigupdatenedevelope尝试将字节解封为*cb.configupdate
func UnmarshalConfigUpdateEnvelope(data []byte) (*cb.ConfigUpdateEnvelope, error) {
	configUpdateEnvelope := &cb.ConfigUpdateEnvelope{}
	err := proto.Unmarshal(data, configUpdateEnvelope)
	if err != nil {
		return nil, err
	}
	return configUpdateEnvelope, nil
}

//unmashalconfigupdatenedevelopeorpanic尝试将字节unmashal为*cb.configupdatenedevelope或出错时出现的混乱
func UnmarshalConfigUpdateEnvelopeOrPanic(data []byte) *cb.ConfigUpdateEnvelope {
	result, err := UnmarshalConfigUpdateEnvelope(data)
	if err != nil {
		panic(err)
	}
	return result
}

//unmashalconfigendevelope尝试将字节解封为*cb.configendevelope
func UnmarshalConfigEnvelope(data []byte) (*cb.ConfigEnvelope, error) {
	configEnv := &cb.ConfigEnvelope{}
	err := proto.Unmarshal(data, configEnv)
	if err != nil {
		return nil, err
	}
	return configEnv, nil
}

//unmashalconfigendeveloperpanic尝试将字节解封为*cb.configendevelope或出错时的panics
func UnmarshalConfigEnvelopeOrPanic(data []byte) *cb.ConfigEnvelope {
	result, err := UnmarshalConfigEnvelope(data)
	if err != nil {
		panic(err)
	}
	return result
}

//unmashalconfigupdatefrompayload从给定的负载取消配置更新
func UnmarshalConfigUpdateFromPayload(payload *cb.Payload) (*cb.ConfigUpdate, error) {
	configEnv, err := UnmarshalConfigEnvelope(payload.Data)
	if err != nil {
		return nil, err
	}
	configUpdateEnv, err := utils.EnvelopeToConfigUpdate(configEnv.LastUpdate)
	if err != nil {
		return nil, err
	}

	return UnmarshalConfigUpdate(configUpdateEnv.ConfigUpdate)
}

