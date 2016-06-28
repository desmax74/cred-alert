// This file was generated by counterfeiter
package loggingfakes

import (
	"cred-alert/logging"
	"sync"

	"github.com/pivotal-golang/lager"
)

type FakeCounter struct {
	IncStub        func(lager.Logger)
	incMutex       sync.RWMutex
	incArgsForCall []struct {
		arg1 lager.Logger
	}
	IncNStub        func(lager.Logger, int)
	incNMutex       sync.RWMutex
	incNArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCounter) Inc(arg1 lager.Logger) {
	fake.incMutex.Lock()
	fake.incArgsForCall = append(fake.incArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Inc", []interface{}{arg1})
	fake.incMutex.Unlock()
	if fake.IncStub != nil {
		fake.IncStub(arg1)
	}
}

func (fake *FakeCounter) IncCallCount() int {
	fake.incMutex.RLock()
	defer fake.incMutex.RUnlock()
	return len(fake.incArgsForCall)
}

func (fake *FakeCounter) IncArgsForCall(i int) lager.Logger {
	fake.incMutex.RLock()
	defer fake.incMutex.RUnlock()
	return fake.incArgsForCall[i].arg1
}

func (fake *FakeCounter) IncN(arg1 lager.Logger, arg2 int) {
	fake.incNMutex.Lock()
	fake.incNArgsForCall = append(fake.incNArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("IncN", []interface{}{arg1, arg2})
	fake.incNMutex.Unlock()
	if fake.IncNStub != nil {
		fake.IncNStub(arg1, arg2)
	}
}

func (fake *FakeCounter) IncNCallCount() int {
	fake.incNMutex.RLock()
	defer fake.incNMutex.RUnlock()
	return len(fake.incNArgsForCall)
}

func (fake *FakeCounter) IncNArgsForCall(i int) (lager.Logger, int) {
	fake.incNMutex.RLock()
	defer fake.incNMutex.RUnlock()
	return fake.incNArgsForCall[i].arg1, fake.incNArgsForCall[i].arg2
}

func (fake *FakeCounter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.incMutex.RLock()
	defer fake.incMutex.RUnlock()
	fake.incNMutex.RLock()
	defer fake.incNMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCounter) recordInvocation(key string, args []interface{}) {
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

var _ logging.Counter = new(FakeCounter)