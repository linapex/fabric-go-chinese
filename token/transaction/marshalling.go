
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:38</date>
//</624456133221814272>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package transaction

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/common"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/token"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/hyperledger/fabric/token/identity"
	"github.com/pkg/errors"
)

func UnmarshalTokenTransaction(raw []byte) (*cb.ChannelHeader, *token.TokenTransaction, identity.PublicInfo, error) {
//有效载荷…
	payload := &common.Payload{}
	err := proto.Unmarshal(raw, payload)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "error unmarshaling Payload")
	}

//签名头中的创建者
	sh, err := utils.GetSignatureHeader(payload.Header.SignatureHeader)
	if err != nil {
		return nil, nil, nil, err
	}
	creatorInfo := &TxCreatorInfo{public: sh.Creator}

	chdr, err := utils.UnmarshalChannelHeader(payload.Header.ChannelHeader)
	if err != nil {
		return nil, nil, nil, err
	}

//验证有效负载类型
	if common.HeaderType(chdr.Type) != common.HeaderType_TOKEN_TRANSACTION {
		return nil, nil, nil, errors.Errorf("only token transactions are supported, provided type: %d", chdr.Type)
	}

	ttx := &token.TokenTransaction{}
	err = proto.Unmarshal(payload.Data, ttx)
	if err != nil {
		return nil, nil, nil, errors.Errorf("failed getting token token transaction, %s", err)
	}

	return chdr, ttx, creatorInfo, nil
}

