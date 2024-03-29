// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_PostRet

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
	//手机号
	Mobile string `protobuf:"bytes,1,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	//密码
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	//短信验证码
	SmsCode              string   `protobuf:"bytes,3,opt,name=Sms_code,json=SmsCode,proto3" json:"Sms_code,omitempty"`
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

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Request) GetSmsCode() string {
	if m != nil {
		return m.SmsCode
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	SessionID            string   `protobuf:"bytes,3,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
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

func (m *Response) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
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
	proto.RegisterType((*Message)(nil), "go.micro.srv.PostRet.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostRet.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostRet.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.PostRet.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.PostRet.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.PostRet.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.PostRet.Pong")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x5f, 0x4b, 0xe3, 0x40,
	0x14, 0xc5, 0x9b, 0xed, 0xb6, 0x49, 0xef, 0x53, 0x77, 0x28, 0x4b, 0x37, 0x5d, 0x8b, 0xcc, 0x83,
	0xd6, 0x97, 0x58, 0xf4, 0x23, 0x68, 0x41, 0x85, 0x42, 0x49, 0x40, 0x7c, 0x11, 0x49, 0xdb, 0x4b,
	0x08, 0x36, 0x73, 0xeb, 0xdc, 0xa9, 0x7f, 0xbe, 0xbc, 0x48, 0x26, 0x53, 0x2b, 0x92, 0xe2, 0x53,
	0x72, 0xe6, 0x77, 0x72, 0x72, 0xee, 0xcc, 0xc0, 0x60, 0xad, 0xc9, 0xd0, 0x29, 0xbe, 0xa6, 0xc5,
	0x7a, 0x85, 0xdb, 0x67, 0x64, 0x57, 0x45, 0x2f, 0xa3, 0xa8, 0xc8, 0x17, 0x9a, 0x22, 0xd6, 0xcf,
	0xd1, 0x8c, 0xd8, 0xc4, 0x68, 0xe4, 0x00, 0xfc, 0x29, 0x32, 0xa7, 0x19, 0x8a, 0x2e, 0x34, 0x39,
	0x7d, 0xeb, 0x7b, 0x87, 0xde, 0xa8, 0x13, 0x97, 0xaf, 0xf2, 0x0e, 0xfc, 0x18, 0x9f, 0x36, 0xc8,
	0x46, 0xfc, 0x85, 0xf6, 0x94, 0xe6, 0xf9, 0x0a, 0x1d, 0x77, 0x4a, 0x84, 0x10, 0xcc, 0x52, 0xe6,
	0x17, 0xd2, 0xcb, 0xfe, 0x2f, 0x4b, 0x3e, 0xb5, 0xf8, 0x07, 0x41, 0x52, 0xf0, 0xc3, 0x82, 0x96,
	0xd8, 0x6f, 0x5a, 0xe6, 0x27, 0x05, 0x5f, 0xd0, 0x12, 0xe5, 0x2d, 0x04, 0x31, 0xf2, 0x9a, 0x14,
	0xa3, 0xe8, 0x41, 0x6b, 0xa2, 0xb5, 0x22, 0x97, 0x5c, 0x89, 0xf2, 0x87, 0x13, 0xad, 0x0b, 0xce,
	0x5c, 0xac, 0x53, 0xe2, 0x3f, 0x74, 0x12, 0x64, 0xce, 0x49, 0x5d, 0x5f, 0xba, 0xd4, 0xdd, 0x82,
	0x1c, 0x41, 0x37, 0x31, 0x1a, 0xd3, 0x22, 0x57, 0xd9, 0xb6, 0x7a, 0x0f, 0x5a, 0x0b, 0xda, 0x28,
	0x63, 0xf3, 0x9b, 0x71, 0x25, 0xe4, 0x09, 0xfc, 0xf9, 0xe2, 0xdc, 0x55, 0xa9, 0xb1, 0x0e, 0xe1,
	0xf7, 0x2c, 0x57, 0x59, 0x59, 0x89, 0x8d, 0xa6, 0x47, 0x74, 0xd8, 0x29, 0xcb, 0x69, 0x3f, 0x3f,
	0x7b, 0xf7, 0xc0, 0x9f, 0x54, 0x67, 0x21, 0x6e, 0xc0, 0x77, 0x5b, 0x2f, 0x0e, 0xa2, 0xba, 0x13,
	0x89, 0x5c, 0xed, 0x70, 0xb8, 0x0f, 0x57, 0x5d, 0x65, 0x43, 0xdc, 0x43, 0xbb, 0x1a, 0x41, 0x1c,
	0xd5, 0x7b, 0xbf, 0x6f, 0x45, 0x78, 0xfc, 0xa3, 0x6f, 0x1b, 0x3e, 0xf6, 0xc4, 0x15, 0x04, 0xe5,
	0xd8, 0x76, 0xb4, 0xb0, 0xfe, 0xc3, 0x92, 0x87, 0xfb, 0x18, 0xa9, 0x4c, 0x36, 0x46, 0xde, 0xd8,
	0x9b, 0xb7, 0xed, 0x0d, 0x3c, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x4b, 0x7c, 0x28, 0xa0,
	0x02, 0x00, 0x00,
}
