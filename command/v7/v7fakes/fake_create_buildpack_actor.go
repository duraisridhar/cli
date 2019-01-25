// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	sync "sync"

	v7action "code.cloudfoundry.org/cli/actor/v7action"
	ccv3 "code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeCreateBuildpackActor struct {
	CreateBuildpackStub        func(v7action.Buildpack) (v7action.Buildpack, v7action.Warnings, error)
	createBuildpackMutex       sync.RWMutex
	createBuildpackArgsForCall []struct {
		arg1 v7action.Buildpack
	}
	createBuildpackReturns struct {
		result1 v7action.Buildpack
		result2 v7action.Warnings
		result3 error
	}
	createBuildpackReturnsOnCall map[int]struct {
		result1 v7action.Buildpack
		result2 v7action.Warnings
		result3 error
	}
	PollUploadBuildpackJobStub        func(ccv3.JobURL) (v7action.Warnings, error)
	pollUploadBuildpackJobMutex       sync.RWMutex
	pollUploadBuildpackJobArgsForCall []struct {
		arg1 ccv3.JobURL
	}
	pollUploadBuildpackJobReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	pollUploadBuildpackJobReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	PrepareBuildpackBitsStub        func(string, string, v7action.Downloader) (string, error)
	prepareBuildpackBitsMutex       sync.RWMutex
	prepareBuildpackBitsArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 v7action.Downloader
	}
	prepareBuildpackBitsReturns struct {
		result1 string
		result2 error
	}
	prepareBuildpackBitsReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UploadBuildpackStub        func(string, string, v7action.SimpleProgressBar) (ccv3.JobURL, v7action.Warnings, error)
	uploadBuildpackMutex       sync.RWMutex
	uploadBuildpackArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 v7action.SimpleProgressBar
	}
	uploadBuildpackReturns struct {
		result1 ccv3.JobURL
		result2 v7action.Warnings
		result3 error
	}
	uploadBuildpackReturnsOnCall map[int]struct {
		result1 ccv3.JobURL
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreateBuildpackActor) CreateBuildpack(arg1 v7action.Buildpack) (v7action.Buildpack, v7action.Warnings, error) {
	fake.createBuildpackMutex.Lock()
	ret, specificReturn := fake.createBuildpackReturnsOnCall[len(fake.createBuildpackArgsForCall)]
	fake.createBuildpackArgsForCall = append(fake.createBuildpackArgsForCall, struct {
		arg1 v7action.Buildpack
	}{arg1})
	fake.recordInvocation("CreateBuildpack", []interface{}{arg1})
	fake.createBuildpackMutex.Unlock()
	if fake.CreateBuildpackStub != nil {
		return fake.CreateBuildpackStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createBuildpackReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCreateBuildpackActor) CreateBuildpackCallCount() int {
	fake.createBuildpackMutex.RLock()
	defer fake.createBuildpackMutex.RUnlock()
	return len(fake.createBuildpackArgsForCall)
}

func (fake *FakeCreateBuildpackActor) CreateBuildpackCalls(stub func(v7action.Buildpack) (v7action.Buildpack, v7action.Warnings, error)) {
	fake.createBuildpackMutex.Lock()
	defer fake.createBuildpackMutex.Unlock()
	fake.CreateBuildpackStub = stub
}

func (fake *FakeCreateBuildpackActor) CreateBuildpackArgsForCall(i int) v7action.Buildpack {
	fake.createBuildpackMutex.RLock()
	defer fake.createBuildpackMutex.RUnlock()
	argsForCall := fake.createBuildpackArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCreateBuildpackActor) CreateBuildpackReturns(result1 v7action.Buildpack, result2 v7action.Warnings, result3 error) {
	fake.createBuildpackMutex.Lock()
	defer fake.createBuildpackMutex.Unlock()
	fake.CreateBuildpackStub = nil
	fake.createBuildpackReturns = struct {
		result1 v7action.Buildpack
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateBuildpackActor) CreateBuildpackReturnsOnCall(i int, result1 v7action.Buildpack, result2 v7action.Warnings, result3 error) {
	fake.createBuildpackMutex.Lock()
	defer fake.createBuildpackMutex.Unlock()
	fake.CreateBuildpackStub = nil
	if fake.createBuildpackReturnsOnCall == nil {
		fake.createBuildpackReturnsOnCall = make(map[int]struct {
			result1 v7action.Buildpack
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.createBuildpackReturnsOnCall[i] = struct {
		result1 v7action.Buildpack
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateBuildpackActor) PollUploadBuildpackJob(arg1 ccv3.JobURL) (v7action.Warnings, error) {
	fake.pollUploadBuildpackJobMutex.Lock()
	ret, specificReturn := fake.pollUploadBuildpackJobReturnsOnCall[len(fake.pollUploadBuildpackJobArgsForCall)]
	fake.pollUploadBuildpackJobArgsForCall = append(fake.pollUploadBuildpackJobArgsForCall, struct {
		arg1 ccv3.JobURL
	}{arg1})
	fake.recordInvocation("PollUploadBuildpackJob", []interface{}{arg1})
	fake.pollUploadBuildpackJobMutex.Unlock()
	if fake.PollUploadBuildpackJobStub != nil {
		return fake.PollUploadBuildpackJobStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.pollUploadBuildpackJobReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreateBuildpackActor) PollUploadBuildpackJobCallCount() int {
	fake.pollUploadBuildpackJobMutex.RLock()
	defer fake.pollUploadBuildpackJobMutex.RUnlock()
	return len(fake.pollUploadBuildpackJobArgsForCall)
}

func (fake *FakeCreateBuildpackActor) PollUploadBuildpackJobCalls(stub func(ccv3.JobURL) (v7action.Warnings, error)) {
	fake.pollUploadBuildpackJobMutex.Lock()
	defer fake.pollUploadBuildpackJobMutex.Unlock()
	fake.PollUploadBuildpackJobStub = stub
}

func (fake *FakeCreateBuildpackActor) PollUploadBuildpackJobArgsForCall(i int) ccv3.JobURL {
	fake.pollUploadBuildpackJobMutex.RLock()
	defer fake.pollUploadBuildpackJobMutex.RUnlock()
	argsForCall := fake.pollUploadBuildpackJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCreateBuildpackActor) PollUploadBuildpackJobReturns(result1 v7action.Warnings, result2 error) {
	fake.pollUploadBuildpackJobMutex.Lock()
	defer fake.pollUploadBuildpackJobMutex.Unlock()
	fake.PollUploadBuildpackJobStub = nil
	fake.pollUploadBuildpackJobReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateBuildpackActor) PollUploadBuildpackJobReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.pollUploadBuildpackJobMutex.Lock()
	defer fake.pollUploadBuildpackJobMutex.Unlock()
	fake.PollUploadBuildpackJobStub = nil
	if fake.pollUploadBuildpackJobReturnsOnCall == nil {
		fake.pollUploadBuildpackJobReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.pollUploadBuildpackJobReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateBuildpackActor) PrepareBuildpackBits(arg1 string, arg2 string, arg3 v7action.Downloader) (string, error) {
	fake.prepareBuildpackBitsMutex.Lock()
	ret, specificReturn := fake.prepareBuildpackBitsReturnsOnCall[len(fake.prepareBuildpackBitsArgsForCall)]
	fake.prepareBuildpackBitsArgsForCall = append(fake.prepareBuildpackBitsArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 v7action.Downloader
	}{arg1, arg2, arg3})
	fake.recordInvocation("PrepareBuildpackBits", []interface{}{arg1, arg2, arg3})
	fake.prepareBuildpackBitsMutex.Unlock()
	if fake.PrepareBuildpackBitsStub != nil {
		return fake.PrepareBuildpackBitsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.prepareBuildpackBitsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreateBuildpackActor) PrepareBuildpackBitsCallCount() int {
	fake.prepareBuildpackBitsMutex.RLock()
	defer fake.prepareBuildpackBitsMutex.RUnlock()
	return len(fake.prepareBuildpackBitsArgsForCall)
}

func (fake *FakeCreateBuildpackActor) PrepareBuildpackBitsCalls(stub func(string, string, v7action.Downloader) (string, error)) {
	fake.prepareBuildpackBitsMutex.Lock()
	defer fake.prepareBuildpackBitsMutex.Unlock()
	fake.PrepareBuildpackBitsStub = stub
}

func (fake *FakeCreateBuildpackActor) PrepareBuildpackBitsArgsForCall(i int) (string, string, v7action.Downloader) {
	fake.prepareBuildpackBitsMutex.RLock()
	defer fake.prepareBuildpackBitsMutex.RUnlock()
	argsForCall := fake.prepareBuildpackBitsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCreateBuildpackActor) PrepareBuildpackBitsReturns(result1 string, result2 error) {
	fake.prepareBuildpackBitsMutex.Lock()
	defer fake.prepareBuildpackBitsMutex.Unlock()
	fake.PrepareBuildpackBitsStub = nil
	fake.prepareBuildpackBitsReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateBuildpackActor) PrepareBuildpackBitsReturnsOnCall(i int, result1 string, result2 error) {
	fake.prepareBuildpackBitsMutex.Lock()
	defer fake.prepareBuildpackBitsMutex.Unlock()
	fake.PrepareBuildpackBitsStub = nil
	if fake.prepareBuildpackBitsReturnsOnCall == nil {
		fake.prepareBuildpackBitsReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.prepareBuildpackBitsReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateBuildpackActor) UploadBuildpack(arg1 string, arg2 string, arg3 v7action.SimpleProgressBar) (ccv3.JobURL, v7action.Warnings, error) {
	fake.uploadBuildpackMutex.Lock()
	ret, specificReturn := fake.uploadBuildpackReturnsOnCall[len(fake.uploadBuildpackArgsForCall)]
	fake.uploadBuildpackArgsForCall = append(fake.uploadBuildpackArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 v7action.SimpleProgressBar
	}{arg1, arg2, arg3})
	fake.recordInvocation("UploadBuildpack", []interface{}{arg1, arg2, arg3})
	fake.uploadBuildpackMutex.Unlock()
	if fake.UploadBuildpackStub != nil {
		return fake.UploadBuildpackStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.uploadBuildpackReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCreateBuildpackActor) UploadBuildpackCallCount() int {
	fake.uploadBuildpackMutex.RLock()
	defer fake.uploadBuildpackMutex.RUnlock()
	return len(fake.uploadBuildpackArgsForCall)
}

func (fake *FakeCreateBuildpackActor) UploadBuildpackCalls(stub func(string, string, v7action.SimpleProgressBar) (ccv3.JobURL, v7action.Warnings, error)) {
	fake.uploadBuildpackMutex.Lock()
	defer fake.uploadBuildpackMutex.Unlock()
	fake.UploadBuildpackStub = stub
}

func (fake *FakeCreateBuildpackActor) UploadBuildpackArgsForCall(i int) (string, string, v7action.SimpleProgressBar) {
	fake.uploadBuildpackMutex.RLock()
	defer fake.uploadBuildpackMutex.RUnlock()
	argsForCall := fake.uploadBuildpackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCreateBuildpackActor) UploadBuildpackReturns(result1 ccv3.JobURL, result2 v7action.Warnings, result3 error) {
	fake.uploadBuildpackMutex.Lock()
	defer fake.uploadBuildpackMutex.Unlock()
	fake.UploadBuildpackStub = nil
	fake.uploadBuildpackReturns = struct {
		result1 ccv3.JobURL
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateBuildpackActor) UploadBuildpackReturnsOnCall(i int, result1 ccv3.JobURL, result2 v7action.Warnings, result3 error) {
	fake.uploadBuildpackMutex.Lock()
	defer fake.uploadBuildpackMutex.Unlock()
	fake.UploadBuildpackStub = nil
	if fake.uploadBuildpackReturnsOnCall == nil {
		fake.uploadBuildpackReturnsOnCall = make(map[int]struct {
			result1 ccv3.JobURL
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.uploadBuildpackReturnsOnCall[i] = struct {
		result1 ccv3.JobURL
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateBuildpackActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createBuildpackMutex.RLock()
	defer fake.createBuildpackMutex.RUnlock()
	fake.pollUploadBuildpackJobMutex.RLock()
	defer fake.pollUploadBuildpackJobMutex.RUnlock()
	fake.prepareBuildpackBitsMutex.RLock()
	defer fake.prepareBuildpackBitsMutex.RUnlock()
	fake.uploadBuildpackMutex.RLock()
	defer fake.uploadBuildpackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreateBuildpackActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.CreateBuildpackActor = new(FakeCreateBuildpackActor)
