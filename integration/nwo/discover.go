
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456078440009728>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package nwo

import (
	"encoding/json"

	"github.com/hyperledger/fabric/integration/nwo/commands"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

//DiscoveredPeer定义了一个结构，用于使用发现服务发现对等点。
//结果中的每个对等方都将具有这些字段
type DiscoveredPeer struct {
	MSPID      string   `yaml:"mspid,omitempty"`
	Endpoint   string   `yaml:"endpoint,omitempty"`
	Identity   string   `yaml:"identity,omitempty"`
	Chaincodes []string `yaml:"chaincodes,omitempty"`
}

//按照中的指定，使用通道名和用户对对等机运行发现服务命令发现对等机
//函数参数。返回发现的对等点的切片
func DiscoverPeers(n *Network, p *Peer, user, channelName string) func() []DiscoveredPeer {
	return func() []DiscoveredPeer {
		peers := commands.Peers{
			UserCert: n.PeerUserCert(p, user),
			UserKey:  n.PeerUserKey(p, user),
			MSPID:    n.Organization(p.Organization).MSPID,
			Server:   n.PeerAddress(p, ListenPort),
			Channel:  channelName,
		}
		sess, err := n.Discover(peers)
		Expect(err).NotTo(HaveOccurred())
		Eventually(sess).Should(gexec.Exit(0))

		var discovered []DiscoveredPeer
		err = json.Unmarshal(sess.Out.Contents(), &discovered)
		Expect(err).NotTo(HaveOccurred())
		return discovered
	}
}

