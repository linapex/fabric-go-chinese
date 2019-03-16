
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455992293199872>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package entities

import (
	"fmt"
	"sync"

	b "github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
)

var bccspInst b.BCCSP
var o sync.Once

func initOnce() {
	factory.InitFactories(nil)
	bccspInst = factory.GetDefault()
}

func GetEncrypterEntityForTest(id string) (EncrypterEntity, error) {
	o.Do(initOnce)

	sk, err := bccspInst.KeyGen(&b.AES256KeyGenOpts{Temporary: true})
	if err != nil {
		return nil, fmt.Errorf("GetEncrypterEntityForTest error: KeyGen returned %s", err)
	}

	ent, err := NewEncrypterEntity(id, bccspInst, sk, &b.AESCBCPKCS7ModeOpts{}, &b.AESCBCPKCS7ModeOpts{})
	if err != nil {
		return nil, fmt.Errorf("GetEncrypterEntityForTest error: NewEncrypterEntity returned %s", err)
	}

	return ent, nil
}

func GetEncrypterSignerEntityForTest(id string) (EncrypterSignerEntity, error) {
	o.Do(initOnce)

	sk_enc, err := bccspInst.KeyGen(&b.AES256KeyGenOpts{Temporary: true})
	if err != nil {
		return nil, fmt.Errorf("GetEncrypterSignerEntityForTest error: KeyGen returned %s", err)
	}

	sk_sig, err := bccspInst.KeyGen(&b.ECDSAP256KeyGenOpts{Temporary: true})
	if err != nil {
		return nil, fmt.Errorf("GetEncrypterSignerEntityForTest error: KeyGen returned %s", err)
	}

	ent, err := NewEncrypterSignerEntity(id, bccspInst, sk_enc, sk_sig, &b.AESCBCPKCS7ModeOpts{}, &b.AESCBCPKCS7ModeOpts{}, nil, &b.SHA256Opts{})
	if err != nil {
		return nil, fmt.Errorf("GetEncrypterSignerEntityForTest error: NewEncrypterSignerEntity returned %s", err)
	}

	return ent, nil
}

