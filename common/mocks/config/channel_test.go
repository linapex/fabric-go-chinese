
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455965554511872>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package config

import (
	"testing"

	"github.com/hyperledger/fabric/common/channelconfig"
)

func TestChannelConfigInterface(t *testing.T) {
	_ = channelconfig.Channel(&Channel{})
}

