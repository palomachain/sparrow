// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	context "context"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	keyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	types "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"
)

// IonClient is an autogenerated mock type for the IonClient type
type IonClient struct {
	mock.Mock
}

// DecodeBech32ValAddr provides a mock function with given fields: _a0
func (_m *IonClient) DecodeBech32ValAddr(_a0 string) (types.ValAddress, error) {
	ret := _m.Called(_a0)

	var r0 types.ValAddress
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (types.ValAddress, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) types.ValAddress); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ValAddress)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKeybase provides a mock function with given fields:
func (_m *IonClient) GetKeybase() keyring.Keyring {
	ret := _m.Called()

	var r0 keyring.Keyring
	if rf, ok := ret.Get(0).(func() keyring.Keyring); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keyring.Keyring)
		}
	}

	return r0
}

// SendMsg provides a mock function with given fields: ctx, msg, memo
func (_m *IonClient) SendMsg(ctx context.Context, msg types.Msg, memo string) (*types.TxResponse, error) {
	ret := _m.Called(ctx, msg, memo)

	var r0 *types.TxResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Msg, string) (*types.TxResponse, error)); ok {
		return rf(ctx, msg, memo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.Msg, string) *types.TxResponse); ok {
		r0 = rf(ctx, msg, memo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TxResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.Msg, string) error); ok {
		r1 = rf(ctx, msg, memo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetSDKContext provides a mock function with given fields:
func (_m *IonClient) SetSDKContext() {
	_m.Called()
}

// Status provides a mock function with given fields: _a0
func (_m *IonClient) Status(_a0 context.Context) (*coretypes.ResultStatus, error) {
	ret := _m.Called(_a0)

	var r0 *coretypes.ResultStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*coretypes.ResultStatus, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *coretypes.ResultStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIonClient creates a new instance of IonClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIonClient(t interface {
	mock.TestingT
	Cleanup(func())
},
) *IonClient {
	mock := &IonClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
