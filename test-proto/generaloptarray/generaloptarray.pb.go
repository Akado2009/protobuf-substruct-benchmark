// Code generated by protoc-gen-go. DO NOT EDIT.
// source: generaloptarray/generaloptarray.proto

package generaloptarray // import "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/generaloptarray"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import firstmessage "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/firstmessage"
import secondmessage "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/secondmessage"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GeneralOptArrays struct {
	Fmsgs                []*firstmessage.FirstMessage   `protobuf:"bytes,1,rep,name=fmsgs,proto3" json:"fmsgs,omitempty"`
	Smsgs                []*secondmessage.SecondMessage `protobuf:"bytes,2,rep,name=smsgs,proto3" json:"smsgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GeneralOptArrays) Reset()         { *m = GeneralOptArrays{} }
func (m *GeneralOptArrays) String() string { return proto.CompactTextString(m) }
func (*GeneralOptArrays) ProtoMessage()    {}
func (*GeneralOptArrays) Descriptor() ([]byte, []int) {
	return fileDescriptor_generaloptarray_e97c7444cf1a322d, []int{0}
}
func (m *GeneralOptArrays) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralOptArrays.Unmarshal(m, b)
}
func (m *GeneralOptArrays) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralOptArrays.Marshal(b, m, deterministic)
}
func (dst *GeneralOptArrays) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralOptArrays.Merge(dst, src)
}
func (m *GeneralOptArrays) XXX_Size() int {
	return xxx_messageInfo_GeneralOptArrays.Size(m)
}
func (m *GeneralOptArrays) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralOptArrays.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralOptArrays proto.InternalMessageInfo

func NewGeneralOptArrays() *GeneralOptArrays {
	m := &GeneralOptArrays{}
	return m
}

func (m *GeneralOptArrays) GetFmsgs() []*firstmessage.FirstMessage {
	return m.Fmsgs
}

func (m *GeneralOptArrays) GetSmsgs() []*secondmessage.SecondMessage {
	return m.Smsgs
}

func (m *GeneralOptArrays) SetFmsgs(value []*firstmessage.FirstMessage) *GeneralOptArrays {
	m.Fmsgs = value
	return m
}

func (m *GeneralOptArrays) SetSmsgs(value []*secondmessage.SecondMessage) *GeneralOptArrays {
	m.Smsgs = value
	return m
}

func init() {
	proto.RegisterType((*GeneralOptArrays)(nil), "general.GeneralOptArrays")
}

func init() {
	proto.RegisterFile("generaloptarray/generaloptarray.proto", fileDescriptor_generaloptarray_e97c7444cf1a322d)
}

var fileDescriptor_generaloptarray_e97c7444cf1a322d = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x4f, 0xcd, 0x4b,
	0x2d, 0x4a, 0xcc, 0xc9, 0x2f, 0x28, 0x49, 0x2c, 0x2a, 0x4a, 0xac, 0xd4, 0x47, 0xe3, 0xeb, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xb1, 0x43, 0x85, 0xa5, 0xe4, 0xd3, 0x32, 0x8b, 0x8a, 0x4b, 0x72,
	0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0xf5, 0x91, 0x39, 0x10, 0x95, 0x52, 0x8a, 0xc5, 0xa9, 0xc9,
	0xf9, 0x79, 0x29, 0x30, 0x15, 0x28, 0x3c, 0x88, 0x12, 0xa5, 0x0a, 0x2e, 0x01, 0x77, 0x88, 0x71,
	0xfe, 0x05, 0x25, 0x8e, 0x20, 0x5b, 0x8a, 0x85, 0x0c, 0xb8, 0x58, 0xd3, 0x72, 0x8b, 0xd3, 0x8b,
	0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xa4, 0xf4, 0x50, 0x8c, 0x76, 0x03, 0x71, 0x7c, 0x21,
	0x9c, 0x20, 0x88, 0x42, 0x21, 0x23, 0x2e, 0xd6, 0x62, 0xb0, 0x0e, 0x26, 0xb0, 0x0e, 0x19, 0x3d,
	0x54, 0xab, 0x82, 0xc1, 0x3c, 0xb8, 0x1e, 0xb0, 0x52, 0x27, 0xbf, 0x28, 0x9f, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc7, 0xec, 0xc4, 0x94, 0x7c, 0x23, 0x03, 0x03,
	0x4b, 0x7d, 0xb0, 0xbb, 0x92, 0x4a, 0xd3, 0x74, 0x8b, 0x4b, 0x93, 0x8a, 0x4b, 0x8a, 0x4a, 0x93,
	0x4b, 0x74, 0x93, 0x52, 0xf3, 0x92, 0x33, 0x72, 0x13, 0x8b, 0xb2, 0xf5, 0x4b, 0x52, 0x8b, 0x4b,
	0x74, 0xc1, 0x2a, 0xd0, 0x03, 0x27, 0x89, 0x0d, 0x2c, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x6b, 0xd4, 0xe4, 0xa2, 0x46, 0x01, 0x00, 0x00,
}
