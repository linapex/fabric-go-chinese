
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455949117034496>

/*
版权所有IBM Corp.2017保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomBytes(t *testing.T) {
	_, err := GetRandomBytes(10)

	assert.NoError(t, err, "GetRandomBytes fails")
}

func TestGetRandomNonce(t *testing.T) {
	_, err := GetRandomNonce()

	assert.NoError(t, err, "GetRandomNonce fails")
}

