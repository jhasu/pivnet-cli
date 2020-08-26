// Code generated by counterfeiter. DO NOT EDIT.
package loginfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/commands/login"
)

type FakeRCHandler struct {
	SaveProfileStub        func(string, string, string, string, int64) error
	saveProfileMutex       sync.RWMutex
	saveProfileArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 int64
	}
	saveProfileReturns struct {
		result1 error
	}
	saveProfileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRCHandler) SaveProfile(arg1 string, arg2 string, arg3 string, arg4 string, arg5 int64) error {
	fake.saveProfileMutex.Lock()
	ret, specificReturn := fake.saveProfileReturnsOnCall[len(fake.saveProfileArgsForCall)]
	fake.saveProfileArgsForCall = append(fake.saveProfileArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 int64
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("SaveProfile", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.saveProfileMutex.Unlock()
	if fake.SaveProfileStub != nil {
		return fake.SaveProfileStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.saveProfileReturns
	return fakeReturns.result1
}

func (fake *FakeRCHandler) SaveProfileCallCount() int {
	fake.saveProfileMutex.RLock()
	defer fake.saveProfileMutex.RUnlock()
	return len(fake.saveProfileArgsForCall)
}

func (fake *FakeRCHandler) SaveProfileCalls(stub func(string, string, string, string, int64) error) {
	fake.saveProfileMutex.Lock()
	defer fake.saveProfileMutex.Unlock()
	fake.SaveProfileStub = stub
}

func (fake *FakeRCHandler) SaveProfileArgsForCall(i int) (string, string, string, string, int64) {
	fake.saveProfileMutex.RLock()
	defer fake.saveProfileMutex.RUnlock()
	argsForCall := fake.saveProfileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeRCHandler) SaveProfileReturns(result1 error) {
	fake.saveProfileMutex.Lock()
	defer fake.saveProfileMutex.Unlock()
	fake.SaveProfileStub = nil
	fake.saveProfileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRCHandler) SaveProfileReturnsOnCall(i int, result1 error) {
	fake.saveProfileMutex.Lock()
	defer fake.saveProfileMutex.Unlock()
	fake.SaveProfileStub = nil
	if fake.saveProfileReturnsOnCall == nil {
		fake.saveProfileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveProfileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRCHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.saveProfileMutex.RLock()
	defer fake.saveProfileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRCHandler) recordInvocation(key string, args []interface{}) {
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

var _ login.RCHandler = new(FakeRCHandler)
