
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:35</date>
//</624456118856323072>

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


package msp

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func (mp *MSPPrincipal) VariablyOpaqueFields() []string {
	return []string{"principal"}
}

func (mp *MSPPrincipal) VariablyOpaqueFieldProto(name string) (proto.Message, error) {
	if name != mp.VariablyOpaqueFields()[0] {
		return nil, fmt.Errorf("not a marshaled field: %s", name)
	}
	switch mp.PrincipalClassification {
	case MSPPrincipal_ROLE:
		return &MSPRole{}, nil
	case MSPPrincipal_ORGANIZATION_UNIT:
		return &OrganizationUnit{}, nil
	case MSPPrincipal_IDENTITY:
		return nil, fmt.Errorf("unable to decode MSP type IDENTITY until the protos are fixed to include the IDENTITY proto in protos/msp")
	default:
		return nil, fmt.Errorf("unable to decode MSP type: %v", mp.PrincipalClassification)
	}
}

