
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455949632933888>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package tlsgen

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCertEncoding(t *testing.T) {
	pair, err := newCertKeyPair(false, false, "", nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, pair)
	assert.NotEmpty(t, pair.PrivKeyString())
	assert.NotEmpty(t, pair.PubKeyString())
	pair2, err := CertKeyPairFromString(pair.PrivKeyString(), pair.PubKeyString())
	assert.Equal(t, pair.Key, pair2.Key)
	assert.Equal(t, pair.Cert, pair2.Cert)
}

func TestLoadCert(t *testing.T) {
	pair, err := newCertKeyPair(false, false, "", nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, pair)
	tlsCertPair, err := tls.X509KeyPair(pair.Cert, pair.Key)
	assert.NoError(t, err)
	assert.NotNil(t, tlsCertPair)
	block, _ := pem.Decode(pair.Cert)
	cert, err := x509.ParseCertificate(block.Bytes)
	assert.NoError(t, err)
	assert.NotNil(t, cert)
}

