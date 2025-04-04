// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: controller/api/services/v1/session_service.proto

package services

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	sessions "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/sessions"
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

type GetSessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSessionRequest) Reset() {
	*x = GetSessionRequest{}
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSessionRequest) ProtoMessage() {}

func (x *GetSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSessionRequest.ProtoReflect.Descriptor instead.
func (*GetSessionRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_session_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSessionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSessionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          *sessions.Session      `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSessionResponse) Reset() {
	*x = GetSessionResponse{}
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSessionResponse) ProtoMessage() {}

func (x *GetSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSessionResponse.ProtoReflect.Descriptor instead.
func (*GetSessionResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_session_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetSessionResponse) GetItem() *sessions.Session {
	if x != nil {
		return x.Item
	}
	return nil
}

type ListSessionsRequest struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	ScopeId   string                 `protobuf:"bytes,1,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	Recursive bool                   `protobuf:"varint,20,opt,name=recursive,proto3" json:"recursive,omitempty" class:"public" eventstream:"observation"`          // @gotags: `class:"public" eventstream:"observation"`
	// You can specify that the filter should only return items that match.
	// Refer to [filter expressions](https://developer.hashicorp.com/boundary/docs/concepts/filtering) for more information.
	Filter string `protobuf:"bytes,30,opt,name=filter,proto3" json:"filter,omitempty" class:"public"` // @gotags: `class:"public"`
	// Deprecated. Should be removed in 0.18.0 since pagination removes the need
	// for the optimization that this provided.
	// By default only non-terminated (i.e. pending, active, canceling) are returned.
	// Set this option to include terminated sessions as well.  If not set paginated
	// responses may not include terminated sessions either in the response or in
	// the removed_ids.
	//
	// Deprecated: Marked as deprecated in controller/api/services/v1/session_service.proto.
	IncludeTerminated bool `protobuf:"varint,40,opt,name=include_terminated,proto3" json:"include_terminated,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// An opaque token that Boundary uses to continue an existing iteration or
	// request updated items. If you do not specify a token, pagination
	// starts from the beginning. To learn more about list pagination
	// in Boundary, refer to [list pagination](https://developer.hashicorp.com/boundary/docs/api-clients/api/pagination).
	ListToken string `protobuf:"bytes,50,opt,name=list_token,proto3" json:"list_token,omitempty" class:"public"` // @gotags: `class:"public"`
	// The maximum size of a page in this iteration.
	// If unset, the default page size configured will be used.
	// If the page_size is greater than the default page configured,
	// the page size will be truncated to this number..
	PageSize      uint32 `protobuf:"varint,60,opt,name=page_size,proto3" json:"page_size,omitempty" class:"public"` // @gotags: `class:"public"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSessionsRequest) Reset() {
	*x = ListSessionsRequest{}
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSessionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSessionsRequest) ProtoMessage() {}

func (x *ListSessionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSessionsRequest.ProtoReflect.Descriptor instead.
func (*ListSessionsRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_session_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListSessionsRequest) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *ListSessionsRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

func (x *ListSessionsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// Deprecated: Marked as deprecated in controller/api/services/v1/session_service.proto.
func (x *ListSessionsRequest) GetIncludeTerminated() bool {
	if x != nil {
		return x.IncludeTerminated
	}
	return false
}

func (x *ListSessionsRequest) GetListToken() string {
	if x != nil {
		return x.ListToken
	}
	return ""
}

func (x *ListSessionsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListSessionsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Items []*sessions.Session    `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// The type of response, either "delta" or "complete".
	// Delta signifies that this is part of a paginated result
	// or an update to a previously completed pagination.
	// Complete signifies that it is the last page.
	ResponseType string `protobuf:"bytes,2,opt,name=response_type,proto3" json:"response_type,omitempty" class:"public"` // @gotags: `class:"public"`
	// An opaque token used to continue an existing pagination or
	// request updated items. Use this token in the next list request
	// to request the next page.
	ListToken string `protobuf:"bytes,3,opt,name=list_token,proto3" json:"list_token,omitempty" class:"public"` // @gotags: `class:"public"`
	// The name of the field which the items are sorted by.
	SortBy string `protobuf:"bytes,4,opt,name=sort_by,proto3" json:"sort_by,omitempty" class:"public"` // @gotags: `class:"public"`
	// The direction of the sort, either "asc" or "desc".
	SortDir string `protobuf:"bytes,5,opt,name=sort_dir,proto3" json:"sort_dir,omitempty" class:"public"` // @gotags: `class:"public"`
	// A list of item IDs that have been removed since they were returned
	// as part of a pagination. They should be dropped from any client cache.
	// This may contain items that are not known to the cache, if they were
	// created and deleted between listings.
	RemovedIds []string `protobuf:"bytes,6,rep,name=removed_ids,proto3" json:"removed_ids,omitempty" class:"public"` // @gotags: `class:"public"`
	// An estimate at the total items available. This may change during pagination.
	EstItemCount  uint32 `protobuf:"varint,7,opt,name=est_item_count,proto3" json:"est_item_count,omitempty" class:"public"` // @gotags: `class:"public"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSessionsResponse) Reset() {
	*x = ListSessionsResponse{}
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSessionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSessionsResponse) ProtoMessage() {}

func (x *ListSessionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSessionsResponse.ProtoReflect.Descriptor instead.
func (*ListSessionsResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_session_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListSessionsResponse) GetItems() []*sessions.Session {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListSessionsResponse) GetResponseType() string {
	if x != nil {
		return x.ResponseType
	}
	return ""
}

func (x *ListSessionsResponse) GetListToken() string {
	if x != nil {
		return x.ListToken
	}
	return ""
}

func (x *ListSessionsResponse) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *ListSessionsResponse) GetSortDir() string {
	if x != nil {
		return x.SortDir
	}
	return ""
}

func (x *ListSessionsResponse) GetRemovedIds() []string {
	if x != nil {
		return x.RemovedIds
	}
	return nil
}

func (x *ListSessionsResponse) GetEstItemCount() uint32 {
	if x != nil {
		return x.EstItemCount
	}
	return 0
}

type CancelSessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" class:"public" eventstream:"observation"`            // @gotags: `class:"public" eventstream:"observation"`
	Version       uint32                 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty" class:"public"` // @gotags: `class:"public"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelSessionRequest) Reset() {
	*x = CancelSessionRequest{}
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelSessionRequest) ProtoMessage() {}

func (x *CancelSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelSessionRequest.ProtoReflect.Descriptor instead.
func (*CancelSessionRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_session_service_proto_rawDescGZIP(), []int{4}
}

func (x *CancelSessionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CancelSessionRequest) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type CancelSessionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          *sessions.Session      `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelSessionResponse) Reset() {
	*x = CancelSessionResponse{}
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelSessionResponse) ProtoMessage() {}

func (x *CancelSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_session_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelSessionResponse.ProtoReflect.Descriptor instead.
func (*CancelSessionResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_session_service_proto_rawDescGZIP(), []int{5}
}

func (x *CancelSessionResponse) GetItem() *sessions.Session {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_controller_api_services_v1_session_service_proto protoreflect.FileDescriptor

var file_controller_api_services_v1_session_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x32,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0xd8,
	0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xa1, 0x02, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x5f,
	0x64, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x5f,
	0x64, 0x69, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x5f, 0x69, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x40, 0x0a,
	0x14, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x5a, 0x0a, 0x15, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x32, 0xe8, 0x06, 0x0a, 0x0e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa7,
	0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x92, 0x41,
	0x18, 0x12, 0x16, 0x47, 0x65, 0x74, 0x73, 0x20, 0x61, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x20, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x62,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x9f, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x92, 0x41,
	0x15, 0x12, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xb6, 0x01, 0x0a, 0x0d, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x40, 0x92, 0x41, 0x14, 0x12, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x73, 0x20,
	0x61, 0x20, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23,
	0x3a, 0x01, 0x2a, 0x62, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x1a, 0xd0, 0x02, 0x92, 0x41, 0xcc, 0x02, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb5, 0x01, 0x41, 0x20,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x73, 0x20, 0x61, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x2c, 0x20,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x20, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x62, 0x65, 0x74,
	0x77, 0x65, 0x65, 0x6e, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x6e, 0x64, 0x20,
	0x61, 0x20, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x20, 0x54, 0x68, 0x65, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x6c,
	0x65, 0x74, 0x20, 0x79, 0x6f, 0x75, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x20, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x2e, 0x1a, 0x80, 0x01, 0x0a, 0x30, 0x52, 0x65, 0x61, 0x64, 0x20, 0x61, 0x62, 0x6f,
	0x75, 0x74, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x20, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x4c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x73,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_services_v1_session_service_proto_rawDescOnce sync.Once
	file_controller_api_services_v1_session_service_proto_rawDescData = file_controller_api_services_v1_session_service_proto_rawDesc
)

func file_controller_api_services_v1_session_service_proto_rawDescGZIP() []byte {
	file_controller_api_services_v1_session_service_proto_rawDescOnce.Do(func() {
		file_controller_api_services_v1_session_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_services_v1_session_service_proto_rawDescData)
	})
	return file_controller_api_services_v1_session_service_proto_rawDescData
}

var file_controller_api_services_v1_session_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_controller_api_services_v1_session_service_proto_goTypes = []any{
	(*GetSessionRequest)(nil),     // 0: controller.api.services.v1.GetSessionRequest
	(*GetSessionResponse)(nil),    // 1: controller.api.services.v1.GetSessionResponse
	(*ListSessionsRequest)(nil),   // 2: controller.api.services.v1.ListSessionsRequest
	(*ListSessionsResponse)(nil),  // 3: controller.api.services.v1.ListSessionsResponse
	(*CancelSessionRequest)(nil),  // 4: controller.api.services.v1.CancelSessionRequest
	(*CancelSessionResponse)(nil), // 5: controller.api.services.v1.CancelSessionResponse
	(*sessions.Session)(nil),      // 6: controller.api.resources.sessions.v1.Session
}
var file_controller_api_services_v1_session_service_proto_depIdxs = []int32{
	6, // 0: controller.api.services.v1.GetSessionResponse.item:type_name -> controller.api.resources.sessions.v1.Session
	6, // 1: controller.api.services.v1.ListSessionsResponse.items:type_name -> controller.api.resources.sessions.v1.Session
	6, // 2: controller.api.services.v1.CancelSessionResponse.item:type_name -> controller.api.resources.sessions.v1.Session
	0, // 3: controller.api.services.v1.SessionService.GetSession:input_type -> controller.api.services.v1.GetSessionRequest
	2, // 4: controller.api.services.v1.SessionService.ListSessions:input_type -> controller.api.services.v1.ListSessionsRequest
	4, // 5: controller.api.services.v1.SessionService.CancelSession:input_type -> controller.api.services.v1.CancelSessionRequest
	1, // 6: controller.api.services.v1.SessionService.GetSession:output_type -> controller.api.services.v1.GetSessionResponse
	3, // 7: controller.api.services.v1.SessionService.ListSessions:output_type -> controller.api.services.v1.ListSessionsResponse
	5, // 8: controller.api.services.v1.SessionService.CancelSession:output_type -> controller.api.services.v1.CancelSessionResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_controller_api_services_v1_session_service_proto_init() }
func file_controller_api_services_v1_session_service_proto_init() {
	if File_controller_api_services_v1_session_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_api_services_v1_session_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controller_api_services_v1_session_service_proto_goTypes,
		DependencyIndexes: file_controller_api_services_v1_session_service_proto_depIdxs,
		MessageInfos:      file_controller_api_services_v1_session_service_proto_msgTypes,
	}.Build()
	File_controller_api_services_v1_session_service_proto = out.File
	file_controller_api_services_v1_session_service_proto_rawDesc = nil
	file_controller_api_services_v1_session_service_proto_goTypes = nil
	file_controller_api_services_v1_session_service_proto_depIdxs = nil
}
