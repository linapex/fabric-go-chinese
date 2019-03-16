
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:55</date>
//</624455950274662400>

//伪造者生成的代码。不要编辑。
package mock

import (
	sync "sync"

	common "github.com/hyperledger/fabric/protos/common"
)

type BlockIterator struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	NextStub        func() (*common.Block, common.Status)
	nextMutex       sync.RWMutex
	nextArgsForCall []struct {
	}
	nextReturns struct {
		result1 *common.Block
		result2 common.Status
	}
	nextReturnsOnCall map[int]struct {
		result1 *common.Block
		result2 common.Status
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BlockIterator) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *BlockIterator) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *BlockIterator) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *BlockIterator) Next() (*common.Block, common.Status) {
	fake.nextMutex.Lock()
	ret, specificReturn := fake.nextReturnsOnCall[len(fake.nextArgsForCall)]
	fake.nextArgsForCall = append(fake.nextArgsForCall, struct {
	}{})
	fake.recordInvocation("Next", []interface{}{})
	fake.nextMutex.Unlock()
	if fake.NextStub != nil {
		return fake.NextStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.nextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BlockIterator) NextCallCount() int {
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	return len(fake.nextArgsForCall)
}

func (fake *BlockIterator) NextCalls(stub func() (*common.Block, common.Status)) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = stub
}

func (fake *BlockIterator) NextReturns(result1 *common.Block, result2 common.Status) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	fake.nextReturns = struct {
		result1 *common.Block
		result2 common.Status
	}{result1, result2}
}

func (fake *BlockIterator) NextReturnsOnCall(i int, result1 *common.Block, result2 common.Status) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	if fake.nextReturnsOnCall == nil {
		fake.nextReturnsOnCall = make(map[int]struct {
			result1 *common.Block
			result2 common.Status
		})
	}
	fake.nextReturnsOnCall[i] = struct {
		result1 *common.Block
		result2 common.Status
	}{result1, result2}
}

func (fake *BlockIterator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BlockIterator) recordInvocation(key string, args []interface{}) {
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

