// Code generated by mockery v2.43.1. DO NOT EDIT.

package evm

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	mock "github.com/stretchr/testify/mock"
)

// mockMevClient is an autogenerated mock type for the mevClient type
type mockMevClient struct {
	mock.Mock
}

// Relay provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockMevClient) Relay(_a0 context.Context, _a1 *big.Int, _a2 *types.Transaction) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Relay")
	}

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, *types.Transaction) (common.Hash, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, *types.Transaction) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int, *types.Transaction) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockMevClient creates a new instance of mockMevClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockMevClient(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockMevClient {
	mock := &mockMevClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
