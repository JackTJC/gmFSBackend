// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: get_node.proto

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

type GetNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId  int64    `protobuf:"varint,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"` // 节点id
	BaseReq *BaseReq `protobuf:"bytes,255,opt,name=baseReq,proto3" json:"baseReq,omitempty"`
}

func (x *GetNodeRequest) Reset() {
	*x = GetNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeRequest.ProtoReflect.Descriptor instead.
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return file_get_node_proto_rawDescGZIP(), []int{0}
}

func (x *GetNodeRequest) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *GetNodeRequest) GetBaseReq() *BaseReq {
	if x != nil {
		return x.BaseReq
	}
	return nil
}

type GetNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node     *Node     `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`         // 节点
	SubNodes []*Node   `protobuf:"bytes,2,rep,name=subNodes,proto3" json:"subNodes,omitempty"` // 子节点
	BaseResp *BaseResp `protobuf:"bytes,255,opt,name=baseResp,proto3" json:"baseResp,omitempty"`
}

func (x *GetNodeResponse) Reset() {
	*x = GetNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeResponse) ProtoMessage() {}

func (x *GetNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeResponse.ProtoReflect.Descriptor instead.
func (*GetNodeResponse) Descriptor() ([]byte, []int) {
	return file_get_node_proto_rawDescGZIP(), []int{1}
}

func (x *GetNodeResponse) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *GetNodeResponse) GetSubNodes() []*Node {
	if x != nil {
		return x.SubNodes
	}
	return nil
}

func (x *GetNodeResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

var File_get_node_proto protoreflect.FileDescriptor

var file_get_node_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x22, 0x77, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x08, 0x73,
	0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x73, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x46, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72,
	0x61, 0x64, 0x75, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x42, 0x07, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x63, 0x6b, 0x54, 0x4a, 0x43, 0x2f, 0x67, 0x6d, 0x46, 0x53, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_node_proto_rawDescOnce sync.Once
	file_get_node_proto_rawDescData = file_get_node_proto_rawDesc
)

func file_get_node_proto_rawDescGZIP() []byte {
	file_get_node_proto_rawDescOnce.Do(func() {
		file_get_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_node_proto_rawDescData)
	})
	return file_get_node_proto_rawDescData
}

var file_get_node_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_get_node_proto_goTypes = []interface{}{
	(*GetNodeRequest)(nil),  // 0: GetNodeRequest
	(*GetNodeResponse)(nil), // 1: GetNodeResponse
	(*BaseReq)(nil),         // 2: BaseReq
	(*Node)(nil),            // 3: Node
	(*BaseResp)(nil),        // 4: BaseResp
}
var file_get_node_proto_depIdxs = []int32{
	2, // 0: GetNodeRequest.baseReq:type_name -> BaseReq
	3, // 1: GetNodeResponse.node:type_name -> Node
	3, // 2: GetNodeResponse.subNodes:type_name -> Node
	4, // 3: GetNodeResponse.baseResp:type_name -> BaseResp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_get_node_proto_init() }
func file_get_node_proto_init() {
	if File_get_node_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_get_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeRequest); i {
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
		file_get_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeResponse); i {
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
			RawDescriptor: file_get_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_get_node_proto_goTypes,
		DependencyIndexes: file_get_node_proto_depIdxs,
		MessageInfos:      file_get_node_proto_msgTypes,
	}.Build()
	File_get_node_proto = out.File
	file_get_node_proto_rawDesc = nil
	file_get_node_proto_goTypes = nil
	file_get_node_proto_depIdxs = nil
}
