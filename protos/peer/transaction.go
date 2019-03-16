
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:36</date>
//</624456122073354240>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package peer

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/common"
)

func init() {
	common.PayloadDataMap[int32(common.HeaderType_ENDORSER_TRANSACTION)] = &Transaction{}
}

func (ta *TransactionAction) StaticallyOpaqueFields() []string {
	return []string{"header", "payload"}
}

func (ta *TransactionAction) StaticallyOpaqueFieldProto(name string) (proto.Message, error) {
	switch name {
	case ta.StaticallyOpaqueFields()[0]:
		return &common.SignatureHeader{}, nil
	case ta.StaticallyOpaqueFields()[1]:
		return &ChaincodeActionPayload{}, nil
	default:
		return nil, fmt.Errorf("not a marshaled field: %s", name)
	}
}

func (cap *ChaincodeActionPayload) StaticallyOpaqueFields() []string {
	return []string{"chaincode_proposal_payload"}
}

func (cap *ChaincodeActionPayload) StaticallyOpaqueFieldProto(name string) (proto.Message, error) {
	if name != cap.StaticallyOpaqueFields()[0] {
		return nil, fmt.Errorf("not a marshaled field: %s", name)
	}
	return &ChaincodeProposalPayload{}, nil
}

func (cae *ChaincodeEndorsedAction) StaticallyOpaqueFields() []string {
	return []string{"proposal_response_payload"}
}

func (cae *ChaincodeEndorsedAction) StaticallyOpaqueFieldProto(name string) (proto.Message, error) {
	if name != cae.StaticallyOpaqueFields()[0] {
		return nil, fmt.Errorf("not a marshaled field: %s", name)
	}
	return &ProposalResponsePayload{}, nil
}

