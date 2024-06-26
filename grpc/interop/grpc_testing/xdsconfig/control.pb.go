// Copyright 2024 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: protos/grpc/testing/xdsconfig/control.proto

package xdsconfig

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StopOnRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceName string `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *StopOnRequestRequest) Reset() {
	*x = StopOnRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopOnRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopOnRequestRequest) ProtoMessage() {}

func (x *StopOnRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopOnRequestRequest.ProtoReflect.Descriptor instead.
func (*StopOnRequestRequest) Descriptor() ([]byte, []int) {
	return file_protos_grpc_testing_xdsconfig_control_proto_rawDescGZIP(), []int{0}
}

func (x *StopOnRequestRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *StopOnRequestRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

type StopOnRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters []*StopOnRequestResponse_ResourceFilter `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *StopOnRequestResponse) Reset() {
	*x = StopOnRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopOnRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopOnRequestResponse) ProtoMessage() {}

func (x *StopOnRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopOnRequestResponse.ProtoReflect.Descriptor instead.
func (*StopOnRequestResponse) Descriptor() ([]byte, []int) {
	return file_protos_grpc_testing_xdsconfig_control_proto_rawDescGZIP(), []int{1}
}

func (x *StopOnRequestResponse) GetFilters() []*StopOnRequestResponse_ResourceFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type UpsertResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listener     *string `protobuf:"bytes,1,opt,name=listener,proto3,oneof" json:"listener,omitempty"`
	Cluster      string  `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	UpstreamHost string  `protobuf:"bytes,3,opt,name=upstream_host,json=upstreamHost,proto3" json:"upstream_host,omitempty"`
	UpstreamPort uint32  `protobuf:"varint,4,opt,name=upstream_port,json=upstreamPort,proto3" json:"upstream_port,omitempty"`
}

func (x *UpsertResourcesRequest) Reset() {
	*x = UpsertResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertResourcesRequest) ProtoMessage() {}

func (x *UpsertResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertResourcesRequest.ProtoReflect.Descriptor instead.
func (*UpsertResourcesRequest) Descriptor() ([]byte, []int) {
	return file_protos_grpc_testing_xdsconfig_control_proto_rawDescGZIP(), []int{2}
}

func (x *UpsertResourcesRequest) GetListener() string {
	if x != nil && x.Listener != nil {
		return *x.Listener
	}
	return ""
}

func (x *UpsertResourcesRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *UpsertResourcesRequest) GetUpstreamHost() string {
	if x != nil {
		return x.UpstreamHost
	}
	return ""
}

func (x *UpsertResourcesRequest) GetUpstreamPort() uint32 {
	if x != nil {
		return x.UpstreamPort
	}
	return 0
}

type UpsertResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource []*anypb.Any `protobuf:"bytes,1,rep,name=resource,proto3" json:"resource,omitempty"`
}

func (x *UpsertResourcesResponse) Reset() {
	*x = UpsertResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertResourcesResponse) ProtoMessage() {}

func (x *UpsertResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertResourcesResponse.ProtoReflect.Descriptor instead.
func (*UpsertResourcesResponse) Descriptor() ([]byte, []int) {
	return file_protos_grpc_testing_xdsconfig_control_proto_rawDescGZIP(), []int{3}
}

func (x *UpsertResourcesResponse) GetResource() []*anypb.Any {
	if x != nil {
		return x.Resource
	}
	return nil
}

type StopOnRequestResponse_ResourceFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceName string `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *StopOnRequestResponse_ResourceFilter) Reset() {
	*x = StopOnRequestResponse_ResourceFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopOnRequestResponse_ResourceFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopOnRequestResponse_ResourceFilter) ProtoMessage() {}

func (x *StopOnRequestResponse_ResourceFilter) ProtoReflect() protoreflect.Message {
	mi := &file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopOnRequestResponse_ResourceFilter.ProtoReflect.Descriptor instead.
func (*StopOnRequestResponse_ResourceFilter) Descriptor() ([]byte, []int) {
	return file_protos_grpc_testing_xdsconfig_control_proto_rawDescGZIP(), []int{1, 0}
}

func (x *StopOnRequestResponse_ResourceFilter) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *StopOnRequestResponse_ResourceFilter) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_protos_grpc_testing_xdsconfig_control_proto protoreflect.FileDescriptor

var file_protos_grpc_testing_xdsconfig_control_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x78, 0x64, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x78, 0x64, 0x73, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x60, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x70, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x70, 0x4f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x78, 0x64, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x4f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x1a, 0x5a, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xaa, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x08, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x75,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x6f, 0x72, 0x74,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x22, 0x4b, 0x0a,
	0x17, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x78, 0x64, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_grpc_testing_xdsconfig_control_proto_rawDescOnce sync.Once
	file_protos_grpc_testing_xdsconfig_control_proto_rawDescData = file_protos_grpc_testing_xdsconfig_control_proto_rawDesc
)

func file_protos_grpc_testing_xdsconfig_control_proto_rawDescGZIP() []byte {
	file_protos_grpc_testing_xdsconfig_control_proto_rawDescOnce.Do(func() {
		file_protos_grpc_testing_xdsconfig_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_grpc_testing_xdsconfig_control_proto_rawDescData)
	})
	return file_protos_grpc_testing_xdsconfig_control_proto_rawDescData
}

var file_protos_grpc_testing_xdsconfig_control_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_grpc_testing_xdsconfig_control_proto_goTypes = []any{
	(*StopOnRequestRequest)(nil),                 // 0: grpc.testing.xdsconfig.StopOnRequestRequest
	(*StopOnRequestResponse)(nil),                // 1: grpc.testing.xdsconfig.StopOnRequestResponse
	(*UpsertResourcesRequest)(nil),               // 2: grpc.testing.xdsconfig.UpsertResourcesRequest
	(*UpsertResourcesResponse)(nil),              // 3: grpc.testing.xdsconfig.UpsertResourcesResponse
	(*StopOnRequestResponse_ResourceFilter)(nil), // 4: grpc.testing.xdsconfig.StopOnRequestResponse.ResourceFilter
	(*anypb.Any)(nil),                            // 5: google.protobuf.Any
}
var file_protos_grpc_testing_xdsconfig_control_proto_depIdxs = []int32{
	4, // 0: grpc.testing.xdsconfig.StopOnRequestResponse.filters:type_name -> grpc.testing.xdsconfig.StopOnRequestResponse.ResourceFilter
	5, // 1: grpc.testing.xdsconfig.UpsertResourcesResponse.resource:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_grpc_testing_xdsconfig_control_proto_init() }
func file_protos_grpc_testing_xdsconfig_control_proto_init() {
	if File_protos_grpc_testing_xdsconfig_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StopOnRequestRequest); i {
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
		file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StopOnRequestResponse); i {
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
		file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpsertResourcesRequest); i {
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
		file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpsertResourcesResponse); i {
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
		file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*StopOnRequestResponse_ResourceFilter); i {
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
	file_protos_grpc_testing_xdsconfig_control_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_grpc_testing_xdsconfig_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_grpc_testing_xdsconfig_control_proto_goTypes,
		DependencyIndexes: file_protos_grpc_testing_xdsconfig_control_proto_depIdxs,
		MessageInfos:      file_protos_grpc_testing_xdsconfig_control_proto_msgTypes,
	}.Build()
	File_protos_grpc_testing_xdsconfig_control_proto = out.File
	file_protos_grpc_testing_xdsconfig_control_proto_rawDesc = nil
	file_protos_grpc_testing_xdsconfig_control_proto_goTypes = nil
	file_protos_grpc_testing_xdsconfig_control_proto_depIdxs = nil
}
