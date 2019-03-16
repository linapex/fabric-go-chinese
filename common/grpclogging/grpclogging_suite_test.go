
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:56</date>
//</624455955219746816>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package grpclogging_test

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"testing"
	"time"

	"github.com/hyperledger/fabric/common/grpclogging/testpb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate protoc--proto_path=$go path/src/github.com/hyperledger/fabric/common/grpclogging/testpb--go_out=plugins=grpc:$go path/src$go path/src/github.com/hyperledger/fabric/common/grpclogging/testpb/echo.proto

func TestGrpclogging(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grpclogging Suite")
}

//go：生成仿冒者-o fakes/echo_service.go--echoserviceserver的假名。EchoServiceServer（EchoServiceServer）

type echoServiceServer interface {
	testpb.EchoServiceServer
}

func newTemplate(subjectCN string, hosts ...string) x509.Certificate {
	notBefore := time.Now().Add(-1 * time.Minute)
	notAfter := time.Now().Add(time.Duration(365 * 24 * time.Hour))

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	Expect(err).NotTo(HaveOccurred())

	template := x509.Certificate{
		Subject:               pkix.Name{CommonName: subjectCN},
		SerialNumber:          serialNumber,
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	return template
}

func pemEncode(derCert []byte, key *ecdsa.PrivateKey) (pemCert, pemKey []byte) {
	certBuf := &bytes.Buffer{}
	err := pem.Encode(certBuf, &pem.Block{Type: "CERTIFICATE", Bytes: derCert})
	Expect(err).NotTo(HaveOccurred())

	keyBytes, err := x509.MarshalECPrivateKey(key)
	Expect(err).NotTo(HaveOccurred())

	keyBuf := &bytes.Buffer{}
	err = pem.Encode(keyBuf, &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes})
	Expect(err).NotTo(HaveOccurred())

	return certBuf.Bytes(), keyBuf.Bytes()
}

func generateCA(subjectCN string, hosts ...string) (pemCert, pemKey []byte) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	Expect(err).NotTo(HaveOccurred())
	publicKey := privateKey.Public()

	template := newTemplate(subjectCN, hosts...)
	template.KeyUsage |= x509.KeyUsageCertSign
	template.IsCA = true

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey, privateKey)
	Expect(err).NotTo(HaveOccurred())

	return pemEncode(derBytes, privateKey)
}

func issueCertificate(caCert, caKey []byte, subjectCN string, hosts ...string) (pemCert, pemKey []byte) {
	tlsCert, err := tls.X509KeyPair(caCert, caKey)
	Expect(err).NotTo(HaveOccurred())

	ca, err := x509.ParseCertificate(tlsCert.Certificate[0])
	Expect(err).NotTo(HaveOccurred())

	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	Expect(err).NotTo(HaveOccurred())
	publicKey := privateKey.Public()

	template := newTemplate(subjectCN, hosts...)
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, ca, publicKey, tlsCert.PrivateKey)
	Expect(err).NotTo(HaveOccurred())

	return pemEncode(derBytes, privateKey)
}

