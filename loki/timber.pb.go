package loki

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
	return fileDescriptor_timber_7e218a50bdf694bf, []int{0}
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
	return fileDescriptor_timber_7e218a50bdf694bf, []int{1}
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
	return fileDescriptor_timber_7e218a50bdf694bf, []int{2}
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
	proto.RegisterType((*PushRequest)(nil), "loki.PushRequest")
	proto.RegisterType((*Stream)(nil), "loki.Stream")
	proto.RegisterType((*Entry)(nil), "loki.Entry")
}

func init() { proto.RegisterFile("timber.proto", fileDescriptor_timber_7e218a50bdf694bf) }

var fileDescriptor_timber_7e218a50bdf694bf = []byte{
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xd9, 0x75, 0xed, 0xb2, 0x93, 0x3d, 0xe5, 0x20, 0x65, 0x2f, 0x4a, 0x41, 0xe9, 0x29,
	0x85, 0x8a, 0xe0, 0x0b, 0x88, 0x57, 0x89, 0xfa, 0x00, 0x0d, 0x8c, 0x35, 0x98, 0x34, 0x35, 0x99,
	0x1e, 0x7c, 0x7b, 0xe9, 0xc4, 0xb8, 0xb7, 0xe4, 0x9b, 0x7f, 0xbe, 0xf9, 0xe1, 0x48, 0xd6, 0x1b,
	0x8c, 0x6a, 0x8e, 0x81, 0x82, 0xdc, 0xb9, 0xf0, 0x65, 0x4f, 0xd7, 0x63, 0x08, 0xa3, 0xc3, 0x8e,
	0x99, 0x59, 0x3e, 0x3a, 0xb2, 0x1e, 0x13, 0x0d, 0x7e, 0xce, 0xb1, 0xe6, 0x01, 0xc4, 0xcb, 0x92,
	0x3e, 0x35, 0x7e, 0x2f, 0x98, 0x48, 0xde, 0xc1, 0x3e, 0x51, 0xc4, 0xc1, 0xa7, 0x7a, 0x73, 0x73,
	0xd1, 0x8a, 0xfe, 0xa8, 0x56, 0x8f, 0x7a, 0x65, 0xa8, 0xcb, 0xb0, 0x79, 0x86, 0x2a, 0x23, 0x79,
	0x05, 0x95, 0x1b, 0x0c, 0xba, 0x75, 0x61, 0xd3, 0x1e, 0xf4, 0xdf, 0x4f, 0xde, 0xc2, 0x1e, 0x27,
	0x8a, 0x16, 0x53, 0xbd, 0x65, 0x93, 0xc8, 0xa6, 0xa7, 0x89, 0xe2, 0x8f, 0x2e, 0xb3, 0xe6, 0x1d,
	0x2e, 0x99, 0xc8, 0x47, 0x38, 0xfc, 0x77, 0x63, 0x95, 0xe8, 0x4f, 0x2a, 0xb7, 0x57, 0xa5, 0xbd,
	0x7a, 0x2b, 0x09, 0x7d, 0x0e, 0x4b, 0x09, 0x3b, 0x67, 0x27, 0xac, 0xb7, 0x7c, 0x9f, 0xdf, 0xa6,
	0xe2, 0x95, 0xfb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x0d, 0xc4, 0x7c, 0x14, 0x01, 0x00,
	0x00,
}
