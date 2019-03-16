
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:26</date>
//</624456082827251712>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package msptesttools

import (
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/config/configtest"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/msp/mgmt"
)

//loadtestmspsetup设置本地msp
//和默认链的链MSP
func LoadMSPSetupForTesting() error {
	dir, err := configtest.GetDevMspDir()
	if err != nil {
		return err
	}
	conf, err := msp.GetLocalMspConfig(dir, nil, "SampleOrg")
	if err != nil {
		return err
	}

	err = mgmt.GetLocalMSP().Setup(conf)
	if err != nil {
		return err
	}

	err = mgmt.GetManagerForChain(util.GetTestChainID()).Setup([]msp.MSP{mgmt.GetLocalMSP()})
	if err != nil {
		return err
	}

	return nil
}

//加载开发本地MSP以用于测试。对于生产/运行时上下文无效
func LoadDevMsp() error {
	mspDir, err := configtest.GetDevMspDir()
	if err != nil {
		return err
	}

	return mgmt.LoadLocalMsp(mspDir, nil, "SampleOrg")
}

