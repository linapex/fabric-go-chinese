
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:26</date>
//</624456083934547968>

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


package msp

import (
	"path/filepath"
	"testing"

	"github.com/hyperledger/fabric/bccsp/sw"
	"github.com/stretchr/testify/assert"
)

func TestBadConfigOU(t *testing.T) {
//测试数据/badconfiguo：
//配置是这样的，只有标识
//当ou=cop2并且由根CA签名时，应该验证
	thisMSP := getLocalMSP(t, "testdata/badconfigou")

	id, err := thisMSP.GetDefaultSigningIdentity()
	assert.NoError(t, err)

//默认签名标识ou是cop，但配置了msp
//仅验证ou为cop2的标识
	err = id.Validate()
	assert.Error(t, err)
}

func TestBadConfigOUCert(t *testing.T) {
//测试数据/badconfigoucert:
//ou标识符的配置指向
//既不是CA也不是MSP的中间CA的证书。
	conf, err := GetLocalMspConfig("testdata/badconfigoucert", nil, "SampleOrg")
	assert.NoError(t, err)

	thisMSP, err := newBccspMsp(MSPv1_0)
	assert.NoError(t, err)

	err = thisMSP.Setup(conf)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed adding OU. Certificate [")
	assert.Contains(t, err.Error(), "] not in root or intermediate certs.")
}

func TestValidateIntermediateConfigOU(t *testing.T) {
//测试数据/外部：
//配置是这样的，只有
//ou=应验证由中间CA签署的超级账本测试
	thisMSP := getLocalMSP(t, "testdata/external")

	id, err := thisMSP.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	err = id.Validate()
	assert.NoError(t, err)

	conf, err := GetLocalMspConfig("testdata/external", nil, "SampleOrg")
	assert.NoError(t, err)

	thisMSP, err = newBccspMsp(MSPv1_0)
	assert.NoError(t, err)
	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join("testdata/external", "keystore"), true)
	assert.NoError(t, err)
	csp, err := sw.NewWithParams(256, "SHA2", ks)
	assert.NoError(t, err)
	thisMSP.(*bccspmsp).bccsp = csp

	err = thisMSP.Setup(conf)
	assert.NoError(t, err)
}

