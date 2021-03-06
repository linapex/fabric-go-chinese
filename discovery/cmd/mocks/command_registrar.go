
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:18</date>
//</624456047142113280>

//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import common "github.com/hyperledger/fabric/cmd/common"

import kingpin "gopkg.in/alecthomas/kingpin.v2"
import mock "github.com/stretchr/testify/mock"

//commandRegistrar是commandRegistrar类型的自动生成的模拟类型
type CommandRegistrar struct {
	mock.Mock
}

//命令提供具有给定字段的模拟函数：name、help、oncommand
func (_m *CommandRegistrar) Command(name string, help string, onCommand common.CLICommand) *kingpin.CmdClause {
	ret := _m.Called(name, help, onCommand)

	var r0 *kingpin.CmdClause
	if rf, ok := ret.Get(0).(func(string, string, common.CLICommand) *kingpin.CmdClause); ok {
		r0 = rf(name, help, onCommand)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kingpin.CmdClause)
		}
	}

	return r0
}

