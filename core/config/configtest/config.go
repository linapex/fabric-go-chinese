
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:08</date>
//</624456004708339712>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package configtest

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

//adddevconfigPath将devconfigdir添加到viper路径。
func AddDevConfigPath(v *viper.Viper) error {
	devPath, err := GetDevConfigDir()
	if err != nil {
		return err
	}

	if v != nil {
		v.AddConfigPath(devPath)
	} else {
		viper.AddConfigPath(devPath)
	}

	return nil
}

func dirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

//getdevconfigdir获取默认配置的路径，即
//使用源树维护。这只能用于
//测试/开发环境。
func GetDevConfigDir() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", fmt.Errorf("GOPATH not set")
	}

	for _, p := range filepath.SplitList(gopath) {
		devPath := filepath.Join(p, "src/github.com/hyperledger/fabric/sampleconfig")
		if !dirExists(devPath) {
			continue
		}

		return devPath, nil
	}

	return "", fmt.Errorf("DevConfigDir not found in %s", gopath)
}

//getdevmspdir获取维护的sampleconfig/msp树的路径
//使用源树。这只能在测试/开发中使用
//语境。
func GetDevMspDir() (string, error) {
	devDir, err := GetDevConfigDir()
	if err != nil {
		return "", fmt.Errorf("Error obtaining DevConfigDir: %s", devDir)
	}

	return filepath.Join(devDir, "msp"), nil
}

func SetDevFabricConfigPath(t *testing.T) (cleanup func()) {
	t.Helper()

	oldFabricCfgPath, resetFabricCfgPath := os.LookupEnv("FABRIC_CFG_PATH")
	devConfigDir, err := GetDevConfigDir()
	if err != nil {
		t.Fatalf("failed to get dev config dir: %s", err)
	}

	err = os.Setenv("FABRIC_CFG_PATH", devConfigDir)
	if resetFabricCfgPath {
		return func() {
			err := os.Setenv("FABRIC_CFG_PATH", oldFabricCfgPath)
			assert.NoError(t, err)
		}
	}

	return func() {
		err := os.Unsetenv("FABRIC_CFG_PATH")
		assert.NoError(t, err)
	}
}

