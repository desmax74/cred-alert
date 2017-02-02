// This file was generated by counterfeiter
package notificationsfakes

import (
	"cred-alert/notifications"
	"sync"
)

type FakeSlackNotificationFormatter struct {
	FormatNotificationsStub        func(batch []notifications.Notification) []notifications.SlackMessage
	formatNotificationsMutex       sync.RWMutex
	formatNotificationsArgsForCall []struct {
		batch []notifications.Notification
	}
	formatNotificationsReturns struct {
		result1 []notifications.SlackMessage
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSlackNotificationFormatter) FormatNotifications(batch []notifications.Notification) []notifications.SlackMessage {
	var batchCopy []notifications.Notification
	if batch != nil {
		batchCopy = make([]notifications.Notification, len(batch))
		copy(batchCopy, batch)
	}
	fake.formatNotificationsMutex.Lock()
	fake.formatNotificationsArgsForCall = append(fake.formatNotificationsArgsForCall, struct {
		batch []notifications.Notification
	}{batchCopy})
	fake.recordInvocation("FormatNotifications", []interface{}{batchCopy})
	fake.formatNotificationsMutex.Unlock()
	if fake.FormatNotificationsStub != nil {
		return fake.FormatNotificationsStub(batch)
	}
	return fake.formatNotificationsReturns.result1
}

func (fake *FakeSlackNotificationFormatter) FormatNotificationsCallCount() int {
	fake.formatNotificationsMutex.RLock()
	defer fake.formatNotificationsMutex.RUnlock()
	return len(fake.formatNotificationsArgsForCall)
}

func (fake *FakeSlackNotificationFormatter) FormatNotificationsArgsForCall(i int) []notifications.Notification {
	fake.formatNotificationsMutex.RLock()
	defer fake.formatNotificationsMutex.RUnlock()
	return fake.formatNotificationsArgsForCall[i].batch
}

func (fake *FakeSlackNotificationFormatter) FormatNotificationsReturns(result1 []notifications.SlackMessage) {
	fake.FormatNotificationsStub = nil
	fake.formatNotificationsReturns = struct {
		result1 []notifications.SlackMessage
	}{result1}
}

func (fake *FakeSlackNotificationFormatter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.formatNotificationsMutex.RLock()
	defer fake.formatNotificationsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSlackNotificationFormatter) recordInvocation(key string, args []interface{}) {
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

var _ notifications.SlackNotificationFormatter = new(FakeSlackNotificationFormatter)
