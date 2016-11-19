// This file was generated by counterfeiter
package apihubfakes

import (
	"sync"
	"time"

	"github.com/apihub/apihub"
)

type FakeService struct {
	HostStub        func() string
	hostMutex       sync.RWMutex
	hostArgsForCall []struct{}
	hostReturns     struct {
		result1 string
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct{}
	stopReturns     struct {
		result1 error
	}
	InfoStub        func() (apihub.ServiceSpec, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct{}
	infoReturns     struct {
		result1 apihub.ServiceSpec
		result2 error
	}
	BackendsStub        func() ([]apihub.BackendInfo, error)
	backendsMutex       sync.RWMutex
	backendsArgsForCall []struct{}
	backendsReturns     struct {
		result1 []apihub.BackendInfo
		result2 error
	}
	SetTimeoutStub        func(time.Duration) error
	setTimeoutMutex       sync.RWMutex
	setTimeoutArgsForCall []struct {
		arg1 time.Duration
	}
	setTimeoutReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeService) Host() string {
	fake.hostMutex.Lock()
	fake.hostArgsForCall = append(fake.hostArgsForCall, struct{}{})
	fake.recordInvocation("Host", []interface{}{})
	fake.hostMutex.Unlock()
	if fake.HostStub != nil {
		return fake.HostStub()
	} else {
		return fake.hostReturns.result1
	}
}

func (fake *FakeService) HostCallCount() int {
	fake.hostMutex.RLock()
	defer fake.hostMutex.RUnlock()
	return len(fake.hostArgsForCall)
}

func (fake *FakeService) HostReturns(result1 string) {
	fake.HostStub = nil
	fake.hostReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeService) Start() error {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	} else {
		return fake.startReturns.result1
	}
}

func (fake *FakeService) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeService) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) Stop() error {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub()
	} else {
		return fake.stopReturns.result1
	}
}

func (fake *FakeService) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeService) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) Info() (apihub.ServiceSpec, error) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct{}{})
	fake.recordInvocation("Info", []interface{}{})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub()
	} else {
		return fake.infoReturns.result1, fake.infoReturns.result2
	}
}

func (fake *FakeService) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeService) InfoReturns(result1 apihub.ServiceSpec, result2 error) {
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 apihub.ServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeService) Backends() ([]apihub.BackendInfo, error) {
	fake.backendsMutex.Lock()
	fake.backendsArgsForCall = append(fake.backendsArgsForCall, struct{}{})
	fake.recordInvocation("Backends", []interface{}{})
	fake.backendsMutex.Unlock()
	if fake.BackendsStub != nil {
		return fake.BackendsStub()
	} else {
		return fake.backendsReturns.result1, fake.backendsReturns.result2
	}
}

func (fake *FakeService) BackendsCallCount() int {
	fake.backendsMutex.RLock()
	defer fake.backendsMutex.RUnlock()
	return len(fake.backendsArgsForCall)
}

func (fake *FakeService) BackendsReturns(result1 []apihub.BackendInfo, result2 error) {
	fake.BackendsStub = nil
	fake.backendsReturns = struct {
		result1 []apihub.BackendInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeService) SetTimeout(arg1 time.Duration) error {
	fake.setTimeoutMutex.Lock()
	fake.setTimeoutArgsForCall = append(fake.setTimeoutArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	fake.recordInvocation("SetTimeout", []interface{}{arg1})
	fake.setTimeoutMutex.Unlock()
	if fake.SetTimeoutStub != nil {
		return fake.SetTimeoutStub(arg1)
	} else {
		return fake.setTimeoutReturns.result1
	}
}

func (fake *FakeService) SetTimeoutCallCount() int {
	fake.setTimeoutMutex.RLock()
	defer fake.setTimeoutMutex.RUnlock()
	return len(fake.setTimeoutArgsForCall)
}

func (fake *FakeService) SetTimeoutArgsForCall(i int) time.Duration {
	fake.setTimeoutMutex.RLock()
	defer fake.setTimeoutMutex.RUnlock()
	return fake.setTimeoutArgsForCall[i].arg1
}

func (fake *FakeService) SetTimeoutReturns(result1 error) {
	fake.SetTimeoutStub = nil
	fake.setTimeoutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.hostMutex.RLock()
	defer fake.hostMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.backendsMutex.RLock()
	defer fake.backendsMutex.RUnlock()
	fake.setTimeoutMutex.RLock()
	defer fake.setTimeoutMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeService) recordInvocation(key string, args []interface{}) {
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

var _ apihub.Service = new(FakeService)
