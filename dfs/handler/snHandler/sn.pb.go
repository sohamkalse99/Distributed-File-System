// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.14.0
// source: sn.proto

package snHandler

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

type Registration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageNodeName   string `protobuf:"bytes,1,opt,name=storage_node_name,json=storageNodeName,proto3" json:"storage_node_name,omitempty"`
	StoragePortNumber string `protobuf:"bytes,2,opt,name=storage_port_number,json=storagePortNumber,proto3" json:"storage_port_number,omitempty"`
	SnFellowPN        string `protobuf:"bytes,3,opt,name=snFellowPN,proto3" json:"snFellowPN,omitempty"`
	Status            string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"` // ok or not ok
}

func (x *Registration) Reset() {
	*x = Registration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registration) ProtoMessage() {}

func (x *Registration) ProtoReflect() protoreflect.Message {
	mi := &file_sn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registration.ProtoReflect.Descriptor instead.
func (*Registration) Descriptor() ([]byte, []int) {
	return file_sn_proto_rawDescGZIP(), []int{0}
}

func (x *Registration) GetStorageNodeName() string {
	if x != nil {
		return x.StorageNodeName
	}
	return ""
}

func (x *Registration) GetStoragePortNumber() string {
	if x != nil {
		return x.StoragePortNumber
	}
	return ""
}

func (x *Registration) GetSnFellowPN() string {
	if x != nil {
		return x.SnFellowPN
	}
	return ""
}

func (x *Registration) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Add name with the heartbeat
	StorageNodeName       string `protobuf:"bytes,1,opt,name=storage_node_name,json=storageNodeName,proto3" json:"storage_node_name,omitempty"`
	StoragePortNumber     string `protobuf:"bytes,2,opt,name=storage_port_number,json=storagePortNumber,proto3" json:"storage_port_number,omitempty"`
	SnFellowPN            string `protobuf:"bytes,3,opt,name=snFellowPN,proto3" json:"snFellowPN,omitempty"`
	SpaceAvailability     uint64 `protobuf:"varint,4,opt,name=space_availability,json=spaceAvailability,proto3" json:"space_availability,omitempty"`
	RequestProcessedCount int32  `protobuf:"varint,5,opt,name=request_processed_count,json=requestProcessedCount,proto3" json:"request_processed_count,omitempty"`
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_sn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_sn_proto_rawDescGZIP(), []int{1}
}

func (x *Heartbeat) GetStorageNodeName() string {
	if x != nil {
		return x.StorageNodeName
	}
	return ""
}

func (x *Heartbeat) GetStoragePortNumber() string {
	if x != nil {
		return x.StoragePortNumber
	}
	return ""
}

func (x *Heartbeat) GetSnFellowPN() string {
	if x != nil {
		return x.SnFellowPN
	}
	return ""
}

func (x *Heartbeat) GetSpaceAvailability() uint64 {
	if x != nil {
		return x.SpaceAvailability
	}
	return 0
}

func (x *Heartbeat) GetRequestProcessedCount() int32 {
	if x != nil {
		return x.RequestProcessedCount
	}
	return 0
}

type FellowSNMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data            [][]byte          `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	ReplicaNames    []string          `protobuf:"bytes,2,rep,name=replicaNames,proto3" json:"replicaNames,omitempty"`
	ClientsnPortMap map[string]string `protobuf:"bytes,3,rep,name=clientsnPortMap,proto3" json:"clientsnPortMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SnName          string            `protobuf:"bytes,4,opt,name=snName,proto3" json:"snName,omitempty"`
}

func (x *FellowSNMsg) Reset() {
	*x = FellowSNMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FellowSNMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FellowSNMsg) ProtoMessage() {}

func (x *FellowSNMsg) ProtoReflect() protoreflect.Message {
	mi := &file_sn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FellowSNMsg.ProtoReflect.Descriptor instead.
func (*FellowSNMsg) Descriptor() ([]byte, []int) {
	return file_sn_proto_rawDescGZIP(), []int{2}
}

func (x *FellowSNMsg) GetData() [][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FellowSNMsg) GetReplicaNames() []string {
	if x != nil {
		return x.ReplicaNames
	}
	return nil
}

func (x *FellowSNMsg) GetClientsnPortMap() map[string]string {
	if x != nil {
		return x.ClientsnPortMap
	}
	return nil
}

func (x *FellowSNMsg) GetSnName() string {
	if x != nil {
		return x.SnName
	}
	return ""
}

type Wrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Task:
	//
	//	*Wrapper_RegTask
	//	*Wrapper_HeartbeatTask
	//	*Wrapper_FellowSNTask
	Task isWrapper_Task `protobuf_oneof:"task"`
}

func (x *Wrapper) Reset() {
	*x = Wrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wrapper) ProtoMessage() {}

func (x *Wrapper) ProtoReflect() protoreflect.Message {
	mi := &file_sn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wrapper.ProtoReflect.Descriptor instead.
func (*Wrapper) Descriptor() ([]byte, []int) {
	return file_sn_proto_rawDescGZIP(), []int{3}
}

func (m *Wrapper) GetTask() isWrapper_Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (x *Wrapper) GetRegTask() *Registration {
	if x, ok := x.GetTask().(*Wrapper_RegTask); ok {
		return x.RegTask
	}
	return nil
}

func (x *Wrapper) GetHeartbeatTask() *Heartbeat {
	if x, ok := x.GetTask().(*Wrapper_HeartbeatTask); ok {
		return x.HeartbeatTask
	}
	return nil
}

func (x *Wrapper) GetFellowSNTask() *FellowSNMsg {
	if x, ok := x.GetTask().(*Wrapper_FellowSNTask); ok {
		return x.FellowSNTask
	}
	return nil
}

type isWrapper_Task interface {
	isWrapper_Task()
}

type Wrapper_RegTask struct {
	RegTask *Registration `protobuf:"bytes,1,opt,name=reg_task,json=regTask,proto3,oneof"`
}

type Wrapper_HeartbeatTask struct {
	HeartbeatTask *Heartbeat `protobuf:"bytes,2,opt,name=heartbeat_task,json=heartbeatTask,proto3,oneof"`
}

type Wrapper_FellowSNTask struct {
	FellowSNTask *FellowSNMsg `protobuf:"bytes,3,opt,name=fellowSN_task,json=fellowSNTask,proto3,oneof"`
}

func (*Wrapper_RegTask) isWrapper_Task() {}

func (*Wrapper_HeartbeatTask) isWrapper_Task() {}

func (*Wrapper_FellowSNTask) isWrapper_Task() {}

var File_sn_proto protoreflect.FileDescriptor

var file_sn_proto_rawDesc = []byte{
	0x0a, 0x08, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x0c, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x72,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6e, 0x46, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x50, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x46,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x4e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xee, 0x01, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x2a, 0x0a,
	0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6e, 0x46,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x6e, 0x46, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x4e, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x17, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xee, 0x01, 0x0a, 0x0b, 0x46, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e, 0x4d, 0x73, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x46, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e, 0x4d, 0x73, 0x67, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x6e, 0x50, 0x6f,
	0x72, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x42, 0x0a,
	0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xa7, 0x01, 0x0a, 0x07, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x08, 0x72, 0x65, 0x67, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x07, 0x72, 0x65, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x33, 0x0a, 0x0e, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x48, 0x00, 0x52,
	0x0d, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x33,
	0x0a, 0x0d, 0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e,
	0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x4e, 0x54,
	0x61, 0x73, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x42, 0x15, 0x5a, 0x13, 0x2e,
	0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sn_proto_rawDescOnce sync.Once
	file_sn_proto_rawDescData = file_sn_proto_rawDesc
)

func file_sn_proto_rawDescGZIP() []byte {
	file_sn_proto_rawDescOnce.Do(func() {
		file_sn_proto_rawDescData = protoimpl.X.CompressGZIP(file_sn_proto_rawDescData)
	})
	return file_sn_proto_rawDescData
}

var file_sn_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sn_proto_goTypes = []interface{}{
	(*Registration)(nil), // 0: Registration
	(*Heartbeat)(nil),    // 1: Heartbeat
	(*FellowSNMsg)(nil),  // 2: FellowSNMsg
	(*Wrapper)(nil),      // 3: Wrapper
	nil,                  // 4: FellowSNMsg.ClientsnPortMapEntry
}
var file_sn_proto_depIdxs = []int32{
	4, // 0: FellowSNMsg.clientsnPortMap:type_name -> FellowSNMsg.ClientsnPortMapEntry
	0, // 1: Wrapper.reg_task:type_name -> Registration
	1, // 2: Wrapper.heartbeat_task:type_name -> Heartbeat
	2, // 3: Wrapper.fellowSN_task:type_name -> FellowSNMsg
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sn_proto_init() }
func file_sn_proto_init() {
	if File_sn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registration); i {
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
		file_sn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
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
		file_sn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FellowSNMsg); i {
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
		file_sn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wrapper); i {
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
	file_sn_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Wrapper_RegTask)(nil),
		(*Wrapper_HeartbeatTask)(nil),
		(*Wrapper_FellowSNTask)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sn_proto_goTypes,
		DependencyIndexes: file_sn_proto_depIdxs,
		MessageInfos:      file_sn_proto_msgTypes,
	}.Build()
	File_sn_proto = out.File
	file_sn_proto_rawDesc = nil
	file_sn_proto_goTypes = nil
	file_sn_proto_depIdxs = nil
}
