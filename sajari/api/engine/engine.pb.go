// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sajari/api/engine/engine.proto

package enginepb

import (
	rpc "code.sajari.com/protogen-go/sajari/rpc"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Collection as defined when creating/loading/deleting/listing.
type Collection struct {
	// Name of the collection.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba45a2afb52d7ebe, []int{0}
}

func (m *Collection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collection.Unmarshal(m, b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return xxx_messageInfo_Collection.Size(m)
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Collection)(nil), "sajari.api.engine.Collection")
}

func init() { proto.RegisterFile("sajari/api/engine/engine.proto", fileDescriptor_ba45a2afb52d7ebe) }

var fileDescriptor_ba45a2afb52d7ebe = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4e, 0xcc, 0x4a,
	0x2c, 0xca, 0xd4, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0x85, 0x52, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x82, 0x10, 0x79, 0xbd, 0xc4, 0x82, 0x4c, 0x3d, 0x88, 0x84,
	0x94, 0x18, 0x54, 0x4b, 0x51, 0x41, 0xb2, 0x7e, 0x6a, 0x6e, 0x41, 0x49, 0x25, 0x44, 0xa9, 0x92,
	0x02, 0x17, 0x97, 0x73, 0x7e, 0x4e, 0x4e, 0x6a, 0x72, 0x49, 0x66, 0x7e, 0x9e, 0x90, 0x10, 0x17,
	0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x6d, 0x34, 0x85,
	0x91, 0x8b, 0xcd, 0x15, 0x6c, 0x88, 0x90, 0x0b, 0x97, 0x80, 0x73, 0x51, 0x6a, 0x62, 0x49, 0x2a,
	0x92, 0x16, 0x59, 0x3d, 0x0c, 0xcb, 0xf4, 0x10, 0xd2, 0x52, 0x82, 0x30, 0xe9, 0xa2, 0x82, 0x64,
	0x3d, 0x57, 0x90, 0xc5, 0x20, 0x53, 0x5c, 0x52, 0x73, 0x52, 0x29, 0x33, 0xc5, 0xc9, 0x22, 0xca,
	0x2c, 0x39, 0x3f, 0x25, 0x15, 0x26, 0x91, 0x9c, 0x9f, 0xab, 0x0f, 0xf6, 0x51, 0x7a, 0x6a, 0x9e,
	0x6e, 0x7a, 0xbe, 0x3e, 0x46, 0x08, 0x59, 0x43, 0xa8, 0x82, 0xa4, 0x24, 0x36, 0xb0, 0x3a, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x86, 0xed, 0x8e, 0x46, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EngineClient is the client API for Engine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EngineClient interface {
	// Create a collection.
	CreateCollection(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*rpc.Empty, error)
	// Deletes all resources associated with this collection. Collections must first
	// be unloaded before they can be deleted.
	DeleteCollection(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*rpc.Empty, error)
}

type engineClient struct {
	cc *grpc.ClientConn
}

func NewEngineClient(cc *grpc.ClientConn) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) CreateCollection(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*rpc.Empty, error) {
	out := new(rpc.Empty)
	err := c.cc.Invoke(ctx, "/sajari.api.engine.Engine/CreateCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) DeleteCollection(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*rpc.Empty, error) {
	out := new(rpc.Empty)
	err := c.cc.Invoke(ctx, "/sajari.api.engine.Engine/DeleteCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineServer is the server API for Engine service.
type EngineServer interface {
	// Create a collection.
	CreateCollection(context.Context, *Collection) (*rpc.Empty, error)
	// Deletes all resources associated with this collection. Collections must first
	// be unloaded before they can be deleted.
	DeleteCollection(context.Context, *Collection) (*rpc.Empty, error)
}

func RegisterEngineServer(s *grpc.Server, srv EngineServer) {
	s.RegisterService(&_Engine_serviceDesc, srv)
}

func _Engine_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Collection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.api.engine.Engine/CreateCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).CreateCollection(ctx, req.(*Collection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_DeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Collection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).DeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.api.engine.Engine/DeleteCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).DeleteCollection(ctx, req.(*Collection))
	}
	return interceptor(ctx, in, info, handler)
}

var _Engine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sajari.api.engine.Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollection",
			Handler:    _Engine_CreateCollection_Handler,
		},
		{
			MethodName: "DeleteCollection",
			Handler:    _Engine_DeleteCollection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sajari/api/engine/engine.proto",
}
