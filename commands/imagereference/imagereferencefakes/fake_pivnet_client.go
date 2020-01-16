// Code generated by counterfeiter. DO NOT EDIT.
package imagereferencefakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v4"
	"github.com/pivotal-cf/pivnet-cli/commands/imagereference"
)

type FakePivnetClient struct {
	AddImageReferenceToReleaseStub        func(string, int, int) error
	addImageReferenceToReleaseMutex       sync.RWMutex
	addImageReferenceToReleaseArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	addImageReferenceToReleaseReturns struct {
		result1 error
	}
	addImageReferenceToReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	CreateImageReferenceStub        func(pivnet.CreateImageReferenceConfig) (pivnet.ImageReference, error)
	createImageReferenceMutex       sync.RWMutex
	createImageReferenceArgsForCall []struct {
		arg1 pivnet.CreateImageReferenceConfig
	}
	createImageReferenceReturns struct {
		result1 pivnet.ImageReference
		result2 error
	}
	createImageReferenceReturnsOnCall map[int]struct {
		result1 pivnet.ImageReference
		result2 error
	}
	DeleteImageReferenceStub        func(string, int) (pivnet.ImageReference, error)
	deleteImageReferenceMutex       sync.RWMutex
	deleteImageReferenceArgsForCall []struct {
		arg1 string
		arg2 int
	}
	deleteImageReferenceReturns struct {
		result1 pivnet.ImageReference
		result2 error
	}
	deleteImageReferenceReturnsOnCall map[int]struct {
		result1 pivnet.ImageReference
		result2 error
	}
	ImageReferenceStub        func(string, int) (pivnet.ImageReference, error)
	imageReferenceMutex       sync.RWMutex
	imageReferenceArgsForCall []struct {
		arg1 string
		arg2 int
	}
	imageReferenceReturns struct {
		result1 pivnet.ImageReference
		result2 error
	}
	imageReferenceReturnsOnCall map[int]struct {
		result1 pivnet.ImageReference
		result2 error
	}
	ImageReferenceForReleaseStub        func(string, int, int) (pivnet.ImageReference, error)
	imageReferenceForReleaseMutex       sync.RWMutex
	imageReferenceForReleaseArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	imageReferenceForReleaseReturns struct {
		result1 pivnet.ImageReference
		result2 error
	}
	imageReferenceForReleaseReturnsOnCall map[int]struct {
		result1 pivnet.ImageReference
		result2 error
	}
	ImageReferencesStub        func(string) ([]pivnet.ImageReference, error)
	imageReferencesMutex       sync.RWMutex
	imageReferencesArgsForCall []struct {
		arg1 string
	}
	imageReferencesReturns struct {
		result1 []pivnet.ImageReference
		result2 error
	}
	imageReferencesReturnsOnCall map[int]struct {
		result1 []pivnet.ImageReference
		result2 error
	}
	ImageReferencesForDigestStub        func(string, string) ([]pivnet.ImageReference, error)
	imageReferencesForDigestMutex       sync.RWMutex
	imageReferencesForDigestArgsForCall []struct {
		arg1 string
		arg2 string
	}
	imageReferencesForDigestReturns struct {
		result1 []pivnet.ImageReference
		result2 error
	}
	imageReferencesForDigestReturnsOnCall map[int]struct {
		result1 []pivnet.ImageReference
		result2 error
	}
	ImageReferencesForReleaseStub        func(string, int) ([]pivnet.ImageReference, error)
	imageReferencesForReleaseMutex       sync.RWMutex
	imageReferencesForReleaseArgsForCall []struct {
		arg1 string
		arg2 int
	}
	imageReferencesForReleaseReturns struct {
		result1 []pivnet.ImageReference
		result2 error
	}
	imageReferencesForReleaseReturnsOnCall map[int]struct {
		result1 []pivnet.ImageReference
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
	RemoveImageReferenceFromReleaseStub        func(string, int, int) error
	removeImageReferenceFromReleaseMutex       sync.RWMutex
	removeImageReferenceFromReleaseArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	removeImageReferenceFromReleaseReturns struct {
		result1 error
	}
	removeImageReferenceFromReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateImageReferenceStub        func(string, pivnet.ImageReference) (pivnet.ImageReference, error)
	updateImageReferenceMutex       sync.RWMutex
	updateImageReferenceArgsForCall []struct {
		arg1 string
		arg2 pivnet.ImageReference
	}
	updateImageReferenceReturns struct {
		result1 pivnet.ImageReference
		result2 error
	}
	updateImageReferenceReturnsOnCall map[int]struct {
		result1 pivnet.ImageReference
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) AddImageReferenceToRelease(arg1 string, arg2 int, arg3 int) error {
	fake.addImageReferenceToReleaseMutex.Lock()
	ret, specificReturn := fake.addImageReferenceToReleaseReturnsOnCall[len(fake.addImageReferenceToReleaseArgsForCall)]
	fake.addImageReferenceToReleaseArgsForCall = append(fake.addImageReferenceToReleaseArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("AddImageReferenceToRelease", []interface{}{arg1, arg2, arg3})
	fake.addImageReferenceToReleaseMutex.Unlock()
	if fake.AddImageReferenceToReleaseStub != nil {
		return fake.AddImageReferenceToReleaseStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addImageReferenceToReleaseReturns
	return fakeReturns.result1
}

func (fake *FakePivnetClient) AddImageReferenceToReleaseCallCount() int {
	fake.addImageReferenceToReleaseMutex.RLock()
	defer fake.addImageReferenceToReleaseMutex.RUnlock()
	return len(fake.addImageReferenceToReleaseArgsForCall)
}

func (fake *FakePivnetClient) AddImageReferenceToReleaseCalls(stub func(string, int, int) error) {
	fake.addImageReferenceToReleaseMutex.Lock()
	defer fake.addImageReferenceToReleaseMutex.Unlock()
	fake.AddImageReferenceToReleaseStub = stub
}

func (fake *FakePivnetClient) AddImageReferenceToReleaseArgsForCall(i int) (string, int, int) {
	fake.addImageReferenceToReleaseMutex.RLock()
	defer fake.addImageReferenceToReleaseMutex.RUnlock()
	argsForCall := fake.addImageReferenceToReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePivnetClient) AddImageReferenceToReleaseReturns(result1 error) {
	fake.addImageReferenceToReleaseMutex.Lock()
	defer fake.addImageReferenceToReleaseMutex.Unlock()
	fake.AddImageReferenceToReleaseStub = nil
	fake.addImageReferenceToReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AddImageReferenceToReleaseReturnsOnCall(i int, result1 error) {
	fake.addImageReferenceToReleaseMutex.Lock()
	defer fake.addImageReferenceToReleaseMutex.Unlock()
	fake.AddImageReferenceToReleaseStub = nil
	if fake.addImageReferenceToReleaseReturnsOnCall == nil {
		fake.addImageReferenceToReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addImageReferenceToReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) CreateImageReference(arg1 pivnet.CreateImageReferenceConfig) (pivnet.ImageReference, error) {
	fake.createImageReferenceMutex.Lock()
	ret, specificReturn := fake.createImageReferenceReturnsOnCall[len(fake.createImageReferenceArgsForCall)]
	fake.createImageReferenceArgsForCall = append(fake.createImageReferenceArgsForCall, struct {
		arg1 pivnet.CreateImageReferenceConfig
	}{arg1})
	fake.recordInvocation("CreateImageReference", []interface{}{arg1})
	fake.createImageReferenceMutex.Unlock()
	if fake.CreateImageReferenceStub != nil {
		return fake.CreateImageReferenceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createImageReferenceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) CreateImageReferenceCallCount() int {
	fake.createImageReferenceMutex.RLock()
	defer fake.createImageReferenceMutex.RUnlock()
	return len(fake.createImageReferenceArgsForCall)
}

func (fake *FakePivnetClient) CreateImageReferenceCalls(stub func(pivnet.CreateImageReferenceConfig) (pivnet.ImageReference, error)) {
	fake.createImageReferenceMutex.Lock()
	defer fake.createImageReferenceMutex.Unlock()
	fake.CreateImageReferenceStub = stub
}

func (fake *FakePivnetClient) CreateImageReferenceArgsForCall(i int) pivnet.CreateImageReferenceConfig {
	fake.createImageReferenceMutex.RLock()
	defer fake.createImageReferenceMutex.RUnlock()
	argsForCall := fake.createImageReferenceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePivnetClient) CreateImageReferenceReturns(result1 pivnet.ImageReference, result2 error) {
	fake.createImageReferenceMutex.Lock()
	defer fake.createImageReferenceMutex.Unlock()
	fake.CreateImageReferenceStub = nil
	fake.createImageReferenceReturns = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) CreateImageReferenceReturnsOnCall(i int, result1 pivnet.ImageReference, result2 error) {
	fake.createImageReferenceMutex.Lock()
	defer fake.createImageReferenceMutex.Unlock()
	fake.CreateImageReferenceStub = nil
	if fake.createImageReferenceReturnsOnCall == nil {
		fake.createImageReferenceReturnsOnCall = make(map[int]struct {
			result1 pivnet.ImageReference
			result2 error
		})
	}
	fake.createImageReferenceReturnsOnCall[i] = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DeleteImageReference(arg1 string, arg2 int) (pivnet.ImageReference, error) {
	fake.deleteImageReferenceMutex.Lock()
	ret, specificReturn := fake.deleteImageReferenceReturnsOnCall[len(fake.deleteImageReferenceArgsForCall)]
	fake.deleteImageReferenceArgsForCall = append(fake.deleteImageReferenceArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("DeleteImageReference", []interface{}{arg1, arg2})
	fake.deleteImageReferenceMutex.Unlock()
	if fake.DeleteImageReferenceStub != nil {
		return fake.DeleteImageReferenceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteImageReferenceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) DeleteImageReferenceCallCount() int {
	fake.deleteImageReferenceMutex.RLock()
	defer fake.deleteImageReferenceMutex.RUnlock()
	return len(fake.deleteImageReferenceArgsForCall)
}

func (fake *FakePivnetClient) DeleteImageReferenceCalls(stub func(string, int) (pivnet.ImageReference, error)) {
	fake.deleteImageReferenceMutex.Lock()
	defer fake.deleteImageReferenceMutex.Unlock()
	fake.DeleteImageReferenceStub = stub
}

func (fake *FakePivnetClient) DeleteImageReferenceArgsForCall(i int) (string, int) {
	fake.deleteImageReferenceMutex.RLock()
	defer fake.deleteImageReferenceMutex.RUnlock()
	argsForCall := fake.deleteImageReferenceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) DeleteImageReferenceReturns(result1 pivnet.ImageReference, result2 error) {
	fake.deleteImageReferenceMutex.Lock()
	defer fake.deleteImageReferenceMutex.Unlock()
	fake.DeleteImageReferenceStub = nil
	fake.deleteImageReferenceReturns = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DeleteImageReferenceReturnsOnCall(i int, result1 pivnet.ImageReference, result2 error) {
	fake.deleteImageReferenceMutex.Lock()
	defer fake.deleteImageReferenceMutex.Unlock()
	fake.DeleteImageReferenceStub = nil
	if fake.deleteImageReferenceReturnsOnCall == nil {
		fake.deleteImageReferenceReturnsOnCall = make(map[int]struct {
			result1 pivnet.ImageReference
			result2 error
		})
	}
	fake.deleteImageReferenceReturnsOnCall[i] = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReference(arg1 string, arg2 int) (pivnet.ImageReference, error) {
	fake.imageReferenceMutex.Lock()
	ret, specificReturn := fake.imageReferenceReturnsOnCall[len(fake.imageReferenceArgsForCall)]
	fake.imageReferenceArgsForCall = append(fake.imageReferenceArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("ImageReference", []interface{}{arg1, arg2})
	fake.imageReferenceMutex.Unlock()
	if fake.ImageReferenceStub != nil {
		return fake.ImageReferenceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.imageReferenceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ImageReferenceCallCount() int {
	fake.imageReferenceMutex.RLock()
	defer fake.imageReferenceMutex.RUnlock()
	return len(fake.imageReferenceArgsForCall)
}

func (fake *FakePivnetClient) ImageReferenceCalls(stub func(string, int) (pivnet.ImageReference, error)) {
	fake.imageReferenceMutex.Lock()
	defer fake.imageReferenceMutex.Unlock()
	fake.ImageReferenceStub = stub
}

func (fake *FakePivnetClient) ImageReferenceArgsForCall(i int) (string, int) {
	fake.imageReferenceMutex.RLock()
	defer fake.imageReferenceMutex.RUnlock()
	argsForCall := fake.imageReferenceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) ImageReferenceReturns(result1 pivnet.ImageReference, result2 error) {
	fake.imageReferenceMutex.Lock()
	defer fake.imageReferenceMutex.Unlock()
	fake.ImageReferenceStub = nil
	fake.imageReferenceReturns = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReferenceReturnsOnCall(i int, result1 pivnet.ImageReference, result2 error) {
	fake.imageReferenceMutex.Lock()
	defer fake.imageReferenceMutex.Unlock()
	fake.ImageReferenceStub = nil
	if fake.imageReferenceReturnsOnCall == nil {
		fake.imageReferenceReturnsOnCall = make(map[int]struct {
			result1 pivnet.ImageReference
			result2 error
		})
	}
	fake.imageReferenceReturnsOnCall[i] = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReferenceForRelease(arg1 string, arg2 int, arg3 int) (pivnet.ImageReference, error) {
	fake.imageReferenceForReleaseMutex.Lock()
	ret, specificReturn := fake.imageReferenceForReleaseReturnsOnCall[len(fake.imageReferenceForReleaseArgsForCall)]
	fake.imageReferenceForReleaseArgsForCall = append(fake.imageReferenceForReleaseArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("ImageReferenceForRelease", []interface{}{arg1, arg2, arg3})
	fake.imageReferenceForReleaseMutex.Unlock()
	if fake.ImageReferenceForReleaseStub != nil {
		return fake.ImageReferenceForReleaseStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.imageReferenceForReleaseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ImageReferenceForReleaseCallCount() int {
	fake.imageReferenceForReleaseMutex.RLock()
	defer fake.imageReferenceForReleaseMutex.RUnlock()
	return len(fake.imageReferenceForReleaseArgsForCall)
}

func (fake *FakePivnetClient) ImageReferenceForReleaseCalls(stub func(string, int, int) (pivnet.ImageReference, error)) {
	fake.imageReferenceForReleaseMutex.Lock()
	defer fake.imageReferenceForReleaseMutex.Unlock()
	fake.ImageReferenceForReleaseStub = stub
}

func (fake *FakePivnetClient) ImageReferenceForReleaseArgsForCall(i int) (string, int, int) {
	fake.imageReferenceForReleaseMutex.RLock()
	defer fake.imageReferenceForReleaseMutex.RUnlock()
	argsForCall := fake.imageReferenceForReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePivnetClient) ImageReferenceForReleaseReturns(result1 pivnet.ImageReference, result2 error) {
	fake.imageReferenceForReleaseMutex.Lock()
	defer fake.imageReferenceForReleaseMutex.Unlock()
	fake.ImageReferenceForReleaseStub = nil
	fake.imageReferenceForReleaseReturns = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReferenceForReleaseReturnsOnCall(i int, result1 pivnet.ImageReference, result2 error) {
	fake.imageReferenceForReleaseMutex.Lock()
	defer fake.imageReferenceForReleaseMutex.Unlock()
	fake.ImageReferenceForReleaseStub = nil
	if fake.imageReferenceForReleaseReturnsOnCall == nil {
		fake.imageReferenceForReleaseReturnsOnCall = make(map[int]struct {
			result1 pivnet.ImageReference
			result2 error
		})
	}
	fake.imageReferenceForReleaseReturnsOnCall[i] = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReferences(arg1 string) ([]pivnet.ImageReference, error) {
	fake.imageReferencesMutex.Lock()
	ret, specificReturn := fake.imageReferencesReturnsOnCall[len(fake.imageReferencesArgsForCall)]
	fake.imageReferencesArgsForCall = append(fake.imageReferencesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ImageReferences", []interface{}{arg1})
	fake.imageReferencesMutex.Unlock()
	if fake.ImageReferencesStub != nil {
		return fake.ImageReferencesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.imageReferencesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ImageReferencesCallCount() int {
	fake.imageReferencesMutex.RLock()
	defer fake.imageReferencesMutex.RUnlock()
	return len(fake.imageReferencesArgsForCall)
}

func (fake *FakePivnetClient) ImageReferencesCalls(stub func(string) ([]pivnet.ImageReference, error)) {
	fake.imageReferencesMutex.Lock()
	defer fake.imageReferencesMutex.Unlock()
	fake.ImageReferencesStub = stub
}

func (fake *FakePivnetClient) ImageReferencesArgsForCall(i int) string {
	fake.imageReferencesMutex.RLock()
	defer fake.imageReferencesMutex.RUnlock()
	argsForCall := fake.imageReferencesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePivnetClient) ImageReferencesReturns(result1 []pivnet.ImageReference, result2 error) {
	fake.imageReferencesMutex.Lock()
	defer fake.imageReferencesMutex.Unlock()
	fake.ImageReferencesStub = nil
	fake.imageReferencesReturns = struct {
		result1 []pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReferencesReturnsOnCall(i int, result1 []pivnet.ImageReference, result2 error) {
	fake.imageReferencesMutex.Lock()
	defer fake.imageReferencesMutex.Unlock()
	fake.ImageReferencesStub = nil
	if fake.imageReferencesReturnsOnCall == nil {
		fake.imageReferencesReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ImageReference
			result2 error
		})
	}
	fake.imageReferencesReturnsOnCall[i] = struct {
		result1 []pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReferencesForDigest(arg1 string, arg2 string) ([]pivnet.ImageReference, error) {
	fake.imageReferencesForDigestMutex.Lock()
	ret, specificReturn := fake.imageReferencesForDigestReturnsOnCall[len(fake.imageReferencesForDigestArgsForCall)]
	fake.imageReferencesForDigestArgsForCall = append(fake.imageReferencesForDigestArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ImageReferencesForDigest", []interface{}{arg1, arg2})
	fake.imageReferencesForDigestMutex.Unlock()
	if fake.ImageReferencesForDigestStub != nil {
		return fake.ImageReferencesForDigestStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.imageReferencesForDigestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ImageReferencesForDigestCallCount() int {
	fake.imageReferencesForDigestMutex.RLock()
	defer fake.imageReferencesForDigestMutex.RUnlock()
	return len(fake.imageReferencesForDigestArgsForCall)
}

func (fake *FakePivnetClient) ImageReferencesForDigestCalls(stub func(string, string) ([]pivnet.ImageReference, error)) {
	fake.imageReferencesForDigestMutex.Lock()
	defer fake.imageReferencesForDigestMutex.Unlock()
	fake.ImageReferencesForDigestStub = stub
}

func (fake *FakePivnetClient) ImageReferencesForDigestArgsForCall(i int) (string, string) {
	fake.imageReferencesForDigestMutex.RLock()
	defer fake.imageReferencesForDigestMutex.RUnlock()
	argsForCall := fake.imageReferencesForDigestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) ImageReferencesForDigestReturns(result1 []pivnet.ImageReference, result2 error) {
	fake.imageReferencesForDigestMutex.Lock()
	defer fake.imageReferencesForDigestMutex.Unlock()
	fake.ImageReferencesForDigestStub = nil
	fake.imageReferencesForDigestReturns = struct {
		result1 []pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReferencesForDigestReturnsOnCall(i int, result1 []pivnet.ImageReference, result2 error) {
	fake.imageReferencesForDigestMutex.Lock()
	defer fake.imageReferencesForDigestMutex.Unlock()
	fake.ImageReferencesForDigestStub = nil
	if fake.imageReferencesForDigestReturnsOnCall == nil {
		fake.imageReferencesForDigestReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ImageReference
			result2 error
		})
	}
	fake.imageReferencesForDigestReturnsOnCall[i] = struct {
		result1 []pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReferencesForRelease(arg1 string, arg2 int) ([]pivnet.ImageReference, error) {
	fake.imageReferencesForReleaseMutex.Lock()
	ret, specificReturn := fake.imageReferencesForReleaseReturnsOnCall[len(fake.imageReferencesForReleaseArgsForCall)]
	fake.imageReferencesForReleaseArgsForCall = append(fake.imageReferencesForReleaseArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("ImageReferencesForRelease", []interface{}{arg1, arg2})
	fake.imageReferencesForReleaseMutex.Unlock()
	if fake.ImageReferencesForReleaseStub != nil {
		return fake.ImageReferencesForReleaseStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.imageReferencesForReleaseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ImageReferencesForReleaseCallCount() int {
	fake.imageReferencesForReleaseMutex.RLock()
	defer fake.imageReferencesForReleaseMutex.RUnlock()
	return len(fake.imageReferencesForReleaseArgsForCall)
}

func (fake *FakePivnetClient) ImageReferencesForReleaseCalls(stub func(string, int) ([]pivnet.ImageReference, error)) {
	fake.imageReferencesForReleaseMutex.Lock()
	defer fake.imageReferencesForReleaseMutex.Unlock()
	fake.ImageReferencesForReleaseStub = stub
}

func (fake *FakePivnetClient) ImageReferencesForReleaseArgsForCall(i int) (string, int) {
	fake.imageReferencesForReleaseMutex.RLock()
	defer fake.imageReferencesForReleaseMutex.RUnlock()
	argsForCall := fake.imageReferencesForReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) ImageReferencesForReleaseReturns(result1 []pivnet.ImageReference, result2 error) {
	fake.imageReferencesForReleaseMutex.Lock()
	defer fake.imageReferencesForReleaseMutex.Unlock()
	fake.ImageReferencesForReleaseStub = nil
	fake.imageReferencesForReleaseReturns = struct {
		result1 []pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReferencesForReleaseReturnsOnCall(i int, result1 []pivnet.ImageReference, result2 error) {
	fake.imageReferencesForReleaseMutex.Lock()
	defer fake.imageReferencesForReleaseMutex.Unlock()
	fake.ImageReferencesForReleaseStub = nil
	if fake.imageReferencesForReleaseReturnsOnCall == nil {
		fake.imageReferencesForReleaseReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ImageReference
			result2 error
		})
	}
	fake.imageReferencesForReleaseReturnsOnCall[i] = struct {
		result1 []pivnet.ImageReference
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

func (fake *FakePivnetClient) RemoveImageReferenceFromRelease(arg1 string, arg2 int, arg3 int) error {
	fake.removeImageReferenceFromReleaseMutex.Lock()
	ret, specificReturn := fake.removeImageReferenceFromReleaseReturnsOnCall[len(fake.removeImageReferenceFromReleaseArgsForCall)]
	fake.removeImageReferenceFromReleaseArgsForCall = append(fake.removeImageReferenceFromReleaseArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("RemoveImageReferenceFromRelease", []interface{}{arg1, arg2, arg3})
	fake.removeImageReferenceFromReleaseMutex.Unlock()
	if fake.RemoveImageReferenceFromReleaseStub != nil {
		return fake.RemoveImageReferenceFromReleaseStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeImageReferenceFromReleaseReturns
	return fakeReturns.result1
}

func (fake *FakePivnetClient) RemoveImageReferenceFromReleaseCallCount() int {
	fake.removeImageReferenceFromReleaseMutex.RLock()
	defer fake.removeImageReferenceFromReleaseMutex.RUnlock()
	return len(fake.removeImageReferenceFromReleaseArgsForCall)
}

func (fake *FakePivnetClient) RemoveImageReferenceFromReleaseCalls(stub func(string, int, int) error) {
	fake.removeImageReferenceFromReleaseMutex.Lock()
	defer fake.removeImageReferenceFromReleaseMutex.Unlock()
	fake.RemoveImageReferenceFromReleaseStub = stub
}

func (fake *FakePivnetClient) RemoveImageReferenceFromReleaseArgsForCall(i int) (string, int, int) {
	fake.removeImageReferenceFromReleaseMutex.RLock()
	defer fake.removeImageReferenceFromReleaseMutex.RUnlock()
	argsForCall := fake.removeImageReferenceFromReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePivnetClient) RemoveImageReferenceFromReleaseReturns(result1 error) {
	fake.removeImageReferenceFromReleaseMutex.Lock()
	defer fake.removeImageReferenceFromReleaseMutex.Unlock()
	fake.RemoveImageReferenceFromReleaseStub = nil
	fake.removeImageReferenceFromReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveImageReferenceFromReleaseReturnsOnCall(i int, result1 error) {
	fake.removeImageReferenceFromReleaseMutex.Lock()
	defer fake.removeImageReferenceFromReleaseMutex.Unlock()
	fake.RemoveImageReferenceFromReleaseStub = nil
	if fake.removeImageReferenceFromReleaseReturnsOnCall == nil {
		fake.removeImageReferenceFromReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeImageReferenceFromReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) UpdateImageReference(arg1 string, arg2 pivnet.ImageReference) (pivnet.ImageReference, error) {
	fake.updateImageReferenceMutex.Lock()
	ret, specificReturn := fake.updateImageReferenceReturnsOnCall[len(fake.updateImageReferenceArgsForCall)]
	fake.updateImageReferenceArgsForCall = append(fake.updateImageReferenceArgsForCall, struct {
		arg1 string
		arg2 pivnet.ImageReference
	}{arg1, arg2})
	fake.recordInvocation("UpdateImageReference", []interface{}{arg1, arg2})
	fake.updateImageReferenceMutex.Unlock()
	if fake.UpdateImageReferenceStub != nil {
		return fake.UpdateImageReferenceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateImageReferenceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) UpdateImageReferenceCallCount() int {
	fake.updateImageReferenceMutex.RLock()
	defer fake.updateImageReferenceMutex.RUnlock()
	return len(fake.updateImageReferenceArgsForCall)
}

func (fake *FakePivnetClient) UpdateImageReferenceCalls(stub func(string, pivnet.ImageReference) (pivnet.ImageReference, error)) {
	fake.updateImageReferenceMutex.Lock()
	defer fake.updateImageReferenceMutex.Unlock()
	fake.UpdateImageReferenceStub = stub
}

func (fake *FakePivnetClient) UpdateImageReferenceArgsForCall(i int) (string, pivnet.ImageReference) {
	fake.updateImageReferenceMutex.RLock()
	defer fake.updateImageReferenceMutex.RUnlock()
	argsForCall := fake.updateImageReferenceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) UpdateImageReferenceReturns(result1 pivnet.ImageReference, result2 error) {
	fake.updateImageReferenceMutex.Lock()
	defer fake.updateImageReferenceMutex.Unlock()
	fake.UpdateImageReferenceStub = nil
	fake.updateImageReferenceReturns = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpdateImageReferenceReturnsOnCall(i int, result1 pivnet.ImageReference, result2 error) {
	fake.updateImageReferenceMutex.Lock()
	defer fake.updateImageReferenceMutex.Unlock()
	fake.UpdateImageReferenceStub = nil
	if fake.updateImageReferenceReturnsOnCall == nil {
		fake.updateImageReferenceReturnsOnCall = make(map[int]struct {
			result1 pivnet.ImageReference
			result2 error
		})
	}
	fake.updateImageReferenceReturnsOnCall[i] = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addImageReferenceToReleaseMutex.RLock()
	defer fake.addImageReferenceToReleaseMutex.RUnlock()
	fake.createImageReferenceMutex.RLock()
	defer fake.createImageReferenceMutex.RUnlock()
	fake.deleteImageReferenceMutex.RLock()
	defer fake.deleteImageReferenceMutex.RUnlock()
	fake.imageReferenceMutex.RLock()
	defer fake.imageReferenceMutex.RUnlock()
	fake.imageReferenceForReleaseMutex.RLock()
	defer fake.imageReferenceForReleaseMutex.RUnlock()
	fake.imageReferencesMutex.RLock()
	defer fake.imageReferencesMutex.RUnlock()
	fake.imageReferencesForDigestMutex.RLock()
	defer fake.imageReferencesForDigestMutex.RUnlock()
	fake.imageReferencesForReleaseMutex.RLock()
	defer fake.imageReferencesForReleaseMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	fake.removeImageReferenceFromReleaseMutex.RLock()
	defer fake.removeImageReferenceFromReleaseMutex.RUnlock()
	fake.updateImageReferenceMutex.RLock()
	defer fake.updateImageReferenceMutex.RUnlock()
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

var _ imagereference.PivnetClient = new(FakePivnetClient)
