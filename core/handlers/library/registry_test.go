
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:10</date>
//</624456012711071744>

/*
版权所有IBM Corp，SecureKey Technologies Inc.保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package library

import (
	"testing"

	"github.com/hyperledger/fabric/core/handlers/auth"
	"github.com/hyperledger/fabric/core/handlers/decoration"
	"github.com/stretchr/testify/assert"
)

func TestInitRegistry(t *testing.T) {
	r := InitRegistry(Config{
		AuthFilters: []*HandlerConfig{{Name: "DefaultAuth"}},
		Decorators:  []*HandlerConfig{{Name: "DefaultDecorator"}},
	})
	assert.NotNil(t, r)
	authHandlers := r.Lookup(Auth)
	assert.NotNil(t, authHandlers)
	filters, isAuthFilters := authHandlers.([]auth.Filter)
	assert.True(t, isAuthFilters)
	assert.Len(t, filters, 1)

	decorationHandlers := r.Lookup(Decoration)
	assert.NotNil(t, decorationHandlers)
	decorators, isDecorators := decorationHandlers.([]decoration.Decorator)
	assert.True(t, isDecorators)
	assert.Len(t, decorators, 1)
}

func TestLoadCompiledInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic with invalid factory method")
		}
	}()

	testReg := registry{}
	testReg.loadCompiled("InvalidFactory", Auth)
}

