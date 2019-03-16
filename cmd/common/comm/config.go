
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:52</date>
//</624455940145418240>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package comm

import (
	"io/ioutil"
	"time"

	"github.com/hyperledger/fabric/common/crypto/tlsgen"
	"github.com/hyperledger/fabric/core/comm"
	"github.com/pkg/errors"
)

type genTLSCertFunc func() (*tlsgen.CertKeyPair, error)

//config定义客户端的配置
type Config struct {
	CertPath       string
	KeyPath        string
	PeerCACertPath string
	Timeout        time.Duration
}

//ToSecureOptions将此配置转换为SecureOptions。
//如果
//配置中不存在TLS证书和密钥
func (conf Config) ToSecureOptions(newSelfSignedTLSCert genTLSCertFunc) (*comm.SecureOptions, error) {
	if conf.PeerCACertPath == "" {
		return &comm.SecureOptions{}, nil
	}
	caBytes, err := loadFile(conf.PeerCACertPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var keyBytes, certBytes []byte
//如果未提供TLS密钥和证书，则动态生成自签名密钥和证书
	if conf.KeyPath == "" && conf.CertPath == "" {
		tlsCert, err := newSelfSignedTLSCert()
		if err != nil {
			return nil, err
		}
		keyBytes, certBytes = tlsCert.Key, tlsCert.Cert
	} else {
		keyBytes, err = loadFile(conf.KeyPath)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		certBytes, err = loadFile(conf.CertPath)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return &comm.SecureOptions{
		Key:               keyBytes,
		Certificate:       certBytes,
		UseTLS:            true,
		ServerRootCAs:     [][]byte{caBytes},
		RequireClientCert: true,
	}, nil
}

func loadFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Errorf("Failed opening file %s: %v", path, err)
	}
	return b, nil
}

