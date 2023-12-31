// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "noteapp/core/data/model"

	mock "github.com/stretchr/testify/mock"
)

// IUserRepo is an autogenerated mock type for the IUserRepo type
type IUserRepo struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: email
func (_m *IUserRepo) FindByEmail(email string) *model.User {
	ret := _m.Called(email)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	return r0
}

// RegisterUser provides a mock function with given fields: name, email, password, age
func (_m *IUserRepo) RegisterUser(name string, email string, password string, age uint) *model.User {
	ret := _m.Called(name, email, password, age)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string, string, string, uint) *model.User); ok {
		r0 = rf(name, email, password, age)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	return r0
}

type mockConstructorTestingTNewIUserRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserRepo creates a new instance of IUserRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserRepo(t mockConstructorTestingTNewIUserRepo) *IUserRepo {
	mock := &IUserRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
