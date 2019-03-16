
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455962224234496>

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
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

//
//
type Buffer struct {
	buf      *proto.Buffer
	position int
}

//
func NewBuffer(b []byte) *Buffer {
	return &Buffer{proto.NewBuffer(b), 0}
}

//
func (b *Buffer) DecodeVarint() (uint64, error) {
	val, err := b.buf.DecodeVarint()
	if err == nil {
		b.position += proto.SizeVarint(val)
	} else {
		err = errors.Wrap(err, "error decoding varint with proto.Buffer")
	}
	return val, err
}

//
func (b *Buffer) DecodeRawBytes(alloc bool) ([]byte, error) {
	val, err := b.buf.DecodeRawBytes(alloc)
	if err == nil {
		b.position += proto.SizeVarint(uint64(len(val))) + len(val)
	} else {
		err = errors.Wrap(err, "error decoding raw bytes with proto.Buffer")
	}
	return val, err
}

//GetBytesConsumed返回基础[]字节中当前位置的偏移量
func (b *Buffer) GetBytesConsumed() int {
	return b.position
}

