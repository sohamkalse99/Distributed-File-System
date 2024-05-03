// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.14.0
// source: client.proto

package clientHandler

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

type FileOpns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName       string                          `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Action         string                          `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	ChunkSize      int64                           `protobuf:"varint,3,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	FileSize       int64                           `protobuf:"varint,4,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	Checksum       []byte                          `protobuf:"bytes,5,opt,name=checksum,proto3" json:"checksum,omitempty"`
	DstSNList      []string                        `protobuf:"bytes,6,rep,name=dstSNList,proto3" json:"dstSNList,omitempty"`
	SnChunkMap     map[string]*FileOpnsChunks      `protobuf:"bytes,7,rep,name=snChunkMap,proto3" json:"snChunkMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ReplicaNameMap map[string]*FileOpnsReplicaName `protobuf:"bytes,8,rep,name=replicaNameMap,proto3" json:"replicaNameMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Might be used for get
	ReplicaSNChunkMap map[string]*FileOpnsReplicaList   `protobuf:"bytes,9,rep,name=replicaSNChunkMap,proto3" json:"replicaSNChunkMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FellowSNNamesMap  map[string]*FileOpnsFellowSNNames `protobuf:"bytes,10,rep,name=fellowSNNamesMap,proto3" json:"fellowSNNamesMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FellowsnPortMap   map[string]string                 `protobuf:"bytes,11,rep,name=fellowsnPortMap,proto3" json:"fellowsnPortMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ClientsnPortMap   map[string]string                 `protobuf:"bytes,12,rep,name=clientsnPortMap,proto3" json:"clientsnPortMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FileOpns) Reset() {
	*x = FileOpns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOpns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOpns) ProtoMessage() {}

func (x *FileOpns) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOpns.ProtoReflect.Descriptor instead.
func (*FileOpns) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *FileOpns) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileOpns) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *FileOpns) GetChunkSize() int64 {
	if x != nil {
		return x.ChunkSize
	}
	return 0
}

func (x *FileOpns) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FileOpns) GetChecksum() []byte {
	if x != nil {
		return x.Checksum
	}
	return nil
}

func (x *FileOpns) GetDstSNList() []string {
	if x != nil {
		return x.DstSNList
	}
	return nil
}

func (x *FileOpns) GetSnChunkMap() map[string]*FileOpnsChunks {
	if x != nil {
		return x.SnChunkMap
	}
	return nil
}

func (x *FileOpns) GetReplicaNameMap() map[string]*FileOpnsReplicaName {
	if x != nil {
		return x.ReplicaNameMap
	}
	return nil
}

func (x *FileOpns) GetReplicaSNChunkMap() map[string]*FileOpnsReplicaList {
	if x != nil {
		return x.ReplicaSNChunkMap
	}
	return nil
}

func (x *FileOpns) GetFellowSNNamesMap() map[string]*FileOpnsFellowSNNames {
	if x != nil {
		return x.FellowSNNamesMap
	}
	return nil
}

func (x *FileOpns) GetFellowsnPortMap() map[string]string {
	if x != nil {
		return x.FellowsnPortMap
	}
	return nil
}

func (x *FileOpns) GetClientsnPortMap() map[string]string {
	if x != nil {
		return x.ClientsnPortMap
	}
	return nil
}

type FileOpnsChunks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkList []string `protobuf:"bytes,1,rep,name=chunkList,proto3" json:"chunkList,omitempty"`
}

func (x *FileOpnsChunks) Reset() {
	*x = FileOpnsChunks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOpnsChunks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOpnsChunks) ProtoMessage() {}

func (x *FileOpnsChunks) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOpnsChunks.ProtoReflect.Descriptor instead.
func (*FileOpnsChunks) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0, 1}
}

func (x *FileOpnsChunks) GetChunkList() []string {
	if x != nil {
		return x.ChunkList
	}
	return nil
}

type FileOpnsReplicaName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplicaNameList []string `protobuf:"bytes,1,rep,name=replicaNameList,proto3" json:"replicaNameList,omitempty"`
}

func (x *FileOpnsReplicaName) Reset() {
	*x = FileOpnsReplicaName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOpnsReplicaName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOpnsReplicaName) ProtoMessage() {}

func (x *FileOpnsReplicaName) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOpnsReplicaName.ProtoReflect.Descriptor instead.
func (*FileOpnsReplicaName) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0, 3}
}

func (x *FileOpnsReplicaName) GetReplicaNameList() []string {
	if x != nil {
		return x.ReplicaNameList
	}
	return nil
}

type FileOpnsReplicaList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplicaChunkList []string `protobuf:"bytes,1,rep,name=replica_chunk_list,json=replicaChunkList,proto3" json:"replica_chunk_list,omitempty"`
}

func (x *FileOpnsReplicaList) Reset() {
	*x = FileOpnsReplicaList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOpnsReplicaList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOpnsReplicaList) ProtoMessage() {}

func (x *FileOpnsReplicaList) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOpnsReplicaList.ProtoReflect.Descriptor instead.
func (*FileOpnsReplicaList) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0, 5}
}

func (x *FileOpnsReplicaList) GetReplicaChunkList() []string {
	if x != nil {
		return x.ReplicaChunkList
	}
	return nil
}

type FileOpnsFellowSNNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fellow_SNNamesList []string `protobuf:"bytes,1,rep,name=fellow_SNNames_list,json=fellowSNNamesList,proto3" json:"fellow_SNNames_list,omitempty"`
}

func (x *FileOpnsFellowSNNames) Reset() {
	*x = FileOpnsFellowSNNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOpnsFellowSNNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOpnsFellowSNNames) ProtoMessage() {}

func (x *FileOpnsFellowSNNames) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOpnsFellowSNNames.ProtoReflect.Descriptor instead.
func (*FileOpnsFellowSNNames) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0, 7}
}

func (x *FileOpnsFellowSNNames) GetFellow_SNNamesList() []string {
	if x != nil {
		return x.Fellow_SNNamesList
	}
	return nil
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7,
	0x0a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x73, 0x74, 0x53,
	0x4e, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x73, 0x74,
	0x53, 0x4e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x6e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x4d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4f, 0x70, 0x6e, 0x73, 0x2e, 0x53, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x73, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x61,
	0x70, 0x12, 0x45, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4e, 0x61, 0x6d, 0x65,
	0x4d, 0x61, 0x70, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4f, 0x70, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4e, 0x61, 0x6d, 0x65,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x4e, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x53, 0x4e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x61, 0x70, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x6e, 0x73, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x4e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x4e,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x61, 0x70, 0x12, 0x4b, 0x0a, 0x10, 0x66, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x53, 0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x4d, 0x61, 0x70, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x6e, 0x73, 0x2e, 0x46, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x10, 0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x4d, 0x61, 0x70, 0x12, 0x48, 0x0a, 0x0f, 0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x73,
	0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x6e, 0x73, 0x2e, 0x46, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f,
	0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x12,
	0x48, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d,
	0x61, 0x70, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f,
	0x70, 0x6e, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x1a, 0x4f, 0x0a, 0x0f, 0x53, 0x6e, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x6e, 0x73, 0x2e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x26, 0x0a, 0x06, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x1a, 0x58, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4e, 0x61, 0x6d,
	0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4f, 0x70, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x0b,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4e, 0x61, 0x6d,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x5c, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x53, 0x4e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x5c, 0x0a, 0x15, 0x46, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x4f, 0x70, 0x6e, 0x73, 0x2e, 0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3f, 0x0a, 0x0d, 0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x2e, 0x0a, 0x13, 0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x53, 0x4e, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x66,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x42, 0x0a, 0x14, 0x46, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x42, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x6e,
	0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_client_proto_goTypes = []interface{}{
	(*FileOpns)(nil),              // 0: FileOpns
	nil,                           // 1: FileOpns.SnChunkMapEntry
	(*FileOpnsChunks)(nil),        // 2: FileOpns.chunks
	nil,                           // 3: FileOpns.ReplicaNameMapEntry
	(*FileOpnsReplicaName)(nil),   // 4: FileOpns.replicaName
	nil,                           // 5: FileOpns.ReplicaSNChunkMapEntry
	(*FileOpnsReplicaList)(nil),   // 6: FileOpns.replica_list
	nil,                           // 7: FileOpns.FellowSNNamesMapEntry
	(*FileOpnsFellowSNNames)(nil), // 8: FileOpns.fellowSNNames
	nil,                           // 9: FileOpns.FellowsnPortMapEntry
	nil,                           // 10: FileOpns.ClientsnPortMapEntry
}
var file_client_proto_depIdxs = []int32{
	1,  // 0: FileOpns.snChunkMap:type_name -> FileOpns.SnChunkMapEntry
	3,  // 1: FileOpns.replicaNameMap:type_name -> FileOpns.ReplicaNameMapEntry
	5,  // 2: FileOpns.replicaSNChunkMap:type_name -> FileOpns.ReplicaSNChunkMapEntry
	7,  // 3: FileOpns.fellowSNNamesMap:type_name -> FileOpns.FellowSNNamesMapEntry
	9,  // 4: FileOpns.fellowsnPortMap:type_name -> FileOpns.FellowsnPortMapEntry
	10, // 5: FileOpns.clientsnPortMap:type_name -> FileOpns.ClientsnPortMapEntry
	2,  // 6: FileOpns.SnChunkMapEntry.value:type_name -> FileOpns.chunks
	4,  // 7: FileOpns.ReplicaNameMapEntry.value:type_name -> FileOpns.replicaName
	6,  // 8: FileOpns.ReplicaSNChunkMapEntry.value:type_name -> FileOpns.replica_list
	8,  // 9: FileOpns.FellowSNNamesMapEntry.value:type_name -> FileOpns.fellowSNNames
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOpns); i {
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
		file_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOpnsChunks); i {
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
		file_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOpnsReplicaName); i {
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
		file_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOpnsReplicaList); i {
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
		file_client_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOpnsFellowSNNames); i {
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
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}
