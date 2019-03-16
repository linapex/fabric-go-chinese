
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:52</date>
//</624455939566604288>

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


package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"math/big"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDERToX509Certificate(t *testing.T) {
	testExtKeyUsage := []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}
	testUnknownExtKeyUsage := []asn1.ObjectIdentifier{[]int{1, 2, 3}, []int{2, 59, 1}}
	extraExtensionData := []byte("extra extension")
	commonName := "test.example.com"
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{"Σ Acme Co"},
			Country:      []string{"US"},
			ExtraNames: []pkix.AttributeTypeAndValue{
				{
					Type:  []int{2, 5, 4, 42},
					Value: "Gopher",
				},
//这应该覆盖国家，如上所述。
				{
					Type:  []int{2, 5, 4, 6},
					Value: "NL",
				},
			},
		},
		NotBefore: time.Now().Add(-1 * time.Hour),
		NotAfter:  time.Now().Add(1 * time.Hour),

		SignatureAlgorithm: x509.ECDSAWithSHA256,

		SubjectKeyId: []byte{1, 2, 3, 4},
		KeyUsage:     x509.KeyUsageCertSign,

		ExtKeyUsage:        testExtKeyUsage,
		UnknownExtKeyUsage: testUnknownExtKeyUsage,

		BasicConstraintsValid: true,
		IsCA:                  true,

OCSPServer:            []string{"http://occurrentbccsp.example.com“，
IssuingCertificateURL: []string{"http://crt.example.com/ca1.crt“，

		DNSNames:       []string{"test.example.com"},
		EmailAddresses: []string{"gopher@golang.org"},
		IPAddresses:    []net.IP{net.IPv4(127, 0, 0, 1).To4(), net.ParseIP("2001:4860:0:2001::68")},

		PolicyIdentifiers:   []asn1.ObjectIdentifier{[]int{1, 2, 3}},
		PermittedDNSDomains: []string{".example.com", "example.com"},

CRLDistributionPoints: []string{"http://crl1.example.com/ca1.crl“，”http://crl2.example.com/ca1.crl”，

		ExtraExtensions: []pkix.Extension{
			{
				Id:    []int{1, 2, 3, 4},
				Value: extraExtensionData,
			},
		},
	}

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	certRaw, err := x509.CreateCertificate(rand.Reader, &template, &template, key.Public(), key)
	assert.NoError(t, err)

	cert, err := DERToX509Certificate(certRaw)
	assert.NoError(t, err)
	assert.NotNil(t, cert)
	assert.Equal(t, cert.Raw, certRaw)

}

