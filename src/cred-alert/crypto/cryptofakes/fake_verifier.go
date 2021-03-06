// This file was generated by counterfeiter
package cryptofakes

import (
	"cred-alert/crypto"
	"sync"
)

type FakeVerifier struct {
	VerifyStub        func([]byte, []byte) error
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct {
		arg1 []byte
		arg2 []byte
	}
	verifyReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVerifier) Verify(arg1 []byte, arg2 []byte) error {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.verifyMutex.Lock()
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct {
		arg1 []byte
		arg2 []byte
	}{arg1Copy, arg2Copy})
	fake.recordInvocation("Verify", []interface{}{arg1Copy, arg2Copy})
	fake.verifyMutex.Unlock()
	if fake.VerifyStub != nil {
		return fake.VerifyStub(arg1, arg2)
	}
	return fake.verifyReturns.result1
}

func (fake *FakeVerifier) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *FakeVerifier) VerifyArgsForCall(i int) ([]byte, []byte) {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return fake.verifyArgsForCall[i].arg1, fake.verifyArgsForCall[i].arg2
}

func (fake *FakeVerifier) VerifyReturns(result1 error) {
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVerifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVerifier) recordInvocation(key string, args []interface{}) {
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

var _ crypto.Verifier = new(FakeVerifier)
