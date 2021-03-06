// This file was generated by counterfeiter
package queuefakes

import (
	"cred-alert/queue"
	"sync"
)

type FakeTask struct {
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 string
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct{}
	typeReturns     struct {
		result1 string
	}
	PayloadStub        func() string
	payloadMutex       sync.RWMutex
	payloadArgsForCall []struct{}
	payloadReturns     struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTask) ID() string {
	fake.iDMutex.Lock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	return fake.iDReturns.result1
}

func (fake *FakeTask) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeTask) IDReturns(result1 string) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTask) Type() string {
	fake.typeMutex.Lock()
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct{}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	return fake.typeReturns.result1
}

func (fake *FakeTask) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeTask) TypeReturns(result1 string) {
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTask) Payload() string {
	fake.payloadMutex.Lock()
	fake.payloadArgsForCall = append(fake.payloadArgsForCall, struct{}{})
	fake.recordInvocation("Payload", []interface{}{})
	fake.payloadMutex.Unlock()
	if fake.PayloadStub != nil {
		return fake.PayloadStub()
	}
	return fake.payloadReturns.result1
}

func (fake *FakeTask) PayloadCallCount() int {
	fake.payloadMutex.RLock()
	defer fake.payloadMutex.RUnlock()
	return len(fake.payloadArgsForCall)
}

func (fake *FakeTask) PayloadReturns(result1 string) {
	fake.PayloadStub = nil
	fake.payloadReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTask) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.payloadMutex.RLock()
	defer fake.payloadMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTask) recordInvocation(key string, args []interface{}) {
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

var _ queue.Task = new(FakeTask)
