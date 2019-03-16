
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:07</date>
//</624456002606993408>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package sysccprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	chaincodeInstance := ChaincodeInstance{
		ChainID:          "ChainID",
		ChaincodeName:    "ChaincodeName",
		ChaincodeVersion: "ChaincodeVersion",
	}

	assert.NotNil(t, chaincodeInstance.String(), "str should not be nil")
	assert.Equal(t, chaincodeInstance.String(), "ChainID.ChaincodeName#ChaincodeVersion", "str should be the correct value")
}

