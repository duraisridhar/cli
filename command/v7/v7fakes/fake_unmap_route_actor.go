// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeUnmapRouteActor struct {
	GetApplicationByNameAndSpaceStub        func(string, string) (v7action.Application, v7action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	GetDomainByNameStub        func(string) (v7action.Domain, v7action.Warnings, error)
	getDomainByNameMutex       sync.RWMutex
	getDomainByNameArgsForCall []struct {
		arg1 string
	}
	getDomainByNameReturns struct {
		result1 v7action.Domain
		result2 v7action.Warnings
		result3 error
	}
	getDomainByNameReturnsOnCall map[int]struct {
		result1 v7action.Domain
		result2 v7action.Warnings
		result3 error
	}
	GetRouteByAttributesStub        func(string, string, string, string) (v7action.Route, v7action.Warnings, error)
	getRouteByAttributesMutex       sync.RWMutex
	getRouteByAttributesArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	getRouteByAttributesReturns struct {
		result1 v7action.Route
		result2 v7action.Warnings
		result3 error
	}
	getRouteByAttributesReturnsOnCall map[int]struct {
		result1 v7action.Route
		result2 v7action.Warnings
		result3 error
	}
	GetRouteDestinationByAppGUIDStub        func(string, string) (v7action.RouteDestination, v7action.Warnings, error)
	getRouteDestinationByAppGUIDMutex       sync.RWMutex
	getRouteDestinationByAppGUIDArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getRouteDestinationByAppGUIDReturns struct {
		result1 v7action.RouteDestination
		result2 v7action.Warnings
		result3 error
	}
	getRouteDestinationByAppGUIDReturnsOnCall map[int]struct {
		result1 v7action.RouteDestination
		result2 v7action.Warnings
		result3 error
	}
	UnmapRouteStub        func(string, string) (v7action.Warnings, error)
	unmapRouteMutex       sync.RWMutex
	unmapRouteArgsForCall []struct {
		arg1 string
		arg2 string
	}
	unmapRouteReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	unmapRouteReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnmapRouteActor) GetApplicationByNameAndSpace(arg1 string, arg2 string) (v7action.Application, v7action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeUnmapRouteActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeUnmapRouteActor) GetApplicationByNameAndSpaceCalls(stub func(string, string) (v7action.Application, v7action.Warnings, error)) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = stub
}

func (fake *FakeUnmapRouteActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUnmapRouteActor) GetApplicationByNameAndSpaceReturns(result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUnmapRouteActor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Application
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUnmapRouteActor) GetDomainByName(arg1 string) (v7action.Domain, v7action.Warnings, error) {
	fake.getDomainByNameMutex.Lock()
	ret, specificReturn := fake.getDomainByNameReturnsOnCall[len(fake.getDomainByNameArgsForCall)]
	fake.getDomainByNameArgsForCall = append(fake.getDomainByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetDomainByName", []interface{}{arg1})
	fake.getDomainByNameMutex.Unlock()
	if fake.GetDomainByNameStub != nil {
		return fake.GetDomainByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getDomainByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeUnmapRouteActor) GetDomainByNameCallCount() int {
	fake.getDomainByNameMutex.RLock()
	defer fake.getDomainByNameMutex.RUnlock()
	return len(fake.getDomainByNameArgsForCall)
}

func (fake *FakeUnmapRouteActor) GetDomainByNameCalls(stub func(string) (v7action.Domain, v7action.Warnings, error)) {
	fake.getDomainByNameMutex.Lock()
	defer fake.getDomainByNameMutex.Unlock()
	fake.GetDomainByNameStub = stub
}

func (fake *FakeUnmapRouteActor) GetDomainByNameArgsForCall(i int) string {
	fake.getDomainByNameMutex.RLock()
	defer fake.getDomainByNameMutex.RUnlock()
	argsForCall := fake.getDomainByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUnmapRouteActor) GetDomainByNameReturns(result1 v7action.Domain, result2 v7action.Warnings, result3 error) {
	fake.getDomainByNameMutex.Lock()
	defer fake.getDomainByNameMutex.Unlock()
	fake.GetDomainByNameStub = nil
	fake.getDomainByNameReturns = struct {
		result1 v7action.Domain
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUnmapRouteActor) GetDomainByNameReturnsOnCall(i int, result1 v7action.Domain, result2 v7action.Warnings, result3 error) {
	fake.getDomainByNameMutex.Lock()
	defer fake.getDomainByNameMutex.Unlock()
	fake.GetDomainByNameStub = nil
	if fake.getDomainByNameReturnsOnCall == nil {
		fake.getDomainByNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Domain
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getDomainByNameReturnsOnCall[i] = struct {
		result1 v7action.Domain
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUnmapRouteActor) GetRouteByAttributes(arg1 string, arg2 string, arg3 string, arg4 string) (v7action.Route, v7action.Warnings, error) {
	fake.getRouteByAttributesMutex.Lock()
	ret, specificReturn := fake.getRouteByAttributesReturnsOnCall[len(fake.getRouteByAttributesArgsForCall)]
	fake.getRouteByAttributesArgsForCall = append(fake.getRouteByAttributesArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetRouteByAttributes", []interface{}{arg1, arg2, arg3, arg4})
	fake.getRouteByAttributesMutex.Unlock()
	if fake.GetRouteByAttributesStub != nil {
		return fake.GetRouteByAttributesStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getRouteByAttributesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeUnmapRouteActor) GetRouteByAttributesCallCount() int {
	fake.getRouteByAttributesMutex.RLock()
	defer fake.getRouteByAttributesMutex.RUnlock()
	return len(fake.getRouteByAttributesArgsForCall)
}

func (fake *FakeUnmapRouteActor) GetRouteByAttributesCalls(stub func(string, string, string, string) (v7action.Route, v7action.Warnings, error)) {
	fake.getRouteByAttributesMutex.Lock()
	defer fake.getRouteByAttributesMutex.Unlock()
	fake.GetRouteByAttributesStub = stub
}

func (fake *FakeUnmapRouteActor) GetRouteByAttributesArgsForCall(i int) (string, string, string, string) {
	fake.getRouteByAttributesMutex.RLock()
	defer fake.getRouteByAttributesMutex.RUnlock()
	argsForCall := fake.getRouteByAttributesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeUnmapRouteActor) GetRouteByAttributesReturns(result1 v7action.Route, result2 v7action.Warnings, result3 error) {
	fake.getRouteByAttributesMutex.Lock()
	defer fake.getRouteByAttributesMutex.Unlock()
	fake.GetRouteByAttributesStub = nil
	fake.getRouteByAttributesReturns = struct {
		result1 v7action.Route
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUnmapRouteActor) GetRouteByAttributesReturnsOnCall(i int, result1 v7action.Route, result2 v7action.Warnings, result3 error) {
	fake.getRouteByAttributesMutex.Lock()
	defer fake.getRouteByAttributesMutex.Unlock()
	fake.GetRouteByAttributesStub = nil
	if fake.getRouteByAttributesReturnsOnCall == nil {
		fake.getRouteByAttributesReturnsOnCall = make(map[int]struct {
			result1 v7action.Route
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getRouteByAttributesReturnsOnCall[i] = struct {
		result1 v7action.Route
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUnmapRouteActor) GetRouteDestinationByAppGUID(arg1 string, arg2 string) (v7action.RouteDestination, v7action.Warnings, error) {
	fake.getRouteDestinationByAppGUIDMutex.Lock()
	ret, specificReturn := fake.getRouteDestinationByAppGUIDReturnsOnCall[len(fake.getRouteDestinationByAppGUIDArgsForCall)]
	fake.getRouteDestinationByAppGUIDArgsForCall = append(fake.getRouteDestinationByAppGUIDArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetRouteDestinationByAppGUID", []interface{}{arg1, arg2})
	fake.getRouteDestinationByAppGUIDMutex.Unlock()
	if fake.GetRouteDestinationByAppGUIDStub != nil {
		return fake.GetRouteDestinationByAppGUIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getRouteDestinationByAppGUIDReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeUnmapRouteActor) GetRouteDestinationByAppGUIDCallCount() int {
	fake.getRouteDestinationByAppGUIDMutex.RLock()
	defer fake.getRouteDestinationByAppGUIDMutex.RUnlock()
	return len(fake.getRouteDestinationByAppGUIDArgsForCall)
}

func (fake *FakeUnmapRouteActor) GetRouteDestinationByAppGUIDCalls(stub func(string, string) (v7action.RouteDestination, v7action.Warnings, error)) {
	fake.getRouteDestinationByAppGUIDMutex.Lock()
	defer fake.getRouteDestinationByAppGUIDMutex.Unlock()
	fake.GetRouteDestinationByAppGUIDStub = stub
}

func (fake *FakeUnmapRouteActor) GetRouteDestinationByAppGUIDArgsForCall(i int) (string, string) {
	fake.getRouteDestinationByAppGUIDMutex.RLock()
	defer fake.getRouteDestinationByAppGUIDMutex.RUnlock()
	argsForCall := fake.getRouteDestinationByAppGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUnmapRouteActor) GetRouteDestinationByAppGUIDReturns(result1 v7action.RouteDestination, result2 v7action.Warnings, result3 error) {
	fake.getRouteDestinationByAppGUIDMutex.Lock()
	defer fake.getRouteDestinationByAppGUIDMutex.Unlock()
	fake.GetRouteDestinationByAppGUIDStub = nil
	fake.getRouteDestinationByAppGUIDReturns = struct {
		result1 v7action.RouteDestination
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUnmapRouteActor) GetRouteDestinationByAppGUIDReturnsOnCall(i int, result1 v7action.RouteDestination, result2 v7action.Warnings, result3 error) {
	fake.getRouteDestinationByAppGUIDMutex.Lock()
	defer fake.getRouteDestinationByAppGUIDMutex.Unlock()
	fake.GetRouteDestinationByAppGUIDStub = nil
	if fake.getRouteDestinationByAppGUIDReturnsOnCall == nil {
		fake.getRouteDestinationByAppGUIDReturnsOnCall = make(map[int]struct {
			result1 v7action.RouteDestination
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getRouteDestinationByAppGUIDReturnsOnCall[i] = struct {
		result1 v7action.RouteDestination
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUnmapRouteActor) UnmapRoute(arg1 string, arg2 string) (v7action.Warnings, error) {
	fake.unmapRouteMutex.Lock()
	ret, specificReturn := fake.unmapRouteReturnsOnCall[len(fake.unmapRouteArgsForCall)]
	fake.unmapRouteArgsForCall = append(fake.unmapRouteArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("UnmapRoute", []interface{}{arg1, arg2})
	fake.unmapRouteMutex.Unlock()
	if fake.UnmapRouteStub != nil {
		return fake.UnmapRouteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.unmapRouteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUnmapRouteActor) UnmapRouteCallCount() int {
	fake.unmapRouteMutex.RLock()
	defer fake.unmapRouteMutex.RUnlock()
	return len(fake.unmapRouteArgsForCall)
}

func (fake *FakeUnmapRouteActor) UnmapRouteCalls(stub func(string, string) (v7action.Warnings, error)) {
	fake.unmapRouteMutex.Lock()
	defer fake.unmapRouteMutex.Unlock()
	fake.UnmapRouteStub = stub
}

func (fake *FakeUnmapRouteActor) UnmapRouteArgsForCall(i int) (string, string) {
	fake.unmapRouteMutex.RLock()
	defer fake.unmapRouteMutex.RUnlock()
	argsForCall := fake.unmapRouteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUnmapRouteActor) UnmapRouteReturns(result1 v7action.Warnings, result2 error) {
	fake.unmapRouteMutex.Lock()
	defer fake.unmapRouteMutex.Unlock()
	fake.UnmapRouteStub = nil
	fake.unmapRouteReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnmapRouteActor) UnmapRouteReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.unmapRouteMutex.Lock()
	defer fake.unmapRouteMutex.Unlock()
	fake.UnmapRouteStub = nil
	if fake.unmapRouteReturnsOnCall == nil {
		fake.unmapRouteReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.unmapRouteReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnmapRouteActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getDomainByNameMutex.RLock()
	defer fake.getDomainByNameMutex.RUnlock()
	fake.getRouteByAttributesMutex.RLock()
	defer fake.getRouteByAttributesMutex.RUnlock()
	fake.getRouteDestinationByAppGUIDMutex.RLock()
	defer fake.getRouteDestinationByAppGUIDMutex.RUnlock()
	fake.unmapRouteMutex.RLock()
	defer fake.unmapRouteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnmapRouteActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.UnmapRouteActor = new(FakeUnmapRouteActor)
