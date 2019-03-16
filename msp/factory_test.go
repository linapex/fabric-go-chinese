
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:26</date>
//</624456081795452928>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package msp

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInvalidOpts(t *testing.T) {
	i, err := New(nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid msp.NewOpts instance. It must be either *BCCSPNewOpts or *IdemixNewOpts. It was [<nil>]")
	assert.Nil(t, i)

	i, err = New(&BCCSPNewOpts{NewBaseOpts{Version: -1}})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid *BCCSPNewOpts. Version not recognized [-1]")
	assert.Nil(t, i)

	i, err = New(&IdemixNewOpts{NewBaseOpts{Version: -1}})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid *IdemixNewOpts. Version not recognized [-1]")
	assert.Nil(t, i)
}

func TestNew(t *testing.T) {
	i, err := New(&BCCSPNewOpts{NewBaseOpts{Version: MSPv1_0}})
	assert.NoError(t, err)
	assert.NotNil(t, i)
	assert.Equal(t, MSPVersion(MSPv1_0), i.(*bccspmsp).version)
	assert.Equal(t,
		runtime.FuncForPC(reflect.ValueOf(i.(*bccspmsp).internalSetupFunc).Pointer()).Name(),
		runtime.FuncForPC(reflect.ValueOf(i.(*bccspmsp).setupV1).Pointer()).Name(),
	)
	assert.Equal(t,
		runtime.FuncForPC(reflect.ValueOf(i.(*bccspmsp).internalValidateIdentityOusFunc).Pointer()).Name(),
		runtime.FuncForPC(reflect.ValueOf(i.(*bccspmsp).validateIdentityOUsV1).Pointer()).Name(),
	)

	i, err = New(&BCCSPNewOpts{NewBaseOpts{Version: MSPv1_1}})
	assert.NoError(t, err)
	assert.NotNil(t, i)
	assert.Equal(t, MSPVersion(MSPv1_1), i.(*bccspmsp).version)
	assert.Equal(t,
		runtime.FuncForPC(reflect.ValueOf(i.(*bccspmsp).internalSetupFunc).Pointer()).Name(),
		runtime.FuncForPC(reflect.ValueOf(i.(*bccspmsp).setupV11).Pointer()).Name(),
	)
	assert.Equal(t,
		runtime.FuncForPC(reflect.ValueOf(i.(*bccspmsp).internalValidateIdentityOusFunc).Pointer()).Name(),
		runtime.FuncForPC(reflect.ValueOf(i.(*bccspmsp).validateIdentityOUsV11).Pointer()).Name(),
	)

	i, err = New(&IdemixNewOpts{NewBaseOpts{Version: MSPv1_0}})
	assert.Error(t, err)
	assert.Nil(t, i)
	assert.Contains(t, err.Error(), "Invalid *IdemixNewOpts. Version not recognized [0]")

	i, err = New(&IdemixNewOpts{NewBaseOpts{Version: MSPv1_1}})
	assert.NoError(t, err)
	assert.NotNil(t, i)
}

