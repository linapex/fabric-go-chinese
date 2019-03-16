
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:26</date>
//</624456082583982080>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package mgmt

import (
	"testing"

	"github.com/hyperledger/fabric/core/config/configtest"
)

func TestLocalMSP(t *testing.T) {
	mspDir, err := configtest.GetDevMspDir()
	if err != nil {
		t.Fatalf("GetDevMspDir failed, err %s", err)
	}

	err = LoadLocalMsp(mspDir, nil, "SampleOrg")
	if err != nil {
		t.Fatalf("LoadLocalMsp failed, err %s", err)
	}

	_, err = GetLocalMSP().GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetDefaultSigningIdentity failed, err %s", err)
	}
}

