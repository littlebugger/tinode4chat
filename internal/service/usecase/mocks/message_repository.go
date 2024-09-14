// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/littlebugger/tinode4chat/internal/service/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockMessageRepository is an autogenerated mock type for the MessageRepository type
type MockMessageRepository struct {
	mock.Mock
}

type MockMessageRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMessageRepository) EXPECT() *MockMessageRepository_Expecter {
	return &MockMessageRepository_Expecter{mock: &_m.Mock}
}

// CheckIfUserInRoom provides a mock function with given fields: ctx, roomID, userID
func (_m *MockMessageRepository) CheckIfUserInRoom(ctx context.Context, roomID string, userID string) (bool, error) {
	ret := _m.Called(ctx, roomID, userID)

	if len(ret) == 0 {
		panic("no return value specified for CheckIfUserInRoom")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, roomID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, roomID, userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, roomID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMessageRepository_CheckIfUserInRoom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckIfUserInRoom'
type MockMessageRepository_CheckIfUserInRoom_Call struct {
	*mock.Call
}

// CheckIfUserInRoom is a helper method to define mock.On call
//   - ctx context.Context
//   - roomID string
//   - userID string
func (_e *MockMessageRepository_Expecter) CheckIfUserInRoom(ctx interface{}, roomID interface{}, userID interface{}) *MockMessageRepository_CheckIfUserInRoom_Call {
	return &MockMessageRepository_CheckIfUserInRoom_Call{Call: _e.mock.On("CheckIfUserInRoom", ctx, roomID, userID)}
}

func (_c *MockMessageRepository_CheckIfUserInRoom_Call) Run(run func(ctx context.Context, roomID string, userID string)) *MockMessageRepository_CheckIfUserInRoom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockMessageRepository_CheckIfUserInRoom_Call) Return(_a0 bool, _a1 error) *MockMessageRepository_CheckIfUserInRoom_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMessageRepository_CheckIfUserInRoom_Call) RunAndReturn(run func(context.Context, string, string) (bool, error)) *MockMessageRepository_CheckIfUserInRoom_Call {
	_c.Call.Return(run)
	return _c
}

// CreateMessage provides a mock function with given fields: ctx, message
func (_m *MockMessageRepository) CreateMessage(ctx context.Context, message entity.Message) (*string, error) {
	ret := _m.Called(ctx, message)

	if len(ret) == 0 {
		panic("no return value specified for CreateMessage")
	}

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Message) (*string, error)); ok {
		return rf(ctx, message)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Message) *string); ok {
		r0 = rf(ctx, message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Message) error); ok {
		r1 = rf(ctx, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMessageRepository_CreateMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMessage'
type MockMessageRepository_CreateMessage_Call struct {
	*mock.Call
}

// CreateMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - message entity.Message
func (_e *MockMessageRepository_Expecter) CreateMessage(ctx interface{}, message interface{}) *MockMessageRepository_CreateMessage_Call {
	return &MockMessageRepository_CreateMessage_Call{Call: _e.mock.On("CreateMessage", ctx, message)}
}

func (_c *MockMessageRepository_CreateMessage_Call) Run(run func(ctx context.Context, message entity.Message)) *MockMessageRepository_CreateMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.Message))
	})
	return _c
}

func (_c *MockMessageRepository_CreateMessage_Call) Return(_a0 *string, _a1 error) *MockMessageRepository_CreateMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMessageRepository_CreateMessage_Call) RunAndReturn(run func(context.Context, entity.Message) (*string, error)) *MockMessageRepository_CreateMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetMessagesByChatRoom provides a mock function with given fields: ctx, roomID
func (_m *MockMessageRepository) GetMessagesByChatRoom(ctx context.Context, roomID string) ([]entity.Message, error) {
	ret := _m.Called(ctx, roomID)

	if len(ret) == 0 {
		panic("no return value specified for GetMessagesByChatRoom")
	}

	var r0 []entity.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.Message, error)); ok {
		return rf(ctx, roomID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Message); ok {
		r0 = rf(ctx, roomID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, roomID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMessageRepository_GetMessagesByChatRoom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMessagesByChatRoom'
type MockMessageRepository_GetMessagesByChatRoom_Call struct {
	*mock.Call
}

// GetMessagesByChatRoom is a helper method to define mock.On call
//   - ctx context.Context
//   - roomID string
func (_e *MockMessageRepository_Expecter) GetMessagesByChatRoom(ctx interface{}, roomID interface{}) *MockMessageRepository_GetMessagesByChatRoom_Call {
	return &MockMessageRepository_GetMessagesByChatRoom_Call{Call: _e.mock.On("GetMessagesByChatRoom", ctx, roomID)}
}

func (_c *MockMessageRepository_GetMessagesByChatRoom_Call) Run(run func(ctx context.Context, roomID string)) *MockMessageRepository_GetMessagesByChatRoom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockMessageRepository_GetMessagesByChatRoom_Call) Return(_a0 []entity.Message, _a1 error) *MockMessageRepository_GetMessagesByChatRoom_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMessageRepository_GetMessagesByChatRoom_Call) RunAndReturn(run func(context.Context, string) ([]entity.Message, error)) *MockMessageRepository_GetMessagesByChatRoom_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMessageRepository creates a new instance of MockMessageRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMessageRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMessageRepository {
	mock := &MockMessageRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
