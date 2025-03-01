// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: cosmos/bank/query.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Coin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Coin) Reset() {
	*x = Coin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coin.ProtoReflect.Descriptor instead.
func (*Coin) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_query_proto_rawDescGZIP(), []int{0}
}

func (x *Coin) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Coin) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key is a value returned in PageResponse.next_key to begin
	// querying the next page most efficiently. Only one of offset or key
	// should be set.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// offset is a numeric offset that can be used when key is unavailable.
	// It is less efficient than using key. Only one of offset or key should
	// be set.
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// limit is the total number of results to be returned in the result page.
	// If left empty it will default to a value to be set by each app.
	Limit uint64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// count_total is set to true  to indicate that the result set should include
	// a count of the total number of items available for pagination in UIs. count_total
	// is only respected when offset is used. It is ignored when key is set.
	CountTotal bool `protobuf:"varint,4,opt,name=count_total,json=countTotal,proto3" json:"count_total,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_query_proto_rawDescGZIP(), []int{1}
}

func (x *PageRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PageRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *PageRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PageRequest) GetCountTotal() bool {
	if x != nil {
		return x.CountTotal
	}
	return false
}

type PageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	NextKey []byte `protobuf:"bytes,1,opt,name=next_key,json=nextKey,proto3" json:"next_key,omitempty"`
	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total uint64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PageResponse) Reset() {
	*x = PageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageResponse) ProtoMessage() {}

func (x *PageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageResponse.ProtoReflect.Descriptor instead.
func (*PageResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_query_proto_rawDescGZIP(), []int{2}
}

func (x *PageResponse) GetNextKey() []byte {
	if x != nil {
		return x.NextKey
	}
	return nil
}

func (x *PageResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// QueryBalanceRequest is the request type for the Query/AllBalances RPC method.
type QueryAllBalancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the address to query balances for.
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllBalancesRequest) Reset() {
	*x = QueryAllBalancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllBalancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllBalancesRequest) ProtoMessage() {}

func (x *QueryAllBalancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllBalancesRequest.ProtoReflect.Descriptor instead.
func (*QueryAllBalancesRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryAllBalancesRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *QueryAllBalancesRequest) GetPagination() *PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryAllBalancesResponse is the response type for the Query/AllBalances RPC method.
type QueryAllBalancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// balances is the balances of all the coins.
	Balances []*Coin `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllBalancesResponse) Reset() {
	*x = QueryAllBalancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllBalancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllBalancesResponse) ProtoMessage() {}

func (x *QueryAllBalancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllBalancesResponse.ProtoReflect.Descriptor instead.
func (*QueryAllBalancesResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_query_proto_rawDescGZIP(), []int{4}
}

func (x *QueryAllBalancesResponse) GetBalances() []*Coin {
	if x != nil {
		return x.Balances
	}
	return nil
}

func (x *QueryAllBalancesResponse) GetPagination() *PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryTotalSupplyRequest is the request type for the Query/TotalSupply RPC method.
type QueryTotalSupplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryTotalSupplyRequest) Reset() {
	*x = QueryTotalSupplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTotalSupplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTotalSupplyRequest) ProtoMessage() {}

func (x *QueryTotalSupplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTotalSupplyRequest.ProtoReflect.Descriptor instead.
func (*QueryTotalSupplyRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_query_proto_rawDescGZIP(), []int{5}
}

// QueryTotalSupplyResponse is the response type for the Query/TotalSupply RPC method
type QueryTotalSupplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// supply is the supply of the coins
	Supply []*Coin `protobuf:"bytes,1,rep,name=supply,proto3" json:"supply,omitempty"`
}

func (x *QueryTotalSupplyResponse) Reset() {
	*x = QueryTotalSupplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTotalSupplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTotalSupplyResponse) ProtoMessage() {}

func (x *QueryTotalSupplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTotalSupplyResponse.ProtoReflect.Descriptor instead.
func (*QueryTotalSupplyResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_query_proto_rawDescGZIP(), []int{6}
}

func (x *QueryTotalSupplyResponse) GetSupply() []*Coin {
	if x != nil {
		return x.Supply
	}
	return nil
}

var File_cosmos_bank_query_proto protoreflect.FileDescriptor

var file_cosmos_bank_query_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x04,
	0x43, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x6e, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x3f, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x75, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x94, 0x01, 0x0a, 0x18, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x43, 0x6f, 0x69, 0x6e, 0x52, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x41,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x19, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x18,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x69, 0x6e, 0x52, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x32, 0x81, 0x03, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0xc2, 0x01, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x6c, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x56, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x92,
	0x41, 0x28, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x20, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x20, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x12, 0xb2, 0x01, 0x0a, 0x0b, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12,
	0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x6e,
	0x6b, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x92, 0x41, 0x24, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x0c, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x20, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x1a, 0x0d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x20, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x42,
	0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x69,
	0x72, 0x61, 0x43, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x6b, 0x61, 0x69, 0x2f, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x58, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x92, 0x41, 0x44, 0x12, 0x05, 0x32, 0x03,
	0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01, 0x72, 0x38, 0x0a, 0x0c, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x69, 0x72, 0x61, 0x43,
	0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x6b, 0x61, 0x69, 0x2f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x58,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_bank_query_proto_rawDescOnce sync.Once
	file_cosmos_bank_query_proto_rawDescData = file_cosmos_bank_query_proto_rawDesc
)

func file_cosmos_bank_query_proto_rawDescGZIP() []byte {
	file_cosmos_bank_query_proto_rawDescOnce.Do(func() {
		file_cosmos_bank_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_bank_query_proto_rawDescData)
	})
	return file_cosmos_bank_query_proto_rawDescData
}

var file_cosmos_bank_query_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cosmos_bank_query_proto_goTypes = []interface{}{
	(*Coin)(nil),                     // 0: cosmos.bank.v1beta1.Coin
	(*PageRequest)(nil),              // 1: cosmos.bank.v1beta1.PageRequest
	(*PageResponse)(nil),             // 2: cosmos.bank.v1beta1.PageResponse
	(*QueryAllBalancesRequest)(nil),  // 3: cosmos.bank.v1beta1.QueryAllBalancesRequest
	(*QueryAllBalancesResponse)(nil), // 4: cosmos.bank.v1beta1.QueryAllBalancesResponse
	(*QueryTotalSupplyRequest)(nil),  // 5: cosmos.bank.v1beta1.QueryTotalSupplyRequest
	(*QueryTotalSupplyResponse)(nil), // 6: cosmos.bank.v1beta1.QueryTotalSupplyResponse
}
var file_cosmos_bank_query_proto_depIdxs = []int32{
	1, // 0: cosmos.bank.v1beta1.QueryAllBalancesRequest.pagination:type_name -> cosmos.bank.v1beta1.PageRequest
	0, // 1: cosmos.bank.v1beta1.QueryAllBalancesResponse.balances:type_name -> cosmos.bank.v1beta1.Coin
	2, // 2: cosmos.bank.v1beta1.QueryAllBalancesResponse.pagination:type_name -> cosmos.bank.v1beta1.PageResponse
	0, // 3: cosmos.bank.v1beta1.QueryTotalSupplyResponse.supply:type_name -> cosmos.bank.v1beta1.Coin
	3, // 4: cosmos.bank.v1beta1.Query.AllBalances:input_type -> cosmos.bank.v1beta1.QueryAllBalancesRequest
	5, // 5: cosmos.bank.v1beta1.Query.TotalSupply:input_type -> cosmos.bank.v1beta1.QueryTotalSupplyRequest
	4, // 6: cosmos.bank.v1beta1.Query.AllBalances:output_type -> cosmos.bank.v1beta1.QueryAllBalancesResponse
	6, // 7: cosmos.bank.v1beta1.Query.TotalSupply:output_type -> cosmos.bank.v1beta1.QueryTotalSupplyResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cosmos_bank_query_proto_init() }
func file_cosmos_bank_query_proto_init() {
	if File_cosmos_bank_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_bank_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_bank_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_bank_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_bank_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllBalancesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_bank_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllBalancesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_bank_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTotalSupplyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_bank_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTotalSupplyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_bank_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_bank_query_proto_goTypes,
		DependencyIndexes: file_cosmos_bank_query_proto_depIdxs,
		MessageInfos:      file_cosmos_bank_query_proto_msgTypes,
	}.Build()
	File_cosmos_bank_query_proto = out.File
	file_cosmos_bank_query_proto_rawDesc = nil
	file_cosmos_bank_query_proto_goTypes = nil
	file_cosmos_bank_query_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// AllBalances queries the balance of all coins for a single account.
	AllBalances(ctx context.Context, in *QueryAllBalancesRequest, opts ...grpc.CallOption) (*QueryAllBalancesResponse, error)
	TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) AllBalances(ctx context.Context, in *QueryAllBalancesRequest, opts ...grpc.CallOption) (*QueryAllBalancesResponse, error) {
	out := new(QueryAllBalancesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/AllBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error) {
	out := new(QueryTotalSupplyResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/TotalSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// AllBalances queries the balance of all coins for a single account.
	AllBalances(context.Context, *QueryAllBalancesRequest) (*QueryAllBalancesResponse, error)
	TotalSupply(context.Context, *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) AllBalances(context.Context, *QueryAllBalancesRequest) (*QueryAllBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllBalances not implemented")
}
func (*UnimplementedQueryServer) TotalSupply(context.Context, *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalSupply not implemented")
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_AllBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/AllBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllBalances(ctx, req.(*QueryAllBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/TotalSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalSupply(ctx, req.(*QueryTotalSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.bank.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllBalances",
			Handler:    _Query_AllBalances_Handler,
		},
		{
			MethodName: "TotalSupply",
			Handler:    _Query_TotalSupply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/bank/query.proto",
}
