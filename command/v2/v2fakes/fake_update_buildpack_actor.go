// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
	"code.cloudfoundry.org/cli/types"
)

type FakeUpdateBuildpackActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	UpdateBuildpackByNameStub        func(name string, position types.NullInt, locked types.NullBool, enabled types.NullBool) (string, v2action.Warnings, error)
	updateBuildpackByNameMutex       sync.RWMutex
	updateBuildpackByNameArgsForCall []struct {
		name     string
		position types.NullInt
		locked   types.NullBool
		enabled  types.NullBool
	}
	updateBuildpackByNameReturns struct {
		result1 string
		result2 v2action.Warnings
		result3 error
	}
	updateBuildpackByNameReturnsOnCall map[int]struct {
		result1 string
		result2 v2action.Warnings
		result3 error
	}
	PrepareBuildpackBitsStub        func(inputPath string, tmpDirPath string, downloader v2action.Downloader) (string, error)
	prepareBuildpackBitsMutex       sync.RWMutex
	prepareBuildpackBitsArgsForCall []struct {
		inputPath  string
		tmpDirPath string
		downloader v2action.Downloader
	}
	prepareBuildpackBitsReturns struct {
		result1 string
		result2 error
	}
	prepareBuildpackBitsReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UploadBuildpackStub        func(GUID string, path string, progBar v2action.SimpleProgressBar) (v2action.Warnings, error)
	uploadBuildpackMutex       sync.RWMutex
	uploadBuildpackArgsForCall []struct {
		GUID    string
		path    string
		progBar v2action.SimpleProgressBar
	}
	uploadBuildpackReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	uploadBuildpackReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUpdateBuildpackActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeUpdateBuildpackActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeUpdateBuildpackActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeUpdateBuildpackActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeUpdateBuildpackActor) UpdateBuildpackByName(name string, position types.NullInt, locked types.NullBool, enabled types.NullBool) (string, v2action.Warnings, error) {
	fake.updateBuildpackByNameMutex.Lock()
	ret, specificReturn := fake.updateBuildpackByNameReturnsOnCall[len(fake.updateBuildpackByNameArgsForCall)]
	fake.updateBuildpackByNameArgsForCall = append(fake.updateBuildpackByNameArgsForCall, struct {
		name     string
		position types.NullInt
		locked   types.NullBool
		enabled  types.NullBool
	}{name, position, locked, enabled})
	fake.recordInvocation("UpdateBuildpackByName", []interface{}{name, position, locked, enabled})
	fake.updateBuildpackByNameMutex.Unlock()
	if fake.UpdateBuildpackByNameStub != nil {
		return fake.UpdateBuildpackByNameStub(name, position, locked, enabled)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.updateBuildpackByNameReturns.result1, fake.updateBuildpackByNameReturns.result2, fake.updateBuildpackByNameReturns.result3
}

func (fake *FakeUpdateBuildpackActor) UpdateBuildpackByNameCallCount() int {
	fake.updateBuildpackByNameMutex.RLock()
	defer fake.updateBuildpackByNameMutex.RUnlock()
	return len(fake.updateBuildpackByNameArgsForCall)
}

func (fake *FakeUpdateBuildpackActor) UpdateBuildpackByNameArgsForCall(i int) (string, types.NullInt, types.NullBool, types.NullBool) {
	fake.updateBuildpackByNameMutex.RLock()
	defer fake.updateBuildpackByNameMutex.RUnlock()
	return fake.updateBuildpackByNameArgsForCall[i].name, fake.updateBuildpackByNameArgsForCall[i].position, fake.updateBuildpackByNameArgsForCall[i].locked, fake.updateBuildpackByNameArgsForCall[i].enabled
}

func (fake *FakeUpdateBuildpackActor) UpdateBuildpackByNameReturns(result1 string, result2 v2action.Warnings, result3 error) {
	fake.UpdateBuildpackByNameStub = nil
	fake.updateBuildpackByNameReturns = struct {
		result1 string
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUpdateBuildpackActor) UpdateBuildpackByNameReturnsOnCall(i int, result1 string, result2 v2action.Warnings, result3 error) {
	fake.UpdateBuildpackByNameStub = nil
	if fake.updateBuildpackByNameReturnsOnCall == nil {
		fake.updateBuildpackByNameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.updateBuildpackByNameReturnsOnCall[i] = struct {
		result1 string
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUpdateBuildpackActor) PrepareBuildpackBits(inputPath string, tmpDirPath string, downloader v2action.Downloader) (string, error) {
	fake.prepareBuildpackBitsMutex.Lock()
	ret, specificReturn := fake.prepareBuildpackBitsReturnsOnCall[len(fake.prepareBuildpackBitsArgsForCall)]
	fake.prepareBuildpackBitsArgsForCall = append(fake.prepareBuildpackBitsArgsForCall, struct {
		inputPath  string
		tmpDirPath string
		downloader v2action.Downloader
	}{inputPath, tmpDirPath, downloader})
	fake.recordInvocation("PrepareBuildpackBits", []interface{}{inputPath, tmpDirPath, downloader})
	fake.prepareBuildpackBitsMutex.Unlock()
	if fake.PrepareBuildpackBitsStub != nil {
		return fake.PrepareBuildpackBitsStub(inputPath, tmpDirPath, downloader)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.prepareBuildpackBitsReturns.result1, fake.prepareBuildpackBitsReturns.result2
}

func (fake *FakeUpdateBuildpackActor) PrepareBuildpackBitsCallCount() int {
	fake.prepareBuildpackBitsMutex.RLock()
	defer fake.prepareBuildpackBitsMutex.RUnlock()
	return len(fake.prepareBuildpackBitsArgsForCall)
}

func (fake *FakeUpdateBuildpackActor) PrepareBuildpackBitsArgsForCall(i int) (string, string, v2action.Downloader) {
	fake.prepareBuildpackBitsMutex.RLock()
	defer fake.prepareBuildpackBitsMutex.RUnlock()
	return fake.prepareBuildpackBitsArgsForCall[i].inputPath, fake.prepareBuildpackBitsArgsForCall[i].tmpDirPath, fake.prepareBuildpackBitsArgsForCall[i].downloader
}

func (fake *FakeUpdateBuildpackActor) PrepareBuildpackBitsReturns(result1 string, result2 error) {
	fake.PrepareBuildpackBitsStub = nil
	fake.prepareBuildpackBitsReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUpdateBuildpackActor) PrepareBuildpackBitsReturnsOnCall(i int, result1 string, result2 error) {
	fake.PrepareBuildpackBitsStub = nil
	if fake.prepareBuildpackBitsReturnsOnCall == nil {
		fake.prepareBuildpackBitsReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.prepareBuildpackBitsReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUpdateBuildpackActor) UploadBuildpack(GUID string, path string, progBar v2action.SimpleProgressBar) (v2action.Warnings, error) {
	fake.uploadBuildpackMutex.Lock()
	ret, specificReturn := fake.uploadBuildpackReturnsOnCall[len(fake.uploadBuildpackArgsForCall)]
	fake.uploadBuildpackArgsForCall = append(fake.uploadBuildpackArgsForCall, struct {
		GUID    string
		path    string
		progBar v2action.SimpleProgressBar
	}{GUID, path, progBar})
	fake.recordInvocation("UploadBuildpack", []interface{}{GUID, path, progBar})
	fake.uploadBuildpackMutex.Unlock()
	if fake.UploadBuildpackStub != nil {
		return fake.UploadBuildpackStub(GUID, path, progBar)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.uploadBuildpackReturns.result1, fake.uploadBuildpackReturns.result2
}

func (fake *FakeUpdateBuildpackActor) UploadBuildpackCallCount() int {
	fake.uploadBuildpackMutex.RLock()
	defer fake.uploadBuildpackMutex.RUnlock()
	return len(fake.uploadBuildpackArgsForCall)
}

func (fake *FakeUpdateBuildpackActor) UploadBuildpackArgsForCall(i int) (string, string, v2action.SimpleProgressBar) {
	fake.uploadBuildpackMutex.RLock()
	defer fake.uploadBuildpackMutex.RUnlock()
	return fake.uploadBuildpackArgsForCall[i].GUID, fake.uploadBuildpackArgsForCall[i].path, fake.uploadBuildpackArgsForCall[i].progBar
}

func (fake *FakeUpdateBuildpackActor) UploadBuildpackReturns(result1 v2action.Warnings, result2 error) {
	fake.UploadBuildpackStub = nil
	fake.uploadBuildpackReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUpdateBuildpackActor) UploadBuildpackReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.UploadBuildpackStub = nil
	if fake.uploadBuildpackReturnsOnCall == nil {
		fake.uploadBuildpackReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.uploadBuildpackReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUpdateBuildpackActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.updateBuildpackByNameMutex.RLock()
	defer fake.updateBuildpackByNameMutex.RUnlock()
	fake.prepareBuildpackBitsMutex.RLock()
	defer fake.prepareBuildpackBitsMutex.RUnlock()
	fake.uploadBuildpackMutex.RLock()
	defer fake.uploadBuildpackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUpdateBuildpackActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.UpdateBuildpackActor = new(FakeUpdateBuildpackActor)
