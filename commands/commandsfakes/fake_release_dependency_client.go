// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/v3/commands"
)

type FakeReleaseDependencyClient struct {
	AddStub        func(string, string, string, string) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	addReturns struct {
		result1 error
	}
	addReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(string, string) error
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listReturns struct {
		result1 error
	}
	listReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveStub        func(string, string, string, string) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseDependencyClient) Add(arg1 string, arg2 string, arg3 string, arg4 string) error {
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Add", []interface{}{arg1, arg2, arg3, arg4})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addReturns
	return fakeReturns.result1
}

func (fake *FakeReleaseDependencyClient) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeReleaseDependencyClient) AddCalls(stub func(string, string, string, string) error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *FakeReleaseDependencyClient) AddArgsForCall(i int) (string, string, string, string) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeReleaseDependencyClient) AddReturns(result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseDependencyClient) AddReturnsOnCall(i int, result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseDependencyClient) List(arg1 string, arg2 string) error {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("List", []interface{}{arg1, arg2})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1
}

func (fake *FakeReleaseDependencyClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeReleaseDependencyClient) ListCalls(stub func(string, string) error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeReleaseDependencyClient) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeReleaseDependencyClient) ListReturns(result1 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseDependencyClient) ListReturnsOnCall(i int, result1 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseDependencyClient) Remove(arg1 string, arg2 string, arg3 string, arg4 string) error {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Remove", []interface{}{arg1, arg2, arg3, arg4})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeReturns
	return fakeReturns.result1
}

func (fake *FakeReleaseDependencyClient) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeReleaseDependencyClient) RemoveCalls(stub func(string, string, string, string) error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = stub
}

func (fake *FakeReleaseDependencyClient) RemoveArgsForCall(i int) (string, string, string, string) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	argsForCall := fake.removeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeReleaseDependencyClient) RemoveReturns(result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseDependencyClient) RemoveReturnsOnCall(i int, result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseDependencyClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReleaseDependencyClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.ReleaseDependencyClient = new(FakeReleaseDependencyClient)
