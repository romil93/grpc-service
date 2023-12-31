// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: logger/logger.proto

package logger

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

type Status_StatusCode int32

const (
	Status_UNKOWN Status_StatusCode = 0
	Status_OK     Status_StatusCode = 1
	Status_NOT_OK Status_StatusCode = 2
)

// Enum value maps for Status_StatusCode.
var (
	Status_StatusCode_name = map[int32]string{
		0: "UNKOWN",
		1: "OK",
		2: "NOT_OK",
	}
	Status_StatusCode_value = map[string]int32{
		"UNKOWN": 0,
		"OK":     1,
		"NOT_OK": 2,
	}
)

func (x Status_StatusCode) Enum() *Status_StatusCode {
	p := new(Status_StatusCode)
	*p = x
	return p
}

func (x Status_StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_logger_logger_proto_enumTypes[0].Descriptor()
}

func (Status_StatusCode) Type() protoreflect.EnumType {
	return &file_logger_logger_proto_enumTypes[0]
}

func (x Status_StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_StatusCode.Descriptor instead.
func (Status_StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_logger_logger_proto_rawDescGZIP(), []int{1, 0}
}

type LoggerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UseCase string `protobuf:"bytes,1,opt,name=use_case,json=useCase,proto3" json:"use_case,omitempty"`
	Log     string `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *LoggerMessage) Reset() {
	*x = LoggerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggerMessage) ProtoMessage() {}

func (x *LoggerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_logger_logger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggerMessage.ProtoReflect.Descriptor instead.
func (*LoggerMessage) Descriptor() ([]byte, []int) {
	return file_logger_logger_proto_rawDescGZIP(), []int{0}
}

func (x *LoggerMessage) GetUseCase() string {
	if x != nil {
		return x.UseCase
	}
	return ""
}

func (x *LoggerMessage) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode Status_StatusCode `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3,enum=logger.Status_StatusCode" json:"status_code,omitempty"`
	Errors     []string          `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_logger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_logger_logger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_logger_logger_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetStatusCode() Status_StatusCode {
	if x != nil {
		return x.StatusCode
	}
	return Status_UNKOWN
}

func (x *Status) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_logger_logger_proto protoreflect.FileDescriptor

var file_logger_logger_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x22, 0x3c, 0x0a,
	0x0d, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x75, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0x8a, 0x01, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x2c, 0x0a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x4b, 0x10, 0x02, 0x32, 0x3f, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x12, 0x35, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x15, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x5f, 0x0a, 0x26, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x6f, 0x6d, 0x69, 0x6c, 0x39, 0x33, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x42, 0x0b, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x6f, 0x6d, 0x69, 0x6c, 0x39, 0x33, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_logger_logger_proto_rawDescOnce sync.Once
	file_logger_logger_proto_rawDescData = file_logger_logger_proto_rawDesc
)

func file_logger_logger_proto_rawDescGZIP() []byte {
	file_logger_logger_proto_rawDescOnce.Do(func() {
		file_logger_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_logger_logger_proto_rawDescData)
	})
	return file_logger_logger_proto_rawDescData
}

var file_logger_logger_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_logger_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_logger_logger_proto_goTypes = []interface{}{
	(Status_StatusCode)(0), // 0: logger.Status.StatusCode
	(*LoggerMessage)(nil),  // 1: logger.LoggerMessage
	(*Status)(nil),         // 2: logger.Status
}
var file_logger_logger_proto_depIdxs = []int32{
	0, // 0: logger.Status.status_code:type_name -> logger.Status.StatusCode
	1, // 1: logger.Logger.LogMessage:input_type -> logger.LoggerMessage
	2, // 2: logger.Logger.LogMessage:output_type -> logger.Status
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_logger_logger_proto_init() }
func file_logger_logger_proto_init() {
	if File_logger_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logger_logger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggerMessage); i {
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
		file_logger_logger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_logger_logger_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logger_logger_proto_goTypes,
		DependencyIndexes: file_logger_logger_proto_depIdxs,
		EnumInfos:         file_logger_logger_proto_enumTypes,
		MessageInfos:      file_logger_logger_proto_msgTypes,
	}.Build()
	File_logger_logger_proto = out.File
	file_logger_logger_proto_rawDesc = nil
	file_logger_logger_proto_goTypes = nil
	file_logger_logger_proto_depIdxs = nil
}
