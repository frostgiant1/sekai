// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: kira/staking/query.proto

package proto

import (
	context "context"
	_ "github.com/gogo/protobuf/gogoproto"
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

// ValidatorsRequest is the request type for the Query/AllValidators RPC method.
type ValidatorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Valkey   string `protobuf:"bytes,2,opt,name=valkey,proto3" json:"valkey,omitempty"`
	Pubkey   string `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Moniker  string `protobuf:"bytes,4,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Status   string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Proposer string `protobuf:"bytes,6,opt,name=proposer,proto3" json:"proposer,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *PageRequest `protobuf:"bytes,7,opt,name=pagination,proto3" json:"pagination,omitempty"`
	All        bool         `protobuf:"varint,8,opt,name=all,proto3" json:"all,omitempty"`
}

func (x *ValidatorsRequest) Reset() {
	*x = ValidatorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kira_staking_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorsRequest) ProtoMessage() {}

func (x *ValidatorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kira_staking_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorsRequest.ProtoReflect.Descriptor instead.
func (*ValidatorsRequest) Descriptor() ([]byte, []int) {
	return file_kira_staking_query_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatorsRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ValidatorsRequest) GetValkey() string {
	if x != nil {
		return x.Valkey
	}
	return ""
}

func (x *ValidatorsRequest) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

func (x *ValidatorsRequest) GetMoniker() string {
	if x != nil {
		return x.Moniker
	}
	return ""
}

func (x *ValidatorsRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ValidatorsRequest) GetProposer() string {
	if x != nil {
		return x.Proposer
	}
	return ""
}

func (x *ValidatorsRequest) GetPagination() *PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ValidatorsRequest) GetAll() bool {
	if x != nil {
		return x.All
	}
	return false
}

type QueryValidator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Valkey     string `protobuf:"bytes,2,opt,name=valkey,proto3" json:"valkey,omitempty"`
	Pubkey     string `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Proposer   string `protobuf:"bytes,4,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Moniker    string `protobuf:"bytes,5,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Website    string `protobuf:"bytes,6,opt,name=website,proto3" json:"website,omitempty"`
	Social     string `protobuf:"bytes,7,opt,name=social,proto3" json:"social,omitempty"`
	Identity   string `protobuf:"bytes,8,opt,name=identity,proto3" json:"identity,omitempty"`
	Commission string `protobuf:"bytes,9,opt,name=commission,proto3" json:"commission,omitempty"`
	Status     string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Rank       int64  `protobuf:"varint,11,opt,name=rank,proto3" json:"rank,omitempty"`
	Streak     int64  `protobuf:"varint,12,opt,name=streak,proto3" json:"streak,omitempty"`
	Mischance  int64  `protobuf:"varint,13,opt,name=mischance,proto3" json:"mischance,omitempty"`
}

func (x *QueryValidator) Reset() {
	*x = QueryValidator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kira_staking_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidator) ProtoMessage() {}

func (x *QueryValidator) ProtoReflect() protoreflect.Message {
	mi := &file_kira_staking_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryValidator.ProtoReflect.Descriptor instead.
func (*QueryValidator) Descriptor() ([]byte, []int) {
	return file_kira_staking_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryValidator) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *QueryValidator) GetValkey() string {
	if x != nil {
		return x.Valkey
	}
	return ""
}

func (x *QueryValidator) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

func (x *QueryValidator) GetProposer() string {
	if x != nil {
		return x.Proposer
	}
	return ""
}

func (x *QueryValidator) GetMoniker() string {
	if x != nil {
		return x.Moniker
	}
	return ""
}

func (x *QueryValidator) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *QueryValidator) GetSocial() string {
	if x != nil {
		return x.Social
	}
	return ""
}

func (x *QueryValidator) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *QueryValidator) GetCommission() string {
	if x != nil {
		return x.Commission
	}
	return ""
}

func (x *QueryValidator) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *QueryValidator) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *QueryValidator) GetStreak() int64 {
	if x != nil {
		return x.Streak
	}
	return 0
}

func (x *QueryValidator) GetMischance() int64 {
	if x != nil {
		return x.Mischance
	}
	return 0
}

// ValidatorsResponse is response type for the Query/Validators RPC method
type ValidatorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validators contains all the queried validators.
	Validators []*QueryValidator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	Actors     []string          `protobuf:"bytes,2,rep,name=actors,proto3" json:"actors,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *PageResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ValidatorsResponse) Reset() {
	*x = ValidatorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kira_staking_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorsResponse) ProtoMessage() {}

func (x *ValidatorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kira_staking_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorsResponse.ProtoReflect.Descriptor instead.
func (*ValidatorsResponse) Descriptor() ([]byte, []int) {
	return file_kira_staking_query_proto_rawDescGZIP(), []int{2}
}

func (x *ValidatorsResponse) GetValidators() []*QueryValidator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *ValidatorsResponse) GetActors() []string {
	if x != nil {
		return x.Actors
	}
	return nil
}

func (x *ValidatorsResponse) GetPagination() *PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_kira_staking_query_proto protoreflect.FileDescriptor

var file_kira_staking_query_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6b, 0x69, 0x72, 0x61, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6b, 0x69, 0x72, 0x61,
	0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x1d, 0x6b, 0x69, 0x72, 0x61, 0x2f, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x02, 0x0a, 0x11,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f,
	0x6e, 0x69, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12, 0x73, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6b, 0x69, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x38, 0xfa, 0xde, 0x1f, 0x34, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x6c, 0x6c,
	0x22, 0xe6, 0x02, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e,
	0x69, 0x6b, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x69,
	0x6b, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x73, 0x63, 0x68, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x69, 0x73, 0x63, 0x68, 0x61,
	0x6e, 0x63, 0x65, 0x3a, 0x04, 0xe8, 0xa0, 0x1f, 0x01, 0x22, 0xe1, 0x01, 0x0a, 0x12, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x69, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x75, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x69, 0x72,
	0x61, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0xfa, 0xde, 0x1f, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xa3, 0x01,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x99, 0x01, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x6b, 0x69, 0x72, 0x61, 0x2e, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6b, 0x69, 0x72, 0x61, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x61, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73,
	0x92, 0x41, 0x30, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x20, 0x41, 0x6c, 0x6c, 0x20, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x1a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x20, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x2e, 0x42, 0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4b, 0x69, 0x72, 0x61, 0x43, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x6b, 0x61, 0x69,
	0x2f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x58, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x92, 0x41, 0x44,
	0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01, 0x72, 0x38, 0x0a, 0x0c, 0x67, 0x52,
	0x50, 0x43, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x28, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b,
	0x69, 0x72, 0x61, 0x43, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x6b, 0x61, 0x69, 0x2f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x58, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kira_staking_query_proto_rawDescOnce sync.Once
	file_kira_staking_query_proto_rawDescData = file_kira_staking_query_proto_rawDesc
)

func file_kira_staking_query_proto_rawDescGZIP() []byte {
	file_kira_staking_query_proto_rawDescOnce.Do(func() {
		file_kira_staking_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_kira_staking_query_proto_rawDescData)
	})
	return file_kira_staking_query_proto_rawDescData
}

var file_kira_staking_query_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kira_staking_query_proto_goTypes = []interface{}{
	(*ValidatorsRequest)(nil),  // 0: kira.staking.ValidatorsRequest
	(*QueryValidator)(nil),     // 1: kira.staking.QueryValidator
	(*ValidatorsResponse)(nil), // 2: kira.staking.ValidatorsResponse
	(*PageRequest)(nil),        // 3: kira.staking.PageRequest
	(*PageResponse)(nil),       // 4: kira.staking.PageResponse
}
var file_kira_staking_query_proto_depIdxs = []int32{
	3, // 0: kira.staking.ValidatorsRequest.pagination:type_name -> kira.staking.PageRequest
	1, // 1: kira.staking.ValidatorsResponse.validators:type_name -> kira.staking.QueryValidator
	4, // 2: kira.staking.ValidatorsResponse.pagination:type_name -> kira.staking.PageResponse
	0, // 3: kira.staking.Query.Validators:input_type -> kira.staking.ValidatorsRequest
	2, // 4: kira.staking.Query.Validators:output_type -> kira.staking.ValidatorsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kira_staking_query_proto_init() }
func file_kira_staking_query_proto_init() {
	if File_kira_staking_query_proto != nil {
		return
	}
	file_kira_staking_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kira_staking_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorsRequest); i {
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
		file_kira_staking_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidator); i {
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
		file_kira_staking_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorsResponse); i {
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
			RawDescriptor: file_kira_staking_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kira_staking_query_proto_goTypes,
		DependencyIndexes: file_kira_staking_query_proto_depIdxs,
		MessageInfos:      file_kira_staking_query_proto_msgTypes,
	}.Build()
	File_kira_staking_query_proto = out.File
	file_kira_staking_query_proto_rawDesc = nil
	file_kira_staking_query_proto_goTypes = nil
	file_kira_staking_query_proto_depIdxs = nil
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
	// Validators queries all validators that match the given status.
	Validators(ctx context.Context, in *ValidatorsRequest, opts ...grpc.CallOption) (*ValidatorsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Validators(ctx context.Context, in *ValidatorsRequest, opts ...grpc.CallOption) (*ValidatorsResponse, error) {
	out := new(ValidatorsResponse)
	err := c.cc.Invoke(ctx, "/kira.staking.Query/Validators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Validators queries all validators that match the given status.
	Validators(context.Context, *ValidatorsRequest) (*ValidatorsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Validators(context.Context, *ValidatorsRequest) (*ValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validators not implemented")
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Validators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.staking.Query/Validators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validators(ctx, req.(*ValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kira.staking.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validators",
			Handler:    _Query_Validators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kira/staking/query.proto",
}
