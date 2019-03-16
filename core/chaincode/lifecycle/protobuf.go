
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:02</date>
//</624455982293979136>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package lifecycle

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

//protobuf定义protobuf生命周期需求的子集，并允许
//用于注入模拟封送错误。
type Protobuf interface {
	Marshal(msg proto.Message) (marshaled []byte, err error)
	Unmarshal(marshaled []byte, msg proto.Message) error
}

//Protobufimpl是用于Protobuf的标准实现
type ProtobufImpl struct{}

//元帅传给原定元帅
func (p ProtobufImpl) Marshal(msg proto.Message) ([]byte, error) {
	res, err := proto.Marshal(msg)
	return res, errors.WithStack(err)
}

//解组传递给Proto。解组
func (p ProtobufImpl) Unmarshal(marshaled []byte, msg proto.Message) error {
	return errors.WithStack(proto.Unmarshal(marshaled, msg))
}

