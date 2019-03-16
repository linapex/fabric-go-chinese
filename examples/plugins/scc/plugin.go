
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:22</date>
//</624456064183570432>

/*
版权所有SecureKey Technologies Inc.保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//new返回chaincode接口的实现
func New() shim.Chaincode {
	return &scc{}
}

type scc struct{}

//init实现chaincode填充程序接口
func (s *scc) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

//invoke实现chaincode填充程序接口
func (s *scc) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func main() {}

