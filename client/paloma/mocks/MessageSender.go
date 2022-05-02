// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	types "github.com/cosmos/cosmos-sdk/types"
)

// MessageSender is an autogenerated mock type for the MessageSender type
type MessageSender struct {
	mock.Mock
}

// SendMsg provides a mock function with given fields: ctx, msg
func (_m *MessageSender) SendMsg(ctx context.Context, msg types.Msg) (*types.TxResponse, error) {
	ret := _m.Called(ctx, msg)

	var r0 *types.TxResponse
	if rf, ok := ret.Get(0).(func(context.Context, types.Msg) *types.TxResponse); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TxResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.Msg) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMessageSender creates a new instance of MessageSender. It also registers a cleanup function to assert the mocks expectations.
func NewMessageSender(t testing.TB) *MessageSender {
	mock := &MessageSender{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
