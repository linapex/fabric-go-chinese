
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:31</date>
//</624456103513559040>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMockPublicPrivateKeyPairPEM(t *testing.T) {
	_, _, err := GenerateMockPublicPrivateKeyPairPEM(false)
	assert.NoError(t, err, "Unable to generate a public/private key pair: %v", err)
}

func TestGenerateMockPublicPrivateKeyPairPEMWhenCASet(t *testing.T) {
	_, _, err := GenerateMockPublicPrivateKeyPairPEM(true)
	assert.NoError(t, err, "Unable to generate a signer certificate: %v", err)
}

