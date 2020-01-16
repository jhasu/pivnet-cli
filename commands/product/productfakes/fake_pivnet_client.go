// Code generated by counterfeiter. DO NOT EDIT.
package productfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v4"
	"github.com/pivotal-cf/pivnet-cli/commands/product"
)

type FakePivnetClient struct {
	FindProductForSlugStub        func(string) (pivnet.Product, error)
	findProductForSlugMutex       sync.RWMutex
	findProductForSlugArgsForCall []struct {
		arg1 string
	}
	findProductForSlugReturns struct {
		result1 pivnet.Product
		result2 error
	}
	findProductForSlugReturnsOnCall map[int]struct {
		result1 pivnet.Product
		result2 error
	}
	ProductsStub        func() ([]pivnet.Product, error)
	productsMutex       sync.RWMutex
	productsArgsForCall []struct {
	}
	productsReturns struct {
		result1 []pivnet.Product
		result2 error
	}
	productsReturnsOnCall map[int]struct {
		result1 []pivnet.Product
		result2 error
	}
	SlugAliasStub        func(string) (pivnet.SlugAliasResponse, error)
	slugAliasMutex       sync.RWMutex
	slugAliasArgsForCall []struct {
		arg1 string
	}
	slugAliasReturns struct {
		result1 pivnet.SlugAliasResponse
		result2 error
	}
	slugAliasReturnsOnCall map[int]struct {
		result1 pivnet.SlugAliasResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) FindProductForSlug(arg1 string) (pivnet.Product, error) {
	fake.findProductForSlugMutex.Lock()
	ret, specificReturn := fake.findProductForSlugReturnsOnCall[len(fake.findProductForSlugArgsForCall)]
	fake.findProductForSlugArgsForCall = append(fake.findProductForSlugArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindProductForSlug", []interface{}{arg1})
	fake.findProductForSlugMutex.Unlock()
	if fake.FindProductForSlugStub != nil {
		return fake.FindProductForSlugStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findProductForSlugReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) FindProductForSlugCallCount() int {
	fake.findProductForSlugMutex.RLock()
	defer fake.findProductForSlugMutex.RUnlock()
	return len(fake.findProductForSlugArgsForCall)
}

func (fake *FakePivnetClient) FindProductForSlugCalls(stub func(string) (pivnet.Product, error)) {
	fake.findProductForSlugMutex.Lock()
	defer fake.findProductForSlugMutex.Unlock()
	fake.FindProductForSlugStub = stub
}

func (fake *FakePivnetClient) FindProductForSlugArgsForCall(i int) string {
	fake.findProductForSlugMutex.RLock()
	defer fake.findProductForSlugMutex.RUnlock()
	argsForCall := fake.findProductForSlugArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePivnetClient) FindProductForSlugReturns(result1 pivnet.Product, result2 error) {
	fake.findProductForSlugMutex.Lock()
	defer fake.findProductForSlugMutex.Unlock()
	fake.FindProductForSlugStub = nil
	fake.findProductForSlugReturns = struct {
		result1 pivnet.Product
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) FindProductForSlugReturnsOnCall(i int, result1 pivnet.Product, result2 error) {
	fake.findProductForSlugMutex.Lock()
	defer fake.findProductForSlugMutex.Unlock()
	fake.FindProductForSlugStub = nil
	if fake.findProductForSlugReturnsOnCall == nil {
		fake.findProductForSlugReturnsOnCall = make(map[int]struct {
			result1 pivnet.Product
			result2 error
		})
	}
	fake.findProductForSlugReturnsOnCall[i] = struct {
		result1 pivnet.Product
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Products() ([]pivnet.Product, error) {
	fake.productsMutex.Lock()
	ret, specificReturn := fake.productsReturnsOnCall[len(fake.productsArgsForCall)]
	fake.productsArgsForCall = append(fake.productsArgsForCall, struct {
	}{})
	fake.recordInvocation("Products", []interface{}{})
	fake.productsMutex.Unlock()
	if fake.ProductsStub != nil {
		return fake.ProductsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.productsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ProductsCallCount() int {
	fake.productsMutex.RLock()
	defer fake.productsMutex.RUnlock()
	return len(fake.productsArgsForCall)
}

func (fake *FakePivnetClient) ProductsCalls(stub func() ([]pivnet.Product, error)) {
	fake.productsMutex.Lock()
	defer fake.productsMutex.Unlock()
	fake.ProductsStub = stub
}

func (fake *FakePivnetClient) ProductsReturns(result1 []pivnet.Product, result2 error) {
	fake.productsMutex.Lock()
	defer fake.productsMutex.Unlock()
	fake.ProductsStub = nil
	fake.productsReturns = struct {
		result1 []pivnet.Product
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductsReturnsOnCall(i int, result1 []pivnet.Product, result2 error) {
	fake.productsMutex.Lock()
	defer fake.productsMutex.Unlock()
	fake.ProductsStub = nil
	if fake.productsReturnsOnCall == nil {
		fake.productsReturnsOnCall = make(map[int]struct {
			result1 []pivnet.Product
			result2 error
		})
	}
	fake.productsReturnsOnCall[i] = struct {
		result1 []pivnet.Product
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) SlugAlias(arg1 string) (pivnet.SlugAliasResponse, error) {
	fake.slugAliasMutex.Lock()
	ret, specificReturn := fake.slugAliasReturnsOnCall[len(fake.slugAliasArgsForCall)]
	fake.slugAliasArgsForCall = append(fake.slugAliasArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SlugAlias", []interface{}{arg1})
	fake.slugAliasMutex.Unlock()
	if fake.SlugAliasStub != nil {
		return fake.SlugAliasStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.slugAliasReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) SlugAliasCallCount() int {
	fake.slugAliasMutex.RLock()
	defer fake.slugAliasMutex.RUnlock()
	return len(fake.slugAliasArgsForCall)
}

func (fake *FakePivnetClient) SlugAliasCalls(stub func(string) (pivnet.SlugAliasResponse, error)) {
	fake.slugAliasMutex.Lock()
	defer fake.slugAliasMutex.Unlock()
	fake.SlugAliasStub = stub
}

func (fake *FakePivnetClient) SlugAliasArgsForCall(i int) string {
	fake.slugAliasMutex.RLock()
	defer fake.slugAliasMutex.RUnlock()
	argsForCall := fake.slugAliasArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePivnetClient) SlugAliasReturns(result1 pivnet.SlugAliasResponse, result2 error) {
	fake.slugAliasMutex.Lock()
	defer fake.slugAliasMutex.Unlock()
	fake.SlugAliasStub = nil
	fake.slugAliasReturns = struct {
		result1 pivnet.SlugAliasResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) SlugAliasReturnsOnCall(i int, result1 pivnet.SlugAliasResponse, result2 error) {
	fake.slugAliasMutex.Lock()
	defer fake.slugAliasMutex.Unlock()
	fake.SlugAliasStub = nil
	if fake.slugAliasReturnsOnCall == nil {
		fake.slugAliasReturnsOnCall = make(map[int]struct {
			result1 pivnet.SlugAliasResponse
			result2 error
		})
	}
	fake.slugAliasReturnsOnCall[i] = struct {
		result1 pivnet.SlugAliasResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findProductForSlugMutex.RLock()
	defer fake.findProductForSlugMutex.RUnlock()
	fake.productsMutex.RLock()
	defer fake.productsMutex.RUnlock()
	fake.slugAliasMutex.RLock()
	defer fake.slugAliasMutex.RUnlock()
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

var _ product.PivnetClient = new(FakePivnetClient)
