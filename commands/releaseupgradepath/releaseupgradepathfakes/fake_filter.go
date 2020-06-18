// Code generated by counterfeiter. DO NOT EDIT.
package releaseupgradepathfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v5"
	"github.com/pivotal-cf/pivnet-cli/commands/releaseupgradepath"
)

type FakeFilter struct {
	ReleasesByVersionStub        func([]pivnet.Release, string) ([]pivnet.Release, error)
	releasesByVersionMutex       sync.RWMutex
	releasesByVersionArgsForCall []struct {
		arg1 []pivnet.Release
		arg2 string
	}
	releasesByVersionReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	releasesByVersionReturnsOnCall map[int]struct {
		result1 []pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFilter) ReleasesByVersion(arg1 []pivnet.Release, arg2 string) ([]pivnet.Release, error) {
	var arg1Copy []pivnet.Release
	if arg1 != nil {
		arg1Copy = make([]pivnet.Release, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.releasesByVersionMutex.Lock()
	ret, specificReturn := fake.releasesByVersionReturnsOnCall[len(fake.releasesByVersionArgsForCall)]
	fake.releasesByVersionArgsForCall = append(fake.releasesByVersionArgsForCall, struct {
		arg1 []pivnet.Release
		arg2 string
	}{arg1Copy, arg2})
	fake.recordInvocation("ReleasesByVersion", []interface{}{arg1Copy, arg2})
	fake.releasesByVersionMutex.Unlock()
	if fake.ReleasesByVersionStub != nil {
		return fake.ReleasesByVersionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releasesByVersionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFilter) ReleasesByVersionCallCount() int {
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	return len(fake.releasesByVersionArgsForCall)
}

func (fake *FakeFilter) ReleasesByVersionCalls(stub func([]pivnet.Release, string) ([]pivnet.Release, error)) {
	fake.releasesByVersionMutex.Lock()
	defer fake.releasesByVersionMutex.Unlock()
	fake.ReleasesByVersionStub = stub
}

func (fake *FakeFilter) ReleasesByVersionArgsForCall(i int) ([]pivnet.Release, string) {
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	argsForCall := fake.releasesByVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFilter) ReleasesByVersionReturns(result1 []pivnet.Release, result2 error) {
	fake.releasesByVersionMutex.Lock()
	defer fake.releasesByVersionMutex.Unlock()
	fake.ReleasesByVersionStub = nil
	fake.releasesByVersionReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) ReleasesByVersionReturnsOnCall(i int, result1 []pivnet.Release, result2 error) {
	fake.releasesByVersionMutex.Lock()
	defer fake.releasesByVersionMutex.Unlock()
	fake.ReleasesByVersionStub = nil
	if fake.releasesByVersionReturnsOnCall == nil {
		fake.releasesByVersionReturnsOnCall = make(map[int]struct {
			result1 []pivnet.Release
			result2 error
		})
	}
	fake.releasesByVersionReturnsOnCall[i] = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFilter) recordInvocation(key string, args []interface{}) {
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

var _ releaseupgradepath.Filter = new(FakeFilter)
