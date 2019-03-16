
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456128801017856>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package server

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/protos/token"
	"github.com/pkg/errors"
)

//unmashalcommand unmashal token.command消息
func UnmarshalCommand(raw []byte) (*token.Command, error) {
	command := &token.Command{}
	err := proto.Unmarshal(raw, command)
	if err != nil {
		return nil, err
	}

	return command, nil
}

type TimeFunc func() time.Time

//responseMarshaler生成token.signedMandResponse
type ResponseMarshaler struct {
	Signer  Signer
	Creator []byte
	Time    TimeFunc
}

func NewResponseMarshaler(signerID SignerIdentity) (*ResponseMarshaler, error) {
	creator, err := signerID.Serialize()
	if err != nil {
		return nil, err
	}

	return &ResponseMarshaler{
		Signer:  signerID,
		Creator: creator,
		Time:    time.Now,
	}, nil
}

func (s *ResponseMarshaler) MarshalCommandResponse(command []byte, responsePayload interface{}) (*token.SignedCommandResponse, error) {
	cr, err := commandResponseFromPayload(responsePayload)
	if err != nil {
		return nil, err
	}

	ts, err := ptypes.TimestampProto(s.Time())
	if err != nil {
		return nil, err
	}

	cr.Header = &token.CommandResponseHeader{
		Creator:     s.Creator,
		CommandHash: util.ComputeSHA256(command),
		Timestamp:   ts,
	}

	return s.createSignedCommandResponse(cr)
}

func (s *ResponseMarshaler) createSignedCommandResponse(cr *token.CommandResponse) (*token.SignedCommandResponse, error) {
	raw, err := proto.Marshal(cr)
	if err != nil {
		return nil, err
	}

	signature, err := s.Signer.Sign(raw)
	if err != nil {
		return nil, err
	}

	return &token.SignedCommandResponse{
		Response:  raw,
		Signature: signature,
	}, nil
}

func commandResponseFromPayload(payload interface{}) (*token.CommandResponse, error) {
	switch t := payload.(type) {
	case *token.CommandResponse_TokenTransaction:
		return &token.CommandResponse{Payload: t}, nil
	case *token.CommandResponse_Err:
		return &token.CommandResponse{Payload: t}, nil
	case *token.CommandResponse_UnspentTokens:
		return &token.CommandResponse{Payload: t}, nil
	default:
		return nil, errors.Errorf("command type not recognized: %T", t)
	}
}

