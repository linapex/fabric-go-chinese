
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456128662605824>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package server

import (
	"github.com/hyperledger/fabric/token/ledger"
	"github.com/hyperledger/fabric/token/tms/plain"
	"github.com/pkg/errors"
)

//管理器实现令牌/服务器/tmsmanager接口
//TODO:将在lscc baased tms配置可用后更新
type Manager struct {
	LedgerManager ledger.LedgerManager
}

//现在它返回一个普通的发行者。
//当基于LSCC的TMS配置可用后，它将被更新。
//返回为特定通道配置的颁发者
func (manager *Manager) GetIssuer(channel string, privateCredential, publicCredential []byte) (Issuer, error) {
	return &plain.Issuer{}, nil
}

//GetTransactor返回绑定到已传递通道的事务处理程序及其凭据
//是元组（privatecredential、publiccredential）。
func (manager *Manager) GetTransactor(channel string, privateCredential, publicCredential []byte) (Transactor, error) {
	ledger, err := manager.LedgerManager.GetLedgerReader(channel)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting ledger for channel: %s", channel)
	}
	return &plain.Transactor{Ledger: ledger, PublicCredential: publicCredential}, nil
}

