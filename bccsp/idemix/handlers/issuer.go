
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:50</date>
//</624455930011979776>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package handlers

import (
	"github.com/hyperledger/fabric/bccsp"
	"github.com/pkg/errors"
)

//issuer secret key包含颁发者密钥
//并实现bccsp.key接口
type issuerSecretKey struct {
//sk是对颁发者密钥的IDemix引用
	sk IssuerSecretKey
//可导出如果为真，则可以通过bytes函数导出sk
	exportable bool
}

func NewIssuerSecretKey(sk IssuerSecretKey, exportable bool) *issuerSecretKey {
	return &issuerSecretKey{sk: sk, exportable: exportable}
}

func (k *issuerSecretKey) Bytes() ([]byte, error) {
	if k.exportable {
		return k.sk.Bytes()
	}

	return nil, errors.New("not exportable")
}

func (k *issuerSecretKey) SKI() []byte {
	pk, err := k.PublicKey()
	if err != nil {
		return nil
	}

	return pk.SKI()
}

func (*issuerSecretKey) Symmetric() bool {
	return false
}

func (*issuerSecretKey) Private() bool {
	return true
}

func (k *issuerSecretKey) PublicKey() (bccsp.Key, error) {
	return &issuerPublicKey{k.sk.Public()}, nil
}

//IssuerPublicKey包含颁发者公钥
//并实现bccsp.key接口
type issuerPublicKey struct {
	pk IssuerPublicKey
}

func NewIssuerPublicKey(pk IssuerPublicKey) *issuerPublicKey {
	return &issuerPublicKey{pk}
}

func (k *issuerPublicKey) Bytes() ([]byte, error) {
	return k.pk.Bytes()
}

func (k *issuerPublicKey) SKI() []byte {
	return k.pk.Hash()
}

func (*issuerPublicKey) Symmetric() bool {
	return false
}

func (*issuerPublicKey) Private() bool {
	return false
}

func (k *issuerPublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}

//issuerkeygen生成颁发者密钥。
type IssuerKeyGen struct {
//exportable是允许颁发者密钥标记为exportable的标志。
//如果密钥标记为可导出，则其byte s方法将返回密钥的字节表示形式。
	Exportable bool
//颁发者实现基础加密算法
	Issuer Issuer
}

func (g *IssuerKeyGen) KeyGen(opts bccsp.KeyGenOpts) (k bccsp.Key, err error) {
	o, ok := opts.(*bccsp.IdemixIssuerKeyGenOpts)
	if !ok {
		return nil, errors.New("invalid options, expected *bccsp.IdemixIssuerKeyGenOpts")
	}

//创建新的密钥对
	key, err := g.Issuer.NewKey(o.AttributeNames)
	if err != nil {
		return nil, err
	}

	return &issuerSecretKey{exportable: g.Exportable, sk: key}, nil
}

//IssuerPublickeyImporter导入颁发者公钥
type IssuerPublicKeyImporter struct {
//颁发者实现基础加密算法
	Issuer Issuer
}

func (i *IssuerPublicKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (k bccsp.Key, err error) {
	der, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("invalid raw, expected byte array")
	}

	if len(der) == 0 {
		return nil, errors.New("invalid raw, it must not be nil")
	}

	o, ok := opts.(*bccsp.IdemixIssuerPublicKeyImportOpts)
	if !ok {
		return nil, errors.New("invalid options, expected *bccsp.IdemixIssuerPublicKeyImportOpts")
	}

	pk, err := i.Issuer.NewPublicKeyFromBytes(raw.([]byte), o.AttributeNames)
	if err != nil {
		return nil, err
	}

	return &issuerPublicKey{pk}, nil
}

