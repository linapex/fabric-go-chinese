
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:52</date>
//</624455938169901056>

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


package sw

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
)

type rsaSigner struct{}

func (s *rsaSigner) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	if opts == nil {
		return nil, errors.New("Invalid options. Must be different from nil.")
	}

	return k.(*rsaPrivateKey).privKey.Sign(rand.Reader, digest, opts)
}

type rsaPrivateKeyVerifier struct{}

func (v *rsaPrivateKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	if opts == nil {
		return false, errors.New("Invalid options. It must not be nil.")
	}
	switch opts.(type) {
	case *rsa.PSSOptions:
		err := rsa.VerifyPSS(&(k.(*rsaPrivateKey).privKey.PublicKey),
			(opts.(*rsa.PSSOptions)).Hash,
			digest, signature, opts.(*rsa.PSSOptions))

		return err == nil, err
	default:
		return false, fmt.Errorf("Opts type not recognized [%s]", opts)
	}
}

type rsaPublicKeyKeyVerifier struct{}

func (v *rsaPublicKeyKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	if opts == nil {
		return false, errors.New("Invalid options. It must not be nil.")
	}
	switch opts.(type) {
	case *rsa.PSSOptions:
		err := rsa.VerifyPSS(k.(*rsaPublicKey).pubKey,
			(opts.(*rsa.PSSOptions)).Hash,
			digest, signature, opts.(*rsa.PSSOptions))

		return err == nil, err
	default:
		return false, fmt.Errorf("Opts type not recognized [%s]", opts)
	}
}

