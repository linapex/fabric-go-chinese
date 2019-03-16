
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456113152069632>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package version

import (
	"fmt"
	"runtime"

	"github.com/hyperledger/fabric/common/metadata"
	"github.com/spf13/cobra"
)

//程序名
const ProgramName = "peer"

//cmd返回版本的cobra命令
func Cmd() *cobra.Command {
	return cobraCommand
}

var cobraCommand = &cobra.Command{
	Use:   "version",
	Short: "Print fabric peer version.",
	Long:  `Print current version of the fabric peer server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}
//对命令行的分析已完成，因此沉默命令用法
		cmd.SilenceUsage = true
		fmt.Print(GetInfo())
		return nil
	},
}

//GetInfo返回对等机的版本信息
func GetInfo() string {
	if metadata.Version == "" {
		metadata.Version = "development build"
	}

	ccinfo := fmt.Sprintf(" Base Image Version: %s\n"+
		"  Base Docker Namespace: %s\n"+
		"  Base Docker Label: %s\n"+
		"  Docker Namespace: %s\n",
		metadata.BaseVersion, metadata.BaseDockerNamespace,
		metadata.BaseDockerLabel, metadata.DockerNamespace)

	return fmt.Sprintf("%s:\n Version: %s\n Commit SHA: %s\n Go version: %s\n"+
		" OS/Arch: %s\n"+
		" Chaincode:\n %s\n",
		ProgramName, metadata.Version, metadata.CommitSHA, runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH), ccinfo)
}

