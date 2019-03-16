
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:51</date>
//</624455934327918592>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package pkcs11

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/utils"
)

func (csp *impl) signECDSA(k ecdsaPrivateKey, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	r, s, err := csp.signP11ECDSA(k.ski, digest)
	if err != nil {
		return nil, err
	}

	s, _, err = utils.ToLowS(k.pub.pub, s)
	if err != nil {
		return nil, err
	}

	return utils.MarshalECDSASignature(r, s)
}

func (csp *impl) verifyECDSA(k ecdsaPublicKey, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	r, s, err := utils.UnmarshalECDSASignature(signature)
	if err != nil {
		return false, fmt.Errorf("Failed unmashalling signature [%s]", err)
	}

	lowS, err := utils.IsLowS(k.pub, s)
	if err != nil {
		return false, err
	}

	if !lowS {
		return false, fmt.Errorf("Invalid S. Must be smaller than half the order [%s][%s]", s, utils.GetCurveHalfOrdersAt(k.pub.Curve))
	}

	if csp.softVerify {
		return ecdsa.Verify(k.pub, digest, r, s), nil
	}
	return csp.verifyP11ECDSA(k.ski, digest, r, s, k.pub.Curve.Params().BitSize/8)

}

