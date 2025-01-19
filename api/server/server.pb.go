// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: server/server.proto

package server

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

type CreateServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateServerRequest) Reset() {
	*x = CreateServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServerRequest) ProtoMessage() {}

func (x *CreateServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServerRequest.ProtoReflect.Descriptor instead.
func (*CreateServerRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{0}
}

type CreateServerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateServerReply) Reset() {
	*x = CreateServerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServerReply) ProtoMessage() {}

func (x *CreateServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServerReply.ProtoReflect.Descriptor instead.
func (*CreateServerReply) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{1}
}

type UpdateServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateServerRequest) Reset() {
	*x = UpdateServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServerRequest) ProtoMessage() {}

func (x *UpdateServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServerRequest.ProtoReflect.Descriptor instead.
func (*UpdateServerRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{2}
}

type UpdateServerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateServerReply) Reset() {
	*x = UpdateServerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServerReply) ProtoMessage() {}

func (x *UpdateServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServerReply.ProtoReflect.Descriptor instead.
func (*UpdateServerReply) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{3}
}

type DeleteServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteServerRequest) Reset() {
	*x = DeleteServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServerRequest) ProtoMessage() {}

func (x *DeleteServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServerRequest.ProtoReflect.Descriptor instead.
func (*DeleteServerRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{4}
}

type DeleteServerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteServerReply) Reset() {
	*x = DeleteServerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServerReply) ProtoMessage() {}

func (x *DeleteServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServerReply.ProtoReflect.Descriptor instead.
func (*DeleteServerReply) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{5}
}

type GetServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetServerRequest) Reset() {
	*x = GetServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerRequest) ProtoMessage() {}

func (x *GetServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerRequest.ProtoReflect.Descriptor instead.
func (*GetServerRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{6}
}

type GetServerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetServerReply) Reset() {
	*x = GetServerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerReply) ProtoMessage() {}

func (x *GetServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerReply.ProtoReflect.Descriptor instead.
func (*GetServerReply) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{7}
}

type ListServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListServerRequest) Reset() {
	*x = ListServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServerRequest) ProtoMessage() {}

func (x *ListServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServerRequest.ProtoReflect.Descriptor instead.
func (*ListServerRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{8}
}

type ListServerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListServerReply) Reset() {
	*x = ListServerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServerReply) ProtoMessage() {}

func (x *ListServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServerReply.ProtoReflect.Descriptor instead.
func (*ListServerReply) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{9}
}

var File_server_server_proto protoreflect.FileDescriptor

var file_server_server_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0x89, 0x03, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x4e,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x45,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x48, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x29, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x01, 0x5a,
	0x19, 0x6c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_server_server_proto_rawDescOnce sync.Once
	file_server_server_proto_rawDescData = file_server_server_proto_rawDesc
)

func file_server_server_proto_rawDescGZIP() []byte {
	file_server_server_proto_rawDescOnce.Do(func() {
		file_server_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_server_proto_rawDescData)
	})
	return file_server_server_proto_rawDescData
}

var file_server_server_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_server_server_proto_goTypes = []interface{}{
	(*CreateServerRequest)(nil), // 0: api.server.CreateServerRequest
	(*CreateServerReply)(nil),   // 1: api.server.CreateServerReply
	(*UpdateServerRequest)(nil), // 2: api.server.UpdateServerRequest
	(*UpdateServerReply)(nil),   // 3: api.server.UpdateServerReply
	(*DeleteServerRequest)(nil), // 4: api.server.DeleteServerRequest
	(*DeleteServerReply)(nil),   // 5: api.server.DeleteServerReply
	(*GetServerRequest)(nil),    // 6: api.server.GetServerRequest
	(*GetServerReply)(nil),      // 7: api.server.GetServerReply
	(*ListServerRequest)(nil),   // 8: api.server.ListServerRequest
	(*ListServerReply)(nil),     // 9: api.server.ListServerReply
}
var file_server_server_proto_depIdxs = []int32{
	0, // 0: api.server.Server.CreateServer:input_type -> api.server.CreateServerRequest
	2, // 1: api.server.Server.UpdateServer:input_type -> api.server.UpdateServerRequest
	4, // 2: api.server.Server.DeleteServer:input_type -> api.server.DeleteServerRequest
	6, // 3: api.server.Server.GetServer:input_type -> api.server.GetServerRequest
	8, // 4: api.server.Server.ListServer:input_type -> api.server.ListServerRequest
	1, // 5: api.server.Server.CreateServer:output_type -> api.server.CreateServerReply
	3, // 6: api.server.Server.UpdateServer:output_type -> api.server.UpdateServerReply
	5, // 7: api.server.Server.DeleteServer:output_type -> api.server.DeleteServerReply
	7, // 8: api.server.Server.GetServer:output_type -> api.server.GetServerReply
	9, // 9: api.server.Server.ListServer:output_type -> api.server.ListServerReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_server_proto_init() }
func file_server_server_proto_init() {
	if File_server_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServerRequest); i {
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
		file_server_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServerReply); i {
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
		file_server_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServerRequest); i {
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
		file_server_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServerReply); i {
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
		file_server_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServerRequest); i {
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
		file_server_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServerReply); i {
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
		file_server_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerRequest); i {
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
		file_server_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerReply); i {
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
		file_server_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServerRequest); i {
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
		file_server_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServerReply); i {
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
			RawDescriptor: file_server_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_server_proto_goTypes,
		DependencyIndexes: file_server_server_proto_depIdxs,
		MessageInfos:      file_server_server_proto_msgTypes,
	}.Build()
	File_server_server_proto = out.File
	file_server_server_proto_rawDesc = nil
	file_server_server_proto_goTypes = nil
	file_server_server_proto_depIdxs = nil
}
