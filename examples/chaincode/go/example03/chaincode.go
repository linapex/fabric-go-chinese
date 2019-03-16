
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:21</date>
//</624456060895236096>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


//这个程序是一个错误的链式代码程序，它试图将状态放入查询上下文中-查询应该返回错误
package example03

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//simple chaincode示例simple chaincode实现
type SimpleChaincode struct{}

//init接受一个字符串和int。它们在状态中作为键/值对存储。
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
var A string //实体
var Aval int //资产持有
	var err error
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

//初始化链码
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d\n", Aval)

//将状态写入分类帐-此放置在运行中是合法的
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//invoke是一个no op
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "query" {
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"query\"")
}

func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
var A string //实体
var Aval int //资产持有
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d\n", Aval)

//将状态写入分类帐-此放置在运行中是非法的
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		jsonResp := "{\"Error\":\"Cannot put state within chaincode query\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(nil)
}

