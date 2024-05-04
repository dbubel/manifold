// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: manifold.proto

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

type Priority int32

const (
	Priority_NORMAL Priority = 0
	Priority_HIGH   Priority = 1
)

// Enum value maps for Priority.
var (
	Priority_name = map[int32]string{
		0: "NORMAL",
		1: "HIGH",
	}
	Priority_value = map[string]int32{
		"NORMAL": 0,
		"HIGH":   1,
	}
)

func (x Priority) Enum() *Priority {
	p := new(Priority)
	*p = x
	return p
}

func (x Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_manifold_proto_enumTypes[0].Descriptor()
}

func (Priority) Type() protoreflect.EnumType {
	return &file_manifold_proto_enumTypes[0]
}

func (x Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Priority.Descriptor instead.
func (Priority) EnumDescriptor() ([]byte, []int) {
	return file_manifold_proto_rawDescGZIP(), []int{0}
}

type EnqueueMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicName string   `protobuf:"bytes,1,opt,name=topicName,proto3" json:"topicName,omitempty"`
	Data      []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Priority  Priority `protobuf:"varint,3,opt,name=priority,proto3,enum=manifold.Priority" json:"priority,omitempty"`
}

func (x *EnqueueMsg) Reset() {
	*x = EnqueueMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifold_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueMsg) ProtoMessage() {}

func (x *EnqueueMsg) ProtoReflect() protoreflect.Message {
	mi := &file_manifold_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueMsg.ProtoReflect.Descriptor instead.
func (*EnqueueMsg) Descriptor() ([]byte, []int) {
	return file_manifold_proto_rawDescGZIP(), []int{0}
}

func (x *EnqueueMsg) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *EnqueueMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *EnqueueMsg) GetPriority() Priority {
	if x != nil {
		return x.Priority
	}
	return Priority_NORMAL
}

type EnqueueAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EnqueueAck) Reset() {
	*x = EnqueueAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifold_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueAck) ProtoMessage() {}

func (x *EnqueueAck) ProtoReflect() protoreflect.Message {
	mi := &file_manifold_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueAck.ProtoReflect.Descriptor instead.
func (*EnqueueAck) Descriptor() ([]byte, []int) {
	return file_manifold_proto_rawDescGZIP(), []int{1}
}

func (x *EnqueueAck) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DequeueMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicName string `protobuf:"bytes,1,opt,name=topicName,proto3" json:"topicName,omitempty"`
}

func (x *DequeueMsg) Reset() {
	*x = DequeueMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifold_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DequeueMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequeueMsg) ProtoMessage() {}

func (x *DequeueMsg) ProtoReflect() protoreflect.Message {
	mi := &file_manifold_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequeueMsg.ProtoReflect.Descriptor instead.
func (*DequeueMsg) Descriptor() ([]byte, []int) {
	return file_manifold_proto_rawDescGZIP(), []int{2}
}

func (x *DequeueMsg) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

type DequeueAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DequeueAck) Reset() {
	*x = DequeueAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifold_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DequeueAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequeueAck) ProtoMessage() {}

func (x *DequeueAck) ProtoReflect() protoreflect.Message {
	mi := &file_manifold_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequeueAck.ProtoReflect.Descriptor instead.
func (*DequeueAck) Descriptor() ([]byte, []int) {
	return file_manifold_proto_rawDescGZIP(), []int{3}
}

func (x *DequeueAck) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifold_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_manifold_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_manifold_proto_rawDescGZIP(), []int{4}
}

type StringList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyMap map[string]int32 `protobuf:"bytes,1,rep,name=my_map,json=myMap,proto3" json:"my_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *StringList) Reset() {
	*x = StringList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifold_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringList) ProtoMessage() {}

func (x *StringList) ProtoReflect() protoreflect.Message {
	mi := &file_manifold_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringList.ProtoReflect.Descriptor instead.
func (*StringList) Descriptor() ([]byte, []int) {
	return file_manifold_proto_rawDescGZIP(), []int{5}
}

func (x *StringList) GetMyMap() map[string]int32 {
	if x != nil {
		return x.MyMap
	}
	return nil
}

type Length struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *Length) Reset() {
	*x = Length{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifold_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Length) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Length) ProtoMessage() {}

func (x *Length) ProtoReflect() protoreflect.Message {
	mi := &file_manifold_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Length.ProtoReflect.Descriptor instead.
func (*Length) Descriptor() ([]byte, []int) {
	return file_manifold_proto_rawDescGZIP(), []int{6}
}

func (x *Length) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type DeleteTopicMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicName string `protobuf:"bytes,1,opt,name=topicName,proto3" json:"topicName,omitempty"`
}

func (x *DeleteTopicMsg) Reset() {
	*x = DeleteTopicMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifold_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTopicMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTopicMsg) ProtoMessage() {}

func (x *DeleteTopicMsg) ProtoReflect() protoreflect.Message {
	mi := &file_manifold_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTopicMsg.ProtoReflect.Descriptor instead.
func (*DeleteTopicMsg) Descriptor() ([]byte, []int) {
	return file_manifold_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTopicMsg) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

var File_manifold_proto protoreflect.FileDescriptor

var file_manifold_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64, 0x22, 0x6e, 0x0a, 0x0a, 0x45, 0x6e,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d,
	0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x0a, 0x45, 0x6e,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2a, 0x0a, 0x0a,
	0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0a, 0x44, 0x65, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x7e, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x06, 0x6d, 0x79, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x4d, 0x79, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x6d, 0x79, 0x4d, 0x61, 0x70, 0x1a, 0x38, 0x0a, 0x0a, 0x4d, 0x79, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x20, 0x0a, 0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x2e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x20, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x01, 0x32, 0xe7, 0x02, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x69,
	0x66, 0x6f, 0x6c, 0x64, 0x12, 0x37, 0x0a, 0x07, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12,
	0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64, 0x2e, 0x45, 0x6e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64,
	0x2e, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x07, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x6f, 0x6c, 0x64, 0x2e, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x14,
	0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64, 0x2e, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f,
	0x6c, 0x64, 0x2e, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e,
	0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64, 0x2e, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x41, 0x63, 0x6b, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x12, 0x0f, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x2e, 0x6d, 0x61,
	0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x4d, 0x73, 0x67, 0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f,
	0x6c, 0x64, 0x2e, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x10, 0x2e,
	0x6d, 0x61, 0x6e, 0x69, 0x66, 0x6f, 0x6c, 0x64, 0x2e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22,
	0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manifold_proto_rawDescOnce sync.Once
	file_manifold_proto_rawDescData = file_manifold_proto_rawDesc
)

func file_manifold_proto_rawDescGZIP() []byte {
	file_manifold_proto_rawDescOnce.Do(func() {
		file_manifold_proto_rawDescData = protoimpl.X.CompressGZIP(file_manifold_proto_rawDescData)
	})
	return file_manifold_proto_rawDescData
}

var file_manifold_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_manifold_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_manifold_proto_goTypes = []interface{}{
	(Priority)(0),          // 0: manifold.Priority
	(*EnqueueMsg)(nil),     // 1: manifold.EnqueueMsg
	(*EnqueueAck)(nil),     // 2: manifold.EnqueueAck
	(*DequeueMsg)(nil),     // 3: manifold.DequeueMsg
	(*DequeueAck)(nil),     // 4: manifold.DequeueAck
	(*Empty)(nil),          // 5: manifold.Empty
	(*StringList)(nil),     // 6: manifold.StringList
	(*Length)(nil),         // 7: manifold.Length
	(*DeleteTopicMsg)(nil), // 8: manifold.DeleteTopicMsg
	nil,                    // 9: manifold.StringList.MyMapEntry
}
var file_manifold_proto_depIdxs = []int32{
	0, // 0: manifold.EnqueueMsg.priority:type_name -> manifold.Priority
	9, // 1: manifold.StringList.my_map:type_name -> manifold.StringList.MyMapEntry
	1, // 2: manifold.Manifold.Enqueue:input_type -> manifold.EnqueueMsg
	3, // 3: manifold.Manifold.Dequeue:input_type -> manifold.DequeueMsg
	3, // 4: manifold.Manifold.StreamDequeue:input_type -> manifold.DequeueMsg
	5, // 5: manifold.Manifold.ListTopics:input_type -> manifold.Empty
	8, // 6: manifold.Manifold.DeleteTopic:input_type -> manifold.DeleteTopicMsg
	3, // 7: manifold.Manifold.TopicLength:input_type -> manifold.DequeueMsg
	2, // 8: manifold.Manifold.Enqueue:output_type -> manifold.EnqueueAck
	4, // 9: manifold.Manifold.Dequeue:output_type -> manifold.DequeueAck
	4, // 10: manifold.Manifold.StreamDequeue:output_type -> manifold.DequeueAck
	6, // 11: manifold.Manifold.ListTopics:output_type -> manifold.StringList
	5, // 12: manifold.Manifold.DeleteTopic:output_type -> manifold.Empty
	7, // 13: manifold.Manifold.TopicLength:output_type -> manifold.Length
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_manifold_proto_init() }
func file_manifold_proto_init() {
	if File_manifold_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manifold_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueMsg); i {
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
		file_manifold_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueAck); i {
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
		file_manifold_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DequeueMsg); i {
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
		file_manifold_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DequeueAck); i {
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
		file_manifold_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_manifold_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringList); i {
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
		file_manifold_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Length); i {
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
		file_manifold_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTopicMsg); i {
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
			RawDescriptor: file_manifold_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manifold_proto_goTypes,
		DependencyIndexes: file_manifold_proto_depIdxs,
		EnumInfos:         file_manifold_proto_enumTypes,
		MessageInfos:      file_manifold_proto_msgTypes,
	}.Build()
	File_manifold_proto = out.File
	file_manifold_proto_rawDesc = nil
	file_manifold_proto_goTypes = nil
	file_manifold_proto_depIdxs = nil
}
