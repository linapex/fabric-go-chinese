
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:49</date>
//</624455927029829632>

//+构建PKCS11

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

package factory

import (
	"os"
	"testing"

	"github.com/hyperledger/fabric/bccsp/pkcs11"
	"github.com/stretchr/testify/assert"
)

func TestPKCS11FactoryName(t *testing.T) {
	f := &PKCS11Factory{}
	assert.Equal(t, f.Name(), PKCS11BasedFactoryName)
}

func TestPKCS11FactoryGetInvalidArgs(t *testing.T) {
	f := &PKCS11Factory{}

	_, err := f.Get(nil)
	assert.Error(t, err, "Invalid config. It must not be nil.")

	_, err = f.Get(&FactoryOpts{})
	assert.Error(t, err, "Invalid config. It must not be nil.")

	opts := &FactoryOpts{
		Pkcs11Opts: &pkcs11.PKCS11Opts{},
	}
	_, err = f.Get(opts)
	assert.Error(t, err, "CSP:500 - Failed initializing configuration at [0,]")
}

func TestPKCS11FactoryGet(t *testing.T) {
	f := &PKCS11Factory{}
	lib, pin, label := pkcs11.FindPKCS11Lib()

	opts := &FactoryOpts{
		Pkcs11Opts: &pkcs11.PKCS11Opts{
			SecLevel:   256,
			HashFamily: "SHA2",
			Library:    lib,
			Pin:        pin,
			Label:      label,
		},
	}
	csp, err := f.Get(opts)
	assert.NoError(t, err)
	assert.NotNil(t, csp)

	opts = &FactoryOpts{
		Pkcs11Opts: &pkcs11.PKCS11Opts{
			SecLevel:     256,
			HashFamily:   "SHA2",
			FileKeystore: &pkcs11.FileKeystoreOpts{KeyStorePath: os.TempDir()},
			Library:      lib,
			Pin:          pin,
			Label:        label,
		},
	}
	csp, err = f.Get(opts)
	assert.NoError(t, err)
	assert.NotNil(t, csp)

	opts = &FactoryOpts{
		Pkcs11Opts: &pkcs11.PKCS11Opts{
			SecLevel:   256,
			HashFamily: "SHA2",
			Ephemeral:  true,
			Library:    lib,
			Pin:        pin,
			Label:      label,
		},
	}
	csp, err = f.Get(opts)
	assert.NoError(t, err)
	assert.NotNil(t, csp)
}

