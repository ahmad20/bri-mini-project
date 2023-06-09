// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/ahmad20/bri-mini-project/entities"
	mock "github.com/stretchr/testify/mock"
)

// UseCaseInterface is an autogenerated mock type for the UseCaseInterface type
type UseCaseInterface struct {
	mock.Mock
}

type UseCaseInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *UseCaseInterface) EXPECT() *UseCaseInterface_Expecter {
	return &UseCaseInterface_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: _a0
func (_m *UseCaseInterface) Delete(_a0 *entities.Account) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Account) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UseCaseInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type UseCaseInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 *entities.Account
func (_e *UseCaseInterface_Expecter) Delete(_a0 interface{}) *UseCaseInterface_Delete_Call {
	return &UseCaseInterface_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *UseCaseInterface_Delete_Call) Run(run func(_a0 *entities.Account)) *UseCaseInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entities.Account))
	})
	return _c
}

func (_c *UseCaseInterface_Delete_Call) Return(_a0 error) *UseCaseInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCaseInterface_Delete_Call) RunAndReturn(run func(*entities.Account) error) *UseCaseInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAdminsWithConditions provides a mock function with given fields: keyword, page, limit
func (_m *UseCaseInterface) GetAdminsWithConditions(keyword string, page string, limit string) ([]*entities.Account, error) {
	ret := _m.Called(keyword, page, limit)

	var r0 []*entities.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) ([]*entities.Account, error)); ok {
		return rf(keyword, page, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) []*entities.Account); ok {
		r0 = rf(keyword, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(keyword, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseCaseInterface_GetAdminsWithConditions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAdminsWithConditions'
type UseCaseInterface_GetAdminsWithConditions_Call struct {
	*mock.Call
}

// GetAdminsWithConditions is a helper method to define mock.On call
//   - keyword string
//   - page string
//   - limit string
func (_e *UseCaseInterface_Expecter) GetAdminsWithConditions(keyword interface{}, page interface{}, limit interface{}) *UseCaseInterface_GetAdminsWithConditions_Call {
	return &UseCaseInterface_GetAdminsWithConditions_Call{Call: _e.mock.On("GetAdminsWithConditions", keyword, page, limit)}
}

func (_c *UseCaseInterface_GetAdminsWithConditions_Call) Run(run func(keyword string, page string, limit string)) *UseCaseInterface_GetAdminsWithConditions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *UseCaseInterface_GetAdminsWithConditions_Call) Return(_a0 []*entities.Account, _a1 error) *UseCaseInterface_GetAdminsWithConditions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UseCaseInterface_GetAdminsWithConditions_Call) RunAndReturn(run func(string, string, string) ([]*entities.Account, error)) *UseCaseInterface_GetAdminsWithConditions_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *UseCaseInterface) GetAll() ([]*entities.Account, error) {
	ret := _m.Called()

	var r0 []*entities.Account
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*entities.Account, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*entities.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Account)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseCaseInterface_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type UseCaseInterface_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *UseCaseInterface_Expecter) GetAll() *UseCaseInterface_GetAll_Call {
	return &UseCaseInterface_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *UseCaseInterface_GetAll_Call) Run(run func()) *UseCaseInterface_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UseCaseInterface_GetAll_Call) Return(_a0 []*entities.Account, _a1 error) *UseCaseInterface_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UseCaseInterface_GetAll_Call) RunAndReturn(run func() ([]*entities.Account, error)) *UseCaseInterface_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: id
func (_m *UseCaseInterface) GetById(id string) (*entities.Account, error) {
	ret := _m.Called(id)

	var r0 *entities.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.Account, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.Account); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseCaseInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type UseCaseInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - id string
func (_e *UseCaseInterface_Expecter) GetById(id interface{}) *UseCaseInterface_GetById_Call {
	return &UseCaseInterface_GetById_Call{Call: _e.mock.On("GetById", id)}
}

func (_c *UseCaseInterface_GetById_Call) Run(run func(id string)) *UseCaseInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UseCaseInterface_GetById_Call) Return(_a0 *entities.Account, _a1 error) *UseCaseInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UseCaseInterface_GetById_Call) RunAndReturn(run func(string) (*entities.Account, error)) *UseCaseInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetWaitingApproval provides a mock function with given fields:
func (_m *UseCaseInterface) GetWaitingApproval() ([]*entities.Account, error) {
	ret := _m.Called()

	var r0 []*entities.Account
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*entities.Account, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*entities.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Account)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseCaseInterface_GetWaitingApproval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWaitingApproval'
type UseCaseInterface_GetWaitingApproval_Call struct {
	*mock.Call
}

// GetWaitingApproval is a helper method to define mock.On call
func (_e *UseCaseInterface_Expecter) GetWaitingApproval() *UseCaseInterface_GetWaitingApproval_Call {
	return &UseCaseInterface_GetWaitingApproval_Call{Call: _e.mock.On("GetWaitingApproval")}
}

func (_c *UseCaseInterface_GetWaitingApproval_Call) Run(run func()) *UseCaseInterface_GetWaitingApproval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UseCaseInterface_GetWaitingApproval_Call) Return(_a0 []*entities.Account, _a1 error) *UseCaseInterface_GetWaitingApproval_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UseCaseInterface_GetWaitingApproval_Call) RunAndReturn(run func() ([]*entities.Account, error)) *UseCaseInterface_GetWaitingApproval_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: _a0
func (_m *UseCaseInterface) Register(_a0 *entities.Account) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Account) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UseCaseInterface_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type UseCaseInterface_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - _a0 *entities.Account
func (_e *UseCaseInterface_Expecter) Register(_a0 interface{}) *UseCaseInterface_Register_Call {
	return &UseCaseInterface_Register_Call{Call: _e.mock.On("Register", _a0)}
}

func (_c *UseCaseInterface_Register_Call) Run(run func(_a0 *entities.Account)) *UseCaseInterface_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entities.Account))
	})
	return _c
}

func (_c *UseCaseInterface_Register_Call) Return(_a0 error) *UseCaseInterface_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCaseInterface_Register_Call) RunAndReturn(run func(*entities.Account) error) *UseCaseInterface_Register_Call {
	_c.Call.Return(run)
	return _c
}

// SearchByUsername provides a mock function with given fields: username
func (_m *UseCaseInterface) SearchByUsername(username string) (*entities.Account, error) {
	ret := _m.Called(username)

	var r0 *entities.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.Account, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.Account); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseCaseInterface_SearchByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchByUsername'
type UseCaseInterface_SearchByUsername_Call struct {
	*mock.Call
}

// SearchByUsername is a helper method to define mock.On call
//   - username string
func (_e *UseCaseInterface_Expecter) SearchByUsername(username interface{}) *UseCaseInterface_SearchByUsername_Call {
	return &UseCaseInterface_SearchByUsername_Call{Call: _e.mock.On("SearchByUsername", username)}
}

func (_c *UseCaseInterface_SearchByUsername_Call) Run(run func(username string)) *UseCaseInterface_SearchByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UseCaseInterface_SearchByUsername_Call) Return(_a0 *entities.Account, _a1 error) *UseCaseInterface_SearchByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UseCaseInterface_SearchByUsername_Call) RunAndReturn(run func(string) (*entities.Account, error)) *UseCaseInterface_SearchByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateApproval provides a mock function with given fields: status, _a1
func (_m *UseCaseInterface) UpdateApproval(status string, _a1 *entities.Account) error {
	ret := _m.Called(status, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *entities.Account) error); ok {
		r0 = rf(status, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UseCaseInterface_UpdateApproval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateApproval'
type UseCaseInterface_UpdateApproval_Call struct {
	*mock.Call
}

// UpdateApproval is a helper method to define mock.On call
//   - status string
//   - _a1 *entities.Account
func (_e *UseCaseInterface_Expecter) UpdateApproval(status interface{}, _a1 interface{}) *UseCaseInterface_UpdateApproval_Call {
	return &UseCaseInterface_UpdateApproval_Call{Call: _e.mock.On("UpdateApproval", status, _a1)}
}

func (_c *UseCaseInterface_UpdateApproval_Call) Run(run func(status string, _a1 *entities.Account)) *UseCaseInterface_UpdateApproval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*entities.Account))
	})
	return _c
}

func (_c *UseCaseInterface_UpdateApproval_Call) Return(_a0 error) *UseCaseInterface_UpdateApproval_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCaseInterface_UpdateApproval_Call) RunAndReturn(run func(string, *entities.Account) error) *UseCaseInterface_UpdateApproval_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: status, _a1
func (_m *UseCaseInterface) UpdateStatus(status string, _a1 *entities.Account) error {
	ret := _m.Called(status, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *entities.Account) error); ok {
		r0 = rf(status, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UseCaseInterface_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type UseCaseInterface_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - status string
//   - _a1 *entities.Account
func (_e *UseCaseInterface_Expecter) UpdateStatus(status interface{}, _a1 interface{}) *UseCaseInterface_UpdateStatus_Call {
	return &UseCaseInterface_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", status, _a1)}
}

func (_c *UseCaseInterface_UpdateStatus_Call) Run(run func(status string, _a1 *entities.Account)) *UseCaseInterface_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*entities.Account))
	})
	return _c
}

func (_c *UseCaseInterface_UpdateStatus_Call) Return(_a0 error) *UseCaseInterface_UpdateStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCaseInterface_UpdateStatus_Call) RunAndReturn(run func(string, *entities.Account) error) *UseCaseInterface_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUseCaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCaseInterface creates a new instance of UseCaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCaseInterface(t mockConstructorTestingTNewUseCaseInterface) *UseCaseInterface {
	mock := &UseCaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}