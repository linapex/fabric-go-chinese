
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:19</date>
//</624456050396893184>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package support

import "github.com/hyperledger/fabric/discovery"

//DiscoverySupport聚合了发现服务所需的所有支持
type DiscoverySupport struct {
	discovery.AccessControlSupport
	discovery.GossipSupport
	discovery.EndorsementSupport
	discovery.ConfigSupport
	discovery.ConfigSequenceSupport
}

//NewDiscoverySupport返回聚合发现支持
func NewDiscoverySupport(
	access discovery.AccessControlSupport,
	gossip discovery.GossipSupport,
	endorsement discovery.EndorsementSupport,
	config discovery.ConfigSupport,
	sequence discovery.ConfigSequenceSupport,
) *DiscoverySupport {
	return &DiscoverySupport{
		AccessControlSupport:  access,
		GossipSupport:         gossip,
		EndorsementSupport:    endorsement,
		ConfigSupport:         config,
		ConfigSequenceSupport: sequence,
	}
}

