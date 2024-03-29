// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sajari/engine/schema/schema.proto

package schema

import (
	rpc "code.sajari.com/protogen-go/sajari/rpc"
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

// Type represents the underlying data type of the field. Default is a string.
type Field_Type int32

const (
	Field_STRING    Field_Type = 0
	Field_INTEGER   Field_Type = 1
	Field_FLOAT     Field_Type = 2
	Field_DOUBLE    Field_Type = 5
	Field_BOOLEAN   Field_Type = 3
	Field_TIMESTAMP Field_Type = 4
)

var Field_Type_name = map[int32]string{
	0: "STRING",
	1: "INTEGER",
	2: "FLOAT",
	5: "DOUBLE",
	3: "BOOLEAN",
	4: "TIMESTAMP",
}

var Field_Type_value = map[string]int32{
	"STRING":    0,
	"INTEGER":   1,
	"FLOAT":     2,
	"DOUBLE":    5,
	"BOOLEAN":   3,
	"TIMESTAMP": 4,
}

func (x Field_Type) String() string {
	return proto.EnumName(Field_Type_name, int32(x))
}

func (Field_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c268116a1997bf59, []int{1, 0}
}

// Mode is an enumeration of modes for a field.
type Field_Mode int32

const (
	// Nullable fields do not need to be set.
	Field_NULLABLE Field_Mode = 0
	// Required fields must be specified (cannot be null).
	Field_REQUIRED Field_Mode = 1
	// Unique fields must be specified, and must be unique.
	Field_UNIQUE Field_Mode = 2
)

var Field_Mode_name = map[int32]string{
	0: "NULLABLE",
	1: "REQUIRED",
	2: "UNIQUE",
}

var Field_Mode_value = map[string]int32{
	"NULLABLE": 0,
	"REQUIRED": 1,
	"UNIQUE":   2,
}

func (x Field_Mode) String() string {
	return proto.EnumName(Field_Mode_name, int32(x))
}

func (Field_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c268116a1997bf59, []int{1, 1}
}

// Fields is a list of Fields.
type Fields struct {
	Fields               []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fields) Reset()         { *m = Fields{} }
func (m *Fields) String() string { return proto.CompactTextString(m) }
func (*Fields) ProtoMessage()    {}
func (*Fields) Descriptor() ([]byte, []int) {
	return fileDescriptor_c268116a1997bf59, []int{0}
}

func (m *Fields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fields.Unmarshal(m, b)
}
func (m *Fields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fields.Marshal(b, m, deterministic)
}
func (m *Fields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fields.Merge(m, src)
}
func (m *Fields) XXX_Size() int {
	return xxx_messageInfo_Fields.Size(m)
}
func (m *Fields) XXX_DiscardUnknown() {
	xxx_messageInfo_Fields.DiscardUnknown(m)
}

var xxx_messageInfo_Fields proto.InternalMessageInfo

func (m *Fields) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

// Field defines the properties of of a field in the schema.
type Field struct {
	// ID used internally
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the field.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the field.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Type of the field.
	Type Field_Type `protobuf:"varint,4,opt,name=type,proto3,enum=sajari.engine.schema.Field_Type" json:"type,omitempty"`
	// Repeated sets that fields are repeated values (i.e. a list of values).
	Repeated bool `protobuf:"varint,5,opt,name=repeated,proto3" json:"repeated,omitempty"`
	// Repeated length sets the required length for the repeated field.
	RepeatedLen int32 `protobuf:"varint,11,opt,name=repeated_len,json=repeatedLen,proto3" json:"repeated_len,omitempty"`
	// Required sets that the field must be specified when adding new records.
	// DEPRECATED: use mode instead.
	Required bool `protobuf:"varint,6,opt,name=required,proto3" json:"required,omitempty"`
	// Store is set to false if original values should not be stored. Ideal for large text
	// fields that need to be indexed but will never be returned in results.
	// DEPRECATED: this field is now ignored.
	Store bool `protobuf:"varint,7,opt,name=store,proto3" json:"store,omitempty"`
	// Indexed is set to true if this field should be indexed into a reverse index.
	// Only string and string array fields can be indexed.
	// DEPRECATED: this field is now ignored for incoming values.
	Indexed bool `protobuf:"varint,8,opt,name=indexed,proto3" json:"indexed,omitempty"`
	// Unique is set to true if this field should contain unique values. Unique
	// fields also support indexing to assist with fast record lookups.
	// DEPRECATED: use mode instead.
	Unique bool `protobuf:"varint,9,opt,name=unique,proto3" json:"unique,omitempty"`
	// Mode of the field.
	Mode Field_Mode `protobuf:"varint,12,opt,name=mode,proto3,enum=sajari.engine.schema.Field_Mode" json:"mode,omitempty"`
	// Indexes is a list of indexes defined for the field.
	Indexes              []*FieldIndex `protobuf:"bytes,10,rep,name=indexes,proto3" json:"indexes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_c268116a1997bf59, []int{1}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

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

func (m *Field) GetRepeatedLen() int32 {
	if m != nil {
		return m.RepeatedLen
	}
	return 0
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

func (m *Field) GetMode() Field_Mode {
	if m != nil {
		return m.Mode
	}
	return Field_NULLABLE
}

func (m *Field) GetIndexes() []*FieldIndex {
	if m != nil {
		return m.Indexes
	}
	return nil
}

// FieldIndex
type FieldIndex struct {
	// ID used internally
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Specification of how it was created.
	Spec string `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// Description of the index.
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldIndex) Reset()         { *m = FieldIndex{} }
func (m *FieldIndex) String() string { return proto.CompactTextString(m) }
func (*FieldIndex) ProtoMessage()    {}
func (*FieldIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_c268116a1997bf59, []int{2}
}

func (m *FieldIndex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldIndex.Unmarshal(m, b)
}
func (m *FieldIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldIndex.Marshal(b, m, deterministic)
}
func (m *FieldIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldIndex.Merge(m, src)
}
func (m *FieldIndex) XXX_Size() int {
	return xxx_messageInfo_FieldIndex.Size(m)
}
func (m *FieldIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldIndex.DiscardUnknown(m)
}

var xxx_messageInfo_FieldIndex proto.InternalMessageInfo

func (m *FieldIndex) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FieldIndex) GetSpec() string {
	if m != nil {
		return m.Spec
	}
	return ""
}

func (m *FieldIndex) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Response is the message type returned from AddFields and MutateField calls.
type Response struct {
	Status               []*rpc.Status `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c268116a1997bf59, []int{3}
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

func (m *Response) GetStatus() []*rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// MutateFieldRequest is a message type passed to the MutateField method.
type MutateFieldRequest struct {
	// Name of the field to mutate.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of mutations to apply to the field.  Each is taken individually, if there
	// are any errors then the rest are ignored.
	Mutations            []*MutateFieldRequest_Mutation `protobuf:"bytes,2,rep,name=mutations,proto3" json:"mutations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MutateFieldRequest) Reset()         { *m = MutateFieldRequest{} }
func (m *MutateFieldRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFieldRequest) ProtoMessage()    {}
func (*MutateFieldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c268116a1997bf59, []int{4}
}

func (m *MutateFieldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFieldRequest.Unmarshal(m, b)
}
func (m *MutateFieldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFieldRequest.Marshal(b, m, deterministic)
}
func (m *MutateFieldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFieldRequest.Merge(m, src)
}
func (m *MutateFieldRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFieldRequest.Size(m)
}
func (m *MutateFieldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFieldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFieldRequest proto.InternalMessageInfo

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
	//	*MutateFieldRequest_Mutation_AddIndex
	Mutation             isMutateFieldRequest_Mutation_Mutation `protobuf_oneof:"mutation"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *MutateFieldRequest_Mutation) Reset()         { *m = MutateFieldRequest_Mutation{} }
func (m *MutateFieldRequest_Mutation) String() string { return proto.CompactTextString(m) }
func (*MutateFieldRequest_Mutation) ProtoMessage()    {}
func (*MutateFieldRequest_Mutation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c268116a1997bf59, []int{4, 0}
}

func (m *MutateFieldRequest_Mutation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFieldRequest_Mutation.Unmarshal(m, b)
}
func (m *MutateFieldRequest_Mutation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFieldRequest_Mutation.Marshal(b, m, deterministic)
}
func (m *MutateFieldRequest_Mutation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFieldRequest_Mutation.Merge(m, src)
}
func (m *MutateFieldRequest_Mutation) XXX_Size() int {
	return xxx_messageInfo_MutateFieldRequest_Mutation.Size(m)
}
func (m *MutateFieldRequest_Mutation) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFieldRequest_Mutation.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFieldRequest_Mutation proto.InternalMessageInfo

type isMutateFieldRequest_Mutation_Mutation interface {
	isMutateFieldRequest_Mutation_Mutation()
}

type MutateFieldRequest_Mutation_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type MutateFieldRequest_Mutation_Description struct {
	Description string `protobuf:"bytes,2,opt,name=description,proto3,oneof"`
}

type MutateFieldRequest_Mutation_Type struct {
	Type Field_Type `protobuf:"varint,3,opt,name=type,proto3,enum=sajari.engine.schema.Field_Type,oneof"`
}

type MutateFieldRequest_Mutation_Repeated struct {
	Repeated bool `protobuf:"varint,4,opt,name=repeated,proto3,oneof"`
}

type MutateFieldRequest_Mutation_Required struct {
	Required bool `protobuf:"varint,5,opt,name=required,proto3,oneof"`
}

type MutateFieldRequest_Mutation_Unique struct {
	Unique bool `protobuf:"varint,6,opt,name=unique,proto3,oneof"`
}

type MutateFieldRequest_Mutation_Indexed struct {
	Indexed bool `protobuf:"varint,7,opt,name=indexed,proto3,oneof"`
}

type MutateFieldRequest_Mutation_AddIndex struct {
	AddIndex *FieldIndex `protobuf:"bytes,8,opt,name=add_index,json=addIndex,proto3,oneof"`
}

func (*MutateFieldRequest_Mutation_Name) isMutateFieldRequest_Mutation_Mutation() {}

func (*MutateFieldRequest_Mutation_Description) isMutateFieldRequest_Mutation_Mutation() {}

func (*MutateFieldRequest_Mutation_Type) isMutateFieldRequest_Mutation_Mutation() {}

func (*MutateFieldRequest_Mutation_Repeated) isMutateFieldRequest_Mutation_Mutation() {}

func (*MutateFieldRequest_Mutation_Required) isMutateFieldRequest_Mutation_Mutation() {}

func (*MutateFieldRequest_Mutation_Unique) isMutateFieldRequest_Mutation_Mutation() {}

func (*MutateFieldRequest_Mutation_Indexed) isMutateFieldRequest_Mutation_Mutation() {}

func (*MutateFieldRequest_Mutation_AddIndex) isMutateFieldRequest_Mutation_Mutation() {}

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

func (m *MutateFieldRequest_Mutation) GetAddIndex() *FieldIndex {
	if x, ok := m.GetMutation().(*MutateFieldRequest_Mutation_AddIndex); ok {
		return x.AddIndex
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MutateFieldRequest_Mutation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MutateFieldRequest_Mutation_Name)(nil),
		(*MutateFieldRequest_Mutation_Description)(nil),
		(*MutateFieldRequest_Mutation_Type)(nil),
		(*MutateFieldRequest_Mutation_Repeated)(nil),
		(*MutateFieldRequest_Mutation_Required)(nil),
		(*MutateFieldRequest_Mutation_Unique)(nil),
		(*MutateFieldRequest_Mutation_Indexed)(nil),
		(*MutateFieldRequest_Mutation_AddIndex)(nil),
	}
}

func init() {
	proto.RegisterEnum("sajari.engine.schema.Field_Type", Field_Type_name, Field_Type_value)
	proto.RegisterEnum("sajari.engine.schema.Field_Mode", Field_Mode_name, Field_Mode_value)
	proto.RegisterType((*Fields)(nil), "sajari.engine.schema.Fields")
	proto.RegisterType((*Field)(nil), "sajari.engine.schema.Field")
	proto.RegisterType((*FieldIndex)(nil), "sajari.engine.schema.FieldIndex")
	proto.RegisterType((*Response)(nil), "sajari.engine.schema.Response")
	proto.RegisterType((*MutateFieldRequest)(nil), "sajari.engine.schema.MutateFieldRequest")
	proto.RegisterType((*MutateFieldRequest_Mutation)(nil), "sajari.engine.schema.MutateFieldRequest.Mutation")
}

func init() { proto.RegisterFile("sajari/engine/schema/schema.proto", fileDescriptor_c268116a1997bf59) }

var fileDescriptor_c268116a1997bf59 = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xda, 0x4a,
	0x10, 0xc6, 0x06, 0x8c, 0x3d, 0x24, 0x91, 0xcf, 0x2a, 0xca, 0x59, 0x71, 0xa2, 0x23, 0xc7, 0x57,
	0xe8, 0x48, 0xc7, 0xb4, 0xa4, 0xca, 0x45, 0xd5, 0xaa, 0x02, 0xc5, 0x49, 0x90, 0xf8, 0x69, 0x16,
	0x50, 0xa5, 0xde, 0x44, 0x2e, 0x3b, 0x4d, 0x5d, 0x05, 0xdb, 0xb1, 0x8d, 0xd4, 0xbc, 0x4b, 0xdf,
	0xa3, 0x17, 0x7d, 0xa3, 0x3e, 0x45, 0xb5, 0x6b, 0x1b, 0x1c, 0x15, 0x91, 0x5c, 0xb1, 0x33, 0xdf,
	0x37, 0x33, 0xeb, 0xdd, 0xef, 0x5b, 0xe0, 0x24, 0xf1, 0xbe, 0x7a, 0xb1, 0xdf, 0xc1, 0xe0, 0xd6,
	0x0f, 0xb0, 0x93, 0x2c, 0xbe, 0xe0, 0xd2, 0xcb, 0x7f, 0x9c, 0x28, 0x0e, 0xd3, 0x90, 0x1c, 0x66,
	0x14, 0x27, 0xa3, 0x38, 0x19, 0xd6, 0x3a, 0xca, 0x0b, 0xe3, 0x68, 0xd1, 0xc1, 0x65, 0x94, 0x3e,
	0x64, 0xec, 0xd6, 0xdf, 0xa5, 0x7c, 0x92, 0x7a, 0xe9, 0x2a, 0xc9, 0x00, 0xfb, 0x2d, 0x68, 0x17,
	0x3e, 0xde, 0xf1, 0x84, 0x9c, 0x82, 0xf6, 0x59, 0xae, 0xa8, 0x62, 0x55, 0xdb, 0xcd, 0xee, 0x3f,
	0xce, 0xb6, 0x09, 0x8e, 0x64, 0xb3, 0x9c, 0x6a, 0x7f, 0xaf, 0x41, 0x5d, 0x66, 0xc8, 0x01, 0xa8,
	0x3e, 0xa7, 0x8a, 0xa5, 0xb4, 0xf7, 0x99, 0xea, 0x73, 0x42, 0xa0, 0x16, 0x78, 0x4b, 0xa4, 0xaa,
	0xa5, 0xb4, 0x0d, 0x26, 0xd7, 0xc4, 0x82, 0x26, 0xc7, 0x64, 0x11, 0xfb, 0x51, 0xea, 0x87, 0x01,
	0xad, 0x4a, 0xa8, 0x9c, 0x22, 0xaf, 0xa0, 0x96, 0x3e, 0x44, 0x48, 0x6b, 0x96, 0xd2, 0x3e, 0xe8,
	0x5a, 0x3b, 0xb6, 0xe0, 0xcc, 0x1e, 0x22, 0x64, 0x92, 0x4d, 0x5a, 0xa0, 0xc7, 0x18, 0xa1, 0x97,
	0x22, 0xa7, 0x75, 0x4b, 0x69, 0xeb, 0x6c, 0x1d, 0x93, 0x13, 0xd8, 0x2b, 0xd6, 0x37, 0x77, 0x18,
	0xd0, 0xa6, 0xa5, 0xb4, 0xeb, 0xac, 0x59, 0xe4, 0x86, 0x18, 0x64, 0xe5, 0xf7, 0x2b, 0x3f, 0x46,
	0x4e, 0xb5, 0xa2, 0x3c, 0x8b, 0xc9, 0x21, 0xd4, 0x93, 0x34, 0x8c, 0x91, 0x36, 0x24, 0x90, 0x05,
	0x84, 0x42, 0xc3, 0x0f, 0x38, 0x7e, 0x43, 0x4e, 0x75, 0x99, 0x2f, 0x42, 0x72, 0x04, 0xda, 0x2a,
	0xf0, 0xef, 0x57, 0x48, 0x0d, 0x09, 0xe4, 0x91, 0xf8, 0xb0, 0x65, 0xc8, 0x91, 0xee, 0x3d, 0xfd,
	0x61, 0xa3, 0x90, 0x23, 0x93, 0x6c, 0xf2, 0xba, 0x98, 0x93, 0x50, 0x90, 0x97, 0xb2, 0xab, 0x70,
	0x20, 0x98, 0xc5, 0x4e, 0x12, 0x9b, 0x41, 0x4d, 0x1c, 0x11, 0x01, 0xd0, 0xa6, 0x33, 0x36, 0x18,
	0x5f, 0x9a, 0x15, 0xd2, 0x84, 0xc6, 0x60, 0x3c, 0x73, 0x2f, 0x5d, 0x66, 0x2a, 0xc4, 0x80, 0xfa,
	0xc5, 0x70, 0xd2, 0x9b, 0x99, 0xaa, 0xe0, 0x9c, 0x4f, 0xe6, 0xfd, 0xa1, 0x6b, 0xd6, 0x05, 0xa7,
	0x3f, 0x99, 0x0c, 0xdd, 0xde, 0xd8, 0xac, 0x92, 0x7d, 0x30, 0x66, 0x83, 0x91, 0x3b, 0x9d, 0xf5,
	0x46, 0xef, 0xcd, 0x9a, 0xed, 0x40, 0x4d, 0xec, 0x8e, 0xec, 0x81, 0x3e, 0x9e, 0x0f, 0x87, 0x3d,
	0x51, 0x51, 0x11, 0x11, 0x73, 0xaf, 0xe7, 0x03, 0xe6, 0x9e, 0x9b, 0x8a, 0xe8, 0x35, 0x1f, 0x0f,
	0xae, 0xe7, 0xae, 0xa9, 0xda, 0x0c, 0x60, 0xb3, 0xb5, 0x6d, 0x12, 0x49, 0x22, 0x5c, 0x14, 0x12,
	0x11, 0xeb, 0xa7, 0x25, 0x62, 0x9f, 0x81, 0xce, 0x30, 0x89, 0xc2, 0x20, 0x41, 0xf2, 0x1f, 0x68,
	0x99, 0x9a, 0x73, 0xcd, 0x92, 0xe2, 0x78, 0xe2, 0x68, 0xe1, 0x4c, 0x25, 0xc2, 0x72, 0x86, 0xfd,
	0xa3, 0x0a, 0x64, 0xb4, 0x4a, 0xbd, 0x14, 0x33, 0x09, 0xe3, 0xfd, 0x0a, 0x93, 0x74, 0xad, 0x53,
	0xa5, 0xa4, 0xd3, 0x09, 0x18, 0x4b, 0xc1, 0xf4, 0xc3, 0x20, 0xa1, 0xaa, 0xec, 0xfc, 0x72, 0xfb,
	0xc1, 0xff, 0xd9, 0x30, 0x4b, 0xf9, 0x61, 0xc0, 0x36, 0x3d, 0x5a, 0x3f, 0x55, 0xd0, 0x8b, 0x3c,
	0x39, 0x2c, 0x4f, 0xbc, 0xaa, 0xe4, 0x33, 0xed, 0xc7, 0x1f, 0xae, 0xe6, 0xe0, 0x23, 0x77, 0x9c,
	0xe5, 0xee, 0xa8, 0x3e, 0xcf, 0x1d, 0xa2, 0xb7, 0xf4, 0xc7, 0x71, 0xc9, 0x1f, 0xc2, 0x59, 0xfa,
	0x55, 0xa5, 0xe4, 0x90, 0xe3, 0x92, 0xfc, 0xeb, 0x1b, 0x34, 0x37, 0x00, 0x5d, 0x0b, 0x5a, 0xcb,
	0xb1, 0x42, 0xd2, 0xad, 0x8d, 0x09, 0x1a, 0x39, 0xb4, 0xb6, 0xc1, 0x3b, 0x30, 0x3c, 0xce, 0x6f,
	0x64, 0x28, 0x2d, 0xf2, 0x0c, 0xe9, 0x8a, 0xb1, 0x1e, 0xcf, 0xd6, 0x7d, 0x00, 0xbd, 0x38, 0xbe,
	0xee, 0x2f, 0x05, 0xb4, 0xa9, 0x64, 0x93, 0x37, 0x60, 0x5c, 0x62, 0x9a, 0xbf, 0x58, 0x7f, 0x95,
	0x6f, 0xdb, 0x15, 0xaf, 0x5d, 0xeb, 0x78, 0xc7, 0x90, 0x84, 0x0c, 0xc0, 0xe8, 0x71, 0x9e, 0x07,
	0x3b, 0xa9, 0xad, 0x7f, 0xb7, 0xa3, 0x6b, 0xe5, 0x7d, 0x80, 0x66, 0xe9, 0xee, 0x49, 0xfb, 0xb9,
	0xf2, 0x78, 0xaa, 0x71, 0xbf, 0xfb, 0xf1, 0xc5, 0x22, 0xe4, 0x58, 0xb0, 0x16, 0xe1, 0xb2, 0x23,
	0x5f, 0xea, 0x5b, 0x0c, 0xfe, 0xbf, 0x0d, 0x3b, 0xdb, 0xfe, 0x18, 0x3e, 0x69, 0x92, 0x71, 0xfa,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x2b, 0xeb, 0x7f, 0x37, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SchemaClient is the client API for Schema service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchemaClient interface {
	// GetFields returns the fields in the schema.
	GetFields(ctx context.Context, in *rpc.Empty, opts ...grpc.CallOption) (*Fields, error)
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

func (c *schemaClient) GetFields(ctx context.Context, in *rpc.Empty, opts ...grpc.CallOption) (*Fields, error) {
	out := new(Fields)
	err := c.cc.Invoke(ctx, "/sajari.engine.schema.Schema/GetFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaClient) AddFields(ctx context.Context, in *Fields, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sajari.engine.schema.Schema/AddFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaClient) MutateField(ctx context.Context, in *MutateFieldRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sajari.engine.schema.Schema/MutateField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServer is the server API for Schema service.
type SchemaServer interface {
	// GetFields returns the fields in the schema.
	GetFields(context.Context, *rpc.Empty) (*Fields, error)
	// AddFields adds new fields to the schema.
	AddFields(context.Context, *Fields) (*Response, error)
	// MutateField mutates a field in the schema.
	MutateField(context.Context, *MutateFieldRequest) (*Response, error)
}

// UnimplementedSchemaServer can be embedded to have forward compatible implementations.
type UnimplementedSchemaServer struct {
}

func (*UnimplementedSchemaServer) GetFields(ctx context.Context, req *rpc.Empty) (*Fields, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFields not implemented")
}
func (*UnimplementedSchemaServer) AddFields(ctx context.Context, req *Fields) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFields not implemented")
}
func (*UnimplementedSchemaServer) MutateField(ctx context.Context, req *MutateFieldRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateField not implemented")
}

func RegisterSchemaServer(s *grpc.Server, srv SchemaServer) {
	s.RegisterService(&_Schema_serviceDesc, srv)
}

func _Schema_GetFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.Empty)
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
		return srv.(SchemaServer).GetFields(ctx, req.(*rpc.Empty))
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
