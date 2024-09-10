// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/littlebugger/tinode4chat/internal/service/entity"
	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *MockUserRepository) CreateUser(ctx context.Context, user entity.User) (primitive.ObjectID, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 primitive.ObjectID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) (primitive.ObjectID, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) primitive.ObjectID); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(primitive.ObjectID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockUserRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user entity.User
func (_e *MockUserRepository_Expecter) CreateUser(ctx interface{}, user interface{}) *MockUserRepository_CreateUser_Call {
	return &MockUserRepository_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, user)}
}

func (_c *MockUserRepository_CreateUser_Call) Run(run func(ctx context.Context, user entity.User)) *MockUserRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.User))
	})
	return _c
}

func (_c *MockUserRepository_CreateUser_Call) Return(_a0 primitive.ObjectID, _a1 error) *MockUserRepository_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_CreateUser_Call) RunAndReturn(run func(context.Context, entity.User) (primitive.ObjectID, error)) *MockUserRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type MockUserRepository_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *MockUserRepository_Expecter) GetUserByEmail(ctx interface{}, email interface{}) *MockUserRepository_GetUserByEmail_Call {
	return &MockUserRepository_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", ctx, email)}
}

func (_c *MockUserRepository_GetUserByEmail_Call) Run(run func(ctx context.Context, email string)) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserRepository_GetUserByEmail_Call) Return(_a0 *entity.User, _a1 error) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetUserByEmail_Call) RunAndReturn(run func(context.Context, string) (*entity.User, error)) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByID provides a mock function with given fields: ctx, id
func (_m *MockUserRepository) GetUserByID(ctx context.Context, id primitive.ObjectID) (*entity.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) (*entity.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *entity.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByID'
type MockUserRepository_GetUserByID_Call struct {
	*mock.Call
}

// GetUserByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id primitive.ObjectID
func (_e *MockUserRepository_Expecter) GetUserByID(ctx interface{}, id interface{}) *MockUserRepository_GetUserByID_Call {
	return &MockUserRepository_GetUserByID_Call{Call: _e.mock.On("GetUserByID", ctx, id)}
}

func (_c *MockUserRepository_GetUserByID_Call) Run(run func(ctx context.Context, id primitive.ObjectID)) *MockUserRepository_GetUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(primitive.ObjectID))
	})
	return _c
}

func (_c *MockUserRepository_GetUserByID_Call) Return(_a0 *entity.User, _a1 error) *MockUserRepository_GetUserByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetUserByID_Call) RunAndReturn(run func(context.Context, primitive.ObjectID) (*entity.User, error)) *MockUserRepository_GetUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: ctx, user
func (_m *MockUserRepository) UpdateUser(ctx context.Context, user entity.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type MockUserRepository_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user entity.User
func (_e *MockUserRepository_Expecter) UpdateUser(ctx interface{}, user interface{}) *MockUserRepository_UpdateUser_Call {
	return &MockUserRepository_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, user)}
}

func (_c *MockUserRepository_UpdateUser_Call) Run(run func(ctx context.Context, user entity.User)) *MockUserRepository_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.User))
	})
	return _c
}

func (_c *MockUserRepository_UpdateUser_Call) Return(_a0 error) *MockUserRepository_UpdateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_UpdateUser_Call) RunAndReturn(run func(context.Context, entity.User) error) *MockUserRepository_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}