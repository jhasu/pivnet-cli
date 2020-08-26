// Code generated by counterfeiter. DO NOT EDIT.
package dependencyspecifierfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v6"
	"github.com/pivotal-cf/pivnet-cli/commands/dependencyspecifier"
)

type FakePivnetClient struct {
	CreateDependencySpecifierStub        func(string, int, string, string) (pivnet.DependencySpecifier, error)
	createDependencySpecifierMutex       sync.RWMutex
	createDependencySpecifierArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 string
		arg4 string
	}
	createDependencySpecifierReturns struct {
		result1 pivnet.DependencySpecifier
		result2 error
	}
	createDependencySpecifierReturnsOnCall map[int]struct {
		result1 pivnet.DependencySpecifier
		result2 error
	}
	DeleteDependencySpecifierStub        func(string, int, int) error
	deleteDependencySpecifierMutex       sync.RWMutex
	deleteDependencySpecifierArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	deleteDependencySpecifierReturns struct {
		result1 error
	}
	deleteDependencySpecifierReturnsOnCall map[int]struct {
		result1 error
	}
	DependencySpecifierStub        func(string, int, int) (pivnet.DependencySpecifier, error)
	dependencySpecifierMutex       sync.RWMutex
	dependencySpecifierArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	dependencySpecifierReturns struct {
		result1 pivnet.DependencySpecifier
		result2 error
	}
	dependencySpecifierReturnsOnCall map[int]struct {
		result1 pivnet.DependencySpecifier
		result2 error
	}
	DependencySpecifiersStub        func(string, int) ([]pivnet.DependencySpecifier, error)
	dependencySpecifiersMutex       sync.RWMutex
	dependencySpecifiersArgsForCall []struct {
		arg1 string
		arg2 int
	}
	dependencySpecifiersReturns struct {
		result1 []pivnet.DependencySpecifier
		result2 error
	}
	dependencySpecifiersReturnsOnCall map[int]struct {
		result1 []pivnet.DependencySpecifier
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
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) CreateDependencySpecifier(arg1 string, arg2 int, arg3 string, arg4 string) (pivnet.DependencySpecifier, error) {
	fake.createDependencySpecifierMutex.Lock()
	ret, specificReturn := fake.createDependencySpecifierReturnsOnCall[len(fake.createDependencySpecifierArgsForCall)]
	fake.createDependencySpecifierArgsForCall = append(fake.createDependencySpecifierArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CreateDependencySpecifier", []interface{}{arg1, arg2, arg3, arg4})
	fake.createDependencySpecifierMutex.Unlock()
	if fake.CreateDependencySpecifierStub != nil {
		return fake.CreateDependencySpecifierStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createDependencySpecifierReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) CreateDependencySpecifierCallCount() int {
	fake.createDependencySpecifierMutex.RLock()
	defer fake.createDependencySpecifierMutex.RUnlock()
	return len(fake.createDependencySpecifierArgsForCall)
}

func (fake *FakePivnetClient) CreateDependencySpecifierCalls(stub func(string, int, string, string) (pivnet.DependencySpecifier, error)) {
	fake.createDependencySpecifierMutex.Lock()
	defer fake.createDependencySpecifierMutex.Unlock()
	fake.CreateDependencySpecifierStub = stub
}

func (fake *FakePivnetClient) CreateDependencySpecifierArgsForCall(i int) (string, int, string, string) {
	fake.createDependencySpecifierMutex.RLock()
	defer fake.createDependencySpecifierMutex.RUnlock()
	argsForCall := fake.createDependencySpecifierArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakePivnetClient) CreateDependencySpecifierReturns(result1 pivnet.DependencySpecifier, result2 error) {
	fake.createDependencySpecifierMutex.Lock()
	defer fake.createDependencySpecifierMutex.Unlock()
	fake.CreateDependencySpecifierStub = nil
	fake.createDependencySpecifierReturns = struct {
		result1 pivnet.DependencySpecifier
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) CreateDependencySpecifierReturnsOnCall(i int, result1 pivnet.DependencySpecifier, result2 error) {
	fake.createDependencySpecifierMutex.Lock()
	defer fake.createDependencySpecifierMutex.Unlock()
	fake.CreateDependencySpecifierStub = nil
	if fake.createDependencySpecifierReturnsOnCall == nil {
		fake.createDependencySpecifierReturnsOnCall = make(map[int]struct {
			result1 pivnet.DependencySpecifier
			result2 error
		})
	}
	fake.createDependencySpecifierReturnsOnCall[i] = struct {
		result1 pivnet.DependencySpecifier
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DeleteDependencySpecifier(arg1 string, arg2 int, arg3 int) error {
	fake.deleteDependencySpecifierMutex.Lock()
	ret, specificReturn := fake.deleteDependencySpecifierReturnsOnCall[len(fake.deleteDependencySpecifierArgsForCall)]
	fake.deleteDependencySpecifierArgsForCall = append(fake.deleteDependencySpecifierArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("DeleteDependencySpecifier", []interface{}{arg1, arg2, arg3})
	fake.deleteDependencySpecifierMutex.Unlock()
	if fake.DeleteDependencySpecifierStub != nil {
		return fake.DeleteDependencySpecifierStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteDependencySpecifierReturns
	return fakeReturns.result1
}

func (fake *FakePivnetClient) DeleteDependencySpecifierCallCount() int {
	fake.deleteDependencySpecifierMutex.RLock()
	defer fake.deleteDependencySpecifierMutex.RUnlock()
	return len(fake.deleteDependencySpecifierArgsForCall)
}

func (fake *FakePivnetClient) DeleteDependencySpecifierCalls(stub func(string, int, int) error) {
	fake.deleteDependencySpecifierMutex.Lock()
	defer fake.deleteDependencySpecifierMutex.Unlock()
	fake.DeleteDependencySpecifierStub = stub
}

func (fake *FakePivnetClient) DeleteDependencySpecifierArgsForCall(i int) (string, int, int) {
	fake.deleteDependencySpecifierMutex.RLock()
	defer fake.deleteDependencySpecifierMutex.RUnlock()
	argsForCall := fake.deleteDependencySpecifierArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePivnetClient) DeleteDependencySpecifierReturns(result1 error) {
	fake.deleteDependencySpecifierMutex.Lock()
	defer fake.deleteDependencySpecifierMutex.Unlock()
	fake.DeleteDependencySpecifierStub = nil
	fake.deleteDependencySpecifierReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) DeleteDependencySpecifierReturnsOnCall(i int, result1 error) {
	fake.deleteDependencySpecifierMutex.Lock()
	defer fake.deleteDependencySpecifierMutex.Unlock()
	fake.DeleteDependencySpecifierStub = nil
	if fake.deleteDependencySpecifierReturnsOnCall == nil {
		fake.deleteDependencySpecifierReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteDependencySpecifierReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) DependencySpecifier(arg1 string, arg2 int, arg3 int) (pivnet.DependencySpecifier, error) {
	fake.dependencySpecifierMutex.Lock()
	ret, specificReturn := fake.dependencySpecifierReturnsOnCall[len(fake.dependencySpecifierArgsForCall)]
	fake.dependencySpecifierArgsForCall = append(fake.dependencySpecifierArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("DependencySpecifier", []interface{}{arg1, arg2, arg3})
	fake.dependencySpecifierMutex.Unlock()
	if fake.DependencySpecifierStub != nil {
		return fake.DependencySpecifierStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.dependencySpecifierReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) DependencySpecifierCallCount() int {
	fake.dependencySpecifierMutex.RLock()
	defer fake.dependencySpecifierMutex.RUnlock()
	return len(fake.dependencySpecifierArgsForCall)
}

func (fake *FakePivnetClient) DependencySpecifierCalls(stub func(string, int, int) (pivnet.DependencySpecifier, error)) {
	fake.dependencySpecifierMutex.Lock()
	defer fake.dependencySpecifierMutex.Unlock()
	fake.DependencySpecifierStub = stub
}

func (fake *FakePivnetClient) DependencySpecifierArgsForCall(i int) (string, int, int) {
	fake.dependencySpecifierMutex.RLock()
	defer fake.dependencySpecifierMutex.RUnlock()
	argsForCall := fake.dependencySpecifierArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePivnetClient) DependencySpecifierReturns(result1 pivnet.DependencySpecifier, result2 error) {
	fake.dependencySpecifierMutex.Lock()
	defer fake.dependencySpecifierMutex.Unlock()
	fake.DependencySpecifierStub = nil
	fake.dependencySpecifierReturns = struct {
		result1 pivnet.DependencySpecifier
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DependencySpecifierReturnsOnCall(i int, result1 pivnet.DependencySpecifier, result2 error) {
	fake.dependencySpecifierMutex.Lock()
	defer fake.dependencySpecifierMutex.Unlock()
	fake.DependencySpecifierStub = nil
	if fake.dependencySpecifierReturnsOnCall == nil {
		fake.dependencySpecifierReturnsOnCall = make(map[int]struct {
			result1 pivnet.DependencySpecifier
			result2 error
		})
	}
	fake.dependencySpecifierReturnsOnCall[i] = struct {
		result1 pivnet.DependencySpecifier
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DependencySpecifiers(arg1 string, arg2 int) ([]pivnet.DependencySpecifier, error) {
	fake.dependencySpecifiersMutex.Lock()
	ret, specificReturn := fake.dependencySpecifiersReturnsOnCall[len(fake.dependencySpecifiersArgsForCall)]
	fake.dependencySpecifiersArgsForCall = append(fake.dependencySpecifiersArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("DependencySpecifiers", []interface{}{arg1, arg2})
	fake.dependencySpecifiersMutex.Unlock()
	if fake.DependencySpecifiersStub != nil {
		return fake.DependencySpecifiersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.dependencySpecifiersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) DependencySpecifiersCallCount() int {
	fake.dependencySpecifiersMutex.RLock()
	defer fake.dependencySpecifiersMutex.RUnlock()
	return len(fake.dependencySpecifiersArgsForCall)
}

func (fake *FakePivnetClient) DependencySpecifiersCalls(stub func(string, int) ([]pivnet.DependencySpecifier, error)) {
	fake.dependencySpecifiersMutex.Lock()
	defer fake.dependencySpecifiersMutex.Unlock()
	fake.DependencySpecifiersStub = stub
}

func (fake *FakePivnetClient) DependencySpecifiersArgsForCall(i int) (string, int) {
	fake.dependencySpecifiersMutex.RLock()
	defer fake.dependencySpecifiersMutex.RUnlock()
	argsForCall := fake.dependencySpecifiersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) DependencySpecifiersReturns(result1 []pivnet.DependencySpecifier, result2 error) {
	fake.dependencySpecifiersMutex.Lock()
	defer fake.dependencySpecifiersMutex.Unlock()
	fake.DependencySpecifiersStub = nil
	fake.dependencySpecifiersReturns = struct {
		result1 []pivnet.DependencySpecifier
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DependencySpecifiersReturnsOnCall(i int, result1 []pivnet.DependencySpecifier, result2 error) {
	fake.dependencySpecifiersMutex.Lock()
	defer fake.dependencySpecifiersMutex.Unlock()
	fake.DependencySpecifiersStub = nil
	if fake.dependencySpecifiersReturnsOnCall == nil {
		fake.dependencySpecifiersReturnsOnCall = make(map[int]struct {
			result1 []pivnet.DependencySpecifier
			result2 error
		})
	}
	fake.dependencySpecifiersReturnsOnCall[i] = struct {
		result1 []pivnet.DependencySpecifier
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

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createDependencySpecifierMutex.RLock()
	defer fake.createDependencySpecifierMutex.RUnlock()
	fake.deleteDependencySpecifierMutex.RLock()
	defer fake.deleteDependencySpecifierMutex.RUnlock()
	fake.dependencySpecifierMutex.RLock()
	defer fake.dependencySpecifierMutex.RUnlock()
	fake.dependencySpecifiersMutex.RLock()
	defer fake.dependencySpecifiersMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
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

var _ dependencyspecifier.PivnetClient = new(FakePivnetClient)
