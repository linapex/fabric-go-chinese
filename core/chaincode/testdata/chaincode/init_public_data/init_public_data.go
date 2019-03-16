
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455993484382208>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//simple chaincode示例simple chaincode实现
type SimpleChaincode struct {
}

//init初始化公共状态
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	if err := stub.PutState("dummyKey", []byte("dummyValue")); err != nil {
		return shim.Error(fmt.Sprintf("put operation failed. Error storing state: %s", err))
	}
	return shim.Success(nil)
}

//invoke是一个no op
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

