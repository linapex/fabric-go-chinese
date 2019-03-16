
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456110958448640>

/*
版权所有IBM Corp.2016-2017保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package common

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/hyperledger/fabric/core/comm"
	ab "github.com/hyperledger/fabric/protos/orderer"
	"github.com/pkg/errors"
)

//OrdererClient表示用于与订单通信的客户端
//服务
type OrdererClient struct {
	commonClient
}

//NewOrdererClientfromenv从
//全局毒蛇实例
func NewOrdererClientFromEnv() (*OrdererClient, error) {
	address, override, clientConfig, err := configFromEnv("orderer")
	if err != nil {
		return nil, errors.WithMessage(err, "failed to load config for OrdererClient")
	}
	gClient, err := comm.NewGRPCClient(clientConfig)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create OrdererClient from config")
	}
	oClient := &OrdererClient{
		commonClient: commonClient{
			GRPCClient: gClient,
			address:    address,
			sn:         override}}
	return oClient, nil
}

//广播返回原子广播服务的广播客户端
func (oc *OrdererClient) Broadcast() (ab.AtomicBroadcast_BroadcastClient, error) {
	conn, err := oc.commonClient.NewConnection(oc.address, oc.sn)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("orderer client failed to connect to %s", oc.address))
	}
//TODO:返回前检查是否应实际处理错误
	return ab.NewAtomicBroadcastClient(conn).Broadcast(context.TODO())
}

//Deliver返回AtomicBroadcast服务的Deliver客户端
func (oc *OrdererClient) Deliver() (ab.AtomicBroadcast_DeliverClient, error) {
	conn, err := oc.commonClient.NewConnection(oc.address, oc.sn)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("orderer client failed to connect to %s", oc.address))
	}
//TODO:返回前检查是否应实际处理错误
	return ab.NewAtomicBroadcastClient(conn).Deliver(context.TODO())

}

//证书返回TLS客户端证书（如果可用）
func (oc *OrdererClient) Certificate() tls.Certificate {
	return oc.commonClient.Certificate()
}

