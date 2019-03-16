
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456127299457024>

//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/token/identity"
)

type Deserializer struct {
	DeserializeIdentityStub        func(serializedIdentity []byte) (msp.Identity, error)
	deserializeIdentityMutex       sync.RWMutex
	deserializeIdentityArgsForCall []struct {
		serializedIdentity []byte
	}
	deserializeIdentityReturns struct {
		result1 msp.Identity
		result2 error
	}
	deserializeIdentityReturnsOnCall map[int]struct {
		result1 msp.Identity
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Deserializer) DeserializeIdentity(serializedIdentity []byte) (msp.Identity, error) {
	var serializedIdentityCopy []byte
	if serializedIdentity != nil {
		serializedIdentityCopy = make([]byte, len(serializedIdentity))
		copy(serializedIdentityCopy, serializedIdentity)
	}
	fake.deserializeIdentityMutex.Lock()
	ret, specificReturn := fake.deserializeIdentityReturnsOnCall[len(fake.deserializeIdentityArgsForCall)]
	fake.deserializeIdentityArgsForCall = append(fake.deserializeIdentityArgsForCall, struct {
		serializedIdentity []byte
	}{serializedIdentityCopy})
	fake.recordInvocation("DeserializeIdentity", []interface{}{serializedIdentityCopy})
	fake.deserializeIdentityMutex.Unlock()
	if fake.DeserializeIdentityStub != nil {
		return fake.DeserializeIdentityStub(serializedIdentity)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deserializeIdentityReturns.result1, fake.deserializeIdentityReturns.result2
}

func (fake *Deserializer) DeserializeIdentityCallCount() int {
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	return len(fake.deserializeIdentityArgsForCall)
}

func (fake *Deserializer) DeserializeIdentityArgsForCall(i int) []byte {
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	return fake.deserializeIdentityArgsForCall[i].serializedIdentity
}

func (fake *Deserializer) DeserializeIdentityReturns(result1 msp.Identity, result2 error) {
	fake.DeserializeIdentityStub = nil
	fake.deserializeIdentityReturns = struct {
		result1 msp.Identity
		result2 error
	}{result1, result2}
}

func (fake *Deserializer) DeserializeIdentityReturnsOnCall(i int, result1 msp.Identity, result2 error) {
	fake.DeserializeIdentityStub = nil
	if fake.deserializeIdentityReturnsOnCall == nil {
		fake.deserializeIdentityReturnsOnCall = make(map[int]struct {
			result1 msp.Identity
			result2 error
		})
	}
	fake.deserializeIdentityReturnsOnCall[i] = struct {
		result1 msp.Identity
		result2 error
	}{result1, result2}
}

func (fake *Deserializer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Deserializer) recordInvocation(key string, args []interface{}) {
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

var _ identity.Deserializer = new(Deserializer)

