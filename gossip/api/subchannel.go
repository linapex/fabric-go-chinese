
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:22</date>
//</624456064498143232>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package api

import "github.com/hyperledger/fabric/gossip/common"

//routingfilter定义哪些对等方应接收特定消息，
//或者哪些对等方有资格接收特定消息
type RoutingFilter func(peerIdentity PeerIdentityType) bool

//子通道选择标准描述了从子通道选择对等点的方法
//给了他们签名
type SubChannelSelectionCriteria func(signature PeerSignature) bool

//RoutingFilterFactory定义了一个对象，该对象给定了一个CollectionCriteria和一个通道，
//它可以确定哪些对等方应该知道与
//收藏标准。
type RoutingFilterFactory interface {
//对等方返回给定chainID和collectioncriteria的路由筛选器
	Peers(common.ChainID, SubChannelSelectionCriteria) RoutingFilter
}

