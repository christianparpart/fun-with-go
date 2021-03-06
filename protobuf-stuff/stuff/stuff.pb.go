// Code generated by protoc-gen-go.
// source: stuff.proto
// DO NOT EDIT!

/*
Package stuff is a generated protocol buffer package.

It is generated from these files:
	stuff.proto

It has these top-level messages:
	Person
*/
package stuff

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Person struct {
	Id               *int64   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name             *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Email            []string `protobuf:"bytes,3,rep,name=email" json:"email,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetEmail() []string {
	if m != nil {
		return m.Email
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "Person")
}

var fileDescriptor0 = []byte{
	// 87 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2e, 0x29, 0x4d,
	0x4b, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe4, 0x62, 0x0b, 0x48, 0x2d, 0x2a, 0xce,
	0xcf, 0x13, 0xe2, 0xe2, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd2, 0x60, 0x16, 0xe2, 0xe1,
	0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0xf2, 0x38, 0x85, 0x78, 0xb9, 0x58, 0x53, 0x73,
	0x13, 0x33, 0x73, 0x24, 0x98, 0x15, 0x98, 0x35, 0x38, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61,
	0x7f, 0xb6, 0xb1, 0x40, 0x00, 0x00, 0x00,
}
