// This file was generated by counterfeiter
package queuefakes

import (
	"cred-alert/queue"
	"sync"
)

type FakeAckTask struct {
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
	AckStub        func() error
	ackMutex       sync.RWMutex
	ackArgsForCall []struct{}
	ackReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAckTask) Type() string {
	fake.typeMutex.Lock()
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct{}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	} else {
		return fake.typeReturns.result1
	}
}

func (fake *FakeAckTask) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeAckTask) TypeReturns(result1 string) {
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAckTask) Payload() string {
	fake.payloadMutex.Lock()
	fake.payloadArgsForCall = append(fake.payloadArgsForCall, struct{}{})
	fake.recordInvocation("Payload", []interface{}{})
	fake.payloadMutex.Unlock()
	if fake.PayloadStub != nil {
		return fake.PayloadStub()
	} else {
		return fake.payloadReturns.result1
	}
}

func (fake *FakeAckTask) PayloadCallCount() int {
	fake.payloadMutex.RLock()
	defer fake.payloadMutex.RUnlock()
	return len(fake.payloadArgsForCall)
}

func (fake *FakeAckTask) PayloadReturns(result1 string) {
	fake.PayloadStub = nil
	fake.payloadReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAckTask) Ack() error {
	fake.ackMutex.Lock()
	fake.ackArgsForCall = append(fake.ackArgsForCall, struct{}{})
	fake.recordInvocation("Ack", []interface{}{})
	fake.ackMutex.Unlock()
	if fake.AckStub != nil {
		return fake.AckStub()
	} else {
		return fake.ackReturns.result1
	}
}

func (fake *FakeAckTask) AckCallCount() int {
	fake.ackMutex.RLock()
	defer fake.ackMutex.RUnlock()
	return len(fake.ackArgsForCall)
}

func (fake *FakeAckTask) AckReturns(result1 error) {
	fake.AckStub = nil
	fake.ackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAckTask) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.payloadMutex.RLock()
	defer fake.payloadMutex.RUnlock()
	fake.ackMutex.RLock()
	defer fake.ackMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAckTask) recordInvocation(key string, args []interface{}) {
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

var _ queue.AckTask = new(FakeAckTask)
