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

	ClientName int32 `protobuf:"varint,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
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

	MAX int32 `protobuf:"varint,1,opt,name=MAX,proto3" json:"MAX,omitempty"`
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

	ClientName int32  `protobuf:"varint,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
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

	ClientName int32  `protobuf:"varint,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
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

	ClientName int32 `protobuf:"varint,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	Sender     int32 `protobuf:"varint,2,opt,name=sender,proto3" json:"sender,omitempty"`
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

func (x *IAmCoordinator) GetClientName() int32 {
	if x != nil {
		return x.ClientName
	}
	return 0
}

func (x *IAmCoordinator) GetSender() int32 {
	if x != nil {
		return x.Sender
	}
	return 0
}

type SendsAllegiance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName int32 `protobuf:"varint,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
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

func (x *SendsAllegiance) GetClientName() int32 {
	if x != nil {
		return x.ClientName
	}
	return 0
}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port int32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	mi := &file_grpc_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_proto_msgTypes[6]
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
	return file_grpc_service_proto_rawDescGZIP(), []int{6}
}

func (x *JoinRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	mi := &file_grpc_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_grpc_service_proto_rawDescGZIP(), []int{7}
}

func (x *JoinResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_grpc_service_proto protoreflect.FileDescriptor

var file_grpc_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x41, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x4d, 0x41, 0x58, 0x22, 0x45, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x53,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x0a, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x53, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x48, 0x0a, 0x0e, 0x49, 0x41, 0x6d, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x0f,
	0x53, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x6c, 0x6c, 0x65, 0x67, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x21, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xbf, 0x01, 0x0a,
	0x0b, 0x49, 0x54, 0x55, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0a,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x0c, 0x2e, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x08, 0x45, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x07, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x26, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12,
	0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x53, 0x1a, 0x0b, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x53, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0f, 0x2e, 0x49, 0x41, 0x6d, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x10, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x73, 0x41, 0x6c, 0x6c, 0x65, 0x67, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x42, 0x0e,
	0x5a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_service_proto_goTypes = []any{
	(*RequestElection)(nil), // 0: RequestElection
	(*Answer)(nil),          // 1: Answer
	(*RequestCS)(nil),       // 2: RequestCS
	(*ResponseCS)(nil),      // 3: ResponseCS
	(*IAmCoordinator)(nil),  // 4: IAmCoordinator
	(*SendsAllegiance)(nil), // 5: SendsAllegiance
	(*JoinRequest)(nil),     // 6: JoinRequest
	(*JoinResponse)(nil),    // 7: JoinResponse
}
var file_grpc_service_proto_depIdxs = []int32{
	6, // 0: ITUDatabase.NotifyJoin:input_type -> JoinRequest
	0, // 1: ITUDatabase.Election:input_type -> RequestElection
	2, // 2: ITUDatabase.Broadcast:input_type -> RequestCS
	4, // 3: ITUDatabase.Coordinator:input_type -> IAmCoordinator
	7, // 4: ITUDatabase.NotifyJoin:output_type -> JoinResponse
	1, // 5: ITUDatabase.Election:output_type -> Answer
	3, // 6: ITUDatabase.Broadcast:output_type -> ResponseCS
	5, // 7: ITUDatabase.Coordinator:output_type -> SendsAllegiance
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
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
			NumMessages:   8,
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
