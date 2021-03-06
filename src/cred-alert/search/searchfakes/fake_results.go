// This file was generated by counterfeiter
package searchfakes

import (
	"cred-alert/search"
	"sync"
)

type FakeResults struct {
	CStub        func() <-chan search.Result
	cMutex       sync.RWMutex
	cArgsForCall []struct{}
	cReturns     struct {
		result1 <-chan search.Result
	}
	ErrStub        func() error
	errMutex       sync.RWMutex
	errArgsForCall []struct{}
	errReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResults) C() <-chan search.Result {
	fake.cMutex.Lock()
	fake.cArgsForCall = append(fake.cArgsForCall, struct{}{})
	fake.recordInvocation("C", []interface{}{})
	fake.cMutex.Unlock()
	if fake.CStub != nil {
		return fake.CStub()
	}
	return fake.cReturns.result1
}

func (fake *FakeResults) CCallCount() int {
	fake.cMutex.RLock()
	defer fake.cMutex.RUnlock()
	return len(fake.cArgsForCall)
}

func (fake *FakeResults) CReturns(result1 <-chan search.Result) {
	fake.CStub = nil
	fake.cReturns = struct {
		result1 <-chan search.Result
	}{result1}
}

func (fake *FakeResults) Err() error {
	fake.errMutex.Lock()
	fake.errArgsForCall = append(fake.errArgsForCall, struct{}{})
	fake.recordInvocation("Err", []interface{}{})
	fake.errMutex.Unlock()
	if fake.ErrStub != nil {
		return fake.ErrStub()
	}
	return fake.errReturns.result1
}

func (fake *FakeResults) ErrCallCount() int {
	fake.errMutex.RLock()
	defer fake.errMutex.RUnlock()
	return len(fake.errArgsForCall)
}

func (fake *FakeResults) ErrReturns(result1 error) {
	fake.ErrStub = nil
	fake.errReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResults) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cMutex.RLock()
	defer fake.cMutex.RUnlock()
	fake.errMutex.RLock()
	defer fake.errMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeResults) recordInvocation(key string, args []interface{}) {
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

var _ search.Results = new(FakeResults)
