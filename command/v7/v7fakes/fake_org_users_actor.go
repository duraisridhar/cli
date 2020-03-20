// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeOrgUsersActor struct {
	GetOrgUsersByRoleTypeStub        func(string) (map[constant.RoleType][]v7action.User, v7action.Warnings, error)
	getOrgUsersByRoleTypeMutex       sync.RWMutex
	getOrgUsersByRoleTypeArgsForCall []struct {
		arg1 string
	}
	getOrgUsersByRoleTypeReturns struct {
		result1 map[constant.RoleType][]v7action.User
		result2 v7action.Warnings
		result3 error
	}
	getOrgUsersByRoleTypeReturnsOnCall map[int]struct {
		result1 map[constant.RoleType][]v7action.User
		result2 v7action.Warnings
		result3 error
	}
	GetOrganizationByNameStub        func(string) (v7action.Organization, v7action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		arg1 string
	}
	getOrganizationByNameReturns struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrgUsersActor) GetOrgUsersByRoleType(arg1 string) (map[constant.RoleType][]v7action.User, v7action.Warnings, error) {
	fake.getOrgUsersByRoleTypeMutex.Lock()
	ret, specificReturn := fake.getOrgUsersByRoleTypeReturnsOnCall[len(fake.getOrgUsersByRoleTypeArgsForCall)]
	fake.getOrgUsersByRoleTypeArgsForCall = append(fake.getOrgUsersByRoleTypeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrgUsersByRoleType", []interface{}{arg1})
	fake.getOrgUsersByRoleTypeMutex.Unlock()
	if fake.GetOrgUsersByRoleTypeStub != nil {
		return fake.GetOrgUsersByRoleTypeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrgUsersByRoleTypeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeOrgUsersActor) GetOrgUsersByRoleTypeCallCount() int {
	fake.getOrgUsersByRoleTypeMutex.RLock()
	defer fake.getOrgUsersByRoleTypeMutex.RUnlock()
	return len(fake.getOrgUsersByRoleTypeArgsForCall)
}

func (fake *FakeOrgUsersActor) GetOrgUsersByRoleTypeCalls(stub func(string) (map[constant.RoleType][]v7action.User, v7action.Warnings, error)) {
	fake.getOrgUsersByRoleTypeMutex.Lock()
	defer fake.getOrgUsersByRoleTypeMutex.Unlock()
	fake.GetOrgUsersByRoleTypeStub = stub
}

func (fake *FakeOrgUsersActor) GetOrgUsersByRoleTypeArgsForCall(i int) string {
	fake.getOrgUsersByRoleTypeMutex.RLock()
	defer fake.getOrgUsersByRoleTypeMutex.RUnlock()
	argsForCall := fake.getOrgUsersByRoleTypeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOrgUsersActor) GetOrgUsersByRoleTypeReturns(result1 map[constant.RoleType][]v7action.User, result2 v7action.Warnings, result3 error) {
	fake.getOrgUsersByRoleTypeMutex.Lock()
	defer fake.getOrgUsersByRoleTypeMutex.Unlock()
	fake.GetOrgUsersByRoleTypeStub = nil
	fake.getOrgUsersByRoleTypeReturns = struct {
		result1 map[constant.RoleType][]v7action.User
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgUsersActor) GetOrgUsersByRoleTypeReturnsOnCall(i int, result1 map[constant.RoleType][]v7action.User, result2 v7action.Warnings, result3 error) {
	fake.getOrgUsersByRoleTypeMutex.Lock()
	defer fake.getOrgUsersByRoleTypeMutex.Unlock()
	fake.GetOrgUsersByRoleTypeStub = nil
	if fake.getOrgUsersByRoleTypeReturnsOnCall == nil {
		fake.getOrgUsersByRoleTypeReturnsOnCall = make(map[int]struct {
			result1 map[constant.RoleType][]v7action.User
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrgUsersByRoleTypeReturnsOnCall[i] = struct {
		result1 map[constant.RoleType][]v7action.User
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgUsersActor) GetOrganizationByName(arg1 string) (v7action.Organization, v7action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationByName", []interface{}{arg1})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeOrgUsersActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeOrgUsersActor) GetOrganizationByNameCalls(stub func(string) (v7action.Organization, v7action.Warnings, error)) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = stub
}

func (fake *FakeOrgUsersActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	argsForCall := fake.getOrganizationByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOrgUsersActor) GetOrganizationByNameReturns(result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgUsersActor) GetOrganizationByNameReturnsOnCall(i int, result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Organization
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgUsersActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrgUsersByRoleTypeMutex.RLock()
	defer fake.getOrgUsersByRoleTypeMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOrgUsersActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.OrgUsersActor = new(FakeOrgUsersActor)
