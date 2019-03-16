
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455962282954752>

/*
版权所有IBM Corp.2016保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
*/


package util

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestBuffer(t *testing.T) {
	pb := proto.NewBuffer(nil)
	pb.EncodeVarint(10)
	pos1 := len(pb.Bytes())
	pb.EncodeRawBytes([]byte("JunkText"))
	pos2 := len(pb.Bytes())
	pb.EncodeRawBytes([]byte("YetAnotherJunkText"))
	pos3 := len(pb.Bytes())
	pb.EncodeVarint(1000000)
	pos4 := len(pb.Bytes())

	b := NewBuffer(pb.Bytes())
	b.DecodeVarint()
	assert.Equal(t, pos1, b.GetBytesConsumed())
	b.DecodeRawBytes(false)
	assert.Equal(t, pos2, b.GetBytesConsumed())
	b.DecodeRawBytes(false)
	assert.Equal(t, pos3, b.GetBytesConsumed())
	b.DecodeVarint()
	assert.Equal(t, pos4, b.GetBytesConsumed())
}

