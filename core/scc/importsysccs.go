
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:16</date>
//</624456040854851584>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package scc

import (
	"github.com/hyperledger/fabric/core/common/ccprovider"
)

//deployysccs是系统链码的钩子，系统链码在结构中注册。
//注意，chaincode必须像用户chaincode一样部署和启动。
func (p *Provider) DeploySysCCs(chainID string, ccp ccprovider.ChaincodeProvider) {
	for _, sysCC := range p.SysCCs {
		deploySysCC(chainID, ccp, sysCC)
	}
}

//DEDEPLoySyscs用于单元测试中，在
//在同一进程中重新启动它们。这允许系统干净启动。
//在同一过程中
func (p *Provider) DeDeploySysCCs(chainID string, ccp ccprovider.ChaincodeProvider) {
	for _, sysCC := range p.SysCCs {
		deDeploySysCC(chainID, ccp, sysCC)
	}
}

