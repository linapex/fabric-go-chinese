
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:29</date>
//</624456096018337792>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package multichannel

import (
	"testing"

	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/deliver/mock"
	"github.com/hyperledger/fabric/common/ledger/blockledger/mocks"
	"github.com/hyperledger/fabric/common/mocks/config"
	"github.com/hyperledger/fabric/common/mocks/configtx"
	mockpolicies "github.com/hyperledger/fabric/common/mocks/policies"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/common/tools/configtxgen/configtxgentest"
	"github.com/hyperledger/fabric/common/tools/configtxgen/encoder"
	"github.com/hyperledger/fabric/common/tools/configtxgen/localconfig"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/orderer"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestChainSupportBlock(t *testing.T) {
	ledger := &mocks.ReadWriter{}
	ledger.On("Height").Return(uint64(100))
	iterator := &mock.BlockIterator{}
	iterator.NextReturns(&common.Block{Header: &common.BlockHeader{Number: 99}}, common.Status_SUCCESS)
	ledger.On("Iterator", &orderer.SeekPosition{
		Type: &orderer.SeekPosition_Specified{
			Specified: &orderer.SeekSpecified{Number: 99},
		},
	}).Return(iterator, uint64(99))
	cs := &ChainSupport{ledgerResources: &ledgerResources{ReadWriter: ledger}}

	assert.Nil(t, cs.Block(100))
	assert.Equal(t, uint64(99), cs.Block(99).Header.Number)
}

type mutableResourcesMock struct {
	config.Resources
}

func (*mutableResourcesMock) Update(*channelconfig.Bundle) {
	panic("implement me")
}

func TestVerifyBlockSignature(t *testing.T) {
	policyMgr := &mockpolicies.Manager{
		PolicyMap: make(map[string]policies.Policy),
	}
	ms := &mutableResourcesMock{
		Resources: config.Resources{
			ConfigtxValidatorVal: &configtx.Validator{ChainIDVal: "mychannel"},
			PolicyManagerVal:     policyMgr,
		},
	}
	cs := &ChainSupport{
		ledgerResources: &ledgerResources{
			configResources: &configResources{
				mutableResources: ms,
			},
		},
	}

//方案一：策略管理器未初始化
//因此找不到政策
	err := cs.VerifyBlockSignature([]*common.SignedData{}, nil)
	assert.EqualError(t, err, "policy /Channel/Orderer/BlockValidation wasn't found")

//场景二：策略管理器找到策略，但它评估
//出错。
	policyMgr.PolicyMap["/Channel/Orderer/BlockValidation"] = &mockpolicies.Policy{
		Err: errors.New("invalid signature"),
	}
	err = cs.VerifyBlockSignature([]*common.SignedData{}, nil)
	assert.EqualError(t, err, "block verification failed: invalid signature")

//场景三：策略管理器找到策略，并对成功进行评估
	policyMgr.PolicyMap["/Channel/Orderer/BlockValidation"] = &mockpolicies.Policy{
		Err: nil,
	}
	assert.NoError(t, cs.VerifyBlockSignature([]*common.SignedData{}, nil))

//场景四：传递了一个错误的配置信封
	err = cs.VerifyBlockSignature([]*common.SignedData{}, &common.ConfigEnvelope{})
	assert.EqualError(t, err, "channelconfig Config cannot be nil")

//场景V：传递有效的配置信封
	assert.NoError(t, cs.VerifyBlockSignature([]*common.SignedData{}, testConfigEnvelope(t)))

}

func testConfigEnvelope(t *testing.T) *common.ConfigEnvelope {
	config := configtxgentest.Load(localconfig.SampleInsecureSoloProfile)
	group, err := encoder.NewChannelGroup(config)
	assert.NoError(t, err)
	assert.NotNil(t, group)
	return &common.ConfigEnvelope{
		Config: &common.Config{
			ChannelGroup: group,
		},
	}
}

