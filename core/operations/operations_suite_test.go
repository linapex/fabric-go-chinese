
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456037759455232>

/*
版权所有IBM公司保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package operations_test

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/hyperledger/fabric-lib-go/healthz"
	"github.com/hyperledger/fabric/common/crypto/tlsgen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOperations(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Operations Suite")
}

func generateCertificates(tempDir string) {
	serverCA, err := tlsgen.NewCA()
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "server-ca.pem"), serverCA.CertBytes(), 0640)
	Expect(err).NotTo(HaveOccurred())
	serverKeyPair, err := serverCA.NewServerCertKeyPair("127.0.0.1")
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "server-cert.pem"), serverKeyPair.Cert, 0640)
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "server-key.pem"), serverKeyPair.Key, 0640)
	Expect(err).NotTo(HaveOccurred())

	clientCA, err := tlsgen.NewCA()
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "client-ca.pem"), clientCA.CertBytes(), 0640)
	Expect(err).NotTo(HaveOccurred())
	clientKeyPair, err := clientCA.NewClientCertKeyPair()
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "client-cert.pem"), clientKeyPair.Cert, 0640)
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "client-key.pem"), clientKeyPair.Key, 0640)
	Expect(err).NotTo(HaveOccurred())
}

func newHTTPClient(tlsDir string, withClientCert bool) *http.Client {
	clientCertPool := x509.NewCertPool()
	caCert, err := ioutil.ReadFile(filepath.Join(tlsDir, "server-ca.pem"))
	Expect(err).NotTo(HaveOccurred())
	clientCertPool.AppendCertsFromPEM(caCert)

	tlsClientConfig := &tls.Config{
		RootCAs: clientCertPool,
	}
	if withClientCert {
		clientCert, err := tls.LoadX509KeyPair(
			filepath.Join(tlsDir, "client-cert.pem"),
			filepath.Join(tlsDir, "client-key.pem"),
		)
		Expect(err).NotTo(HaveOccurred())
		tlsClientConfig.Certificates = []tls.Certificate{clientCert}
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsClientConfig,
		},
	}
}

//go：生成伪造者-o fakes/healthchecker.go-fake name healthchecker。健康检查者
type healthChecker interface {
	healthz.HealthChecker
}

