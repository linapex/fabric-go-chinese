
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455947183460352>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channelconfig_test

import (
	"testing"

	newchannelconfig "github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/tools/configtxgen/configtxgentest"
	"github.com/hyperledger/fabric/common/tools/configtxgen/encoder"
	genesisconfig "github.com/hyperledger/fabric/common/tools/configtxgen/localconfig"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/stretchr/testify/assert"
)

func TestWithRealConfigtx(t *testing.T) {
	conf := configtxgentest.Load(genesisconfig.SampleSingleMSPSoloProfile)

//没有一个示例配置文件定义应用程序配置节
//在一个创世时期（因为这是个坏主意），但我们把它们结合起来
//这里是为了更好地练习代码。
	conf.Application = &genesisconfig.Application{
		Organizations: []*genesisconfig.Organization{
			conf.Orderer.Organizations[0],
		},
	}
	conf.Application.Organizations[0].AnchorPeers = []*genesisconfig.AnchorPeer{
		{
			Host: "foo",
			Port: 7,
		},
	}
	gb := encoder.New(conf).GenesisBlockForChannel("foo")
	env := utils.ExtractEnvelopeOrPanic(gb, 0)
	_, err := newchannelconfig.NewBundleFromEnvelope(env)
	assert.NoError(t, err)
}

