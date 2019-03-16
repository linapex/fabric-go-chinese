
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:55</date>
//</624455953546219520>

//伪造者生成的代码。不要编辑。
package mock

import (
	sync "sync"

	flogging "github.com/hyperledger/fabric/common/flogging"
	zapcore "go.uber.org/zap/zapcore"
)

type Observer struct {
	CheckStub        func(zapcore.Entry, *zapcore.CheckedEntry)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 zapcore.Entry
		arg2 *zapcore.CheckedEntry
	}
	WriteEntryStub        func(zapcore.Entry, []zapcore.Field)
	writeEntryMutex       sync.RWMutex
	writeEntryArgsForCall []struct {
		arg1 zapcore.Entry
		arg2 []zapcore.Field
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Observer) Check(arg1 zapcore.Entry, arg2 *zapcore.CheckedEntry) {
	fake.checkMutex.Lock()
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 zapcore.Entry
		arg2 *zapcore.CheckedEntry
	}{arg1, arg2})
	fake.recordInvocation("Check", []interface{}{arg1, arg2})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		fake.CheckStub(arg1, arg2)
	}
}

func (fake *Observer) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *Observer) CheckCalls(stub func(zapcore.Entry, *zapcore.CheckedEntry)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *Observer) CheckArgsForCall(i int) (zapcore.Entry, *zapcore.CheckedEntry) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Observer) WriteEntry(arg1 zapcore.Entry, arg2 []zapcore.Field) {
	var arg2Copy []zapcore.Field
	if arg2 != nil {
		arg2Copy = make([]zapcore.Field, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeEntryMutex.Lock()
	fake.writeEntryArgsForCall = append(fake.writeEntryArgsForCall, struct {
		arg1 zapcore.Entry
		arg2 []zapcore.Field
	}{arg1, arg2Copy})
	fake.recordInvocation("WriteEntry", []interface{}{arg1, arg2Copy})
	fake.writeEntryMutex.Unlock()
	if fake.WriteEntryStub != nil {
		fake.WriteEntryStub(arg1, arg2)
	}
}

func (fake *Observer) WriteEntryCallCount() int {
	fake.writeEntryMutex.RLock()
	defer fake.writeEntryMutex.RUnlock()
	return len(fake.writeEntryArgsForCall)
}

func (fake *Observer) WriteEntryCalls(stub func(zapcore.Entry, []zapcore.Field)) {
	fake.writeEntryMutex.Lock()
	defer fake.writeEntryMutex.Unlock()
	fake.WriteEntryStub = stub
}

func (fake *Observer) WriteEntryArgsForCall(i int) (zapcore.Entry, []zapcore.Field) {
	fake.writeEntryMutex.RLock()
	defer fake.writeEntryMutex.RUnlock()
	argsForCall := fake.writeEntryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Observer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.writeEntryMutex.RLock()
	defer fake.writeEntryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Observer) recordInvocation(key string, args []interface{}) {
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

var _ flogging.Observer = new(Observer)
