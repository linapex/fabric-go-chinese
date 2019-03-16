
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:31</date>
//</624456103702302720>

//版权所有IBM公司。保留所有权利。
//SPDX许可证标识符：Apache-2.0

package main

import (
	"github.com/hyperledger/fabric/common/localmsp"
	"github.com/hyperledger/fabric/common/tools/configtxgen/encoder"
	genesisconfig "github.com/hyperledger/fabric/common/tools/configtxgen/localconfig"
	cb "github.com/hyperledger/fabric/protos/common"
)

func newChainRequest(consensusType, creationPolicy, newChannelID string) *cb.Envelope {
	env, err := encoder.MakeChannelCreationTransaction(newChannelID, localmsp.NewSigner(), genesisconfig.Load(genesisconfig.SampleSingleMSPChannelProfile))
	if err != nil {
		panic(err)
	}
	return env
}

