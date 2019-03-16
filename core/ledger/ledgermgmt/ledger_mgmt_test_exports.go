
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:14</date>
//</624456030973071360>

/*
版权所有IBM公司。保留所有权利。
SPDX许可证标识符：Apache-2.0
**/


package ledgermgmt

import (
	"fmt"
	"os"

	"github.com/hyperledger/fabric/common/metrics/disabled"
	"github.com/hyperledger/fabric/core/chaincode/platforms"
	"github.com/hyperledger/fabric/core/chaincode/platforms/golang"
	"github.com/hyperledger/fabric/core/ledger/ledgerconfig"
	"github.com/hyperledger/fabric/core/ledger/mock"
)

//initializetestenv初始化测试的ledgermgmt
func InitializeTestEnv() {
	remove()
	InitializeTestEnvWithInitializer(nil)
}

//initializeTestInvWithInitializer使用提供的初始值设定项初始化测试的ledgermgmt
func InitializeTestEnvWithInitializer(initializer *Initializer) {
	remove()
	InitializeExistingTestEnvWithInitializer(initializer)
}

//InitializeExistingTestenvWithInitializer为具有现有分类帐的测试初始化LedgerMgmt
//此功能不会删除现有分类帐，并在升级测试中使用。
//Todo Ledgermgmt应重新编写，以将包范围的函数移动到结构
func InitializeExistingTestEnvWithInitializer(initializer *Initializer) {
	if initializer == nil {
		initializer = &Initializer{}
	}
	if initializer.DeployedChaincodeInfoProvider == nil {
		initializer.DeployedChaincodeInfoProvider = &mock.DeployedChaincodeInfoProvider{}
	}
	if initializer.MetricsProvider == nil {
		initializer.MetricsProvider = &disabled.Provider{}
	}
	if initializer.PlatformRegistry == nil {
		initializer.PlatformRegistry = platforms.NewRegistry(&golang.Platform{})
	}
	initialize(initializer)
}

//cleanuptestenv关闭ledgermagmt并删除存储目录
func CleanupTestEnv() {
	Close()
	remove()
}

func remove() {
	path := ledgerconfig.GetRootPath()
	fmt.Printf("removing dir = %s\n", path)
	err := os.RemoveAll(path)
	if err != nil {
		logger.Errorf("Error: %s", err)
	}
}

