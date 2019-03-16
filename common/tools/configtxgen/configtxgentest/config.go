
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:59</date>
//</624455968264032256>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package configtxgentest

import (
	"fmt"

	"github.com/hyperledger/fabric/common/tools/configtxgen/localconfig"
	"github.com/hyperledger/fabric/core/config/configtest"
)

func Load(profile string) *localconfig.Profile {
	devConfigDir, err := configtest.GetDevConfigDir()
	if err != nil {
		panic(fmt.Sprintf("failed to get dev config dir: %s", err))
	}
	return localconfig.Load(profile, devConfigDir)
}

func LoadTopLevel() *localconfig.TopLevel {
	devConfigDir, err := configtest.GetDevConfigDir()
	if err != nil {
		panic(fmt.Sprintf("failed to get dev config dir: %s", err))
	}
	return localconfig.LoadTopLevel(devConfigDir)
}

