
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:16</date>
//</624456038522818560>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package peer

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/protos/common"
)

//computefullconfig计算给定当前资源束和事务（包含增量）的完整资源配置。
func computeFullConfig(currentConfigBundle *channelconfig.Bundle, channelConfTx *common.Envelope) (*common.Config, error) {
	fullChannelConfigEnv, err := currentConfigBundle.ConfigtxValidator().ProposeConfigUpdate(channelConfTx)
	if err != nil {
		return nil, err
	}
	return fullChannelConfigEnv.Config, nil
}

//TODO最好使序列化/反序列化具有确定性
func serialize(resConfig *common.Config) ([]byte, error) {
	return proto.Marshal(resConfig)
}

func deserialize(serializedConf []byte) (*common.Config, error) {
	conf := &common.Config{}
	if err := proto.Unmarshal(serializedConf, conf); err != nil {
		return nil, err
	}
	return conf, nil
}

//RetrievePersistedChannelConfig从StateDB检索持久化通道配置
func retrievePersistedChannelConfig(ledger ledger.PeerLedger) (*common.Config, error) {
	qe, err := ledger.NewQueryExecutor()
	if err != nil {
		return nil, err
	}
	defer qe.Done()
	return retrievePersistedConf(qe, channelConfigKey)
}

