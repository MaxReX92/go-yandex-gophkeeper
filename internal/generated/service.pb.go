// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: service.proto

package generated

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

type EventType int32

const (
	EventType_INITIAL EventType = 0
	EventType_ADD     EventType = 1
	EventType_EDIT    EventType = 2
	EventType_REMOVE  EventType = 3
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "INITIAL",
		1: "ADD",
		2: "EDIT",
		3: "REMOVE",
	}
	EventType_value = map[string]int32{
		"INITIAL": 0,
		"ADD":     1,
		"EDIT":    2,
		"REMOVE":  3,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type SecretType int32

const (
	SecretType_BINARY     SecretType = 0
	SecretType_CARD       SecretType = 1
	SecretType_CREDENTIAL SecretType = 2
	SecretType_NOTE       SecretType = 3
)

// Enum value maps for SecretType.
var (
	SecretType_name = map[int32]string{
		0: "BINARY",
		1: "CARD",
		2: "CREDENTIAL",
		3: "NOTE",
	}
	SecretType_value = map[string]int32{
		"BINARY":     0,
		"CARD":       1,
		"CREDENTIAL": 2,
		"NOTE":       3,
	}
)

func (x SecretType) Enum() *SecretType {
	p := new(SecretType)
	*p = x
	return p
}

func (x SecretType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecretType) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[1].Descriptor()
}

func (SecretType) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[1]
}

func (x SecretType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecretType.Descriptor instead.
func (SecretType) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Content  []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *Secret) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Secret) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type SecretEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   EventType `protobuf:"varint,1,opt,name=type,proto3,enum=com.github.MaxReX92.go_yandex_gophkeeper.EventType" json:"type,omitempty"`
	Secret *Secret   `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *SecretEvent) Reset() {
	*x = SecretEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretEvent) ProtoMessage() {}

func (x *SecretEvent) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretEvent.ProtoReflect.Descriptor instead.
func (*SecretEvent) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *SecretEvent) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_INITIAL
}

func (x *SecretEvent) GetSecret() *Secret {
	if x != nil {
		return x.Secret
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

type SecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Secret *Secret `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *SecretRequest) Reset() {
	*x = SecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretRequest) ProtoMessage() {}

func (x *SecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretRequest.ProtoReflect.Descriptor instead.
func (*SecretRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *SecretRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SecretRequest) GetSecret() *Secret {
	if x != nil {
		return x.Secret
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52,
	0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67,
	0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69,
	0x64, 0x22, 0x3e, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0xa0, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x47, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78,
	0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f,
	0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32,
	0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x22, 0x22, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x9d, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67,
	0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x48,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52,
	0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67,
	0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2a, 0x37, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x45,
	0x44, 0x49, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10,
	0x03, 0x2a, 0x3c, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x43,
	0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54,
	0x49, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x54, 0x45, 0x10, 0x03, 0x32,
	0xd8, 0x04, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x66, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e,
	0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e,
	0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x74, 0x0a, 0x09, 0x41, 0x64, 0x64,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78,
	0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f,
	0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12,
	0x77, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12,
	0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78,
	0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f,
	0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67,
	0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x77, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67,
	0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d,
	0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x12, 0x77, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d,
	0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d,
	0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x42, 0x14, 0x5a, 0x12, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_proto_goTypes = []interface{}{
	(EventType)(0),        // 0: com.github.MaxReX92.go_yandex_gophkeeper.EventType
	(SecretType)(0),       // 1: com.github.MaxReX92.go_yandex_gophkeeper.SecretType
	(*Void)(nil),          // 2: com.github.MaxReX92.go_yandex_gophkeeper.Void
	(*Secret)(nil),        // 3: com.github.MaxReX92.go_yandex_gophkeeper.Secret
	(*SecretEvent)(nil),   // 4: com.github.MaxReX92.go_yandex_gophkeeper.SecretEvent
	(*User)(nil),          // 5: com.github.MaxReX92.go_yandex_gophkeeper.User
	(*SecretRequest)(nil), // 6: com.github.MaxReX92.go_yandex_gophkeeper.SecretRequest
}
var file_service_proto_depIdxs = []int32{
	0, // 0: com.github.MaxReX92.go_yandex_gophkeeper.SecretEvent.type:type_name -> com.github.MaxReX92.go_yandex_gophkeeper.EventType
	3, // 1: com.github.MaxReX92.go_yandex_gophkeeper.SecretEvent.secret:type_name -> com.github.MaxReX92.go_yandex_gophkeeper.Secret
	5, // 2: com.github.MaxReX92.go_yandex_gophkeeper.SecretRequest.user:type_name -> com.github.MaxReX92.go_yandex_gophkeeper.User
	3, // 3: com.github.MaxReX92.go_yandex_gophkeeper.SecretRequest.secret:type_name -> com.github.MaxReX92.go_yandex_gophkeeper.Secret
	2, // 4: com.github.MaxReX92.go_yandex_gophkeeper.SecretService.Ping:input_type -> com.github.MaxReX92.go_yandex_gophkeeper.Void
	6, // 5: com.github.MaxReX92.go_yandex_gophkeeper.SecretService.AddSecret:input_type -> com.github.MaxReX92.go_yandex_gophkeeper.SecretRequest
	6, // 6: com.github.MaxReX92.go_yandex_gophkeeper.SecretService.ChangeSecret:input_type -> com.github.MaxReX92.go_yandex_gophkeeper.SecretRequest
	6, // 7: com.github.MaxReX92.go_yandex_gophkeeper.SecretService.RemoveSecret:input_type -> com.github.MaxReX92.go_yandex_gophkeeper.SecretRequest
	5, // 8: com.github.MaxReX92.go_yandex_gophkeeper.SecretService.SecretEvents:input_type -> com.github.MaxReX92.go_yandex_gophkeeper.User
	2, // 9: com.github.MaxReX92.go_yandex_gophkeeper.SecretService.Ping:output_type -> com.github.MaxReX92.go_yandex_gophkeeper.Void
	2, // 10: com.github.MaxReX92.go_yandex_gophkeeper.SecretService.AddSecret:output_type -> com.github.MaxReX92.go_yandex_gophkeeper.Void
	2, // 11: com.github.MaxReX92.go_yandex_gophkeeper.SecretService.ChangeSecret:output_type -> com.github.MaxReX92.go_yandex_gophkeeper.Void
	2, // 12: com.github.MaxReX92.go_yandex_gophkeeper.SecretService.RemoveSecret:output_type -> com.github.MaxReX92.go_yandex_gophkeeper.Void
	4, // 13: com.github.MaxReX92.go_yandex_gophkeeper.SecretService.SecretEvents:output_type -> com.github.MaxReX92.go_yandex_gophkeeper.SecretEvent
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretEvent); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretRequest); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		EnumInfos:         file_service_proto_enumTypes,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
