// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/palomachain/pigeon/types/paloma/x/valset/types"
	mock "github.com/stretchr/testify/mock"
)

// MsgServer is an autogenerated mock type for the MsgServer type
type MsgServer struct {
	mock.Mock
}

// AddExternalChainInfoForValidator provides a mock function with given fields: _a0, _a1
func (_m *MsgServer) AddExternalChainInfoForValidator(_a0 context.Context, _a1 *types.MsgAddExternalChainInfoForValidator) (*types.MsgAddExternalChainInfoForValidatorResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.MsgAddExternalChainInfoForValidatorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.MsgAddExternalChainInfoForValidator) *types.MsgAddExternalChainInfoForValidatorResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.MsgAddExternalChainInfoForValidatorResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.MsgAddExternalChainInfoForValidator) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMsgServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMsgServer creates a new instance of MsgServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMsgServer(t mockConstructorTestingTNewMsgServer) *MsgServer {
	mock := &MsgServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
