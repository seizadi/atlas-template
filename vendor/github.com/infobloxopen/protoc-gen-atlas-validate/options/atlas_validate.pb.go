// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto

/*
Package options is a generated protocol buffer package.

It is generated from these files:
	github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto

It has these top-level messages:
	AtlasValidateFileOption
	AtlasValidateMethodOption
	AtlasValidateServiceOption
*/
package options

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AtlasValidateFileOption struct {
	AllowUnknownFields bool `protobuf:"varint,1,opt,name=allow_unknown_fields,json=allowUnknownFields,proto3" json:"allow_unknown_fields,omitempty"`
}

func (m *AtlasValidateFileOption) Reset()         { *m = AtlasValidateFileOption{} }
func (m *AtlasValidateFileOption) String() string { return proto.CompactTextString(m) }
func (*AtlasValidateFileOption) ProtoMessage()    {}
func (*AtlasValidateFileOption) Descriptor() ([]byte, []int) {
	return fileDescriptorAtlasValidate, []int{0}
}

func (m *AtlasValidateFileOption) GetAllowUnknownFields() bool {
	if m != nil {
		return m.AllowUnknownFields
	}
	return false
}

type AtlasValidateMethodOption struct {
	AllowUnknownFields bool `protobuf:"varint,1,opt,name=allow_unknown_fields,json=allowUnknownFields,proto3" json:"allow_unknown_fields,omitempty"`
}

func (m *AtlasValidateMethodOption) Reset()         { *m = AtlasValidateMethodOption{} }
func (m *AtlasValidateMethodOption) String() string { return proto.CompactTextString(m) }
func (*AtlasValidateMethodOption) ProtoMessage()    {}
func (*AtlasValidateMethodOption) Descriptor() ([]byte, []int) {
	return fileDescriptorAtlasValidate, []int{1}
}

func (m *AtlasValidateMethodOption) GetAllowUnknownFields() bool {
	if m != nil {
		return m.AllowUnknownFields
	}
	return false
}

type AtlasValidateServiceOption struct {
	AllowUnknownFields bool `protobuf:"varint,1,opt,name=allow_unknown_fields,json=allowUnknownFields,proto3" json:"allow_unknown_fields,omitempty"`
}

func (m *AtlasValidateServiceOption) Reset()         { *m = AtlasValidateServiceOption{} }
func (m *AtlasValidateServiceOption) String() string { return proto.CompactTextString(m) }
func (*AtlasValidateServiceOption) ProtoMessage()    {}
func (*AtlasValidateServiceOption) Descriptor() ([]byte, []int) {
	return fileDescriptorAtlasValidate, []int{2}
}

func (m *AtlasValidateServiceOption) GetAllowUnknownFields() bool {
	if m != nil {
		return m.AllowUnknownFields
	}
	return false
}

var E_File = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*AtlasValidateFileOption)(nil),
	Field:         52219,
	Name:          "atlas_validate.file",
	Tag:           "bytes,52219,opt,name=file",
	Filename:      "github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto",
}

var E_Method = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*AtlasValidateMethodOption)(nil),
	Field:         52219,
	Name:          "atlas_validate.method",
	Tag:           "bytes,52219,opt,name=method",
	Filename:      "github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto",
}

var E_Service = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*AtlasValidateServiceOption)(nil),
	Field:         52219,
	Name:          "atlas_validate.service",
	Tag:           "bytes,52219,opt,name=service",
	Filename:      "github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto",
}

func init() {
	proto.RegisterType((*AtlasValidateFileOption)(nil), "atlas_validate.AtlasValidateFileOption")
	proto.RegisterType((*AtlasValidateMethodOption)(nil), "atlas_validate.AtlasValidateMethodOption")
	proto.RegisterType((*AtlasValidateServiceOption)(nil), "atlas_validate.AtlasValidateServiceOption")
	proto.RegisterExtension(E_File)
	proto.RegisterExtension(E_Method)
	proto.RegisterExtension(E_Service)
}

func init() {
	proto.RegisterFile("github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto", fileDescriptorAtlasValidate)
}

var fileDescriptorAtlasValidate = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x29, 0xca, 0x94, 0x08, 0x1e, 0x82, 0xe0, 0x1c, 0xa2, 0x65, 0x17, 0xa7, 0xd0, 0x54,
	0xf4, 0x56, 0x4f, 0x53, 0xd8, 0x45, 0x36, 0xa1, 0xa2, 0x07, 0x3d, 0x94, 0xfe, 0x48, 0xbb, 0x60,
	0x96, 0x57, 0x9a, 0x74, 0xf3, 0x3f, 0xf1, 0x8f, 0xf5, 0x22, 0xa6, 0x2d, 0x5d, 0x1c, 0xee, 0xd0,
	0x53, 0xe1, 0xbd, 0x7e, 0x3f, 0xdf, 0xbe, 0x0f, 0x45, 0xb3, 0x8c, 0xa9, 0x79, 0x19, 0x91, 0x18,
	0x16, 0x2e, 0x13, 0x29, 0x44, 0x1c, 0x3e, 0x21, 0xa7, 0xc2, 0xcd, 0x0b, 0x50, 0x10, 0x3b, 0x19,
	0x15, 0x4e, 0xa8, 0x78, 0x28, 0x9d, 0x65, 0xc8, 0x59, 0x12, 0x2a, 0xea, 0x42, 0xae, 0x18, 0x08,
	0xe9, 0xea, 0x71, 0xd0, 0x8c, 0x89, 0x0e, 0xe0, 0x43, 0x73, 0x3a, 0xb0, 0x33, 0x80, 0x8c, 0xd3,
	0x0a, 0x17, 0x95, 0xa9, 0x9b, 0x50, 0x19, 0x17, 0x2c, 0x57, 0x50, 0x54, 0x89, 0xe1, 0x23, 0x3a,
	0x1e, 0xff, 0x66, 0x5e, 0xeb, 0xc8, 0x84, 0x71, 0xfa, 0xa4, 0x2b, 0xf0, 0x35, 0x3a, 0x0a, 0x39,
	0x87, 0x55, 0x50, 0x8a, 0x0f, 0x01, 0x2b, 0x11, 0xa4, 0x8c, 0xf2, 0x44, 0xf6, 0x2d, 0xdb, 0x1a,
	0xed, 0xfb, 0x58, 0xef, 0x5e, 0xaa, 0xd5, 0x44, 0x6f, 0x86, 0x53, 0x74, 0x62, 0xc0, 0xa6, 0x54,
	0xcd, 0x21, 0xe9, 0x8c, 0x9b, 0xa1, 0x81, 0x81, 0x7b, 0xa6, 0xc5, 0x92, 0xc5, 0x9d, 0x3f, 0xcf,
	0x7b, 0x47, 0xbb, 0x29, 0xe3, 0x14, 0x9f, 0x92, 0x4a, 0x0b, 0x69, 0xb4, 0x90, 0xf6, 0x6a, 0xd9,
	0xff, 0xfe, 0xda, 0xb1, 0xad, 0xd1, 0xc1, 0xcd, 0x05, 0xf9, 0xa3, 0xf8, 0x1f, 0x4f, 0xbe, 0x86,
	0x7a, 0x31, 0xea, 0x2d, 0xf4, 0xb9, 0xf8, 0x6c, 0x03, 0xbf, 0xee, 0xa1, 0x2d, 0xb8, 0xdc, 0x5a,
	0xb0, 0x9e, 0xf1, 0x6b, 0xb4, 0x97, 0xa1, 0x3d, 0x59, 0x49, 0xc0, 0xe7, 0x1b, 0x2d, 0x86, 0x9e,
	0xb6, 0xe6, 0x6a, 0x6b, 0x8d, 0x11, 0xf2, 0x1b, 0xfa, 0xfd, 0xc3, 0xdb, 0xb8, 0xf3, 0xaf, 0x79,
	0x57, 0x3f, 0xa3, 0x9e, 0x7e, 0xf5, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x22, 0x84, 0xde, 0xbf,
	0xe6, 0x02, 0x00, 0x00,
}
