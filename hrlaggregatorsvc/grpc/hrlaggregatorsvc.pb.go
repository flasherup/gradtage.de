// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hrlaggregatorsvc.proto

package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Status struct {
	Station              string   `protobuf:"bytes,1,opt,name=station,proto3" json:"station,omitempty"`
	Update               string   `protobuf:"bytes,2,opt,name=update,proto3" json:"update,omitempty"`
	Temperature          float32  `protobuf:"fixed32,3,opt,name=temperature,proto3" json:"temperature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_hrlaggregatorsvc_18a6040e2606bff4, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetStation() string {
	if m != nil {
		return m.Station
	}
	return ""
}

func (m *Status) GetUpdate() string {
	if m != nil {
		return m.Update
	}
	return ""
}

func (m *Status) GetTemperature() float32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

type GetStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatusRequest) Reset()         { *m = GetStatusRequest{} }
func (m *GetStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatusRequest) ProtoMessage()    {}
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hrlaggregatorsvc_18a6040e2606bff4, []int{1}
}
func (m *GetStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatusRequest.Unmarshal(m, b)
}
func (m *GetStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatusRequest.Marshal(b, m, deterministic)
}
func (dst *GetStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatusRequest.Merge(dst, src)
}
func (m *GetStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatusRequest.Size(m)
}
func (m *GetStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatusRequest proto.InternalMessageInfo

type GetStatusResponse struct {
	Status               []*Status `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty"`
	Err                  string    `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetStatusResponse) Reset()         { *m = GetStatusResponse{} }
func (m *GetStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatusResponse) ProtoMessage()    {}
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hrlaggregatorsvc_18a6040e2606bff4, []int{2}
}
func (m *GetStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatusResponse.Unmarshal(m, b)
}
func (m *GetStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatusResponse.Marshal(b, m, deterministic)
}
func (dst *GetStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatusResponse.Merge(dst, src)
}
func (m *GetStatusResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatusResponse.Size(m)
}
func (m *GetStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatusResponse proto.InternalMessageInfo

func (m *GetStatusResponse) GetStatus() []*Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetStatusResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "grpc.Status")
	proto.RegisterType((*GetStatusRequest)(nil), "grpc.GetStatusRequest")
	proto.RegisterType((*GetStatusResponse)(nil), "grpc.GetStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HrlAggregatorSVCClient is the client API for HrlAggregatorSVC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HrlAggregatorSVCClient interface {
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
}

type hrlAggregatorSVCClient struct {
	cc *grpc.ClientConn
}

func NewHrlAggregatorSVCClient(cc *grpc.ClientConn) HrlAggregatorSVCClient {
	return &hrlAggregatorSVCClient{cc}
}

func (c *hrlAggregatorSVCClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/grpc.HrlAggregatorSVC/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HrlAggregatorSVCServer is the server API for HrlAggregatorSVC service.
type HrlAggregatorSVCServer interface {
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
}

func RegisterHrlAggregatorSVCServer(s *grpc.Server, srv HrlAggregatorSVCServer) {
	s.RegisterService(&_HrlAggregatorSVC_serviceDesc, srv)
}

func _HrlAggregatorSVC_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HrlAggregatorSVCServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.HrlAggregatorSVC/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HrlAggregatorSVCServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HrlAggregatorSVC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.HrlAggregatorSVC",
	HandlerType: (*HrlAggregatorSVCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _HrlAggregatorSVC_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hrlaggregatorsvc.proto",
}

func init() {
	proto.RegisterFile("hrlaggregatorsvc.proto", fileDescriptor_hrlaggregatorsvc_18a6040e2606bff4)
}

var fileDescriptor_hrlaggregatorsvc_18a6040e2606bff4 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x49, 0x83, 0x82, 0x7a, 0x65, 0x08, 0x37, 0x04, 0x8b, 0xc9, 0xb2, 0x18, 0x32, 0x65,
	0x28, 0x3b, 0x12, 0x62, 0x00, 0x89, 0xcd, 0x95, 0x98, 0x58, 0x4c, 0x39, 0x19, 0xa4, 0x52, 0x9b,
	0xf3, 0x99, 0xdf, 0x8f, 0x9a, 0x98, 0x2a, 0x82, 0xcd, 0xf7, 0x3d, 0xcb, 0xef, 0xf3, 0x41, 0xf7,
	0xce, 0x3b, 0xe7, 0x3d, 0x93, 0x77, 0x12, 0x38, 0x7d, 0x6f, 0x87, 0xc8, 0x41, 0x02, 0x9e, 0x7a,
	0x8e, 0x5b, 0xf3, 0x02, 0xcd, 0x46, 0x9c, 0xe4, 0x84, 0x0a, 0xce, 0x92, 0x38, 0xf9, 0x08, 0x7b,
	0x55, 0xe9, 0xaa, 0x5f, 0xda, 0xdf, 0x11, 0x3b, 0x68, 0x72, 0x7c, 0x73, 0x42, 0x6a, 0x31, 0x06,
	0x65, 0x42, 0x0d, 0x2b, 0xa1, 0xcf, 0x48, 0xec, 0x24, 0x33, 0xa9, 0x5a, 0x57, 0xfd, 0xc2, 0xce,
	0x91, 0x41, 0x68, 0x1f, 0x48, 0xa6, 0x02, 0x4b, 0x5f, 0x99, 0x92, 0x98, 0x27, 0xb8, 0x98, 0xb1,
	0x14, 0xc3, 0x3e, 0x11, 0x5e, 0x43, 0x93, 0x46, 0xa2, 0x2a, 0x5d, 0xf7, 0xab, 0xf5, 0xf9, 0x70,
	0xb0, 0x1b, 0xca, 0xad, 0x92, 0x61, 0x0b, 0x35, 0x31, 0x17, 0x8b, 0xc3, 0x71, 0x6d, 0xa1, 0x7d,
	0xe4, 0xdd, 0xdd, 0xf1, 0x7b, 0x9b, 0xe7, 0x7b, 0xbc, 0x85, 0xe5, 0xb1, 0x00, 0xbb, 0xe9, 0xa1,
	0xbf, 0x16, 0x57, 0x97, 0xff, 0xf8, 0x64, 0x62, 0x4e, 0x5e, 0x9b, 0x71, 0x3f, 0x37, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xaa, 0x23, 0x18, 0x7a, 0x39, 0x01, 0x00, 0x00,
}
