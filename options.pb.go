// Code generated by protoc-gen-go. DO NOT EDIT.
// source: x-mod/protobuf/options.proto

package options // import "github.com/x-mod/options"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServiceOption struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceOption) Reset()         { *m = ServiceOption{} }
func (m *ServiceOption) String() string { return proto.CompactTextString(m) }
func (*ServiceOption) ProtoMessage()    {}
func (*ServiceOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_options_def6cfb740cc3b48, []int{0}
}
func (m *ServiceOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceOption.Unmarshal(m, b)
}
func (m *ServiceOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceOption.Marshal(b, m, deterministic)
}
func (dst *ServiceOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceOption.Merge(dst, src)
}
func (m *ServiceOption) XXX_Size() int {
	return xxx_messageInfo_ServiceOption.Size(m)
}
func (m *ServiceOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceOption.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceOption proto.InternalMessageInfo

func (m *ServiceOption) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type HttpOption struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Uri                  string   `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpOption) Reset()         { *m = HttpOption{} }
func (m *HttpOption) String() string { return proto.CompactTextString(m) }
func (*HttpOption) ProtoMessage()    {}
func (*HttpOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_options_def6cfb740cc3b48, []int{1}
}
func (m *HttpOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpOption.Unmarshal(m, b)
}
func (m *HttpOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpOption.Marshal(b, m, deterministic)
}
func (dst *HttpOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpOption.Merge(dst, src)
}
func (m *HttpOption) XXX_Size() int {
	return xxx_messageInfo_HttpOption.Size(m)
}
func (m *HttpOption) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpOption.DiscardUnknown(m)
}

var xxx_messageInfo_HttpOption proto.InternalMessageInfo

func (m *HttpOption) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HttpOption) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

var E_Service = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*ServiceOption)(nil),
	Field:         91020818,
	Name:          "options.service",
	Tag:           "bytes,91020818,opt,name=service",
	Filename:      "x-mod/protobuf/options.proto",
}

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*HttpOption)(nil),
	Field:         91020819,
	Name:          "options.http",
	Tag:           "bytes,91020819,opt,name=http",
	Filename:      "x-mod/protobuf/options.proto",
}

func init() {
	proto.RegisterType((*ServiceOption)(nil), "options.ServiceOption")
	proto.RegisterType((*HttpOption)(nil), "options.HttpOption")
	proto.RegisterExtension(E_Service)
	proto.RegisterExtension(E_Http)
}

func init() {
	proto.RegisterFile("x-mod/protobuf/options.proto", fileDescriptor_options_def6cfb740cc3b48)
}

var fileDescriptor_options_def6cfb740cc3b48 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xa9, 0xd0, 0xcd, 0xcd,
	0x4f, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcf, 0x2f, 0x28, 0xc9, 0xcc,
	0xcf, 0x2b, 0xd6, 0x03, 0x0b, 0x08, 0xb1, 0x43, 0xb9, 0x52, 0x0a, 0xe9, 0xf9, 0xf9, 0xe9, 0x39,
	0xa9, 0x08, 0x75, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0xf9, 0x45, 0x10, 0xa5, 0x4a,
	0x9a, 0x5c, 0xbc, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0xfe, 0x60, 0x3d, 0x42, 0x12, 0x5c,
	0xec, 0x65, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30,
	0xae, 0x92, 0x19, 0x17, 0x97, 0x47, 0x49, 0x49, 0x01, 0x54, 0x9d, 0x18, 0x17, 0x5b, 0x6e, 0x6a,
	0x49, 0x46, 0x7e, 0x0a, 0x54, 0x19, 0x94, 0x27, 0x24, 0xc0, 0xc5, 0x5c, 0x5a, 0x94, 0x29, 0xc1,
	0x04, 0x16, 0x04, 0x31, 0xad, 0x42, 0xb8, 0xd8, 0x8b, 0x21, 0x56, 0x08, 0xc9, 0xeb, 0x41, 0x1c,
	0xa4, 0x07, 0x73, 0x90, 0x1e, 0x8a, 0xe5, 0xc5, 0x12, 0x93, 0xf6, 0x6c, 0xd6, 0x56, 0x60, 0xd4,
	0xe0, 0x36, 0x12, 0xd3, 0x83, 0x79, 0x09, 0x45, 0x45, 0x10, 0xcc, 0x28, 0x2b, 0x2f, 0x2e, 0x96,
	0x8c, 0x92, 0x92, 0x02, 0x21, 0x39, 0x0c, 0x23, 0x7d, 0xc1, 0x0e, 0x81, 0x99, 0x38, 0x19, 0x66,
	0xa2, 0x30, 0xdc, 0x44, 0x84, 0x2f, 0x82, 0xc0, 0x66, 0x38, 0x49, 0x45, 0x49, 0xa4, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x82, 0x16, 0xaa, 0x38, 0x89, 0x0d, 0x6c,
	0xae, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x06, 0xf5, 0x80, 0x74, 0x72, 0x01, 0x00, 0x00,
}
