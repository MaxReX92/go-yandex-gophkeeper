// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: authService.proto

package generated

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_authService_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_authService_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_authService_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity      string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Token         string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	PersonalToken string `protobuf:"bytes,3,opt,name=personalToken,proto3" json:"personalToken,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_authService_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetPersonalToken() string {
	if x != nil {
		return x.PersonalToken
	}
	return ""
}

type ProlongTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ProlongTokenRequest) Reset() {
	*x = ProlongTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProlongTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProlongTokenRequest) ProtoMessage() {}

func (x *ProlongTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProlongTokenRequest.ProtoReflect.Descriptor instead.
func (*ProlongTokenRequest) Descriptor() ([]byte, []int) {
	return file_authService_proto_rawDescGZIP(), []int{4}
}

func (x *ProlongTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ProlongTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ProlongTokenResponse) Reset() {
	*x = ProlongTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProlongTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProlongTokenResponse) ProtoMessage() {}

func (x *ProlongTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProlongTokenResponse.ProtoReflect.Descriptor instead.
func (*ProlongTokenResponse) Descriptor() ([]byte, []int) {
	return file_authService_proto_rawDescGZIP(), []int{5}
}

func (x *ProlongTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_authService_proto protoreflect.FileDescriptor

var file_authService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x4d, 0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x22, 0x41, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x2e, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x3e, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x67, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2b, 0x0a, 0x13, 0x50, 0x72, 0x6f,
	0x6c, 0x6f, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2c, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x6c, 0x6f, 0x6e,
	0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x96, 0x03, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d,
	0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x58,
	0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x6f, 0x70,
	0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d,
	0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x58, 0x39, 0x32, 0x2e,
	0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x3d,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52,
	0x65, 0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67,
	0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x6c, 0x6f, 0x6e,
	0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65,
	0x58, 0x39, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x6f,
	0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x6c, 0x6f, 0x6e, 0x67,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a,
	0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authService_proto_rawDescOnce sync.Once
	file_authService_proto_rawDescData = file_authService_proto_rawDesc
)

func file_authService_proto_rawDescGZIP() []byte {
	file_authService_proto_rawDescOnce.Do(func() {
		file_authService_proto_rawDescData = protoimpl.X.CompressGZIP(file_authService_proto_rawDescData)
	})
	return file_authService_proto_rawDescData
}

var file_authService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_authService_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),      // 0: com.github.MaxReX92.go_yandex_gophkeeper.RegisterRequest
	(*RegisterResponse)(nil),     // 1: com.github.MaxReX92.go_yandex_gophkeeper.RegisterResponse
	(*LoginRequest)(nil),         // 2: com.github.MaxReX92.go_yandex_gophkeeper.LoginRequest
	(*LoginResponse)(nil),        // 3: com.github.MaxReX92.go_yandex_gophkeeper.LoginResponse
	(*ProlongTokenRequest)(nil),  // 4: com.github.MaxReX92.go_yandex_gophkeeper.ProlongTokenRequest
	(*ProlongTokenResponse)(nil), // 5: com.github.MaxReX92.go_yandex_gophkeeper.ProlongTokenResponse
}
var file_authService_proto_depIdxs = []int32{
	0, // 0: com.github.MaxReX92.go_yandex_gophkeeper.AuthService.Register:input_type -> com.github.MaxReX92.go_yandex_gophkeeper.RegisterRequest
	2, // 1: com.github.MaxReX92.go_yandex_gophkeeper.AuthService.Login:input_type -> com.github.MaxReX92.go_yandex_gophkeeper.LoginRequest
	4, // 2: com.github.MaxReX92.go_yandex_gophkeeper.AuthService.Prolong:input_type -> com.github.MaxReX92.go_yandex_gophkeeper.ProlongTokenRequest
	1, // 3: com.github.MaxReX92.go_yandex_gophkeeper.AuthService.Register:output_type -> com.github.MaxReX92.go_yandex_gophkeeper.RegisterResponse
	3, // 4: com.github.MaxReX92.go_yandex_gophkeeper.AuthService.Login:output_type -> com.github.MaxReX92.go_yandex_gophkeeper.LoginResponse
	5, // 5: com.github.MaxReX92.go_yandex_gophkeeper.AuthService.Prolong:output_type -> com.github.MaxReX92.go_yandex_gophkeeper.ProlongTokenResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authService_proto_init() }
func file_authService_proto_init() {
	if File_authService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_authService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_authService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_authService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_authService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProlongTokenRequest); i {
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
		file_authService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProlongTokenResponse); i {
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
			RawDescriptor: file_authService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authService_proto_goTypes,
		DependencyIndexes: file_authService_proto_depIdxs,
		MessageInfos:      file_authService_proto_msgTypes,
	}.Build()
	File_authService_proto = out.File
	file_authService_proto_rawDesc = nil
	file_authService_proto_goTypes = nil
	file_authService_proto_depIdxs = nil
}
