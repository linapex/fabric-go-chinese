
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:08</date>
//</624456006381867008>

//伪造者生成的代码。不要编辑。
package mock

import (
	sync "sync"

	container "github.com/hyperledger/fabric/core/container"
	ccintf "github.com/hyperledger/fabric/core/container/ccintf"
)

type VMCReq struct {
	DoStub        func(container.VM) error
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		arg1 container.VM
	}
	doReturns struct {
		result1 error
	}
	doReturnsOnCall map[int]struct {
		result1 error
	}
	GetCCIDStub        func() ccintf.CCID
	getCCIDMutex       sync.RWMutex
	getCCIDArgsForCall []struct {
	}
	getCCIDReturns struct {
		result1 ccintf.CCID
	}
	getCCIDReturnsOnCall map[int]struct {
		result1 ccintf.CCID
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *VMCReq) Do(arg1 container.VM) error {
	fake.doMutex.Lock()
	ret, specificReturn := fake.doReturnsOnCall[len(fake.doArgsForCall)]
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		arg1 container.VM
	}{arg1})
	fake.recordInvocation("Do", []interface{}{arg1})
	fake.doMutex.Unlock()
	if fake.DoStub != nil {
		return fake.DoStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.doReturns
	return fakeReturns.result1
}

func (fake *VMCReq) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *VMCReq) DoCalls(stub func(container.VM) error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = stub
}

func (fake *VMCReq) DoArgsForCall(i int) container.VM {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	argsForCall := fake.doArgsForCall[i]
	return argsForCall.arg1
}

func (fake *VMCReq) DoReturns(result1 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 error
	}{result1}
}

func (fake *VMCReq) DoReturnsOnCall(i int, result1 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	if fake.doReturnsOnCall == nil {
		fake.doReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.doReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *VMCReq) GetCCID() ccintf.CCID {
	fake.getCCIDMutex.Lock()
	ret, specificReturn := fake.getCCIDReturnsOnCall[len(fake.getCCIDArgsForCall)]
	fake.getCCIDArgsForCall = append(fake.getCCIDArgsForCall, struct {
	}{})
	fake.recordInvocation("GetCCID", []interface{}{})
	fake.getCCIDMutex.Unlock()
	if fake.GetCCIDStub != nil {
		return fake.GetCCIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getCCIDReturns
	return fakeReturns.result1
}

func (fake *VMCReq) GetCCIDCallCount() int {
	fake.getCCIDMutex.RLock()
	defer fake.getCCIDMutex.RUnlock()
	return len(fake.getCCIDArgsForCall)
}

func (fake *VMCReq) GetCCIDCalls(stub func() ccintf.CCID) {
	fake.getCCIDMutex.Lock()
	defer fake.getCCIDMutex.Unlock()
	fake.GetCCIDStub = stub
}

func (fake *VMCReq) GetCCIDReturns(result1 ccintf.CCID) {
	fake.getCCIDMutex.Lock()
	defer fake.getCCIDMutex.Unlock()
	fake.GetCCIDStub = nil
	fake.getCCIDReturns = struct {
		result1 ccintf.CCID
	}{result1}
}

func (fake *VMCReq) GetCCIDReturnsOnCall(i int, result1 ccintf.CCID) {
	fake.getCCIDMutex.Lock()
	defer fake.getCCIDMutex.Unlock()
	fake.GetCCIDStub = nil
	if fake.getCCIDReturnsOnCall == nil {
		fake.getCCIDReturnsOnCall = make(map[int]struct {
			result1 ccintf.CCID
		})
	}
	fake.getCCIDReturnsOnCall[i] = struct {
		result1 ccintf.CCID
	}{result1}
}

func (fake *VMCReq) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	fake.getCCIDMutex.RLock()
	defer fake.getCCIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *VMCReq) recordInvocation(key string, args []interface{}) {
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
