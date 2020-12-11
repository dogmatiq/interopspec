// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: github.com/dogmatiq/interopspec/config/service.proto

package config

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListApplicationIdentitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListApplicationIdentitiesRequest) Reset() {
	*x = ListApplicationIdentitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationIdentitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationIdentitiesRequest) ProtoMessage() {}

func (x *ListApplicationIdentitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplicationIdentitiesRequest.ProtoReflect.Descriptor instead.
func (*ListApplicationIdentitiesRequest) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_config_service_proto_rawDescGZIP(), []int{0}
}

type ListApplicationIdentitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identities []*Identity `protobuf:"bytes,1,rep,name=identities,proto3" json:"identities,omitempty"`
}

func (x *ListApplicationIdentitiesResponse) Reset() {
	*x = ListApplicationIdentitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationIdentitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationIdentitiesResponse) ProtoMessage() {}

func (x *ListApplicationIdentitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplicationIdentitiesResponse.ProtoReflect.Descriptor instead.
func (*ListApplicationIdentitiesResponse) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListApplicationIdentitiesResponse) GetIdentities() []*Identity {
	if x != nil {
		return x.Identities
	}
	return nil
}

type ListApplicationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListApplicationsRequest) Reset() {
	*x = ListApplicationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationsRequest) ProtoMessage() {}

func (x *ListApplicationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplicationsRequest.ProtoReflect.Descriptor instead.
func (*ListApplicationsRequest) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_config_service_proto_rawDescGZIP(), []int{2}
}

type ListApplicationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications []*Application `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty"`
}

func (x *ListApplicationsResponse) Reset() {
	*x = ListApplicationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationsResponse) ProtoMessage() {}

func (x *ListApplicationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplicationsResponse.ProtoReflect.Descriptor instead.
func (*ListApplicationsResponse) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListApplicationsResponse) GetApplications() []*Application {
	if x != nil {
		return x.Applications
	}
	return nil
}

var File_github_com_dogmatiq_interopspec_config_service_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_interopspec_config_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65,
	0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d,
	0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65, 0x63,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x66, 0x0a, 0x21, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x19, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x64, 0x0a, 0x18, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64,
	0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x32, 0x96, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x92, 0x01, 0x0a, 0x19,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x39, 0x2e, 0x64, 0x6f, 0x67, 0x6d,
	0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x77, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x74, 0x69, 0x71,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dogmatiq_interopspec_config_service_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_interopspec_config_service_proto_rawDescData = file_github_com_dogmatiq_interopspec_config_service_proto_rawDesc
)

func file_github_com_dogmatiq_interopspec_config_service_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_interopspec_config_service_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_interopspec_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_interopspec_config_service_proto_rawDescData)
	})
	return file_github_com_dogmatiq_interopspec_config_service_proto_rawDescData
}

var file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_dogmatiq_interopspec_config_service_proto_goTypes = []interface{}{
	(*ListApplicationIdentitiesRequest)(nil),  // 0: dogma.interop.v1.config.ListApplicationIdentitiesRequest
	(*ListApplicationIdentitiesResponse)(nil), // 1: dogma.interop.v1.config.ListApplicationIdentitiesResponse
	(*ListApplicationsRequest)(nil),           // 2: dogma.interop.v1.config.ListApplicationsRequest
	(*ListApplicationsResponse)(nil),          // 3: dogma.interop.v1.config.ListApplicationsResponse
	(*Identity)(nil),                          // 4: dogma.interop.v1.config.Identity
	(*Application)(nil),                       // 5: dogma.interop.v1.config.Application
}
var file_github_com_dogmatiq_interopspec_config_service_proto_depIdxs = []int32{
	4, // 0: dogma.interop.v1.config.ListApplicationIdentitiesResponse.identities:type_name -> dogma.interop.v1.config.Identity
	5, // 1: dogma.interop.v1.config.ListApplicationsResponse.applications:type_name -> dogma.interop.v1.config.Application
	0, // 2: dogma.interop.v1.config.Config.ListApplicationIdentities:input_type -> dogma.interop.v1.config.ListApplicationIdentitiesRequest
	2, // 3: dogma.interop.v1.config.Config.ListApplications:input_type -> dogma.interop.v1.config.ListApplicationsRequest
	1, // 4: dogma.interop.v1.config.Config.ListApplicationIdentities:output_type -> dogma.interop.v1.config.ListApplicationIdentitiesResponse
	3, // 5: dogma.interop.v1.config.Config.ListApplications:output_type -> dogma.interop.v1.config.ListApplicationsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_interopspec_config_service_proto_init() }
func file_github_com_dogmatiq_interopspec_config_service_proto_init() {
	if File_github_com_dogmatiq_interopspec_config_service_proto != nil {
		return
	}
	file_github_com_dogmatiq_interopspec_config_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplicationIdentitiesRequest); i {
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
		file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplicationIdentitiesResponse); i {
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
		file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplicationsRequest); i {
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
		file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplicationsResponse); i {
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
			RawDescriptor: file_github_com_dogmatiq_interopspec_config_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_dogmatiq_interopspec_config_service_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_interopspec_config_service_proto_depIdxs,
		MessageInfos:      file_github_com_dogmatiq_interopspec_config_service_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_interopspec_config_service_proto = out.File
	file_github_com_dogmatiq_interopspec_config_service_proto_rawDesc = nil
	file_github_com_dogmatiq_interopspec_config_service_proto_goTypes = nil
	file_github_com_dogmatiq_interopspec_config_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConfigClient is the client API for Config service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigClient interface {
	// ListApplicationIdentities returns the identity of all applications.
	ListApplicationIdentities(ctx context.Context, in *ListApplicationIdentitiesRequest, opts ...grpc.CallOption) (*ListApplicationIdentitiesResponse, error)
	// ListApplications returns the full configuration of all applications.
	ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error)
}

type configClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigClient(cc grpc.ClientConnInterface) ConfigClient {
	return &configClient{cc}
}

func (c *configClient) ListApplicationIdentities(ctx context.Context, in *ListApplicationIdentitiesRequest, opts ...grpc.CallOption) (*ListApplicationIdentitiesResponse, error) {
	out := new(ListApplicationIdentitiesResponse)
	err := c.cc.Invoke(ctx, "/dogma.interop.v1.config.Config/ListApplicationIdentities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error) {
	out := new(ListApplicationsResponse)
	err := c.cc.Invoke(ctx, "/dogma.interop.v1.config.Config/ListApplications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServer is the server API for Config service.
type ConfigServer interface {
	// ListApplicationIdentities returns the identity of all applications.
	ListApplicationIdentities(context.Context, *ListApplicationIdentitiesRequest) (*ListApplicationIdentitiesResponse, error)
	// ListApplications returns the full configuration of all applications.
	ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error)
}

// UnimplementedConfigServer can be embedded to have forward compatible implementations.
type UnimplementedConfigServer struct {
}

func (*UnimplementedConfigServer) ListApplicationIdentities(context.Context, *ListApplicationIdentitiesRequest) (*ListApplicationIdentitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplicationIdentities not implemented")
}
func (*UnimplementedConfigServer) ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplications not implemented")
}

func RegisterConfigServer(s *grpc.Server, srv ConfigServer) {
	s.RegisterService(&_Config_serviceDesc, srv)
}

func _Config_ListApplicationIdentities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationIdentitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).ListApplicationIdentities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dogma.interop.v1.config.Config/ListApplicationIdentities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).ListApplicationIdentities(ctx, req.(*ListApplicationIdentitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dogma.interop.v1.config.Config/ListApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).ListApplications(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Config_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dogma.interop.v1.config.Config",
	HandlerType: (*ConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListApplicationIdentities",
			Handler:    _Config_ListApplicationIdentities_Handler,
		},
		{
			MethodName: "ListApplications",
			Handler:    _Config_ListApplications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/dogmatiq/interopspec/config/service.proto",
}
