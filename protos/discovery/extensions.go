
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:34</date>
//</624456115303747584>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package discovery

import (
	"github.com/golang/protobuf/proto"
)

//querytype定义服务发现请求的类型
type QueryType uint8

const (
	InvalidQueryType QueryType = iota
	ConfigQueryType
	PeerMembershipQueryType
	ChaincodeQueryType
	LocalMembershipQueryType
)

//GetType返回请求的类型
func (q *Query) GetType() QueryType {
	if q.GetCcQuery() != nil {
		return ChaincodeQueryType
	}
	if q.GetConfigQuery() != nil {
		return ConfigQueryType
	}
	if q.GetPeerQuery() != nil {
		return PeerMembershipQueryType
	}
	if q.GetLocalPeers() != nil {
		return LocalMembershipQueryType
	}
	return InvalidQueryType
}

//ToRequest反序列化此SignedRequest的负载
//并以其对象形式返回序列化请求。
//如果操作失败，则返回错误。
func (sr *SignedRequest) ToRequest() (*Request, error) {
	req := &Request{}
	return req, proto.Unmarshal(sr.Payload, req)
}

//configat返回响应中给定索引的configresult，
//或错误（如果存在）。
func (m *Response) ConfigAt(i int) (*ConfigResult, *Error) {
	r := m.Results[i]
	return r.GetConfigResult(), r.GetError()
}

//membershipat返回响应中给定索引处的peermembershipresult，
//或错误（如果存在）。
func (m *Response) MembershipAt(i int) (*PeerMembershipResult, *Error) {
	r := m.Results[i]
	return r.GetMembers(), r.GetError()
}

//背书返回响应中给定索引处的对等成员身份结果，
//或错误（如果存在）。
func (m *Response) EndorsersAt(i int) (*ChaincodeQueryResult, *Error) {
	r := m.Results[i]
	return r.GetCcQueryRes(), r.GetError()
}

