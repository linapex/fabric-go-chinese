
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456079446642688>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package main

import (
	"github.com/hyperledger/fabric/core/handlers/endorsement/api"
	"github.com/hyperledger/fabric/core/handlers/endorsement/builtin"
	"github.com/hyperledger/fabric/integration/pluggable"
)

//go build-buildmode=plugin-o插件。

//NewPluginFactory是插件基础结构运行的函数，用于创建认可插件工厂。
func NewPluginFactory() endorsement.PluginFactory {
	pluggable.PublishEndorsementPluginActivation()
	return &builtin.DefaultEndorsementFactory{}
}

