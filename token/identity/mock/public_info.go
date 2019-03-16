
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456127643389952>

//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/token/identity"
)

type PublicInfo struct {
	PublicStub        func() []byte
	publicMutex       sync.RWMutex
	publicArgsForCall []struct{}
	publicReturns     struct {
		result1 []byte
	}
	publicReturnsOnCall map[int]struct {
		result1 []byte
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PublicInfo) Public() []byte {
	fake.publicMutex.Lock()
	ret, specificReturn := fake.publicReturnsOnCall[len(fake.publicArgsForCall)]
	fake.publicArgsForCall = append(fake.publicArgsForCall, struct{}{})
	fake.recordInvocation("Public", []interface{}{})
	fake.publicMutex.Unlock()
	if fake.PublicStub != nil {
		return fake.PublicStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.publicReturns.result1
}

func (fake *PublicInfo) PublicCallCount() int {
	fake.publicMutex.RLock()
	defer fake.publicMutex.RUnlock()
	return len(fake.publicArgsForCall)
}

func (fake *PublicInfo) PublicReturns(result1 []byte) {
	fake.PublicStub = nil
	fake.publicReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *PublicInfo) PublicReturnsOnCall(i int, result1 []byte) {
	fake.PublicStub = nil
	if fake.publicReturnsOnCall == nil {
		fake.publicReturnsOnCall = make(map[int]struct {
			result1 []byte
		})
	}
	fake.publicReturnsOnCall[i] = struct {
		result1 []byte
	}{result1}
}

func (fake *PublicInfo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.publicMutex.RLock()
	defer fake.publicMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PublicInfo) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ identity.PublicInfo = new(PublicInfo)

