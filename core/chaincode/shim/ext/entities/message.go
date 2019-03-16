
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455992154787840>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package entities

import (
	"encoding/json"

	"github.com/pkg/errors"
)

//SignedMessage是一个包含空格的简单结构
//为了有效载荷和它上面的签名，以及方便
//签名、验证、封送和取消封送的功能
type SignedMessage struct {
//ID包含对此消息签名的实体的描述
	ID []byte `json:"id"`

//有效负载包含已签名的消息
	Payload []byte `json:"payload"`

//
	Sig []byte `json:"sig"`
}

//
func (m *SignedMessage) Sign(signer Signer) error {
	if signer == nil {
		return errors.New("nil signer")
	}

	m.Sig = nil
	bytes, err := json.Marshal(m)
	if err != nil {
		return errors.Wrap(err, "sign error: json.Marshal returned")
	}
	sig, err := signer.Sign(bytes)
	if err != nil {
		return errors.WithMessage(err, "sign error: signer.Sign returned")
	}
	m.Sig = sig

	return nil
}

//
func (m *SignedMessage) Verify(verifier Signer) (bool, error) {
	if verifier == nil {
		return false, errors.New("nil verifier")
	}

	sig := m.Sig
	m.Sig = nil
	defer func() {
		m.Sig = sig
	}()

	bytes, err := json.Marshal(m)
	if err != nil {
		return false, errors.Wrap(err, "sign error: json.Marshal returned")
	}

	return verifier.Verify(sig, bytes)
}

//
func (m *SignedMessage) ToBytes() ([]byte, error) {
	return json.Marshal(m)
}

//
func (m *SignedMessage) FromBytes(d []byte) error {
	return json.Unmarshal(d, m)
}

