// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sajari/api/pipeline/v1/record.proto

package piplinepb

import (
	record "code.sajari.com/protogen-go/sajari/engine/store/record"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// SearchRequest is a request to perform a search using a pipeline.
type AddRequest struct {
	// Pipeline to run.
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	// Values is a mapping of key -> value which should be substituted
	// into the pipeline inputs.
	Values map[string]string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// List of records to add.
	Records              []*record.Record `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1139fe48e665c6f, []int{0}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *AddRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *AddRequest) GetRecords() []*record.Record {
	if m != nil {
		return m.Records
	}
	return nil
}

// SearchResponse is a response to an Add call.
type AddResponse struct {
	Response             *record.AddResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1139fe48e665c6f, []int{1}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetResponse() *record.AddResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type ReplaceRequest struct {
	Pipeline             *Pipeline                          `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	Values               map[string]string                  `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	KeyRecords           []*record.ReplaceRequest_KeyRecord `protobuf:"bytes,3,rep,name=key_records,json=keyRecords,proto3" json:"key_records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ReplaceRequest) Reset()         { *m = ReplaceRequest{} }
func (m *ReplaceRequest) String() string { return proto.CompactTextString(m) }
func (*ReplaceRequest) ProtoMessage()    {}
func (*ReplaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1139fe48e665c6f, []int{2}
}

func (m *ReplaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceRequest.Unmarshal(m, b)
}
func (m *ReplaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceRequest.Marshal(b, m, deterministic)
}
func (m *ReplaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceRequest.Merge(m, src)
}
func (m *ReplaceRequest) XXX_Size() int {
	return xxx_messageInfo_ReplaceRequest.Size(m)
}
func (m *ReplaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceRequest proto.InternalMessageInfo

func (m *ReplaceRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *ReplaceRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ReplaceRequest) GetKeyRecords() []*record.ReplaceRequest_KeyRecord {
	if m != nil {
		return m.KeyRecords
	}
	return nil
}

type ReplaceResponse struct {
	Response             *record.ReplaceResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ReplaceResponse) Reset()         { *m = ReplaceResponse{} }
func (m *ReplaceResponse) String() string { return proto.CompactTextString(m) }
func (*ReplaceResponse) ProtoMessage()    {}
func (*ReplaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1139fe48e665c6f, []int{3}
}

func (m *ReplaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceResponse.Unmarshal(m, b)
}
func (m *ReplaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceResponse.Marshal(b, m, deterministic)
}
func (m *ReplaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceResponse.Merge(m, src)
}
func (m *ReplaceResponse) XXX_Size() int {
	return xxx_messageInfo_ReplaceResponse.Size(m)
}
func (m *ReplaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceResponse proto.InternalMessageInfo

func (m *ReplaceResponse) GetResponse() *record.ReplaceResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "sajari.api.pipeline.v1.AddRequest")
	proto.RegisterMapType((map[string]string)(nil), "sajari.api.pipeline.v1.AddRequest.ValuesEntry")
	proto.RegisterType((*AddResponse)(nil), "sajari.api.pipeline.v1.AddResponse")
	proto.RegisterType((*ReplaceRequest)(nil), "sajari.api.pipeline.v1.ReplaceRequest")
	proto.RegisterMapType((map[string]string)(nil), "sajari.api.pipeline.v1.ReplaceRequest.ValuesEntry")
	proto.RegisterType((*ReplaceResponse)(nil), "sajari.api.pipeline.v1.ReplaceResponse")
}

func init() {
	proto.RegisterFile("sajari/api/pipeline/v1/record.proto", fileDescriptor_d1139fe48e665c6f)
}

var fileDescriptor_d1139fe48e665c6f = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xa6, 0x09, 0xbd, 0xf7, 0x7a, 0x02, 0x2a, 0x83, 0x48, 0xc8, 0xaa, 0xe4, 0xaa, 0x2d, 0x14,
	0x27, 0x34, 0xba, 0xf0, 0xa7, 0x08, 0x55, 0x54, 0x50, 0x10, 0x19, 0xa9, 0x8b, 0x6e, 0x24, 0x4d,
	0x0e, 0x21, 0xb6, 0x66, 0xc6, 0x4c, 0x5a, 0xc8, 0x33, 0xb8, 0xf0, 0x45, 0x7c, 0x48, 0xc9, 0x64,
	0x12, 0x9b, 0xda, 0xb4, 0x05, 0x5d, 0xe5, 0x24, 0xf9, 0x7e, 0xce, 0xf9, 0x4e, 0x32, 0x70, 0x2d,
	0x83, 0xaf, 0x41, 0x96, 0x78, 0x81, 0x48, 0x3c, 0x91, 0x08, 0x5c, 0x27, 0x29, 0x7a, 0xdb, 0x89,
	0x97, 0x61, 0xc8, 0xb3, 0x88, 0x8a, 0x8c, 0xe7, 0x9c, 0xdc, 0xad, 0x40, 0x34, 0x10, 0x09, 0xad,
	0x41, 0x74, 0x3b, 0x71, 0xee, 0x77, 0x90, 0x1b, 0x8c, 0xa2, 0x3b, 0x43, 0x0d, 0xc3, 0x34, 0x2e,
	0x01, 0x32, 0xe7, 0x19, 0x6a, 0x83, 0x96, 0x8f, 0xe3, 0x76, 0xe8, 0x6d, 0x64, 0x10, 0x6b, 0x31,
	0xf7, 0x87, 0x01, 0x30, 0x8b, 0x22, 0x86, 0xdf, 0x37, 0x28, 0x73, 0x32, 0x85, 0xab, 0x1a, 0x69,
	0xf7, 0x06, 0xbd, 0x91, 0xe5, 0x0f, 0xe8, 0xe1, 0x6e, 0xe9, 0x47, 0x5d, 0xb3, 0x86, 0x41, 0xde,
	0xc0, 0xc5, 0x36, 0x58, 0x6f, 0x50, 0xda, 0xc6, 0xc0, 0x1c, 0x59, 0x3e, 0xed, 0xe2, 0xfe, 0x71,
	0xa4, 0x9f, 0x15, 0xe1, 0x75, 0x9a, 0x67, 0x05, 0xd3, 0x6c, 0x32, 0x85, 0xcb, 0x6a, 0x10, 0x69,
	0x9b, 0x4a, 0xc8, 0xad, 0x85, 0xaa, 0x99, 0xa9, 0x9a, 0x99, 0xea, 0x61, 0x99, 0xba, 0xb0, 0x9a,
	0xe2, 0x3c, 0x05, 0x6b, 0x47, 0x94, 0xdc, 0x06, 0x73, 0x85, 0x85, 0x9a, 0xe6, 0x06, 0x2b, 0x4b,
	0x72, 0x07, 0xfa, 0xca, 0xc8, 0x36, 0xd4, 0xb3, 0xea, 0xe6, 0x99, 0xf1, 0xa4, 0xe7, 0x32, 0xb0,
	0x54, 0x6b, 0x52, 0xf0, 0x54, 0x22, 0x79, 0x05, 0x57, 0x99, 0xae, 0x75, 0x1a, 0xc3, 0x63, 0x8d,
	0xec, 0x50, 0x59, 0x43, 0x74, 0x7f, 0x19, 0x70, 0x93, 0xa1, 0x58, 0x07, 0x21, 0xfe, 0x9f, 0x94,
	0xdf, 0xed, 0xa5, 0xec, 0x77, 0x71, 0xdb, 0xae, 0x07, 0x93, 0x9e, 0x83, 0xb5, 0xc2, 0xe2, 0x4b,
	0x3b, 0xed, 0xc7, 0xc7, 0xd3, 0x6e, 0x89, 0xbe, 0xc7, 0x42, 0xe7, 0x0f, 0xab, 0xba, 0xfc, 0xa7,
	0x15, 0x2c, 0xe0, 0x56, 0x63, 0xa1, 0xd7, 0xf0, 0xf6, 0xaf, 0x35, 0x8c, 0xcf, 0xea, 0x70, 0x7f,
	0x15, 0xfe, 0x4f, 0x03, 0xfa, 0x9f, 0x4a, 0x28, 0xf9, 0x00, 0xe6, 0x2c, 0x8a, 0x88, 0x7b, 0xfa,
	0x03, 0x75, 0xae, 0x8f, 0x62, 0x74, 0x8b, 0x0b, 0xb8, 0xd4, 0xb6, 0xe4, 0xc1, 0x79, 0xeb, 0x70,
	0x86, 0x27, 0x71, 0x8d, 0x76, 0x7f, 0x5e, 0xfe, 0xb1, 0xe4, 0x5e, 0x17, 0x43, 0xbd, 0xae, 0x75,
	0xc7, 0xdd, 0xba, 0x65, 0x3c, 0x1a, 0x5b, 0x69, 0xbf, 0x7c, 0xb1, 0x98, 0x86, 0x3c, 0xc2, 0x9a,
	0x12, 0xf2, 0x6f, 0x9e, 0x3a, 0x17, 0x62, 0x4c, 0x1f, 0xc6, 0xdc, 0x3b, 0x7c, 0x80, 0x3c, 0x17,
	0x89, 0x28, 0x4b, 0xb1, 0x5c, 0x5e, 0x28, 0xf4, 0xa3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x59,
	0x4f, 0x12, 0x1a, 0xf8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreClient interface {
	// Add adds a list of records to a collection using a store pipeline.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Replace records in the collection using a store pipeline.
	Replace(ctx context.Context, in *ReplaceRequest, opts ...grpc.CallOption) (*ReplaceResponse, error)
	Usage(ctx context.Context, in *UsageRequest, opts ...grpc.CallOption) (*RecordUsageResponse, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/sajari.api.pipeline.v1.Store/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Replace(ctx context.Context, in *ReplaceRequest, opts ...grpc.CallOption) (*ReplaceResponse, error) {
	out := new(ReplaceResponse)
	err := c.cc.Invoke(ctx, "/sajari.api.pipeline.v1.Store/Replace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Usage(ctx context.Context, in *UsageRequest, opts ...grpc.CallOption) (*RecordUsageResponse, error) {
	out := new(RecordUsageResponse)
	err := c.cc.Invoke(ctx, "/sajari.api.pipeline.v1.Store/Usage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	// Add adds a list of records to a collection using a store pipeline.
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Replace records in the collection using a store pipeline.
	Replace(context.Context, *ReplaceRequest) (*ReplaceResponse, error)
	Usage(context.Context, *UsageRequest) (*RecordUsageResponse, error)
}

// UnimplementedStoreServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (*UnimplementedStoreServer) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedStoreServer) Replace(ctx context.Context, req *ReplaceRequest) (*ReplaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
}
func (*UnimplementedStoreServer) Usage(ctx context.Context, req *UsageRequest) (*RecordUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Usage not implemented")
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.api.pipeline.v1.Store/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Replace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Replace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.api.pipeline.v1.Store/Replace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Replace(ctx, req.(*ReplaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Usage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Usage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.api.pipeline.v1.Store/Usage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Usage(ctx, req.(*UsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sajari.api.pipeline.v1.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Store_Add_Handler,
		},
		{
			MethodName: "Replace",
			Handler:    _Store_Replace_Handler,
		},
		{
			MethodName: "Usage",
			Handler:    _Store_Usage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sajari/api/pipeline/v1/record.proto",
}
