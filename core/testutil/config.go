
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:17</date>
//</624456043249799168>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package testutil

import (
	"flag"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"

	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/core/config/configtest"
	"github.com/hyperledger/fabric/msp"
	"github.com/spf13/viper"
)

var configLogger = flogging.MustGetLogger("config")

//SETUPTESTCONFIG在测试执行期间设置配置
func SetupTestConfig() {
	flag.Parse()

//现在设置配置文件
	viper.SetEnvPrefix("CORE")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
viper.SetConfigName("core") //配置文件名（不带扩展名）
	err := configtest.AddDevConfigPath(nil)
	if err != nil {
		panic(fmt.Errorf("Fatal error adding DevConfigPath: %s \n", err))
	}

err = viper.ReadInConfig() //查找并读取配置文件
if err != nil {            //处理读取配置文件时的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

//设置maxprocs的数目
	var numProcsDesired = viper.GetInt("peer.gomaxprocs")
	configLogger.Debugf("setting Number of procs to %d, was %d\n", numProcsDesired, runtime.GOMAXPROCS(numProcsDesired))

//初始化BCCSP
	var bccspConfig *factory.FactoryOpts
	err = viper.UnmarshalKey("peer.BCCSP", &bccspConfig)
	if err != nil {
		bccspConfig = nil
	}

	tmpKeyStore, err := ioutil.TempDir("/tmp", "msp-keystore")
	if err != nil {
		panic(fmt.Errorf("Could not create temporary directory: %s\n", tmpKeyStore))
	}

	msp.SetupBCCSPKeystoreConfig(bccspConfig, tmpKeyStore)

	err = factory.InitFactories(bccspConfig)
	if err != nil {
		panic(fmt.Errorf("Could not initialize BCCSP Factories [%s]", err))
	}
}

