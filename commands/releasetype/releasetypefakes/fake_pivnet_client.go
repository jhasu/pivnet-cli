// Code generated by counterfeiter. DO NOT EDIT.
package releasetypefakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v7"
	"github.com/pivotal-cf/pivnet-cli/v3/commands/releasetype"
)

type FakePivnetClient struct {
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
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
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

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
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

var _ releasetype.PivnetClient = new(FakePivnetClient)
