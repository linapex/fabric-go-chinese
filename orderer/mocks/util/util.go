
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:31</date>
//</624456103442255872>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
)

//generateMockPublicPrivateKeyPairItem返回已编码的公钥/私钥对
//作为PEM字符串。
func GenerateMockPublicPrivateKeyPairPEM(isCA bool) (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return "", "", err
	}
	privateKeyPEM := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	))

	template := x509.Certificate{
		SerialNumber: big.NewInt(100),
		Subject: pkix.Name{
			Organization: []string{"Hyperledger Fabric"},
		},
	}
	if isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	publicKeyCert, err := x509.CreateCertificate(rand.Reader, &template, &template, privateKey.Public(), privateKey)
	if err != nil {
		return "", "", err
	}

	publicKeyCertPEM := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: publicKeyCert,
		},
	))

	return publicKeyCertPEM, privateKeyPEM, nil
}

