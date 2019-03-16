
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:18</date>
//</624456049440591872>

//伪造者生成的代码。不要编辑。
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric/discovery/support/config"
	"github.com/hyperledger/fabric/protos/common"
)

type ConfigBlockGetter struct {
	GetCurrConfigBlockStub        func(channel string) *common.Block
	getCurrConfigBlockMutex       sync.RWMutex
	getCurrConfigBlockArgsForCall []struct {
		channel string
	}
	getCurrConfigBlockReturns struct {
		result1 *common.Block
	}
	getCurrConfigBlockReturnsOnCall map[int]struct {
		result1 *common.Block
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ConfigBlockGetter) GetCurrConfigBlock(channel string) *common.Block {
	fake.getCurrConfigBlockMutex.Lock()
	ret, specificReturn := fake.getCurrConfigBlockReturnsOnCall[len(fake.getCurrConfigBlockArgsForCall)]
	fake.getCurrConfigBlockArgsForCall = append(fake.getCurrConfigBlockArgsForCall, struct {
		channel string
	}{channel})
	fake.recordInvocation("GetCurrConfigBlock", []interface{}{channel})
	fake.getCurrConfigBlockMutex.Unlock()
	if fake.GetCurrConfigBlockStub != nil {
		return fake.GetCurrConfigBlockStub(channel)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getCurrConfigBlockReturns.result1
}

func (fake *ConfigBlockGetter) GetCurrConfigBlockCallCount() int {
	fake.getCurrConfigBlockMutex.RLock()
	defer fake.getCurrConfigBlockMutex.RUnlock()
	return len(fake.getCurrConfigBlockArgsForCall)
}

func (fake *ConfigBlockGetter) GetCurrConfigBlockArgsForCall(i int) string {
	fake.getCurrConfigBlockMutex.RLock()
	defer fake.getCurrConfigBlockMutex.RUnlock()
	return fake.getCurrConfigBlockArgsForCall[i].channel
}

func (fake *ConfigBlockGetter) GetCurrConfigBlockReturns(result1 *common.Block) {
	fake.GetCurrConfigBlockStub = nil
	fake.getCurrConfigBlockReturns = struct {
		result1 *common.Block
	}{result1}
}

func (fake *ConfigBlockGetter) GetCurrConfigBlockReturnsOnCall(i int, result1 *common.Block) {
	fake.GetCurrConfigBlockStub = nil
	if fake.getCurrConfigBlockReturnsOnCall == nil {
		fake.getCurrConfigBlockReturnsOnCall = make(map[int]struct {
			result1 *common.Block
		})
	}
	fake.getCurrConfigBlockReturnsOnCall[i] = struct {
		result1 *common.Block
	}{result1}
}

func (fake *ConfigBlockGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCurrConfigBlockMutex.RLock()
	defer fake.getCurrConfigBlockMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ConfigBlockGetter) recordInvocation(key string, args []interface{}) {
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

var _ config.CurrentConfigBlockGetter = new(ConfigBlockGetter)
