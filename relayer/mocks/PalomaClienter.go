// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	chain "github.com/palomachain/pigeon/chain"

	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"

	evmtypes "github.com/palomachain/paloma/x/evm/types"

	gravitytypes "github.com/palomachain/paloma/x/gravity/types"

	mock "github.com/stretchr/testify/mock"

	paloma "github.com/palomachain/pigeon/chain/paloma"

	proto "github.com/cosmos/gogoproto/proto"

	time "time"

	types "github.com/cosmos/cosmos-sdk/x/staking/types"

	valsettypes "github.com/palomachain/paloma/x/valset/types"
)

// PalomaClienter is an autogenerated mock type for the PalomaClienter type
type PalomaClienter struct {
	mock.Mock
}

// AddExternalChainInfo provides a mock function with given fields: ctx, chainInfos
func (_m *PalomaClienter) AddExternalChainInfo(ctx context.Context, chainInfos ...paloma.ChainInfoIn) error {
	_va := make([]interface{}, len(chainInfos))
	for _i := range chainInfos {
		_va[_i] = chainInfos[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...paloma.ChainInfoIn) error); ok {
		r0 = rf(ctx, chainInfos...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddMessageEvidence provides a mock function with given fields: ctx, queueTypeName, messageID, proof
func (_m *PalomaClienter) AddMessageEvidence(ctx context.Context, queueTypeName string, messageID uint64, proof proto.Message) error {
	ret := _m.Called(ctx, queueTypeName, messageID, proof)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, proto.Message) error); ok {
		r0 = rf(ctx, queueTypeName, messageID, proof)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlockHeight provides a mock function with given fields: _a0
func (_m *PalomaClienter) BlockHeight(_a0 context.Context) (int64, error) {
	ret := _m.Called(_a0)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int64, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BroadcastMessageSignatures provides a mock function with given fields: ctx, signatures
func (_m *PalomaClienter) BroadcastMessageSignatures(ctx context.Context, signatures ...paloma.BroadcastMessageSignatureIn) error {
	_va := make([]interface{}, len(signatures))
	for _i := range signatures {
		_va[_i] = signatures[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...paloma.BroadcastMessageSignatureIn) error); ok {
		r0 = rf(ctx, signatures...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetValidator provides a mock function with given fields: ctx
func (_m *PalomaClienter) GetValidator(ctx context.Context) (*types.Validator, error) {
	ret := _m.Called(ctx)

	var r0 *types.Validator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*types.Validator, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *types.Validator); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Validator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValidatorAddress provides a mock function with given fields:
func (_m *PalomaClienter) GetValidatorAddress() cosmos_sdktypes.ValAddress {
	ret := _m.Called()

	var r0 cosmos_sdktypes.ValAddress
	if rf, ok := ret.Get(0).(func() cosmos_sdktypes.ValAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cosmos_sdktypes.ValAddress)
		}
	}

	return r0
}

// GravityConfirmBatches provides a mock function with given fields: ctx, signatures
func (_m *PalomaClienter) GravityConfirmBatches(ctx context.Context, signatures ...chain.SignedGravityOutgoingTxBatch) error {
	_va := make([]interface{}, len(signatures))
	for _i := range signatures {
		_va[_i] = signatures[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...chain.SignedGravityOutgoingTxBatch) error); ok {
		r0 = rf(ctx, signatures...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GravityQueryBatchesForRelaying provides a mock function with given fields: ctx, chainReferenceID
func (_m *PalomaClienter) GravityQueryBatchesForRelaying(ctx context.Context, chainReferenceID string) ([]chain.GravityBatchWithSignatures, error) {
	ret := _m.Called(ctx, chainReferenceID)

	var r0 []chain.GravityBatchWithSignatures
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]chain.GravityBatchWithSignatures, error)); ok {
		return rf(ctx, chainReferenceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []chain.GravityBatchWithSignatures); ok {
		r0 = rf(ctx, chainReferenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chain.GravityBatchWithSignatures)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, chainReferenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GravityQueryLastUnsignedBatch provides a mock function with given fields: ctx, chainReferenceID
func (_m *PalomaClienter) GravityQueryLastUnsignedBatch(ctx context.Context, chainReferenceID string) ([]gravitytypes.OutgoingTxBatch, error) {
	ret := _m.Called(ctx, chainReferenceID)

	var r0 []gravitytypes.OutgoingTxBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]gravitytypes.OutgoingTxBatch, error)); ok {
		return rf(ctx, chainReferenceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []gravitytypes.OutgoingTxBatch); ok {
		r0 = rf(ctx, chainReferenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gravitytypes.OutgoingTxBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, chainReferenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GravityRequestBatch provides a mock function with given fields: ctx, chainReferenceId
func (_m *PalomaClienter) GravityRequestBatch(ctx context.Context, chainReferenceId string) error {
	ret := _m.Called(ctx, chainReferenceId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, chainReferenceId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KeepValidatorAlive provides a mock function with given fields: ctx, appVersion
func (_m *PalomaClienter) KeepValidatorAlive(ctx context.Context, appVersion string) error {
	ret := _m.Called(ctx, appVersion)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, appVersion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryGetEVMChainInfos provides a mock function with given fields: ctx
func (_m *PalomaClienter) QueryGetEVMChainInfos(ctx context.Context) ([]*evmtypes.ChainInfo, error) {
	ret := _m.Called(ctx)

	var r0 []*evmtypes.ChainInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*evmtypes.ChainInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*evmtypes.ChainInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*evmtypes.ChainInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryGetEVMValsetByID provides a mock function with given fields: ctx, id, chainID
func (_m *PalomaClienter) QueryGetEVMValsetByID(ctx context.Context, id uint64, chainID string) (*evmtypes.Valset, error) {
	ret := _m.Called(ctx, id, chainID)

	var r0 *evmtypes.Valset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) (*evmtypes.Valset, error)); ok {
		return rf(ctx, id, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) *evmtypes.Valset); ok {
		r0 = rf(ctx, id, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*evmtypes.Valset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, string) error); ok {
		r1 = rf(ctx, id, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryGetSnapshotByID provides a mock function with given fields: ctx, id
func (_m *PalomaClienter) QueryGetSnapshotByID(ctx context.Context, id uint64) (*valsettypes.Snapshot, error) {
	ret := _m.Called(ctx, id)

	var r0 *valsettypes.Snapshot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*valsettypes.Snapshot, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *valsettypes.Snapshot); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*valsettypes.Snapshot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryGetValidatorAliveUntil provides a mock function with given fields: ctx
func (_m *PalomaClienter) QueryGetValidatorAliveUntil(ctx context.Context) (time.Time, error) {
	ret := _m.Called(ctx)

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (time.Time, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) time.Time); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryMessagesForAttesting provides a mock function with given fields: ctx, queueTypeName
func (_m *PalomaClienter) QueryMessagesForAttesting(ctx context.Context, queueTypeName string) ([]chain.MessageWithSignatures, error) {
	ret := _m.Called(ctx, queueTypeName)

	var r0 []chain.MessageWithSignatures
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]chain.MessageWithSignatures, error)); ok {
		return rf(ctx, queueTypeName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []chain.MessageWithSignatures); ok {
		r0 = rf(ctx, queueTypeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chain.MessageWithSignatures)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, queueTypeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryMessagesForRelaying provides a mock function with given fields: ctx, queueTypeName
func (_m *PalomaClienter) QueryMessagesForRelaying(ctx context.Context, queueTypeName string) ([]chain.MessageWithSignatures, error) {
	ret := _m.Called(ctx, queueTypeName)

	var r0 []chain.MessageWithSignatures
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]chain.MessageWithSignatures, error)); ok {
		return rf(ctx, queueTypeName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []chain.MessageWithSignatures); ok {
		r0 = rf(ctx, queueTypeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chain.MessageWithSignatures)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, queueTypeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryMessagesForSigning provides a mock function with given fields: ctx, queueTypeName
func (_m *PalomaClienter) QueryMessagesForSigning(ctx context.Context, queueTypeName string) ([]chain.QueuedMessage, error) {
	ret := _m.Called(ctx, queueTypeName)

	var r0 []chain.QueuedMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]chain.QueuedMessage, error)); ok {
		return rf(ctx, queueTypeName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []chain.QueuedMessage); ok {
		r0 = rf(ctx, queueTypeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chain.QueuedMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, queueTypeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryValidatorInfo provides a mock function with given fields: ctx
func (_m *PalomaClienter) QueryValidatorInfo(ctx context.Context) ([]*valsettypes.ExternalChainInfo, error) {
	ret := _m.Called(ctx)

	var r0 []*valsettypes.ExternalChainInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*valsettypes.ExternalChainInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*valsettypes.ExternalChainInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*valsettypes.ExternalChainInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetErrorData provides a mock function with given fields: ctx, queueTypeName, messageID, data
func (_m *PalomaClienter) SetErrorData(ctx context.Context, queueTypeName string, messageID uint64, data []byte) error {
	ret := _m.Called(ctx, queueTypeName, messageID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, []byte) error); ok {
		r0 = rf(ctx, queueTypeName, messageID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPublicAccessData provides a mock function with given fields: ctx, queueTypeName, messageID, data
func (_m *PalomaClienter) SetPublicAccessData(ctx context.Context, queueTypeName string, messageID uint64, data []byte) error {
	ret := _m.Called(ctx, queueTypeName, messageID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, []byte) error); ok {
		r0 = rf(ctx, queueTypeName, messageID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPalomaClienter interface {
	mock.TestingT
	Cleanup(func())
}

// NewPalomaClienter creates a new instance of PalomaClienter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPalomaClienter(t mockConstructorTestingTNewPalomaClienter) *PalomaClienter {
	mock := &PalomaClienter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
