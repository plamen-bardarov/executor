// This file was generated by counterfeiter
package containerstorefakes

import (
	"sync"

	"code.cloudfoundry.org/executor"
	"code.cloudfoundry.org/executor/depot/containerstore"
	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager"
)

type FakeCredManager struct {
	CreateCredDirStub        func(lager.Logger, executor.Container) ([]garden.BindMount, error)
	createCredDirMutex       sync.RWMutex
	createCredDirArgsForCall []struct {
		arg1 lager.Logger
		arg2 executor.Container
	}
	createCredDirReturns struct {
		result1 []garden.BindMount
		result2 error
	}
	GenerateCredsStub        func(lager.Logger, executor.Container) error
	generateCredsMutex       sync.RWMutex
	generateCredsArgsForCall []struct {
		arg1 lager.Logger
		arg2 executor.Container
	}
	generateCredsReturns struct {
		result1 error
	}
	RemoveCredsStub        func(lager.Logger, executor.Container) error
	removeCredsMutex       sync.RWMutex
	removeCredsArgsForCall []struct {
		arg1 lager.Logger
		arg2 executor.Container
	}
	removeCredsReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredManager) CreateCredDir(arg1 lager.Logger, arg2 executor.Container) ([]garden.BindMount, error) {
	fake.createCredDirMutex.Lock()
	fake.createCredDirArgsForCall = append(fake.createCredDirArgsForCall, struct {
		arg1 lager.Logger
		arg2 executor.Container
	}{arg1, arg2})
	fake.recordInvocation("CreateCredDir", []interface{}{arg1, arg2})
	fake.createCredDirMutex.Unlock()
	if fake.CreateCredDirStub != nil {
		return fake.CreateCredDirStub(arg1, arg2)
	} else {
		return fake.createCredDirReturns.result1, fake.createCredDirReturns.result2
	}
}

func (fake *FakeCredManager) CreateCredDirCallCount() int {
	fake.createCredDirMutex.RLock()
	defer fake.createCredDirMutex.RUnlock()
	return len(fake.createCredDirArgsForCall)
}

func (fake *FakeCredManager) CreateCredDirArgsForCall(i int) (lager.Logger, executor.Container) {
	fake.createCredDirMutex.RLock()
	defer fake.createCredDirMutex.RUnlock()
	return fake.createCredDirArgsForCall[i].arg1, fake.createCredDirArgsForCall[i].arg2
}

func (fake *FakeCredManager) CreateCredDirReturns(result1 []garden.BindMount, result2 error) {
	fake.CreateCredDirStub = nil
	fake.createCredDirReturns = struct {
		result1 []garden.BindMount
		result2 error
	}{result1, result2}
}

func (fake *FakeCredManager) GenerateCreds(arg1 lager.Logger, arg2 executor.Container) error {
	fake.generateCredsMutex.Lock()
	fake.generateCredsArgsForCall = append(fake.generateCredsArgsForCall, struct {
		arg1 lager.Logger
		arg2 executor.Container
	}{arg1, arg2})
	fake.recordInvocation("GenerateCreds", []interface{}{arg1, arg2})
	fake.generateCredsMutex.Unlock()
	if fake.GenerateCredsStub != nil {
		return fake.GenerateCredsStub(arg1, arg2)
	} else {
		return fake.generateCredsReturns.result1
	}
}

func (fake *FakeCredManager) GenerateCredsCallCount() int {
	fake.generateCredsMutex.RLock()
	defer fake.generateCredsMutex.RUnlock()
	return len(fake.generateCredsArgsForCall)
}

func (fake *FakeCredManager) GenerateCredsArgsForCall(i int) (lager.Logger, executor.Container) {
	fake.generateCredsMutex.RLock()
	defer fake.generateCredsMutex.RUnlock()
	return fake.generateCredsArgsForCall[i].arg1, fake.generateCredsArgsForCall[i].arg2
}

func (fake *FakeCredManager) GenerateCredsReturns(result1 error) {
	fake.GenerateCredsStub = nil
	fake.generateCredsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredManager) RemoveCreds(arg1 lager.Logger, arg2 executor.Container) error {
	fake.removeCredsMutex.Lock()
	fake.removeCredsArgsForCall = append(fake.removeCredsArgsForCall, struct {
		arg1 lager.Logger
		arg2 executor.Container
	}{arg1, arg2})
	fake.recordInvocation("RemoveCreds", []interface{}{arg1, arg2})
	fake.removeCredsMutex.Unlock()
	if fake.RemoveCredsStub != nil {
		return fake.RemoveCredsStub(arg1, arg2)
	} else {
		return fake.removeCredsReturns.result1
	}
}

func (fake *FakeCredManager) RemoveCredsCallCount() int {
	fake.removeCredsMutex.RLock()
	defer fake.removeCredsMutex.RUnlock()
	return len(fake.removeCredsArgsForCall)
}

func (fake *FakeCredManager) RemoveCredsArgsForCall(i int) (lager.Logger, executor.Container) {
	fake.removeCredsMutex.RLock()
	defer fake.removeCredsMutex.RUnlock()
	return fake.removeCredsArgsForCall[i].arg1, fake.removeCredsArgsForCall[i].arg2
}

func (fake *FakeCredManager) RemoveCredsReturns(result1 error) {
	fake.RemoveCredsStub = nil
	fake.removeCredsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCredDirMutex.RLock()
	defer fake.createCredDirMutex.RUnlock()
	fake.generateCredsMutex.RLock()
	defer fake.generateCredsMutex.RUnlock()
	fake.removeCredsMutex.RLock()
	defer fake.removeCredsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCredManager) recordInvocation(key string, args []interface{}) {
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

var _ containerstore.CredManager = new(FakeCredManager)
