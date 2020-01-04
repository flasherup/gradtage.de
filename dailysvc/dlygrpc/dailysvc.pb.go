// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dailysvc.proto

package dlygrpc

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
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{0}
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
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{1}
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
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{2}
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
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{3}
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
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{4}
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
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{5}
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
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{6}
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

type UpdateAvgForYearRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAvgForYearRequest) Reset()         { *m = UpdateAvgForYearRequest{} }
func (m *UpdateAvgForYearRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAvgForYearRequest) ProtoMessage()    {}
func (*UpdateAvgForYearRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{7}
}
func (m *UpdateAvgForYearRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAvgForYearRequest.Unmarshal(m, b)
}
func (m *UpdateAvgForYearRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAvgForYearRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateAvgForYearRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAvgForYearRequest.Merge(dst, src)
}
func (m *UpdateAvgForYearRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAvgForYearRequest.Size(m)
}
func (m *UpdateAvgForYearRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAvgForYearRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAvgForYearRequest proto.InternalMessageInfo

func (m *UpdateAvgForYearRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateAvgForYearResponse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAvgForYearResponse) Reset()         { *m = UpdateAvgForYearResponse{} }
func (m *UpdateAvgForYearResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateAvgForYearResponse) ProtoMessage()    {}
func (*UpdateAvgForYearResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{8}
}
func (m *UpdateAvgForYearResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAvgForYearResponse.Unmarshal(m, b)
}
func (m *UpdateAvgForYearResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAvgForYearResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateAvgForYearResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAvgForYearResponse.Merge(dst, src)
}
func (m *UpdateAvgForYearResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateAvgForYearResponse.Size(m)
}
func (m *UpdateAvgForYearResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAvgForYearResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAvgForYearResponse proto.InternalMessageInfo

func (m *UpdateAvgForYearResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type UpdateAvgForDOYRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Doy                  int32    `protobuf:"varint,2,opt,name=doy,proto3" json:"doy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAvgForDOYRequest) Reset()         { *m = UpdateAvgForDOYRequest{} }
func (m *UpdateAvgForDOYRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAvgForDOYRequest) ProtoMessage()    {}
func (*UpdateAvgForDOYRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{9}
}
func (m *UpdateAvgForDOYRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAvgForDOYRequest.Unmarshal(m, b)
}
func (m *UpdateAvgForDOYRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAvgForDOYRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateAvgForDOYRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAvgForDOYRequest.Merge(dst, src)
}
func (m *UpdateAvgForDOYRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAvgForDOYRequest.Size(m)
}
func (m *UpdateAvgForDOYRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAvgForDOYRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAvgForDOYRequest proto.InternalMessageInfo

func (m *UpdateAvgForDOYRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateAvgForDOYRequest) GetDoy() int32 {
	if m != nil {
		return m.Doy
	}
	return 0
}

type UpdateAvgForDOYResponse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAvgForDOYResponse) Reset()         { *m = UpdateAvgForDOYResponse{} }
func (m *UpdateAvgForDOYResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateAvgForDOYResponse) ProtoMessage()    {}
func (*UpdateAvgForDOYResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{10}
}
func (m *UpdateAvgForDOYResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAvgForDOYResponse.Unmarshal(m, b)
}
func (m *UpdateAvgForDOYResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAvgForDOYResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateAvgForDOYResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAvgForDOYResponse.Merge(dst, src)
}
func (m *UpdateAvgForDOYResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateAvgForDOYResponse.Size(m)
}
func (m *UpdateAvgForDOYResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAvgForDOYResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAvgForDOYResponse proto.InternalMessageInfo

func (m *UpdateAvgForDOYResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetAvgRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAvgRequest) Reset()         { *m = GetAvgRequest{} }
func (m *GetAvgRequest) String() string { return proto.CompactTextString(m) }
func (*GetAvgRequest) ProtoMessage()    {}
func (*GetAvgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{11}
}
func (m *GetAvgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAvgRequest.Unmarshal(m, b)
}
func (m *GetAvgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAvgRequest.Marshal(b, m, deterministic)
}
func (dst *GetAvgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAvgRequest.Merge(dst, src)
}
func (m *GetAvgRequest) XXX_Size() int {
	return xxx_messageInfo_GetAvgRequest.Size(m)
}
func (m *GetAvgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAvgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAvgRequest proto.InternalMessageInfo

func (m *GetAvgRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetAvgResponse struct {
	Temps                []*Temperature `protobuf:"bytes,1,rep,name=temps,proto3" json:"temps,omitempty"`
	Err                  string         `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetAvgResponse) Reset()         { *m = GetAvgResponse{} }
func (m *GetAvgResponse) String() string { return proto.CompactTextString(m) }
func (*GetAvgResponse) ProtoMessage()    {}
func (*GetAvgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dailysvc_5c8ba1da71bbf467, []int{12}
}
func (m *GetAvgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAvgResponse.Unmarshal(m, b)
}
func (m *GetAvgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAvgResponse.Marshal(b, m, deterministic)
}
func (dst *GetAvgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAvgResponse.Merge(dst, src)
}
func (m *GetAvgResponse) XXX_Size() int {
	return xxx_messageInfo_GetAvgResponse.Size(m)
}
func (m *GetAvgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAvgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAvgResponse proto.InternalMessageInfo

func (m *GetAvgResponse) GetTemps() []*Temperature {
	if m != nil {
		return m.Temps
	}
	return nil
}

func (m *GetAvgResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Temperature)(nil), "dlygrpc.Temperature")
	proto.RegisterType((*GetPeriodRequest)(nil), "dlygrpc.GetPeriodRequest")
	proto.RegisterType((*GetPeriodResponse)(nil), "dlygrpc.GetPeriodResponse")
	proto.RegisterType((*PushPeriodRequest)(nil), "dlygrpc.PushPeriodRequest")
	proto.RegisterType((*PushPeriodResponse)(nil), "dlygrpc.PushPeriodResponse")
	proto.RegisterType((*GetUpdateDateRequest)(nil), "dlygrpc.GetUpdateDateRequest")
	proto.RegisterType((*GetUpdateDateResponse)(nil), "dlygrpc.GetUpdateDateResponse")
	proto.RegisterMapType((map[string]string)(nil), "dlygrpc.GetUpdateDateResponse.DatesEntry")
	proto.RegisterType((*UpdateAvgForYearRequest)(nil), "dlygrpc.UpdateAvgForYearRequest")
	proto.RegisterType((*UpdateAvgForYearResponse)(nil), "dlygrpc.UpdateAvgForYearResponse")
	proto.RegisterType((*UpdateAvgForDOYRequest)(nil), "dlygrpc.UpdateAvgForDOYRequest")
	proto.RegisterType((*UpdateAvgForDOYResponse)(nil), "dlygrpc.UpdateAvgForDOYResponse")
	proto.RegisterType((*GetAvgRequest)(nil), "dlygrpc.GetAvgRequest")
	proto.RegisterType((*GetAvgResponse)(nil), "dlygrpc.GetAvgResponse")
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
	UpdateAvgForYear(ctx context.Context, in *UpdateAvgForYearRequest, opts ...grpc.CallOption) (*UpdateAvgForYearResponse, error)
	UpdateAvgForDOY(ctx context.Context, in *UpdateAvgForDOYRequest, opts ...grpc.CallOption) (*UpdateAvgForDOYResponse, error)
	GetAvg(ctx context.Context, in *GetAvgRequest, opts ...grpc.CallOption) (*GetAvgResponse, error)
}

type dailySVCClient struct {
	cc *grpc.ClientConn
}

func NewDailySVCClient(cc *grpc.ClientConn) DailySVCClient {
	return &dailySVCClient{cc}
}

func (c *dailySVCClient) GetPeriod(ctx context.Context, in *GetPeriodRequest, opts ...grpc.CallOption) (*GetPeriodResponse, error) {
	out := new(GetPeriodResponse)
	err := c.cc.Invoke(ctx, "/dlygrpc.DailySVC/GetPeriod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dailySVCClient) PushPeriod(ctx context.Context, in *PushPeriodRequest, opts ...grpc.CallOption) (*PushPeriodResponse, error) {
	out := new(PushPeriodResponse)
	err := c.cc.Invoke(ctx, "/dlygrpc.DailySVC/PushPeriod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dailySVCClient) GetUpdateDate(ctx context.Context, in *GetUpdateDateRequest, opts ...grpc.CallOption) (*GetUpdateDateResponse, error) {
	out := new(GetUpdateDateResponse)
	err := c.cc.Invoke(ctx, "/dlygrpc.DailySVC/GetUpdateDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dailySVCClient) UpdateAvgForYear(ctx context.Context, in *UpdateAvgForYearRequest, opts ...grpc.CallOption) (*UpdateAvgForYearResponse, error) {
	out := new(UpdateAvgForYearResponse)
	err := c.cc.Invoke(ctx, "/dlygrpc.DailySVC/UpdateAvgForYear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dailySVCClient) UpdateAvgForDOY(ctx context.Context, in *UpdateAvgForDOYRequest, opts ...grpc.CallOption) (*UpdateAvgForDOYResponse, error) {
	out := new(UpdateAvgForDOYResponse)
	err := c.cc.Invoke(ctx, "/dlygrpc.DailySVC/UpdateAvgForDOY", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dailySVCClient) GetAvg(ctx context.Context, in *GetAvgRequest, opts ...grpc.CallOption) (*GetAvgResponse, error) {
	out := new(GetAvgResponse)
	err := c.cc.Invoke(ctx, "/dlygrpc.DailySVC/GetAvg", in, out, opts...)
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
	UpdateAvgForYear(context.Context, *UpdateAvgForYearRequest) (*UpdateAvgForYearResponse, error)
	UpdateAvgForDOY(context.Context, *UpdateAvgForDOYRequest) (*UpdateAvgForDOYResponse, error)
	GetAvg(context.Context, *GetAvgRequest) (*GetAvgResponse, error)
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
		FullMethod: "/dlygrpc.DailySVC/GetPeriod",
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
		FullMethod: "/dlygrpc.DailySVC/PushPeriod",
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
		FullMethod: "/dlygrpc.DailySVC/GetUpdateDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DailySVCServer).GetUpdateDate(ctx, req.(*GetUpdateDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DailySVC_UpdateAvgForYear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvgForYearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DailySVCServer).UpdateAvgForYear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dlygrpc.DailySVC/UpdateAvgForYear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DailySVCServer).UpdateAvgForYear(ctx, req.(*UpdateAvgForYearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DailySVC_UpdateAvgForDOY_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvgForDOYRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DailySVCServer).UpdateAvgForDOY(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dlygrpc.DailySVC/UpdateAvgForDOY",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DailySVCServer).UpdateAvgForDOY(ctx, req.(*UpdateAvgForDOYRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DailySVC_GetAvg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DailySVCServer).GetAvg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dlygrpc.DailySVC/GetAvg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DailySVCServer).GetAvg(ctx, req.(*GetAvgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DailySVC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dlygrpc.DailySVC",
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
		{
			MethodName: "UpdateAvgForYear",
			Handler:    _DailySVC_UpdateAvgForYear_Handler,
		},
		{
			MethodName: "UpdateAvgForDOY",
			Handler:    _DailySVC_UpdateAvgForDOY_Handler,
		},
		{
			MethodName: "GetAvg",
			Handler:    _DailySVC_GetAvg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dailysvc.proto",
}

func init() { proto.RegisterFile("dailysvc.proto", fileDescriptor_dailysvc_5c8ba1da71bbf467) }

var fileDescriptor_dailysvc_5c8ba1da71bbf467 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xed, 0xa6, 0x90, 0x89, 0x08, 0xe9, 0x28, 0xa4, 0xc6, 0x08, 0x62, 0xf6, 0x80, 0x52,
	0x7e, 0x5c, 0xa9, 0x5c, 0xaa, 0x72, 0x40, 0xa5, 0x81, 0x48, 0x1c, 0x68, 0x30, 0x50, 0x29, 0x47,
	0x27, 0x5e, 0x5c, 0xab, 0x69, 0x6c, 0xd6, 0xeb, 0x48, 0x7e, 0x1f, 0x1e, 0x86, 0x77, 0x40, 0xe2,
	0x59, 0xd0, 0xae, 0x7f, 0x9b, 0xc4, 0xe1, 0xc2, 0x6d, 0x66, 0xf6, 0x9b, 0xf9, 0x3e, 0xcf, 0x7c,
	0x09, 0xb4, 0x5d, 0xc7, 0x9f, 0x27, 0xd1, 0x72, 0x66, 0x85, 0x2c, 0xe0, 0x01, 0xde, 0x71, 0xe7,
	0x89, 0xc7, 0xc2, 0x99, 0xf1, 0xca, 0xf3, 0xf9, 0x55, 0x3c, 0xb5, 0x66, 0xc1, 0xcd, 0x91, 0x17,
	0x78, 0xc1, 0x91, 0x7c, 0x9f, 0xc6, 0xdf, 0x65, 0x26, 0x13, 0x19, 0xa5, 0x7d, 0xe4, 0x1c, 0x5a,
	0x5f, 0xe9, 0x4d, 0x48, 0x99, 0xc3, 0x63, 0x46, 0x11, 0x61, 0xd7, 0x75, 0x38, 0xd5, 0x15, 0x53,
	0x19, 0x34, 0x6d, 0x19, 0xa3, 0x09, 0x2d, 0x5e, 0x42, 0x74, 0xd5, 0x54, 0x06, 0xaa, 0x5d, 0x2d,
	0x91, 0x8f, 0xd0, 0x19, 0x51, 0x3e, 0xa6, 0xcc, 0x0f, 0x5c, 0x9b, 0xfe, 0x88, 0x69, 0xc4, 0xb1,
	0x0d, 0xaa, 0xef, 0x66, 0x73, 0x54, 0xdf, 0xc5, 0x2e, 0x34, 0x22, 0xee, 0x30, 0x2e, 0xfb, 0x9b,
	0x76, 0x9a, 0x60, 0x07, 0x34, 0xba, 0x70, 0x75, 0x4d, 0xd6, 0x44, 0x48, 0x3e, 0xc3, 0x7e, 0x65,
	0x56, 0x14, 0x06, 0x8b, 0x88, 0xe2, 0x73, 0x68, 0x08, 0xbe, 0x48, 0x57, 0x4c, 0x6d, 0xd0, 0x3a,
	0xee, 0x5a, 0xd9, 0xd7, 0x5a, 0x15, 0xed, 0x76, 0x0a, 0x91, 0x23, 0x19, 0xcb, 0x68, 0x44, 0x48,
	0x2e, 0x60, 0x7f, 0x1c, 0x47, 0x57, 0xdb, 0xf5, 0x15, 0x14, 0xea, 0x3f, 0x29, 0xc8, 0x33, 0xc0,
	0xea, 0xc0, 0x4c, 0x64, 0x46, 0xac, 0x94, 0xc4, 0x16, 0x74, 0x47, 0x94, 0x7f, 0x0b, 0xc5, 0x1a,
	0x87, 0x0e, 0xa7, 0x39, 0x77, 0x0f, 0x34, 0xdf, 0x4d, 0x3f, 0xa6, 0xf9, 0x6e, 0xf7, 0xd7, 0x9f,
	0xfe, 0x8e, 0x2d, 0x0a, 0xe4, 0xa7, 0x02, 0x0f, 0x56, 0x1a, 0xb2, 0xd9, 0x6f, 0xa1, 0x21, 0x6a,
	0xf9, 0x02, 0x0e, 0x0b, 0x75, 0x1b, 0xe1, 0x96, 0x48, 0xa2, 0xf7, 0x0b, 0xce, 0x12, 0x3b, 0xed,
	0x5b, 0xdf, 0x8a, 0x71, 0x02, 0x50, 0xc2, 0xc4, 0xfb, 0x35, 0x4d, 0x72, 0xf1, 0xd7, 0x34, 0x11,
	0x07, 0x5b, 0x3a, 0xf3, 0x98, 0xe6, 0x07, 0x93, 0xc9, 0xa9, 0x7a, 0xa2, 0x90, 0x43, 0x38, 0x48,
	0x39, 0xcf, 0x96, 0xde, 0x87, 0x80, 0x4d, 0xa8, 0xc3, 0x6a, 0xb6, 0x4a, 0x5e, 0x82, 0xbe, 0x0e,
	0xad, 0xdd, 0xd7, 0x29, 0xf4, 0xaa, 0xe8, 0xe1, 0xc5, 0xa4, 0xee, 0x5a, 0x1d, 0xd0, 0xdc, 0x20,
	0x91, 0xd2, 0x1a, 0xb6, 0x08, 0xc9, 0x8b, 0xdb, 0xa2, 0x64, 0x6f, 0x2d, 0x51, 0x1f, 0xee, 0x8d,
	0x28, 0x3f, 0x5b, 0x7a, 0x75, 0xba, 0x3f, 0x41, 0x3b, 0x07, 0xfc, 0x0f, 0x0b, 0x1e, 0xff, 0xd6,
	0xe0, 0xee, 0x50, 0xfc, 0x62, 0xbf, 0x5c, 0x9e, 0xe3, 0x10, 0x9a, 0x85, 0xc5, 0xf1, 0x61, 0xf5,
	0x94, 0xb7, 0x2c, 0x6a, 0x18, 0x9b, 0x9e, 0x52, 0x39, 0x64, 0x07, 0x47, 0x00, 0xa5, 0x09, 0xb1,
	0xc4, 0xae, 0x59, 0xdd, 0x78, 0xb4, 0xf1, 0xad, 0x18, 0x34, 0x96, 0xcb, 0x28, 0x5d, 0x84, 0x8f,
	0xeb, 0xdc, 0x95, 0x8e, 0x7b, 0xb2, 0xdd, 0x7c, 0x64, 0x07, 0x27, 0xd0, 0x59, 0xbd, 0x3a, 0x9a,
	0x45, 0x57, 0x8d, 0x77, 0x8c, 0xa7, 0x5b, 0x10, 0xc5, 0xe8, 0x4b, 0xb8, 0xbf, 0x72, 0x66, 0xec,
	0x6f, 0xec, 0x2b, 0xcd, 0x63, 0x98, 0xf5, 0x80, 0x62, 0xee, 0x1b, 0xd8, 0x4b, 0x0f, 0x8e, 0xbd,
	0xea, 0xe7, 0x95, 0x16, 0x31, 0x0e, 0xd6, 0xea, 0x79, 0xf3, 0x74, 0x4f, 0xfe, 0x97, 0xbe, 0xfe,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x1e, 0x14, 0xa4, 0x95, 0x05, 0x00, 0x00,
}