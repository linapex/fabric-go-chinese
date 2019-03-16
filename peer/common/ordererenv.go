
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456111134609408>

/*
版权所有IBM Corp.2016-2017保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package common

import (
	"os"
	"time"

	"github.com/hyperledger/fabric/common/flogging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	OrderingEndpoint           string
	tlsEnabled                 bool
	clientAuth                 bool
	caFile                     string
	keyFile                    string
	certFile                   string
	ordererTLSHostnameOverride string
	connTimeout                time.Duration
)

//setorderenv将特定于医嘱者的设置添加到全局viper环境中
func SetOrdererEnv(cmd *cobra.Command, args []string) {
//读取旧的日志记录级别设置，如果设置，
//通知用户fabric_logging_spec env变量
	var loggingLevel string
	if viper.GetString("logging_level") != "" {
		loggingLevel = viper.GetString("logging_level")
	} else {
		loggingLevel = viper.GetString("logging.level")
	}
	if loggingLevel != "" {
		mainLogger.Warning("CORE_LOGGING_LEVEL is no longer supported, please use the FABRIC_LOGGING_SPEC environment variable")
	}
//由于COBRA当前不支持，需要在此处初始化日志记录
//链接PersistentPrerun函数
	loggingSpec := os.Getenv("FABRIC_LOGGING_SPEC")
	flogging.InitFromSpec(loggingSpec)
//从标志设置排序器环境
	viper.Set("orderer.tls.rootcert.file", caFile)
	viper.Set("orderer.tls.clientKey.file", keyFile)
	viper.Set("orderer.tls.clientCert.file", certFile)
	viper.Set("orderer.address", OrderingEndpoint)
	viper.Set("orderer.tls.serverhostoverride", ordererTLSHostnameOverride)
	viper.Set("orderer.tls.enabled", tlsEnabled)
	viper.Set("orderer.tls.clientAuthRequired", clientAuth)
	viper.Set("orderer.client.connTimeout", connTimeout)
}

//addorderflags为与排序器相关的命令添加标志
func AddOrdererFlags(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()

	flags.StringVarP(&OrderingEndpoint, "orderer", "o", "", "Ordering service endpoint")
	flags.BoolVarP(&tlsEnabled, "tls", "", false, "Use TLS when communicating with the orderer endpoint")
	flags.BoolVarP(&clientAuth, "clientauth", "", false,
		"Use mutual TLS when communicating with the orderer endpoint")
	flags.StringVarP(&caFile, "cafile", "", "",
		"Path to file containing PEM-encoded trusted certificate(s) for the ordering endpoint")
	flags.StringVarP(&keyFile, "keyfile", "", "",
		"Path to file containing PEM-encoded private key to use for mutual TLS "+
			"communication with the orderer endpoint")
	flags.StringVarP(&certFile, "certfile", "", "",
		"Path to file containing PEM-encoded X509 public key to use for "+
			"mutual TLS communication with the orderer endpoint")
	flags.StringVarP(&ordererTLSHostnameOverride, "ordererTLSHostnameOverride",
		"", "", "The hostname override to use when validating the TLS connection to the orderer.")
	flags.DurationVarP(&connTimeout, "connTimeout",
		"", 3*time.Second, "Timeout for client to connect")
}

