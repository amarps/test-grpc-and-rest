// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pb/log.proto

package Log

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

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_pb_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_pb_log_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Log) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_pb_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_pb_log_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_pb_log_proto protoreflect.FileDescriptor

var file_pb_log_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x62, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x4c, 0x6f, 0x67, 0x22, 0x2b, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x43, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x07, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x49,
	0x44, 0x1a, 0x08, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x1c, 0x0a,
	0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67,
	0x1a, 0x07, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x74,
	0x65, 0x73, 0x74, 0x47, 0x52, 0x50, 0x43, 0x61, 0x6e, 0x64, 0x52, 0x45, 0x53, 0x54, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x4c,
	0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_log_proto_rawDescOnce sync.Once
	file_pb_log_proto_rawDescData = file_pb_log_proto_rawDesc
)

func file_pb_log_proto_rawDescGZIP() []byte {
	file_pb_log_proto_rawDescOnce.Do(func() {
		file_pb_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_log_proto_rawDescData)
	})
	return file_pb_log_proto_rawDescData
}

var file_pb_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_log_proto_goTypes = []interface{}{
	(*Log)(nil), // 0: Log.Log
	(*ID)(nil),  // 1: Log.ID
}
var file_pb_log_proto_depIdxs = []int32{
	1, // 0: Log.Logger.Read:input_type -> Log.ID
	0, // 1: Log.Logger.Write:input_type -> Log.Log
	0, // 2: Log.Logger.Read:output_type -> Log.Log
	1, // 3: Log.Logger.Write:output_type -> Log.ID
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_log_proto_init() }
func file_pb_log_proto_init() {
	if File_pb_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_pb_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
			RawDescriptor: file_pb_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_log_proto_goTypes,
		DependencyIndexes: file_pb_log_proto_depIdxs,
		MessageInfos:      file_pb_log_proto_msgTypes,
	}.Build()
	File_pb_log_proto = out.File
	file_pb_log_proto_rawDesc = nil
	file_pb_log_proto_goTypes = nil
	file_pb_log_proto_depIdxs = nil
}