// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: rpc/core.proto

package core

import (
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

type AvatarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *AvatarInfo) Reset() {
	*x = AvatarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarInfo) ProtoMessage() {}

func (x *AvatarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarInfo.ProtoReflect.Descriptor instead.
func (*AvatarInfo) Descriptor() ([]byte, []int) {
	return file_rpc_core_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AvatarInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// base message
type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_rpc_core_proto_rawDescGZIP(), []int{1}
}

func (x *BaseResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type BaseUUIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *BaseUUIDResp) Reset() {
	*x = BaseUUIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseUUIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseUUIDResp) ProtoMessage() {}

func (x *BaseUUIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseUUIDResp.ProtoReflect.Descriptor instead.
func (*BaseUUIDResp) Descriptor() ([]byte, []int) {
	return file_rpc_core_proto_rawDescGZIP(), []int{2}
}

func (x *BaseUUIDResp) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *BaseUUIDResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type IdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdReq) Reset() {
	*x = IdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReq.ProtoReflect.Descriptor instead.
func (*IdReq) Descriptor() ([]byte, []int) {
	return file_rpc_core_proto_rawDescGZIP(), []int{3}
}

func (x *IdReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    *string `protobuf:"bytes,1,opt,name=username,proto3,oneof" json:"username,omitempty"`
	Password    *string `protobuf:"bytes,2,opt,name=password,proto3,oneof" json:"password,omitempty"`
	LoginStatus *uint64 `protobuf:"varint,3,opt,name=login_status,json=loginStatus,proto3,oneof" json:"login_status,omitempty"`
	Avatar      *string `protobuf:"bytes,4,opt,name=avatar,proto3,oneof" json:"avatar,omitempty"`
	Email       *string `protobuf:"bytes,5,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Phone       *string `protobuf:"bytes,6,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_rpc_core_proto_rawDescGZIP(), []int{4}
}

func (x *UserInfo) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *UserInfo) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *UserInfo) GetLoginStatus() uint64 {
	if x != nil && x.LoginStatus != nil {
		return *x.LoginStatus
	}
	return 0
}

func (x *UserInfo) GetAvatar() string {
	if x != nil && x.Avatar != nil {
		return *x.Avatar
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UserInfo) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

type UserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *UserListReq) Reset() {
	*x = UserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListReq) ProtoMessage() {}

func (x *UserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListReq.ProtoReflect.Descriptor instead.
func (*UserListReq) Descriptor() ([]byte, []int) {
	return file_rpc_core_proto_rawDescGZIP(), []int{5}
}

func (x *UserListReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *UserListReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type UserListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint64      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*UserInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *UserListResp) Reset() {
	*x = UserListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResp) ProtoMessage() {}

func (x *UserListResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResp.ProtoReflect.Descriptor instead.
func (*UserListResp) Descriptor() ([]byte, []int) {
	return file_rpc_core_proto_rawDescGZIP(), []int{6}
}

func (x *UserListResp) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserListResp) GetData() []*UserInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type UsernameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UsernameReq) Reset() {
	*x = UsernameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsernameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernameReq) ProtoMessage() {}

func (x *UsernameReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernameReq.ProtoReflect.Descriptor instead.
func (*UsernameReq) Descriptor() ([]byte, []int) {
	return file_rpc_core_proto_rawDescGZIP(), []int{7}
}

func (x *UsernameReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_rpc_core_proto protoreflect.FileDescriptor

var file_rpc_core_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x39, 0x0a, 0x0a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0x1c, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x34, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x91,
	0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a,
	0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x22, 0x3e, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x48, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xdd, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x72, 0x65,
	0x12, 0x30, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x12,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2c, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x30, 0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x34, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_core_proto_rawDescOnce sync.Once
	file_rpc_core_proto_rawDescData = file_rpc_core_proto_rawDesc
)

func file_rpc_core_proto_rawDescGZIP() []byte {
	file_rpc_core_proto_rawDescOnce.Do(func() {
		file_rpc_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_core_proto_rawDescData)
	})
	return file_rpc_core_proto_rawDescData
}

var file_rpc_core_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_core_proto_goTypes = []interface{}{
	(*AvatarInfo)(nil),   // 0: core.AvatarInfo
	(*BaseResp)(nil),     // 1: core.BaseResp
	(*BaseUUIDResp)(nil), // 2: core.BaseUUIDResp
	(*IdReq)(nil),        // 3: core.IdReq
	(*UserInfo)(nil),     // 4: core.UserInfo
	(*UserListReq)(nil),  // 5: core.UserListReq
	(*UserListResp)(nil), // 6: core.UserListResp
	(*UsernameReq)(nil),  // 7: core.UsernameReq
}
var file_rpc_core_proto_depIdxs = []int32{
	4, // 0: core.UserListResp.data:type_name -> core.UserInfo
	4, // 1: core.Core.createUser:input_type -> core.UserInfo
	4, // 2: core.Core.updateUser:input_type -> core.UserInfo
	0, // 3: core.Core.uploadAvatar:input_type -> core.AvatarInfo
	5, // 4: core.Core.getUserList:input_type -> core.UserListReq
	3, // 5: core.Core.getUserById:input_type -> core.IdReq
	7, // 6: core.Core.getUserByUsername:input_type -> core.UsernameReq
	3, // 7: core.Core.deleteUser:input_type -> core.IdReq
	2, // 8: core.Core.createUser:output_type -> core.BaseUUIDResp
	1, // 9: core.Core.updateUser:output_type -> core.BaseResp
	1, // 10: core.Core.uploadAvatar:output_type -> core.BaseResp
	6, // 11: core.Core.getUserList:output_type -> core.UserListResp
	4, // 12: core.Core.getUserById:output_type -> core.UserInfo
	4, // 13: core.Core.getUserByUsername:output_type -> core.UserInfo
	1, // 14: core.Core.deleteUser:output_type -> core.BaseResp
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_core_proto_init() }
func file_rpc_core_proto_init() {
	if File_rpc_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarInfo); i {
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
		file_rpc_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_rpc_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseUUIDResp); i {
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
		file_rpc_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdReq); i {
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
		file_rpc_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_rpc_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListReq); i {
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
		file_rpc_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListResp); i {
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
		file_rpc_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsernameReq); i {
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
	file_rpc_core_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_core_proto_goTypes,
		DependencyIndexes: file_rpc_core_proto_depIdxs,
		MessageInfos:      file_rpc_core_proto_msgTypes,
	}.Build()
	File_rpc_core_proto = out.File
	file_rpc_core_proto_rawDesc = nil
	file_rpc_core_proto_goTypes = nil
	file_rpc_core_proto_depIdxs = nil
}
