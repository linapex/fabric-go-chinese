
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456010253209600>

/*
版权所有IBM Corp，SecureKey Technologies Inc.保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package auth

import (
	"github.com/hyperledger/fabric/protos/peer"
)

//筛选器定义一个身份验证筛选器，用于截取
//流程建议方法
type Filter interface {
	peer.EndorserServer
//init用下一个背书服务器初始化筛选器
	Init(next peer.EndorserServer)
}

//chainfilters按提供的顺序链接给定的auth筛选器。
//最后一个过滤器总是转发给背书人。
func ChainFilters(endorser peer.EndorserServer, filters ...Filter) peer.EndorserServer {
	if len(filters) == 0 {
		return endorser
	}

//每个过滤器向前到下一个
	for i := 0; i < len(filters)-1; i++ {
		filters[i].Init(filters[i+1])
	}

//最后一个过滤器转发给背书人
	filters[len(filters)-1].Init(endorser)

	return filters[0]
}

