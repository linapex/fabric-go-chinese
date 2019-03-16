
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456075948593152>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package main

import (
	"fmt"
	"os"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/integration/chaincode/keylevelep"
)

func main() {
	err := shim.Start(&keylevelep.EndorsementCC{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exiting SBE chaincode: %s", err)
		os.Exit(2)
	}
}

