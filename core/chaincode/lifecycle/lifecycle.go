
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:02</date>
//</624455981522227200>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package lifecycle

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/persistence"
	"github.com/pkg/errors"
)

//chaincodestore提供了一种持久化chaincode的方法
type ChaincodeStore interface {
	Save(name, version string, ccInstallPkg []byte) (hash []byte, err error)
	RetrieveHash(name, version string) (hash []byte, err error)
}

type PackageParser interface {
	Parse(data []byte) (*persistence.ChaincodePackage, error)
}

//
//由SCC以及内部
type Lifecycle struct {
	ChaincodeStore ChaincodeStore
	PackageParser  PackageParser
}

//
//它返回哈希以通过引用链码，或者在失败时返回一个错误。
func (l *Lifecycle) InstallChaincode(name, version string, chaincodeInstallPackage []byte) ([]byte, error) {
//我们先验证chaincodeinstallpackage的格式是否正确，然后再编写它。
	_, err := l.PackageParser.Parse(chaincodeInstallPackage)
	if err != nil {
		return nil, errors.WithMessage(err, "could not parse as a chaincode install package")
	}

	hash, err := l.ChaincodeStore.Save(name, version, chaincodeInstallPackage)
	if err != nil {
		return nil, errors.WithMessage(err, "could not save cc install package")
	}

	return hash, nil
}

//QueryInstalledChaincode返回给定名称和版本的已安装链码的哈希。
func (l *Lifecycle) QueryInstalledChaincode(name, version string) ([]byte, error) {
	hash, err := l.ChaincodeStore.RetrieveHash(name, version)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("could not retrieve hash for chaincode '%s:%s'", name, version))
	}

	return hash, nil
}

