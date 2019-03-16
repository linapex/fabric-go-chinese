
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455962492669952>

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
**/


package localmsp

import (
	"fmt"

	"github.com/hyperledger/fabric/common/crypto"
	mspmgmt "github.com/hyperledger/fabric/msp/mgmt"
	cb "github.com/hyperledger/fabric/protos/common"
)

type mspSigner struct {
}

//NewSigner返回基于MSP的LocalSigner的新实例。
//
//
func NewSigner() crypto.LocalSigner {
	return &mspSigner{}
}

//NewSignatureHeader创建具有正确签名标识和有效nonce的SignatureHeader
func (s *mspSigner) NewSignatureHeader() (*cb.SignatureHeader, error) {
	signer, err := mspmgmt.GetLocalMSP().GetDefaultSigningIdentity()
	if err != nil {
		return nil, fmt.Errorf("Failed getting MSP-based signer [%s]", err)
	}

	creatorIdentityRaw, err := signer.Serialize()
	if err != nil {
		return nil, fmt.Errorf("Failed serializing creator public identity [%s]", err)
	}

	nonce, err := crypto.GetRandomNonce()
	if err != nil {
		return nil, fmt.Errorf("Failed creating nonce [%s]", err)
	}

	sh := &cb.SignatureHeader{}
	sh.Creator = creatorIdentityRaw
	sh.Nonce = nonce

	return sh, nil
}

//对应嵌入由NewSignatureHeader创建的签名头的邮件进行签名
func (s *mspSigner) Sign(message []byte) ([]byte, error) {
	signer, err := mspmgmt.GetLocalMSP().GetDefaultSigningIdentity()
	if err != nil {
		return nil, fmt.Errorf("Failed getting MSP-based signer [%s]", err)
	}

	signature, err := signer.Sign(message)
	if err != nil {
		return nil, fmt.Errorf("Failed generating signature [%s]", err)
	}

	return signature, nil
}

