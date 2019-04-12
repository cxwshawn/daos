// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv.proto

package srv

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NotifyReadyReq struct {
	Uri                  string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Nctxs                uint32   `protobuf:"varint,2,opt,name=nctxs,proto3" json:"nctxs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyReadyReq) Reset()         { *m = NotifyReadyReq{} }
func (m *NotifyReadyReq) String() string { return proto.CompactTextString(m) }
func (*NotifyReadyReq) ProtoMessage()    {}
func (*NotifyReadyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_srv_064e53ee8ccaf2f4, []int{0}
}
func (m *NotifyReadyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyReadyReq.Unmarshal(m, b)
}
func (m *NotifyReadyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyReadyReq.Marshal(b, m, deterministic)
}
func (dst *NotifyReadyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyReadyReq.Merge(dst, src)
}
func (m *NotifyReadyReq) XXX_Size() int {
	return xxx_messageInfo_NotifyReadyReq.Size(m)
}
func (m *NotifyReadyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyReadyReq.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyReadyReq proto.InternalMessageInfo

func (m *NotifyReadyReq) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *NotifyReadyReq) GetNctxs() uint32 {
	if m != nil {
		return m.Nctxs
	}
	return 0
}

func init() {
	proto.RegisterType((*NotifyReadyReq)(nil), "srv.NotifyReadyReq")
}

func init() { proto.RegisterFile("srv.proto", fileDescriptor_srv_064e53ee8ccaf2f4) }

var fileDescriptor_srv_064e53ee8ccaf2f4 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2e, 0x2a, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2e, 0x2a, 0x53, 0xb2, 0xe0, 0xe2, 0xf3, 0xcb,
	0x2f, 0xc9, 0x4c, 0xab, 0x0c, 0x4a, 0x4d, 0x4c, 0xa9, 0x0c, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62,
	0x2e, 0x2d, 0xca, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58,
	0xf3, 0x92, 0x4b, 0x2a, 0x8a, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x83, 0x20, 0x9c, 0x24, 0x36,
	0xb0, 0x29, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x62, 0xdc, 0x0f, 0x52, 0x00, 0x00,
	0x00,
}
