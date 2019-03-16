
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456110882951168>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package common_test

import (
	"testing"

	"github.com/hyperledger/fabric/peer/common"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	assert := assert.New(t)

//失败-文件名为空
	networkConfig, err := common.GetConfig("")
	assert.Error(err)
	assert.Nil(networkConfig)

//失败-文件不存在
	networkConfig, err = common.GetConfig("fakefile.yaml")
	assert.Error(err)
	assert.Nil(networkConfig)

//失败-连接配置文件中一些bool的意外值
	networkConfig, err = common.GetConfig("testdata/connectionprofile-bad.yaml")
	assert.Error(err, "error should have been nil")
	assert.Nil(networkConfig, "network config should be set")

//成功
	networkConfig, err = common.GetConfig("testdata/connectionprofile.yaml")
	assert.NoError(err, "error should have been nil")
	assert.NotNil(networkConfig, "network config should be set")
	assert.Equal(networkConfig.Name, "connection-profile")

	channelPeers := networkConfig.Channels["mychannel"].Peers
	assert.Equal(len(channelPeers), 2)
	for _, peer := range channelPeers {
		assert.True(peer.EndorsingPeer)
	}

	peers := networkConfig.Peers
	assert.Equal(len(peers), 2)
	for _, peer := range peers {
		assert.NotEmpty(peer.TLSCACerts.Path)
	}
}

