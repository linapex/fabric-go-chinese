
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456111474348032>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package common

import (
	"context"

	ccapi "github.com/hyperledger/fabric/peer/chaincode/api"
	pb "github.com/hyperledger/fabric/protos/peer"
	grpc "google.golang.org/grpc"
)

//PeerDeliverClient保存连接客户端所需的信息
//到对等交付服务
type PeerDeliverClient struct {
	Client pb.DeliverClient
}

//传递将客户端连接到传递RPC
func (dc PeerDeliverClient) Deliver(ctx context.Context, opts ...grpc.CallOption) (ccapi.Deliver, error) {
	d, err := dc.Client.Deliver(ctx, opts...)
	return d, err
}

//deliverfiltered将客户端连接到deliverfiltered rpc
func (dc PeerDeliverClient) DeliverFiltered(ctx context.Context, opts ...grpc.CallOption) (ccapi.Deliver, error) {
	df, err := dc.Client.DeliverFiltered(ctx, opts...)
	return df, err
}

