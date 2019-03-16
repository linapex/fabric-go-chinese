
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456077508874240>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package commands

type OutputBlock struct {
	ChannelID   string
	Profile     string
	ConfigPath  string
	OutputBlock string
}

func (o OutputBlock) SessionName() string {
	return "configtxgen-output-block"
}

func (o OutputBlock) Args() []string {
	return []string{
		"-channelID", o.ChannelID,
		"-profile", o.Profile,
		"-configPath", o.ConfigPath,
		"-outputBlock", o.OutputBlock,
	}
}

type CreateChannelTx struct {
	ChannelID             string
	Profile               string
	ConfigPath            string
	OutputCreateChannelTx string
}

func (c CreateChannelTx) SessionName() string {
	return "configtxgen-create-channel-tx"
}

func (c CreateChannelTx) Args() []string {
	return []string{
		"-channelID", c.ChannelID,
		"-profile", c.Profile,
		"-configPath", c.ConfigPath,
		"-outputCreateChannelTx", c.OutputCreateChannelTx,
	}
}

type OutputAnchorPeersUpdate struct {
	ChannelID               string
	Profile                 string
	ConfigPath              string
	AsOrg                   string
	OutputAnchorPeersUpdate string
}

func (o OutputAnchorPeersUpdate) SessionName() string {
	return "configtxgen-output-anchor-peers-update"
}

func (o OutputAnchorPeersUpdate) Args() []string {
	return []string{
		"-channelID", o.ChannelID,
		"-profile", o.Profile,
		"-configPath", o.ConfigPath,
		"-asOrg", o.AsOrg,
		"-outputAnchorPeersUpdate", o.OutputAnchorPeersUpdate,
	}
}

