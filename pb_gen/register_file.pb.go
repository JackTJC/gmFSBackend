// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: register_file.proto

package pb_gen

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

type RegisterFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId  int64    `protobuf:"varint,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	DirId   int64    `protobuf:"varint,2,opt,name=dirId,proto3" json:"dirId,omitempty"`
	BaseReq *BaseReq `protobuf:"bytes,255,opt,name=baseReq,proto3" json:"baseReq,omitempty"`
}

func (x *RegisterFileRequest) Reset() {
	*x = RegisterFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterFileRequest) ProtoMessage() {}

func (x *RegisterFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_register_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterFileRequest.ProtoReflect.Descriptor instead.
func (*RegisterFileRequest) Descriptor() ([]byte, []int) {
	return file_register_file_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterFileRequest) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *RegisterFileRequest) GetDirId() int64 {
	if x != nil {
		return x.DirId
	}
	return 0
}

func (x *RegisterFileRequest) GetBaseReq() *BaseReq {
	if x != nil {
		return x.BaseReq
	}
	return nil
}

type RegisterFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,255,opt,name=baseResp,proto3" json:"baseResp,omitempty"`
}

func (x *RegisterFileResponse) Reset() {
	*x = RegisterFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterFileResponse) ProtoMessage() {}

func (x *RegisterFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_register_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterFileResponse.ProtoReflect.Descriptor instead.
func (*RegisterFileResponse) Descriptor() ([]byte, []int) {
	return file_register_file_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterFileResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

var File_register_file_proto protoreflect.FileDescriptor

var file_register_file_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x69, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x64, 0x69, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x22, 0x3e, 0x0a,
	0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x4b, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x75, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x42, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69,
	0x6c, 0x65, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a,
	0x61, 0x63, 0x6b, 0x54, 0x4a, 0x43, 0x2f, 0x67, 0x6d, 0x46, 0x53, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_register_file_proto_rawDescOnce sync.Once
	file_register_file_proto_rawDescData = file_register_file_proto_rawDesc
)

func file_register_file_proto_rawDescGZIP() []byte {
	file_register_file_proto_rawDescOnce.Do(func() {
		file_register_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_register_file_proto_rawDescData)
	})
	return file_register_file_proto_rawDescData
}

var file_register_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_register_file_proto_goTypes = []interface{}{
	(*RegisterFileRequest)(nil),  // 0: RegisterFileRequest
	(*RegisterFileResponse)(nil), // 1: RegisterFileResponse
	(*BaseReq)(nil),              // 2: BaseReq
	(*BaseResp)(nil),             // 3: BaseResp
}
var file_register_file_proto_depIdxs = []int32{
	2, // 0: RegisterFileRequest.baseReq:type_name -> BaseReq
	3, // 1: RegisterFileResponse.baseResp:type_name -> BaseResp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_register_file_proto_init() }
func file_register_file_proto_init() {
	if File_register_file_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_register_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterFileRequest); i {
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
		file_register_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterFileResponse); i {
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
			RawDescriptor: file_register_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_register_file_proto_goTypes,
		DependencyIndexes: file_register_file_proto_depIdxs,
		MessageInfos:      file_register_file_proto_msgTypes,
	}.Build()
	File_register_file_proto = out.File
	file_register_file_proto_rawDesc = nil
	file_register_file_proto_goTypes = nil
	file_register_file_proto_depIdxs = nil
}
