// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/getamis/alice/crypto/zkproof/message.proto

package zkproof

import (
	fmt "fmt"
	ecpointgrouplaw "github.com/getamis/alice/crypto/ecpointgrouplaw"
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

type IntegerFactorizationProofMessage struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Challenge            []byte   `protobuf:"bytes,2,opt,name=challenge,proto3" json:"challenge,omitempty"`
	Proof                []byte   `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegerFactorizationProofMessage) Reset()         { *m = IntegerFactorizationProofMessage{} }
func (m *IntegerFactorizationProofMessage) String() string { return proto.CompactTextString(m) }
func (*IntegerFactorizationProofMessage) ProtoMessage()    {}
func (*IntegerFactorizationProofMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7463df78901cfd5c, []int{0}
}

func (m *IntegerFactorizationProofMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegerFactorizationProofMessage.Unmarshal(m, b)
}
func (m *IntegerFactorizationProofMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegerFactorizationProofMessage.Marshal(b, m, deterministic)
}
func (m *IntegerFactorizationProofMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegerFactorizationProofMessage.Merge(m, src)
}
func (m *IntegerFactorizationProofMessage) XXX_Size() int {
	return xxx_messageInfo_IntegerFactorizationProofMessage.Size(m)
}
func (m *IntegerFactorizationProofMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegerFactorizationProofMessage.DiscardUnknown(m)
}

var xxx_messageInfo_IntegerFactorizationProofMessage proto.InternalMessageInfo

func (m *IntegerFactorizationProofMessage) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *IntegerFactorizationProofMessage) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func (m *IntegerFactorizationProofMessage) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type SchnorrProofMessage struct {
	Blake2BKey           []byte                          `protobuf:"bytes,1,opt,name=blake2bKey,proto3" json:"blake2bKey,omitempty"`
	V                    *ecpointgrouplaw.EcPointMessage `protobuf:"bytes,2,opt,name=V,proto3" json:"V,omitempty"`
	Alpha                *ecpointgrouplaw.EcPointMessage `protobuf:"bytes,3,opt,name=alpha,proto3" json:"alpha,omitempty"`
	U                    []byte                          `protobuf:"bytes,4,opt,name=u,proto3" json:"u,omitempty"`
	T                    []byte                          `protobuf:"bytes,5,opt,name=t,proto3" json:"t,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *SchnorrProofMessage) Reset()         { *m = SchnorrProofMessage{} }
func (m *SchnorrProofMessage) String() string { return proto.CompactTextString(m) }
func (*SchnorrProofMessage) ProtoMessage()    {}
func (*SchnorrProofMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7463df78901cfd5c, []int{1}
}

func (m *SchnorrProofMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchnorrProofMessage.Unmarshal(m, b)
}
func (m *SchnorrProofMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchnorrProofMessage.Marshal(b, m, deterministic)
}
func (m *SchnorrProofMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchnorrProofMessage.Merge(m, src)
}
func (m *SchnorrProofMessage) XXX_Size() int {
	return xxx_messageInfo_SchnorrProofMessage.Size(m)
}
func (m *SchnorrProofMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SchnorrProofMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SchnorrProofMessage proto.InternalMessageInfo

func (m *SchnorrProofMessage) GetBlake2BKey() []byte {
	if m != nil {
		return m.Blake2BKey
	}
	return nil
}

func (m *SchnorrProofMessage) GetV() *ecpointgrouplaw.EcPointMessage {
	if m != nil {
		return m.V
	}
	return nil
}

func (m *SchnorrProofMessage) GetAlpha() *ecpointgrouplaw.EcPointMessage {
	if m != nil {
		return m.Alpha
	}
	return nil
}

func (m *SchnorrProofMessage) GetU() []byte {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *SchnorrProofMessage) GetT() []byte {
	if m != nil {
		return m.T
	}
	return nil
}

func init() {
	proto.RegisterType((*IntegerFactorizationProofMessage)(nil), "zkproof.IntegerFactorizationProofMessage")
	proto.RegisterType((*SchnorrProofMessage)(nil), "zkproof.SchnorrProofMessage")
}

func init() {
	proto.RegisterFile("github.com/getamis/alice/crypto/zkproof/message.proto", fileDescriptor_7463df78901cfd5c)
}

var fileDescriptor_7463df78901cfd5c = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0xba, 0x8a, 0xb1, 0xa7, 0xe8, 0x61, 0x11, 0xd1, 0xd2, 0x93, 0x17, 0x37, 0x50,
	0xe9, 0xc9, 0xb3, 0x82, 0x88, 0x50, 0x2a, 0xf4, 0x9e, 0x0d, 0x63, 0x36, 0x34, 0xbb, 0x13, 0xb2,
	0x13, 0xa4, 0x7d, 0x31, 0x5f, 0x4f, 0x36, 0x29, 0x56, 0xbd, 0xd8, 0xe3, 0xf7, 0x4f, 0xe6, 0xff,
	0x92, 0xf0, 0x99, 0xb1, 0xd4, 0xc4, 0xba, 0xd2, 0xd8, 0x4a, 0x03, 0xa4, 0x5a, 0xdb, 0x4b, 0xe5,
	0xac, 0x06, 0xa9, 0xc3, 0xda, 0x13, 0xca, 0xcd, 0xca, 0x07, 0xc4, 0x77, 0xd9, 0x42, 0xdf, 0x2b,
	0x03, 0x95, 0x0f, 0x48, 0x28, 0x4e, 0xb6, 0xf1, 0xe5, 0xc3, 0x7f, 0xfb, 0xa0, 0x3d, 0xda, 0x8e,
	0x4c, 0xc0, 0xe8, 0x9d, 0xfa, 0x90, 0x89, 0x72, 0xcb, 0x84, 0xf8, 0xf8, 0xb9, 0x23, 0x30, 0x10,
	0x9e, 0x94, 0x26, 0x0c, 0x76, 0xa3, 0xc8, 0x62, 0x37, 0x1f, 0x9a, 0x5f, 0xb3, 0x4f, 0x5c, 0xf1,
	0x53, 0x1f, 0x6b, 0x67, 0xf5, 0x0b, 0xac, 0x4b, 0x36, 0x66, 0xb7, 0xa3, 0xc5, 0x2e, 0x18, 0xa6,
	0xba, 0x51, 0xce, 0x41, 0x67, 0xa0, 0x3c, 0xc8, 0xd3, 0xef, 0x40, 0x5c, 0xf0, 0x22, 0xdd, 0xb2,
	0x3c, 0x4c, 0x93, 0x0c, 0x93, 0x4f, 0xc6, 0xcf, 0xdf, 0x74, 0xd3, 0x61, 0x08, 0xbf, 0x4c, 0xd7,
	0x9c, 0xd7, 0x4e, 0xad, 0x60, 0x5a, 0xef, 0x54, 0x3f, 0x12, 0x71, 0xc7, 0xd9, 0x32, 0x39, 0xce,
	0xa6, 0x37, 0xd5, 0x9f, 0x67, 0x55, 0x8f, 0x7a, 0x3e, 0xf0, 0xb6, 0x6b, 0xc1, 0x96, 0x62, 0xc6,
	0x0b, 0xe5, 0x7c, 0xa3, 0x92, 0x7c, 0x8f, 0x95, 0x7c, 0x5a, 0x8c, 0x38, 0x8b, 0xe5, 0x51, 0x92,
	0xb3, 0x38, 0x10, 0x95, 0x45, 0x26, 0xaa, 0x8f, 0xd3, 0xb7, 0xdd, 0x7f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x11, 0xd2, 0xfa, 0xe4, 0xb5, 0x01, 0x00, 0x00,
}
