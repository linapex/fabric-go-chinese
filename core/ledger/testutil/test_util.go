
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456034907328512>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package testutil

import (
	"flag"
	"fmt"
	mathRand "math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/config/configtest"
	"github.com/spf13/viper"
)

//测试随机数生成器用于测试的随机数生成器
type TestRandomNumberGenerator struct {
	rand      *mathRand.Rand
	maxNumber int
}

//newTestRandomNumberGenerator构造新的“TestRandomNumberGenerator”
func NewTestRandomNumberGenerator(maxNumber int) *TestRandomNumberGenerator {
	return &TestRandomNumberGenerator{
		mathRand.New(mathRand.NewSource(time.Now().UnixNano())),
		maxNumber,
	}
}

//下一个生成下一个随机数
func (randNumGenerator *TestRandomNumberGenerator) Next() int {
	return randNumGenerator.rand.Intn(randNumGenerator.maxNumber)
}

//SetupTestConfig sets up configurations for tetsing
func SetupTestConfig() {
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetDefault("peer.ledger.test.loadYAML", true)
	loadYAML := viper.GetBool("peer.ledger.test.loadYAML")
	if loadYAML {
		viper.SetConfigName("test")
		err := viper.ReadInConfig()
if err != nil { //处理读取配置文件时的错误
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

//设置CoreYamlConfig设置测试配置
func SetupCoreYAMLConfig() {
	viper.SetConfigName("core")
	viper.SetEnvPrefix("CORE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := configtest.AddDevConfigPath(nil)
	if err != nil {
		panic(fmt.Errorf("Fatal error adding dev dir: %s \n", err))
	}

	err = viper.ReadInConfig()
if err != nil { //处理读取配置文件时的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

//resetconfigtodefaultValues将配置选项重置回默认值
func ResetConfigToDefaultValues() {
//重置为默认值
	viper.Set("ledger.state.totalQueryLimit", 10000)
	viper.Set("ledger.state.couchDBConfig.internalQueryLimit", 1000)
	viper.Set("ledger.state.stateDatabase", "goleveldb")
	viper.Set("ledger.history.enableHistoryDatabase", false)
	viper.Set("ledger.state.couchDBConfig.autoWarmIndexes", true)
	viper.Set("ledger.state.couchDBConfig.warmIndexesAfterNBlocks", 1)
	viper.Set("peer.fileSystemPath", "/var/hyperledger/production")
}

//parseTestParams分析测试参数
func ParseTestParams() []string {
	testParams := flag.String("testParams", "", "Test specific parameters")
	flag.Parse()
	regex, err := regexp.Compile(",(\\s+)?")
	if err != nil {
		panic(fmt.Errorf("err = %s\n", err))
	}
	paramsArray := regex.Split(*testParams, -1)
	return paramsArray
}

