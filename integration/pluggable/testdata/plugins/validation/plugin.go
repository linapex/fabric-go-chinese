
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456079538917376>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package main

import (
	"github.com/hyperledger/fabric/core/handlers/validation/api"
	"github.com/hyperledger/fabric/core/handlers/validation/builtin"
	"github.com/hyperledger/fabric/integration/pluggable"
)

//go build-buildmode=plugin-o插件。

//NewPluginFactory是插件基础结构运行的用于创建验证插件工厂的函数。
func NewPluginFactory() validation.PluginFactory {
	pluggable.PublishValidationPluginActivation()
	return &builtin.DefaultValidationFactory{}
}

