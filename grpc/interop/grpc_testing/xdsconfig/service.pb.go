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
// source: protos/grpc/testing/xdsconfig/service.proto

package xdsconfig

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protos_grpc_testing_xdsconfig_service_proto protoreflect.FileDescriptor

var file_protos_grpc_testing_xdsconfig_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x78, 0x64, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x78, 0x64, 0x73, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x78, 0x64, 0x73, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xfb, 0x01, 0x0a, 0x17, 0x58, 0x64, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c,
	0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x70, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x78,
	0x64, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x4f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x78, 0x64, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x4f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x0f,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x78,
	0x64, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x78,
	0x64, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x78, 0x64,
	0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_grpc_testing_xdsconfig_service_proto_goTypes = []any{
	(*StopOnRequestRequest)(nil),    // 0: grpc.testing.xdsconfig.StopOnRequestRequest
	(*UpsertResourcesRequest)(nil),  // 1: grpc.testing.xdsconfig.UpsertResourcesRequest
	(*StopOnRequestResponse)(nil),   // 2: grpc.testing.xdsconfig.StopOnRequestResponse
	(*UpsertResourcesResponse)(nil), // 3: grpc.testing.xdsconfig.UpsertResourcesResponse
}
var file_protos_grpc_testing_xdsconfig_service_proto_depIdxs = []int32{
	0, // 0: grpc.testing.xdsconfig.XdsConfigControlService.StopOnRequest:input_type -> grpc.testing.xdsconfig.StopOnRequestRequest
	1, // 1: grpc.testing.xdsconfig.XdsConfigControlService.UpsertResources:input_type -> grpc.testing.xdsconfig.UpsertResourcesRequest
	2, // 2: grpc.testing.xdsconfig.XdsConfigControlService.StopOnRequest:output_type -> grpc.testing.xdsconfig.StopOnRequestResponse
	3, // 3: grpc.testing.xdsconfig.XdsConfigControlService.UpsertResources:output_type -> grpc.testing.xdsconfig.UpsertResourcesResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_grpc_testing_xdsconfig_service_proto_init() }
func file_protos_grpc_testing_xdsconfig_service_proto_init() {
	if File_protos_grpc_testing_xdsconfig_service_proto != nil {
		return
	}
	file_protos_grpc_testing_xdsconfig_control_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_grpc_testing_xdsconfig_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_grpc_testing_xdsconfig_service_proto_goTypes,
		DependencyIndexes: file_protos_grpc_testing_xdsconfig_service_proto_depIdxs,
	}.Build()
	File_protos_grpc_testing_xdsconfig_service_proto = out.File
	file_protos_grpc_testing_xdsconfig_service_proto_rawDesc = nil
	file_protos_grpc_testing_xdsconfig_service_proto_goTypes = nil
	file_protos_grpc_testing_xdsconfig_service_proto_depIdxs = nil
}
