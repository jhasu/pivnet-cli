// Code generated by counterfeiter. DO NOT EDIT.
package releasefakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v4"
	"github.com/pivotal-cf/pivnet-cli/commands/release"
)

type FakePivnetClient struct {
	CreateReleaseStub        func(pivnet.CreateReleaseConfig) (pivnet.Release, error)
	createReleaseMutex       sync.RWMutex
	createReleaseArgsForCall []struct {
		arg1 pivnet.CreateReleaseConfig
	}
	createReleaseReturns struct {
		result1 pivnet.Release
		result2 error
	}
	createReleaseReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	DeleteReleaseStub        func(string, pivnet.Release) error
	deleteReleaseMutex       sync.RWMutex
	deleteReleaseArgsForCall []struct {
		arg1 string
		arg2 pivnet.Release
	}
	deleteReleaseReturns struct {
		result1 error
	}
	deleteReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	EULAsStub        func() ([]pivnet.EULA, error)
	eULAsMutex       sync.RWMutex
	eULAsArgsForCall []struct {
	}
	eULAsReturns struct {
		result1 []pivnet.EULA
		result2 error
	}
	eULAsReturnsOnCall map[int]struct {
		result1 []pivnet.EULA
		result2 error
	}
	ReleaseForVersionStub        func(string, string) (pivnet.Release, error)
	releaseForVersionMutex       sync.RWMutex
	releaseForVersionArgsForCall []struct {
		arg1 string
		arg2 string
	}
	releaseForVersionReturns struct {
		result1 pivnet.Release
		result2 error
	}
	releaseForVersionReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	ReleaseTypesStub        func() ([]pivnet.ReleaseType, error)
	releaseTypesMutex       sync.RWMutex
	releaseTypesArgsForCall []struct {
	}
	releaseTypesReturns struct {
		result1 []pivnet.ReleaseType
		result2 error
	}
	releaseTypesReturnsOnCall map[int]struct {
		result1 []pivnet.ReleaseType
		result2 error
	}
	ReleasesForProductSlugStub        func(string, ...pivnet.QueryParameter) ([]pivnet.Release, error)
	releasesForProductSlugMutex       sync.RWMutex
	releasesForProductSlugArgsForCall []struct {
		arg1 string
		arg2 []pivnet.QueryParameter
	}
	releasesForProductSlugReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	releasesForProductSlugReturnsOnCall map[int]struct {
		result1 []pivnet.Release
		result2 error
	}
	UpdateReleaseStub        func(string, pivnet.Release) (pivnet.Release, error)
	updateReleaseMutex       sync.RWMutex
	updateReleaseArgsForCall []struct {
		arg1 string
		arg2 pivnet.Release
	}
	updateReleaseReturns struct {
		result1 pivnet.Release
		result2 error
	}
	updateReleaseReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) CreateRelease(arg1 pivnet.CreateReleaseConfig) (pivnet.Release, error) {
	fake.createReleaseMutex.Lock()
	ret, specificReturn := fake.createReleaseReturnsOnCall[len(fake.createReleaseArgsForCall)]
	fake.createReleaseArgsForCall = append(fake.createReleaseArgsForCall, struct {
		arg1 pivnet.CreateReleaseConfig
	}{arg1})
	fake.recordInvocation("CreateRelease", []interface{}{arg1})
	fake.createReleaseMutex.Unlock()
	if fake.CreateReleaseStub != nil {
		return fake.CreateReleaseStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReleaseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) CreateReleaseCallCount() int {
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	return len(fake.createReleaseArgsForCall)
}

func (fake *FakePivnetClient) CreateReleaseCalls(stub func(pivnet.CreateReleaseConfig) (pivnet.Release, error)) {
	fake.createReleaseMutex.Lock()
	defer fake.createReleaseMutex.Unlock()
	fake.CreateReleaseStub = stub
}

func (fake *FakePivnetClient) CreateReleaseArgsForCall(i int) pivnet.CreateReleaseConfig {
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	argsForCall := fake.createReleaseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePivnetClient) CreateReleaseReturns(result1 pivnet.Release, result2 error) {
	fake.createReleaseMutex.Lock()
	defer fake.createReleaseMutex.Unlock()
	fake.CreateReleaseStub = nil
	fake.createReleaseReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) CreateReleaseReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.createReleaseMutex.Lock()
	defer fake.createReleaseMutex.Unlock()
	fake.CreateReleaseStub = nil
	if fake.createReleaseReturnsOnCall == nil {
		fake.createReleaseReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.createReleaseReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DeleteRelease(arg1 string, arg2 pivnet.Release) error {
	fake.deleteReleaseMutex.Lock()
	ret, specificReturn := fake.deleteReleaseReturnsOnCall[len(fake.deleteReleaseArgsForCall)]
	fake.deleteReleaseArgsForCall = append(fake.deleteReleaseArgsForCall, struct {
		arg1 string
		arg2 pivnet.Release
	}{arg1, arg2})
	fake.recordInvocation("DeleteRelease", []interface{}{arg1, arg2})
	fake.deleteReleaseMutex.Unlock()
	if fake.DeleteReleaseStub != nil {
		return fake.DeleteReleaseStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReleaseReturns
	return fakeReturns.result1
}

func (fake *FakePivnetClient) DeleteReleaseCallCount() int {
	fake.deleteReleaseMutex.RLock()
	defer fake.deleteReleaseMutex.RUnlock()
	return len(fake.deleteReleaseArgsForCall)
}

func (fake *FakePivnetClient) DeleteReleaseCalls(stub func(string, pivnet.Release) error) {
	fake.deleteReleaseMutex.Lock()
	defer fake.deleteReleaseMutex.Unlock()
	fake.DeleteReleaseStub = stub
}

func (fake *FakePivnetClient) DeleteReleaseArgsForCall(i int) (string, pivnet.Release) {
	fake.deleteReleaseMutex.RLock()
	defer fake.deleteReleaseMutex.RUnlock()
	argsForCall := fake.deleteReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) DeleteReleaseReturns(result1 error) {
	fake.deleteReleaseMutex.Lock()
	defer fake.deleteReleaseMutex.Unlock()
	fake.DeleteReleaseStub = nil
	fake.deleteReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) DeleteReleaseReturnsOnCall(i int, result1 error) {
	fake.deleteReleaseMutex.Lock()
	defer fake.deleteReleaseMutex.Unlock()
	fake.DeleteReleaseStub = nil
	if fake.deleteReleaseReturnsOnCall == nil {
		fake.deleteReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) EULAs() ([]pivnet.EULA, error) {
	fake.eULAsMutex.Lock()
	ret, specificReturn := fake.eULAsReturnsOnCall[len(fake.eULAsArgsForCall)]
	fake.eULAsArgsForCall = append(fake.eULAsArgsForCall, struct {
	}{})
	fake.recordInvocation("EULAs", []interface{}{})
	fake.eULAsMutex.Unlock()
	if fake.EULAsStub != nil {
		return fake.EULAsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.eULAsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) EULAsCallCount() int {
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	return len(fake.eULAsArgsForCall)
}

func (fake *FakePivnetClient) EULAsCalls(stub func() ([]pivnet.EULA, error)) {
	fake.eULAsMutex.Lock()
	defer fake.eULAsMutex.Unlock()
	fake.EULAsStub = stub
}

func (fake *FakePivnetClient) EULAsReturns(result1 []pivnet.EULA, result2 error) {
	fake.eULAsMutex.Lock()
	defer fake.eULAsMutex.Unlock()
	fake.EULAsStub = nil
	fake.eULAsReturns = struct {
		result1 []pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) EULAsReturnsOnCall(i int, result1 []pivnet.EULA, result2 error) {
	fake.eULAsMutex.Lock()
	defer fake.eULAsMutex.Unlock()
	fake.EULAsStub = nil
	if fake.eULAsReturnsOnCall == nil {
		fake.eULAsReturnsOnCall = make(map[int]struct {
			result1 []pivnet.EULA
			result2 error
		})
	}
	fake.eULAsReturnsOnCall[i] = struct {
		result1 []pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseForVersion(arg1 string, arg2 string) (pivnet.Release, error) {
	fake.releaseForVersionMutex.Lock()
	ret, specificReturn := fake.releaseForVersionReturnsOnCall[len(fake.releaseForVersionArgsForCall)]
	fake.releaseForVersionArgsForCall = append(fake.releaseForVersionArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ReleaseForVersion", []interface{}{arg1, arg2})
	fake.releaseForVersionMutex.Unlock()
	if fake.ReleaseForVersionStub != nil {
		return fake.ReleaseForVersionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releaseForVersionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ReleaseForVersionCallCount() int {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return len(fake.releaseForVersionArgsForCall)
}

func (fake *FakePivnetClient) ReleaseForVersionCalls(stub func(string, string) (pivnet.Release, error)) {
	fake.releaseForVersionMutex.Lock()
	defer fake.releaseForVersionMutex.Unlock()
	fake.ReleaseForVersionStub = stub
}

func (fake *FakePivnetClient) ReleaseForVersionArgsForCall(i int) (string, string) {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	argsForCall := fake.releaseForVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) ReleaseForVersionReturns(result1 pivnet.Release, result2 error) {
	fake.releaseForVersionMutex.Lock()
	defer fake.releaseForVersionMutex.Unlock()
	fake.ReleaseForVersionStub = nil
	fake.releaseForVersionReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseForVersionReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.releaseForVersionMutex.Lock()
	defer fake.releaseForVersionMutex.Unlock()
	fake.ReleaseForVersionStub = nil
	if fake.releaseForVersionReturnsOnCall == nil {
		fake.releaseForVersionReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.releaseForVersionReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseTypes() ([]pivnet.ReleaseType, error) {
	fake.releaseTypesMutex.Lock()
	ret, specificReturn := fake.releaseTypesReturnsOnCall[len(fake.releaseTypesArgsForCall)]
	fake.releaseTypesArgsForCall = append(fake.releaseTypesArgsForCall, struct {
	}{})
	fake.recordInvocation("ReleaseTypes", []interface{}{})
	fake.releaseTypesMutex.Unlock()
	if fake.ReleaseTypesStub != nil {
		return fake.ReleaseTypesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releaseTypesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ReleaseTypesCallCount() int {
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	return len(fake.releaseTypesArgsForCall)
}

func (fake *FakePivnetClient) ReleaseTypesCalls(stub func() ([]pivnet.ReleaseType, error)) {
	fake.releaseTypesMutex.Lock()
	defer fake.releaseTypesMutex.Unlock()
	fake.ReleaseTypesStub = stub
}

func (fake *FakePivnetClient) ReleaseTypesReturns(result1 []pivnet.ReleaseType, result2 error) {
	fake.releaseTypesMutex.Lock()
	defer fake.releaseTypesMutex.Unlock()
	fake.ReleaseTypesStub = nil
	fake.releaseTypesReturns = struct {
		result1 []pivnet.ReleaseType
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseTypesReturnsOnCall(i int, result1 []pivnet.ReleaseType, result2 error) {
	fake.releaseTypesMutex.Lock()
	defer fake.releaseTypesMutex.Unlock()
	fake.ReleaseTypesStub = nil
	if fake.releaseTypesReturnsOnCall == nil {
		fake.releaseTypesReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ReleaseType
			result2 error
		})
	}
	fake.releaseTypesReturnsOnCall[i] = struct {
		result1 []pivnet.ReleaseType
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleasesForProductSlug(arg1 string, arg2 ...pivnet.QueryParameter) ([]pivnet.Release, error) {
	fake.releasesForProductSlugMutex.Lock()
	ret, specificReturn := fake.releasesForProductSlugReturnsOnCall[len(fake.releasesForProductSlugArgsForCall)]
	fake.releasesForProductSlugArgsForCall = append(fake.releasesForProductSlugArgsForCall, struct {
		arg1 string
		arg2 []pivnet.QueryParameter
	}{arg1, arg2})
	fake.recordInvocation("ReleasesForProductSlug", []interface{}{arg1, arg2})
	fake.releasesForProductSlugMutex.Unlock()
	if fake.ReleasesForProductSlugStub != nil {
		return fake.ReleasesForProductSlugStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releasesForProductSlugReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ReleasesForProductSlugCallCount() int {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return len(fake.releasesForProductSlugArgsForCall)
}

func (fake *FakePivnetClient) ReleasesForProductSlugCalls(stub func(string, ...pivnet.QueryParameter) ([]pivnet.Release, error)) {
	fake.releasesForProductSlugMutex.Lock()
	defer fake.releasesForProductSlugMutex.Unlock()
	fake.ReleasesForProductSlugStub = stub
}

func (fake *FakePivnetClient) ReleasesForProductSlugArgsForCall(i int) (string, []pivnet.QueryParameter) {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	argsForCall := fake.releasesForProductSlugArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) ReleasesForProductSlugReturns(result1 []pivnet.Release, result2 error) {
	fake.releasesForProductSlugMutex.Lock()
	defer fake.releasesForProductSlugMutex.Unlock()
	fake.ReleasesForProductSlugStub = nil
	fake.releasesForProductSlugReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleasesForProductSlugReturnsOnCall(i int, result1 []pivnet.Release, result2 error) {
	fake.releasesForProductSlugMutex.Lock()
	defer fake.releasesForProductSlugMutex.Unlock()
	fake.ReleasesForProductSlugStub = nil
	if fake.releasesForProductSlugReturnsOnCall == nil {
		fake.releasesForProductSlugReturnsOnCall = make(map[int]struct {
			result1 []pivnet.Release
			result2 error
		})
	}
	fake.releasesForProductSlugReturnsOnCall[i] = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpdateRelease(arg1 string, arg2 pivnet.Release) (pivnet.Release, error) {
	fake.updateReleaseMutex.Lock()
	ret, specificReturn := fake.updateReleaseReturnsOnCall[len(fake.updateReleaseArgsForCall)]
	fake.updateReleaseArgsForCall = append(fake.updateReleaseArgsForCall, struct {
		arg1 string
		arg2 pivnet.Release
	}{arg1, arg2})
	fake.recordInvocation("UpdateRelease", []interface{}{arg1, arg2})
	fake.updateReleaseMutex.Unlock()
	if fake.UpdateReleaseStub != nil {
		return fake.UpdateReleaseStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateReleaseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) UpdateReleaseCallCount() int {
	fake.updateReleaseMutex.RLock()
	defer fake.updateReleaseMutex.RUnlock()
	return len(fake.updateReleaseArgsForCall)
}

func (fake *FakePivnetClient) UpdateReleaseCalls(stub func(string, pivnet.Release) (pivnet.Release, error)) {
	fake.updateReleaseMutex.Lock()
	defer fake.updateReleaseMutex.Unlock()
	fake.UpdateReleaseStub = stub
}

func (fake *FakePivnetClient) UpdateReleaseArgsForCall(i int) (string, pivnet.Release) {
	fake.updateReleaseMutex.RLock()
	defer fake.updateReleaseMutex.RUnlock()
	argsForCall := fake.updateReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) UpdateReleaseReturns(result1 pivnet.Release, result2 error) {
	fake.updateReleaseMutex.Lock()
	defer fake.updateReleaseMutex.Unlock()
	fake.UpdateReleaseStub = nil
	fake.updateReleaseReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpdateReleaseReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.updateReleaseMutex.Lock()
	defer fake.updateReleaseMutex.Unlock()
	fake.UpdateReleaseStub = nil
	if fake.updateReleaseReturnsOnCall == nil {
		fake.updateReleaseReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.updateReleaseReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	fake.deleteReleaseMutex.RLock()
	defer fake.deleteReleaseMutex.RUnlock()
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	fake.updateReleaseMutex.RLock()
	defer fake.updateReleaseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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

var _ release.PivnetClient = new(FakePivnetClient)
