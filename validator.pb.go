//copy from https://github.com/mwitkow/go-proto-validators/blob/master/validator.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: validator.proto

package validator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FieldValidator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Uses a Golang RE2-syntax regex to match the field contents.
	Regex *string `protobuf:"bytes,1,opt,name=regex" json:"regex,omitempty"`
	// Field value of integer strictly greater than this value.
	IntGt *int64 `protobuf:"varint,2,opt,name=int_gt,json=intGt" json:"int_gt,omitempty"`
	// Field value of integer strictly smaller than this value.
	IntLt *int64 `protobuf:"varint,3,opt,name=int_lt,json=intLt" json:"int_lt,omitempty"`
	// Field value of double strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatGt *float64 `protobuf:"fixed64,6,opt,name=float_gt,json=floatGt" json:"float_gt,omitempty"`
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatLt *float64 `protobuf:"fixed64,7,opt,name=float_lt,json=floatLt" json:"float_lt,omitempty"`
	// Field value of double describing the epsilon within which
	// any comparison should be considered to be true. For example,
	// when using float_gt = 0.35, using a float_epsilon of 0.05
	// would mean that any value above 0.30 is acceptable. It can be
	// thought of as a {float_value_condition} +- {float_epsilon}.
	// If unset, no correction for floating point inaccuracies in
	// comparisons will be attempted.
	FloatEpsilon *float64 `protobuf:"fixed64,8,opt,name=float_epsilon,json=floatEpsilon" json:"float_epsilon,omitempty"`
	// Floating-point value compared to which the field content should be greater or equal.
	FloatGte *float64 `protobuf:"fixed64,9,opt,name=float_gte,json=floatGte" json:"float_gte,omitempty"`
	// Floating-point value compared to which the field content should be smaller or equal.
	FloatLte *float64 `protobuf:"fixed64,10,opt,name=float_lte,json=floatLte" json:"float_lte,omitempty"`
	// Used for string fields, requires the string to be not empty (i.e different from "").
	StringNotEmpty *bool `protobuf:"varint,11,opt,name=string_not_empty,json=stringNotEmpty" json:"string_not_empty,omitempty"`
	// Repeated field with at least this number of elements.
	RepeatedCountMin *int64 `protobuf:"varint,12,opt,name=repeated_count_min,json=repeatedCountMin" json:"repeated_count_min,omitempty"`
	// Repeated field with at most this number of elements.
	RepeatedCountMax *int64 `protobuf:"varint,13,opt,name=repeated_count_max,json=repeatedCountMax" json:"repeated_count_max,omitempty"`
	// Field value of length greater than this value.
	LengthGt *int64 `protobuf:"varint,14,opt,name=length_gt,json=lengthGt" json:"length_gt,omitempty"`
	// Field value of length smaller than this value.
	LengthLt *int64 `protobuf:"varint,15,opt,name=length_lt,json=lengthLt" json:"length_lt,omitempty"`
	// Field value of length strictly equal to this value.
	LengthEq *int64 `protobuf:"varint,16,opt,name=length_eq,json=lengthEq" json:"length_eq,omitempty"`
	// Requires that the value is in the enum.
	IsInEnum *bool `protobuf:"varint,17,opt,name=is_in_enum,json=isInEnum" json:"is_in_enum,omitempty"`
}

func (x *FieldValidator) Reset() {
	*x = FieldValidator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldValidator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldValidator) ProtoMessage() {}

func (x *FieldValidator) ProtoReflect() protoreflect.Message {
	mi := &file_validator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldValidator.ProtoReflect.Descriptor instead.
func (*FieldValidator) Descriptor() ([]byte, []int) {
	return file_validator_proto_rawDescGZIP(), []int{0}
}

func (x *FieldValidator) GetRegex() string {
	if x != nil && x.Regex != nil {
		return *x.Regex
	}
	return ""
}

func (x *FieldValidator) GetIntGt() int64 {
	if x != nil && x.IntGt != nil {
		return *x.IntGt
	}
	return 0
}

func (x *FieldValidator) GetIntLt() int64 {
	if x != nil && x.IntLt != nil {
		return *x.IntLt
	}
	return 0
}

func (x *FieldValidator) GetFloatGt() float64 {
	if x != nil && x.FloatGt != nil {
		return *x.FloatGt
	}
	return 0
}

func (x *FieldValidator) GetFloatLt() float64 {
	if x != nil && x.FloatLt != nil {
		return *x.FloatLt
	}
	return 0
}

func (x *FieldValidator) GetFloatEpsilon() float64 {
	if x != nil && x.FloatEpsilon != nil {
		return *x.FloatEpsilon
	}
	return 0
}

func (x *FieldValidator) GetFloatGte() float64 {
	if x != nil && x.FloatGte != nil {
		return *x.FloatGte
	}
	return 0
}

func (x *FieldValidator) GetFloatLte() float64 {
	if x != nil && x.FloatLte != nil {
		return *x.FloatLte
	}
	return 0
}

func (x *FieldValidator) GetStringNotEmpty() bool {
	if x != nil && x.StringNotEmpty != nil {
		return *x.StringNotEmpty
	}
	return false
}

func (x *FieldValidator) GetRepeatedCountMin() int64 {
	if x != nil && x.RepeatedCountMin != nil {
		return *x.RepeatedCountMin
	}
	return 0
}

func (x *FieldValidator) GetRepeatedCountMax() int64 {
	if x != nil && x.RepeatedCountMax != nil {
		return *x.RepeatedCountMax
	}
	return 0
}

func (x *FieldValidator) GetLengthGt() int64 {
	if x != nil && x.LengthGt != nil {
		return *x.LengthGt
	}
	return 0
}

func (x *FieldValidator) GetLengthLt() int64 {
	if x != nil && x.LengthLt != nil {
		return *x.LengthLt
	}
	return 0
}

func (x *FieldValidator) GetLengthEq() int64 {
	if x != nil && x.LengthEq != nil {
		return *x.LengthEq
	}
	return 0
}

func (x *FieldValidator) GetIsInEnum() bool {
	if x != nil && x.IsInEnum != nil {
		return *x.IsInEnum
	}
	return false
}

var file_validator_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldValidator)(nil),
		Field:         65020,
		Name:          "validator.field",
		Tag:           "bytes,65020,opt,name=field",
		Filename:      "validator.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional validator.FieldValidator field = 65020;
	E_Field = &file_validator_proto_extTypes[0]
)

var File_validator_proto protoreflect.FileDescriptor

var file_validator_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4,
	0x03, 0x0a, 0x0e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x5f, 0x67,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x47, 0x74, 0x12, 0x15,
	0x0a, 0x06, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x69, 0x6e, 0x74, 0x4c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x67,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x47, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x65, 0x70, 0x73, 0x69, 0x6c, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x45, 0x70, 0x73, 0x69, 0x6c, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x67, 0x74, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x47, 0x74, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x6c, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d,
	0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x78,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x67, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x47, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x6c, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x4c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x5f, 0x65, 0x71, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x45, 0x71, 0x12, 0x1c, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x69, 0x6e,
	0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x49,
	0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x3a, 0x50, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfc, 0xfb,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
}

var (
	file_validator_proto_rawDescOnce sync.Once
	file_validator_proto_rawDescData = file_validator_proto_rawDesc
)

func file_validator_proto_rawDescGZIP() []byte {
	file_validator_proto_rawDescOnce.Do(func() {
		file_validator_proto_rawDescData = protoimpl.X.CompressGZIP(file_validator_proto_rawDescData)
	})
	return file_validator_proto_rawDescData
}

var file_validator_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_validator_proto_goTypes = []interface{}{
	(*FieldValidator)(nil),            // 0: validator.FieldValidator
	(*descriptorpb.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_validator_proto_depIdxs = []int32{
	1, // 0: validator.field:extendee -> google.protobuf.FieldOptions
	0, // 1: validator.field:type_name -> validator.FieldValidator
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_validator_proto_init() }
func file_validator_proto_init() {
	if File_validator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_validator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldValidator); i {
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
			RawDescriptor: file_validator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_validator_proto_goTypes,
		DependencyIndexes: file_validator_proto_depIdxs,
		MessageInfos:      file_validator_proto_msgTypes,
		ExtensionInfos:    file_validator_proto_extTypes,
	}.Build()
	File_validator_proto = out.File
	file_validator_proto_rawDesc = nil
	file_validator_proto_goTypes = nil
	file_validator_proto_depIdxs = nil
}