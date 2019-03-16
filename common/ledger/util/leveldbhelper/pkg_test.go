
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455962157125632>

/*
版权所有IBM Corp.2017保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/


package leveldbhelper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDBPath = "/tmp/fabric/ledgertests/util/leveldbhelper"

type testDBEnv struct {
	t    *testing.T
	path string
	db   *DB
}

type testDBProviderEnv struct {
	t        *testing.T
	path     string
	provider *Provider
}

func newTestDBEnv(t *testing.T, path string) *testDBEnv {
	testDBEnv := &testDBEnv{t: t, path: path}
	testDBEnv.cleanup()
	testDBEnv.db = CreateDB(&Conf{path})
	return testDBEnv
}

func newTestProviderEnv(t *testing.T, path string) *testDBProviderEnv {
	testProviderEnv := &testDBProviderEnv{t: t, path: path}
	testProviderEnv.cleanup()
	testProviderEnv.provider = NewProvider(&Conf{path})
	return testProviderEnv
}

func (dbEnv *testDBEnv) cleanup() {
	if dbEnv.db != nil {
		dbEnv.db.Close()
	}
	assert.NoError(dbEnv.t, os.RemoveAll(dbEnv.path))
}

func (providerEnv *testDBProviderEnv) cleanup() {
	if providerEnv.provider != nil {
		providerEnv.provider.Close()
	}
	assert.NoError(providerEnv.t, os.RemoveAll(providerEnv.path))
}

