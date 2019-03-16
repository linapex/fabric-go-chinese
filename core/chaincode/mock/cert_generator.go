
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:02</date>
//</624455982856015872>

//伪造者生成的代码。不要编辑。
package mock

import (
	sync "sync"

	accesscontrol "github.com/hyperledger/fabric/core/chaincode/accesscontrol"
)

type CertGenerator struct {
	GenerateStub        func(string) (*accesscontrol.CertAndPrivKeyPair, error)
	generateMutex       sync.RWMutex
	generateArgsForCall []struct {
		arg1 string
	}
	generateReturns struct {
		result1 *accesscontrol.CertAndPrivKeyPair
		result2 error
	}
	generateReturnsOnCall map[int]struct {
		result1 *accesscontrol.CertAndPrivKeyPair
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CertGenerator) Generate(arg1 string) (*accesscontrol.CertAndPrivKeyPair, error) {
	fake.generateMutex.Lock()
	ret, specificReturn := fake.generateReturnsOnCall[len(fake.generateArgsForCall)]
	fake.generateArgsForCall = append(fake.generateArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Generate", []interface{}{arg1})
	fake.generateMutex.Unlock()
	if fake.GenerateStub != nil {
		return fake.GenerateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.generateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CertGenerator) GenerateCallCount() int {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return len(fake.generateArgsForCall)
}

func (fake *CertGenerator) GenerateCalls(stub func(string) (*accesscontrol.CertAndPrivKeyPair, error)) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = stub
}

func (fake *CertGenerator) GenerateArgsForCall(i int) string {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	argsForCall := fake.generateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CertGenerator) GenerateReturns(result1 *accesscontrol.CertAndPrivKeyPair, result2 error) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	fake.generateReturns = struct {
		result1 *accesscontrol.CertAndPrivKeyPair
		result2 error
	}{result1, result2}
}

func (fake *CertGenerator) GenerateReturnsOnCall(i int, result1 *accesscontrol.CertAndPrivKeyPair, result2 error) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	if fake.generateReturnsOnCall == nil {
		fake.generateReturnsOnCall = make(map[int]struct {
			result1 *accesscontrol.CertAndPrivKeyPair
			result2 error
		})
	}
	fake.generateReturnsOnCall[i] = struct {
		result1 *accesscontrol.CertAndPrivKeyPair
		result2 error
	}{result1, result2}
}

func (fake *CertGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CertGenerator) recordInvocation(key string, args []interface{}) {
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

