
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:56</date>
//</624455954301194240>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertex(t *testing.T) {
	v := NewVertex("1", "1")
	assert.Equal(t, "1", v.Data)
	assert.Equal(t, "1", v.Id)
	u := NewVertex("2", "2")
	v.AddNeighbor(u)
	assert.Contains(t, u.Neighbors(), v)
	assert.Contains(t, v.Neighbors(), u)
	assert.Equal(t, u, v.NeighborById("2"))
	assert.Nil(t, v.NeighborById("3"))
}

