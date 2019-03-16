
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455948894736384>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package crypto

import (
	"crypto/x509"
	"encoding/pem"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/msp"
)

//expiresat在给定标识过期或零时间后返回。time
//如果我们不能确定
func ExpiresAt(identityBytes []byte) time.Time {
	sId := &msp.SerializedIdentity{}
//如果protobuf解析失败，我们将不决定到期时间。
	if err := proto.Unmarshal(identityBytes, sId); err != nil {
		return time.Time{}
	}
	bl, _ := pem.Decode(sId.IdBytes)
	if bl == nil {
//如果身份不是PEM块，我们就不会决定过期时间。
		return time.Time{}
	}
	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		return time.Time{}
	}
	return cert.NotAfter
}

