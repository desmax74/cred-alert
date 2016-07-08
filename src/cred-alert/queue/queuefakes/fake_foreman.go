// This file was generated by counterfeiter
package queuefakes

import (
	"cred-alert/queue"
	"sync"
)

type FakeForeman struct {
	BuildJobStub        func(queue.Task) (queue.Job, error)
	buildJobMutex       sync.RWMutex
	buildJobArgsForCall []struct {
		arg1 queue.Task
	}
	buildJobReturns struct {
		result1 queue.Job
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeForeman) BuildJob(arg1 queue.Task) (queue.Job, error) {
	fake.buildJobMutex.Lock()
	fake.buildJobArgsForCall = append(fake.buildJobArgsForCall, struct {
		arg1 queue.Task
	}{arg1})
	fake.recordInvocation("BuildJob", []interface{}{arg1})
	fake.buildJobMutex.Unlock()
	if fake.BuildJobStub != nil {
		return fake.BuildJobStub(arg1)
	} else {
		return fake.buildJobReturns.result1, fake.buildJobReturns.result2
	}
}

func (fake *FakeForeman) BuildJobCallCount() int {
	fake.buildJobMutex.RLock()
	defer fake.buildJobMutex.RUnlock()
	return len(fake.buildJobArgsForCall)
}

func (fake *FakeForeman) BuildJobArgsForCall(i int) queue.Task {
	fake.buildJobMutex.RLock()
	defer fake.buildJobMutex.RUnlock()
	return fake.buildJobArgsForCall[i].arg1
}

func (fake *FakeForeman) BuildJobReturns(result1 queue.Job, result2 error) {
	fake.BuildJobStub = nil
	fake.buildJobReturns = struct {
		result1 queue.Job
		result2 error
	}{result1, result2}
}

func (fake *FakeForeman) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildJobMutex.RLock()
	defer fake.buildJobMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeForeman) recordInvocation(key string, args []interface{}) {
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

var _ queue.Foreman = new(FakeForeman)
