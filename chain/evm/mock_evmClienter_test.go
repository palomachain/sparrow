// Code generated by mockery v2.14.0. DO NOT EDIT.

package evm

import (
	big "math/big"

	abi "github.com/ethereum/go-ethereum/accounts/abi"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// mockEvmClienter is an autogenerated mock type for the evmClienter type
type mockEvmClienter struct {
	mock.Mock
}

// DeployContract provides a mock function with given fields: ctx, chainID, contractAbi, bytecode, constructorInput
func (_m *mockEvmClienter) DeployContract(ctx context.Context, chainID *big.Int, contractAbi abi.ABI, bytecode []byte, constructorInput []byte) (common.Address, *types.Transaction, error) {
	ret := _m.Called(ctx, chainID, contractAbi, bytecode, constructorInput)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, abi.ABI, []byte, []byte) common.Address); ok {
		r0 = rf(ctx, chainID, contractAbi, bytecode, constructorInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 *types.Transaction
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int, abi.ABI, []byte, []byte) *types.Transaction); ok {
		r1 = rf(ctx, chainID, contractAbi, bytecode, constructorInput)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Transaction)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *big.Int, abi.ABI, []byte, []byte) error); ok {
		r2 = rf(ctx, chainID, contractAbi, bytecode, constructorInput)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ExecuteSmartContract provides a mock function with given fields: ctx, chainID, contractAbi, addr, method, arguments
func (_m *mockEvmClienter) ExecuteSmartContract(ctx context.Context, chainID *big.Int, contractAbi abi.ABI, addr common.Address, method string, arguments []interface{}) (*types.Transaction, error) {
	ret := _m.Called(ctx, chainID, contractAbi, addr, method, arguments)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, abi.ABI, common.Address, string, []interface{}) *types.Transaction); ok {
		r0 = rf(ctx, chainID, contractAbi, addr, method, arguments)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int, abi.ABI, common.Address, string, []interface{}) error); ok {
		r1 = rf(ctx, chainID, contractAbi, addr, method, arguments)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterLogs provides a mock function with given fields: ctx, fq, currBlockHeight, fn
func (_m *mockEvmClienter) FilterLogs(ctx context.Context, fq ethereum.FilterQuery, currBlockHeight *big.Int, fn func([]types.Log) bool) (bool, error) {
	ret := _m.Called(ctx, fq, currBlockHeight, fn)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, *big.Int, func([]types.Log) bool) bool); ok {
		r0 = rf(ctx, fq, currBlockHeight, fn)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery, *big.Int, func([]types.Log) bool) error); ok {
		r1 = rf(ctx, fq, currBlockHeight, fn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionByHash provides a mock function with given fields: ctx, txHash
func (_m *mockEvmClienter) TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	ret := _m.Called(ctx, txHash)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Transaction); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) bool); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, common.Hash) error); ok {
		r2 = rf(ctx, txHash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTnewMockEvmClienter interface {
	mock.TestingT
	Cleanup(func())
}

// newMockEvmClienter creates a new instance of mockEvmClienter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockEvmClienter(t mockConstructorTestingTnewMockEvmClienter) *mockEvmClienter {
	mock := &mockEvmClienter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
