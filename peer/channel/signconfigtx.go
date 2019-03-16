
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:32</date>
//</624456108844519424>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channel

import (
	"io/ioutil"

	"github.com/hyperledger/fabric/protos/utils"
	"github.com/spf13/cobra"
)

func signconfigtxCmd(cf *ChannelCmdFactory) *cobra.Command {
	signconfigtxCmd := &cobra.Command{
		Use:   "signconfigtx",
		Short: "Signs a configtx update.",
		Long:  "Signs the supplied configtx update file in place on the filesystem. Requires '-f'.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return sign(cmd, args, cf)
		},
	}
	flagList := []string{
		"file",
	}
	attachFlags(signconfigtxCmd, flagList)

	return signconfigtxCmd
}

func sign(cmd *cobra.Command, args []string, cf *ChannelCmdFactory) error {
	if channelTxFile == "" {
		return InvalidCreateTx("No configtx file name supplied")
	}
//对命令行的分析已完成，因此沉默命令用法
	cmd.SilenceUsage = true

	var err error
	if cf == nil {
		cf, err = InitCmdFactory(EndorserNotRequired, PeerDeliverNotRequired, OrdererNotRequired)
		if err != nil {
			return err
		}
	}

	fileData, err := ioutil.ReadFile(channelTxFile)
	if err != nil {
		return ConfigTxFileNotFound(err.Error())
	}

	ctxEnv, err := utils.UnmarshalEnvelope(fileData)
	if err != nil {
		return err
	}

	sCtxEnv, err := sanityCheckAndSignConfigTx(ctxEnv)
	if err != nil {
		return err
	}

	sCtxEnvData := utils.MarshalOrPanic(sCtxEnv)

	return ioutil.WriteFile(channelTxFile, sCtxEnvData, 0660)
}

