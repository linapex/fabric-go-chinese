
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456012237115392>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package main

import (
	"github.com/hyperledger/fabric/core/handlers/endorsement/api"
	"github.com/hyperledger/fabric/protos/peer"
)

type NoOpEndorser struct {
}

func (*NoOpEndorser) Endorse(payload []byte, sp *peer.SignedProposal) (*peer.Endorsement, []byte, error) {
	return nil, payload, nil
}

func (*NoOpEndorser) Init(dependencies ...endorsement.Dependency) error {
	return nil
}

type NoOpEndorserFactory struct {
}

func (*NoOpEndorserFactory) New() endorsement.Plugin {
	return &NoOpEndorser{}
}

//NewPluginFactory是插件基础结构运行的函数，用于创建认可插件工厂。
func NewPluginFactory() endorsement.PluginFactory {
	return &NoOpEndorserFactory{}
}

