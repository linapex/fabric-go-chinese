
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455946810167296>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channelconfig

import (
	"testing"

	"github.com/hyperledger/fabric/core/config/configtest"
	"github.com/hyperledger/fabric/msp"
	mspprotos "github.com/hyperledger/fabric/protos/msp"
	"github.com/stretchr/testify/assert"
)

func TestMSPConfigManager(t *testing.T) {
	mspDir, err := configtest.GetDevMspDir()
	assert.NoError(t, err)
	conf, err := msp.GetLocalMspConfig(mspDir, nil, "SampleOrg")
	assert.NoError(t, err)

//测试成功：

	mspVers := []msp.MSPVersion{msp.MSPv1_0, msp.MSPv1_1}

	for _, ver := range mspVers {
		mspCH := NewMSPConfigHandler(ver)

		_, err = mspCH.ProposeMSP(conf)
		assert.NoError(t, err)

		mgr, err := mspCH.CreateMSPManager()
		assert.NoError(t, err)
		assert.NotNil(t, mgr)

		msps, err := mgr.GetMSPs()
		assert.NoError(t, err)

		if msps == nil || len(msps) == 0 {
			t.Fatalf("There are no MSPS in the manager")
		}

		for _, mspInst := range msps {
			assert.Equal(t, mspInst.GetVersion(), msp.MSPVersion(ver))
		}
	}
}

func TestMSPConfigFailure(t *testing.T) {
	mspCH := NewMSPConfigHandler(msp.MSPv1_0)

//开始/提议/承诺
	t.Run("Bad proto", func(t *testing.T) {
		_, err := mspCH.ProposeMSP(&mspprotos.MSPConfig{Config: []byte("BARF!")})
		assert.Error(t, err)
	})

	t.Run("Bad MSP Type", func(t *testing.T) {
		_, err := mspCH.ProposeMSP(&mspprotos.MSPConfig{Type: int32(10)})
		assert.Error(t, err)
	})
}

