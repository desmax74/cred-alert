// This file was generated by counterfeiter
package metricsfakes

import (
	"cred-alert/metrics"
	"sync"
)

type FakeEmitter struct {
	CounterStub        func(name string) metrics.Counter
	counterMutex       sync.RWMutex
	counterArgsForCall []struct {
		name string
	}
	counterReturns struct {
		result1 metrics.Counter
	}
	GaugeStub        func(name string) metrics.Gauge
	gaugeMutex       sync.RWMutex
	gaugeArgsForCall []struct {
		name string
	}
	gaugeReturns struct {
		result1 metrics.Gauge
	}
	TimerStub        func(name string) metrics.Timer
	timerMutex       sync.RWMutex
	timerArgsForCall []struct {
		name string
	}
	timerReturns struct {
		result1 metrics.Timer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEmitter) Counter(name string) metrics.Counter {
	fake.counterMutex.Lock()
	fake.counterArgsForCall = append(fake.counterArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Counter", []interface{}{name})
	fake.counterMutex.Unlock()
	if fake.CounterStub != nil {
		return fake.CounterStub(name)
	} else {
		return fake.counterReturns.result1
	}
}

func (fake *FakeEmitter) CounterCallCount() int {
	fake.counterMutex.RLock()
	defer fake.counterMutex.RUnlock()
	return len(fake.counterArgsForCall)
}

func (fake *FakeEmitter) CounterArgsForCall(i int) string {
	fake.counterMutex.RLock()
	defer fake.counterMutex.RUnlock()
	return fake.counterArgsForCall[i].name
}

func (fake *FakeEmitter) CounterReturns(result1 metrics.Counter) {
	fake.CounterStub = nil
	fake.counterReturns = struct {
		result1 metrics.Counter
	}{result1}
}

func (fake *FakeEmitter) Gauge(name string) metrics.Gauge {
	fake.gaugeMutex.Lock()
	fake.gaugeArgsForCall = append(fake.gaugeArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Gauge", []interface{}{name})
	fake.gaugeMutex.Unlock()
	if fake.GaugeStub != nil {
		return fake.GaugeStub(name)
	} else {
		return fake.gaugeReturns.result1
	}
}

func (fake *FakeEmitter) GaugeCallCount() int {
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	return len(fake.gaugeArgsForCall)
}

func (fake *FakeEmitter) GaugeArgsForCall(i int) string {
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	return fake.gaugeArgsForCall[i].name
}

func (fake *FakeEmitter) GaugeReturns(result1 metrics.Gauge) {
	fake.GaugeStub = nil
	fake.gaugeReturns = struct {
		result1 metrics.Gauge
	}{result1}
}

func (fake *FakeEmitter) Timer(name string) metrics.Timer {
	fake.timerMutex.Lock()
	fake.timerArgsForCall = append(fake.timerArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Timer", []interface{}{name})
	fake.timerMutex.Unlock()
	if fake.TimerStub != nil {
		return fake.TimerStub(name)
	} else {
		return fake.timerReturns.result1
	}
}

func (fake *FakeEmitter) TimerCallCount() int {
	fake.timerMutex.RLock()
	defer fake.timerMutex.RUnlock()
	return len(fake.timerArgsForCall)
}

func (fake *FakeEmitter) TimerArgsForCall(i int) string {
	fake.timerMutex.RLock()
	defer fake.timerMutex.RUnlock()
	return fake.timerArgsForCall[i].name
}

func (fake *FakeEmitter) TimerReturns(result1 metrics.Timer) {
	fake.TimerStub = nil
	fake.timerReturns = struct {
		result1 metrics.Timer
	}{result1}
}

func (fake *FakeEmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.counterMutex.RLock()
	defer fake.counterMutex.RUnlock()
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	fake.timerMutex.RLock()
	defer fake.timerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeEmitter) recordInvocation(key string, args []interface{}) {
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

var _ metrics.Emitter = new(FakeEmitter)
