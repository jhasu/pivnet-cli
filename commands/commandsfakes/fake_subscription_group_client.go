// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeSubscriptionGroupClient struct {
	AddMemberStub        func(int, string, string) error
	addMemberMutex       sync.RWMutex
	addMemberArgsForCall []struct {
		arg1 int
		arg2 string
		arg3 string
	}
	addMemberReturns struct {
		result1 error
	}
	addMemberReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(int) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 int
	}
	getReturns struct {
		result1 error
	}
	getReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func() error
	listMutex       sync.RWMutex
	listArgsForCall []struct {
	}
	listReturns struct {
		result1 error
	}
	listReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveMemberStub        func(int, string) error
	removeMemberMutex       sync.RWMutex
	removeMemberArgsForCall []struct {
		arg1 int
		arg2 string
	}
	removeMemberReturns struct {
		result1 error
	}
	removeMemberReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSubscriptionGroupClient) AddMember(arg1 int, arg2 string, arg3 string) error {
	fake.addMemberMutex.Lock()
	ret, specificReturn := fake.addMemberReturnsOnCall[len(fake.addMemberArgsForCall)]
	fake.addMemberArgsForCall = append(fake.addMemberArgsForCall, struct {
		arg1 int
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("AddMember", []interface{}{arg1, arg2, arg3})
	fake.addMemberMutex.Unlock()
	if fake.AddMemberStub != nil {
		return fake.AddMemberStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addMemberReturns
	return fakeReturns.result1
}

func (fake *FakeSubscriptionGroupClient) AddMemberCallCount() int {
	fake.addMemberMutex.RLock()
	defer fake.addMemberMutex.RUnlock()
	return len(fake.addMemberArgsForCall)
}

func (fake *FakeSubscriptionGroupClient) AddMemberCalls(stub func(int, string, string) error) {
	fake.addMemberMutex.Lock()
	defer fake.addMemberMutex.Unlock()
	fake.AddMemberStub = stub
}

func (fake *FakeSubscriptionGroupClient) AddMemberArgsForCall(i int) (int, string, string) {
	fake.addMemberMutex.RLock()
	defer fake.addMemberMutex.RUnlock()
	argsForCall := fake.addMemberArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSubscriptionGroupClient) AddMemberReturns(result1 error) {
	fake.addMemberMutex.Lock()
	defer fake.addMemberMutex.Unlock()
	fake.AddMemberStub = nil
	fake.addMemberReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubscriptionGroupClient) AddMemberReturnsOnCall(i int, result1 error) {
	fake.addMemberMutex.Lock()
	defer fake.addMemberMutex.Unlock()
	fake.AddMemberStub = nil
	if fake.addMemberReturnsOnCall == nil {
		fake.addMemberReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addMemberReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubscriptionGroupClient) Get(arg1 int) error {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1
}

func (fake *FakeSubscriptionGroupClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeSubscriptionGroupClient) GetCalls(stub func(int) error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeSubscriptionGroupClient) GetArgsForCall(i int) int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSubscriptionGroupClient) GetReturns(result1 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubscriptionGroupClient) GetReturnsOnCall(i int, result1 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubscriptionGroupClient) List() error {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
	}{})
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1
}

func (fake *FakeSubscriptionGroupClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeSubscriptionGroupClient) ListCalls(stub func() error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeSubscriptionGroupClient) ListReturns(result1 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubscriptionGroupClient) ListReturnsOnCall(i int, result1 error) {
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

func (fake *FakeSubscriptionGroupClient) RemoveMember(arg1 int, arg2 string) error {
	fake.removeMemberMutex.Lock()
	ret, specificReturn := fake.removeMemberReturnsOnCall[len(fake.removeMemberArgsForCall)]
	fake.removeMemberArgsForCall = append(fake.removeMemberArgsForCall, struct {
		arg1 int
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("RemoveMember", []interface{}{arg1, arg2})
	fake.removeMemberMutex.Unlock()
	if fake.RemoveMemberStub != nil {
		return fake.RemoveMemberStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeMemberReturns
	return fakeReturns.result1
}

func (fake *FakeSubscriptionGroupClient) RemoveMemberCallCount() int {
	fake.removeMemberMutex.RLock()
	defer fake.removeMemberMutex.RUnlock()
	return len(fake.removeMemberArgsForCall)
}

func (fake *FakeSubscriptionGroupClient) RemoveMemberCalls(stub func(int, string) error) {
	fake.removeMemberMutex.Lock()
	defer fake.removeMemberMutex.Unlock()
	fake.RemoveMemberStub = stub
}

func (fake *FakeSubscriptionGroupClient) RemoveMemberArgsForCall(i int) (int, string) {
	fake.removeMemberMutex.RLock()
	defer fake.removeMemberMutex.RUnlock()
	argsForCall := fake.removeMemberArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSubscriptionGroupClient) RemoveMemberReturns(result1 error) {
	fake.removeMemberMutex.Lock()
	defer fake.removeMemberMutex.Unlock()
	fake.RemoveMemberStub = nil
	fake.removeMemberReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubscriptionGroupClient) RemoveMemberReturnsOnCall(i int, result1 error) {
	fake.removeMemberMutex.Lock()
	defer fake.removeMemberMutex.Unlock()
	fake.RemoveMemberStub = nil
	if fake.removeMemberReturnsOnCall == nil {
		fake.removeMemberReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeMemberReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubscriptionGroupClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMemberMutex.RLock()
	defer fake.addMemberMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.removeMemberMutex.RLock()
	defer fake.removeMemberMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSubscriptionGroupClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.SubscriptionGroupClient = new(FakeSubscriptionGroupClient)
