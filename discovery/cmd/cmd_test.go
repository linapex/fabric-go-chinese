
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:18</date>
//</624456046412304384>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package discovery_test

import (
	"testing"

	"github.com/hyperledger/fabric/discovery/cmd"
	"github.com/hyperledger/fabric/discovery/cmd/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/alecthomas/kingpin.v2"
)

func TestAddCommands(t *testing.T) {
	app := kingpin.New("foo", "bar")
	cli := &mocks.CommandRegistrar{}
	configFunc := mock.AnythingOfType("common.CLICommand")
	cli.On("Command", discovery.PeersCommand, mock.Anything, configFunc).Return(app.Command(discovery.PeersCommand, ""))
	cli.On("Command", discovery.ConfigCommand, mock.Anything, configFunc).Return(app.Command(discovery.ConfigCommand, ""))
	cli.On("Command", discovery.EndorsersCommand, mock.Anything, configFunc).Return(app.Command(discovery.EndorsersCommand, ""))
	discovery.AddCommands(cli)
//确保为子命令配置了服务和通道标志
	for _, cmd := range []string{discovery.PeersCommand, discovery.ConfigCommand, discovery.EndorsersCommand} {
		assert.NotNil(t, app.GetCommand(cmd).GetFlag("server"))
		assert.NotNil(t, app.GetCommand(cmd).GetFlag("channel"))
	}
//确保为背书人调用了chaincode和collection标志。
	assert.NotNil(t, app.GetCommand(discovery.EndorsersCommand).GetFlag("chaincode"))
	assert.NotNil(t, app.GetCommand(discovery.EndorsersCommand).GetFlag("collection"))
}

