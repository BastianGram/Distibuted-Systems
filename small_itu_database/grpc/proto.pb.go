// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: grpc/proto.proto

package proto

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

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	mi := &file_grpc_proto_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type JoinResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *JoinResponce) Reset() {
	*x = JoinResponce{}
	mi := &file_grpc_proto_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponce) ProtoMessage() {}

func (x *JoinResponce) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponce.ProtoReflect.Descriptor instead.
func (*JoinResponce) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{1}
}

func (x *JoinResponce) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ServerBroadcast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ClientName string `protobuf:"bytes,2,opt,name=clientName,proto3" json:"clientName,omitempty"`
}

func (x *ServerBroadcast) Reset() {
	*x = ServerBroadcast{}
	mi := &file_grpc_proto_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerBroadcast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerBroadcast) ProtoMessage() {}

func (x *ServerBroadcast) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerBroadcast.ProtoReflect.Descriptor instead.
func (*ServerBroadcast) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{2}
}

func (x *ServerBroadcast) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ServerBroadcast) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

type BroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ClientName string `protobuf:"bytes,2,opt,name=clientName,proto3" json:"clientName,omitempty"`
}

func (x *BroadcastRequest) Reset() {
	*x = BroadcastRequest{}
	mi := &file_grpc_proto_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRequest) ProtoMessage() {}

func (x *BroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRequest.ProtoReflect.Descriptor instead.
func (*BroadcastRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{3}
}

func (x *BroadcastRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BroadcastRequest) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

type ClientLeaves struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName string `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
}

func (x *ClientLeaves) Reset() {
	*x = ClientLeaves{}
	mi := &file_grpc_proto_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientLeaves) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientLeaves) ProtoMessage() {}

func (x *ClientLeaves) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientLeaves.ProtoReflect.Descriptor instead.
func (*ClientLeaves) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{4}
}

func (x *ClientLeaves) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

type ServerClientLeaves struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName string `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
}

func (x *ServerClientLeaves) Reset() {
	*x = ServerClientLeaves{}
	mi := &file_grpc_proto_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerClientLeaves) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerClientLeaves) ProtoMessage() {}

func (x *ServerClientLeaves) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerClientLeaves.ProtoReflect.Descriptor instead.
func (*ServerClientLeaves) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{5}
}

func (x *ServerClientLeaves) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

var File_grpc_proto_proto protoreflect.FileDescriptor

var file_grpc_proto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x0f, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x9f, 0x01, 0x0a, 0x0b, 0x49,
	0x54, 0x55, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x4a, 0x6f,
	0x69, 0x6e, 0x12, 0x0c, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x11,
	0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c,
	0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x73, 0x1a, 0x13, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_proto_proto_rawDescOnce sync.Once
	file_grpc_proto_proto_rawDescData = file_grpc_proto_proto_rawDesc
)

func file_grpc_proto_proto_rawDescGZIP() []byte {
	file_grpc_proto_proto_rawDescOnce.Do(func() {
		file_grpc_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_proto_rawDescData)
	})
	return file_grpc_proto_proto_rawDescData
}

var file_grpc_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_proto_proto_goTypes = []any{
	(*JoinRequest)(nil),        // 0: JoinRequest
	(*JoinResponce)(nil),       // 1: JoinResponce
	(*ServerBroadcast)(nil),    // 2: ServerBroadcast
	(*BroadcastRequest)(nil),   // 3: BroadcastRequest
	(*ClientLeaves)(nil),       // 4: ClientLeaves
	(*ServerClientLeaves)(nil), // 5: ServerClientLeaves
}
var file_grpc_proto_proto_depIdxs = []int32{
	0, // 0: ITUDatabase.Join:input_type -> JoinRequest
	3, // 1: ITUDatabase.Broadcast:input_type -> BroadcastRequest
	4, // 2: ITUDatabase.ClientLeaving:input_type -> ClientLeaves
	1, // 3: ITUDatabase.Join:output_type -> JoinResponce
	2, // 4: ITUDatabase.Broadcast:output_type -> ServerBroadcast
	5, // 5: ITUDatabase.ClientLeaving:output_type -> ServerClientLeaves
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_proto_init() }
func file_grpc_proto_proto_init() {
	if File_grpc_proto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_proto_goTypes,
		DependencyIndexes: file_grpc_proto_proto_depIdxs,
		MessageInfos:      file_grpc_proto_proto_msgTypes,
	}.Build()
	File_grpc_proto_proto = out.File
	file_grpc_proto_proto_rawDesc = nil
	file_grpc_proto_proto_goTypes = nil
	file_grpc_proto_proto_depIdxs = nil
}
