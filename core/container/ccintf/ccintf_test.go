
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:08</date>
//</624456004989358080>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package ccintf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	t.Run("Complete", func(t *testing.T) {
		ccid := &CCID{Name: "ccname", Version: "ver"}
		name := ccid.GetName()
		assert.Equal(t, "ccname-ver", name, "unexpected name")
	})

	t.Run("MissingVersion", func(t *testing.T) {
		ccid := &CCID{Name: "ccname"}
		name := ccid.GetName()
		assert.Equal(t, "ccname", name)
	})
}

