
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:04</date>
//</624455989243940864>

/*
 *版权所有Greg Haskins保留所有权利
 *
 *SPDX许可证标识符：Apache-2.0
 *
 **/


package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//simple chaincode示例simple chaincode实现
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

