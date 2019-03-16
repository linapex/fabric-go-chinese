
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455946063581184>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channelconfig

import (
	cb "github.com/hyperledger/fabric/protos/common"
)

const (
//consortiumsgroupkey是consortiums配置的组名
	ConsortiumsGroupKey = "Consortiums"
)

//CONSORITUMSCONFIG保存联盟配置信息
type ConsortiumsConfig struct {
	consortiums map[string]Consortium
}

//newconsortiumsconfig创建consoritums配置的新实例
func NewConsortiumsConfig(consortiumsGroup *cb.ConfigGroup, mspConfig *MSPConfigHandler) (*ConsortiumsConfig, error) {
	cc := &ConsortiumsConfig{
		consortiums: make(map[string]Consortium),
	}

	for consortiumName, consortiumGroup := range consortiumsGroup.Groups {
		var err error
		if cc.consortiums[consortiumName], err = NewConsortiumConfig(consortiumGroup, mspConfig); err != nil {
			return nil, err
		}
	}
	return cc, nil
}

//联合体返回当前联合体的映射
func (cc *ConsortiumsConfig) Consortiums() map[string]Consortium {
	return cc.consortiums
}

