// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dailysvc.proto

package grpc

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
	return fileDescriptor_dailysvc_d219c32d03b8d1a3, []int{0}
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
	return fileDescriptor_dailysvc_d219c32d03b8d1a3, []int{1}
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
	return fileDescriptor_dailysvc_d219c32d03b8d1a3, []int{2}
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

type PushPeriodRequest struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Temps                []*Temperature `protobuf:"bytes,2,rep,name=temps,proto3" json:"temps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PushPeriodRequest) Reset()         { *m = PushPeriodRequest{} }
func (m *PushPeriodRequest) String() string { return proto.CompactTextString(m) }
func (*PushPeriodRequest) ProtoMessage()    {}
func (*PushPeriodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dailysvc_d219c32d03b8d1a3, []int{3}
}
func (m *PushPeriodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushPeriodRequest.Unmarshal(m, b)
}
func (m *PushPeriodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushPeriodRequest.Marshal(b, m, deterministic)
}
func (dst *PushPeriodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushPeriodRequest.Merge(dst, src)
}
func (m *PushPeriodRequest) XXX_Size() int {
	return xxx_messageInfo_PushPeriodRequest.Size(m)
}
func (m *PushPeriodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushPeriodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushPeriodRequest proto.InternalMessageInfo

func (m *PushPeriodRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PushPeriodRequest) GetTemps() []*Temperature {
	if m != nil {
		return m.Temps
	}
	return nil
}

type PushPeriodResponse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushPeriodResponse) Reset()         { *m = PushPeriodResponse{} }
func (m *PushPeriodResponse) String() string { return proto.CompactTextString(m) }
func (*PushPeriodResponse) ProtoMessage()    {}
func (*PushPeriodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dailysvc_d219c32d03b8d1a3, []int{4}
}
func (m *PushPeriodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushPeriodResponse.Unmarshal(m, b)
}
func (m *PushPeriodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushPeriodResponse.Marshal(b, m, deterministic)
}
func (dst *PushPeriodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushPeriodResponse.Merge(dst, src)
}
func (m *PushPeriodResponse) XXX_Size() int {
	return xxx_messageInfo_PushPeriodResponse.Size(m)
}
func (m *PushPeriodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushPeriodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushPeriodResponse proto.InternalMessageInfo

func (m *PushPeriodResponse) GetErr() string {
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
	return fileDescriptor_dailysvc_d219c32d03b8d1a3, []int{5}
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
	return fileDescriptor_dailysvc_d219c32d03b8d1a3, []int{6}
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
	proto.RegisterType((*Temperature)(nil), "grpc.Temperature")
	proto.RegisterType((*GetPeriodRequest)(nil), "grpc.GetPeriodRequest")
	proto.RegisterType((*GetPeriodResponse)(nil), "grpc.GetPeriodResponse")
	proto.RegisterType((*PushPeriodRequest)(nil), "grpc.PushPeriodRequest")
	proto.RegisterType((*PushPeriodResponse)(nil), "grpc.PushPeriodResponse")
	proto.RegisterType((*GetUpdateDateRequest)(nil), "grpc.GetUpdateDateRequest")
	proto.RegisterType((*GetUpdateDateResponse)(nil), "grpc.GetUpdateDateResponse")
	proto.RegisterMapType((map[string]string)(nil), "grpc.GetUpdateDateResponse.DatesEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DailySVCClient is the client API for DailySVC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DailySVCClient interface {
	GetPeriod(ctx context.Context, in *GetPeriodRequest, opts ...grpc.CallOption) (*GetPeriodResponse, error)
	PushPeriod(ctx context.Context, in *PushPeriodRequest, opts ...grpc.CallOption) (*PushPeriodResponse, error)
	GetUpdateDate(ctx context.Context, in *GetUpdateDateRequest, opts ...grpc.CallOption) (*GetUpdateDateResponse, error)
}

type dailySVCClient struct {
	cc *grpc.ClientConn
}

func NewDailySVCClient(cc *grpc.ClientConn) DailySVCClient {
	return &dailySVCClient{cc}
}

func (c *dailySVCClient) GetPeriod(ctx context.Context, in *GetPeriodRequest, opts ...grpc.CallOption) (*GetPeriodResponse, error) {
	out := new(GetPeriodResponse)
	err := c.cc.Invoke(ctx, "/grpc.DailySVC/GetPeriod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dailySVCClient) PushPeriod(ctx context.Context, in *PushPeriodRequest, opts ...grpc.CallOption) (*PushPeriodResponse, error) {
	out := new(PushPeriodResponse)
	err := c.cc.Invoke(ctx, "/grpc.DailySVC/PushPeriod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dailySVCClient) GetUpdateDate(ctx context.Context, in *GetUpdateDateRequest, opts ...grpc.CallOption) (*GetUpdateDateResponse, error) {
	out := new(GetUpdateDateResponse)
	err := c.cc.Invoke(ctx, "/grpc.DailySVC/GetUpdateDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DailySVCServer is the server API for DailySVC service.
type DailySVCServer interface {
	GetPeriod(context.Context, *GetPeriodRequest) (*GetPeriodResponse, error)
	PushPeriod(context.Context, *PushPeriodRequest) (*PushPeriodResponse, error)
	GetUpdateDate(context.Context, *GetUpdateDateRequest) (*GetUpdateDateResponse, error)
}

func RegisterDailySVCServer(s *grpc.Server, srv DailySVCServer) {
	s.RegisterService(&_DailySVC_serviceDesc, srv)
}

func _DailySVC_GetPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DailySVCServer).GetPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DailySVC/GetPeriod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DailySVCServer).GetPeriod(ctx, req.(*GetPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DailySVC_PushPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DailySVCServer).PushPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DailySVC/PushPeriod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DailySVCServer).PushPeriod(ctx, req.(*PushPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DailySVC_GetUpdateDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpdateDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DailySVCServer).GetUpdateDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DailySVC/GetUpdateDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DailySVCServer).GetUpdateDate(ctx, req.(*GetUpdateDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DailySVC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.DailySVC",
	HandlerType: (*DailySVCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPeriod",
			Handler:    _DailySVC_GetPeriod_Handler,
		},
		{
			MethodName: "PushPeriod",
			Handler:    _DailySVC_PushPeriod_Handler,
		},
		{
			MethodName: "GetUpdateDate",
			Handler:    _DailySVC_GetUpdateDate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dailysvc.proto",
}

func init() { proto.RegisterFile("dailysvc.proto", fileDescriptor_dailysvc_d219c32d03b8d1a3) }

var fileDescriptor_dailysvc_d219c32d03b8d1a3 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcb, 0xca, 0xd3, 0x40,
	0x18, 0xed, 0x24, 0xad, 0x98, 0xaf, 0xf8, 0xf3, 0x77, 0xf8, 0xad, 0x21, 0x2e, 0x0c, 0xb3, 0xa8,
	0xdd, 0x98, 0x42, 0xdd, 0x14, 0x11, 0x41, 0x5b, 0x29, 0x14, 0x91, 0x12, 0x2f, 0xfb, 0xb4, 0x19,
	0xd3, 0xd0, 0x4b, 0xe2, 0xcc, 0xa4, 0x90, 0xc7, 0xf1, 0x6d, 0x7c, 0x05, 0x37, 0x3e, 0x8b, 0xcc,
	0x4c, 0x6e, 0x35, 0xa5, 0xbb, 0xef, 0x7a, 0xce, 0x99, 0xf3, 0x25, 0x70, 0x17, 0x06, 0xf1, 0x21,
	0xe7, 0xe7, 0xad, 0x97, 0xb2, 0x44, 0x24, 0xb8, 0x1b, 0xb1, 0x74, 0xeb, 0xbc, 0x8a, 0x62, 0xb1,
	0xcb, 0x36, 0xde, 0x36, 0x39, 0x4e, 0xa2, 0x24, 0x4a, 0x26, 0xaa, 0xb9, 0xc9, 0x7e, 0xa8, 0x4c,
	0x25, 0x2a, 0xd2, 0x4b, 0x64, 0x0e, 0xfd, 0xaf, 0xf4, 0x98, 0x52, 0x16, 0x88, 0x8c, 0x51, 0x8c,
	0xa1, 0x1b, 0x06, 0x82, 0xda, 0xc8, 0x45, 0x63, 0xcb, 0x57, 0x31, 0x76, 0xa1, 0x2f, 0xea, 0x11,
	0xdb, 0x70, 0xd1, 0xd8, 0xf0, 0x9b, 0x25, 0xb2, 0x82, 0xfb, 0x25, 0x15, 0x6b, 0xca, 0xe2, 0x24,
	0xf4, 0xe9, 0xcf, 0x8c, 0x72, 0x81, 0xef, 0xc0, 0x88, 0xc3, 0x02, 0xc7, 0x88, 0x43, 0xfc, 0x00,
	0x3d, 0x2e, 0x02, 0x26, 0xd4, 0xbe, 0xe5, 0xeb, 0x04, 0xdf, 0x83, 0x49, 0x4f, 0xa1, 0x6d, 0xaa,
	0x9a, 0x0c, 0xc9, 0x67, 0x18, 0x34, 0xb0, 0x78, 0x9a, 0x9c, 0x38, 0xc5, 0x2f, 0xa1, 0x27, 0xf9,
	0xb8, 0x8d, 0x5c, 0x73, 0xdc, 0x9f, 0x0e, 0x3c, 0xf9, 0x54, 0xaf, 0x21, 0xdc, 0xd7, 0x7d, 0x85,
	0xc7, 0x58, 0xc1, 0x21, 0x43, 0xf2, 0x09, 0x06, 0xeb, 0x8c, 0xef, 0x6e, 0x8b, 0xab, 0xf0, 0x8d,
	0xdb, 0xf8, 0x64, 0x04, 0xb8, 0x89, 0x56, 0xc8, 0x2b, 0x58, 0x51, 0xcd, 0xea, 0xc1, 0xc3, 0x92,
	0x8a, 0x6f, 0xa9, 0x34, 0x70, 0x11, 0x08, 0x5a, 0x12, 0x0f, 0xc1, 0x8c, 0x43, 0xfd, 0x0c, 0xeb,
	0x43, 0xf7, 0xf7, 0xdf, 0x17, 0x1d, 0x5f, 0x16, 0xc8, 0x2f, 0x04, 0x4f, 0xff, 0x5b, 0x28, 0xb0,
	0xdf, 0x42, 0x4f, 0xd6, 0xca, 0xa7, 0x8f, 0xb4, 0xb4, 0xab, 0xb3, 0x9e, 0x4c, 0xf8, 0xc7, 0x93,
	0x60, 0xb9, 0xaf, 0x97, 0xda, 0x7e, 0x38, 0x33, 0x80, 0x7a, 0x4c, 0xf6, 0xf7, 0x34, 0x2f, 0x95,
	0xef, 0x69, 0x2e, 0xef, 0x74, 0x0e, 0x0e, 0x19, 0x2d, 0xef, 0xa4, 0x92, 0x37, 0xc6, 0x0c, 0x4d,
	0xff, 0x20, 0x78, 0xbc, 0x90, 0x9f, 0xdc, 0x97, 0xef, 0x73, 0xfc, 0x0e, 0xac, 0xea, 0x4c, 0x78,
	0x58, 0x89, 0xba, 0xb0, 0xd9, 0x79, 0xd6, 0xaa, 0x6b, 0xa1, 0xa4, 0x83, 0xdf, 0x03, 0xd4, 0x46,
	0xe2, 0x62, 0xb0, 0x75, 0x28, 0xc7, 0x6e, 0x37, 0x2a, 0x88, 0x15, 0x3c, 0xb9, 0xb0, 0x01, 0x3b,
	0x57, 0xbd, 0xd1, 0x40, 0xcf, 0x6f, 0xf8, 0x46, 0x3a, 0x9b, 0x47, 0xea, 0x6f, 0x78, 0xfd, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0x26, 0x19, 0xc9, 0xc5, 0x54, 0x03, 0x00, 0x00,
}
