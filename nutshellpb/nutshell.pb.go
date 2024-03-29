// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nutshellpb/nutshell.proto

package nutshellpb

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

type CalculateRequest struct {
	Operation            string   `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Input1               int64    `protobuf:"varint,2,opt,name=input1,proto3" json:"input1,omitempty"`
	Input2               int64    `protobuf:"varint,3,opt,name=input2,proto3" json:"input2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculateRequest) Reset()         { *m = CalculateRequest{} }
func (m *CalculateRequest) String() string { return proto.CompactTextString(m) }
func (*CalculateRequest) ProtoMessage()    {}
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd8ae89d5945015a, []int{0}
}

func (m *CalculateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculateRequest.Unmarshal(m, b)
}
func (m *CalculateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculateRequest.Marshal(b, m, deterministic)
}
func (m *CalculateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculateRequest.Merge(m, src)
}
func (m *CalculateRequest) XXX_Size() int {
	return xxx_messageInfo_CalculateRequest.Size(m)
}
func (m *CalculateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculateRequest proto.InternalMessageInfo

func (m *CalculateRequest) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *CalculateRequest) GetInput1() int64 {
	if m != nil {
		return m.Input1
	}
	return 0
}

func (m *CalculateRequest) GetInput2() int64 {
	if m != nil {
		return m.Input2
	}
	return 0
}

type CalculateResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculateResponse) Reset()         { *m = CalculateResponse{} }
func (m *CalculateResponse) String() string { return proto.CompactTextString(m) }
func (*CalculateResponse) ProtoMessage()    {}
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd8ae89d5945015a, []int{1}
}

func (m *CalculateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculateResponse.Unmarshal(m, b)
}
func (m *CalculateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculateResponse.Marshal(b, m, deterministic)
}
func (m *CalculateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculateResponse.Merge(m, src)
}
func (m *CalculateResponse) XXX_Size() int {
	return xxx_messageInfo_CalculateResponse.Size(m)
}
func (m *CalculateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculateResponse proto.InternalMessageInfo

func (m *CalculateResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalculateRequest)(nil), "nutshellpb.CalculateRequest")
	proto.RegisterType((*CalculateResponse)(nil), "nutshellpb.CalculateResponse")
}

func init() { proto.RegisterFile("nutshellpb/nutshell.proto", fileDescriptor_cd8ae89d5945015a) }

var fileDescriptor_cd8ae89d5945015a = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x2b, 0x2d, 0x29,
	0xce, 0x48, 0xcd, 0xc9, 0x29, 0x48, 0xd2, 0x87, 0x31, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xb8, 0x10, 0x52, 0x4a, 0x09, 0x5c, 0x02, 0xce, 0x89, 0x39, 0xc9, 0xa5, 0x39, 0x89, 0x25, 0xa9,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x32, 0x5c, 0x9c, 0xf9, 0x05, 0xa9, 0x45, 0x89,
	0x25, 0x99, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0x21, 0x31, 0x2e,
	0xb6, 0xcc, 0xbc, 0x82, 0xd2, 0x12, 0x43, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x28, 0x0f,
	0x2e, 0x6e, 0x24, 0xc1, 0x8c, 0x24, 0x6e, 0xa4, 0xa4, 0xcd, 0x25, 0x88, 0x64, 0x43, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0x2a, 0x48, 0x71, 0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09, 0xd8, 0x7c, 0xe6, 0x20,
	0x28, 0xcf, 0x28, 0x84, 0x8b, 0x03, 0xe6, 0x38, 0x21, 0x0f, 0x2e, 0x4e, 0xb8, 0x46, 0x21, 0x19,
	0x3d, 0x84, 0xa3, 0xf5, 0xd0, 0x5d, 0x2c, 0x25, 0x8b, 0x43, 0x16, 0x62, 0x5b, 0x12, 0x1b, 0xd8,
	0xdf, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x9e, 0x67, 0xf3, 0x14, 0x01, 0x00, 0x00,
}
