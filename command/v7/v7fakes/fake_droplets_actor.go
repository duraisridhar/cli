// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeDropletsActor struct {
	GetApplicationDropletsStub        func(string, string) ([]v7action.Droplet, v7action.Warnings, error)
	getApplicationDropletsMutex       sync.RWMutex
	getApplicationDropletsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationDropletsReturns struct {
		result1 []v7action.Droplet
		result2 v7action.Warnings
		result3 error
	}
	getApplicationDropletsReturnsOnCall map[int]struct {
		result1 []v7action.Droplet
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDropletsActor) GetApplicationDroplets(arg1 string, arg2 string) ([]v7action.Droplet, v7action.Warnings, error) {
	fake.getApplicationDropletsMutex.Lock()
	ret, specificReturn := fake.getApplicationDropletsReturnsOnCall[len(fake.getApplicationDropletsArgsForCall)]
	fake.getApplicationDropletsArgsForCall = append(fake.getApplicationDropletsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationDroplets", []interface{}{arg1, arg2})
	fake.getApplicationDropletsMutex.Unlock()
	if fake.GetApplicationDropletsStub != nil {
		return fake.GetApplicationDropletsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationDropletsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeDropletsActor) GetApplicationDropletsCallCount() int {
	fake.getApplicationDropletsMutex.RLock()
	defer fake.getApplicationDropletsMutex.RUnlock()
	return len(fake.getApplicationDropletsArgsForCall)
}

func (fake *FakeDropletsActor) GetApplicationDropletsCalls(stub func(string, string) ([]v7action.Droplet, v7action.Warnings, error)) {
	fake.getApplicationDropletsMutex.Lock()
	defer fake.getApplicationDropletsMutex.Unlock()
	fake.GetApplicationDropletsStub = stub
}

func (fake *FakeDropletsActor) GetApplicationDropletsArgsForCall(i int) (string, string) {
	fake.getApplicationDropletsMutex.RLock()
	defer fake.getApplicationDropletsMutex.RUnlock()
	argsForCall := fake.getApplicationDropletsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDropletsActor) GetApplicationDropletsReturns(result1 []v7action.Droplet, result2 v7action.Warnings, result3 error) {
	fake.getApplicationDropletsMutex.Lock()
	defer fake.getApplicationDropletsMutex.Unlock()
	fake.GetApplicationDropletsStub = nil
	fake.getApplicationDropletsReturns = struct {
		result1 []v7action.Droplet
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDropletsActor) GetApplicationDropletsReturnsOnCall(i int, result1 []v7action.Droplet, result2 v7action.Warnings, result3 error) {
	fake.getApplicationDropletsMutex.Lock()
	defer fake.getApplicationDropletsMutex.Unlock()
	fake.GetApplicationDropletsStub = nil
	if fake.getApplicationDropletsReturnsOnCall == nil {
		fake.getApplicationDropletsReturnsOnCall = make(map[int]struct {
			result1 []v7action.Droplet
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationDropletsReturnsOnCall[i] = struct {
		result1 []v7action.Droplet
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDropletsActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationDropletsMutex.RLock()
	defer fake.getApplicationDropletsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDropletsActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.DropletsActor = new(FakeDropletsActor)
