// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sajari/api/pipeline/v1/query.proto

package piplinepb

import (
	v1 "code.sajari.com/protogen-go/sajari/api/query/v1"
	engine "code.sajari.com/protogen-go/sajari/engine"
	v11 "code.sajari.com/protogen-go/sajari/engine/query/v1"
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
type SearchRequest struct {
	// Pipeline to run.
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	// Values is a mapping of key -> value which should be substituted
	// into the algorithm inputs.
	Values map[string]string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Tracking is the tracking configuration.
	Tracking             *v1.SearchRequest_Tracking `protobuf:"bytes,3,opt,name=tracking,proto3" json:"tracking,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e834e5a8b9e4b9, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *SearchRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *SearchRequest) GetTracking() *v1.SearchRequest_Tracking {
	if m != nil {
		return m.Tracking
	}
	return nil
}

// SearchResponse is a response to a Search call.
type SearchResponse struct {
	// Input values with updates/modifications by querying system.
	Values map[string]string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// SearchResponse from the engine request.
	SearchResponse *v11.SearchResponse `protobuf:"bytes,2,opt,name=search_response,json=searchResponse,proto3" json:"search_response,omitempty"`
	// Tokens which correspond to the result records.
	Tokens               []*v1.Token `protobuf:"bytes,3,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e834e5a8b9e4b9, []int{1}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *SearchResponse) GetSearchResponse() *v11.SearchResponse {
	if m != nil {
		return m.SearchResponse
	}
	return nil
}

func (m *SearchResponse) GetTokens() []*v1.Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

// EvaluateRequest is a request to perform a Evaluate using a pipeline.
type EvaluateRequest struct {
	// Pipeline to run.
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	// Values is a mapping of key -> value which should be substituted
	// into the algorithm inputs.
	Values map[string]string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Key of the record to run search against.
	Key                  *engine.Key `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EvaluateRequest) Reset()         { *m = EvaluateRequest{} }
func (m *EvaluateRequest) String() string { return proto.CompactTextString(m) }
func (*EvaluateRequest) ProtoMessage()    {}
func (*EvaluateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e834e5a8b9e4b9, []int{2}
}

func (m *EvaluateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateRequest.Unmarshal(m, b)
}
func (m *EvaluateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateRequest.Marshal(b, m, deterministic)
}
func (m *EvaluateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateRequest.Merge(m, src)
}
func (m *EvaluateRequest) XXX_Size() int {
	return xxx_messageInfo_EvaluateRequest.Size(m)
}
func (m *EvaluateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateRequest proto.InternalMessageInfo

func (m *EvaluateRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *EvaluateRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *EvaluateRequest) GetKey() *engine.Key {
	if m != nil {
		return m.Key
	}
	return nil
}

// EvaluateResponse is a response to a Evaluate call.
type EvaluateResponse struct {
	// Input values with updates/modifications by querying system.
	Values map[string]string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// EvaluateResponse from the engine request.
	SearchResponse       *v11.SearchResponse `protobuf:"bytes,2,opt,name=search_response,json=searchResponse,proto3" json:"search_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EvaluateResponse) Reset()         { *m = EvaluateResponse{} }
func (m *EvaluateResponse) String() string { return proto.CompactTextString(m) }
func (*EvaluateResponse) ProtoMessage()    {}
func (*EvaluateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e834e5a8b9e4b9, []int{3}
}

func (m *EvaluateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateResponse.Unmarshal(m, b)
}
func (m *EvaluateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateResponse.Marshal(b, m, deterministic)
}
func (m *EvaluateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateResponse.Merge(m, src)
}
func (m *EvaluateResponse) XXX_Size() int {
	return xxx_messageInfo_EvaluateResponse.Size(m)
}
func (m *EvaluateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateResponse proto.InternalMessageInfo

func (m *EvaluateResponse) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *EvaluateResponse) GetSearchResponse() *v11.SearchResponse {
	if m != nil {
		return m.SearchResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "sajari.api.pipeline.v1.SearchRequest")
	proto.RegisterMapType((map[string]string)(nil), "sajari.api.pipeline.v1.SearchRequest.ValuesEntry")
	proto.RegisterType((*SearchResponse)(nil), "sajari.api.pipeline.v1.SearchResponse")
	proto.RegisterMapType((map[string]string)(nil), "sajari.api.pipeline.v1.SearchResponse.ValuesEntry")
	proto.RegisterType((*EvaluateRequest)(nil), "sajari.api.pipeline.v1.EvaluateRequest")
	proto.RegisterMapType((map[string]string)(nil), "sajari.api.pipeline.v1.EvaluateRequest.ValuesEntry")
	proto.RegisterType((*EvaluateResponse)(nil), "sajari.api.pipeline.v1.EvaluateResponse")
	proto.RegisterMapType((map[string]string)(nil), "sajari.api.pipeline.v1.EvaluateResponse.ValuesEntry")
}

func init() {
	proto.RegisterFile("sajari/api/pipeline/v1/query.proto", fileDescriptor_b1e834e5a8b9e4b9)
}

var fileDescriptor_b1e834e5a8b9e4b9 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x25, 0x59, 0x76, 0x59, 0xef, 0x62, 0x5b, 0x06, 0xd1, 0x90, 0x17, 0x97, 0xa5, 0xad, 0x8b,
	0x62, 0xc2, 0xa6, 0x3e, 0xf8, 0x51, 0x7c, 0x10, 0x8a, 0x68, 0x05, 0x35, 0x56, 0x45, 0x41, 0x64,
	0xba, 0x5e, 0x62, 0xdc, 0x9a, 0x4c, 0x33, 0xd9, 0x85, 0xfc, 0x15, 0xf1, 0x27, 0xfa, 0x17, 0x04,
	0x99, 0xaf, 0xb8, 0x49, 0x13, 0x1a, 0x58, 0xf1, 0x6d, 0x92, 0x39, 0xe7, 0xdc, 0x7b, 0xce, 0xdc,
	0x19, 0x98, 0x70, 0xfa, 0x8d, 0x66, 0xb1, 0x4f, 0x59, 0xec, 0xb3, 0x98, 0xe1, 0x59, 0x9c, 0xa0,
	0xbf, 0x9a, 0xf9, 0xe7, 0x4b, 0xcc, 0x0a, 0x8f, 0x65, 0x69, 0x9e, 0x92, 0xeb, 0x0a, 0xe3, 0x51,
	0x16, 0x7b, 0x06, 0xe3, 0xad, 0x66, 0xee, 0xcd, 0x35, 0xae, 0xc4, 0xd7, 0x88, 0xee, 0x0d, 0x0d,
	0xc0, 0x24, 0x12, 0xb2, 0x0b, 0x34, 0x1b, 0x93, 0xea, 0x46, 0x23, 0x79, 0xaf, 0xa5, 0xb3, 0xb2,
	0x83, 0xaa, 0x54, 0x1d, 0xb6, 0xe4, 0x34, 0xd2, 0x98, 0xc9, 0x4f, 0x1b, 0xae, 0xbe, 0x41, 0x9a,
	0xcd, 0xbf, 0x86, 0x78, 0xbe, 0x44, 0x9e, 0x93, 0x43, 0x18, 0x1a, 0xb0, 0x63, 0x8d, 0xad, 0xe9,
	0x28, 0x18, 0x7b, 0xcd, 0x2e, 0xbd, 0x57, 0x7a, 0x1d, 0x96, 0x0c, 0xf2, 0x0c, 0x06, 0x2b, 0x7a,
	0xb6, 0x44, 0xee, 0xd8, 0xe3, 0xde, 0x74, 0x14, 0xcc, 0xda, 0xb8, 0x95, 0xa2, 0xde, 0x3b, 0xc9,
	0x39, 0x4a, 0xf2, 0xac, 0x08, 0xb5, 0x00, 0x79, 0x0a, 0xc3, 0x3c, 0xa3, 0xf3, 0x45, 0x9c, 0x44,
	0x4e, 0x4f, 0x36, 0x72, 0x67, 0x5d, 0x4c, 0x05, 0x72, 0x41, 0xe9, 0x44, 0x53, 0xc2, 0x92, 0xec,
	0x3e, 0x80, 0xd1, 0x9a, 0x3e, 0xd9, 0x81, 0xde, 0x02, 0x0b, 0xe9, 0xed, 0x4a, 0x28, 0x96, 0xe4,
	0x1a, 0xf4, 0x65, 0x4d, 0xc7, 0x96, 0xff, 0xd4, 0xc7, 0x43, 0xfb, 0xbe, 0x25, 0xe2, 0xd9, 0x32,
	0xfa, 0x9c, 0xa5, 0x09, 0x47, 0xf2, 0xbc, 0x74, 0x68, 0x49, 0x87, 0xc1, 0x65, 0x0e, 0x15, 0xaf,
	0xd1, 0xe2, 0x4b, 0xd8, 0xe6, 0x12, 0xf5, 0x39, 0xd3, 0x30, 0xd9, 0xc2, 0x28, 0xd8, 0x37, 0xa2,
	0x6a, 0x0c, 0x2e, 0x9a, 0x55, 0xe8, 0x70, 0x8b, 0x57, 0x9b, 0x0b, 0x60, 0x90, 0xa7, 0x0b, 0x4c,
	0xb8, 0xd3, 0x93, 0xcd, 0xb9, 0x8d, 0x89, 0x9d, 0x08, 0x48, 0xa8, 0x91, 0x9b, 0xc4, 0xf3, 0xdb,
	0x82, 0xed, 0x23, 0xf1, 0x49, 0x73, 0xfc, 0x37, 0xf3, 0x73, 0x5c, 0x9b, 0x9f, 0x83, 0x36, 0x6e,
	0xad, 0x6c, 0x63, 0xbc, 0xbb, 0xca, 0x8a, 0x1a, 0x1e, 0x52, 0x8b, 0xf4, 0x18, 0x0b, 0x69, 0x6f,
	0x13, 0xff, 0xbf, 0x2c, 0xd8, 0xf9, 0xdb, 0x88, 0x3e, 0x83, 0x17, 0xb5, 0x01, 0xb9, 0x77, 0xb9,
	0x85, 0xff, 0x38, 0x22, 0x1b, 0xd8, 0x0d, 0x7e, 0xd8, 0xd0, 0x7f, 0x2d, 0xca, 0x90, 0xf7, 0x30,
	0x50, 0x65, 0xc8, 0x5e, 0xa7, 0x0b, 0xee, 0xee, 0x77, 0xbb, 0x25, 0xe4, 0x13, 0x0c, 0x4d, 0x2c,
	0xe4, 0x56, 0xc7, 0xb3, 0x77, 0xa7, 0x5d, 0x13, 0x26, 0x1f, 0xa0, 0xff, 0x56, 0xbc, 0x7e, 0x64,
	0xb7, 0x8d, 0x22, 0xb7, 0x8d, 0xf0, 0xed, 0x36, 0x94, 0x4c, 0x41, 0x43, 0x95, 0xf4, 0x93, 0xc7,
	0x1f, 0x0f, 0xe7, 0xe9, 0x17, 0x34, 0x8c, 0x79, 0xfa, 0xdd, 0x97, 0x4f, 0x6c, 0x84, 0xc9, 0xdd,
	0x28, 0xf5, 0x9b, 0xdf, 0xe2, 0x47, 0x2c, 0x66, 0x62, 0xc9, 0x4e, 0x4f, 0x07, 0x12, 0x7d, 0xf0,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x87, 0xb0, 0x3e, 0x08, 0x77, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResponse, error)
	Usage(ctx context.Context, in *UsageRequest, opts ...grpc.CallOption) (*QueryUsageResponse, error)
}

type queryClient struct {
	cc *grpc.ClientConn
}

func NewQueryClient(cc *grpc.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/sajari.api.pipeline.v1.Query/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResponse, error) {
	out := new(EvaluateResponse)
	err := c.cc.Invoke(ctx, "/sajari.api.pipeline.v1.Query/Evaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Usage(ctx context.Context, in *UsageRequest, opts ...grpc.CallOption) (*QueryUsageResponse, error) {
	out := new(QueryUsageResponse)
	err := c.cc.Invoke(ctx, "/sajari.api.pipeline.v1.Query/Usage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Evaluate(context.Context, *EvaluateRequest) (*EvaluateResponse, error)
	Usage(context.Context, *UsageRequest) (*QueryUsageResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedQueryServer) Evaluate(ctx context.Context, req *EvaluateRequest) (*EvaluateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evaluate not implemented")
}
func (*UnimplementedQueryServer) Usage(ctx context.Context, req *UsageRequest) (*QueryUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Usage not implemented")
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.api.pipeline.v1.Query/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.api.pipeline.v1.Query/Evaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Evaluate(ctx, req.(*EvaluateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Usage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Usage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.api.pipeline.v1.Query/Usage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Usage(ctx, req.(*UsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sajari.api.pipeline.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Query_Search_Handler,
		},
		{
			MethodName: "Evaluate",
			Handler:    _Query_Evaluate_Handler,
		},
		{
			MethodName: "Usage",
			Handler:    _Query_Usage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sajari/api/pipeline/v1/query.proto",
}
