// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: wgSvc.proto

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

// The request message
type GetEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEndpointRequest) Reset() {
	*x = GetEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgSvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEndpointRequest) ProtoMessage() {}

func (x *GetEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgSvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEndpointRequest.ProtoReflect.Descriptor instead.
func (*GetEndpointRequest) Descriptor() ([]byte, []int) {
	return file_wgSvc_proto_rawDescGZIP(), []int{0}
}

type GetEndpointReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEndpointReply) Reset() {
	*x = GetEndpointReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgSvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEndpointReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEndpointReply) ProtoMessage() {}

func (x *GetEndpointReply) ProtoReflect() protoreflect.Message {
	mi := &file_wgSvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEndpointReply.ProtoReflect.Descriptor instead.
func (*GetEndpointReply) Descriptor() ([]byte, []int) {
	return file_wgSvc_proto_rawDescGZIP(), []int{1}
}

type RegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegRequest) Reset() {
	*x = RegRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgSvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegRequest) ProtoMessage() {}

func (x *RegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgSvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegRequest.ProtoReflect.Descriptor instead.
func (*RegRequest) Descriptor() ([]byte, []int) {
	return file_wgSvc_proto_rawDescGZIP(), []int{2}
}

type RegReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegReply) Reset() {
	*x = RegReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgSvc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegReply) ProtoMessage() {}

func (x *RegReply) ProtoReflect() protoreflect.Message {
	mi := &file_wgSvc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegReply.ProtoReflect.Descriptor instead.
func (*RegReply) Descriptor() ([]byte, []int) {
	return file_wgSvc_proto_rawDescGZIP(), []int{3}
}

var File_wgSvc_proto protoreflect.FileDescriptor

var file_wgSvc_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x77, 0x67, 0x53, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x77,
	0x67, 0x53, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x0c, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa1,
	0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x4f, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x77, 0x67, 0x53,
	0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x77, 0x67,
	0x53, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x10,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x17, 0x2e, 0x77, 0x67, 0x53, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x77, 0x67, 0x53, 0x76,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x77, 0x67, 0x2d, 0x77, 0x6f, 0x72, 0x6d, 0x68, 0x6f, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wgSvc_proto_rawDescOnce sync.Once
	file_wgSvc_proto_rawDescData = file_wgSvc_proto_rawDesc
)

func file_wgSvc_proto_rawDescGZIP() []byte {
	file_wgSvc_proto_rawDescOnce.Do(func() {
		file_wgSvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgSvc_proto_rawDescData)
	})
	return file_wgSvc_proto_rawDescData
}

var file_wgSvc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wgSvc_proto_goTypes = []interface{}{
	(*GetEndpointRequest)(nil), // 0: wgSvc.proto.GetEndpointRequest
	(*GetEndpointReply)(nil),   // 1: wgSvc.proto.GetEndpointReply
	(*RegRequest)(nil),         // 2: wgSvc.proto.RegRequest
	(*RegReply)(nil),           // 3: wgSvc.proto.RegReply
}
var file_wgSvc_proto_depIdxs = []int32{
	0, // 0: wgSvc.proto.Registry.GetEndpoint:input_type -> wgSvc.proto.GetEndpointRequest
	2, // 1: wgSvc.proto.Registry.RegisterEndpoint:input_type -> wgSvc.proto.RegRequest
	1, // 2: wgSvc.proto.Registry.GetEndpoint:output_type -> wgSvc.proto.GetEndpointReply
	3, // 3: wgSvc.proto.Registry.RegisterEndpoint:output_type -> wgSvc.proto.RegReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wgSvc_proto_init() }
func file_wgSvc_proto_init() {
	if File_wgSvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgSvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEndpointRequest); i {
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
		file_wgSvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEndpointReply); i {
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
		file_wgSvc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegRequest); i {
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
		file_wgSvc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegReply); i {
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
			RawDescriptor: file_wgSvc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wgSvc_proto_goTypes,
		DependencyIndexes: file_wgSvc_proto_depIdxs,
		MessageInfos:      file_wgSvc_proto_msgTypes,
	}.Build()
	File_wgSvc_proto = out.File
	file_wgSvc_proto_rawDesc = nil
	file_wgSvc_proto_goTypes = nil
	file_wgSvc_proto_depIdxs = nil
}
