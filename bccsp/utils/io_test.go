
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:52</date>
//</624455939050704896>

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


package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirExists(t *testing.T) {
	r, err := DirExists("")
	assert.False(t, r)
	assert.NoError(t, err)

	r, err = DirExists(os.TempDir())
	assert.NoError(t, err)
	assert.Equal(t, true, r)

	r, err = DirExists(filepath.Join(os.TempDir(), "7rhf90239vhev90"))
	assert.NoError(t, err)
	assert.Equal(t, false, r)
}

func TestDirMissingOrEmpty(t *testing.T) {
	r, err := DirMissingOrEmpty("")
	assert.NoError(t, err)
	assert.True(t, r)

	r, err = DirMissingOrEmpty(filepath.Join(os.TempDir(), "7rhf90239vhev90"))
	assert.NoError(t, err)
	assert.Equal(t, true, r)
}

func TestDirEmpty(t *testing.T) {
	_, err := DirEmpty("")
	assert.Error(t, err)

	path := filepath.Join(os.TempDir(), "7rhf90239vhev90")
	defer os.Remove(path)
	os.Mkdir(path, os.ModePerm)

	r, err := DirEmpty(path)
	assert.NoError(t, err)
	assert.Equal(t, true, r)

	r, err = DirEmpty(os.TempDir())
	assert.NoError(t, err)
	assert.Equal(t, false, r)

	r, err = DirMissingOrEmpty(os.TempDir())
	assert.NoError(t, err)
	assert.Equal(t, false, r)

	r, err = DirMissingOrEmpty(path)
	assert.NoError(t, err)
	assert.Equal(t, true, r)
}

