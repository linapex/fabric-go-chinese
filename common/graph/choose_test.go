
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:56</date>
//</624455954137616384>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChoose(t *testing.T) {
	assert.Equal(t, 24, factorial(4))
	assert.Equal(t, 1, factorial(0))
	assert.Equal(t, 1, factorial(1))
	assert.Equal(t, 15504, nChooseK(20, 5))
	for n := 1; n < 20; n++ {
		for k := 1; k < n; k++ {
			g := chooseKoutOfN(n, k)
			assert.Equal(t, nChooseK(n, k), len(g), "n=%d, k=%d", n, k)
		}
	}
}

