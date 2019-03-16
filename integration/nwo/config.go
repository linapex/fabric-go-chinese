
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456077953470464>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package nwo

//config保存生成所需的基本信息
//结构配置文件。
type Config struct {
	Organizations []*Organization `yaml:"organizations,omitempty"`
	Consortiums   []*Consortium   `yaml:"consortiums,omitempty"`
	SystemChannel *SystemChannel  `yaml:"system_channel,omitempty"`
	Channels      []*Channel      `yaml:"channels,omitempty"`
	Consensus     *Consensus      `yaml:"consensus,omitempty"`
	Orderers      []*Orderer      `yaml:"orderers,omitempty"`
	Peers         []*Peer         `yaml:"peers,omitempty"`
	Profiles      []*Profile      `yaml:"profiles,omitempty"`
	Templates     *Templates      `yaml:"templates,omitempty"`
}

func (c *Config) RemovePeer(orgName, peerName string) {
	peers := []*Peer{}
	for _, p := range c.Peers {
		if p.Organization != orgName || p.Name != peerName {
			peers = append(peers, p)
		}
	}
	c.Peers = peers
}

