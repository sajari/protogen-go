// Code generated by protoc-gen-go.
// source: sajari/engine/schema/schema.proto
// DO NOT EDIT!

/*
Package sajari_engine_schema is a generated protocol buffer package.

It is generated from these files:
	sajari/engine/schema/schema.proto

It has these top-level messages:
	Fields
	Field
	Response
	MutateFieldRequest
*/
package sajari_engine_schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sajari_rpc "code.sajari.com/protogen-go/sajari/rpc"
import sajari_rpc1 "code.sajari.com/protogen-go/sajari/rpc"

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

// Type represents the underlying data type of the field. Default is a string.
type Field_Type int32

const (
	Field_STRING    Field_Type = 0
	Field_INTEGER   Field_Type = 1
	Field_FLOAT     Field_Type = 2
	Field_BOOLEAN   Field_Type = 3
	Field_TIMESTAMP Field_Type = 4
)

var Field_Type_name = map[int32]string{
	0: "STRING",
	1: "INTEGER",
	2: "FLOAT",
	3: "BOOLEAN",
	4: "TIMESTAMP",
}
var Field_Type_value = map[string]int32{
	"STRING":    0,
	"INTEGER":   1,
	"FLOAT":     2,
	"BOOLEAN":   3,
	"TIMESTAMP": 4,
}

func (x Field_Type) String() string {
	return proto.EnumName(Field_Type_name, int32(x))
}
func (Field_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// Fields is a list of Fields.
type Fields struct {
	Fields []*Field `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty"`
}

func (m *Fields) Reset()                    { *m = Fields{} }
func (m *Fields) String() string            { return proto.CompactTextString(m) }
func (*Fields) ProtoMessage()               {}
func (*Fields) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Fields) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

// Field defines the properties of of a field in the schema.
type Field struct {
	// ID used internally
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name of the field.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Description of the field.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Type of the field.
	Type Field_Type `protobuf:"varint,4,opt,name=type,enum=sajari.engine.schema.Field_Type" json:"type,omitempty"`
	// Repeated sets that fields are repeated values (i.e. a list of values).
	Repeated bool `protobuf:"varint,5,opt,name=repeated" json:"repeated,omitempty"`
	// Required sets that the field must be specified when adding new records.
	Required bool `protobuf:"varint,6,opt,name=required" json:"required,omitempty"`
	// Store is set to false if original values should not be stored. Ideal for large text
	// fields that need to be indexed but will never be returned in results.
	Store bool `protobuf:"varint,7,opt,name=store" json:"store,omitempty"`
	// Indexed is set to true if this field should be indexed into a reverse index.
	// Only string and string array fields can be indexed.
	Indexed bool `protobuf:"varint,8,opt,name=indexed" json:"indexed,omitempty"`
	// Unique is set to true if this field should contain unique values. Unique
	// fields also support indexing to assist with fast record lookups
	Unique bool `protobuf:"varint,9,opt,name=unique" json:"unique,omitempty"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Field) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Field) GetType() Field_Type {
	if m != nil {
		return m.Type
	}
	return Field_STRING
}

func (m *Field) GetRepeated() bool {
	if m != nil {
		return m.Repeated
	}
	return false
}

func (m *Field) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *Field) GetStore() bool {
	if m != nil {
		return m.Store
	}
	return false
}

func (m *Field) GetIndexed() bool {
	if m != nil {
		return m.Indexed
	}
	return false
}

func (m *Field) GetUnique() bool {
	if m != nil {
		return m.Unique
	}
	return false
}

// Response is the message type returned from AddFields and MutateField calls.
type Response struct {
	Status []*sajari_rpc1.Status `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetStatus() []*sajari_rpc1.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// MutateFieldRequest is a message type passed to the MutateField method.
type MutateFieldRequest struct {
	// Name of the field to mutate.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// List of mutations to apply to the field.  Each is taken individually, if there
	// are any errors then the rest are ignored.
	Mutations []*MutateFieldRequest_Mutation `protobuf:"bytes,2,rep,name=mutations" json:"mutations,omitempty"`
}

func (m *MutateFieldRequest) Reset()                    { *m = MutateFieldRequest{} }
func (m *MutateFieldRequest) String() string            { return proto.CompactTextString(m) }
func (*MutateFieldRequest) ProtoMessage()               {}
func (*MutateFieldRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MutateFieldRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MutateFieldRequest) GetMutations() []*MutateFieldRequest_Mutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

// Mutation is a mutation of a schema field.
type MutateFieldRequest_Mutation struct {
	// Types that are valid to be assigned to Mutation:
	//	*MutateFieldRequest_Mutation_Name
	//	*MutateFieldRequest_Mutation_Description
	//	*MutateFieldRequest_Mutation_Type
	//	*MutateFieldRequest_Mutation_Repeated
	//	*MutateFieldRequest_Mutation_Required
	//	*MutateFieldRequest_Mutation_Unique
	//	*MutateFieldRequest_Mutation_Indexed
	Mutation isMutateFieldRequest_Mutation_Mutation `protobuf_oneof:"mutation"`
}

func (m *MutateFieldRequest_Mutation) Reset()                    { *m = MutateFieldRequest_Mutation{} }
func (m *MutateFieldRequest_Mutation) String() string            { return proto.CompactTextString(m) }
func (*MutateFieldRequest_Mutation) ProtoMessage()               {}
func (*MutateFieldRequest_Mutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type isMutateFieldRequest_Mutation_Mutation interface {
	isMutateFieldRequest_Mutation_Mutation()
}

type MutateFieldRequest_Mutation_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,oneof"`
}
type MutateFieldRequest_Mutation_Description struct {
	Description string `protobuf:"bytes,2,opt,name=description,oneof"`
}
type MutateFieldRequest_Mutation_Type struct {
	Type Field_Type `protobuf:"varint,3,opt,name=type,enum=sajari.engine.schema.Field_Type,oneof"`
}
type MutateFieldRequest_Mutation_Repeated struct {
	Repeated bool `protobuf:"varint,4,opt,name=repeated,oneof"`
}
type MutateFieldRequest_Mutation_Required struct {
	Required bool `protobuf:"varint,5,opt,name=required,oneof"`
}
type MutateFieldRequest_Mutation_Unique struct {
	Unique bool `protobuf:"varint,6,opt,name=unique,oneof"`
}
type MutateFieldRequest_Mutation_Indexed struct {
	Indexed bool `protobuf:"varint,7,opt,name=indexed,oneof"`
}

func (*MutateFieldRequest_Mutation_Name) isMutateFieldRequest_Mutation_Mutation()        {}
func (*MutateFieldRequest_Mutation_Description) isMutateFieldRequest_Mutation_Mutation() {}
func (*MutateFieldRequest_Mutation_Type) isMutateFieldRequest_Mutation_Mutation()        {}
func (*MutateFieldRequest_Mutation_Repeated) isMutateFieldRequest_Mutation_Mutation()    {}
func (*MutateFieldRequest_Mutation_Required) isMutateFieldRequest_Mutation_Mutation()    {}
func (*MutateFieldRequest_Mutation_Unique) isMutateFieldRequest_Mutation_Mutation()      {}
func (*MutateFieldRequest_Mutation_Indexed) isMutateFieldRequest_Mutation_Mutation()     {}

func (m *MutateFieldRequest_Mutation) GetMutation() isMutateFieldRequest_Mutation_Mutation {
	if m != nil {
		return m.Mutation
	}
	return nil
}

func (m *MutateFieldRequest_Mutation) GetName() string {
	if x, ok := m.GetMutation().(*MutateFieldRequest_Mutation_Name); ok {
		return x.Name
	}
	return ""
}

func (m *MutateFieldRequest_Mutation) GetDescription() string {
	if x, ok := m.GetMutation().(*MutateFieldRequest_Mutation_Description); ok {
		return x.Description
	}
	return ""
}

func (m *MutateFieldRequest_Mutation) GetType() Field_Type {
	if x, ok := m.GetMutation().(*MutateFieldRequest_Mutation_Type); ok {
		return x.Type
	}
	return Field_STRING
}

func (m *MutateFieldRequest_Mutation) GetRepeated() bool {
	if x, ok := m.GetMutation().(*MutateFieldRequest_Mutation_Repeated); ok {
		return x.Repeated
	}
	return false
}

func (m *MutateFieldRequest_Mutation) GetRequired() bool {
	if x, ok := m.GetMutation().(*MutateFieldRequest_Mutation_Required); ok {
		return x.Required
	}
	return false
}

func (m *MutateFieldRequest_Mutation) GetUnique() bool {
	if x, ok := m.GetMutation().(*MutateFieldRequest_Mutation_Unique); ok {
		return x.Unique
	}
	return false
}

func (m *MutateFieldRequest_Mutation) GetIndexed() bool {
	if x, ok := m.GetMutation().(*MutateFieldRequest_Mutation_Indexed); ok {
		return x.Indexed
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MutateFieldRequest_Mutation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MutateFieldRequest_Mutation_OneofMarshaler, _MutateFieldRequest_Mutation_OneofUnmarshaler, _MutateFieldRequest_Mutation_OneofSizer, []interface{}{
		(*MutateFieldRequest_Mutation_Name)(nil),
		(*MutateFieldRequest_Mutation_Description)(nil),
		(*MutateFieldRequest_Mutation_Type)(nil),
		(*MutateFieldRequest_Mutation_Repeated)(nil),
		(*MutateFieldRequest_Mutation_Required)(nil),
		(*MutateFieldRequest_Mutation_Unique)(nil),
		(*MutateFieldRequest_Mutation_Indexed)(nil),
	}
}

func _MutateFieldRequest_Mutation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MutateFieldRequest_Mutation)
	// mutation
	switch x := m.Mutation.(type) {
	case *MutateFieldRequest_Mutation_Name:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case *MutateFieldRequest_Mutation_Description:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Description)
	case *MutateFieldRequest_Mutation_Type:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Type))
	case *MutateFieldRequest_Mutation_Repeated:
		t := uint64(0)
		if x.Repeated {
			t = 1
		}
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *MutateFieldRequest_Mutation_Required:
		t := uint64(0)
		if x.Required {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *MutateFieldRequest_Mutation_Unique:
		t := uint64(0)
		if x.Unique {
			t = 1
		}
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *MutateFieldRequest_Mutation_Indexed:
		t := uint64(0)
		if x.Indexed {
			t = 1
		}
		b.EncodeVarint(7<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("MutateFieldRequest_Mutation.Mutation has unexpected type %T", x)
	}
	return nil
}

func _MutateFieldRequest_Mutation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MutateFieldRequest_Mutation)
	switch tag {
	case 1: // mutation.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Mutation = &MutateFieldRequest_Mutation_Name{x}
		return true, err
	case 2: // mutation.description
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Mutation = &MutateFieldRequest_Mutation_Description{x}
		return true, err
	case 3: // mutation.type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Mutation = &MutateFieldRequest_Mutation_Type{Field_Type(x)}
		return true, err
	case 4: // mutation.repeated
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Mutation = &MutateFieldRequest_Mutation_Repeated{x != 0}
		return true, err
	case 5: // mutation.required
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Mutation = &MutateFieldRequest_Mutation_Required{x != 0}
		return true, err
	case 6: // mutation.unique
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Mutation = &MutateFieldRequest_Mutation_Unique{x != 0}
		return true, err
	case 7: // mutation.indexed
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Mutation = &MutateFieldRequest_Mutation_Indexed{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _MutateFieldRequest_Mutation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MutateFieldRequest_Mutation)
	// mutation
	switch x := m.Mutation.(type) {
	case *MutateFieldRequest_Mutation_Name:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *MutateFieldRequest_Mutation_Description:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Description)))
		n += len(x.Description)
	case *MutateFieldRequest_Mutation_Type:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Type))
	case *MutateFieldRequest_Mutation_Repeated:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += 1
	case *MutateFieldRequest_Mutation_Required:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += 1
	case *MutateFieldRequest_Mutation_Unique:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += 1
	case *MutateFieldRequest_Mutation_Indexed:
		n += proto.SizeVarint(7<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Fields)(nil), "sajari.engine.schema.Fields")
	proto.RegisterType((*Field)(nil), "sajari.engine.schema.Field")
	proto.RegisterType((*Response)(nil), "sajari.engine.schema.Response")
	proto.RegisterType((*MutateFieldRequest)(nil), "sajari.engine.schema.MutateFieldRequest")
	proto.RegisterType((*MutateFieldRequest_Mutation)(nil), "sajari.engine.schema.MutateFieldRequest.Mutation")
	proto.RegisterEnum("sajari.engine.schema.Field_Type", Field_Type_name, Field_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Schema service

type SchemaClient interface {
	// GetFields returns the fields in the schema.
	GetFields(ctx context.Context, in *sajari_rpc.Empty, opts ...grpc.CallOption) (*Fields, error)
	// AddFields adds new fields to the schema.
	AddFields(ctx context.Context, in *Fields, opts ...grpc.CallOption) (*Response, error)
	// MutateField mutates a field in the schema.
	MutateField(ctx context.Context, in *MutateFieldRequest, opts ...grpc.CallOption) (*Response, error)
}

type schemaClient struct {
	cc *grpc.ClientConn
}

func NewSchemaClient(cc *grpc.ClientConn) SchemaClient {
	return &schemaClient{cc}
}

func (c *schemaClient) GetFields(ctx context.Context, in *sajari_rpc.Empty, opts ...grpc.CallOption) (*Fields, error) {
	out := new(Fields)
	err := grpc.Invoke(ctx, "/sajari.engine.schema.Schema/GetFields", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaClient) AddFields(ctx context.Context, in *Fields, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/sajari.engine.schema.Schema/AddFields", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaClient) MutateField(ctx context.Context, in *MutateFieldRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/sajari.engine.schema.Schema/MutateField", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Schema service

type SchemaServer interface {
	// GetFields returns the fields in the schema.
	GetFields(context.Context, *sajari_rpc.Empty) (*Fields, error)
	// AddFields adds new fields to the schema.
	AddFields(context.Context, *Fields) (*Response, error)
	// MutateField mutates a field in the schema.
	MutateField(context.Context, *MutateFieldRequest) (*Response, error)
}

func RegisterSchemaServer(s *grpc.Server, srv SchemaServer) {
	s.RegisterService(&_Schema_serviceDesc, srv)
}

func _Schema_GetFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sajari_rpc.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServer).GetFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.engine.schema.Schema/GetFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServer).GetFields(ctx, req.(*sajari_rpc.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schema_AddFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fields)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServer).AddFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.engine.schema.Schema/AddFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServer).AddFields(ctx, req.(*Fields))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schema_MutateField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServer).MutateField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.engine.schema.Schema/MutateField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServer).MutateField(ctx, req.(*MutateFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Schema_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sajari.engine.schema.Schema",
	HandlerType: (*SchemaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFields",
			Handler:    _Schema_GetFields_Handler,
		},
		{
			MethodName: "AddFields",
			Handler:    _Schema_AddFields_Handler,
		},
		{
			MethodName: "MutateField",
			Handler:    _Schema_MutateField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sajari/engine/schema/schema.proto",
}

func init() { proto.RegisterFile("sajari/engine/schema/schema.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0x4f, 0x8f, 0xd2, 0x40,
	0x18, 0xc6, 0xb7, 0x05, 0x4a, 0xfb, 0x92, 0xdd, 0xe0, 0x64, 0xb3, 0x4e, 0xea, 0xc6, 0xd4, 0x9e,
	0x1a, 0x0f, 0x25, 0xb2, 0x66, 0x4f, 0x7a, 0x60, 0x13, 0x60, 0x49, 0x16, 0x30, 0xa5, 0x89, 0xe7,
	0x4a, 0x5f, 0x75, 0x8c, 0xb4, 0xa5, 0x33, 0x4d, 0xe4, 0x33, 0xfa, 0x4d, 0xfc, 0x00, 0x1e, 0x3c,
	0x99, 0x4e, 0xa7, 0xd0, 0x8d, 0x04, 0xf7, 0xc4, 0xbc, 0xcf, 0xf3, 0xf0, 0xce, 0x9f, 0xdf, 0x4c,
	0xe1, 0x15, 0x8f, 0xbe, 0x45, 0x39, 0x1b, 0x60, 0xf2, 0x85, 0x25, 0x38, 0xe0, 0xeb, 0xaf, 0xb8,
	0x89, 0xd4, 0x8f, 0x9f, 0xe5, 0xa9, 0x48, 0xc9, 0x65, 0x15, 0xf1, 0xab, 0x88, 0x5f, 0x79, 0xf6,
	0x95, 0xfa, 0x63, 0x9e, 0xad, 0x07, 0xb8, 0xc9, 0xc4, 0xae, 0x4a, 0xdb, 0xcf, 0x1b, 0x3a, 0x17,
	0x91, 0x28, 0x78, 0x65, 0xb8, 0xef, 0xc1, 0x98, 0x30, 0xfc, 0x1e, 0x73, 0x72, 0x03, 0xc6, 0x67,
	0x39, 0xa2, 0x9a, 0xd3, 0xf2, 0x7a, 0xc3, 0x17, 0xfe, 0xb1, 0x19, 0x7c, 0x99, 0x0e, 0x54, 0xd4,
	0xfd, 0xa9, 0x43, 0x47, 0x2a, 0xe4, 0x02, 0x74, 0x16, 0x53, 0xcd, 0xd1, 0xbc, 0xf3, 0x40, 0x67,
	0x31, 0x21, 0xd0, 0x4e, 0xa2, 0x0d, 0x52, 0xdd, 0xd1, 0x3c, 0x2b, 0x90, 0x63, 0xe2, 0x40, 0x2f,
	0x46, 0xbe, 0xce, 0x59, 0x26, 0x58, 0x9a, 0xd0, 0x96, 0xb4, 0x9a, 0x12, 0x79, 0x0b, 0x6d, 0xb1,
	0xcb, 0x90, 0xb6, 0x1d, 0xcd, 0xbb, 0x18, 0x3a, 0x27, 0x96, 0xe0, 0x87, 0xbb, 0x0c, 0x03, 0x99,
	0x26, 0x36, 0x98, 0x39, 0x66, 0x18, 0x09, 0x8c, 0x69, 0xc7, 0xd1, 0x3c, 0x33, 0xd8, 0xd7, 0x95,
	0xb7, 0x2d, 0x58, 0x8e, 0x31, 0x35, 0x6a, 0xaf, 0xaa, 0xc9, 0x25, 0x74, 0xb8, 0x48, 0x73, 0xa4,
	0x5d, 0x69, 0x54, 0x05, 0xa1, 0xd0, 0x65, 0x49, 0x8c, 0x3f, 0x30, 0xa6, 0xa6, 0xd4, 0xeb, 0x92,
	0x5c, 0x81, 0x51, 0x24, 0x6c, 0x5b, 0x20, 0xb5, 0xa4, 0xa1, 0x2a, 0x77, 0x02, 0xed, 0x72, 0x35,
	0x04, 0xc0, 0x58, 0x85, 0xc1, 0x6c, 0x31, 0xed, 0x9f, 0x91, 0x1e, 0x74, 0x67, 0x8b, 0x70, 0x3c,
	0x1d, 0x07, 0x7d, 0x8d, 0x58, 0xd0, 0x99, 0x3c, 0x2c, 0x47, 0x61, 0x5f, 0x2f, 0xf5, 0xbb, 0xe5,
	0xf2, 0x61, 0x3c, 0x5a, 0xf4, 0x5b, 0xe4, 0x1c, 0xac, 0x70, 0x36, 0x1f, 0xaf, 0xc2, 0xd1, 0xfc,
	0x43, 0xbf, 0xed, 0xde, 0x82, 0x19, 0x20, 0xcf, 0xd2, 0x84, 0x23, 0x79, 0x0d, 0x46, 0x05, 0x4a,
	0xe1, 0x20, 0xf5, 0x59, 0xe4, 0xd9, 0xda, 0x5f, 0x49, 0x27, 0x50, 0x09, 0xf7, 0xb7, 0x0e, 0x64,
	0x5e, 0x88, 0x48, 0x60, 0x45, 0x07, 0xb7, 0x05, 0x72, 0xb1, 0x47, 0xa0, 0x35, 0x10, 0x2c, 0xc1,
	0xda, 0x94, 0x49, 0x96, 0x26, 0x9c, 0xea, 0xb2, 0xf3, 0x9b, 0xe3, 0xa7, 0xfc, 0x6f, 0xc3, 0x4a,
	0x62, 0x69, 0x12, 0x1c, 0x7a, 0xd8, 0x7f, 0x34, 0x30, 0x6b, 0x9d, 0x5c, 0x36, 0x67, 0xbc, 0x3f,
	0x53, 0x73, 0xba, 0x8f, 0xb1, 0xeb, 0xca, 0x7c, 0x04, 0xfe, 0x56, 0x81, 0x6f, 0x3d, 0x0d, 0x7c,
	0xd9, 0x5b, 0xa2, 0xbf, 0x6e, 0xa0, 0x2f, 0x2f, 0x8d, 0x79, 0x7f, 0xd6, 0x80, 0x7f, 0xdd, 0x80,
	0xdf, 0x39, 0xb8, 0x0a, 0x3f, 0xdd, 0xe3, 0x34, 0x94, 0xa7, 0x6a, 0x62, 0x1f, 0xae, 0x40, 0x57,
	0x59, 0xb5, 0x70, 0x07, 0x60, 0xd6, 0xbb, 0x1f, 0xfe, 0xd2, 0xc0, 0x58, 0xc9, 0xb5, 0x91, 0x77,
	0x60, 0x4d, 0x51, 0xa8, 0xb7, 0xf4, 0xac, 0x09, 0x6b, 0x5c, 0xbe, 0x43, 0xfb, 0xfa, 0xc4, 0x96,
	0x38, 0x99, 0x81, 0x35, 0x8a, 0x63, 0x55, 0x9c, 0x8c, 0xda, 0x2f, 0x8f, 0xbb, 0xfb, 0x8b, 0xf3,
	0x11, 0x7a, 0x0d, 0x74, 0xc4, 0x7b, 0x2a, 0xdd, 0xff, 0x35, 0xfe, 0x64, 0xc8, 0x2f, 0xc6, 0xcd,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x75, 0x04, 0x9d, 0x9d, 0x04, 0x00, 0x00,
}
