// Code generated by mockery v2.42.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	request "github.com/mattermost/mattermost/server/public/shared/request"
	mock "github.com/stretchr/testify/mock"
)

// AccountMigrationInterface is an autogenerated mock type for the AccountMigrationInterface type
type AccountMigrationInterface struct {
	mock.Mock
}

// MigrateToLdap provides a mock function with given fields: c, fromAuthService, foreignUserFieldNameToMatch, force, dryRun
func (_m *AccountMigrationInterface) MigrateToLdap(c request.CTX, fromAuthService string, foreignUserFieldNameToMatch string, force bool, dryRun bool) *model.AppError {
	ret := _m.Called(c, fromAuthService, foreignUserFieldNameToMatch, force, dryRun)

	if len(ret) == 0 {
		panic("no return value specified for MigrateToLdap")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string, string, bool, bool) *model.AppError); ok {
		r0 = rf(c, fromAuthService, foreignUserFieldNameToMatch, force, dryRun)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// MigrateToSaml provides a mock function with given fields: c, fromAuthService, usersMap, auto, dryRun
func (_m *AccountMigrationInterface) MigrateToSaml(c request.CTX, fromAuthService string, usersMap map[string]string, auto bool, dryRun bool) *model.AppError {
	ret := _m.Called(c, fromAuthService, usersMap, auto, dryRun)

	if len(ret) == 0 {
		panic("no return value specified for MigrateToSaml")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string, map[string]string, bool, bool) *model.AppError); ok {
		r0 = rf(c, fromAuthService, usersMap, auto, dryRun)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// NewAccountMigrationInterface creates a new instance of AccountMigrationInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountMigrationInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountMigrationInterface {
	mock := &AccountMigrationInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
