
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455994256134144>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package comm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	serverOptions := ServerKeepaliveOptions(nil)
	assert.NotNil(t, serverOptions)

	clientOptions := ClientKeepaliveOptions(nil)
	assert.NotNil(t, clientOptions)
}

