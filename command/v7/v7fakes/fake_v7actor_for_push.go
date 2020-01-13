// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeV7ActorForPush struct {
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
	GetDetailedAppSummaryStub        func(string, string, bool) (v7action.DetailedApplicationSummary, v7action.Warnings, error)
	getDetailedAppSummaryMutex       sync.RWMutex
	getDetailedAppSummaryArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
	}
	getDetailedAppSummaryReturns struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}
	getDetailedAppSummaryReturnsOnCall map[int]struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}
	GetStreamingLogsForApplicationByNameAndSpaceStub        func(string, string, sharedaction.LogCacheClient) (<-chan sharedaction.LogMessage, <-chan error, context.CancelFunc, v7action.Warnings, error)
	getStreamingLogsForApplicationByNameAndSpaceMutex       sync.RWMutex
	getStreamingLogsForApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 sharedaction.LogCacheClient
	}
	getStreamingLogsForApplicationByNameAndSpaceReturns struct {
		result1 <-chan sharedaction.LogMessage
		result2 <-chan error
		result3 context.CancelFunc
		result4 v7action.Warnings
		result5 error
	}
	getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 <-chan sharedaction.LogMessage
		result2 <-chan error
		result3 context.CancelFunc
		result4 v7action.Warnings
		result5 error
	}
	RestartApplicationStub        func(string, bool) (v7action.Warnings, error)
	restartApplicationMutex       sync.RWMutex
	restartApplicationArgsForCall []struct {
		arg1 string
		arg2 bool
	}
	restartApplicationReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	restartApplicationReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	SetSpaceManifestStub        func(string, []byte) (v7action.Warnings, error)
	setSpaceManifestMutex       sync.RWMutex
	setSpaceManifestArgsForCall []struct {
		arg1 string
		arg2 []byte
	}
	setSpaceManifestReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	setSpaceManifestReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpace(arg1 string, arg2 string) (v7action.Application, v7action.Warnings, error) {
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

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpaceCalls(stub func(string, string) (v7action.Application, v7action.Warnings, error)) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = stub
}

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpaceReturns(result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7ActorForPush) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v7action.Application, result2 v7action.Warnings, result3 error) {
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

func (fake *FakeV7ActorForPush) GetDetailedAppSummary(arg1 string, arg2 string, arg3 bool) (v7action.DetailedApplicationSummary, v7action.Warnings, error) {
	fake.getDetailedAppSummaryMutex.Lock()
	ret, specificReturn := fake.getDetailedAppSummaryReturnsOnCall[len(fake.getDetailedAppSummaryArgsForCall)]
	fake.getDetailedAppSummaryArgsForCall = append(fake.getDetailedAppSummaryArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetDetailedAppSummary", []interface{}{arg1, arg2, arg3})
	fake.getDetailedAppSummaryMutex.Unlock()
	if fake.GetDetailedAppSummaryStub != nil {
		return fake.GetDetailedAppSummaryStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getDetailedAppSummaryReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7ActorForPush) GetDetailedAppSummaryCallCount() int {
	fake.getDetailedAppSummaryMutex.RLock()
	defer fake.getDetailedAppSummaryMutex.RUnlock()
	return len(fake.getDetailedAppSummaryArgsForCall)
}

func (fake *FakeV7ActorForPush) GetDetailedAppSummaryCalls(stub func(string, string, bool) (v7action.DetailedApplicationSummary, v7action.Warnings, error)) {
	fake.getDetailedAppSummaryMutex.Lock()
	defer fake.getDetailedAppSummaryMutex.Unlock()
	fake.GetDetailedAppSummaryStub = stub
}

func (fake *FakeV7ActorForPush) GetDetailedAppSummaryArgsForCall(i int) (string, string, bool) {
	fake.getDetailedAppSummaryMutex.RLock()
	defer fake.getDetailedAppSummaryMutex.RUnlock()
	argsForCall := fake.getDetailedAppSummaryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV7ActorForPush) GetDetailedAppSummaryReturns(result1 v7action.DetailedApplicationSummary, result2 v7action.Warnings, result3 error) {
	fake.getDetailedAppSummaryMutex.Lock()
	defer fake.getDetailedAppSummaryMutex.Unlock()
	fake.GetDetailedAppSummaryStub = nil
	fake.getDetailedAppSummaryReturns = struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7ActorForPush) GetDetailedAppSummaryReturnsOnCall(i int, result1 v7action.DetailedApplicationSummary, result2 v7action.Warnings, result3 error) {
	fake.getDetailedAppSummaryMutex.Lock()
	defer fake.getDetailedAppSummaryMutex.Unlock()
	fake.GetDetailedAppSummaryStub = nil
	if fake.getDetailedAppSummaryReturnsOnCall == nil {
		fake.getDetailedAppSummaryReturnsOnCall = make(map[int]struct {
			result1 v7action.DetailedApplicationSummary
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getDetailedAppSummaryReturnsOnCall[i] = struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpace(arg1 string, arg2 string, arg3 sharedaction.LogCacheClient) (<-chan sharedaction.LogMessage, <-chan error, context.CancelFunc, v7action.Warnings, error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)]
	fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall = append(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 sharedaction.LogCacheClient
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetStreamingLogsForApplicationByNameAndSpace", []interface{}{arg1, arg2, arg3})
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetStreamingLogsForApplicationByNameAndSpaceStub != nil {
		return fake.GetStreamingLogsForApplicationByNameAndSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4, ret.result5
	}
	fakeReturns := fake.getStreamingLogsForApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4, fakeReturns.result5
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpaceCallCount() int {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpaceCalls(stub func(string, string, sharedaction.LogCacheClient) (<-chan sharedaction.LogMessage, <-chan error, context.CancelFunc, v7action.Warnings, error)) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = stub
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpaceArgsForCall(i int) (string, string, sharedaction.LogCacheClient) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpaceReturns(result1 <-chan sharedaction.LogMessage, result2 <-chan error, result3 context.CancelFunc, result4 v7action.Warnings, result5 error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	fake.getStreamingLogsForApplicationByNameAndSpaceReturns = struct {
		result1 <-chan sharedaction.LogMessage
		result2 <-chan error
		result3 context.CancelFunc
		result4 v7action.Warnings
		result5 error
	}{result1, result2, result3, result4, result5}
}

func (fake *FakeV7ActorForPush) GetStreamingLogsForApplicationByNameAndSpaceReturnsOnCall(i int, result1 <-chan sharedaction.LogMessage, result2 <-chan error, result3 context.CancelFunc, result4 v7action.Warnings, result5 error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	if fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 <-chan sharedaction.LogMessage
			result2 <-chan error
			result3 context.CancelFunc
			result4 v7action.Warnings
			result5 error
		})
	}
	fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 <-chan sharedaction.LogMessage
		result2 <-chan error
		result3 context.CancelFunc
		result4 v7action.Warnings
		result5 error
	}{result1, result2, result3, result4, result5}
}

func (fake *FakeV7ActorForPush) RestartApplication(arg1 string, arg2 bool) (v7action.Warnings, error) {
	fake.restartApplicationMutex.Lock()
	ret, specificReturn := fake.restartApplicationReturnsOnCall[len(fake.restartApplicationArgsForCall)]
	fake.restartApplicationArgsForCall = append(fake.restartApplicationArgsForCall, struct {
		arg1 string
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("RestartApplication", []interface{}{arg1, arg2})
	fake.restartApplicationMutex.Unlock()
	if fake.RestartApplicationStub != nil {
		return fake.RestartApplicationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.restartApplicationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV7ActorForPush) RestartApplicationCallCount() int {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	return len(fake.restartApplicationArgsForCall)
}

func (fake *FakeV7ActorForPush) RestartApplicationCalls(stub func(string, bool) (v7action.Warnings, error)) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = stub
}

func (fake *FakeV7ActorForPush) RestartApplicationArgsForCall(i int) (string, bool) {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	argsForCall := fake.restartApplicationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV7ActorForPush) RestartApplicationReturns(result1 v7action.Warnings, result2 error) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = nil
	fake.restartApplicationReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7ActorForPush) RestartApplicationReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = nil
	if fake.restartApplicationReturnsOnCall == nil {
		fake.restartApplicationReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.restartApplicationReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7ActorForPush) SetSpaceManifest(arg1 string, arg2 []byte) (v7action.Warnings, error) {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.setSpaceManifestMutex.Lock()
	ret, specificReturn := fake.setSpaceManifestReturnsOnCall[len(fake.setSpaceManifestArgsForCall)]
	fake.setSpaceManifestArgsForCall = append(fake.setSpaceManifestArgsForCall, struct {
		arg1 string
		arg2 []byte
	}{arg1, arg2Copy})
	fake.recordInvocation("SetSpaceManifest", []interface{}{arg1, arg2Copy})
	fake.setSpaceManifestMutex.Unlock()
	if fake.SetSpaceManifestStub != nil {
		return fake.SetSpaceManifestStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setSpaceManifestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV7ActorForPush) SetSpaceManifestCallCount() int {
	fake.setSpaceManifestMutex.RLock()
	defer fake.setSpaceManifestMutex.RUnlock()
	return len(fake.setSpaceManifestArgsForCall)
}

func (fake *FakeV7ActorForPush) SetSpaceManifestCalls(stub func(string, []byte) (v7action.Warnings, error)) {
	fake.setSpaceManifestMutex.Lock()
	defer fake.setSpaceManifestMutex.Unlock()
	fake.SetSpaceManifestStub = stub
}

func (fake *FakeV7ActorForPush) SetSpaceManifestArgsForCall(i int) (string, []byte) {
	fake.setSpaceManifestMutex.RLock()
	defer fake.setSpaceManifestMutex.RUnlock()
	argsForCall := fake.setSpaceManifestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV7ActorForPush) SetSpaceManifestReturns(result1 v7action.Warnings, result2 error) {
	fake.setSpaceManifestMutex.Lock()
	defer fake.setSpaceManifestMutex.Unlock()
	fake.SetSpaceManifestStub = nil
	fake.setSpaceManifestReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7ActorForPush) SetSpaceManifestReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.setSpaceManifestMutex.Lock()
	defer fake.setSpaceManifestMutex.Unlock()
	fake.SetSpaceManifestStub = nil
	if fake.setSpaceManifestReturnsOnCall == nil {
		fake.setSpaceManifestReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.setSpaceManifestReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7ActorForPush) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getDetailedAppSummaryMutex.RLock()
	defer fake.getDetailedAppSummaryMutex.RUnlock()
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	fake.setSpaceManifestMutex.RLock()
	defer fake.setSpaceManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV7ActorForPush) recordInvocation(key string, args []interface{}) {
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

var _ v7.V7ActorForPush = new(FakeV7ActorForPush)
