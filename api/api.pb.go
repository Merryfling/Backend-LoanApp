// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_api_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50101,
		Name:          "api.form",
		Tag:           "bytes,50101,opt,name=form",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50102,
		Name:          "api.query",
		Tag:           "bytes,50102,opt,name=query",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50103,
		Name:          "api.header",
		Tag:           "bytes,50103,opt,name=header",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50104,
		Name:          "api.cookie",
		Tag:           "bytes,50104,opt,name=cookie",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50105,
		Name:          "api.body",
		Tag:           "bytes,50105,opt,name=body",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50106,
		Name:          "api.path",
		Tag:           "bytes,50106,opt,name=path",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50201,
		Name:          "api.get",
		Tag:           "bytes,50201,opt,name=get",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50202,
		Name:          "api.post",
		Tag:           "bytes,50202,opt,name=post",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50203,
		Name:          "api.put",
		Tag:           "bytes,50203,opt,name=put",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50204,
		Name:          "api.delete",
		Tag:           "bytes,50204,opt,name=delete",
		Filename:      "api.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string form = 50101;
	E_Form = &file_api_proto_extTypes[0]
	// optional string query = 50102;
	E_Query = &file_api_proto_extTypes[1]
	// optional string header = 50103;
	E_Header = &file_api_proto_extTypes[2]
	// optional string cookie = 50104;
	E_Cookie = &file_api_proto_extTypes[3]
	// optional string body = 50105;
	E_Body = &file_api_proto_extTypes[4]
	// optional string path = 50106;
	E_Path = &file_api_proto_extTypes[5]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional string get = 50201;
	E_Get = &file_api_proto_extTypes[6]
	// optional string post = 50202;
	E_Post = &file_api_proto_extTypes[7]
	// optional string put = 50203;
	E_Put = &file_api_proto_extTypes[8]
	// optional string delete = 50204;
	E_Delete = &file_api_proto_extTypes[9]
)

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3a, 0x36, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb5, 0x87, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x3a, 0x38, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xb6, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x88, 0x01, 0x01, 0x3a, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb7, 0x87,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x3a, 0x3a, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb8, 0x87, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x36, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xb9, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x88, 0x01, 0x01, 0x3a, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xba, 0x87, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x3a, 0x35, 0x0a, 0x03,
	0x67, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x99, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x65, 0x74,
	0x88, 0x01, 0x01, 0x3a, 0x37, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9a, 0x88, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x88, 0x01, 0x01, 0x3a, 0x35, 0x0a, 0x03,
	0x70, 0x75, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x9b, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x75, 0x74,
	0x88, 0x01, 0x01, 0x3a, 0x3b, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9c, 0x88,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_proto_goTypes = []any{
	(*descriptorpb.FieldOptions)(nil),  // 0: google.protobuf.FieldOptions
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: api.form:extendee -> google.protobuf.FieldOptions
	0,  // 1: api.query:extendee -> google.protobuf.FieldOptions
	0,  // 2: api.header:extendee -> google.protobuf.FieldOptions
	0,  // 3: api.cookie:extendee -> google.protobuf.FieldOptions
	0,  // 4: api.body:extendee -> google.protobuf.FieldOptions
	0,  // 5: api.path:extendee -> google.protobuf.FieldOptions
	1,  // 6: api.get:extendee -> google.protobuf.MethodOptions
	1,  // 7: api.post:extendee -> google.protobuf.MethodOptions
	1,  // 8: api.put:extendee -> google.protobuf.MethodOptions
	1,  // 9: api.delete:extendee -> google.protobuf.MethodOptions
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	0,  // [0:10] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 10,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		ExtensionInfos:    file_api_proto_extTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
