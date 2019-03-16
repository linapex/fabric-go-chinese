
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:04</date>
//</624455988983894016>

/*
 *版权所有Greg Haskins保留所有权利
 *
 *SPDX许可证标识符：Apache-2.0
 *
 *本测试代码的目的是证明系统正确包装。
 *向上依赖。因此，我们综合了一个场景，其中链代码
 *直接和间接导入非标准依赖项，然后
 *希望进行单元测试，以验证包中是否包含所需的所有内容。
 *并最终正确构建。
 *
 **/


package main

import (
	"chaincodes/AutoVendor/directdep"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//simple chaincode示例simple chaincode实现
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Error("NOT IMPL")
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Error("NOT IMPL")
}

func main() {
	directdep.PointlessFunction()

	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

