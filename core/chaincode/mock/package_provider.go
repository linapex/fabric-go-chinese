
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:03</date>
//</624455983816511488>

//伪造者生成的代码。不要编辑。
package mock

import (
	sync "sync"
)

type PackageProvider struct {
	GetChaincodeCodePackageStub        func(string, string) ([]byte, error)
	getChaincodeCodePackageMutex       sync.RWMutex
	getChaincodeCodePackageArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getChaincodeCodePackageReturns struct {
		result1 []byte
		result2 error
	}
	getChaincodeCodePackageReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PackageProvider) GetChaincodeCodePackage(arg1 string, arg2 string) ([]byte, error) {
	fake.getChaincodeCodePackageMutex.Lock()
	ret, specificReturn := fake.getChaincodeCodePackageReturnsOnCall[len(fake.getChaincodeCodePackageArgsForCall)]
	fake.getChaincodeCodePackageArgsForCall = append(fake.getChaincodeCodePackageArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetChaincodeCodePackage", []interface{}{arg1, arg2})
	fake.getChaincodeCodePackageMutex.Unlock()
	if fake.GetChaincodeCodePackageStub != nil {
		return fake.GetChaincodeCodePackageStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getChaincodeCodePackageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PackageProvider) GetChaincodeCodePackageCallCount() int {
	fake.getChaincodeCodePackageMutex.RLock()
	defer fake.getChaincodeCodePackageMutex.RUnlock()
	return len(fake.getChaincodeCodePackageArgsForCall)
}

func (fake *PackageProvider) GetChaincodeCodePackageCalls(stub func(string, string) ([]byte, error)) {
	fake.getChaincodeCodePackageMutex.Lock()
	defer fake.getChaincodeCodePackageMutex.Unlock()
	fake.GetChaincodeCodePackageStub = stub
}

func (fake *PackageProvider) GetChaincodeCodePackageArgsForCall(i int) (string, string) {
	fake.getChaincodeCodePackageMutex.RLock()
	defer fake.getChaincodeCodePackageMutex.RUnlock()
	argsForCall := fake.getChaincodeCodePackageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *PackageProvider) GetChaincodeCodePackageReturns(result1 []byte, result2 error) {
	fake.getChaincodeCodePackageMutex.Lock()
	defer fake.getChaincodeCodePackageMutex.Unlock()
	fake.GetChaincodeCodePackageStub = nil
	fake.getChaincodeCodePackageReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *PackageProvider) GetChaincodeCodePackageReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getChaincodeCodePackageMutex.Lock()
	defer fake.getChaincodeCodePackageMutex.Unlock()
	fake.GetChaincodeCodePackageStub = nil
	if fake.getChaincodeCodePackageReturnsOnCall == nil {
		fake.getChaincodeCodePackageReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getChaincodeCodePackageReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *PackageProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getChaincodeCodePackageMutex.RLock()
	defer fake.getChaincodeCodePackageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PackageProvider) recordInvocation(key string, args []interface{}) {
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

