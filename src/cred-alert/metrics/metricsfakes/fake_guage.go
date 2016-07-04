// This file was generated by counterfeiter
package metricsfakes

import (
	"cred-alert/metrics"
	"sync"

	"github.com/pivotal-golang/lager"
)

type FakeGuage struct {
	UpdateStub        func(lager.Logger, float32)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 lager.Logger
		arg2 float32
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGuage) Update(arg1 lager.Logger, arg2 float32) {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 lager.Logger
		arg2 float32
	}{arg1, arg2})
	fake.recordInvocation("Update", []interface{}{arg1, arg2})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		fake.UpdateStub(arg1, arg2)
	}
}

func (fake *FakeGuage) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeGuage) UpdateArgsForCall(i int) (lager.Logger, float32) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].arg1, fake.updateArgsForCall[i].arg2
}

func (fake *FakeGuage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeGuage) recordInvocation(key string, args []interface{}) {
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

var _ metrics.Guage = new(FakeGuage)
