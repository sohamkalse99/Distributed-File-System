// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.14.0
// source: clientSN.proto

package clientSNHandler

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

type ChunkDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName   string   `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	ChunkArray [][]byte `protobuf:"bytes,2,rep,name=chunk_array,json=chunkArray,proto3" json:"chunk_array,omitempty"`
	Action     string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *ChunkDetails) Reset() {
	*x = ChunkDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientSN_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkDetails) ProtoMessage() {}

func (x *ChunkDetails) ProtoReflect() protoreflect.Message {
	mi := &file_clientSN_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkDetails.ProtoReflect.Descriptor instead.
func (*ChunkDetails) Descriptor() ([]byte, []int) {
	return file_clientSN_proto_rawDescGZIP(), []int{0}
}

func (x *ChunkDetails) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ChunkDetails) GetChunkArray() [][]byte {
	if x != nil {
		return x.ChunkArray
	}
	return nil
}

func (x *ChunkDetails) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

var File_clientSN_proto protoreflect.FileDescriptor

var file_clientSN_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x64, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x4e, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clientSN_proto_rawDescOnce sync.Once
	file_clientSN_proto_rawDescData = file_clientSN_proto_rawDesc
)

func file_clientSN_proto_rawDescGZIP() []byte {
	file_clientSN_proto_rawDescOnce.Do(func() {
		file_clientSN_proto_rawDescData = protoimpl.X.CompressGZIP(file_clientSN_proto_rawDescData)
	})
	return file_clientSN_proto_rawDescData
}

var file_clientSN_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_clientSN_proto_goTypes = []interface{}{
	(*ChunkDetails)(nil), // 0: ChunkDetails
}
var file_clientSN_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_clientSN_proto_init() }
func file_clientSN_proto_init() {
	if File_clientSN_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clientSN_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkDetails); i {
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
			RawDescriptor: file_clientSN_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_clientSN_proto_goTypes,
		DependencyIndexes: file_clientSN_proto_depIdxs,
		MessageInfos:      file_clientSN_proto_msgTypes,
	}.Build()
	File_clientSN_proto = out.File
	file_clientSN_proto_rawDesc = nil
	file_clientSN_proto_goTypes = nil
	file_clientSN_proto_depIdxs = nil
}
