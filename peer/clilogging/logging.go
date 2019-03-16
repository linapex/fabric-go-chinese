
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456109452693504>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package clilogging

import (
	"fmt"

	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/peer/common"
	"github.com/spf13/cobra"
)

const (
	loggingFuncName = "logging"
	loggingCmdDes   = "Logging configuration: getlevel|setlevel|getlogspec|setlogspec|revertlevels."
)

var logger = flogging.MustGetLogger("cli.logging")

//命令返回用于日志记录的COBRA命令
func Cmd(cf *LoggingCmdFactory) *cobra.Command {
	loggingCmd.AddCommand(getLevelCmd(cf))
	loggingCmd.AddCommand(setLevelCmd(cf))
	loggingCmd.AddCommand(revertLevelsCmd(cf))
	loggingCmd.AddCommand(getLogSpecCmd(cf))
	loggingCmd.AddCommand(setLogSpecCmd(cf))

	return loggingCmd
}

var loggingCmd = &cobra.Command{
	Use:              loggingFuncName,
	Short:            fmt.Sprint(loggingCmdDes),
	Long:             fmt.Sprint(loggingCmdDes),
	PersistentPreRun: common.InitCmd,
}

