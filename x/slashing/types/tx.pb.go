// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgActivate defines the Msg/Activate request type
type MsgActivate struct {
	ValidatorAddr string `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"address" yaml:"address"`
}

func (m *MsgActivate) Reset()         { *m = MsgActivate{} }
func (m *MsgActivate) String() string { return proto.CompactTextString(m) }
func (*MsgActivate) ProtoMessage()    {}
func (*MsgActivate) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{0}
}
func (m *MsgActivate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgActivate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgActivate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgActivate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgActivate.Merge(m, src)
}
func (m *MsgActivate) XXX_Size() int {
	return m.Size()
}
func (m *MsgActivate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgActivate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgActivate proto.InternalMessageInfo

// MsgActivateResponse defines the Msg/Activate response type
type MsgActivateResponse struct {
}

func (m *MsgActivateResponse) Reset()         { *m = MsgActivateResponse{} }
func (m *MsgActivateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgActivateResponse) ProtoMessage()    {}
func (*MsgActivateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{1}
}
func (m *MsgActivateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgActivateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgActivateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgActivateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgActivateResponse.Merge(m, src)
}
func (m *MsgActivateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgActivateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgActivateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgActivateResponse proto.InternalMessageInfo

// MsgPause defines the Msg/Pause request type
type MsgPause struct {
	ValidatorAddr string `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"address" yaml:"address"`
}

func (m *MsgPause) Reset()         { *m = MsgPause{} }
func (m *MsgPause) String() string { return proto.CompactTextString(m) }
func (*MsgPause) ProtoMessage()    {}
func (*MsgPause) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{2}
}
func (m *MsgPause) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPause) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPause.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPause) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPause.Merge(m, src)
}
func (m *MsgPause) XXX_Size() int {
	return m.Size()
}
func (m *MsgPause) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPause.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPause proto.InternalMessageInfo

// MsgUnpauseResponse defines the Msg/Unpause response type
type MsgPauseResponse struct {
}

func (m *MsgPauseResponse) Reset()         { *m = MsgPauseResponse{} }
func (m *MsgPauseResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPauseResponse) ProtoMessage()    {}
func (*MsgPauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{3}
}
func (m *MsgPauseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPauseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPauseResponse.Merge(m, src)
}
func (m *MsgPauseResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPauseResponse proto.InternalMessageInfo

// MsgUnpause defines the Msg/Unpause request type
type MsgUnpause struct {
	ValidatorAddr string `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"address" yaml:"address"`
}

func (m *MsgUnpause) Reset()         { *m = MsgUnpause{} }
func (m *MsgUnpause) String() string { return proto.CompactTextString(m) }
func (*MsgUnpause) ProtoMessage()    {}
func (*MsgUnpause) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{4}
}
func (m *MsgUnpause) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnpause) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnpause.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnpause) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnpause.Merge(m, src)
}
func (m *MsgUnpause) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnpause) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnpause.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnpause proto.InternalMessageInfo

// MsgUnpauseResponse defines the Msg/Unpause response type
type MsgUnpauseResponse struct {
}

func (m *MsgUnpauseResponse) Reset()         { *m = MsgUnpauseResponse{} }
func (m *MsgUnpauseResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUnpauseResponse) ProtoMessage()    {}
func (*MsgUnpauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{5}
}
func (m *MsgUnpauseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnpauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnpauseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnpauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnpauseResponse.Merge(m, src)
}
func (m *MsgUnpauseResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnpauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnpauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnpauseResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgActivate)(nil), "kira.slashing.MsgActivate")
	proto.RegisterType((*MsgActivateResponse)(nil), "kira.slashing.MsgActivateResponse")
	proto.RegisterType((*MsgPause)(nil), "kira.slashing.MsgPause")
	proto.RegisterType((*MsgPauseResponse)(nil), "kira.slashing.MsgPauseResponse")
	proto.RegisterType((*MsgUnpause)(nil), "kira.slashing.MsgUnpause")
	proto.RegisterType((*MsgUnpauseResponse)(nil), "kira.slashing.MsgUnpauseResponse")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor_0fd2153dc07d3b5c) }

var fileDescriptor_0fd2153dc07d3b5c = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0xa9, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcd, 0xce, 0x2c, 0x4a, 0xd4, 0x2b, 0xce, 0x49, 0x2c, 0xce,
	0xc8, 0xcc, 0x4b, 0x97, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xe8, 0x83, 0x58, 0x10, 0x45,
	0x4a, 0xb1, 0x5c, 0xdc, 0xbe, 0xc5, 0xe9, 0x8e, 0xc9, 0x25, 0x99, 0x65, 0x89, 0x25, 0xa9, 0x42,
	0x2e, 0x5c, 0x7c, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xf9, 0x45, 0xf1, 0x89, 0x29, 0x29,
	0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0xb2, 0xaf, 0xee, 0xc9, 0xb3, 0x83, 0xf8, 0xa9,
	0xc5, 0xc5, 0x9f, 0xee, 0xc9, 0xf3, 0x55, 0x26, 0xe6, 0xe6, 0x58, 0x29, 0x41, 0x05, 0x94, 0x82,
	0x78, 0xe1, 0x9a, 0x1c, 0x53, 0x52, 0x8a, 0xac, 0x38, 0x3a, 0x16, 0xc8, 0x33, 0xcc, 0x58, 0x20,
	0xcf, 0xa8, 0x24, 0xca, 0x25, 0x8c, 0x64, 0x7c, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa,
	0x52, 0x14, 0x17, 0x87, 0x6f, 0x71, 0x7a, 0x40, 0x62, 0x69, 0x31, 0xf5, 0xad, 0x14, 0xe2, 0x12,
	0x80, 0x99, 0x0d, 0xb7, 0x2f, 0x86, 0x8b, 0xcb, 0xb7, 0x38, 0x3d, 0x34, 0xaf, 0x80, 0x26, 0x36,
	0x8a, 0x70, 0x09, 0x21, 0x4c, 0x87, 0xd9, 0x69, 0x74, 0x97, 0x91, 0x8b, 0xd9, 0xb7, 0x38, 0x5d,
	0xc8, 0x8b, 0x8b, 0x03, 0x1e, 0xbc, 0x52, 0x7a, 0x28, 0x71, 0xa2, 0x87, 0x14, 0x36, 0x52, 0x4a,
	0xb8, 0xe5, 0x60, 0x66, 0x0a, 0x39, 0x72, 0xb1, 0x42, 0x02, 0x4d, 0x1c, 0x53, 0x31, 0x58, 0x42,
	0x4a, 0x1e, 0x87, 0x04, 0xdc, 0x08, 0x77, 0x2e, 0x76, 0x58, 0x38, 0x48, 0x62, 0xaa, 0x85, 0x4a,
	0x49, 0x29, 0xe2, 0x94, 0x82, 0x19, 0xe4, 0xe4, 0xb1, 0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x69, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0x7b, 0x67, 0x16, 0x25, 0x3a, 0xe7, 0x17, 0xa5, 0xea, 0x17, 0xa7, 0x66, 0x27,
	0x66, 0xea, 0x57, 0xe8, 0xc3, 0x8c, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x27,
	0x45, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0x0c, 0x08, 0x30, 0xbb, 0x02, 0x00, 0x00,
}

func (this *MsgActivate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgActivate)
	if !ok {
		that2, ok := that.(MsgActivate)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ValidatorAddr != that1.ValidatorAddr {
		return false
	}
	return true
}
func (this *MsgActivateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgActivateResponse)
	if !ok {
		that2, ok := that.(MsgActivateResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *MsgPause) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgPause)
	if !ok {
		that2, ok := that.(MsgPause)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ValidatorAddr != that1.ValidatorAddr {
		return false
	}
	return true
}
func (this *MsgPauseResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgPauseResponse)
	if !ok {
		that2, ok := that.(MsgPauseResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *MsgUnpause) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUnpause)
	if !ok {
		that2, ok := that.(MsgUnpause)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ValidatorAddr != that1.ValidatorAddr {
		return false
	}
	return true
}
func (this *MsgUnpauseResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUnpauseResponse)
	if !ok {
		that2, ok := that.(MsgUnpauseResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Activate defines a method for activating an inactivated validator, thus returning
	// them into the joined validator set, so they can begin receiving provisions
	// and rewards again.
	Activate(ctx context.Context, in *MsgActivate, opts ...grpc.CallOption) (*MsgActivateResponse, error)
	// Pause defines a method for pausing an active validator
	Pause(ctx context.Context, in *MsgPause, opts ...grpc.CallOption) (*MsgPauseResponse, error)
	// Unpause defines a method for a paused validator
	Unpause(ctx context.Context, in *MsgUnpause, opts ...grpc.CallOption) (*MsgUnpauseResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Activate(ctx context.Context, in *MsgActivate, opts ...grpc.CallOption) (*MsgActivateResponse, error) {
	out := new(MsgActivateResponse)
	err := c.cc.Invoke(ctx, "/kira.slashing.Msg/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Pause(ctx context.Context, in *MsgPause, opts ...grpc.CallOption) (*MsgPauseResponse, error) {
	out := new(MsgPauseResponse)
	err := c.cc.Invoke(ctx, "/kira.slashing.Msg/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Unpause(ctx context.Context, in *MsgUnpause, opts ...grpc.CallOption) (*MsgUnpauseResponse, error) {
	out := new(MsgUnpauseResponse)
	err := c.cc.Invoke(ctx, "/kira.slashing.Msg/Unpause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Activate defines a method for activating an inactivated validator, thus returning
	// them into the joined validator set, so they can begin receiving provisions
	// and rewards again.
	Activate(context.Context, *MsgActivate) (*MsgActivateResponse, error)
	// Pause defines a method for pausing an active validator
	Pause(context.Context, *MsgPause) (*MsgPauseResponse, error)
	// Unpause defines a method for a paused validator
	Unpause(context.Context, *MsgUnpause) (*MsgUnpauseResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Activate(ctx context.Context, req *MsgActivate) (*MsgActivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (*UnimplementedMsgServer) Pause(ctx context.Context, req *MsgPause) (*MsgPauseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (*UnimplementedMsgServer) Unpause(ctx context.Context, req *MsgUnpause) (*MsgUnpauseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgActivate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.slashing.Msg/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Activate(ctx, req.(*MsgActivate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPause)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.slashing.Msg/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Pause(ctx, req.(*MsgPause))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnpause)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.slashing.Msg/Unpause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Unpause(ctx, req.(*MsgUnpause))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kira.slashing.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activate",
			Handler:    _Msg_Activate_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _Msg_Pause_Handler,
		},
		{
			MethodName: "Unpause",
			Handler:    _Msg_Unpause_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx.proto",
}

func (m *MsgActivate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgActivate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgActivate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgActivateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgActivateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgActivateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgPause) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPause) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPause) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPauseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPauseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPauseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUnpause) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnpause) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnpause) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnpauseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnpauseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnpauseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgActivate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgActivateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgPause) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgPauseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUnpause) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUnpauseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgActivate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgActivate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgActivate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgActivateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgActivateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgActivateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgPause) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPause: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPause: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgPauseResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPauseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPauseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUnpause) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUnpause: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnpause: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUnpauseResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUnpauseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnpauseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
