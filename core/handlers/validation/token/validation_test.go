
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:10</date>
//</624456016494333952>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package token_test

import (
	"testing"

	"github.com/hyperledger/fabric/core/handlers/validation/token"
	"github.com/stretchr/testify/assert"
)

func TestValidationFactory_New(t *testing.T) {
	factory := &token.ValidationFactory{}
	plugin := factory.New()
	assert.NotNil(t, plugin)
}

func TestValidation_Validate(t *testing.T) {
	factory := &token.ValidationFactory{}
	plugin := factory.New()

	err := plugin.Init()
	assert.NoError(t, err)

//验证返回零，无论什么！
	err = plugin.Validate(nil, "", 0, 0, nil)
	assert.NoError(t, err)
}

