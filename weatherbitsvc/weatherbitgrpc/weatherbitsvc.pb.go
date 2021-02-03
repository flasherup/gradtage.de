// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: weatherbitsvc.proto

package weatherbitgrpc

import (
	context "context"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Temperature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date        string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Temperature float64 `protobuf:"fixed64,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *Temperature) Reset() {
	*x = Temperature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weatherbitsvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Temperature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Temperature) ProtoMessage() {}

func (x *Temperature) ProtoReflect() protoreflect.Message {
	mi := &file_weatherbitsvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Temperature.ProtoReflect.Descriptor instead.
func (*Temperature) Descriptor() ([]byte, []int) {
	return file_weatherbitsvc_proto_rawDescGZIP(), []int{0}
}

func (x *Temperature) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Temperature) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

type Temperatures struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temps []*Temperature `protobuf:"bytes,1,rep,name=temps,proto3" json:"temps,omitempty"`
}

func (x *Temperatures) Reset() {
	*x = Temperatures{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weatherbitsvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Temperatures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Temperatures) ProtoMessage() {}

func (x *Temperatures) ProtoReflect() protoreflect.Message {
	mi := &file_weatherbitsvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Temperatures.ProtoReflect.Descriptor instead.
func (*Temperatures) Descriptor() ([]byte, []int) {
	return file_weatherbitsvc_proto_rawDescGZIP(), []int{1}
}

func (x *Temperatures) GetTemps() []*Temperature {
	if x != nil {
		return x.Temps
	}
	return nil
}

type GetPeriodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids   []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Start string   `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End   string   `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetPeriodRequest) Reset() {
	*x = GetPeriodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weatherbitsvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeriodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeriodRequest) ProtoMessage() {}

func (x *GetPeriodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weatherbitsvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeriodRequest.ProtoReflect.Descriptor instead.
func (*GetPeriodRequest) Descriptor() ([]byte, []int) {
	return file_weatherbitsvc_proto_rawDescGZIP(), []int{2}
}

func (x *GetPeriodRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetPeriodRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *GetPeriodRequest) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type GetPeriodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temps map[string]*Temperatures `protobuf:"bytes,1,rep,name=temps,proto3" json:"temps,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Err   string                   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *GetPeriodResponse) Reset() {
	*x = GetPeriodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weatherbitsvc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeriodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeriodResponse) ProtoMessage() {}

func (x *GetPeriodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weatherbitsvc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeriodResponse.ProtoReflect.Descriptor instead.
func (*GetPeriodResponse) Descriptor() ([]byte, []int) {
	return file_weatherbitsvc_proto_rawDescGZIP(), []int{3}
}

func (x *GetPeriodResponse) GetTemps() map[string]*Temperatures {
	if x != nil {
		return x.Temps
	}
	return nil
}

func (x *GetPeriodResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_weatherbitsvc_proto protoreflect.FileDescriptor

var file_weatherbitsvc_proto_rawDesc = []byte{
	0x0a, 0x13, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x62, 0x69, 0x74, 0x73, 0x76, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x62, 0x69,
	0x74, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x41, 0x0a, 0x0c, 0x54, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x74, 0x65, 0x6d,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x62, 0x69, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x73, 0x22, 0x4c, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0xc1, 0x01, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x62, 0x69, 0x74, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x74,
	0x65, 0x6d, 0x70, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x1a, 0x56, 0x0a, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x62,
	0x69, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x6a,
	0x0a, 0x14, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x42, 0x69, 0x74, 0x53, 0x63, 0x72, 0x61,
	0x70, 0x65, 0x72, 0x53, 0x56, 0x43, 0x12, 0x52, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x20, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x62, 0x69, 0x74,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x62,
	0x69, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_weatherbitsvc_proto_rawDescOnce sync.Once
	file_weatherbitsvc_proto_rawDescData = file_weatherbitsvc_proto_rawDesc
)

func file_weatherbitsvc_proto_rawDescGZIP() []byte {
	file_weatherbitsvc_proto_rawDescOnce.Do(func() {
		file_weatherbitsvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_weatherbitsvc_proto_rawDescData)
	})
	return file_weatherbitsvc_proto_rawDescData
}

var file_weatherbitsvc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_weatherbitsvc_proto_goTypes = []interface{}{
	(*Temperature)(nil),       // 0: weatherbitgrpc.Temperature
	(*Temperatures)(nil),      // 1: weatherbitgrpc.Temperatures
	(*GetPeriodRequest)(nil),  // 2: weatherbitgrpc.GetPeriodRequest
	(*GetPeriodResponse)(nil), // 3: weatherbitgrpc.GetPeriodResponse
	nil,                       // 4: weatherbitgrpc.GetPeriodResponse.TempsEntry
}
var file_weatherbitsvc_proto_depIdxs = []int32{
	0, // 0: weatherbitgrpc.Temperatures.temps:type_name -> weatherbitgrpc.Temperature
	4, // 1: weatherbitgrpc.GetPeriodResponse.temps:type_name -> weatherbitgrpc.GetPeriodResponse.TempsEntry
	1, // 2: weatherbitgrpc.GetPeriodResponse.TempsEntry.value:type_name -> weatherbitgrpc.Temperatures
	2, // 3: weatherbitgrpc.WeatherBitScraperSVC.GetPeriod:input_type -> weatherbitgrpc.GetPeriodRequest
	3, // 4: weatherbitgrpc.WeatherBitScraperSVC.GetPeriod:output_type -> weatherbitgrpc.GetPeriodResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_weatherbitsvc_proto_init() }
func file_weatherbitsvc_proto_init() {
	if File_weatherbitsvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weatherbitsvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Temperature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weatherbitsvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Temperatures); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weatherbitsvc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeriodRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weatherbitsvc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeriodResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_weatherbitsvc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weatherbitsvc_proto_goTypes,
		DependencyIndexes: file_weatherbitsvc_proto_depIdxs,
		MessageInfos:      file_weatherbitsvc_proto_msgTypes,
	}.Build()
	File_weatherbitsvc_proto = out.File
	file_weatherbitsvc_proto_rawDesc = nil
	file_weatherbitsvc_proto_goTypes = nil
	file_weatherbitsvc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WeatherBitScraperSVCClient is the client API for WeatherBitScraperSVC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WeatherBitScraperSVCClient interface {
	GetPeriod(ctx context.Context, in *GetPeriodRequest, opts ...grpc.CallOption) (*GetPeriodResponse, error)
}

type weatherBitScraperSVCClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherBitScraperSVCClient(cc grpc.ClientConnInterface) WeatherBitScraperSVCClient {
	return &weatherBitScraperSVCClient{cc}
}

func (c *weatherBitScraperSVCClient) GetPeriod(ctx context.Context, in *GetPeriodRequest, opts ...grpc.CallOption) (*GetPeriodResponse, error) {
	out := new(GetPeriodResponse)
	err := c.cc.Invoke(ctx, "/weatherbitgrpc.WeatherBitScraperSVC/GetPeriod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherBitScraperSVCServer is the server API for WeatherBitScraperSVC service.
type WeatherBitScraperSVCServer interface {
	GetPeriod(context.Context, *GetPeriodRequest) (*GetPeriodResponse, error)
}

// UnimplementedWeatherBitScraperSVCServer can be embedded to have forward compatible implementations.
type UnimplementedWeatherBitScraperSVCServer struct {
}

func (*UnimplementedWeatherBitScraperSVCServer) GetPeriod(context.Context, *GetPeriodRequest) (*GetPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeriod not implemented")
}

func RegisterWeatherBitScraperSVCServer(s *grpc.Server, srv WeatherBitScraperSVCServer) {
	s.RegisterService(&_WeatherBitScraperSVC_serviceDesc, srv)
}

func _WeatherBitScraperSVC_GetPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherBitScraperSVCServer).GetPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weatherbitgrpc.WeatherBitScraperSVC/GetPeriod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherBitScraperSVCServer).GetPeriod(ctx, req.(*GetPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WeatherBitScraperSVC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weatherbitgrpc.WeatherBitScraperSVC",
	HandlerType: (*WeatherBitScraperSVCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPeriod",
			Handler:    _WeatherBitScraperSVC_GetPeriod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weatherbitsvc.proto",
}
