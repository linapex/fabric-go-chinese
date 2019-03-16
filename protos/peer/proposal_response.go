
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:35</date>
//</624456121549066240>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package peer

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func (ppr *ProposalResponsePayload) StaticallyOpaqueFields() []string {
	return []string{"extension"}
}

func (ppr *ProposalResponsePayload) StaticallyOpaqueFieldProto(name string) (proto.Message, error) {
	if name != ppr.StaticallyOpaqueFields()[0] {
		return nil, fmt.Errorf("not a marshaled field: %s", name)
	}
	return &ChaincodeAction{}, nil
}

