// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: grpc/service.proto

package service

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

type RequestElection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTime int32 `protobuf:"varint,1,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
	ClientName  int32 `protobuf:"varint,2,opt,name=clientName,proto3" json:"clientName,omitempty"`
}

func (x *RequestElection) Reset() {
	*x = RequestElection{}
	mi := &file_grpc_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestElection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestElection) ProtoMessage() {}

func (x *RequestElection) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestElection.ProtoReflect.Descriptor instead.
func (*RequestElection) Descriptor() ([]byte, []int) {
	return file_grpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *RequestElection) GetLamportTime() int32 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

func (x *RequestElection) GetClientName() int32 {
	if x != nil {
		return x.ClientName
	}
	return 0
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTime int32 `protobuf:"varint,1,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
	MAX         int32 `protobuf:"varint,2,opt,name=MAX,proto3" json:"MAX,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	mi := &file_grpc_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_grpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *Answer) GetLamportTime() int32 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

func (x *Answer) GetMAX() int32 {
	if x != nil {
		return x.MAX
	}
	return 0
}

type RequestCS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTime int32  `protobuf:"varint,1,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
	ClientName  int32  `protobuf:"varint,2,opt,name=clientName,proto3" json:"clientName,omitempty"`
	Message     string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RequestCS) Reset() {
	*x = RequestCS{}
	mi := &file_grpc_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestCS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCS) ProtoMessage() {}

func (x *RequestCS) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCS.ProtoReflect.Descriptor instead.
func (*RequestCS) Descriptor() ([]byte, []int) {
	return file_grpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *RequestCS) GetLamportTime() int32 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

func (x *RequestCS) GetClientName() int32 {
	if x != nil {
		return x.ClientName
	}
	return 0
}

func (x *RequestCS) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResponseCS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTime int32  `protobuf:"varint,1,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
	ClientName  int32  `protobuf:"varint,2,opt,name=clientName,proto3" json:"clientName,omitempty"`
	Message     string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseCS) Reset() {
	*x = ResponseCS{}
	mi := &file_grpc_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseCS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCS) ProtoMessage() {}

func (x *ResponseCS) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCS.ProtoReflect.Descriptor instead.
func (*ResponseCS) Descriptor() ([]byte, []int) {
	return file_grpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseCS) GetLamportTime() int32 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

func (x *ResponseCS) GetClientName() int32 {
	if x != nil {
		return x.ClientName
	}
	return 0
}

func (x *ResponseCS) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IAmCoordinator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTime int32 `protobuf:"varint,1,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
	ClientName  int32 `protobuf:"varint,2,opt,name=clientName,proto3" json:"clientName,omitempty"`
}

func (x *IAmCoordinator) Reset() {
	*x = IAmCoordinator{}
	mi := &file_grpc_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IAmCoordinator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IAmCoordinator) ProtoMessage() {}

func (x *IAmCoordinator) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IAmCoordinator.ProtoReflect.Descriptor instead.
func (*IAmCoordinator) Descriptor() ([]byte, []int) {
	return file_grpc_service_proto_rawDescGZIP(), []int{4}
}

func (x *IAmCoordinator) GetLamportTime() int32 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

func (x *IAmCoordinator) GetClientName() int32 {
	if x != nil {
		return x.ClientName
	}
	return 0
}

type SendsAllegiance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTime int32 `protobuf:"varint,1,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
	ClientName  int32 `protobuf:"varint,2,opt,name=clientName,proto3" json:"clientName,omitempty"`
}

func (x *SendsAllegiance) Reset() {
	*x = SendsAllegiance{}
	mi := &file_grpc_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendsAllegiance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendsAllegiance) ProtoMessage() {}

func (x *SendsAllegiance) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendsAllegiance.ProtoReflect.Descriptor instead.
func (*SendsAllegiance) Descriptor() ([]byte, []int) {
	return file_grpc_service_proto_rawDescGZIP(), []int{5}
}

func (x *SendsAllegiance) GetLamportTime() int32 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

func (x *SendsAllegiance) GetClientName() int32 {
	if x != nil {
		return x.ClientName
	}
	return 0
}

var File_grpc_service_proto protoreflect.FileDescriptor

var file_grpc_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c,
	0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x06, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x61, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x41, 0x58, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x4d, 0x41, 0x58, 0x22, 0x68, 0x0a, 0x09, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x61,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43,
	0x53, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x53,
	0x0a, 0x0e, 0x49, 0x41, 0x6d, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x6c, 0x6c, 0x65,
	0x67, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x61,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x91, 0x01, 0x0a, 0x0b, 0x49, 0x54,
	0x55, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x45, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x07, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x25, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12,
	0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x53, 0x1a, 0x0a, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x53, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0f, 0x2e, 0x49, 0x41, 0x6d, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x10, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x73, 0x41, 0x6c, 0x6c, 0x65, 0x67, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a,
	0x0c, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_service_proto_rawDescOnce sync.Once
	file_grpc_service_proto_rawDescData = file_grpc_service_proto_rawDesc
)

func file_grpc_service_proto_rawDescGZIP() []byte {
	file_grpc_service_proto_rawDescOnce.Do(func() {
		file_grpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_service_proto_rawDescData)
	})
	return file_grpc_service_proto_rawDescData
}

var file_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_service_proto_goTypes = []any{
	(*RequestElection)(nil), // 0: RequestElection
	(*Answer)(nil),          // 1: Answer
	(*RequestCS)(nil),       // 2: RequestCS
	(*ResponseCS)(nil),      // 3: ResponseCS
	(*IAmCoordinator)(nil),  // 4: IAmCoordinator
	(*SendsAllegiance)(nil), // 5: SendsAllegiance
}
var file_grpc_service_proto_depIdxs = []int32{
	0, // 0: ITUDatabase.Election:input_type -> RequestElection
	2, // 1: ITUDatabase.Broadcast:input_type -> RequestCS
	4, // 2: ITUDatabase.Coordinator:input_type -> IAmCoordinator
	1, // 3: ITUDatabase.Election:output_type -> Answer
	2, // 4: ITUDatabase.Broadcast:output_type -> RequestCS
	5, // 5: ITUDatabase.Coordinator:output_type -> SendsAllegiance
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_service_proto_init() }
func file_grpc_service_proto_init() {
	if File_grpc_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_service_proto_goTypes,
		DependencyIndexes: file_grpc_service_proto_depIdxs,
		MessageInfos:      file_grpc_service_proto_msgTypes,
	}.Build()
	File_grpc_service_proto = out.File
	file_grpc_service_proto_rawDesc = nil
	file_grpc_service_proto_goTypes = nil
	file_grpc_service_proto_depIdxs = nil
}