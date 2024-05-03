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

	FileName              string            `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	ChunkDataArray        [][]byte          `protobuf:"bytes,2,rep,name=chunk_data_array,json=chunkDataArray,proto3" json:"chunk_data_array,omitempty"`
	ChunkNameArray        []string          `protobuf:"bytes,3,rep,name=chunk_name_array,json=chunkNameArray,proto3" json:"chunk_name_array,omitempty"`
	ReplicaNameArray      []string          `protobuf:"bytes,4,rep,name=replica_name_array,json=replicaNameArray,proto3" json:"replica_name_array,omitempty"`
	ReplicaChunkNameArray []string          `protobuf:"bytes,5,rep,name=replica_chunk_name_array,json=replicaChunkNameArray,proto3" json:"replica_chunk_name_array,omitempty"`
	Action                string            `protobuf:"bytes,6,opt,name=action,proto3" json:"action,omitempty"`
	Fellow_SNNamesList    []string          `protobuf:"bytes,7,rep,name=fellow_SNNames_list,json=fellowSNNamesList,proto3" json:"fellow_SNNames_list,omitempty"`
	FellowsnPortMap       map[string]string `protobuf:"bytes,8,rep,name=fellowsnPortMap,proto3" json:"fellowsnPortMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ClientsnPortMap       map[string]string `protobuf:"bytes,9,rep,name=clientsnPortMap,proto3" json:"clientsnPortMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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

func (x *ChunkDetails) GetChunkDataArray() [][]byte {
	if x != nil {
		return x.ChunkDataArray
	}
	return nil
}

func (x *ChunkDetails) GetChunkNameArray() []string {
	if x != nil {
		return x.ChunkNameArray
	}
	return nil
}

func (x *ChunkDetails) GetReplicaNameArray() []string {
	if x != nil {
		return x.ReplicaNameArray
	}
	return nil
}

func (x *ChunkDetails) GetReplicaChunkNameArray() []string {
	if x != nil {
		return x.ReplicaChunkNameArray
	}
	return nil
}

func (x *ChunkDetails) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ChunkDetails) GetFellow_SNNamesList() []string {
	if x != nil {
		return x.Fellow_SNNamesList
	}
	return nil
}

func (x *ChunkDetails) GetFellowsnPortMap() map[string]string {
	if x != nil {
		return x.FellowsnPortMap
	}
	return nil
}

func (x *ChunkDetails) GetClientsnPortMap() map[string]string {
	if x != nil {
		return x.ClientsnPortMap
	}
	return nil
}

var File_clientSN_proto protoreflect.FileDescriptor

var file_clientSN_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd2, 0x04, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x72, 0x72,
	0x61, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x12, 0x37, 0x0a, 0x18, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x15, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x53, 0x4e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11,
	0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x4c, 0x0a, 0x0f, 0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x6e, 0x50, 0x6f, 0x72,
	0x74, 0x4d, 0x61, 0x70, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x46, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f,
	0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x12,
	0x4c, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d,
	0x61, 0x70, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x6e,
	0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x1a, 0x42, 0x0a,
	0x14, 0x46, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x42, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x6e, 0x50, 0x6f, 0x72,
	0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x4e, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_clientSN_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_clientSN_proto_goTypes = []interface{}{
	(*ChunkDetails)(nil), // 0: ChunkDetails
	nil,                  // 1: ChunkDetails.FellowsnPortMapEntry
	nil,                  // 2: ChunkDetails.ClientsnPortMapEntry
}
var file_clientSN_proto_depIdxs = []int32{
	1, // 0: ChunkDetails.fellowsnPortMap:type_name -> ChunkDetails.FellowsnPortMapEntry
	2, // 1: ChunkDetails.clientsnPortMap:type_name -> ChunkDetails.ClientsnPortMapEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
			NumMessages:   3,
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
