
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:34</date>
//</624456114599104512>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	var h *HashingAlgorithm
	var b *BlockDataHashingStructure
	var o *OrdererAddresses
	var c *Consortium

	h = nil
	assert.Equal(t, "", h.GetName())
	h = &HashingAlgorithm{Name: "SHA256"}
	assert.Equal(t, "SHA256", h.GetName())
	h.Reset()
	_ = h.String()
	_, _ = h.Descriptor()
	h.ProtoMessage()
	assert.Equal(t, "", h.GetName())

	b = nil
	assert.Equal(t, uint32(0), b.GetWidth())
	b = &BlockDataHashingStructure{Width: uint32(1)}
	assert.Equal(t, uint32(1), b.GetWidth())
	b.Reset()
	_ = b.String()
	_, _ = b.Descriptor()
	b.ProtoMessage()
	assert.Equal(t, uint32(0), b.GetWidth())

	o = nil
	assert.Nil(t, o.GetAddresses())
	o = &OrdererAddresses{Addresses: []string{"address"}}
	assert.Equal(t, "address", o.GetAddresses()[0])
	o.Reset()
	_ = o.String()
	_, _ = o.Descriptor()
	o.ProtoMessage()
	assert.Nil(t, o.GetAddresses())

	c = nil
	assert.Equal(t, "", c.GetName())
	c = &Consortium{Name: "consortium"}
	assert.Equal(t, "consortium", c.GetName())
	c.Reset()
	_ = c.String()
	_, _ = c.Descriptor()
	c.ProtoMessage()
	assert.Equal(t, "", c.GetName())

}

