package timberproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion2

type PushRequest struct {
	Streams              []*Stream `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PushRequest) Reset()         { *m = PushRequest{} }
func (m *PushRequest) String() string { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()    {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_timber_039dd8e9b63b8f16, []int{0}
}
func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushRequest.Unmarshal(m, b)
}
func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
}
func (dst *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(dst, src)
}
func (m *PushRequest) XXX_Size() int {
	return xxx_messageInfo_PushRequest.Size(m)
}
func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetStreams() []*Stream {
	if m != nil {
		return m.Streams
	}
	return nil
}

type Stream struct {
	Labels               string   `protobuf:"bytes,1,opt,name=labels,proto3" json:"labels,omitempty"`
	Entries              []*Entry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stream) Reset()         { *m = Stream{} }
func (m *Stream) String() string { return proto.CompactTextString(m) }
func (*Stream) ProtoMessage()    {}
func (*Stream) Descriptor() ([]byte, []int) {
	return fileDescriptor_timber_039dd8e9b63b8f16, []int{1}
}
func (m *Stream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stream.Unmarshal(m, b)
}
func (m *Stream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stream.Marshal(b, m, deterministic)
}
func (dst *Stream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stream.Merge(dst, src)
}
func (m *Stream) XXX_Size() int {
	return xxx_messageInfo_Stream.Size(m)
}
func (m *Stream) XXX_DiscardUnknown() {
	xxx_messageInfo_Stream.DiscardUnknown(m)
}

var xxx_messageInfo_Stream proto.InternalMessageInfo

func (m *Stream) GetLabels() string {
	if m != nil {
		return m.Labels
	}
	return ""
}

func (m *Stream) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Entry struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Line                 string               `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_timber_039dd8e9b63b8f16, []int{2}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (dst *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(dst, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Entry) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func init() {
	proto.RegisterType((*PushRequest)(nil), "timberproto.PushRequest")
	proto.RegisterType((*Stream)(nil), "timberproto.Stream")
	proto.RegisterType((*Entry)(nil), "timberproto.Entry")
}

func init() { proto.RegisterFile("timber.proto", fileDescriptor_timber_039dd8e9b63b8f16) }

var fileDescriptor_timber_039dd8e9b63b8f16 = []byte{
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3d, 0x4b, 0x04, 0x31,
	0x10, 0x86, 0xb9, 0x53, 0xf7, 0xb8, 0x89, 0xd5, 0x08, 0xb2, 0x5c, 0xa3, 0x6c, 0x75, 0x85, 0xe6,
	0x60, 0x6d, 0x2c, 0x6c, 0x6d, 0x45, 0xa2, 0xfe, 0x80, 0x0d, 0x8c, 0x6b, 0x20, 0xd9, 0xac, 0xc9,
	0x6c, 0xe1, 0xbf, 0x17, 0x27, 0xc6, 0x8f, 0x2e, 0xef, 0xe4, 0x99, 0x87, 0x79, 0xe1, 0x94, 0x5d,
	0xb0, 0x94, 0xf4, 0x9c, 0x22, 0x47, 0x54, 0x25, 0x49, 0xd8, 0x5d, 0x8c, 0x31, 0x8e, 0x9e, 0x0e,
	0x92, 0xec, 0xf2, 0x7a, 0x60, 0x17, 0x28, 0xf3, 0x10, 0xe6, 0x42, 0x77, 0x77, 0xa0, 0x1e, 0x97,
	0xfc, 0x66, 0xe8, 0x7d, 0xa1, 0xcc, 0x78, 0x0d, 0x9b, 0xcc, 0x89, 0x86, 0x90, 0xdb, 0xd5, 0xe5,
	0xd1, 0x5e, 0xf5, 0x67, 0xfa, 0x8f, 0x4e, 0x3f, 0xc9, 0x9f, 0xa9, 0x4c, 0xf7, 0x00, 0x4d, 0x19,
	0xe1, 0x39, 0x34, 0x7e, 0xb0, 0xe4, 0xbf, 0xf6, 0x56, 0xfb, 0xad, 0xf9, 0x4e, 0x78, 0x05, 0x1b,
	0x9a, 0x38, 0x39, 0xca, 0xed, 0x5a, 0x84, 0xf8, 0x4f, 0x78, 0x3f, 0x71, 0xfa, 0x30, 0x15, 0xe9,
	0x5e, 0xe0, 0x44, 0x26, 0x78, 0x0b, 0xdb, 0x9f, 0x4b, 0xc5, 0xa8, 0xfa, 0x9d, 0x2e, 0x5d, 0x74,
	0xed, 0xa2, 0x9f, 0x2b, 0x61, 0x7e, 0x61, 0x44, 0x38, 0xf6, 0x6e, 0xa2, 0x76, 0x2d, 0x67, 0xc8,
	0xdb, 0x36, 0xb2, 0x72, 0xf3, 0x19, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x6a, 0x08, 0x83, 0x29, 0x01,
	0x00, 0x00,
}
