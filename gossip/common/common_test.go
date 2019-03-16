
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:22</date>
//</624456065936789504>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNotSame(t *testing.T) {
	id := PKIidType("1")
	assert.True(t, id.IsNotSameFilter(PKIidType("2")))
	assert.False(t, id.IsNotSameFilter(PKIidType("1")))
	assert.False(t, id.IsNotSameFilter(id))
}

func TestPKIidTypeStringer(t *testing.T) {
	tests := []struct {
		input    PKIidType
		expected string
	}{
		{nil, "<nil>"},
		{PKIidType{}, ""},
		{PKIidType{0, 1, 2, 3}, "00010203"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.input.String())
	}
}

