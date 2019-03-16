
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:34</date>
//</624456114884317184>

/*
版权所有IBM Corp.2017保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

                 http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/


package common

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func (p *Policy) VariablyOpaqueFields() []string {
	return []string{"value"}
}

func (p *Policy) VariablyOpaqueFieldProto(name string) (proto.Message, error) {
	if name != p.VariablyOpaqueFields()[0] {
		return nil, fmt.Errorf("not a marshaled field: %s", name)
	}
	switch p.Type {
	case int32(Policy_SIGNATURE):
		return &SignaturePolicyEnvelope{}, nil
	case int32(Policy_IMPLICIT_META):
		return &ImplicitMetaPolicy{}, nil
	default:
		return nil, fmt.Errorf("unable to decode policy type: %v", p.Type)
	}
}

