// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	chain "github.com/palomachain/pigeon/chain"

	config "github.com/palomachain/pigeon/config"

	mev "github.com/palomachain/pigeon/internal/mev"

	mock "github.com/stretchr/testify/mock"
)

// EvmFactorier is an autogenerated mock type for the EvmFactorier type
type EvmFactorier struct {
	mock.Mock
}

// Build provides a mock function with given fields: cfg, chainReferenceID, smartContractID, smartContractABIJson, smartContractAddress, chainID, blockHeight, blockHeightHash, minOnChainBalance, mevClient
func (_m *EvmFactorier) Build(cfg config.EVM, chainReferenceID string, smartContractID string, smartContractABIJson string, smartContractAddress string, chainID *big.Int, blockHeight int64, blockHeightHash common.Hash, minOnChainBalance *big.Int, mevClient mev.Client) (chain.Processor, error) {
	ret := _m.Called(cfg, chainReferenceID, smartContractID, smartContractABIJson, smartContractAddress, chainID, blockHeight, blockHeightHash, minOnChainBalance, mevClient)

	var r0 chain.Processor
	var r1 error
	if rf, ok := ret.Get(0).(func(config.EVM, string, string, string, string, *big.Int, int64, common.Hash, *big.Int, mev.Client) (chain.Processor, error)); ok {
		return rf(cfg, chainReferenceID, smartContractID, smartContractABIJson, smartContractAddress, chainID, blockHeight, blockHeightHash, minOnChainBalance, mevClient)
	}
	if rf, ok := ret.Get(0).(func(config.EVM, string, string, string, string, *big.Int, int64, common.Hash, *big.Int, mev.Client) chain.Processor); ok {
		r0 = rf(cfg, chainReferenceID, smartContractID, smartContractABIJson, smartContractAddress, chainID, blockHeight, blockHeightHash, minOnChainBalance, mevClient)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chain.Processor)
		}
	}

	if rf, ok := ret.Get(1).(func(config.EVM, string, string, string, string, *big.Int, int64, common.Hash, *big.Int, mev.Client) error); ok {
		r1 = rf(cfg, chainReferenceID, smartContractID, smartContractABIJson, smartContractAddress, chainID, blockHeight, blockHeightHash, minOnChainBalance, mevClient)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEvmFactorier creates a new instance of EvmFactorier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEvmFactorier(t interface {
	mock.TestingT
	Cleanup(func())
}) *EvmFactorier {
	mock := &EvmFactorier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
