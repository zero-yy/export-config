// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conf.proto

package conf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Test struct {
	Records              map[int32]*Test_Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Crc32                uint32                 `protobuf:"varint,2,opt,name=crc32,proto3" json:"crc32,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6ecbfc68e85c65, []int{0}
}

func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetRecords() map[int32]*Test_Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *Test) GetCrc32() uint32 {
	if m != nil {
		return m.Crc32
	}
	return 0
}

type Test_Record struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TestString           string   `protobuf:"bytes,2,opt,name=test_string,json=testString,proto3" json:"test_string,omitempty"`
	TestDefault          string   `protobuf:"bytes,3,opt,name=test_default,json=testDefault,proto3" json:"test_default,omitempty"`
	TestInt32            int32    `protobuf:"varint,4,opt,name=test_int32,json=testInt32,proto3" json:"test_int32,omitempty"`
	TestUint32           uint32   `protobuf:"varint,5,opt,name=test_uint32,json=testUint32,proto3" json:"test_uint32,omitempty"`
	TestInt64            int64    `protobuf:"varint,6,opt,name=test_int64,json=testInt64,proto3" json:"test_int64,omitempty"`
	TestUint64           uint64   `protobuf:"varint,7,opt,name=test_uint64,json=testUint64,proto3" json:"test_uint64,omitempty"`
	TestFloat32          float32  `protobuf:"fixed32,8,opt,name=test_float32,json=testFloat32,proto3" json:"test_float32,omitempty"`
	TestFloat64          float64  `protobuf:"fixed64,9,opt,name=test_float64,json=testFloat64,proto3" json:"test_float64,omitempty"`
	TestDouble           float64  `protobuf:"fixed64,10,opt,name=test_double,json=testDouble,proto3" json:"test_double,omitempty"`
	TestArys             []int32  `protobuf:"varint,11,rep,packed,name=test_arys,json=testArys,proto3" json:"test_arys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test_Record) Reset()         { *m = Test_Record{} }
func (m *Test_Record) String() string { return proto.CompactTextString(m) }
func (*Test_Record) ProtoMessage()    {}
func (*Test_Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6ecbfc68e85c65, []int{0, 0}
}

func (m *Test_Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test_Record.Unmarshal(m, b)
}
func (m *Test_Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test_Record.Marshal(b, m, deterministic)
}
func (m *Test_Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test_Record.Merge(m, src)
}
func (m *Test_Record) XXX_Size() int {
	return xxx_messageInfo_Test_Record.Size(m)
}
func (m *Test_Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Test_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Test_Record proto.InternalMessageInfo

func (m *Test_Record) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Test_Record) GetTestString() string {
	if m != nil {
		return m.TestString
	}
	return ""
}

func (m *Test_Record) GetTestDefault() string {
	if m != nil {
		return m.TestDefault
	}
	return ""
}

func (m *Test_Record) GetTestInt32() int32 {
	if m != nil {
		return m.TestInt32
	}
	return 0
}

func (m *Test_Record) GetTestUint32() uint32 {
	if m != nil {
		return m.TestUint32
	}
	return 0
}

func (m *Test_Record) GetTestInt64() int64 {
	if m != nil {
		return m.TestInt64
	}
	return 0
}

func (m *Test_Record) GetTestUint64() uint64 {
	if m != nil {
		return m.TestUint64
	}
	return 0
}

func (m *Test_Record) GetTestFloat32() float32 {
	if m != nil {
		return m.TestFloat32
	}
	return 0
}

func (m *Test_Record) GetTestFloat64() float64 {
	if m != nil {
		return m.TestFloat64
	}
	return 0
}

func (m *Test_Record) GetTestDouble() float64 {
	if m != nil {
		return m.TestDouble
	}
	return 0
}

func (m *Test_Record) GetTestArys() []int32 {
	if m != nil {
		return m.TestArys
	}
	return nil
}

type Test2 struct {
	Records              map[string]*Test2_Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Crc32                uint32                   `protobuf:"varint,2,opt,name=crc32,proto3" json:"crc32,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Test2) Reset()         { *m = Test2{} }
func (m *Test2) String() string { return proto.CompactTextString(m) }
func (*Test2) ProtoMessage()    {}
func (*Test2) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6ecbfc68e85c65, []int{1}
}

func (m *Test2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test2.Unmarshal(m, b)
}
func (m *Test2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test2.Marshal(b, m, deterministic)
}
func (m *Test2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test2.Merge(m, src)
}
func (m *Test2) XXX_Size() int {
	return xxx_messageInfo_Test2.Size(m)
}
func (m *Test2) XXX_DiscardUnknown() {
	xxx_messageInfo_Test2.DiscardUnknown(m)
}

var xxx_messageInfo_Test2 proto.InternalMessageInfo

func (m *Test2) GetRecords() map[string]*Test2_Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *Test2) GetCrc32() uint32 {
	if m != nil {
		return m.Crc32
	}
	return 0
}

type Test2_Record struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TestArys             []int32  `protobuf:"varint,2,rep,packed,name=test_arys,json=testArys,proto3" json:"test_arys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test2_Record) Reset()         { *m = Test2_Record{} }
func (m *Test2_Record) String() string { return proto.CompactTextString(m) }
func (*Test2_Record) ProtoMessage()    {}
func (*Test2_Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6ecbfc68e85c65, []int{1, 0}
}

func (m *Test2_Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test2_Record.Unmarshal(m, b)
}
func (m *Test2_Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test2_Record.Marshal(b, m, deterministic)
}
func (m *Test2_Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test2_Record.Merge(m, src)
}
func (m *Test2_Record) XXX_Size() int {
	return xxx_messageInfo_Test2_Record.Size(m)
}
func (m *Test2_Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Test2_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Test2_Record proto.InternalMessageInfo

func (m *Test2_Record) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Test2_Record) GetTestArys() []int32 {
	if m != nil {
		return m.TestArys
	}
	return nil
}

type Test3 struct {
	Records              map[int32]*Test3_Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Crc32                uint32                  `protobuf:"varint,2,opt,name=crc32,proto3" json:"crc32,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Test3) Reset()         { *m = Test3{} }
func (m *Test3) String() string { return proto.CompactTextString(m) }
func (*Test3) ProtoMessage()    {}
func (*Test3) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6ecbfc68e85c65, []int{2}
}

func (m *Test3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test3.Unmarshal(m, b)
}
func (m *Test3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test3.Marshal(b, m, deterministic)
}
func (m *Test3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test3.Merge(m, src)
}
func (m *Test3) XXX_Size() int {
	return xxx_messageInfo_Test3.Size(m)
}
func (m *Test3) XXX_DiscardUnknown() {
	xxx_messageInfo_Test3.DiscardUnknown(m)
}

var xxx_messageInfo_Test3 proto.InternalMessageInfo

func (m *Test3) GetRecords() map[int32]*Test3_Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *Test3) GetCrc32() uint32 {
	if m != nil {
		return m.Crc32
	}
	return 0
}

type Test3_Record struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TestString           string   `protobuf:"bytes,2,opt,name=test_string,json=testString,proto3" json:"test_string,omitempty"`
	TestDefault          string   `protobuf:"bytes,3,opt,name=test_default,json=testDefault,proto3" json:"test_default,omitempty"`
	TestInt32            int32    `protobuf:"varint,4,opt,name=test_int32,json=testInt32,proto3" json:"test_int32,omitempty"`
	TestUint32           uint32   `protobuf:"varint,5,opt,name=test_uint32,json=testUint32,proto3" json:"test_uint32,omitempty"`
	TestInt64            int64    `protobuf:"varint,6,opt,name=test_int64,json=testInt64,proto3" json:"test_int64,omitempty"`
	TestUint64           uint64   `protobuf:"varint,7,opt,name=test_uint64,json=testUint64,proto3" json:"test_uint64,omitempty"`
	TestFloat32          float32  `protobuf:"fixed32,8,opt,name=test_float32,json=testFloat32,proto3" json:"test_float32,omitempty"`
	TestFloat64          float64  `protobuf:"fixed64,9,opt,name=test_float64,json=testFloat64,proto3" json:"test_float64,omitempty"`
	TestDouble           float64  `protobuf:"fixed64,10,opt,name=test_double,json=testDouble,proto3" json:"test_double,omitempty"`
	TestArys             []int32  `protobuf:"varint,11,rep,packed,name=test_arys,json=testArys,proto3" json:"test_arys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test3_Record) Reset()         { *m = Test3_Record{} }
func (m *Test3_Record) String() string { return proto.CompactTextString(m) }
func (*Test3_Record) ProtoMessage()    {}
func (*Test3_Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6ecbfc68e85c65, []int{2, 0}
}

func (m *Test3_Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test3_Record.Unmarshal(m, b)
}
func (m *Test3_Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test3_Record.Marshal(b, m, deterministic)
}
func (m *Test3_Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test3_Record.Merge(m, src)
}
func (m *Test3_Record) XXX_Size() int {
	return xxx_messageInfo_Test3_Record.Size(m)
}
func (m *Test3_Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Test3_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Test3_Record proto.InternalMessageInfo

func (m *Test3_Record) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Test3_Record) GetTestString() string {
	if m != nil {
		return m.TestString
	}
	return ""
}

func (m *Test3_Record) GetTestDefault() string {
	if m != nil {
		return m.TestDefault
	}
	return ""
}

func (m *Test3_Record) GetTestInt32() int32 {
	if m != nil {
		return m.TestInt32
	}
	return 0
}

func (m *Test3_Record) GetTestUint32() uint32 {
	if m != nil {
		return m.TestUint32
	}
	return 0
}

func (m *Test3_Record) GetTestInt64() int64 {
	if m != nil {
		return m.TestInt64
	}
	return 0
}

func (m *Test3_Record) GetTestUint64() uint64 {
	if m != nil {
		return m.TestUint64
	}
	return 0
}

func (m *Test3_Record) GetTestFloat32() float32 {
	if m != nil {
		return m.TestFloat32
	}
	return 0
}

func (m *Test3_Record) GetTestFloat64() float64 {
	if m != nil {
		return m.TestFloat64
	}
	return 0
}

func (m *Test3_Record) GetTestDouble() float64 {
	if m != nil {
		return m.TestDouble
	}
	return 0
}

func (m *Test3_Record) GetTestArys() []int32 {
	if m != nil {
		return m.TestArys
	}
	return nil
}

func init() {
	proto.RegisterType((*Test)(nil), "Conf.test")
	proto.RegisterMapType((map[int32]*Test_Record)(nil), "Conf.test.RecordsEntry")
	proto.RegisterType((*Test_Record)(nil), "Conf.test.Record")
	proto.RegisterType((*Test2)(nil), "Conf.test2")
	proto.RegisterMapType((map[string]*Test2_Record)(nil), "Conf.test2.RecordsEntry")
	proto.RegisterType((*Test2_Record)(nil), "Conf.test2.Record")
	proto.RegisterType((*Test3)(nil), "Conf.test3")
	proto.RegisterMapType((map[int32]*Test3_Record)(nil), "Conf.test3.RecordsEntry")
	proto.RegisterType((*Test3_Record)(nil), "Conf.test3.Record")
}

func init() { proto.RegisterFile("conf.proto", fileDescriptor_0b6ecbfc68e85c65) }

var fileDescriptor_0b6ecbfc68e85c65 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0x4d, 0xab, 0xd4, 0x30,
	0x14, 0x25, 0x69, 0x3b, 0xef, 0xf5, 0xf6, 0x3d, 0xd1, 0x20, 0x18, 0x9e, 0x88, 0x71, 0x36, 0x66,
	0x55, 0xb0, 0xad, 0x45, 0xdc, 0xa9, 0xa3, 0xe0, 0xc2, 0x59, 0x54, 0xdc, 0xb8, 0x19, 0x3a, 0xfd,
	0x90, 0x62, 0x69, 0x25, 0x6d, 0x85, 0x2e, 0x05, 0x7f, 0x8d, 0xbf, 0xc8, 0xa5, 0x3f, 0x45, 0x72,
	0xa3, 0x33, 0xed, 0x8c, 0xa3, 0xb8, 0x73, 0xe1, 0xaa, 0xcd, 0xb9, 0xe7, 0xe4, 0xde, 0x93, 0x1c,
	0x02, 0x90, 0xb5, 0x4d, 0xe9, 0x7f, 0x50, 0x6d, 0xdf, 0x32, 0xfb, 0x59, 0xdb, 0x94, 0xcb, 0x4f,
	0x36, 0xd8, 0x7d, 0xd1, 0xf5, 0xec, 0x01, 0x9c, 0xa9, 0x22, 0x6b, 0x55, 0xde, 0x71, 0x22, 0x2c,
	0xe9, 0x05, 0xb7, 0x7c, 0x4d, 0xf0, 0x75, 0xd1, 0x4f, 0x4c, 0xe5, 0x79, 0xd3, 0xab, 0x31, 0xf9,
	0xc9, 0x63, 0x37, 0xc1, 0xc9, 0x54, 0x16, 0x06, 0x9c, 0x0a, 0x22, 0x2f, 0x13, 0xb3, 0xb8, 0xfa,
	0x46, 0x61, 0x61, 0xf8, 0xec, 0x1a, 0xd0, 0x2a, 0xe7, 0x44, 0x10, 0xe9, 0x24, 0xb4, 0xca, 0xd9,
	0x5d, 0xf0, 0xf4, 0x76, 0x9b, 0xae, 0x57, 0x55, 0xf3, 0x0e, 0x65, 0x6e, 0x02, 0x1a, 0x7a, 0x8d,
	0x08, 0xbb, 0x07, 0x17, 0x48, 0xc8, 0x8b, 0x32, 0x1d, 0xea, 0x9e, 0x5b, 0xc8, 0x40, 0xd1, 0xca,
	0x40, 0xec, 0x0e, 0xa0, 0x60, 0x53, 0x35, 0x7d, 0x18, 0x70, 0x1b, 0xf7, 0x76, 0x35, 0xf2, 0x52,
	0x03, 0xbb, 0x16, 0x83, 0xa9, 0x3b, 0x38, 0x19, 0x2a, 0xde, 0x20, 0x32, 0xd5, 0xc7, 0x11, 0x5f,
	0x08, 0x22, 0xad, 0x9d, 0x3e, 0x8e, 0x66, 0xfa, 0x38, 0xe2, 0x67, 0x82, 0x48, 0x7b, 0xaf, 0x8f,
	0xa3, 0xdd, 0x88, 0x65, 0xdd, 0xa6, 0xba, 0xc3, 0xb9, 0x20, 0x92, 0x9a, 0x11, 0x5f, 0x18, 0x68,
	0x4e, 0x89, 0x23, 0xee, 0x0a, 0x22, 0xc9, 0x84, 0x32, 0x69, 0x93, 0xb7, 0xc3, 0xb6, 0x2e, 0x38,
	0x20, 0x03, 0xdb, 0xac, 0x10, 0x61, 0xb7, 0x01, 0x87, 0xda, 0xa4, 0x6a, 0xec, 0xb8, 0x27, 0x2c,
	0xe9, 0x24, 0xe7, 0x1a, 0x78, 0xa2, 0xc6, 0xee, 0xea, 0x15, 0x5c, 0x4c, 0x6f, 0x84, 0x5d, 0x07,
	0xeb, 0x7d, 0x31, 0xfe, 0x38, 0x68, 0xfd, 0xcb, 0xee, 0x83, 0xf3, 0x31, 0xad, 0x87, 0x02, 0xcf,
	0xd8, 0x0b, 0x6e, 0x1c, 0xdd, 0x65, 0x62, 0xea, 0x8f, 0xe9, 0x23, 0xb2, 0xfc, 0x4a, 0xc0, 0xd1,
	0xa5, 0x80, 0x05, 0x87, 0x21, 0xe0, 0x7b, 0x61, 0xf0, 0x57, 0x29, 0x78, 0xf8, 0x8b, 0x10, 0xb8,
	0x18, 0x82, 0x99, 0x33, 0x7a, 0xe0, 0x6c, 0x7d, 0xda, 0x99, 0x6b, 0x9c, 0xc9, 0xb9, 0x33, 0x76,
	0x3c, 0xe0, 0xd4, 0xda, 0x67, 0xdb, 0x58, 0x0b, 0x7f, 0x6b, 0x2d, 0xfc, 0x1f, 0xf0, 0x7f, 0x2b,
	0xe0, 0xeb, 0x3f, 0x06, 0xfc, 0x74, 0x0c, 0xc2, 0xe3, 0x18, 0x3c, 0xbd, 0x7c, 0x6b, 0xeb, 0x97,
	0xef, 0x0b, 0xc5, 0xcf, 0x76, 0x81, 0x2f, 0x60, 0xf8, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x43, 0x7b,
	0x0f, 0xa9, 0x0f, 0x05, 0x00, 0x00,
}
