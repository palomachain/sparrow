// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: paloma/valset/snapshot.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ValidatorState int32

const (
	ValidatorState_NONE   ValidatorState = 0
	ValidatorState_ACTIVE ValidatorState = 1
	ValidatorState_JAILED ValidatorState = 2
)

var ValidatorState_name = map[int32]string{
	0: "NONE",
	1: "ACTIVE",
	2: "JAILED",
}

var ValidatorState_value = map[string]int32{
	"NONE":   0,
	"ACTIVE": 1,
	"JAILED": 2,
}

func (x ValidatorState) String() string {
	return proto.EnumName(ValidatorState_name, int32(x))
}

func (ValidatorState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5b64fa25356fdf4d, []int{0}
}

type Validator struct {
	ShareCount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=shareCount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"shareCount"`
	// TODO: make this ed25519 pub key instead of bytes
	PubKey             []byte               `protobuf:"bytes,2,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	State              ValidatorState       `protobuf:"varint,3,opt,name=state,proto3,enum=volumefi.paloma.valset.ValidatorState" json:"state,omitempty"`
	ExternalChainInfos []*ExternalChainInfo `protobuf:"bytes,4,rep,name=externalChainInfos,proto3" json:"externalChainInfos,omitempty"`
	Address            string               `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	SignerAddress      string               `protobuf:"bytes,6,opt,name=signerAddress,proto3" json:"signerAddress,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b64fa25356fdf4d, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *Validator) GetState() ValidatorState {
	if m != nil {
		return m.State
	}
	return ValidatorState_NONE
}

func (m *Validator) GetExternalChainInfos() []*ExternalChainInfo {
	if m != nil {
		return m.ExternalChainInfos
	}
	return nil
}

func (m *Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Validator) GetSignerAddress() string {
	if m != nil {
		return m.SignerAddress
	}
	return ""
}

type Snapshot struct {
	Validators  []Validator                            `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators"`
	Height      int64                                  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	TotalShares github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=totalShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"totalShares"`
	CreatedAt   time.Time                              `protobuf:"bytes,4,opt,name=createdAt,proto3,stdtime" json:"createdAt"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b64fa25356fdf4d, []int{1}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetValidators() []Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *Snapshot) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Snapshot) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

type ExternalChainInfo struct {
	ID      uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ChainID string `protobuf:"bytes,2,opt,name=chainID,proto3" json:"chainID,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *ExternalChainInfo) Reset()         { *m = ExternalChainInfo{} }
func (m *ExternalChainInfo) String() string { return proto.CompactTextString(m) }
func (*ExternalChainInfo) ProtoMessage()    {}
func (*ExternalChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b64fa25356fdf4d, []int{2}
}
func (m *ExternalChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExternalChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExternalChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExternalChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalChainInfo.Merge(m, src)
}
func (m *ExternalChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *ExternalChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalChainInfo proto.InternalMessageInfo

func (m *ExternalChainInfo) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ExternalChainInfo) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *ExternalChainInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterEnum("volumefi.paloma.valset.ValidatorState", ValidatorState_name, ValidatorState_value)
	proto.RegisterType((*Validator)(nil), "volumefi.paloma.valset.Validator")
	proto.RegisterType((*Snapshot)(nil), "volumefi.paloma.valset.Snapshot")
	proto.RegisterType((*ExternalChainInfo)(nil), "volumefi.paloma.valset.ExternalChainInfo")
}

func init() { proto.RegisterFile("paloma/valset/snapshot.proto", fileDescriptor_5b64fa25356fdf4d) }

var fileDescriptor_5b64fa25356fdf4d = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0x3a, 0x69, 0x68, 0x36, 0x10, 0x85, 0x15, 0xaa, 0x4c, 0x84, 0x9c, 0x10, 0xa1, 0x2a,
	0x20, 0x58, 0xab, 0xe1, 0xca, 0x25, 0x5f, 0x20, 0x03, 0x0a, 0xc8, 0xa9, 0x8a, 0xe0, 0x82, 0xd6,
	0xc9, 0xc6, 0xb6, 0xb0, 0xbd, 0x96, 0x77, 0x13, 0x35, 0x47, 0xfe, 0x41, 0x7f, 0x56, 0x0f, 0x1c,
	0x7a, 0x44, 0x1c, 0x0a, 0x4a, 0xfe, 0x08, 0xf2, 0xae, 0x0d, 0x49, 0x29, 0x42, 0x9c, 0x3c, 0xe3,
	0x37, 0xf3, 0x3c, 0xf3, 0xe6, 0x19, 0xde, 0x8b, 0x49, 0xc0, 0x42, 0x62, 0x2e, 0x49, 0xc0, 0xa9,
	0x30, 0x79, 0x44, 0x62, 0xee, 0x31, 0x81, 0xe3, 0x84, 0x09, 0x86, 0x0e, 0x96, 0x2c, 0x58, 0x84,
	0x74, 0xee, 0x63, 0x55, 0x86, 0x55, 0x59, 0xe3, 0x8e, 0xcb, 0x5c, 0x26, 0x4b, 0xcc, 0x34, 0x52,
	0xd5, 0x8d, 0xbb, 0x53, 0xc6, 0x43, 0xc6, 0x3f, 0x2a, 0x40, 0x25, 0x19, 0xd4, 0x74, 0x19, 0x73,
	0x03, 0x6a, 0xca, 0xcc, 0x59, 0xcc, 0x4d, 0xe1, 0x87, 0x94, 0x0b, 0x12, 0xc6, 0x79, 0xef, 0xd5,
	0x02, 0x12, 0xad, 0x32, 0xc8, 0x50, 0x4c, 0xa6, 0x43, 0x38, 0x35, 0x97, 0x47, 0x0e, 0x15, 0xe4,
	0xc8, 0x9c, 0x32, 0x3f, 0x52, 0x78, 0xfb, 0x8b, 0x06, 0x2b, 0x27, 0x24, 0xf0, 0x67, 0x44, 0xb0,
	0x04, 0x8d, 0x21, 0xe4, 0x1e, 0x49, 0xe8, 0x80, 0x2d, 0x22, 0xa1, 0x83, 0x16, 0xe8, 0xdc, 0xec,
	0xe3, 0xf3, 0xcb, 0x66, 0xe1, 0xdb, 0x65, 0xf3, 0xd0, 0xf5, 0x85, 0xb7, 0x70, 0xf0, 0x94, 0x85,
	0xd9, 0x78, 0xd9, 0xe3, 0x09, 0x9f, 0x7d, 0x32, 0xc5, 0x2a, 0xa6, 0x1c, 0x5b, 0x91, 0xb0, 0xb7,
	0x18, 0xd0, 0x01, 0x2c, 0xc7, 0x0b, 0xe7, 0x15, 0x5d, 0xe9, 0x5a, 0xca, 0x65, 0x67, 0x19, 0x7a,
	0x06, 0xf7, 0xb8, 0x20, 0x82, 0xea, 0xc5, 0x16, 0xe8, 0xd4, 0xba, 0x87, 0xf8, 0x7a, 0xa9, 0xf0,
	0xaf, 0xc9, 0x26, 0x69, 0xb5, 0xad, 0x9a, 0xd0, 0x7b, 0x88, 0xe8, 0xa9, 0xa0, 0x49, 0x44, 0x82,
	0x81, 0x47, 0xfc, 0xc8, 0x8a, 0xe6, 0x8c, 0xeb, 0xa5, 0x56, 0xb1, 0x53, 0xed, 0x3e, 0xfc, 0x1b,
	0xd5, 0xe8, 0x6a, 0x87, 0x7d, 0x0d, 0x09, 0xd2, 0xe1, 0x0d, 0x32, 0x9b, 0x25, 0x94, 0x73, 0x7d,
	0xaf, 0x05, 0x3a, 0x15, 0x3b, 0x4f, 0xd1, 0x03, 0x78, 0x8b, 0xfb, 0x6e, 0x44, 0x93, 0x5e, 0x86,
	0x97, 0x25, 0xbe, 0xfb, 0xb2, 0xfd, 0x59, 0x83, 0xfb, 0x93, 0xcc, 0x06, 0xe8, 0x05, 0x84, 0xcb,
	0x7c, 0x01, 0xae, 0x03, 0x39, 0xdf, 0xfd, 0x7f, 0xae, 0xda, 0x2f, 0xa5, 0x82, 0xdb, 0x5b, 0xad,
	0xa9, 0x8c, 0x1e, 0xf5, 0x5d, 0x4f, 0x48, 0x19, 0x8b, 0x76, 0x96, 0xa1, 0xb7, 0xb0, 0x2a, 0x98,
	0x20, 0xc1, 0x24, 0x55, 0x9c, 0x4b, 0x31, 0xff, 0xff, 0x5e, 0xdb, 0x14, 0xa8, 0x0f, 0x2b, 0xd3,
	0x84, 0x12, 0x41, 0x67, 0x3d, 0xa1, 0x97, 0x5a, 0xa0, 0x53, 0xed, 0x36, 0xb0, 0x72, 0x17, 0xce,
	0xdd, 0x85, 0x8f, 0x73, 0xfb, 0xf5, 0xf7, 0xd3, 0x6f, 0x9d, 0x7d, 0x6f, 0x02, 0xfb, 0x77, 0x5b,
	0xfb, 0x1d, 0xbc, 0xfd, 0x87, 0xd8, 0xa8, 0x06, 0x35, 0x6b, 0x28, 0x1d, 0x55, 0xb2, 0x35, 0x6b,
	0x98, 0x0a, 0x3d, 0x95, 0xe0, 0x50, 0xee, 0x54, 0xb1, 0xf3, 0x74, 0xfb, 0x04, 0xc5, 0x9d, 0x13,
	0x3c, 0xea, 0xc2, 0xda, 0xae, 0x21, 0xd0, 0x3e, 0x2c, 0x8d, 0xdf, 0x8c, 0x47, 0xf5, 0x02, 0x82,
	0xb0, 0xdc, 0x1b, 0x1c, 0x5b, 0x27, 0xa3, 0x3a, 0x48, 0xe3, 0x97, 0x3d, 0xeb, 0xf5, 0x68, 0x58,
	0xd7, 0xfa, 0xcf, 0xcf, 0xd7, 0x06, 0xb8, 0x58, 0x1b, 0xe0, 0xc7, 0xda, 0x00, 0x67, 0x1b, 0xa3,
	0x70, 0xb1, 0x31, 0x0a, 0x5f, 0x37, 0x46, 0xe1, 0xc3, 0xe3, 0x2d, 0x7d, 0xd4, 0x29, 0xe4, 0x04,
	0x59, 0x6c, 0x9e, 0xe6, 0x7f, 0xb5, 0x54, 0xca, 0x29, 0xcb, 0xed, 0x9f, 0xfe, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0xb9, 0x22, 0x97, 0x97, 0xf3, 0x03, 0x00, 0x00,
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SignerAddress) > 0 {
		i -= len(m.SignerAddress)
		copy(dAtA[i:], m.SignerAddress)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.SignerAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ExternalChainInfos) > 0 {
		for iNdEx := len(m.ExternalChainInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExternalChainInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSnapshot(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.State != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.ShareCount.Size()
		i -= size
		if _, err := m.ShareCount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSnapshot(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Snapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSnapshot(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.TotalShares.Size()
		i -= size
		if _, err := m.TotalShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSnapshot(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Height != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSnapshot(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ExternalChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExternalChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExternalChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSnapshot(dAtA []byte, offset int, v uint64) int {
	offset -= sovSnapshot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ShareCount.Size()
	n += 1 + l + sovSnapshot(uint64(l))
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovSnapshot(uint64(m.State))
	}
	if len(m.ExternalChainInfos) > 0 {
		for _, e := range m.ExternalChainInfos {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = len(m.SignerAddress)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovSnapshot(uint64(m.Height))
	}
	l = m.TotalShares.Size()
	n += 1 + l + sovSnapshot(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovSnapshot(uint64(l))
	return n
}

func (m *ExternalChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovSnapshot(uint64(m.ID))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}

func sovSnapshot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSnapshot(x uint64) (n int) {
	return sovSnapshot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareCount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ShareCount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= ValidatorState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalChainInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalChainInfos = append(m.ExternalChainInfos, &ExternalChainInfo{})
			if err := m.ExternalChainInfos[len(m.ExternalChainInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExternalChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExternalChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExternalChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSnapshot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSnapshot
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSnapshot
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSnapshot
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSnapshot
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSnapshot        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnapshot          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSnapshot = fmt.Errorf("proto: unexpected end of group")
)
