
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455978288418816>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package chaincode

import "sync"

func NewTxKey(channelID, txID string) string { return channelID + txID }

type ActiveTransactions struct {
	mutex sync.Mutex
	ids   map[string]struct{}
}

func NewActiveTransactions() *ActiveTransactions {
	return &ActiveTransactions{
		ids: map[string]struct{}{},
	}
}

func (a *ActiveTransactions) Add(channelID, txID string) bool {
	key := NewTxKey(channelID, txID)
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if _, ok := a.ids[key]; ok {
		return false
	}

	a.ids[key] = struct{}{}
	return true
}

func (a *ActiveTransactions) Remove(channelID, txID string) {
	key := NewTxKey(channelID, txID)
	a.mutex.Lock()
	delete(a.ids, key)
	a.mutex.Unlock()
}

