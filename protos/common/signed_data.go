
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:34</date>
//</624456115156946944>

/*
版权所有IBM Corp.2016保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

                 http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，

有关管理权限和
许可证限制。
**/


package common

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/util"
)

//signedData用于表示验证签名所需的一般三元组。
//这是为了跨加密方案通用，而大多数加密方案将
//在数据中包含签名标识和一个nonce，这留给加密
//实施
type SignedData struct {
	Data      []byte
	Identity  []byte
	Signature []byte
}

//可签名类型是可以将其内容映射到一组签名数据的类型。
type Signable interface {
//as signeddata将结构的签名集作为signeddata返回，或者返回一个错误，指示为什么不可能这样做。
	AsSignedData() ([]*SignedData, error)
}

//as signeddata将configupdateedevelope的签名集返回为signeddata，或者返回一个错误，指示为什么不可能这样做。
func (ce *ConfigUpdateEnvelope) AsSignedData() ([]*SignedData, error) {
	if ce == nil {
		return nil, fmt.Errorf("No signatures for nil SignedConfigItem")
	}

	result := make([]*SignedData, len(ce.Signatures))
	for i, configSig := range ce.Signatures {
		sigHeader := &SignatureHeader{}
		err := proto.Unmarshal(configSig.SignatureHeader, sigHeader)
		if err != nil {
			return nil, err
		}

		result[i] = &SignedData{
			Data:      util.ConcatenateBytes(configSig.SignatureHeader, ce.ConfigUpdate),
			Identity:  sigHeader.Creator,
			Signature: configSig.Signature,
		}

	}

	return result, nil
}

//as signeddata将信封的签名作为长度为1的signeddata切片返回，或者返回一个错误，指出不可能的原因。
func (env *Envelope) AsSignedData() ([]*SignedData, error) {
	if env == nil {
		return nil, fmt.Errorf("No signatures for nil Envelope")
	}

	payload := &Payload{}
	err := proto.Unmarshal(env.Payload, payload)
	if err != nil {
		return nil, err
	}

 /*payload.header==nil/*payload.header.signatureHeader==nil*/
  返回nil，fmt.errorf（“缺少标题”）。
 }

 shdr：=&signatureheader
 err=proto.unmashal（payload.header.signatureheader，shdr）
 如果犯错！= nIL{
  返回nil，fmt.errorf（“GetSignatureHeaderFromBytes失败，错误%s”，错误）
 }

 返回[]*签名日期
  数据：环境有效载荷，
  身份：shdr.creator，
  签名：环境签名，
 }，nIL
}

