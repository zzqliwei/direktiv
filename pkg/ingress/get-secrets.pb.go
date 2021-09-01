// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/ingress/get-secrets.proto

package ingress

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

type GetSecretsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
}

func (x *GetSecretsRequest) Reset() {
	*x = GetSecretsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_get_secrets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretsRequest) ProtoMessage() {}

func (x *GetSecretsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_get_secrets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretsRequest.ProtoReflect.Descriptor instead.
func (*GetSecretsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_get_secrets_proto_rawDescGZIP(), []int{0}
}

func (x *GetSecretsRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

type GetSecretsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secrets []*GetSecretsResponse_Secret `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
}

func (x *GetSecretsResponse) Reset() {
	*x = GetSecretsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_get_secrets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretsResponse) ProtoMessage() {}

func (x *GetSecretsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_get_secrets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretsResponse.ProtoReflect.Descriptor instead.
func (*GetSecretsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_get_secrets_proto_rawDescGZIP(), []int{1}
}

func (x *GetSecretsResponse) GetSecrets() []*GetSecretsResponse_Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

type GetSecretsResponse_Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *GetSecretsResponse_Secret) Reset() {
	*x = GetSecretsResponse_Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_get_secrets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretsResponse_Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretsResponse_Secret) ProtoMessage() {}

func (x *GetSecretsResponse_Secret) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_get_secrets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretsResponse_Secret.ProtoReflect.Descriptor instead.
func (*GetSecretsResponse_Secret) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_get_secrets_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetSecretsResponse_Secret) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var File_pkg_ingress_get_secrets_proto protoreflect.FileDescriptor

var file_pkg_ingress_get_secrets_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x7e,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x1a, 0x2a, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72,
	0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_ingress_get_secrets_proto_rawDescOnce sync.Once
	file_pkg_ingress_get_secrets_proto_rawDescData = file_pkg_ingress_get_secrets_proto_rawDesc
)

func file_pkg_ingress_get_secrets_proto_rawDescGZIP() []byte {
	file_pkg_ingress_get_secrets_proto_rawDescOnce.Do(func() {
		file_pkg_ingress_get_secrets_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_ingress_get_secrets_proto_rawDescData)
	})
	return file_pkg_ingress_get_secrets_proto_rawDescData
}

var file_pkg_ingress_get_secrets_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_ingress_get_secrets_proto_goTypes = []interface{}{
	(*GetSecretsRequest)(nil),         // 0: ingress.GetSecretsRequest
	(*GetSecretsResponse)(nil),        // 1: ingress.GetSecretsResponse
	(*GetSecretsResponse_Secret)(nil), // 2: ingress.GetSecretsResponse.Secret
}
var file_pkg_ingress_get_secrets_proto_depIdxs = []int32{
	2, // 0: ingress.GetSecretsResponse.secrets:type_name -> ingress.GetSecretsResponse.Secret
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_ingress_get_secrets_proto_init() }
func file_pkg_ingress_get_secrets_proto_init() {
	if File_pkg_ingress_get_secrets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_ingress_get_secrets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretsRequest); i {
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
		file_pkg_ingress_get_secrets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretsResponse); i {
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
		file_pkg_ingress_get_secrets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretsResponse_Secret); i {
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
	file_pkg_ingress_get_secrets_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_ingress_get_secrets_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_ingress_get_secrets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_ingress_get_secrets_proto_goTypes,
		DependencyIndexes: file_pkg_ingress_get_secrets_proto_depIdxs,
		MessageInfos:      file_pkg_ingress_get_secrets_proto_msgTypes,
	}.Build()
	File_pkg_ingress_get_secrets_proto = out.File
	file_pkg_ingress_get_secrets_proto_rawDesc = nil
	file_pkg_ingress_get_secrets_proto_goTypes = nil
	file_pkg_ingress_get_secrets_proto_depIdxs = nil
}
