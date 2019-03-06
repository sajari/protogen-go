// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sajari/pipeline/v2/record.proto

package pipelinepb

import (
	v2 "code.sajari.com/protogen-go/sajari/engine/v2"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
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

type CreateRecordRequest struct {
	// Pipeline to run.
	Pipeline *Identifier `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	// Values is a mapping of key -> value used as the pipeline parameters.
	Values *_struct.Struct `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	// Record to create.
	Record               *v2.Record `protobuf:"bytes,3,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateRecordRequest) Reset()         { *m = CreateRecordRequest{} }
func (m *CreateRecordRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRecordRequest) ProtoMessage()    {}
func (*CreateRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e624f97aa72be4d7, []int{0}
}

func (m *CreateRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRecordRequest.Unmarshal(m, b)
}
func (m *CreateRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRecordRequest.Marshal(b, m, deterministic)
}
func (m *CreateRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRecordRequest.Merge(m, src)
}
func (m *CreateRecordRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRecordRequest.Size(m)
}
func (m *CreateRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRecordRequest proto.InternalMessageInfo

func (m *CreateRecordRequest) GetPipeline() *Identifier {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *CreateRecordRequest) GetValues() *_struct.Struct {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *CreateRecordRequest) GetRecord() *v2.Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type CreateRecordResponse struct {
	// Pipeline to run.
	Pipeline *Identifier `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	// Values is a mapping of key -> value pairs output from the pipeline
	// parameters.
	Values *_struct.Struct `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	// Key of the created record.
	Key                  *v2.Key  `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRecordResponse) Reset()         { *m = CreateRecordResponse{} }
func (m *CreateRecordResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRecordResponse) ProtoMessage()    {}
func (*CreateRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e624f97aa72be4d7, []int{1}
}

func (m *CreateRecordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRecordResponse.Unmarshal(m, b)
}
func (m *CreateRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRecordResponse.Marshal(b, m, deterministic)
}
func (m *CreateRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRecordResponse.Merge(m, src)
}
func (m *CreateRecordResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRecordResponse.Size(m)
}
func (m *CreateRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRecordResponse proto.InternalMessageInfo

func (m *CreateRecordResponse) GetPipeline() *Identifier {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *CreateRecordResponse) GetValues() *_struct.Struct {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *CreateRecordResponse) GetKey() *v2.Key {
	if m != nil {
		return m.Key
	}
	return nil
}

type ReplaceRecordRequest struct {
	// Pipeline to run.
	Pipeline *Identifier `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	// Values is a mapping of key -> value used as the pipeline parameters.
	Values *_struct.Struct `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	// Key to identify the record to replace. If the key does not exist, then
	// a record will be created instead of replaced.
	Key *v2.Key `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// Record data to replace with existing.
	Record               *v2.Record `protobuf:"bytes,4,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReplaceRecordRequest) Reset()         { *m = ReplaceRecordRequest{} }
func (m *ReplaceRecordRequest) String() string { return proto.CompactTextString(m) }
func (*ReplaceRecordRequest) ProtoMessage()    {}
func (*ReplaceRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e624f97aa72be4d7, []int{2}
}

func (m *ReplaceRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceRecordRequest.Unmarshal(m, b)
}
func (m *ReplaceRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceRecordRequest.Marshal(b, m, deterministic)
}
func (m *ReplaceRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceRecordRequest.Merge(m, src)
}
func (m *ReplaceRecordRequest) XXX_Size() int {
	return xxx_messageInfo_ReplaceRecordRequest.Size(m)
}
func (m *ReplaceRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceRecordRequest proto.InternalMessageInfo

func (m *ReplaceRecordRequest) GetPipeline() *Identifier {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *ReplaceRecordRequest) GetValues() *_struct.Struct {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ReplaceRecordRequest) GetKey() *v2.Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ReplaceRecordRequest) GetRecord() *v2.Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type ReplaceRecordResponse struct {
	// Pipeline to run.
	Pipeline *Identifier `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	// Values set by the pipeline.
	Values *_struct.Struct `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	// Key of the created record, or empty if an existing record was replaced.
	Key                  *v2.Key  `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplaceRecordResponse) Reset()         { *m = ReplaceRecordResponse{} }
func (m *ReplaceRecordResponse) String() string { return proto.CompactTextString(m) }
func (*ReplaceRecordResponse) ProtoMessage()    {}
func (*ReplaceRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e624f97aa72be4d7, []int{3}
}

func (m *ReplaceRecordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceRecordResponse.Unmarshal(m, b)
}
func (m *ReplaceRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceRecordResponse.Marshal(b, m, deterministic)
}
func (m *ReplaceRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceRecordResponse.Merge(m, src)
}
func (m *ReplaceRecordResponse) XXX_Size() int {
	return xxx_messageInfo_ReplaceRecordResponse.Size(m)
}
func (m *ReplaceRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceRecordResponse proto.InternalMessageInfo

func (m *ReplaceRecordResponse) GetPipeline() *Identifier {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *ReplaceRecordResponse) GetValues() *_struct.Struct {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ReplaceRecordResponse) GetKey() *v2.Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRecordRequest)(nil), "sajari.pipeline.v2.CreateRecordRequest")
	proto.RegisterType((*CreateRecordResponse)(nil), "sajari.pipeline.v2.CreateRecordResponse")
	proto.RegisterType((*ReplaceRecordRequest)(nil), "sajari.pipeline.v2.ReplaceRecordRequest")
	proto.RegisterType((*ReplaceRecordResponse)(nil), "sajari.pipeline.v2.ReplaceRecordResponse")
}

func init() { proto.RegisterFile("sajari/pipeline/v2/record.proto", fileDescriptor_e624f97aa72be4d7) }

var fileDescriptor_e624f97aa72be4d7 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xbf, 0x4e, 0xfb, 0x30,
	0x10, 0xc7, 0xe5, 0x5f, 0x7f, 0x54, 0xe8, 0x80, 0xc5, 0xb4, 0x22, 0x8a, 0xf8, 0xdb, 0xa5, 0x65,
	0xc0, 0x46, 0x61, 0x82, 0x6e, 0x30, 0x21, 0xb6, 0x74, 0x63, 0x4b, 0x93, 0x6b, 0x14, 0x08, 0x71,
	0x70, 0x9c, 0x4a, 0x7d, 0x27, 0xd8, 0x78, 0x10, 0x9e, 0x80, 0x67, 0x41, 0x8d, 0xe3, 0xaa, 0xa5,
	0x96, 0xc8, 0x04, 0x6c, 0x55, 0xf5, 0xf9, 0xfa, 0x3e, 0x77, 0xe7, 0x18, 0x8e, 0x8a, 0xe0, 0x21,
	0x90, 0x09, 0xcf, 0x93, 0x1c, 0xd3, 0x24, 0x43, 0x3e, 0xf5, 0xb8, 0xc4, 0x50, 0xc8, 0x88, 0xe5,
	0x52, 0x28, 0x41, 0xa9, 0x06, 0x98, 0x01, 0xd8, 0xd4, 0x73, 0xf7, 0x63, 0x21, 0xe2, 0x14, 0x79,
	0x45, 0x8c, 0xcb, 0x09, 0x2f, 0x94, 0x2c, 0x43, 0xa5, 0x13, 0xee, 0x89, 0xe5, 0xc8, 0x45, 0x5a,
	0x23, 0x07, 0x35, 0x82, 0x59, 0x6c, 0xa9, 0xd9, 0x7b, 0x23, 0xb0, 0x7b, 0x23, 0x31, 0x50, 0xe8,
	0x57, 0x7f, 0xfb, 0xf8, 0x5c, 0x62, 0xa1, 0xe8, 0x15, 0x6c, 0x9a, 0x83, 0x1c, 0x72, 0x4c, 0x06,
	0x5b, 0xde, 0x21, 0x5b, 0xd7, 0x63, 0xb7, 0x11, 0x66, 0x2a, 0x99, 0x24, 0x28, 0xfd, 0x05, 0x4f,
	0x39, 0xb4, 0xa7, 0x41, 0x5a, 0x62, 0xe1, 0xfc, 0xab, 0x92, 0x7b, 0x4c, 0x37, 0xc1, 0x4c, 0x13,
	0x6c, 0x54, 0x35, 0xe1, 0xd7, 0x18, 0x3d, 0x87, 0xb6, 0x96, 0x72, 0x5a, 0x55, 0xc0, 0x31, 0xa5,
	0xb4, 0xf4, 0xbc, 0x50, 0x6d, 0x57, 0x73, 0xbd, 0x17, 0x02, 0x9d, 0x55, 0xed, 0x22, 0x17, 0x59,
	0x81, 0x3f, 0xeb, 0xdd, 0x87, 0xd6, 0x23, 0xce, 0x6a, 0xe9, 0xee, 0xba, 0xf4, 0x1d, 0xce, 0xfc,
	0x39, 0xd1, 0xfb, 0x20, 0xd0, 0xf1, 0x31, 0x4f, 0x83, 0xf0, 0x37, 0xc7, 0xdc, 0x54, 0x77, 0x69,
	0x1f, 0xff, 0x1b, 0xee, 0xe3, 0x95, 0x40, 0xf7, 0x4b, 0x83, 0x7f, 0x79, 0x21, 0xde, 0x3b, 0x81,
	0x8d, 0x91, 0x12, 0x12, 0x69, 0x00, 0xdb, 0xcb, 0x17, 0x89, 0xf6, 0x6d, 0x76, 0x96, 0x2f, 0xc4,
	0x1d, 0x7c, 0x0f, 0xd6, 0x23, 0x88, 0x60, 0x67, 0x65, 0x36, 0xd4, 0x1a, 0xb5, 0xdd, 0x0f, 0xf7,
	0xb4, 0x01, 0xa9, 0xab, 0x5c, 0x0f, 0xef, 0x2f, 0x43, 0x11, 0xa1, 0x09, 0x84, 0xe2, 0x49, 0x3f,
	0x1a, 0x31, 0x66, 0x67, 0xb1, 0xe0, 0xeb, 0x2f, 0xc5, 0xd0, 0xfc, 0xce, 0xc7, 0xe3, 0x76, 0x85,
	0x5e, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0x54, 0xb5, 0x13, 0x3d, 0xa4, 0x04, 0x00, 0x00,
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
	// Create a records using a record pipeline.
	CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error)
	// Replace a record using a record pipeline.
	ReplaceRecord(ctx context.Context, in *ReplaceRecordRequest, opts ...grpc.CallOption) (*ReplaceRecordResponse, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error) {
	out := new(CreateRecordResponse)
	err := c.cc.Invoke(ctx, "/sajari.pipeline.v2.Store/CreateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) ReplaceRecord(ctx context.Context, in *ReplaceRecordRequest, opts ...grpc.CallOption) (*ReplaceRecordResponse, error) {
	out := new(ReplaceRecordResponse)
	err := c.cc.Invoke(ctx, "/sajari.pipeline.v2.Store/ReplaceRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	// Create a records using a record pipeline.
	CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error)
	// Replace a record using a record pipeline.
	ReplaceRecord(context.Context, *ReplaceRecordRequest) (*ReplaceRecordResponse, error)
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.pipeline.v2.Store/CreateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).CreateRecord(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_ReplaceRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).ReplaceRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.pipeline.v2.Store/ReplaceRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).ReplaceRecord(ctx, req.(*ReplaceRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sajari.pipeline.v2.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecord",
			Handler:    _Store_CreateRecord_Handler,
		},
		{
			MethodName: "ReplaceRecord",
			Handler:    _Store_ReplaceRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sajari/pipeline/v2/record.proto",
}
