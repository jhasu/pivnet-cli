// Code generated by counterfeiter. DO NOT EDIT.
package pivnetversionsfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v7"
	"github.com/pivotal-cf/pivnet-cli/v3/commands/pivnetversions"
)

type FakePivnetClient struct {
	PivnetVersionsStub        func() (pivnet.PivnetVersions, error)
	pivnetVersionsMutex       sync.RWMutex
	pivnetVersionsArgsForCall []struct {
	}
	pivnetVersionsReturns struct {
		result1 pivnet.PivnetVersions
		result2 error
	}
	pivnetVersionsReturnsOnCall map[int]struct {
		result1 pivnet.PivnetVersions
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) PivnetVersions() (pivnet.PivnetVersions, error) {
	fake.pivnetVersionsMutex.Lock()
	ret, specificReturn := fake.pivnetVersionsReturnsOnCall[len(fake.pivnetVersionsArgsForCall)]
	fake.pivnetVersionsArgsForCall = append(fake.pivnetVersionsArgsForCall, struct {
	}{})
	fake.recordInvocation("PivnetVersions", []interface{}{})
	fake.pivnetVersionsMutex.Unlock()
	if fake.PivnetVersionsStub != nil {
		return fake.PivnetVersionsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.pivnetVersionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) PivnetVersionsCallCount() int {
	fake.pivnetVersionsMutex.RLock()
	defer fake.pivnetVersionsMutex.RUnlock()
	return len(fake.pivnetVersionsArgsForCall)
}

func (fake *FakePivnetClient) PivnetVersionsCalls(stub func() (pivnet.PivnetVersions, error)) {
	fake.pivnetVersionsMutex.Lock()
	defer fake.pivnetVersionsMutex.Unlock()
	fake.PivnetVersionsStub = stub
}

func (fake *FakePivnetClient) PivnetVersionsReturns(result1 pivnet.PivnetVersions, result2 error) {
	fake.pivnetVersionsMutex.Lock()
	defer fake.pivnetVersionsMutex.Unlock()
	fake.PivnetVersionsStub = nil
	fake.pivnetVersionsReturns = struct {
		result1 pivnet.PivnetVersions
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) PivnetVersionsReturnsOnCall(i int, result1 pivnet.PivnetVersions, result2 error) {
	fake.pivnetVersionsMutex.Lock()
	defer fake.pivnetVersionsMutex.Unlock()
	fake.PivnetVersionsStub = nil
	if fake.pivnetVersionsReturnsOnCall == nil {
		fake.pivnetVersionsReturnsOnCall = make(map[int]struct {
			result1 pivnet.PivnetVersions
			result2 error
		})
	}
	fake.pivnetVersionsReturnsOnCall[i] = struct {
		result1 pivnet.PivnetVersions
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pivnetVersionsMutex.RLock()
	defer fake.pivnetVersionsMutex.RUnlock()
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

var _ pivnetversions.PivnetClient = new(FakePivnetClient)
