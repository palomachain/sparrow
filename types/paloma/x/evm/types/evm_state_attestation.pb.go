// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: paloma/evm/evm_state_attestation.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_gogo_protobuf_types "github.com/cosmos/gogoproto/types"
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

type ValidatorBalancesAttestation struct {
	HexAddresses  []string                                        `protobuf:"bytes,1,rep,name=hexAddresses,proto3" json:"hexAddresses,omitempty"`
	ValAddresses  []github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,2,rep,name=valAddresses,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"valAddresses,omitempty"`
	FromBlockTime time.Time                                       `protobuf:"bytes,3,opt,name=fromBlockTime,proto3,stdtime" json:"fromBlockTime"`
}

func (m *ValidatorBalancesAttestation) Reset()         { *m = ValidatorBalancesAttestation{} }
func (m *ValidatorBalancesAttestation) String() string { return proto.CompactTextString(m) }
func (*ValidatorBalancesAttestation) ProtoMessage()    {}
func (*ValidatorBalancesAttestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6f19115e0bc1a8, []int{0}
}
func (m *ValidatorBalancesAttestation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorBalancesAttestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorBalancesAttestation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorBalancesAttestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorBalancesAttestation.Merge(m, src)
}
func (m *ValidatorBalancesAttestation) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorBalancesAttestation) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorBalancesAttestation.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorBalancesAttestation proto.InternalMessageInfo

func (m *ValidatorBalancesAttestation) GetHexAddresses() []string {
	if m != nil {
		return m.HexAddresses
	}
	return nil
}

func (m *ValidatorBalancesAttestation) GetValAddresses() []github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValAddresses
	}
	return nil
}

func (m *ValidatorBalancesAttestation) GetFromBlockTime() time.Time {
	if m != nil {
		return m.FromBlockTime
	}
	return time.Time{}
}

type ValidatorBalancesAttestationRes struct {
	BlockHeight uint64   `protobuf:"varint,1,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	Balances    []string `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances,omitempty"`
}

func (m *ValidatorBalancesAttestationRes) Reset()         { *m = ValidatorBalancesAttestationRes{} }
func (m *ValidatorBalancesAttestationRes) String() string { return proto.CompactTextString(m) }
func (*ValidatorBalancesAttestationRes) ProtoMessage()    {}
func (*ValidatorBalancesAttestationRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6f19115e0bc1a8, []int{1}
}
func (m *ValidatorBalancesAttestationRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorBalancesAttestationRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorBalancesAttestationRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorBalancesAttestationRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorBalancesAttestationRes.Merge(m, src)
}
func (m *ValidatorBalancesAttestationRes) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorBalancesAttestationRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorBalancesAttestationRes.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorBalancesAttestationRes proto.InternalMessageInfo

func (m *ValidatorBalancesAttestationRes) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *ValidatorBalancesAttestationRes) GetBalances() []string {
	if m != nil {
		return m.Balances
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidatorBalancesAttestation)(nil), "palomachain.paloma.evm.ValidatorBalancesAttestation")
	proto.RegisterType((*ValidatorBalancesAttestationRes)(nil), "palomachain.paloma.evm.ValidatorBalancesAttestationRes")
}

func init() {
	proto.RegisterFile("paloma/evm/evm_state_attestation.proto", fileDescriptor_1f6f19115e0bc1a8)
}

var fileDescriptor_1f6f19115e0bc1a8 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x3b, 0x1f, 0x5f, 0x0c, 0x0c, 0xb8, 0x69, 0x8c, 0x21, 0x8d, 0x69, 0x1b, 0x16, 0xa6,
	0x2e, 0x98, 0x46, 0x7d, 0x02, 0xea, 0xc6, 0xb8, 0x6c, 0x90, 0x85, 0x1b, 0x32, 0x6d, 0x87, 0x76,
	0x42, 0x87, 0xdb, 0x74, 0x06, 0x82, 0x6f, 0xc1, 0x63, 0xb1, 0x64, 0xe9, 0x0a, 0x0d, 0xc4, 0x97,
	0x70, 0x65, 0xda, 0xf2, 0x77, 0xe3, 0xa2, 0xe9, 0xb9, 0x37, 0xe7, 0xdc, 0xc9, 0xef, 0x5e, 0x7c,
	0x9b, 0xd1, 0x14, 0x04, 0x75, 0xd9, 0x4c, 0x14, 0xdf, 0x50, 0x2a, 0xaa, 0xd8, 0x90, 0x2a, 0xc5,
	0x0a, 0xc5, 0x61, 0x42, 0xb2, 0x1c, 0x14, 0xe8, 0xd7, 0x95, 0x2f, 0x4c, 0x28, 0x9f, 0x90, 0x4a,
	0x13, 0x36, 0x13, 0xc6, 0x55, 0x0c, 0x31, 0x94, 0x16, 0xb7, 0x50, 0x95, 0xdb, 0xb0, 0x62, 0x80,
	0x38, 0x65, 0x6e, 0x59, 0x05, 0xd3, 0x91, 0xab, 0xb8, 0x28, 0x06, 0x8a, 0xac, 0x32, 0x74, 0xbe,
	0x11, 0xbe, 0x19, 0xd0, 0x94, 0x47, 0x54, 0x41, 0xee, 0xd1, 0x94, 0x4e, 0x42, 0x26, 0x7b, 0xc7,
	0x57, 0xf5, 0x0e, 0x6e, 0x25, 0x6c, 0xde, 0x8b, 0xa2, 0x9c, 0x49, 0xc9, 0x64, 0x1b, 0xd9, 0x35,
	0xa7, 0xe1, 0x9f, 0xf5, 0xf4, 0x57, 0xdc, 0x9a, 0xd1, 0xf4, 0xe8, 0xf9, 0x67, 0xd7, 0x9c, 0x96,
	0x77, 0xff, 0xb3, 0xb6, 0xba, 0x31, 0x57, 0xc9, 0x34, 0x20, 0x21, 0x08, 0x37, 0x04, 0x29, 0x40,
	0xee, 0x7e, 0x5d, 0x19, 0x8d, 0x5d, 0xf5, 0x9e, 0x31, 0x49, 0x06, 0x87, 0xa8, 0x7f, 0x36, 0x46,
	0x7f, 0xc1, 0x97, 0xa3, 0x1c, 0x84, 0x97, 0x42, 0x38, 0xee, 0x73, 0xc1, 0xda, 0x35, 0x1b, 0x39,
	0xcd, 0x07, 0x83, 0x54, 0x50, 0x64, 0x0f, 0x45, 0xfa, 0x7b, 0x28, 0xaf, 0xbe, 0x5c, 0x5b, 0xda,
	0xe2, 0xd3, 0x42, 0xfe, 0x79, 0xb4, 0x33, 0xc4, 0xd6, 0x5f, 0x98, 0x3e, 0x93, 0xba, 0x8d, 0x9b,
	0x41, 0xe1, 0x7f, 0x66, 0x3c, 0x4e, 0x54, 0x1b, 0xd9, 0xc8, 0xf9, 0xef, 0x9f, 0xb6, 0x74, 0x03,
	0xd7, 0x83, 0x5d, 0xb6, 0x64, 0x6c, 0xf8, 0x87, 0xda, 0x7b, 0x5a, 0x6e, 0x4c, 0xb4, 0xda, 0x98,
	0xe8, 0x6b, 0x63, 0xa2, 0xc5, 0xd6, 0xd4, 0x56, 0x5b, 0x53, 0xfb, 0xd8, 0x9a, 0xda, 0xdb, 0xdd,
	0xc9, 0x0e, 0x4e, 0x8e, 0xb7, 0xd3, 0xee, 0xbc, 0x3c, 0x79, 0xb9, 0x8a, 0xe0, 0xa2, 0x44, 0x7a,
	0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x48, 0x3f, 0xbe, 0x7e, 0x0d, 0x02, 0x00, 0x00,
}

func (m *ValidatorBalancesAttestation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorBalancesAttestation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorBalancesAttestation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.FromBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.FromBlockTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEvmStateAttestation(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.ValAddresses) > 0 {
		for iNdEx := len(m.ValAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ValAddresses[iNdEx])
			copy(dAtA[i:], m.ValAddresses[iNdEx])
			i = encodeVarintEvmStateAttestation(dAtA, i, uint64(len(m.ValAddresses[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.HexAddresses) > 0 {
		for iNdEx := len(m.HexAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.HexAddresses[iNdEx])
			copy(dAtA[i:], m.HexAddresses[iNdEx])
			i = encodeVarintEvmStateAttestation(dAtA, i, uint64(len(m.HexAddresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorBalancesAttestationRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorBalancesAttestationRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorBalancesAttestationRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Balances) > 0 {
		for iNdEx := len(m.Balances) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Balances[iNdEx])
			copy(dAtA[i:], m.Balances[iNdEx])
			i = encodeVarintEvmStateAttestation(dAtA, i, uint64(len(m.Balances[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintEvmStateAttestation(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvmStateAttestation(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvmStateAttestation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorBalancesAttestation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.HexAddresses) > 0 {
		for _, s := range m.HexAddresses {
			l = len(s)
			n += 1 + l + sovEvmStateAttestation(uint64(l))
		}
	}
	if len(m.ValAddresses) > 0 {
		for _, b := range m.ValAddresses {
			l = len(b)
			n += 1 + l + sovEvmStateAttestation(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.FromBlockTime)
	n += 1 + l + sovEvmStateAttestation(uint64(l))
	return n
}

func (m *ValidatorBalancesAttestationRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovEvmStateAttestation(uint64(m.BlockHeight))
	}
	if len(m.Balances) > 0 {
		for _, s := range m.Balances {
			l = len(s)
			n += 1 + l + sovEvmStateAttestation(uint64(l))
		}
	}
	return n
}

func sovEvmStateAttestation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvmStateAttestation(x uint64) (n int) {
	return sovEvmStateAttestation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorBalancesAttestation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvmStateAttestation
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
			return fmt.Errorf("proto: ValidatorBalancesAttestation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorBalancesAttestation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HexAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmStateAttestation
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
				return ErrInvalidLengthEvmStateAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvmStateAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HexAddresses = append(m.HexAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValAddresses", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmStateAttestation
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
				return ErrInvalidLengthEvmStateAttestation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvmStateAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValAddresses = append(m.ValAddresses, make([]byte, postIndex-iNdEx))
			copy(m.ValAddresses[len(m.ValAddresses)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmStateAttestation
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
				return ErrInvalidLengthEvmStateAttestation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvmStateAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.FromBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvmStateAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvmStateAttestation
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
func (m *ValidatorBalancesAttestationRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvmStateAttestation
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
			return fmt.Errorf("proto: ValidatorBalancesAttestationRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorBalancesAttestationRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmStateAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balances", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmStateAttestation
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
				return ErrInvalidLengthEvmStateAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvmStateAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balances = append(m.Balances, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvmStateAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvmStateAttestation
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
func skipEvmStateAttestation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvmStateAttestation
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
					return 0, ErrIntOverflowEvmStateAttestation
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
					return 0, ErrIntOverflowEvmStateAttestation
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
				return 0, ErrInvalidLengthEvmStateAttestation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvmStateAttestation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvmStateAttestation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvmStateAttestation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvmStateAttestation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvmStateAttestation = fmt.Errorf("proto: unexpected end of group")
)
