
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:31</date>
//</624456100959227904>

//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/common/metrics"
)

type MetricsGauge struct {
	WithStub        func(labelValues ...string) metrics.Gauge
	withMutex       sync.RWMutex
	withArgsForCall []struct {
		labelValues []string
	}
	withReturns struct {
		result1 metrics.Gauge
	}
	withReturnsOnCall map[int]struct {
		result1 metrics.Gauge
	}
	AddStub        func(delta float64)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		delta float64
	}
	SetStub        func(value float64)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		value float64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MetricsGauge) With(labelValues ...string) metrics.Gauge {
	fake.withMutex.Lock()
	ret, specificReturn := fake.withReturnsOnCall[len(fake.withArgsForCall)]
	fake.withArgsForCall = append(fake.withArgsForCall, struct {
		labelValues []string
	}{labelValues})
	fake.recordInvocation("With", []interface{}{labelValues})
	fake.withMutex.Unlock()
	if fake.WithStub != nil {
		return fake.WithStub(labelValues...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.withReturns.result1
}

func (fake *MetricsGauge) WithCallCount() int {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return len(fake.withArgsForCall)
}

func (fake *MetricsGauge) WithArgsForCall(i int) []string {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return fake.withArgsForCall[i].labelValues
}

func (fake *MetricsGauge) WithReturns(result1 metrics.Gauge) {
	fake.WithStub = nil
	fake.withReturns = struct {
		result1 metrics.Gauge
	}{result1}
}

func (fake *MetricsGauge) WithReturnsOnCall(i int, result1 metrics.Gauge) {
	fake.WithStub = nil
	if fake.withReturnsOnCall == nil {
		fake.withReturnsOnCall = make(map[int]struct {
			result1 metrics.Gauge
		})
	}
	fake.withReturnsOnCall[i] = struct {
		result1 metrics.Gauge
	}{result1}
}

func (fake *MetricsGauge) Add(delta float64) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		delta float64
	}{delta})
	fake.recordInvocation("Add", []interface{}{delta})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		fake.AddStub(delta)
	}
}

func (fake *MetricsGauge) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *MetricsGauge) AddArgsForCall(i int) float64 {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].delta
}

func (fake *MetricsGauge) Set(value float64) {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		value float64
	}{value})
	fake.recordInvocation("Set", []interface{}{value})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		fake.SetStub(value)
	}
}

func (fake *MetricsGauge) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *MetricsGauge) SetArgsForCall(i int) float64 {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return fake.setArgsForCall[i].value
}

func (fake *MetricsGauge) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MetricsGauge) recordInvocation(key string, args []interface{}) {
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
