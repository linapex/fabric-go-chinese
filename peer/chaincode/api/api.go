
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:31</date>
//</624456104109150208>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package api

import (
	pcommon "github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//go：生成伪造者-o../mock/deliver.go-伪造名称deliver。递送

//Deliver定义用于传递块的接口
type Deliver interface {
	Send(*pcommon.Envelope) error
	Recv() (*pb.DeliverResponse, error)
	CloseSend() error
}

