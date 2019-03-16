
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:08</date>
//</624456004859334656>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package config

import (
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfig_dirExists(t *testing.T) {
	tmpF := os.TempDir()
	exists := dirExists(tmpF)
	assert.True(t, exists,
		"%s directory exists but dirExists returned false", tmpF)

	tmpF = "/blah-" + time.Now().Format(time.RFC3339Nano)
	exists = dirExists(tmpF)
	assert.False(t, exists,
		"%s directory does not exist but dirExists returned true",
		tmpF)
}

func TestConfig_InitViper(t *testing.T) {
//案例1：使用viper实例调用initviper
	v := viper.New()
	err := InitViper(v, "")
	assert.NoError(t, err, "Error returned by InitViper")

//案例2：调用initviper的默认viper实例
	err = InitViper(nil, "")
	assert.NoError(t, err, "Error returned by InitViper")
}

func TestConfig_GetPath(t *testing.T) {
//案例1：毒蛇财产不存在
	path := GetPath("foo")
	assert.Equal(t, "", path, "GetPath should have returned empty string for path 'foo'")

//Case 2: viper property that has absolute path
	viper.Set("testpath", "/test/config.yml")
	path = GetPath("testpath")
	assert.Equal(t, "/test/config.yml", path)
}

func TestConfig_TranslatePathInPlace(t *testing.T) {
//案例1：相对路径
	p := "foo"
	TranslatePathInPlace(OfficialPath, &p)
	assert.NotEqual(t, "foo", p, "TranslatePathInPlace failed to translate path %s", p)

//案例2：绝对路径
	p = "/foo"
	TranslatePathInPlace(OfficialPath, &p)
	assert.Equal(t, "/foo", p, "TranslatePathInPlace failed to translate path %s", p)
}

