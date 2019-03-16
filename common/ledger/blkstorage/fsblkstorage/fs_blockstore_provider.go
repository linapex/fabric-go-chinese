
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455959003009024>

/*
版权所有IBM Corp.2016保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
*/


package fsblkstorage

import (
	"github.com/hyperledger/fabric/common/ledger/blkstorage"
	"github.com/hyperledger/fabric/common/ledger/util"
	"github.com/hyperledger/fabric/common/ledger/util/leveldbhelper"
)

//
type FsBlockstoreProvider struct {
	conf            *Conf
	indexConfig     *blkstorage.IndexConfig
	leveldbProvider *leveldbhelper.Provider
}

//
func NewProvider(conf *Conf, indexConfig *blkstorage.IndexConfig) blkstorage.BlockStoreProvider {
	p := leveldbhelper.NewProvider(&leveldbhelper.Conf{DBPath: conf.getIndexDir()})
	return &FsBlockstoreProvider{conf, indexConfig, p}
}

//
func (p *FsBlockstoreProvider) CreateBlockStore(ledgerid string) (blkstorage.BlockStore, error) {
	return p.OpenBlockStore(ledgerid)
}

//
//
//
func (p *FsBlockstoreProvider) OpenBlockStore(ledgerid string) (blkstorage.BlockStore, error) {
	indexStoreHandle := p.leveldbProvider.GetDBHandle(ledgerid)
	return newFsBlockStore(ledgerid, p.conf, p.indexConfig, indexStoreHandle), nil
}

//
func (p *FsBlockstoreProvider) Exists(ledgerid string) (bool, error) {
	exists, _, err := util.FileExists(p.conf.getLedgerBlockDir(ledgerid))
	return exists, err
}

//
func (p *FsBlockstoreProvider) List() ([]string, error) {
	return util.ListSubdirs(p.conf.getChainsDir())
}

//
func (p *FsBlockstoreProvider) Close() {
	p.leveldbProvider.Close()
}

