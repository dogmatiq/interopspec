// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: github.com/dogmatiq/interopspec/envelopespec/envelope.proto

package envelopespec

import (
	proto "github.com/golang/protobuf/proto"
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

// Envelope is a container for a marshaled message and it meta-data.
type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MessageId is a unique identifier for the message.
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// CausationId is the ID of the message that was the direct cause of the
	// message in this envelope.
	CausationId string `protobuf:"bytes,2,opt,name=causation_id,json=causationId,proto3" json:"causation_id,omitempty"`
	// CorrelationId is the ID of the first ancestor of the message in this
	// envelope that was not caused by another message.
	CorrelationId string `protobuf:"bytes,3,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	// SourceApplication is the identity of the Dogma application that produced
	// the message in this envelope.
	SourceApplication *Identity `protobuf:"bytes,4,opt,name=source_application,json=sourceApplication,proto3" json:"source_application,omitempty"`
	// SourceHandler is the identity of the Dogma handler that produced the
	// message in this envelope. It is the zero-value if the message was not
	// produced by a handler.
	SourceHandler *Identity `protobuf:"bytes,5,opt,name=source_handler,json=sourceHandler,proto3" json:"source_handler,omitempty"`
	// SourceInstanceId is the ID of the aggregate or process instance that
	// produced the message in this envelope. It is empty if the message was not
	// produced by a handler, or it was produced by an integration handler.
	SourceInstanceId string `protobuf:"bytes,6,opt,name=source_instance_id,json=sourceInstanceId,proto3" json:"source_instance_id,omitempty"`
	// CreatedAt is the time at which the message was created, marshaled in
	// RFC-3339 format, with nanoseconds.
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ScheduledFor is the time at which a timeout message is scheduled to be
	// handled, marshaled in RFC-3339 format, with nanoseconds.
	ScheduledFor string `protobuf:"bytes,8,opt,name=scheduled_for,json=scheduledFor,proto3" json:"scheduled_for,omitempty"`
	// Description is a human-readable description of the message.
	Description string `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	// PortableName is the unique name used to identify messages of this type.
	PortableName string `protobuf:"bytes,10,opt,name=portable_name,json=portableName,proto3" json:"portable_name,omitempty"`
	// MediaType is a MIME media-type describing the content and encoding of the
	// binary message data.
	MediaType string `protobuf:"bytes,11,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	// Data is the binary message data.
	Data []byte `protobuf:"bytes,12,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Envelope) GetCausationId() string {
	if x != nil {
		return x.CausationId
	}
	return ""
}

func (x *Envelope) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *Envelope) GetSourceApplication() *Identity {
	if x != nil {
		return x.SourceApplication
	}
	return nil
}

func (x *Envelope) GetSourceHandler() *Identity {
	if x != nil {
		return x.SourceHandler
	}
	return nil
}

func (x *Envelope) GetSourceInstanceId() string {
	if x != nil {
		return x.SourceInstanceId
	}
	return ""
}

func (x *Envelope) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Envelope) GetScheduledFor() string {
	if x != nil {
		return x.ScheduledFor
	}
	return ""
}

func (x *Envelope) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Envelope) GetPortableName() string {
	if x != nil {
		return x.PortableName
	}
	return ""
}

func (x *Envelope) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *Envelope) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Identity represents the identity of an application or handler.
type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the entity's unique name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Key is the entity's immutable, unique key.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDescGZIP(), []int{1}
}

func (x *Identity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Identity) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_github_com_dogmatiq_interopspec_envelopespec_envelope_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65,
	0x63, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x65,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x64,
	0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0xff, 0x03, 0x0a, 0x08, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x75, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x75, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x52,
	0x0a, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x46, 0x6f, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x08, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x2e, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61,
	0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65, 0x63, 0x2f,
	0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDescData = file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDesc
)

func file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDescData)
	})
	return file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDescData
}

var file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_goTypes = []interface{}{
	(*Envelope)(nil), // 0: dogma.interop.v1.envelope.Envelope
	(*Identity)(nil), // 1: dogma.interop.v1.envelope.Identity
}
var file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_depIdxs = []int32{
	1, // 0: dogma.interop.v1.envelope.Envelope.source_application:type_name -> dogma.interop.v1.envelope.Identity
	1, // 1: dogma.interop.v1.envelope.Envelope.source_handler:type_name -> dogma.interop.v1.envelope.Identity
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_init() }
func file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_init() {
	if File_github_com_dogmatiq_interopspec_envelopespec_envelope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
		file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
			RawDescriptor: file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_depIdxs,
		MessageInfos:      file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_interopspec_envelopespec_envelope_proto = out.File
	file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_rawDesc = nil
	file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_goTypes = nil
	file_github_com_dogmatiq_interopspec_envelopespec_envelope_proto_depIdxs = nil
}
