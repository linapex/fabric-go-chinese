
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:26</date>
//</624456082906943488>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package msptesttools

import (
	"testing"

	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/msp/mgmt"
)

func TestFakeSetup(t *testing.T) {
	err := LoadMSPSetupForTesting()
	if err != nil {
		t.Fatalf("LoadLocalMsp failed, err %s", err)
	}

	_, err = mgmt.GetLocalMSP().GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetDefaultSigningIdentity failed, err %s", err)
	}

	msps, err := mgmt.GetManagerForChain(util.GetTestChainID()).GetMSPs()
	if err != nil {
		t.Fatalf("EnlistedMSPs failed, err %s", err)
	}

	if msps == nil || len(msps) == 0 {
		t.Fatalf("There are no MSPS in the manager for chain %s", util.GetTestChainID())
	}
}

