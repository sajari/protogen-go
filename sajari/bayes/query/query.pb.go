// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sajari/bayes/query/query.proto

package bayesquerypb

import (
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

// Request contains a model name and an array of string based data representing
// the content to be classified by the naive bayes model
type Request struct {
	// Model is the name of the model to be queried
	Model string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	// Data is a list of strings representing the input query to be
	// classified. Normally these would represent words from text. It is the
	// callers responsibility to stem and tokenise into an array of strings.
	Data                 []string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0df142f249badc45, []int{0}
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

func (m *Request) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Request) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

// Response returns information on the classification.
type Response struct {
	// Scores map represents each of the potential classes and their
	// associated probability (Note: only if the probability calculation does
	// not underflow)
	Scores map[string]float64 `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	// Best represents the highest probability class for the input data.
	Best string `protobuf:"bytes,2,opt,name=best,proto3" json:"best,omitempty"`
	// Unique indicates if this classification was the solo highest probability
	// (i.e. not equal with other classes)
	Unique               bool     `protobuf:"varint,3,opt,name=unique,proto3" json:"unique,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0df142f249badc45, []int{1}
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

func (m *Response) GetScores() map[string]float64 {
	if m != nil {
		return m.Scores
	}
	return nil
}

func (m *Response) GetBest() string {
	if m != nil {
		return m.Best
	}
	return ""
}

func (m *Response) GetUnique() bool {
	if m != nil {
		return m.Unique
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "sajari.bayes.query.Request")
	proto.RegisterType((*Response)(nil), "sajari.bayes.query.Response")
	proto.RegisterMapType((map[string]float64)(nil), "sajari.bayes.query.Response.ScoresEntry")
}

func init() { proto.RegisterFile("sajari/bayes/query/query.proto", fileDescriptor_0df142f249badc45) }

var fileDescriptor_0df142f249badc45 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xd9, 0xe4, 0x6b, 0xbe, 0x76, 0x7a, 0x91, 0x45, 0x24, 0x54, 0x91, 0xd0, 0x53, 0x2e,
	0x6e, 0xa0, 0xbd, 0xa8, 0x45, 0x90, 0x82, 0x27, 0x4f, 0xae, 0x37, 0x6f, 0x9b, 0x64, 0x28, 0xd5,
	0x36, 0x9b, 0x64, 0x77, 0x85, 0xfc, 0x2e, 0xff, 0xa0, 0x64, 0x36, 0x81, 0x82, 0xe2, 0x65, 0x79,
	0x86, 0x79, 0x67, 0xe6, 0x9d, 0x59, 0xb8, 0x36, 0xea, 0x5d, 0xb5, 0xfb, 0x2c, 0x57, 0x1d, 0x9a,
	0xac, 0x71, 0xd8, 0x76, 0xfe, 0x15, 0x75, 0xab, 0xad, 0xe6, 0xdc, 0xe7, 0x05, 0xe5, 0x05, 0x65,
	0x96, 0x6b, 0xf8, 0x2f, 0xb1, 0x71, 0x68, 0x2c, 0x3f, 0x87, 0xc9, 0x51, 0x97, 0x78, 0x88, 0x59,
	0xc2, 0xd2, 0x99, 0xf4, 0x01, 0xe7, 0xf0, 0xaf, 0x54, 0x56, 0xc5, 0x41, 0x12, 0xa6, 0x33, 0x49,
	0xbc, 0xfc, 0x62, 0x30, 0x95, 0x68, 0x6a, 0x5d, 0x19, 0xe4, 0x8f, 0x10, 0x99, 0x42, 0xb7, 0x68,
	0x62, 0x96, 0x84, 0xe9, 0x7c, 0x95, 0x8a, 0x9f, 0x63, 0xc4, 0xa8, 0x16, 0xaf, 0x24, 0x7d, 0xaa,
	0x6c, 0xdb, 0xc9, 0xa1, 0xae, 0x1f, 0x91, 0xa3, 0xb1, 0x71, 0x40, 0x73, 0x89, 0xf9, 0x05, 0x44,
	0xae, 0xda, 0x37, 0x0e, 0xe3, 0x30, 0x61, 0xe9, 0x54, 0x0e, 0xd1, 0xe2, 0x0e, 0xe6, 0x27, 0x2d,
	0xf8, 0x19, 0x84, 0x1f, 0xd8, 0x0d, 0x8e, 0x7b, 0xec, 0xb7, 0xf8, 0x54, 0x07, 0x87, 0xd4, 0x8d,
	0x49, 0x1f, 0xdc, 0x07, 0xb7, 0x6c, 0xf5, 0x0c, 0x93, 0x97, 0xde, 0x0c, 0xdf, 0x8e, 0x70, 0xf9,
	0xbb, 0x55, 0x3a, 0xc7, 0xe2, 0xea, 0xaf, 0x3d, 0xb6, 0x0f, 0x6f, 0x9b, 0x42, 0x97, 0x38, 0x6a,
	0x0a, 0x7d, 0xcc, 0xe8, 0xc8, 0x3b, 0xac, 0x6e, 0x76, 0x3a, 0x3b, 0xfd, 0x09, 0x5f, 0xba, 0x21,
	0x26, 0xac, 0xf3, 0x3c, 0x22, 0xf1, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x55, 0xe9, 0xa0,
	0xb3, 0x01, 0x00, 0x00,
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
	// Query takes a model name and an array of strings and returns a naive bayes
	// based classification for the request data using the model specified.
	Query(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type queryClient struct {
	cc *grpc.ClientConn
}

func NewQueryClient(cc *grpc.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Query(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sajari.bayes.query.Query/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Query takes a model name and an array of strings and returns a naive bayes
	// based classification for the request data using the model specified.
	Query(context.Context, *Request) (*Response, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Query(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.bayes.query.Query/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Query(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sajari.bayes.query.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Query_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sajari/bayes/query/query.proto",
}
