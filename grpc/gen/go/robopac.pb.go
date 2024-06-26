// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: robopac.proto

package robopac

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

type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_robopac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_robopac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_robopac_proto_rawDescGZIP(), []int{0}
}

func (x *RunRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *RunRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_robopac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_robopac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_robopac_proto_rawDescGZIP(), []int{1}
}

func (x *RunResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_robopac_proto protoreflect.FileDescriptor

var file_robopac_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x6f, 0x62, 0x6f, 0x70, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x6f, 0x62, 0x6f, 0x70, 0x61, 0x63, 0x22, 0x36, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x29, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x3b, 0x0a, 0x07, 0x52,
	0x6f, 0x62, 0x6f, 0x50, 0x61, 0x63, 0x12, 0x30, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x13, 0x2e,
	0x72, 0x6f, 0x62, 0x6f, 0x70, 0x61, 0x63, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x70, 0x61, 0x63, 0x2e, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x70, 0x61, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x70, 0x61, 0x63, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x70, 0x61, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_robopac_proto_rawDescOnce sync.Once
	file_robopac_proto_rawDescData = file_robopac_proto_rawDesc
)

func file_robopac_proto_rawDescGZIP() []byte {
	file_robopac_proto_rawDescOnce.Do(func() {
		file_robopac_proto_rawDescData = protoimpl.X.CompressGZIP(file_robopac_proto_rawDescData)
	})
	return file_robopac_proto_rawDescData
}

var file_robopac_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_robopac_proto_goTypes = []interface{}{
	(*RunRequest)(nil),  // 0: robopac.RunRequest
	(*RunResponse)(nil), // 1: robopac.RunResponse
}
var file_robopac_proto_depIdxs = []int32{
	0, // 0: robopac.RoboPac.Run:input_type -> robopac.RunRequest
	1, // 1: robopac.RoboPac.Run:output_type -> robopac.RunResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_robopac_proto_init() }
func file_robopac_proto_init() {
	if File_robopac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_robopac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunRequest); i {
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
		file_robopac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
			RawDescriptor: file_robopac_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_robopac_proto_goTypes,
		DependencyIndexes: file_robopac_proto_depIdxs,
		MessageInfos:      file_robopac_proto_msgTypes,
	}.Build()
	File_robopac_proto = out.File
	file_robopac_proto_rawDesc = nil
	file_robopac_proto_goTypes = nil
	file_robopac_proto_depIdxs = nil
}
