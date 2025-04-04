// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: controller/storage/auth/store/v1/auth_method.proto

package store

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

type AuthMethod struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: gorm:"primary_key"
	AuthAccountId string `protobuf:"bytes,1,opt,name=auth_account_id,json=authAccountId,proto3" json:"auth_account_id,omitempty"`
	// @inject_tag: `gorm:"default:null"`
	IamScopeId    string `protobuf:"bytes,2,opt,name=iam_scope_id,json=iamScopeId,proto3" json:"iam_scope_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthMethod) Reset() {
	*x = AuthMethod{}
	mi := &file_controller_storage_auth_store_v1_auth_method_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMethod) ProtoMessage() {}

func (x *AuthMethod) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_auth_store_v1_auth_method_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMethod.ProtoReflect.Descriptor instead.
func (*AuthMethod) Descriptor() ([]byte, []int) {
	return file_controller_storage_auth_store_v1_auth_method_proto_rawDescGZIP(), []int{0}
}

func (x *AuthMethod) GetAuthAccountId() string {
	if x != nil {
		return x.AuthAccountId
	}
	return ""
}

func (x *AuthMethod) GetIamScopeId() string {
	if x != nil {
		return x.IamScopeId
	}
	return ""
}

type AuthAccount struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: gorm:"primary_key"
	AuthAccountId string `protobuf:"bytes,1,opt,name=auth_account_id,json=authAccountId,proto3" json:"auth_account_id,omitempty"`
	// @inject_tag: `gorm:"default:null"`
	AuthMethodId string `protobuf:"bytes,2,opt,name=auth_method_id,json=authMethodId,proto3" json:"auth_method_id,omitempty"`
	// @inject_tag: `gorm:"default:null"`
	IamScopeId string `protobuf:"bytes,3,opt,name=iam_scope_id,json=iamScopeId,proto3" json:"iam_scope_id,omitempty"`
	// @inject_tag: `gorm:"default:null"`
	IamUserId     string `protobuf:"bytes,4,opt,name=iam_user_id,json=iamUserId,proto3" json:"iam_user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthAccount) Reset() {
	*x = AuthAccount{}
	mi := &file_controller_storage_auth_store_v1_auth_method_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthAccount) ProtoMessage() {}

func (x *AuthAccount) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_auth_store_v1_auth_method_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthAccount.ProtoReflect.Descriptor instead.
func (*AuthAccount) Descriptor() ([]byte, []int) {
	return file_controller_storage_auth_store_v1_auth_method_proto_rawDescGZIP(), []int{1}
}

func (x *AuthAccount) GetAuthAccountId() string {
	if x != nil {
		return x.AuthAccountId
	}
	return ""
}

func (x *AuthAccount) GetAuthMethodId() string {
	if x != nil {
		return x.AuthMethodId
	}
	return ""
}

func (x *AuthAccount) GetIamScopeId() string {
	if x != nil {
		return x.IamScopeId
	}
	return ""
}

func (x *AuthAccount) GetIamUserId() string {
	if x != nil {
		return x.IamUserId
	}
	return ""
}

var File_controller_storage_auth_store_v1_auth_method_proto protoreflect.FileDescriptor

var file_controller_storage_auth_store_v1_auth_method_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x56, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61,
	0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c,
	0x69, 0x61, 0x6d, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x61, 0x6d, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x22, 0x9d,
	0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c,
	0x69, 0x61, 0x6d, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x61, 0x6d, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0b, 0x69, 0x61, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x61, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x39,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_controller_storage_auth_store_v1_auth_method_proto_rawDescOnce sync.Once
	file_controller_storage_auth_store_v1_auth_method_proto_rawDescData = file_controller_storage_auth_store_v1_auth_method_proto_rawDesc
)

func file_controller_storage_auth_store_v1_auth_method_proto_rawDescGZIP() []byte {
	file_controller_storage_auth_store_v1_auth_method_proto_rawDescOnce.Do(func() {
		file_controller_storage_auth_store_v1_auth_method_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_auth_store_v1_auth_method_proto_rawDescData)
	})
	return file_controller_storage_auth_store_v1_auth_method_proto_rawDescData
}

var file_controller_storage_auth_store_v1_auth_method_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_storage_auth_store_v1_auth_method_proto_goTypes = []any{
	(*AuthMethod)(nil),  // 0: controller.storage.auth.store.v1.AuthMethod
	(*AuthAccount)(nil), // 1: controller.storage.auth.store.v1.AuthAccount
}
var file_controller_storage_auth_store_v1_auth_method_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_controller_storage_auth_store_v1_auth_method_proto_init() }
func file_controller_storage_auth_store_v1_auth_method_proto_init() {
	if File_controller_storage_auth_store_v1_auth_method_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_storage_auth_store_v1_auth_method_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_auth_store_v1_auth_method_proto_goTypes,
		DependencyIndexes: file_controller_storage_auth_store_v1_auth_method_proto_depIdxs,
		MessageInfos:      file_controller_storage_auth_store_v1_auth_method_proto_msgTypes,
	}.Build()
	File_controller_storage_auth_store_v1_auth_method_proto = out.File
	file_controller_storage_auth_store_v1_auth_method_proto_rawDesc = nil
	file_controller_storage_auth_store_v1_auth_method_proto_goTypes = nil
	file_controller_storage_auth_store_v1_auth_method_proto_depIdxs = nil
}
