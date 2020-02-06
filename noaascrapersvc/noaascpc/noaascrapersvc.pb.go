// Code generated by protoc-gen-go. DO NOT EDIT.
// source: noaascrapersvc.proto

package noaascpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Temperature struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Temperature          float32  `protobuf:"fixed32,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Temperature) Reset()         { *m = Temperature{} }
func (m *Temperature) String() string { return proto.CompactTextString(m) }
func (*Temperature) ProtoMessage()    {}
func (*Temperature) Descriptor() ([]byte, []int) {
	return fileDescriptor_noaascrapersvc_1f1c4aac620752f6, []int{0}
}
func (m *Temperature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Temperature.Unmarshal(m, b)
}
func (m *Temperature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Temperature.Marshal(b, m, deterministic)
}
func (dst *Temperature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Temperature.Merge(dst, src)
}
func (m *Temperature) XXX_Size() int {
	return xxx_messageInfo_Temperature.Size(m)
}
func (m *Temperature) XXX_DiscardUnknown() {
	xxx_messageInfo_Temperature.DiscardUnknown(m)
}

var xxx_messageInfo_Temperature proto.InternalMessageInfo

func (m *Temperature) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Temperature) GetTemperature() float32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

type GetPeriodRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Start                string   `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  string   `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPeriodRequest) Reset()         { *m = GetPeriodRequest{} }
func (m *GetPeriodRequest) String() string { return proto.CompactTextString(m) }
func (*GetPeriodRequest) ProtoMessage()    {}
func (*GetPeriodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_noaascrapersvc_1f1c4aac620752f6, []int{1}
}
func (m *GetPeriodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPeriodRequest.Unmarshal(m, b)
}
func (m *GetPeriodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPeriodRequest.Marshal(b, m, deterministic)
}
func (dst *GetPeriodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPeriodRequest.Merge(dst, src)
}
func (m *GetPeriodRequest) XXX_Size() int {
	return xxx_messageInfo_GetPeriodRequest.Size(m)
}
func (m *GetPeriodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPeriodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPeriodRequest proto.InternalMessageInfo

func (m *GetPeriodRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPeriodRequest) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *GetPeriodRequest) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type GetPeriodResponse struct {
	Temps                []*Temperature `protobuf:"bytes,1,rep,name=temps,proto3" json:"temps,omitempty"`
	Err                  string         `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetPeriodResponse) Reset()         { *m = GetPeriodResponse{} }
func (m *GetPeriodResponse) String() string { return proto.CompactTextString(m) }
func (*GetPeriodResponse) ProtoMessage()    {}
func (*GetPeriodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_noaascrapersvc_1f1c4aac620752f6, []int{2}
}
func (m *GetPeriodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPeriodResponse.Unmarshal(m, b)
}
func (m *GetPeriodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPeriodResponse.Marshal(b, m, deterministic)
}
func (dst *GetPeriodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPeriodResponse.Merge(dst, src)
}
func (m *GetPeriodResponse) XXX_Size() int {
	return xxx_messageInfo_GetPeriodResponse.Size(m)
}
func (m *GetPeriodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPeriodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPeriodResponse proto.InternalMessageInfo

func (m *GetPeriodResponse) GetTemps() []*Temperature {
	if m != nil {
		return m.Temps
	}
	return nil
}

func (m *GetPeriodResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetUpdateDateRequest struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUpdateDateRequest) Reset()         { *m = GetUpdateDateRequest{} }
func (m *GetUpdateDateRequest) String() string { return proto.CompactTextString(m) }
func (*GetUpdateDateRequest) ProtoMessage()    {}
func (*GetUpdateDateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_noaascrapersvc_1f1c4aac620752f6, []int{3}
}
func (m *GetUpdateDateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUpdateDateRequest.Unmarshal(m, b)
}
func (m *GetUpdateDateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUpdateDateRequest.Marshal(b, m, deterministic)
}
func (dst *GetUpdateDateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUpdateDateRequest.Merge(dst, src)
}
func (m *GetUpdateDateRequest) XXX_Size() int {
	return xxx_messageInfo_GetUpdateDateRequest.Size(m)
}
func (m *GetUpdateDateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUpdateDateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUpdateDateRequest proto.InternalMessageInfo

func (m *GetUpdateDateRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetUpdateDateResponse struct {
	Dates                map[string]string `protobuf:"bytes,1,rep,name=dates,proto3" json:"dates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Err                  string            `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetUpdateDateResponse) Reset()         { *m = GetUpdateDateResponse{} }
func (m *GetUpdateDateResponse) String() string { return proto.CompactTextString(m) }
func (*GetUpdateDateResponse) ProtoMessage()    {}
func (*GetUpdateDateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_noaascrapersvc_1f1c4aac620752f6, []int{4}
}
func (m *GetUpdateDateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUpdateDateResponse.Unmarshal(m, b)
}
func (m *GetUpdateDateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUpdateDateResponse.Marshal(b, m, deterministic)
}
func (dst *GetUpdateDateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUpdateDateResponse.Merge(dst, src)
}
func (m *GetUpdateDateResponse) XXX_Size() int {
	return xxx_messageInfo_GetUpdateDateResponse.Size(m)
}
func (m *GetUpdateDateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUpdateDateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUpdateDateResponse proto.InternalMessageInfo

func (m *GetUpdateDateResponse) GetDates() map[string]string {
	if m != nil {
		return m.Dates
	}
	return nil
}

func (m *GetUpdateDateResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Temperature)(nil), "noaascpc.Temperature")
	proto.RegisterType((*GetPeriodRequest)(nil), "noaascpc.GetPeriodRequest")
	proto.RegisterType((*GetPeriodResponse)(nil), "noaascpc.GetPeriodResponse")
	proto.RegisterType((*GetUpdateDateRequest)(nil), "noaascpc.GetUpdateDateRequest")
	proto.RegisterType((*GetUpdateDateResponse)(nil), "noaascpc.GetUpdateDateResponse")
	proto.RegisterMapType((map[string]string)(nil), "noaascpc.GetUpdateDateResponse.DatesEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NoaaScraperSVCClient is the client API for NoaaScraperSVC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NoaaScraperSVCClient interface {
	GetPeriod(ctx context.Context, in *GetPeriodRequest, opts ...grpc.CallOption) (*GetPeriodResponse, error)
	GetUpdateDate(ctx context.Context, in *GetUpdateDateRequest, opts ...grpc.CallOption) (*GetUpdateDateResponse, error)
}

type noaaScraperSVCClient struct {
	cc *grpc.ClientConn
}

func NewNoaaScraperSVCClient(cc *grpc.ClientConn) NoaaScraperSVCClient {
	return &noaaScraperSVCClient{cc}
}

func (c *noaaScraperSVCClient) GetPeriod(ctx context.Context, in *GetPeriodRequest, opts ...grpc.CallOption) (*GetPeriodResponse, error) {
	out := new(GetPeriodResponse)
	err := c.cc.Invoke(ctx, "/noaascpc.NoaaScraperSVC/GetPeriod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noaaScraperSVCClient) GetUpdateDate(ctx context.Context, in *GetUpdateDateRequest, opts ...grpc.CallOption) (*GetUpdateDateResponse, error) {
	out := new(GetUpdateDateResponse)
	err := c.cc.Invoke(ctx, "/noaascpc.NoaaScraperSVC/GetUpdateDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoaaScraperSVCServer is the server API for NoaaScraperSVC service.
type NoaaScraperSVCServer interface {
	GetPeriod(context.Context, *GetPeriodRequest) (*GetPeriodResponse, error)
	GetUpdateDate(context.Context, *GetUpdateDateRequest) (*GetUpdateDateResponse, error)
}

func RegisterNoaaScraperSVCServer(s *grpc.Server, srv NoaaScraperSVCServer) {
	s.RegisterService(&_NoaaScraperSVC_serviceDesc, srv)
}

func _NoaaScraperSVC_GetPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoaaScraperSVCServer).GetPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noaascpc.NoaaScraperSVC/GetPeriod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoaaScraperSVCServer).GetPeriod(ctx, req.(*GetPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoaaScraperSVC_GetUpdateDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpdateDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoaaScraperSVCServer).GetUpdateDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noaascpc.NoaaScraperSVC/GetUpdateDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoaaScraperSVCServer).GetUpdateDate(ctx, req.(*GetUpdateDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NoaaScraperSVC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "noaascpc.NoaaScraperSVC",
	HandlerType: (*NoaaScraperSVCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPeriod",
			Handler:    _NoaaScraperSVC_GetPeriod_Handler,
		},
		{
			MethodName: "GetUpdateDate",
			Handler:    _NoaaScraperSVC_GetUpdateDate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noaascrapersvc.proto",
}

func init() {
	proto.RegisterFile("noaascrapersvc.proto", fileDescriptor_noaascrapersvc_1f1c4aac620752f6)
}

var fileDescriptor_noaascrapersvc_1f1c4aac620752f6 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4e, 0xe3, 0x30,
	0x10, 0x6d, 0x92, 0x76, 0xb5, 0x99, 0x6a, 0xab, 0xae, 0xd5, 0xae, 0xa2, 0xac, 0xb4, 0x8d, 0x72,
	0xaa, 0x16, 0x91, 0x4a, 0xe5, 0x52, 0x71, 0x42, 0x14, 0xa8, 0xc4, 0x01, 0x21, 0x17, 0xb8, 0xbb,
	0x89, 0x29, 0x11, 0x34, 0x0e, 0xb6, 0x53, 0xa9, 0x1f, 0xc4, 0x8d, 0x0f, 0xe1, 0x2b, 0xf8, 0x16,
	0x64, 0x3b, 0xa1, 0x2d, 0x14, 0x71, 0x9b, 0x79, 0x7e, 0xf3, 0xfc, 0xe6, 0xd9, 0xd0, 0xc9, 0x18,
	0x21, 0x22, 0xe6, 0x24, 0xa7, 0x5c, 0x2c, 0xe3, 0x28, 0xe7, 0x4c, 0x32, 0xf4, 0xd3, 0xa0, 0x79,
	0xec, 0xef, 0xcf, 0x53, 0x79, 0x57, 0xcc, 0xa2, 0x98, 0x2d, 0x06, 0x73, 0x36, 0x67, 0x03, 0x4d,
	0x98, 0x15, 0xb7, 0xba, 0xd3, 0x8d, 0xae, 0xcc, 0x60, 0x38, 0x86, 0xe6, 0x15, 0x5d, 0xe4, 0x94,
	0x13, 0x59, 0x70, 0x8a, 0x10, 0xd4, 0x13, 0x22, 0xa9, 0x67, 0x05, 0x56, 0xdf, 0xc5, 0xba, 0x46,
	0x01, 0x34, 0xe5, 0x9a, 0xe2, 0xd9, 0x81, 0xd5, 0xb7, 0xf1, 0x26, 0x14, 0x9e, 0x43, 0x7b, 0x42,
	0xe5, 0x25, 0xe5, 0x29, 0x4b, 0x30, 0x7d, 0x2c, 0xa8, 0x90, 0xa8, 0x05, 0x76, 0x9a, 0x94, 0x3a,
	0x76, 0x9a, 0xa0, 0x0e, 0x34, 0x84, 0x24, 0x5c, 0xea, 0x79, 0x17, 0x9b, 0x06, 0xb5, 0xc1, 0xa1,
	0x59, 0xe2, 0x39, 0x1a, 0x53, 0x65, 0x88, 0xe1, 0xf7, 0x86, 0x96, 0xc8, 0x59, 0x26, 0x28, 0xda,
	0x83, 0x86, 0xba, 0x4f, 0x78, 0x56, 0xe0, 0xf4, 0x9b, 0xc3, 0x6e, 0x54, 0xad, 0x1b, 0x6d, 0x98,
	0xc7, 0x86, 0xa3, 0x35, 0x39, 0x2f, 0xef, 0x51, 0x65, 0x18, 0x41, 0x67, 0x42, 0xe5, 0x75, 0xae,
	0xd6, 0x39, 0x21, 0x92, 0x56, 0x1e, 0xff, 0x80, 0x93, 0x26, 0x46, 0xd4, 0x3d, 0xae, 0xbf, 0xbc,
	0xf6, 0x6a, 0x58, 0x01, 0xe1, 0x93, 0x05, 0xdd, 0x0f, 0x03, 0xa5, 0x91, 0x23, 0x68, 0x28, 0xac,
	0x32, 0xf2, 0x7f, 0x6d, 0x64, 0x27, 0x3f, 0x52, 0x8d, 0x38, 0xcd, 0x24, 0x5f, 0x61, 0x33, 0xf8,
	0xd9, 0x9d, 0x3f, 0x02, 0x58, 0xd3, 0xd4, 0xf9, 0x3d, 0x5d, 0x95, 0xc1, 0xa9, 0x52, 0x25, 0xb7,
	0x24, 0x0f, 0x05, 0xad, 0x92, 0xd3, 0xcd, 0xa1, 0x3d, 0xb2, 0x86, 0xcf, 0x16, 0xb4, 0x2e, 0x18,
	0x21, 0x53, 0xf3, 0x1d, 0xa6, 0x37, 0x63, 0x74, 0x06, 0xee, 0x7b, 0x7c, 0xc8, 0xdf, 0xb2, 0xb7,
	0xf5, 0x3e, 0xfe, 0xdf, 0x9d, 0x67, 0xc6, 0x76, 0x58, 0x43, 0x18, 0x7e, 0x6d, 0x6d, 0x84, 0xfe,
	0x7d, 0xb9, 0xaa, 0xd1, 0xeb, 0x7d, 0x13, 0x45, 0x58, 0x9b, 0xfd, 0xd0, 0x5f, 0xee, 0xe0, 0x2d,
	0x00, 0x00, 0xff, 0xff, 0x0b, 0x09, 0xb9, 0xd8, 0xc3, 0x02, 0x00, 0x00,
}
