
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:22</date>
//</624456064309399552>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package api

import (
	"testing"

	"github.com/hyperledger/fabric/gossip/common"
	"github.com/stretchr/testify/assert"
)

func TestPeerIdentitySetByOrg(t *testing.T) {
	p1 := PeerIdentityInfo{
		Organization: OrgIdentityType("ORG1"),
		Identity:     PeerIdentityType("Peer1"),
	}
	p2 := PeerIdentityInfo{
		Organization: OrgIdentityType("ORG2"),
		Identity:     PeerIdentityType("Peer2"),
	}
	is := PeerIdentitySet{
		p1, p2,
	}
	m := is.ByOrg()
	assert.Len(t, m, 2)
	assert.Equal(t, PeerIdentitySet{p1}, m["ORG1"])
	assert.Equal(t, PeerIdentitySet{p2}, m["ORG2"])
}

func TestPeerIdentitySetByID(t *testing.T) {
	p1 := PeerIdentityInfo{
		Organization: OrgIdentityType("ORG1"),
		PKIId:        common.PKIidType("p1"),
	}
	p2 := PeerIdentityInfo{
		Organization: OrgIdentityType("ORG2"),
		PKIId:        common.PKIidType("p2"),
	}
	is := PeerIdentitySet{
		p1, p2,
	}
	assert.Equal(t, map[string]PeerIdentityInfo{
		"p1": p1,
		"p2": p2,
	}, is.ByID())
}

