// This file was generated by counterfeiter
package searchfakes

import (
	"context"
	"cred-alert/search"
	"cred-alert/sniff/matchers"
	"sync"

	"code.cloudfoundry.org/lager"
)

type FakeSearcher struct {
	SearchCurrentStub        func(ctx context.Context, logger lager.Logger, matcher matchers.Matcher) search.Results
	searchCurrentMutex       sync.RWMutex
	searchCurrentArgsForCall []struct {
		ctx     context.Context
		logger  lager.Logger
		matcher matchers.Matcher
	}
	searchCurrentReturns struct {
		result1 search.Results
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSearcher) SearchCurrent(ctx context.Context, logger lager.Logger, matcher matchers.Matcher) search.Results {
	fake.searchCurrentMutex.Lock()
	fake.searchCurrentArgsForCall = append(fake.searchCurrentArgsForCall, struct {
		ctx     context.Context
		logger  lager.Logger
		matcher matchers.Matcher
	}{ctx, logger, matcher})
	fake.recordInvocation("SearchCurrent", []interface{}{ctx, logger, matcher})
	fake.searchCurrentMutex.Unlock()
	if fake.SearchCurrentStub != nil {
		return fake.SearchCurrentStub(ctx, logger, matcher)
	}
	return fake.searchCurrentReturns.result1
}

func (fake *FakeSearcher) SearchCurrentCallCount() int {
	fake.searchCurrentMutex.RLock()
	defer fake.searchCurrentMutex.RUnlock()
	return len(fake.searchCurrentArgsForCall)
}

func (fake *FakeSearcher) SearchCurrentArgsForCall(i int) (context.Context, lager.Logger, matchers.Matcher) {
	fake.searchCurrentMutex.RLock()
	defer fake.searchCurrentMutex.RUnlock()
	return fake.searchCurrentArgsForCall[i].ctx, fake.searchCurrentArgsForCall[i].logger, fake.searchCurrentArgsForCall[i].matcher
}

func (fake *FakeSearcher) SearchCurrentReturns(result1 search.Results) {
	fake.SearchCurrentStub = nil
	fake.searchCurrentReturns = struct {
		result1 search.Results
	}{result1}
}

func (fake *FakeSearcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.searchCurrentMutex.RLock()
	defer fake.searchCurrentMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSearcher) recordInvocation(key string, args []interface{}) {
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

var _ search.Searcher = new(FakeSearcher)
