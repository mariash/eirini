// This file was generated by counterfeiter
package userfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/commandregistry"
	"code.cloudfoundry.org/cli/cf/commands/user"
	"code.cloudfoundry.org/cli/cf/flags"
	"code.cloudfoundry.org/cli/cf/models"
	"code.cloudfoundry.org/cli/cf/requirements"
)

type FakeOrgRoleSetter struct {
	MetaDataStub        func() commandregistry.CommandMetadata
	metaDataMutex       sync.RWMutex
	metaDataArgsForCall []struct{}
	metaDataReturns     struct {
		result1 commandregistry.CommandMetadata
	}
	SetDependencyStub        func(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command
	setDependencyMutex       sync.RWMutex
	setDependencyArgsForCall []struct {
		deps       commandregistry.Dependency
		pluginCall bool
	}
	setDependencyReturns struct {
		result1 commandregistry.Command
	}
	RequirementsStub        func(requirementsFactory requirements.Factory, context flags.FlagContext) ([]requirements.Requirement, error)
	requirementsMutex       sync.RWMutex
	requirementsArgsForCall []struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}
	requirementsReturns struct {
		result1 []requirements.Requirement
		result2 error
	}
	ExecuteStub        func(context flags.FlagContext) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		context flags.FlagContext
	}
	executeReturns struct {
		result1 error
	}
	SetOrgRoleStub        func(orgGUID string, role models.Role, userGUID, userName string) error
	setOrgRoleMutex       sync.RWMutex
	setOrgRoleArgsForCall []struct {
		orgGUID  string
		role     models.Role
		userGUID string
		userName string
	}
	setOrgRoleReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrgRoleSetter) MetaData() commandregistry.CommandMetadata {
	fake.metaDataMutex.Lock()
	fake.metaDataArgsForCall = append(fake.metaDataArgsForCall, struct{}{})
	fake.recordInvocation("MetaData", []interface{}{})
	fake.metaDataMutex.Unlock()
	if fake.MetaDataStub != nil {
		return fake.MetaDataStub()
	} else {
		return fake.metaDataReturns.result1
	}
}

func (fake *FakeOrgRoleSetter) MetaDataCallCount() int {
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	return len(fake.metaDataArgsForCall)
}

func (fake *FakeOrgRoleSetter) MetaDataReturns(result1 commandregistry.CommandMetadata) {
	fake.MetaDataStub = nil
	fake.metaDataReturns = struct {
		result1 commandregistry.CommandMetadata
	}{result1}
}

func (fake *FakeOrgRoleSetter) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	fake.setDependencyMutex.Lock()
	fake.setDependencyArgsForCall = append(fake.setDependencyArgsForCall, struct {
		deps       commandregistry.Dependency
		pluginCall bool
	}{deps, pluginCall})
	fake.recordInvocation("SetDependency", []interface{}{deps, pluginCall})
	fake.setDependencyMutex.Unlock()
	if fake.SetDependencyStub != nil {
		return fake.SetDependencyStub(deps, pluginCall)
	} else {
		return fake.setDependencyReturns.result1
	}
}

func (fake *FakeOrgRoleSetter) SetDependencyCallCount() int {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return len(fake.setDependencyArgsForCall)
}

func (fake *FakeOrgRoleSetter) SetDependencyArgsForCall(i int) (commandregistry.Dependency, bool) {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return fake.setDependencyArgsForCall[i].deps, fake.setDependencyArgsForCall[i].pluginCall
}

func (fake *FakeOrgRoleSetter) SetDependencyReturns(result1 commandregistry.Command) {
	fake.SetDependencyStub = nil
	fake.setDependencyReturns = struct {
		result1 commandregistry.Command
	}{result1}
}

func (fake *FakeOrgRoleSetter) Requirements(requirementsFactory requirements.Factory, context flags.FlagContext) ([]requirements.Requirement, error) {
	fake.requirementsMutex.Lock()
	fake.requirementsArgsForCall = append(fake.requirementsArgsForCall, struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}{requirementsFactory, context})
	fake.recordInvocation("Requirements", []interface{}{requirementsFactory, context})
	fake.requirementsMutex.Unlock()
	if fake.RequirementsStub != nil {
		return fake.RequirementsStub(requirementsFactory, context)
	} else {
		return fake.requirementsReturns.result1, fake.requirementsReturns.result2
	}
}

func (fake *FakeOrgRoleSetter) RequirementsCallCount() int {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return len(fake.requirementsArgsForCall)
}

func (fake *FakeOrgRoleSetter) RequirementsArgsForCall(i int) (requirements.Factory, flags.FlagContext) {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return fake.requirementsArgsForCall[i].requirementsFactory, fake.requirementsArgsForCall[i].context
}

func (fake *FakeOrgRoleSetter) RequirementsReturns(result1 []requirements.Requirement, result2 error) {
	fake.RequirementsStub = nil
	fake.requirementsReturns = struct {
		result1 []requirements.Requirement
		result2 error
	}{result1, result2}
}

func (fake *FakeOrgRoleSetter) Execute(context flags.FlagContext) error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		context flags.FlagContext
	}{context})
	fake.recordInvocation("Execute", []interface{}{context})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(context)
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeOrgRoleSetter) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeOrgRoleSetter) ExecuteArgsForCall(i int) flags.FlagContext {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].context
}

func (fake *FakeOrgRoleSetter) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrgRoleSetter) SetOrgRole(orgGUID string, role models.Role, userGUID string, userName string) error {
	fake.setOrgRoleMutex.Lock()
	fake.setOrgRoleArgsForCall = append(fake.setOrgRoleArgsForCall, struct {
		orgGUID  string
		role     models.Role
		userGUID string
		userName string
	}{orgGUID, role, userGUID, userName})
	fake.recordInvocation("SetOrgRole", []interface{}{orgGUID, role, userGUID, userName})
	fake.setOrgRoleMutex.Unlock()
	if fake.SetOrgRoleStub != nil {
		return fake.SetOrgRoleStub(orgGUID, role, userGUID, userName)
	} else {
		return fake.setOrgRoleReturns.result1
	}
}

func (fake *FakeOrgRoleSetter) SetOrgRoleCallCount() int {
	fake.setOrgRoleMutex.RLock()
	defer fake.setOrgRoleMutex.RUnlock()
	return len(fake.setOrgRoleArgsForCall)
}

func (fake *FakeOrgRoleSetter) SetOrgRoleArgsForCall(i int) (string, models.Role, string, string) {
	fake.setOrgRoleMutex.RLock()
	defer fake.setOrgRoleMutex.RUnlock()
	return fake.setOrgRoleArgsForCall[i].orgGUID, fake.setOrgRoleArgsForCall[i].role, fake.setOrgRoleArgsForCall[i].userGUID, fake.setOrgRoleArgsForCall[i].userName
}

func (fake *FakeOrgRoleSetter) SetOrgRoleReturns(result1 error) {
	fake.SetOrgRoleStub = nil
	fake.setOrgRoleReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrgRoleSetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.setOrgRoleMutex.RLock()
	defer fake.setOrgRoleMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeOrgRoleSetter) recordInvocation(key string, args []interface{}) {
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

var _ user.OrgRoleSetter = new(FakeOrgRoleSetter)