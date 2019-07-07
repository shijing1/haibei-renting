// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetHouseInfo

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Sessionid            string   `protobuf:"bytes,1,opt,name=Sessionid,proto3" json:"Sessionid,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	Housedata            []byte   `protobuf:"bytes,3,opt,name=Housedata,proto3" json:"Housedata,omitempty"`
	Userid               int64    `protobuf:"varint,4,opt,name=Userid,proto3" json:"Userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetHousedata() []byte {
	if m != nil {
		return m.Housedata
	}
	return nil
}

func (m *Response) GetUserid() int64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.GetHouseInfo.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetHouseInfo.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetHouseInfo.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.GetHouseInfo.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.GetHouseInfo.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.GetHouseInfo.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.GetHouseInfo.Pong")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0x4b, 0xeb, 0x40,
	0x10, 0xed, 0x26, 0xfd, 0x1c, 0xca, 0xa5, 0x77, 0x29, 0x97, 0xdc, 0x56, 0xb4, 0xac, 0x2f, 0x11,
	0x25, 0x16, 0x7d, 0xf0, 0x17, 0x14, 0xed, 0x83, 0x20, 0x5b, 0x7c, 0xf0, 0x31, 0x36, 0xe3, 0x12,
	0x34, 0xbb, 0x75, 0x77, 0x2b, 0xfa, 0x73, 0xfc, 0xa7, 0x92, 0xcd, 0x96, 0x16, 0xb1, 0xc1, 0xa7,
	0xe4, 0xcc, 0x9c, 0xc3, 0x39, 0x33, 0xb3, 0x30, 0x5e, 0x69, 0x65, 0xd5, 0x39, 0xbe, 0xa7, 0xc5,
	0xea, 0x05, 0x37, 0xdf, 0xc4, 0x55, 0xe9, 0x7f, 0xa1, 0x92, 0x22, 0x5f, 0x6a, 0x95, 0x18, 0xfd,
	0x96, 0x5c, 0xa3, 0xbd, 0x51, 0x6b, 0x83, 0x73, 0xf9, 0xa4, 0xd8, 0x18, 0x3a, 0xb7, 0x68, 0x4c,
	0x2a, 0x90, 0x0e, 0x20, 0x34, 0xe9, 0x47, 0x44, 0x26, 0x24, 0xee, 0xf1, 0xf2, 0x97, 0x5d, 0x41,
	0x87, 0xe3, 0xeb, 0x1a, 0x8d, 0xa5, 0x07, 0xd0, 0x5b, 0xa0, 0x31, 0xb9, 0x92, 0x79, 0xe6, 0x29,
	0xdb, 0x02, 0xfd, 0x03, 0xc1, 0x3c, 0x8b, 0x02, 0x57, 0x0e, 0xe6, 0x19, 0x93, 0xd0, 0xe5, 0x68,
	0x56, 0x4a, 0x1a, 0xa4, 0x43, 0x68, 0xcd, 0xb4, 0x96, 0xca, 0xab, 0x2a, 0x40, 0xff, 0x41, 0x7b,
	0xa6, 0x75, 0x61, 0x84, 0x57, 0x79, 0x54, 0xfa, 0xb8, 0x70, 0x59, 0x6a, 0xd3, 0x28, 0x9c, 0x90,
	0xb8, 0xcf, 0xb7, 0x85, 0x52, 0x75, 0x6f, 0x50, 0xe7, 0x59, 0xd4, 0x9c, 0x90, 0x38, 0xe4, 0x1e,
	0xb1, 0x18, 0x06, 0x0b, 0xab, 0x31, 0x2d, 0x72, 0x29, 0x36, 0x89, 0x87, 0xd0, 0x5a, 0xaa, 0xb5,
	0xb4, 0xce, 0x37, 0xe4, 0x15, 0x60, 0x27, 0xf0, 0x77, 0x87, 0xb9, 0x8d, 0xf8, 0x03, 0xf5, 0x10,
	0x9a, 0x77, 0xb9, 0x14, 0xa5, 0xa9, 0xb1, 0x5a, 0x3d, 0xa3, 0x6f, 0x7b, 0xe4, 0xfa, 0x6a, 0x7f,
	0xff, 0xe2, 0x33, 0x80, 0xce, 0xac, 0xba, 0x03, 0x7d, 0x80, 0xfe, 0xee, 0xda, 0x29, 0x4b, 0xf6,
	0x9e, 0x24, 0xf1, 0x03, 0x8c, 0x8e, 0x6b, 0x39, 0x55, 0x74, 0xd6, 0xa0, 0x02, 0xda, 0xd5, 0x44,
	0xf4, 0xb4, 0x46, 0xf0, 0x7d, 0x3d, 0xa3, 0xb3, 0xdf, 0x91, 0x37, 0x36, 0x53, 0x42, 0x39, 0x74,
	0xcb, 0x7d, 0xb8, 0x99, 0x8f, 0x6a, 0xd4, 0x25, 0x69, 0x54, 0x4b, 0x50, 0x52, 0xb0, 0x46, 0x4c,
	0xa6, 0xe4, 0xb1, 0xed, 0x1e, 0xe8, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x62, 0x42,
	0xbd, 0xbf, 0x02, 0x00, 0x00,
}
