// This file was generated by counterfeiter
package dbfakes

import (
	"cred-alert/db"
	"sync"
)

type FakeStatsRepository struct {
	RepositoryCountStub        func() (int, error)
	repositoryCountMutex       sync.RWMutex
	repositoryCountArgsForCall []struct{}
	repositoryCountReturns     struct {
		result1 int
		result2 error
	}
	DisabledRepositoryCountStub        func() (int, error)
	disabledRepositoryCountMutex       sync.RWMutex
	disabledRepositoryCountArgsForCall []struct{}
	disabledRepositoryCountReturns     struct {
		result1 int
		result2 error
	}
	UnclonedRepositoryCountStub        func() (int, error)
	unclonedRepositoryCountMutex       sync.RWMutex
	unclonedRepositoryCountArgsForCall []struct{}
	unclonedRepositoryCountReturns     struct {
		result1 int
		result2 error
	}
	CredentialCountStub        func() (int, error)
	credentialCountMutex       sync.RWMutex
	credentialCountArgsForCall []struct{}
	credentialCountReturns     struct {
		result1 int
		result2 error
	}
	FetchCountStub        func() (int, error)
	fetchCountMutex       sync.RWMutex
	fetchCountArgsForCall []struct{}
	fetchCountReturns     struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStatsRepository) RepositoryCount() (int, error) {
	fake.repositoryCountMutex.Lock()
	fake.repositoryCountArgsForCall = append(fake.repositoryCountArgsForCall, struct{}{})
	fake.recordInvocation("RepositoryCount", []interface{}{})
	fake.repositoryCountMutex.Unlock()
	if fake.RepositoryCountStub != nil {
		return fake.RepositoryCountStub()
	}
	return fake.repositoryCountReturns.result1, fake.repositoryCountReturns.result2
}

func (fake *FakeStatsRepository) RepositoryCountCallCount() int {
	fake.repositoryCountMutex.RLock()
	defer fake.repositoryCountMutex.RUnlock()
	return len(fake.repositoryCountArgsForCall)
}

func (fake *FakeStatsRepository) RepositoryCountReturns(result1 int, result2 error) {
	fake.RepositoryCountStub = nil
	fake.repositoryCountReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeStatsRepository) DisabledRepositoryCount() (int, error) {
	fake.disabledRepositoryCountMutex.Lock()
	fake.disabledRepositoryCountArgsForCall = append(fake.disabledRepositoryCountArgsForCall, struct{}{})
	fake.recordInvocation("DisabledRepositoryCount", []interface{}{})
	fake.disabledRepositoryCountMutex.Unlock()
	if fake.DisabledRepositoryCountStub != nil {
		return fake.DisabledRepositoryCountStub()
	}
	return fake.disabledRepositoryCountReturns.result1, fake.disabledRepositoryCountReturns.result2
}

func (fake *FakeStatsRepository) DisabledRepositoryCountCallCount() int {
	fake.disabledRepositoryCountMutex.RLock()
	defer fake.disabledRepositoryCountMutex.RUnlock()
	return len(fake.disabledRepositoryCountArgsForCall)
}

func (fake *FakeStatsRepository) DisabledRepositoryCountReturns(result1 int, result2 error) {
	fake.DisabledRepositoryCountStub = nil
	fake.disabledRepositoryCountReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeStatsRepository) UnclonedRepositoryCount() (int, error) {
	fake.unclonedRepositoryCountMutex.Lock()
	fake.unclonedRepositoryCountArgsForCall = append(fake.unclonedRepositoryCountArgsForCall, struct{}{})
	fake.recordInvocation("UnclonedRepositoryCount", []interface{}{})
	fake.unclonedRepositoryCountMutex.Unlock()
	if fake.UnclonedRepositoryCountStub != nil {
		return fake.UnclonedRepositoryCountStub()
	}
	return fake.unclonedRepositoryCountReturns.result1, fake.unclonedRepositoryCountReturns.result2
}

func (fake *FakeStatsRepository) UnclonedRepositoryCountCallCount() int {
	fake.unclonedRepositoryCountMutex.RLock()
	defer fake.unclonedRepositoryCountMutex.RUnlock()
	return len(fake.unclonedRepositoryCountArgsForCall)
}

func (fake *FakeStatsRepository) UnclonedRepositoryCountReturns(result1 int, result2 error) {
	fake.UnclonedRepositoryCountStub = nil
	fake.unclonedRepositoryCountReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeStatsRepository) CredentialCount() (int, error) {
	fake.credentialCountMutex.Lock()
	fake.credentialCountArgsForCall = append(fake.credentialCountArgsForCall, struct{}{})
	fake.recordInvocation("CredentialCount", []interface{}{})
	fake.credentialCountMutex.Unlock()
	if fake.CredentialCountStub != nil {
		return fake.CredentialCountStub()
	}
	return fake.credentialCountReturns.result1, fake.credentialCountReturns.result2
}

func (fake *FakeStatsRepository) CredentialCountCallCount() int {
	fake.credentialCountMutex.RLock()
	defer fake.credentialCountMutex.RUnlock()
	return len(fake.credentialCountArgsForCall)
}

func (fake *FakeStatsRepository) CredentialCountReturns(result1 int, result2 error) {
	fake.CredentialCountStub = nil
	fake.credentialCountReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeStatsRepository) FetchCount() (int, error) {
	fake.fetchCountMutex.Lock()
	fake.fetchCountArgsForCall = append(fake.fetchCountArgsForCall, struct{}{})
	fake.recordInvocation("FetchCount", []interface{}{})
	fake.fetchCountMutex.Unlock()
	if fake.FetchCountStub != nil {
		return fake.FetchCountStub()
	}
	return fake.fetchCountReturns.result1, fake.fetchCountReturns.result2
}

func (fake *FakeStatsRepository) FetchCountCallCount() int {
	fake.fetchCountMutex.RLock()
	defer fake.fetchCountMutex.RUnlock()
	return len(fake.fetchCountArgsForCall)
}

func (fake *FakeStatsRepository) FetchCountReturns(result1 int, result2 error) {
	fake.FetchCountStub = nil
	fake.fetchCountReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeStatsRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.repositoryCountMutex.RLock()
	defer fake.repositoryCountMutex.RUnlock()
	fake.disabledRepositoryCountMutex.RLock()
	defer fake.disabledRepositoryCountMutex.RUnlock()
	fake.unclonedRepositoryCountMutex.RLock()
	defer fake.unclonedRepositoryCountMutex.RUnlock()
	fake.credentialCountMutex.RLock()
	defer fake.credentialCountMutex.RUnlock()
	fake.fetchCountMutex.RLock()
	defer fake.fetchCountMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStatsRepository) recordInvocation(key string, args []interface{}) {
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

var _ db.StatsRepository = new(FakeStatsRepository)
