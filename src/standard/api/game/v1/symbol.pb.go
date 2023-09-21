// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: game/v1/symbol.proto

package gamev1

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

type Symbol int32

const (
	Symbol_SYMBOL_UNKNOWN_UNSPECIFIED Symbol = 0
	Symbol_SYMBOL_EMPTY               Symbol = 1
	Symbol_SYMBOL_CIRCLE              Symbol = 2
	Symbol_SYMBOL_CROSS               Symbol = 3
	Symbol_SYMBOL_NONE                Symbol = 4
)

// Enum value maps for Symbol.
var (
	Symbol_name = map[int32]string{
		0: "SYMBOL_UNKNOWN_UNSPECIFIED",
		1: "SYMBOL_EMPTY",
		2: "SYMBOL_CIRCLE",
		3: "SYMBOL_CROSS",
		4: "SYMBOL_NONE",
	}
	Symbol_value = map[string]int32{
		"SYMBOL_UNKNOWN_UNSPECIFIED": 0,
		"SYMBOL_EMPTY":               1,
		"SYMBOL_CIRCLE":              2,
		"SYMBOL_CROSS":               3,
		"SYMBOL_NONE":                4,
	}
)

func (x Symbol) Enum() *Symbol {
	p := new(Symbol)
	*p = x
	return p
}

func (x Symbol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Symbol) Descriptor() protoreflect.EnumDescriptor {
	return file_game_v1_symbol_proto_enumTypes[0].Descriptor()
}

func (Symbol) Type() protoreflect.EnumType {
	return &file_game_v1_symbol_proto_enumTypes[0]
}

func (x Symbol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Symbol.Descriptor instead.
func (Symbol) EnumDescriptor() ([]byte, []int) {
	return file_game_v1_symbol_proto_rawDescGZIP(), []int{0}
}

var File_game_v1_symbol_proto protoreflect.FileDescriptor

var file_game_v1_symbol_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2a,
	0x70, 0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x59, 0x4d,
	0x42, 0x4f, 0x4c, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x59, 0x4d,
	0x42, 0x4f, 0x4c, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53,
	0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x43, 0x49, 0x52, 0x43, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x43, 0x52, 0x4f, 0x53, 0x53, 0x10, 0x03,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x04, 0x42, 0xa1, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0b, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x77,
	0x61, 0x74, 0x61, 0x6e, 0x61, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2d, 0x74,
	0x69, 0x63, 0x2d, 0x74, 0x61, 0x63, 0x2d, 0x74, 0x6f, 0x65, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x67, 0x61, 0x6d, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58,
	0xaa, 0x02, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x47, 0x61, 0x6d,
	0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x47, 0x61, 0x6d,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_v1_symbol_proto_rawDescOnce sync.Once
	file_game_v1_symbol_proto_rawDescData = file_game_v1_symbol_proto_rawDesc
)

func file_game_v1_symbol_proto_rawDescGZIP() []byte {
	file_game_v1_symbol_proto_rawDescOnce.Do(func() {
		file_game_v1_symbol_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_v1_symbol_proto_rawDescData)
	})
	return file_game_v1_symbol_proto_rawDescData
}

var file_game_v1_symbol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_v1_symbol_proto_goTypes = []interface{}{
	(Symbol)(0), // 0: game.v1.Symbol
}
var file_game_v1_symbol_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_v1_symbol_proto_init() }
func file_game_v1_symbol_proto_init() {
	if File_game_v1_symbol_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_v1_symbol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_v1_symbol_proto_goTypes,
		DependencyIndexes: file_game_v1_symbol_proto_depIdxs,
		EnumInfos:         file_game_v1_symbol_proto_enumTypes,
	}.Build()
	File_game_v1_symbol_proto = out.File
	file_game_v1_symbol_proto_rawDesc = nil
	file_game_v1_symbol_proto_goTypes = nil
	file_game_v1_symbol_proto_depIdxs = nil
}
