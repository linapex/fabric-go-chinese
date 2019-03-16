
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:17</date>
//</624456046001262592>

/*
版权所有IBM公司保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package discovery

import (
	"encoding/hex"
	"sync"

	"github.com/hyperledger/fabric/common/util"
)

//memoizesigner使用相同的签名对邮件进行签名
//如果邮件是最近签名的
type MemoizeSigner struct {
	maxEntries uint
	sync.RWMutex
	memory map[string][]byte
	sign   Signer
}

//NewMemoizeSigner创建一个新的签署MemoizeSigner
//具有给定符号功能的消息
func NewMemoizeSigner(signFunc Signer, maxEntries uint) *MemoizeSigner {
	return &MemoizeSigner{
		maxEntries: maxEntries,
		memory:     make(map[string][]byte),
		sign:       signFunc,
	}
}

//签名者在消息上签名并返回签名和nil，
//或零，失败时出错
func (ms *MemoizeSigner) Sign(msg []byte) ([]byte, error) {
	sig, isInMemory := ms.lookup(msg)
	if isInMemory {
		return sig, nil
	}
	sig, err := ms.sign(msg)
	if err != nil {
		return nil, err
	}
	ms.memorize(msg, sig)
	return sig, nil
}

//lookup looks up the given message in memory and returns
//签名，如果消息在内存中
func (ms *MemoizeSigner) lookup(msg []byte) ([]byte, bool) {
	ms.RLock()
	defer ms.RUnlock()
	sig, exists := ms.memory[msgDigest(msg)]
	return sig, exists
}

func (ms *MemoizeSigner) memorize(msg, signature []byte) {
	if ms.maxEntries == 0 {
		return
	}
	ms.RLock()
	shouldShrink := len(ms.memory) >= (int)(ms.maxEntries)
	ms.RUnlock()

	if shouldShrink {
		ms.shrinkMemory()
	}
	ms.Lock()
	defer ms.Unlock()
	ms.memory[msgDigest(msg)] = signature

}

//逐出从内存中逐出随机消息
//直到其大小小于MaxEntries
func (ms *MemoizeSigner) shrinkMemory() {
	ms.Lock()
	defer ms.Unlock()
	for len(ms.memory) > (int)(ms.maxEntries) {
		ms.evictFromMemory()
	}
}

//从内存中逐出随机消息
func (ms *MemoizeSigner) evictFromMemory() {
	for dig := range ms.memory {
		delete(ms.memory, dig)
		return
	}
}

//msgdigest返回给定消息的摘要
func msgDigest(msg []byte) string {
	return hex.EncodeToString(util.ComputeSHA256(msg))
}

