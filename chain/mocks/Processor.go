// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	chain "github.com/palomachain/pigeon/chain"

	mock "github.com/stretchr/testify/mock"

	queue "github.com/palomachain/pigeon/internal/queue"

	types "github.com/palomachain/paloma/x/gravity/types"
)

// Processor is an autogenerated mock type for the Processor type
type Processor struct {
	mock.Mock
}

// ExternalAccount provides a mock function with given fields:
func (_m *Processor) ExternalAccount() chain.ExternalAccount {
	ret := _m.Called()

	var r0 chain.ExternalAccount
	if rf, ok := ret.Get(0).(func() chain.ExternalAccount); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(chain.ExternalAccount)
	}

	return r0
}

// GetChainReferenceID provides a mock function with given fields:
func (_m *Processor) GetChainReferenceID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GravityRelayBatches provides a mock function with given fields: ctx, batches
func (_m *Processor) GravityRelayBatches(ctx context.Context, batches []chain.GravityBatchWithSignatures) error {
	ret := _m.Called(ctx, batches)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []chain.GravityBatchWithSignatures) error); ok {
		r0 = rf(ctx, batches)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GravitySignBatches provides a mock function with given fields: ctx, batches
func (_m *Processor) GravitySignBatches(ctx context.Context, batches ...types.OutgoingTxBatch) ([]chain.SignedGravityOutgoingTxBatch, error) {
	_va := make([]interface{}, len(batches))
	for _i := range batches {
		_va[_i] = batches[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []chain.SignedGravityOutgoingTxBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...types.OutgoingTxBatch) ([]chain.SignedGravityOutgoingTxBatch, error)); ok {
		return rf(ctx, batches...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...types.OutgoingTxBatch) []chain.SignedGravityOutgoingTxBatch); ok {
		r0 = rf(ctx, batches...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chain.SignedGravityOutgoingTxBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...types.OutgoingTxBatch) error); ok {
		r1 = rf(ctx, batches...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthCheck provides a mock function with given fields: ctx
func (_m *Processor) HealthCheck(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsRightChain provides a mock function with given fields: ctx
func (_m *Processor) IsRightChain(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessMessages provides a mock function with given fields: _a0, _a1, _a2
func (_m *Processor) ProcessMessages(_a0 context.Context, _a1 queue.TypeName, _a2 []chain.MessageWithSignatures) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, queue.TypeName, []chain.MessageWithSignatures) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvideEvidence provides a mock function with given fields: _a0, _a1, _a2
func (_m *Processor) ProvideEvidence(_a0 context.Context, _a1 queue.TypeName, _a2 []chain.MessageWithSignatures) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, queue.TypeName, []chain.MessageWithSignatures) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignMessages provides a mock function with given fields: ctx, messages
func (_m *Processor) SignMessages(ctx context.Context, messages ...chain.QueuedMessage) ([]chain.SignedQueuedMessage, error) {
	_va := make([]interface{}, len(messages))
	for _i := range messages {
		_va[_i] = messages[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []chain.SignedQueuedMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...chain.QueuedMessage) ([]chain.SignedQueuedMessage, error)); ok {
		return rf(ctx, messages...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...chain.QueuedMessage) []chain.SignedQueuedMessage); ok {
		r0 = rf(ctx, messages...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chain.SignedQueuedMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...chain.QueuedMessage) error); ok {
		r1 = rf(ctx, messages...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SupportedQueues provides a mock function with given fields:
func (_m *Processor) SupportedQueues() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

type mockConstructorTestingTNewProcessor interface {
	mock.TestingT
	Cleanup(func())
}

// NewProcessor creates a new instance of Processor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProcessor(t mockConstructorTestingTNewProcessor) *Processor {
	mock := &Processor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
