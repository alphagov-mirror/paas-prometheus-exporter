// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	sync "sync"

	exporter "github.com/alphagov/paas-prometheus-exporter/exporter"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	prometheus "github.com/prometheus/client_golang/prometheus"
)

type FakeWatcherManager struct {
	AddWatcherStub        func(cfclient.App, prometheus.Registerer)
	addWatcherMutex       sync.RWMutex
	addWatcherArgsForCall []struct {
		arg1 cfclient.App
		arg2 prometheus.Registerer
	}
	DeleteWatcherStub        func(string)
	deleteWatcherMutex       sync.RWMutex
	deleteWatcherArgsForCall []struct {
		arg1 string
	}
	UpdateAppInstancesStub        func(cfclient.App)
	updateAppInstancesMutex       sync.RWMutex
	updateAppInstancesArgsForCall []struct {
		arg1 cfclient.App
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWatcherManager) AddWatcher(arg1 cfclient.App, arg2 prometheus.Registerer) {
	fake.addWatcherMutex.Lock()
	fake.addWatcherArgsForCall = append(fake.addWatcherArgsForCall, struct {
		arg1 cfclient.App
		arg2 prometheus.Registerer
	}{arg1, arg2})
	fake.recordInvocation("AddWatcher", []interface{}{arg1, arg2})
	fake.addWatcherMutex.Unlock()
	if fake.AddWatcherStub != nil {
		fake.AddWatcherStub(arg1, arg2)
	}
}

func (fake *FakeWatcherManager) AddWatcherCallCount() int {
	fake.addWatcherMutex.RLock()
	defer fake.addWatcherMutex.RUnlock()
	return len(fake.addWatcherArgsForCall)
}

func (fake *FakeWatcherManager) AddWatcherCalls(stub func(cfclient.App, prometheus.Registerer)) {
	fake.addWatcherMutex.Lock()
	defer fake.addWatcherMutex.Unlock()
	fake.AddWatcherStub = stub
}

func (fake *FakeWatcherManager) AddWatcherArgsForCall(i int) (cfclient.App, prometheus.Registerer) {
	fake.addWatcherMutex.RLock()
	defer fake.addWatcherMutex.RUnlock()
	argsForCall := fake.addWatcherArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeWatcherManager) DeleteWatcher(arg1 string) {
	fake.deleteWatcherMutex.Lock()
	fake.deleteWatcherArgsForCall = append(fake.deleteWatcherArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteWatcher", []interface{}{arg1})
	fake.deleteWatcherMutex.Unlock()
	if fake.DeleteWatcherStub != nil {
		fake.DeleteWatcherStub(arg1)
	}
}

func (fake *FakeWatcherManager) DeleteWatcherCallCount() int {
	fake.deleteWatcherMutex.RLock()
	defer fake.deleteWatcherMutex.RUnlock()
	return len(fake.deleteWatcherArgsForCall)
}

func (fake *FakeWatcherManager) DeleteWatcherCalls(stub func(string)) {
	fake.deleteWatcherMutex.Lock()
	defer fake.deleteWatcherMutex.Unlock()
	fake.DeleteWatcherStub = stub
}

func (fake *FakeWatcherManager) DeleteWatcherArgsForCall(i int) string {
	fake.deleteWatcherMutex.RLock()
	defer fake.deleteWatcherMutex.RUnlock()
	argsForCall := fake.deleteWatcherArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWatcherManager) UpdateAppInstances(arg1 cfclient.App) {
	fake.updateAppInstancesMutex.Lock()
	fake.updateAppInstancesArgsForCall = append(fake.updateAppInstancesArgsForCall, struct {
		arg1 cfclient.App
	}{arg1})
	fake.recordInvocation("UpdateAppInstances", []interface{}{arg1})
	fake.updateAppInstancesMutex.Unlock()
	if fake.UpdateAppInstancesStub != nil {
		fake.UpdateAppInstancesStub(arg1)
	}
}

func (fake *FakeWatcherManager) UpdateAppInstancesCallCount() int {
	fake.updateAppInstancesMutex.RLock()
	defer fake.updateAppInstancesMutex.RUnlock()
	return len(fake.updateAppInstancesArgsForCall)
}

func (fake *FakeWatcherManager) UpdateAppInstancesCalls(stub func(cfclient.App)) {
	fake.updateAppInstancesMutex.Lock()
	defer fake.updateAppInstancesMutex.Unlock()
	fake.UpdateAppInstancesStub = stub
}

func (fake *FakeWatcherManager) UpdateAppInstancesArgsForCall(i int) cfclient.App {
	fake.updateAppInstancesMutex.RLock()
	defer fake.updateAppInstancesMutex.RUnlock()
	argsForCall := fake.updateAppInstancesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWatcherManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addWatcherMutex.RLock()
	defer fake.addWatcherMutex.RUnlock()
	fake.deleteWatcherMutex.RLock()
	defer fake.deleteWatcherMutex.RUnlock()
	fake.updateAppInstancesMutex.RLock()
	defer fake.updateAppInstancesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWatcherManager) recordInvocation(key string, args []interface{}) {
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

var _ exporter.WatcherManager = new(FakeWatcherManager)