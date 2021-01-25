// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: execution_fee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ExecutionFee struct {
	Name              string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TransactionType   string `protobuf:"bytes,2,opt,name=transaction_type,json=transactionType,proto3" json:"transaction_type,omitempty"`
	ExecutionFee      uint64 `protobuf:"varint,3,opt,name=execution_fee,json=executionFee,proto3" json:"execution_fee,omitempty"`
	FailureFee        uint64 `protobuf:"varint,4,opt,name=failure_fee,json=failureFee,proto3" json:"failure_fee,omitempty"`
	Timeout           uint64 `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	DefaultParameters uint64 `protobuf:"varint,6,opt,name=default_parameters,json=defaultParameters,proto3" json:"default_parameters,omitempty"`
}

func (m *ExecutionFee) Reset()         { *m = ExecutionFee{} }
func (m *ExecutionFee) String() string { return proto.CompactTextString(m) }
func (*ExecutionFee) ProtoMessage()    {}
func (*ExecutionFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_76815e69c9dcfd49, []int{0}
}
func (m *ExecutionFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecutionFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExecutionFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExecutionFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionFee.Merge(m, src)
}
func (m *ExecutionFee) XXX_Size() int {
	return m.Size()
}
func (m *ExecutionFee) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionFee.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionFee proto.InternalMessageInfo

func (m *ExecutionFee) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExecutionFee) GetTransactionType() string {
	if m != nil {
		return m.TransactionType
	}
	return ""
}

func (m *ExecutionFee) GetExecutionFee() uint64 {
	if m != nil {
		return m.ExecutionFee
	}
	return 0
}

func (m *ExecutionFee) GetFailureFee() uint64 {
	if m != nil {
		return m.FailureFee
	}
	return 0
}

func (m *ExecutionFee) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *ExecutionFee) GetDefaultParameters() uint64 {
	if m != nil {
		return m.DefaultParameters
	}
	return 0
}

type MsgSetExecutionFee struct {
	Name              string                                        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TransactionType   string                                        `protobuf:"bytes,2,opt,name=transaction_type,json=transactionType,proto3" json:"transaction_type,omitempty"`
	ExecutionFee      uint64                                        `protobuf:"varint,3,opt,name=execution_fee,json=executionFee,proto3" json:"execution_fee,omitempty"`
	FailureFee        uint64                                        `protobuf:"varint,4,opt,name=failure_fee,json=failureFee,proto3" json:"failure_fee,omitempty"`
	Timeout           uint64                                        `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	DefaultParameters uint64                                        `protobuf:"varint,6,opt,name=default_parameters,json=defaultParameters,proto3" json:"default_parameters,omitempty"`
	Proposer          github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,7,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty"`
}

func (m *MsgSetExecutionFee) Reset()         { *m = MsgSetExecutionFee{} }
func (m *MsgSetExecutionFee) String() string { return proto.CompactTextString(m) }
func (*MsgSetExecutionFee) ProtoMessage()    {}
func (*MsgSetExecutionFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_76815e69c9dcfd49, []int{1}
}
func (m *MsgSetExecutionFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetExecutionFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetExecutionFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetExecutionFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetExecutionFee.Merge(m, src)
}
func (m *MsgSetExecutionFee) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetExecutionFee) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetExecutionFee.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetExecutionFee proto.InternalMessageInfo

func (m *MsgSetExecutionFee) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgSetExecutionFee) GetTransactionType() string {
	if m != nil {
		return m.TransactionType
	}
	return ""
}

func (m *MsgSetExecutionFee) GetExecutionFee() uint64 {
	if m != nil {
		return m.ExecutionFee
	}
	return 0
}

func (m *MsgSetExecutionFee) GetFailureFee() uint64 {
	if m != nil {
		return m.FailureFee
	}
	return 0
}

func (m *MsgSetExecutionFee) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *MsgSetExecutionFee) GetDefaultParameters() uint64 {
	if m != nil {
		return m.DefaultParameters
	}
	return 0
}

func (m *MsgSetExecutionFee) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecutionFee)(nil), "kira.gov.ExecutionFee")
	proto.RegisterType((*MsgSetExecutionFee)(nil), "kira.gov.MsgSetExecutionFee")
}

func init() { proto.RegisterFile("execution_fee.proto", fileDescriptor_76815e69c9dcfd49) }

var fileDescriptor_76815e69c9dcfd49 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x92, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0x29, 0x1f, 0x1f, 0x60, 0xc5, 0xa8, 0xd5, 0x45, 0xe3, 0x62, 0x20, 0x18, 0x13, 0x5c,
	0x30, 0x13, 0xe3, 0x03, 0x18, 0x30, 0xba, 0x31, 0x24, 0x06, 0x5d, 0xb9, 0x21, 0x65, 0xe6, 0x32,
	0x36, 0x30, 0xdc, 0xa6, 0xed, 0x10, 0x78, 0x0b, 0x5f, 0xc5, 0xb7, 0x70, 0xc9, 0xd2, 0x85, 0x31,
	0x06, 0xde, 0xc2, 0x95, 0xa1, 0xfc, 0x09, 0xbe, 0x82, 0xab, 0xde, 0x9e, 0xdf, 0xe9, 0x69, 0x4f,
	0x52, 0x7a, 0x04, 0x63, 0x08, 0x53, 0x2b, 0x71, 0xd8, 0xe9, 0x01, 0xf8, 0x4a, 0xa3, 0x45, 0x56,
	0xec, 0x4b, 0x2d, 0xfc, 0x18, 0x47, 0x27, 0xc7, 0x31, 0xc6, 0xe8, 0xc4, 0x60, 0x31, 0x2d, 0x79,
	0xf5, 0x83, 0xd0, 0xd2, 0xcd, 0xfa, 0xdc, 0x2d, 0x00, 0x63, 0x34, 0x37, 0x14, 0x09, 0x70, 0x52,
	0x21, 0xb5, 0x9d, 0xb6, 0x9b, 0xd9, 0x39, 0x3d, 0xb0, 0x5a, 0x0c, 0x8d, 0x08, 0x5d, 0xba, 0x9d,
	0x28, 0xe0, 0x59, 0xc7, 0xf7, 0xb7, 0xf4, 0xc7, 0x89, 0x02, 0x76, 0x4a, 0xf7, 0x7e, 0x3d, 0x83,
	0xff, 0xab, 0x90, 0x5a, 0xae, 0x5d, 0x82, 0xed, 0x3b, 0xca, 0x74, 0xb7, 0x27, 0xe4, 0x20, 0xd5,
	0xe0, 0x2c, 0x39, 0x67, 0xa1, 0x2b, 0x69, 0x61, 0xe0, 0xb4, 0x60, 0x65, 0x02, 0x98, 0x5a, 0xfe,
	0xdf, 0xc1, 0xf5, 0x96, 0xd5, 0x29, 0x8b, 0xa0, 0x27, 0xd2, 0x81, 0xed, 0x28, 0xa1, 0x45, 0x02,
	0x16, 0xb4, 0xe1, 0x79, 0x67, 0x3a, 0x5c, 0x91, 0xfb, 0x0d, 0xa8, 0xbe, 0x66, 0x29, 0x6b, 0x99,
	0xf8, 0x01, 0xec, 0xdf, 0x2d, 0xc9, 0x5a, 0xb4, 0xa8, 0x34, 0x2a, 0x34, 0xa0, 0x79, 0xa1, 0x42,
	0x6a, 0xa5, 0xe6, 0xc5, 0xf7, 0x67, 0xb9, 0x1e, 0x4b, 0xfb, 0x9c, 0x76, 0xfd, 0x10, 0x93, 0x20,
	0x44, 0x93, 0xa0, 0x59, 0x2d, 0x75, 0x13, 0xf5, 0x83, 0x45, 0x3d, 0xe3, 0x37, 0xc2, 0xb0, 0x11,
	0x45, 0x1a, 0x8c, 0x69, 0x6f, 0x22, 0x9a, 0x57, 0x6f, 0x33, 0x8f, 0x4c, 0x67, 0x1e, 0xf9, 0x9a,
	0x79, 0xe4, 0x65, 0xee, 0x65, 0xa6, 0x73, 0x2f, 0xf3, 0x3e, 0xf7, 0x32, 0x4f, 0x67, 0x5b, 0x91,
	0x77, 0x52, 0x8b, 0x6b, 0xd4, 0x10, 0x18, 0xe8, 0x0b, 0x19, 0x8c, 0x83, 0x18, 0x47, 0xcb, 0xd4,
	0x6e, 0xde, 0x7d, 0xad, 0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0xb5, 0xce, 0xa1, 0x91,
	0x02, 0x00, 0x00,
}

func (m *ExecutionFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecutionFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExecutionFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DefaultParameters != 0 {
		i = encodeVarintExecutionFee(dAtA, i, uint64(m.DefaultParameters))
		i--
		dAtA[i] = 0x30
	}
	if m.Timeout != 0 {
		i = encodeVarintExecutionFee(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x28
	}
	if m.FailureFee != 0 {
		i = encodeVarintExecutionFee(dAtA, i, uint64(m.FailureFee))
		i--
		dAtA[i] = 0x20
	}
	if m.ExecutionFee != 0 {
		i = encodeVarintExecutionFee(dAtA, i, uint64(m.ExecutionFee))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TransactionType) > 0 {
		i -= len(m.TransactionType)
		copy(dAtA[i:], m.TransactionType)
		i = encodeVarintExecutionFee(dAtA, i, uint64(len(m.TransactionType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintExecutionFee(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetExecutionFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetExecutionFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetExecutionFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintExecutionFee(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x3a
	}
	if m.DefaultParameters != 0 {
		i = encodeVarintExecutionFee(dAtA, i, uint64(m.DefaultParameters))
		i--
		dAtA[i] = 0x30
	}
	if m.Timeout != 0 {
		i = encodeVarintExecutionFee(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x28
	}
	if m.FailureFee != 0 {
		i = encodeVarintExecutionFee(dAtA, i, uint64(m.FailureFee))
		i--
		dAtA[i] = 0x20
	}
	if m.ExecutionFee != 0 {
		i = encodeVarintExecutionFee(dAtA, i, uint64(m.ExecutionFee))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TransactionType) > 0 {
		i -= len(m.TransactionType)
		copy(dAtA[i:], m.TransactionType)
		i = encodeVarintExecutionFee(dAtA, i, uint64(len(m.TransactionType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintExecutionFee(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExecutionFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovExecutionFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExecutionFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovExecutionFee(uint64(l))
	}
	l = len(m.TransactionType)
	if l > 0 {
		n += 1 + l + sovExecutionFee(uint64(l))
	}
	if m.ExecutionFee != 0 {
		n += 1 + sovExecutionFee(uint64(m.ExecutionFee))
	}
	if m.FailureFee != 0 {
		n += 1 + sovExecutionFee(uint64(m.FailureFee))
	}
	if m.Timeout != 0 {
		n += 1 + sovExecutionFee(uint64(m.Timeout))
	}
	if m.DefaultParameters != 0 {
		n += 1 + sovExecutionFee(uint64(m.DefaultParameters))
	}
	return n
}

func (m *MsgSetExecutionFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovExecutionFee(uint64(l))
	}
	l = len(m.TransactionType)
	if l > 0 {
		n += 1 + l + sovExecutionFee(uint64(l))
	}
	if m.ExecutionFee != 0 {
		n += 1 + sovExecutionFee(uint64(m.ExecutionFee))
	}
	if m.FailureFee != 0 {
		n += 1 + sovExecutionFee(uint64(m.FailureFee))
	}
	if m.Timeout != 0 {
		n += 1 + sovExecutionFee(uint64(m.Timeout))
	}
	if m.DefaultParameters != 0 {
		n += 1 + sovExecutionFee(uint64(m.DefaultParameters))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovExecutionFee(uint64(l))
	}
	return n
}

func sovExecutionFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExecutionFee(x uint64) (n int) {
	return sovExecutionFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExecutionFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecutionFee
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
			return fmt.Errorf("proto: ExecutionFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecutionFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
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
				return ErrInvalidLengthExecutionFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecutionFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
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
				return ErrInvalidLengthExecutionFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecutionFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransactionType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionFee", wireType)
			}
			m.ExecutionFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureFee", wireType)
			}
			m.FailureFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FailureFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultParameters", wireType)
			}
			m.DefaultParameters = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultParameters |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExecutionFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExecutionFee
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExecutionFee
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
func (m *MsgSetExecutionFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecutionFee
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
			return fmt.Errorf("proto: MsgSetExecutionFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetExecutionFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
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
				return ErrInvalidLengthExecutionFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecutionFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
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
				return ErrInvalidLengthExecutionFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecutionFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransactionType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionFee", wireType)
			}
			m.ExecutionFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureFee", wireType)
			}
			m.FailureFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FailureFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultParameters", wireType)
			}
			m.DefaultParameters = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultParameters |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutionFee
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
				return ErrInvalidLengthExecutionFee
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthExecutionFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecutionFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExecutionFee
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExecutionFee
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
func skipExecutionFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExecutionFee
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
					return 0, ErrIntOverflowExecutionFee
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
					return 0, ErrIntOverflowExecutionFee
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
				return 0, ErrInvalidLengthExecutionFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExecutionFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExecutionFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExecutionFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExecutionFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExecutionFee = fmt.Errorf("proto: unexpected end of group")
)
