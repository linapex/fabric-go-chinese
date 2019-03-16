
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:34</date>
//</624456115370856448>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package discovery

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestToRequest(t *testing.T) {
	sr := &SignedRequest{
		Payload: []byte{0},
	}
	r, err := sr.ToRequest()
	assert.Error(t, err)

	req := &Request{}
	b, _ := proto.Marshal(req)
	sr.Payload = b
	r, err = sr.ToRequest()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

type invalidQuery struct {
}

func (*invalidQuery) isQuery_Query() {
}

func TestGetType(t *testing.T) {
	q := &Query{
		Query: &Query_PeerQuery{
			PeerQuery: &PeerMembershipQuery{},
		},
	}
	assert.Equal(t, PeerMembershipQueryType, q.GetType())
	q = &Query{
		Query: &Query_ConfigQuery{
			ConfigQuery: &ConfigQuery{},
		},
	}
	assert.Equal(t, ConfigQueryType, q.GetType())
	q = &Query{
		Query: &Query_CcQuery{
			CcQuery: &ChaincodeQuery{},
		},
	}
	assert.Equal(t, ChaincodeQueryType, q.GetType())

	q = &Query{
		Query: &invalidQuery{},
	}
	assert.Equal(t, InvalidQueryType, q.GetType())
}

