
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:52</date>
//</624455937469452288>

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
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
)

type ecdsaKeyGenerator struct {
	curve elliptic.Curve
}

func (kg *ecdsaKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	privKey, err := ecdsa.GenerateKey(kg.curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("Failed generating ECDSA key for [%v]: [%s]", kg.curve, err)
	}

	return &ecdsaPrivateKey{privKey}, nil
}

type aesKeyGenerator struct {
	length int
}

func (kg *aesKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	lowLevelKey, err := GetRandomBytes(int(kg.length))
	if err != nil {
		return nil, fmt.Errorf("Failed generating AES %d key [%s]", kg.length, err)
	}

	return &aesPrivateKey{lowLevelKey, false}, nil
}

type rsaKeyGenerator struct {
	length int
}

func (kg *rsaKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	lowLevelKey, err := rsa.GenerateKey(rand.Reader, int(kg.length))

	if err != nil {
		return nil, fmt.Errorf("Failed generating RSA %d key [%s]", kg.length, err)
	}

	return &rsaPrivateKey{lowLevelKey}, nil
}

