
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:31</date>
//</624456104629243904>

/*
版权所有IBM Corp.2016-2017保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package chaincode

import (
	"testing"

	"github.com/hyperledger/fabric/peer/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestOrdererFlags(t *testing.T) {
	var (
		ca       = "root.crt"
		key      = "client.key"
		cert     = "client.crt"
		endpoint = "orderer.example.com:7050"
		sn       = "override.example.com"
	)

	testCmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			t.Logf("rootcert: %s", viper.GetString("orderer.tls.rootcert.file"))
			assert.Equal(t, ca, viper.GetString("orderer.tls.rootcert.file"))
			assert.Equal(t, key, viper.GetString("orderer.tls.clientKey.file"))
			assert.Equal(t, cert, viper.GetString("orderer.tls.clientCert.file"))
			assert.Equal(t, endpoint, viper.GetString("orderer.address"))
			assert.Equal(t, sn, viper.GetString("orderer.tls.serverhostoverride"))
			assert.Equal(t, true, viper.GetBool("orderer.tls.enabled"))
			assert.Equal(t, true, viper.GetBool("orderer.tls.clientAuthRequired"))
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			common.SetOrdererEnv(cmd, args)
		},
	}

	runCmd := Cmd(nil)

	runCmd.AddCommand(testCmd)

	runCmd.SetArgs([]string{"test", "--cafile", ca, "--keyfile", key,
		"--certfile", cert, "--orderer", endpoint, "--tls", "--clientauth",
		"--ordererTLSHostnameOverride", sn})
	err := runCmd.Execute()
	assert.NoError(t, err)

//再次检查env
	t.Logf("address: %s", viper.GetString("orderer.address"))
	assert.Equal(t, ca, viper.GetString("orderer.tls.rootcert.file"))
	assert.Equal(t, key, viper.GetString("orderer.tls.clientKey.file"))
	assert.Equal(t, cert, viper.GetString("orderer.tls.clientCert.file"))
	assert.Equal(t, endpoint, viper.GetString("orderer.address"))
	assert.Equal(t, sn, viper.GetString("orderer.tls.serverhostoverride"))
	assert.Equal(t, true, viper.GetBool("orderer.tls.enabled"))
	assert.Equal(t, true, viper.GetBool("orderer.tls.clientAuthRequired"))
}

