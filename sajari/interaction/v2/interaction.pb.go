// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sajari/interaction/v2/interaction.proto

package interactionpb

import (
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

// ConsumeTokenRequest is used in Consume calls.
type ConsumeTokenRequest struct {
	// Token to consume.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Identifier to distinguish different use of tokens.
	Identifier string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Weight assigned to the interaction.
	Weight int32 `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	// Data to be recorded with the interaction.
	Data                 map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ConsumeTokenRequest) Reset()         { *m = ConsumeTokenRequest{} }
func (m *ConsumeTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumeTokenRequest) ProtoMessage()    {}
func (*ConsumeTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82365a02ddcaca4e, []int{0}
}

func (m *ConsumeTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeTokenRequest.Unmarshal(m, b)
}
func (m *ConsumeTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeTokenRequest.Marshal(b, m, deterministic)
}
func (m *ConsumeTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeTokenRequest.Merge(m, src)
}
func (m *ConsumeTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumeTokenRequest.Size(m)
}
func (m *ConsumeTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeTokenRequest proto.InternalMessageInfo

func (m *ConsumeTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ConsumeTokenRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *ConsumeTokenRequest) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *ConsumeTokenRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type ConsumeTokenResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeTokenResponse) Reset()         { *m = ConsumeTokenResponse{} }
func (m *ConsumeTokenResponse) String() string { return proto.CompactTextString(m) }
func (*ConsumeTokenResponse) ProtoMessage()    {}
func (*ConsumeTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82365a02ddcaca4e, []int{1}
}

func (m *ConsumeTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeTokenResponse.Unmarshal(m, b)
}
func (m *ConsumeTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeTokenResponse.Marshal(b, m, deterministic)
}
func (m *ConsumeTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeTokenResponse.Merge(m, src)
}
func (m *ConsumeTokenResponse) XXX_Size() int {
	return xxx_messageInfo_ConsumeTokenResponse.Size(m)
}
func (m *ConsumeTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeTokenResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConsumeTokenRequest)(nil), "sajari.interaction.v2.ConsumeTokenRequest")
	proto.RegisterMapType((map[string]string)(nil), "sajari.interaction.v2.ConsumeTokenRequest.DataEntry")
	proto.RegisterType((*ConsumeTokenResponse)(nil), "sajari.interaction.v2.ConsumeTokenResponse")
}

func init() {
	proto.RegisterFile("sajari/interaction/v2/interaction.proto", fileDescriptor_82365a02ddcaca4e)
}

var fileDescriptor_82365a02ddcaca4e = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4b, 0x4b, 0xc3, 0x40,
	0x10, 0xc7, 0x49, 0x5f, 0xd0, 0xa9, 0x07, 0x59, 0x6b, 0x09, 0x3d, 0x48, 0xe8, 0xc5, 0xa0, 0x98,
	0x40, 0x14, 0x2c, 0xf5, 0xe6, 0x03, 0xf4, 0x1a, 0x3c, 0x79, 0xdb, 0x26, 0x63, 0x5c, 0x6b, 0x77,
	0xe3, 0xee, 0x24, 0xd2, 0x2f, 0xec, 0xe7, 0x90, 0x3c, 0xd4, 0x55, 0x7a, 0xe8, 0x6d, 0x7f, 0x33,
	0xff, 0x9d, 0xf9, 0xc1, 0xc0, 0xb1, 0xe1, 0xaf, 0x5c, 0x8b, 0x50, 0x48, 0x42, 0xcd, 0x13, 0x12,
	0x4a, 0x86, 0x65, 0x64, 0x63, 0x90, 0x6b, 0x45, 0x8a, 0x1d, 0x36, 0xc1, 0xc0, 0xee, 0x94, 0xd1,
	0xec, 0xd3, 0x81, 0x83, 0x1b, 0x25, 0x4d, 0xb1, 0xc6, 0x47, 0xb5, 0x42, 0x19, 0xe3, 0x7b, 0x81,
	0x86, 0xd8, 0x18, 0xfa, 0x54, 0xb1, 0xeb, 0x78, 0x8e, 0x3f, 0x8c, 0x1b, 0x60, 0x47, 0x00, 0x22,
	0x45, 0x49, 0xe2, 0x59, 0xa0, 0x76, 0x3b, 0x75, 0xcb, 0xaa, 0xb0, 0x09, 0x0c, 0x3e, 0x50, 0x64,
	0x2f, 0xe4, 0x76, 0x3d, 0xc7, 0xef, 0xc7, 0x2d, 0xb1, 0x7b, 0xe8, 0xa5, 0x9c, 0xb8, 0xdb, 0xf3,
	0xba, 0xfe, 0x28, 0xba, 0x08, 0xb6, 0xba, 0x04, 0x5b, 0x3c, 0x82, 0x5b, 0x4e, 0xfc, 0x4e, 0x92,
	0xde, 0xc4, 0xf5, 0x84, 0xe9, 0x25, 0x0c, 0x7f, 0x4a, 0x6c, 0x1f, 0xba, 0x2b, 0xdc, 0xb4, 0x8a,
	0xd5, 0xb3, 0xd2, 0x2e, 0xf9, 0x5b, 0x81, 0xad, 0x5b, 0x03, 0x8b, 0xce, 0xdc, 0x99, 0x4d, 0x60,
	0xfc, 0x77, 0xbe, 0xc9, 0x95, 0x34, 0x18, 0x95, 0x30, 0x7a, 0xf8, 0xd5, 0x60, 0x19, 0xec, 0xd9,
	0x31, 0x76, 0xb2, 0xbb, 0xeb, 0xf4, 0x74, 0xa7, 0x6c, 0xb3, 0xf7, 0x7a, 0xf1, 0x34, 0x4f, 0x54,
	0x8a, 0xdf, 0x5f, 0x12, 0xb5, 0x0e, 0xeb, 0x43, 0x65, 0x28, 0xcf, 0x32, 0xf5, 0xef, 0x9e, 0x57,
	0x16, 0xe6, 0xcb, 0xe5, 0xa0, 0x4e, 0x9e, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x85, 0xa4, 0x12,
	0xd1, 0xfd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InteractionClient is the client API for Interaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InteractionClient interface {
	// Consume accepts and processes tokens from interactions.
	ConsumeToken(ctx context.Context, in *ConsumeTokenRequest, opts ...grpc.CallOption) (*ConsumeTokenResponse, error)
}

type interactionClient struct {
	cc *grpc.ClientConn
}

func NewInteractionClient(cc *grpc.ClientConn) InteractionClient {
	return &interactionClient{cc}
}

func (c *interactionClient) ConsumeToken(ctx context.Context, in *ConsumeTokenRequest, opts ...grpc.CallOption) (*ConsumeTokenResponse, error) {
	out := new(ConsumeTokenResponse)
	err := c.cc.Invoke(ctx, "/sajari.interaction.v2.Interaction/ConsumeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionServer is the server API for Interaction service.
type InteractionServer interface {
	// Consume accepts and processes tokens from interactions.
	ConsumeToken(context.Context, *ConsumeTokenRequest) (*ConsumeTokenResponse, error)
}

func RegisterInteractionServer(s *grpc.Server, srv InteractionServer) {
	s.RegisterService(&_Interaction_serviceDesc, srv)
}

func _Interaction_ConsumeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServer).ConsumeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.interaction.v2.Interaction/ConsumeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServer).ConsumeToken(ctx, req.(*ConsumeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Interaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sajari.interaction.v2.Interaction",
	HandlerType: (*InteractionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConsumeToken",
			Handler:    _Interaction_ConsumeToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sajari/interaction/v2/interaction.proto",
}
