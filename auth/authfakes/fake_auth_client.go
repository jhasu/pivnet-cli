// Code generated by counterfeiter. DO NOT EDIT.
package authfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/v3/auth"
)

type FakeAuthClient struct {
	AuthStub        func() (bool, error)
	authMutex       sync.RWMutex
	authArgsForCall []struct {
	}
	authReturns struct {
		result1 bool
		result2 error
	}
	authReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthClient) Auth() (bool, error) {
	fake.authMutex.Lock()
	ret, specificReturn := fake.authReturnsOnCall[len(fake.authArgsForCall)]
	fake.authArgsForCall = append(fake.authArgsForCall, struct {
	}{})
	fake.recordInvocation("Auth", []interface{}{})
	fake.authMutex.Unlock()
	if fake.AuthStub != nil {
		return fake.AuthStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.authReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAuthClient) AuthCallCount() int {
	fake.authMutex.RLock()
	defer fake.authMutex.RUnlock()
	return len(fake.authArgsForCall)
}

func (fake *FakeAuthClient) AuthCalls(stub func() (bool, error)) {
	fake.authMutex.Lock()
	defer fake.authMutex.Unlock()
	fake.AuthStub = stub
}

func (fake *FakeAuthClient) AuthReturns(result1 bool, result2 error) {
	fake.authMutex.Lock()
	defer fake.authMutex.Unlock()
	fake.AuthStub = nil
	fake.authReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthClient) AuthReturnsOnCall(i int, result1 bool, result2 error) {
	fake.authMutex.Lock()
	defer fake.authMutex.Unlock()
	fake.AuthStub = nil
	if fake.authReturnsOnCall == nil {
		fake.authReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.authReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authMutex.RLock()
	defer fake.authMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthClient) recordInvocation(key string, args []interface{}) {
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

var _ auth.AuthClient = new(FakeAuthClient)
