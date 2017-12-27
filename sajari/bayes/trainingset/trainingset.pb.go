// Code generated by protoc-gen-go.
// source: sajari/bayes/trainingset/trainingset.proto
// DO NOT EDIT!

/*
Package sajari_bayes_trainingset is a generated protocol buffer package.

It is generated from these files:
	sajari/bayes/trainingset/trainingset.proto

It has these top-level messages:
	UploadRequest
	UploadResponse
	CreateRequest
	InfoRequest
	InfoResponse
	AddClassRequest
*/
package sajari_bayes_trainingset

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sajari_rpc "code.sajari.com/protogen-go/sajari/rpc"

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

// UploadRequest encapsulates training data representing a document to add to the
// specified training set and class name
type UploadRequest struct {
	// Name of the training set
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Class name to add the data to
	Class string `protobuf:"bytes,2,opt,name=class" json:"class,omitempty"`
	// Data represents the individual strings representing the input query to be
	// classified. Normally these would represent words from text. It is the
	// callers responsibility to stem and tokenise into an array of strings.
	Data []string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty"`
}

func (m *UploadRequest) Reset()                    { *m = UploadRequest{} }
func (m *UploadRequest) String() string            { return proto.CompactTextString(m) }
func (*UploadRequest) ProtoMessage()               {}
func (*UploadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UploadRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UploadRequest) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *UploadRequest) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

// UploadResponse contains a unique hash for the uploaded data to prevent
// duplicate documents distorting the training set probabilities
type UploadResponse struct {
	// unique hash for the uploaded data
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *UploadResponse) Reset()                    { *m = UploadResponse{} }
func (m *UploadResponse) String() string            { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()               {}
func (*UploadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UploadResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// CreateRequest creates a new training set
type CreateRequest struct {
	// Name of the training set to create
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// InfoRequest gets information for a training set
type InfoRequest struct {
	// Name of the training set to get information
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// InfoResponse returns information for a given training set
type InfoResponse struct {
	// List of classes for the specified training set
	Classes []string `protobuf:"bytes,1,rep,name=classes" json:"classes,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InfoResponse) GetClasses() []string {
	if m != nil {
		return m.Classes
	}
	return nil
}

// AddClassRequest creates a new class for the specified training set
type AddClassRequest struct {
	// The name of the training set to add the class to
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Name of the class to add
	Class string `protobuf:"bytes,2,opt,name=class" json:"class,omitempty"`
}

func (m *AddClassRequest) Reset()                    { *m = AddClassRequest{} }
func (m *AddClassRequest) String() string            { return proto.CompactTextString(m) }
func (*AddClassRequest) ProtoMessage()               {}
func (*AddClassRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AddClassRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddClassRequest) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func init() {
	proto.RegisterType((*UploadRequest)(nil), "sajari.bayes.trainingset.UploadRequest")
	proto.RegisterType((*UploadResponse)(nil), "sajari.bayes.trainingset.UploadResponse")
	proto.RegisterType((*CreateRequest)(nil), "sajari.bayes.trainingset.CreateRequest")
	proto.RegisterType((*InfoRequest)(nil), "sajari.bayes.trainingset.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "sajari.bayes.trainingset.InfoResponse")
	proto.RegisterType((*AddClassRequest)(nil), "sajari.bayes.trainingset.AddClassRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TrainingSet service

type TrainingSetClient interface {
	// Uploads a training/testing document for the specified training set name and known class name
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
	// Creates a new training set
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*sajari_rpc.Empty, error)
	// Returns information on the specified training est
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	// Adds a new class to the training set
	AddClass(ctx context.Context, in *AddClassRequest, opts ...grpc.CallOption) (*sajari_rpc.Empty, error)
}

type trainingSetClient struct {
	cc *grpc.ClientConn
}

func NewTrainingSetClient(cc *grpc.ClientConn) TrainingSetClient {
	return &trainingSetClient{cc}
}

func (c *trainingSetClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	out := new(UploadResponse)
	err := grpc.Invoke(ctx, "/sajari.bayes.trainingset.TrainingSet/Upload", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingSetClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*sajari_rpc.Empty, error) {
	out := new(sajari_rpc.Empty)
	err := grpc.Invoke(ctx, "/sajari.bayes.trainingset.TrainingSet/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingSetClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/sajari.bayes.trainingset.TrainingSet/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingSetClient) AddClass(ctx context.Context, in *AddClassRequest, opts ...grpc.CallOption) (*sajari_rpc.Empty, error) {
	out := new(sajari_rpc.Empty)
	err := grpc.Invoke(ctx, "/sajari.bayes.trainingset.TrainingSet/AddClass", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TrainingSet service

type TrainingSetServer interface {
	// Uploads a training/testing document for the specified training set name and known class name
	Upload(context.Context, *UploadRequest) (*UploadResponse, error)
	// Creates a new training set
	Create(context.Context, *CreateRequest) (*sajari_rpc.Empty, error)
	// Returns information on the specified training est
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	// Adds a new class to the training set
	AddClass(context.Context, *AddClassRequest) (*sajari_rpc.Empty, error)
}

func RegisterTrainingSetServer(s *grpc.Server, srv TrainingSetServer) {
	s.RegisterService(&_TrainingSet_serviceDesc, srv)
}

func _TrainingSet_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingSetServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.bayes.trainingset.TrainingSet/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingSetServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingSet_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingSetServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.bayes.trainingset.TrainingSet/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingSetServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingSet_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingSetServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.bayes.trainingset.TrainingSet/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingSetServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingSet_AddClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingSetServer).AddClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.bayes.trainingset.TrainingSet/AddClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingSetServer).AddClass(ctx, req.(*AddClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrainingSet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sajari.bayes.trainingset.TrainingSet",
	HandlerType: (*TrainingSetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _TrainingSet_Upload_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TrainingSet_Create_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _TrainingSet_Info_Handler,
		},
		{
			MethodName: "AddClass",
			Handler:    _TrainingSet_AddClass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sajari/bayes/trainingset/trainingset.proto",
}

func init() { proto.RegisterFile("sajari/bayes/trainingset/trainingset.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x95, 0xb6, 0x04, 0xb8, 0x52, 0x10, 0x16, 0x42, 0x56, 0xa6, 0x12, 0xfe, 0x05, 0x06,
	0x47, 0x82, 0x91, 0x09, 0x15, 0x24, 0x18, 0x58, 0x0a, 0x9d, 0x98, 0xdc, 0xe4, 0xa0, 0x41, 0x6d,
	0x62, 0x6c, 0x33, 0xf4, 0x31, 0x79, 0x23, 0x64, 0x3b, 0x91, 0x52, 0x44, 0x52, 0xb1, 0xdd, 0x25,
	0x9f, 0x7f, 0xbe, 0xfb, 0x12, 0xb8, 0x54, 0xfc, 0x83, 0xcb, 0x2c, 0x9e, 0xf2, 0x25, 0xaa, 0x58,
	0x4b, 0x9e, 0xe5, 0x59, 0xfe, 0xae, 0x50, 0xd7, 0x6b, 0x26, 0x64, 0xa1, 0x0b, 0x42, 0x1d, 0xcb,
	0x2c, 0xcb, 0x6a, 0xef, 0x83, 0xc3, 0x32, 0x45, 0x8a, 0x24, 0xc6, 0x85, 0xd0, 0x4b, 0x77, 0x22,
	0x7c, 0x82, 0xc1, 0x44, 0xcc, 0x0b, 0x9e, 0x8e, 0xf1, 0xf3, 0x0b, 0x95, 0x26, 0x04, 0x7a, 0x39,
	0x5f, 0x20, 0xf5, 0x86, 0x5e, 0xb4, 0x3d, 0xb6, 0x35, 0x39, 0x80, 0x8d, 0x64, 0xce, 0x95, 0xa2,
	0x1d, 0xfb, 0xd0, 0x35, 0x86, 0x4c, 0xb9, 0xe6, 0xb4, 0x3b, 0xec, 0x1a, 0xd2, 0xd4, 0xe1, 0x09,
	0xec, 0x56, 0x71, 0x4a, 0x14, 0xb9, 0x42, 0x43, 0xcd, 0xb8, 0x9a, 0x55, 0x79, 0xa6, 0x0e, 0x8f,
	0x61, 0x30, 0x92, 0xc8, 0x35, 0xb6, 0x5c, 0x1a, 0x1e, 0x41, 0xff, 0x31, 0x7f, 0x2b, 0xda, 0x90,
	0x08, 0x76, 0x1c, 0x52, 0xde, 0x45, 0x61, 0xd3, 0x8e, 0x86, 0x8a, 0x7a, 0x76, 0xa8, 0xaa, 0x0d,
	0x6f, 0x60, 0xef, 0x36, 0x4d, 0x47, 0xa6, 0xfb, 0xf7, 0xa2, 0x57, 0xdf, 0x1d, 0xe8, 0xbf, 0x94,
	0x2e, 0x9f, 0x51, 0x93, 0x57, 0xf0, 0xdd, 0x92, 0xe4, 0x9c, 0x35, 0x09, 0x67, 0x2b, 0x56, 0x83,
	0x68, 0x3d, 0x58, 0xee, 0x70, 0x07, 0xbe, 0x73, 0xd3, 0x16, 0xbe, 0x62, 0x2f, 0xd8, 0xaf, 0x40,
	0x29, 0x12, 0x76, 0x6f, 0x3e, 0x2e, 0x99, 0x40, 0xcf, 0x98, 0x21, 0xa7, 0xcd, 0x19, 0x35, 0xb9,
	0xc1, 0xd9, 0x3a, 0xac, 0x1c, 0xee, 0x01, 0xb6, 0x2a, 0x8d, 0xe4, 0xa2, 0xf9, 0xcc, 0x2f, 0xd5,
	0x7f, 0x0c, 0x38, 0xf5, 0xed, 0xef, 0x77, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x48, 0xae, 0x16,
	0xc8, 0xde, 0x02, 0x00, 0x00,
}
