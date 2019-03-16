
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455949347721216>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package tlsgen

import (
	"crypto"
	"crypto/x509"
)

//certkeypair表示一个tls证书和相应的密钥，
//两个PEM编码
type CertKeyPair struct {
//cert是证书，pem编码
	Cert []byte
//密钥是与证书对应的密钥，PEM编码
	Key []byte

	crypto.Signer
	TLSCert *x509.Certificate
}

//CA定义可以生成的证书颁发机构
//由其签署的证书
type CA interface {
//certbytes返回采用PEM编码的CA证书
	CertBytes() []byte

//newcertkeypair返回证书和私钥对以及nil，
//或零，故障时出错
//证书由CA签名，用于TLS客户端身份验证
	NewClientCertKeyPair() (*CertKeyPair, error)

//newservercertkeypair返回certkeypair和nil，
//具有给定的自定义SAN。
//证书由CA签名。
//返回nil，失败时出错
	NewServerCertKeyPair(host string) (*CertKeyPair, error)
}

type ca struct {
	caCert *CertKeyPair
}

func NewCA() (CA, error) {
	c := &ca{}
	var err error
	c.caCert, err = newCertKeyPair(true, false, "", nil, nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

//certbytes返回采用PEM编码的CA证书
func (c *ca) CertBytes() []byte {
	return c.caCert.Cert
}

//newclientcertkeypair返回证书和私钥对以及nil，
//或零，故障时出错
//该证书由CA签名，并用作客户端TLS证书
func (c *ca) NewClientCertKeyPair() (*CertKeyPair, error) {
	return newCertKeyPair(false, false, "", c.caCert.Signer, c.caCert.TLSCert)
}

//NewServerCertKeyPair返回证书和私钥对以及nil，
//或零，故障时出错
//该证书由CA签名，并用作服务器TLS证书
func (c *ca) NewServerCertKeyPair(host string) (*CertKeyPair, error) {
	keypair, err := newCertKeyPair(false, true, host, c.caCert.Signer, c.caCert.TLSCert)
	if err != nil {
		return nil, err
	}
	return keypair, nil
}

