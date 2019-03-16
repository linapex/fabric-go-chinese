
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:07</date>
//</624456001218678784>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package ccprovider

import (
	"path/filepath"

	"github.com/hyperledger/fabric/core/config"
)

//getchaincodeinstallpathFromViper返回安装链码的路径
func GetChaincodeInstallPathFromViper() string {
	return filepath.Join(config.GetPath("peer.fileSystemPath"), "chaincodes")
}

//loadpackage从文件系统加载chaincode包
func LoadPackage(ccname string, ccversion string, path string) (CCPackage, error) {
	return (&CCInfoFSImpl{}).GetChaincodeFromPath(ccname, ccversion, path)
}

