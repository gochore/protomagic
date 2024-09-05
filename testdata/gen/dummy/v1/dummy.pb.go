// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: dummy/v1/dummy.proto

package dummyv1

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

type TestEnumType int32

const (
	TestEnumType_TEST_ENUM_TYPE_UNSPECIFIED TestEnumType = 0
	TestEnumType_TEST_ENUM_TYPE_FOO         TestEnumType = 1
	TestEnumType_TEST_ENUM_TYPE_BAR         TestEnumType = 2
	TestEnumType_TEST_ENUM_TYPE_HELLO_WORLD TestEnumType = 3   // test case: two words
	TestEnumType_TEST_ENUM_TYPE_FOO2        TestEnumType = 4   // test case: enum value with number
	TestEnumType_TEST_ENUM_TYPE_FOO_3       TestEnumType = 5   // test case: enum value with underscore number
	TestEnumType_TEST_ENUM_TYPE_HELLO       TestEnumType = 100 // test case: non-sequential enum value
)

// Enum value maps for TestEnumType.
var (
	TestEnumType_name = map[int32]string{
		0:   "TEST_ENUM_TYPE_UNSPECIFIED",
		1:   "TEST_ENUM_TYPE_FOO",
		2:   "TEST_ENUM_TYPE_BAR",
		3:   "TEST_ENUM_TYPE_HELLO_WORLD",
		4:   "TEST_ENUM_TYPE_FOO2",
		5:   "TEST_ENUM_TYPE_FOO_3",
		100: "TEST_ENUM_TYPE_HELLO",
	}
	TestEnumType_value = map[string]int32{
		"TEST_ENUM_TYPE_UNSPECIFIED": 0,
		"TEST_ENUM_TYPE_FOO":         1,
		"TEST_ENUM_TYPE_BAR":         2,
		"TEST_ENUM_TYPE_HELLO_WORLD": 3,
		"TEST_ENUM_TYPE_FOO2":        4,
		"TEST_ENUM_TYPE_FOO_3":       5,
		"TEST_ENUM_TYPE_HELLO":       100,
	}
)

func (x TestEnumType) Enum() *TestEnumType {
	p := new(TestEnumType)
	*p = x
	return p
}

func (x TestEnumType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestEnumType) Descriptor() protoreflect.EnumDescriptor {
	return file_dummy_v1_dummy_proto_enumTypes[0].Descriptor()
}

func (TestEnumType) Type() protoreflect.EnumType {
	return &file_dummy_v1_dummy_proto_enumTypes[0]
}

func (x TestEnumType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestEnumType.Descriptor instead.
func (TestEnumType) EnumDescriptor() ([]byte, []int) {
	return file_dummy_v1_dummy_proto_rawDescGZIP(), []int{0}
}

// test case: enum type with number
type TestEnum2Type int32

const (
	TestEnum2Type_TEST_ENUM_2_TYPE_UNSPECIFIED TestEnum2Type = 0
	TestEnum2Type_TEST_ENUM_2_TYPE_FOO         TestEnum2Type = 1
	TestEnum2Type_TEST_ENUM_2_TYPE_BAR         TestEnum2Type = 2
)

// Enum value maps for TestEnum2Type.
var (
	TestEnum2Type_name = map[int32]string{
		0: "TEST_ENUM_2_TYPE_UNSPECIFIED",
		1: "TEST_ENUM_2_TYPE_FOO",
		2: "TEST_ENUM_2_TYPE_BAR",
	}
	TestEnum2Type_value = map[string]int32{
		"TEST_ENUM_2_TYPE_UNSPECIFIED": 0,
		"TEST_ENUM_2_TYPE_FOO":         1,
		"TEST_ENUM_2_TYPE_BAR":         2,
	}
)

func (x TestEnum2Type) Enum() *TestEnum2Type {
	p := new(TestEnum2Type)
	*p = x
	return p
}

func (x TestEnum2Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestEnum2Type) Descriptor() protoreflect.EnumDescriptor {
	return file_dummy_v1_dummy_proto_enumTypes[1].Descriptor()
}

func (TestEnum2Type) Type() protoreflect.EnumType {
	return &file_dummy_v1_dummy_proto_enumTypes[1]
}

func (x TestEnum2Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestEnum2Type.Descriptor instead.
func (TestEnum2Type) EnumDescriptor() ([]byte, []int) {
	return file_dummy_v1_dummy_proto_rawDescGZIP(), []int{1}
}

var File_dummy_v1_dummy_proto protoreflect.FileDescriptor

var file_dummy_v1_dummy_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x76, 0x31,
	0x2a, 0xcb, 0x01, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x46, 0x4f, 0x4f, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x45, 0x53,
	0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41, 0x52, 0x10,
	0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x48, 0x45, 0x4c, 0x4c, 0x4f, 0x5f, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x10,
	0x03, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x46, 0x4f, 0x4f, 0x32, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x45,
	0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4f, 0x4f,
	0x5f, 0x33, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x45, 0x4c, 0x4c, 0x4f, 0x10, 0x64, 0x2a, 0x65,
	0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x32, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x1c, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x32, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x32,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4f, 0x4f, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54,
	0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x32, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x42, 0x41, 0x52, 0x10, 0x02, 0x42, 0x98, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x75,
	0x6d, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x61,
	0x67, 0x69, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x08, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14,
	0x44, 0x75, 0x6d, 0x6d, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dummy_v1_dummy_proto_rawDescOnce sync.Once
	file_dummy_v1_dummy_proto_rawDescData = file_dummy_v1_dummy_proto_rawDesc
)

func file_dummy_v1_dummy_proto_rawDescGZIP() []byte {
	file_dummy_v1_dummy_proto_rawDescOnce.Do(func() {
		file_dummy_v1_dummy_proto_rawDescData = protoimpl.X.CompressGZIP(file_dummy_v1_dummy_proto_rawDescData)
	})
	return file_dummy_v1_dummy_proto_rawDescData
}

var file_dummy_v1_dummy_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_dummy_v1_dummy_proto_goTypes = []any{
	(TestEnumType)(0),  // 0: dummy.v1.TestEnumType
	(TestEnum2Type)(0), // 1: dummy.v1.TestEnum2Type
}
var file_dummy_v1_dummy_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dummy_v1_dummy_proto_init() }
func file_dummy_v1_dummy_proto_init() {
	if File_dummy_v1_dummy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dummy_v1_dummy_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dummy_v1_dummy_proto_goTypes,
		DependencyIndexes: file_dummy_v1_dummy_proto_depIdxs,
		EnumInfos:         file_dummy_v1_dummy_proto_enumTypes,
	}.Build()
	File_dummy_v1_dummy_proto = out.File
	file_dummy_v1_dummy_proto_rawDesc = nil
	file_dummy_v1_dummy_proto_goTypes = nil
	file_dummy_v1_dummy_proto_depIdxs = nil
}
