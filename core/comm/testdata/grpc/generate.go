
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:06</date>
//</624455997309587456>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


//+构建忽略

//go:generate protoc--proto_path=$go path/src/github.com/hyperledger/fabric/core/comm/testdata/grpc--go_out=plugins=grpc:$go path/src test.proto

package grpc

