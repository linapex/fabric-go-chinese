
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:32</date>
//</624456109054234624>

/*
版权所有IBM Corp.2017保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

                 http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/


package channel

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/hyperledger/fabric/peer/common"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/spf13/cobra"
)

func updateCmd(cf *ChannelCmdFactory) *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Send a configtx update.",
		Long:  "Signs and sends the supplied configtx update file to the channel. Requires '-f', '-o', '-c'.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return update(cmd, args, cf)
		},
	}
	flagList := []string{
		"channelID",
		"file",
	}
	attachFlags(updateCmd, flagList)

	return updateCmd
}

func update(cmd *cobra.Command, args []string, cf *ChannelCmdFactory) error {
//由“-c”命令填充的全局chainID
	if channelID == common.UndefinedParamValue {
		return errors.New("Must supply channel ID")
	}

	if channelTxFile == "" {
		return InvalidCreateTx("No configtx file name supplied")
	}
//对命令行的分析已完成，因此沉默命令用法
	cmd.SilenceUsage = true

	var err error
	if cf == nil {
		cf, err = InitCmdFactory(EndorserNotRequired, PeerDeliverNotRequired, OrdererRequired)
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

	var broadcastClient common.BroadcastClient
	broadcastClient, err = cf.BroadcastFactory()
	if err != nil {
		return fmt.Errorf("Error getting broadcast client: %s", err)
	}

	defer broadcastClient.Close()
	err = broadcastClient.Send(sCtxEnv)
	if err != nil {
		return err
	}

	logger.Info("Successfully submitted channel update")
	return nil
}

