
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456109637242880>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package clilogging

import (
	"context"

	"github.com/hyperledger/fabric/protos/peer"
	"github.com/spf13/cobra"
)

func revertLevelsCmd(cf *LoggingCmdFactory) *cobra.Command {
	var loggingRevertLevelsCmd = &cobra.Command{
		Use:   "revertlevels",
		Short: "Reverts the logging spec to the peer's spec at startup.",
		Long:  `Reverts the logging spec to the peer's spec at startup.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return revertLevels(cf, cmd, args)
		},
	}
	return loggingRevertLevelsCmd
}

func revertLevels(cf *LoggingCmdFactory, cmd *cobra.Command, args []string) (err error) {
	err = checkLoggingCmdParams(cmd, args)
	if err == nil {
//对命令行的分析已完成，因此沉默命令用法
		cmd.SilenceUsage = true

		if cf == nil {
			cf, err = InitCmdFactory()
			if err != nil {
				return err
			}
		}
		env := cf.wrapWithEnvelope(&peer.AdminOperation{})
		_, err = cf.AdminClient.RevertLogLevels(context.Background(), env)
		if err != nil {
			return err
		}
		logger.Info("Logging spec reverted to the peer's spec at startup.")
	}
	return err
}

