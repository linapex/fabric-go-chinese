
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455941856694272>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package main

import (
	"os"

	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/cmd/common"
	"github.com/hyperledger/fabric/discovery/cmd"
)

func main() {
	factory.InitFactories(nil)
	cli := common.NewCLI("discover", "Command line client for fabric discovery service")
	discovery.AddCommands(cli)
	cli.Run(os.Args[1:])
}

