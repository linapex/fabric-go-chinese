
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:21</date>
//</624456060329005056>

/*
版权所有IBM Corp.2016保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/


package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//事件发送者示例简单的链代码实现
type EventSender struct {
}

//初始化函数
func (t *EventSender) Init(stub shim.ChaincodeStubInterface) pb.Response {
	err := stub.PutState("noevents", []byte("0"))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//调用函数
func (t *EventSender) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	b, err := stub.GetState("noevents")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	noevts, _ := strconv.Atoi(string(b))

	tosend := "Event " + string(b)
	for _, s := range args {
		tosend = tosend + "," + s
	}

	err = stub.PutState("noevents", []byte(strconv.Itoa(noevts+1)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent("evtsender", []byte(tosend))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//查询函数
func (t *EventSender) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	b, err := stub.GetState("noevents")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	jsonResp := "{\"NoEvents\":\"" + string(b) + "\"}"
	return shim.Success([]byte(jsonResp))
}

func (t *EventSender) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		return t.invoke(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"query\"")
}

func main() {
	err := shim.Start(new(EventSender))
	if err != nil {
		fmt.Printf("Error starting EventSender chaincode: %s", err)
	}
}

