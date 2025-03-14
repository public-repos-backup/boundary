// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: ops/services/v1/health_service.proto

package services

import (
	health "github.com/hashicorp/boundary/internal/gen/worker/health"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetHealthRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Experimental: This field may change or be removed without notice.
	WorkerInfo    bool `protobuf:"varint,1,opt,name=worker_info,proto3" json:"worker_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetHealthRequest) Reset() {
	*x = GetHealthRequest{}
	mi := &file_ops_services_v1_health_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthRequest) ProtoMessage() {}

func (x *GetHealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ops_services_v1_health_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthRequest.ProtoReflect.Descriptor instead.
func (*GetHealthRequest) Descriptor() ([]byte, []int) {
	return file_ops_services_v1_health_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetHealthRequest) GetWorkerInfo() bool {
	if x != nil {
		return x.WorkerInfo
	}
	return false
}

type GetHealthResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Experimental: This field may change or be removed without notice.
	WorkerProcessInfo *health.HealthInfo `protobuf:"bytes,1,opt,name=worker_process_info,proto3" json:"worker_process_info,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetHealthResponse) Reset() {
	*x = GetHealthResponse{}
	mi := &file_ops_services_v1_health_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthResponse) ProtoMessage() {}

func (x *GetHealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ops_services_v1_health_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthResponse.ProtoReflect.Descriptor instead.
func (*GetHealthResponse) Descriptor() ([]byte, []int) {
	return file_ops_services_v1_health_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetHealthResponse) GetWorkerProcessInfo() *health.HealthInfo {
	if x != nil {
		return x.WorkerProcessInfo
	}
	return nil
}

var File_ops_services_v1_health_service_proto protoreflect.FileDescriptor

var file_ops_services_v1_health_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6f, 0x70, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x63, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x74, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x21, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x42, 0x42, 0x5a,
	0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x70, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ops_services_v1_health_service_proto_rawDescOnce sync.Once
	file_ops_services_v1_health_service_proto_rawDescData = file_ops_services_v1_health_service_proto_rawDesc
)

func file_ops_services_v1_health_service_proto_rawDescGZIP() []byte {
	file_ops_services_v1_health_service_proto_rawDescOnce.Do(func() {
		file_ops_services_v1_health_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ops_services_v1_health_service_proto_rawDescData)
	})
	return file_ops_services_v1_health_service_proto_rawDescData
}

var file_ops_services_v1_health_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ops_services_v1_health_service_proto_goTypes = []any{
	(*GetHealthRequest)(nil),  // 0: ops.services.v1.GetHealthRequest
	(*GetHealthResponse)(nil), // 1: ops.services.v1.GetHealthResponse
	(*health.HealthInfo)(nil), // 2: worker.health.v1.HealthInfo
}
var file_ops_services_v1_health_service_proto_depIdxs = []int32{
	2, // 0: ops.services.v1.GetHealthResponse.worker_process_info:type_name -> worker.health.v1.HealthInfo
	0, // 1: ops.services.v1.HealthService.GetHealth:input_type -> ops.services.v1.GetHealthRequest
	1, // 2: ops.services.v1.HealthService.GetHealth:output_type -> ops.services.v1.GetHealthResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ops_services_v1_health_service_proto_init() }
func file_ops_services_v1_health_service_proto_init() {
	if File_ops_services_v1_health_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ops_services_v1_health_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ops_services_v1_health_service_proto_goTypes,
		DependencyIndexes: file_ops_services_v1_health_service_proto_depIdxs,
		MessageInfos:      file_ops_services_v1_health_service_proto_msgTypes,
	}.Build()
	File_ops_services_v1_health_service_proto = out.File
	file_ops_services_v1_health_service_proto_rawDesc = nil
	file_ops_services_v1_health_service_proto_goTypes = nil
	file_ops_services_v1_health_service_proto_depIdxs = nil
}
