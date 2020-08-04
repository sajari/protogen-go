// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sajari/engine/v2/schema.proto

package enginepb

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

// Type represents the underlying data type of the field. Default is a string.
type Field_Type int32

const (
	// String values.
	Field_STRING Field_Type = 0
	// Integer values (64-bit).
	Field_INTEGER Field_Type = 1
	// Floating point values (32-bit).
	Field_FLOAT Field_Type = 2
	// Double floating point values (64-bit).
	Field_DOUBLE Field_Type = 3
	// Boolean values.
	Field_BOOLEAN Field_Type = 4
	// Timestamp values.
	Field_TIMESTAMP Field_Type = 5
)

var Field_Type_name = map[int32]string{
	0: "STRING",
	1: "INTEGER",
	2: "FLOAT",
	3: "DOUBLE",
	4: "BOOLEAN",
	5: "TIMESTAMP",
}

var Field_Type_value = map[string]int32{
	"STRING":    0,
	"INTEGER":   1,
	"FLOAT":     2,
	"DOUBLE":    3,
	"BOOLEAN":   4,
	"TIMESTAMP": 5,
}

func (x Field_Type) String() string {
	return proto.EnumName(Field_Type_name, int32(x))
}

func (Field_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13fad1e859dad9f5, []int{4, 0}
}

// Mode is an enumeration of modes for a field.
type Field_Mode int32

const (
	// Nullable fields do not need to be specified.
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
	return fileDescriptor_13fad1e859dad9f5, []int{4, 1}
}

// ListFieldsRequest.
type ListFieldsRequest struct {
	// Number of fields to return per page.  The default page size is 100.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The `next_page_token` value returned from a previous List request, if any.
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFieldsRequest) Reset()         { *m = ListFieldsRequest{} }
func (m *ListFieldsRequest) String() string { return proto.CompactTextString(m) }
func (*ListFieldsRequest) ProtoMessage()    {}
func (*ListFieldsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_13fad1e859dad9f5, []int{0}
}

func (m *ListFieldsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFieldsRequest.Unmarshal(m, b)
}
func (m *ListFieldsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFieldsRequest.Marshal(b, m, deterministic)
}
func (m *ListFieldsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFieldsRequest.Merge(m, src)
}
func (m *ListFieldsRequest) XXX_Size() int {
	return xxx_messageInfo_ListFieldsRequest.Size(m)
}
func (m *ListFieldsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFieldsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFieldsRequest proto.InternalMessageInfo

func (m *ListFieldsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListFieldsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// ListFieldResponse.
type ListFieldsResponse struct {
	// Fields.
	Fields []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	// Maximum number of fields to return.
	TotalSize int32 `protobuf:"varint,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	// Next page token.
	NextPageToken        string   `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFieldsResponse) Reset()         { *m = ListFieldsResponse{} }
func (m *ListFieldsResponse) String() string { return proto.CompactTextString(m) }
func (*ListFieldsResponse) ProtoMessage()    {}
func (*ListFieldsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13fad1e859dad9f5, []int{1}
}

func (m *ListFieldsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFieldsResponse.Unmarshal(m, b)
}
func (m *ListFieldsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFieldsResponse.Marshal(b, m, deterministic)
}
func (m *ListFieldsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFieldsResponse.Merge(m, src)
}
func (m *ListFieldsResponse) XXX_Size() int {
	return xxx_messageInfo_ListFieldsResponse.Size(m)
}
func (m *ListFieldsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFieldsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFieldsResponse proto.InternalMessageInfo

func (m *ListFieldsResponse) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *ListFieldsResponse) GetTotalSize() int32 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *ListFieldsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// CreateFieldRequest.
type CreateFieldRequest struct {
	// Field to create.
	Field                *Field   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFieldRequest) Reset()         { *m = CreateFieldRequest{} }
func (m *CreateFieldRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFieldRequest) ProtoMessage()    {}
func (*CreateFieldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_13fad1e859dad9f5, []int{2}
}

func (m *CreateFieldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFieldRequest.Unmarshal(m, b)
}
func (m *CreateFieldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFieldRequest.Marshal(b, m, deterministic)
}
func (m *CreateFieldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFieldRequest.Merge(m, src)
}
func (m *CreateFieldRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFieldRequest.Size(m)
}
func (m *CreateFieldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFieldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFieldRequest proto.InternalMessageInfo

func (m *CreateFieldRequest) GetField() *Field {
	if m != nil {
		return m.Field
	}
	return nil
}

type CreateFieldResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFieldResponse) Reset()         { *m = CreateFieldResponse{} }
func (m *CreateFieldResponse) String() string { return proto.CompactTextString(m) }
func (*CreateFieldResponse) ProtoMessage()    {}
func (*CreateFieldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13fad1e859dad9f5, []int{3}
}

func (m *CreateFieldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFieldResponse.Unmarshal(m, b)
}
func (m *CreateFieldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFieldResponse.Marshal(b, m, deterministic)
}
func (m *CreateFieldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFieldResponse.Merge(m, src)
}
func (m *CreateFieldResponse) XXX_Size() int {
	return xxx_messageInfo_CreateFieldResponse.Size(m)
}
func (m *CreateFieldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFieldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFieldResponse proto.InternalMessageInfo

// Field defines the properties of a field in the schema.
type Field struct {
	// Name of the field.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the field.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Type of the field.
	Type Field_Type `protobuf:"varint,3,opt,name=type,proto3,enum=sajari.engine.v2.Field_Type" json:"type,omitempty"`
	// Mode of the field.
	Mode Field_Mode `protobuf:"varint,4,opt,name=mode,proto3,enum=sajari.engine.v2.Field_Mode" json:"mode,omitempty"`
	// Repeated sets that fields are repeated values (i.e. a list of values).
	Repeated bool `protobuf:"varint,5,opt,name=repeated,proto3" json:"repeated,omitempty"`
	// Repeated length sets the required length for the repeated field.
	RepeatedLen int32 `protobuf:"varint,6,opt,name=repeated_len,json=repeatedLen,proto3" json:"repeated_len,omitempty"`
	// Indexes is a list of indexes defined for the field.
	Indexes              []*FieldIndex `protobuf:"bytes,7,rep,name=indexes,proto3" json:"indexes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_13fad1e859dad9f5, []int{4}
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

func (m *Field) GetMode() Field_Mode {
	if m != nil {
		return m.Mode
	}
	return Field_NULLABLE
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

func (m *Field) GetIndexes() []*FieldIndex {
	if m != nil {
		return m.Indexes
	}
	return nil
}

// FieldIndex
type FieldIndex struct {
	// Specification defining how the index was created.
	Spec string `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	// Description of the index.
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldIndex) Reset()         { *m = FieldIndex{} }
func (m *FieldIndex) String() string { return proto.CompactTextString(m) }
func (*FieldIndex) ProtoMessage()    {}
func (*FieldIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_13fad1e859dad9f5, []int{5}
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

// MutateFieldRequest is a message type passed to the MutateField method.
type MutateFieldRequest struct {
	// Name of the field to mutate.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Mutation to apply to the field.
	Mutation             *MutateFieldRequest_Mutation `protobuf:"bytes,2,opt,name=mutation,proto3" json:"mutation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateFieldRequest) Reset()         { *m = MutateFieldRequest{} }
func (m *MutateFieldRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFieldRequest) ProtoMessage()    {}
func (*MutateFieldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_13fad1e859dad9f5, []int{6}
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

func (m *MutateFieldRequest) GetMutation() *MutateFieldRequest_Mutation {
	if m != nil {
		return m.Mutation
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
	//	*MutateFieldRequest_Mutation_Mode
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
	return fileDescriptor_13fad1e859dad9f5, []int{6, 0}
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
	Type Field_Type `protobuf:"varint,3,opt,name=type,proto3,enum=sajari.engine.v2.Field_Type,oneof"`
}

type MutateFieldRequest_Mutation_Repeated struct {
	Repeated bool `protobuf:"varint,4,opt,name=repeated,proto3,oneof"`
}

type MutateFieldRequest_Mutation_Mode struct {
	Mode Field_Mode `protobuf:"varint,5,opt,name=mode,proto3,enum=sajari.engine.v2.Field_Mode,oneof"`
}

type MutateFieldRequest_Mutation_AddIndex struct {
	AddIndex *FieldIndex `protobuf:"bytes,6,opt,name=add_index,json=addIndex,proto3,oneof"`
}

func (*MutateFieldRequest_Mutation_Name) isMutateFieldRequest_Mutation_Mutation() {}

func (*MutateFieldRequest_Mutation_Description) isMutateFieldRequest_Mutation_Mutation() {}

func (*MutateFieldRequest_Mutation_Type) isMutateFieldRequest_Mutation_Mutation() {}

func (*MutateFieldRequest_Mutation_Repeated) isMutateFieldRequest_Mutation_Mutation() {}

func (*MutateFieldRequest_Mutation_Mode) isMutateFieldRequest_Mutation_Mutation() {}

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

func (m *MutateFieldRequest_Mutation) GetMode() Field_Mode {
	if x, ok := m.GetMutation().(*MutateFieldRequest_Mutation_Mode); ok {
		return x.Mode
	}
	return Field_NULLABLE
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
		(*MutateFieldRequest_Mutation_Mode)(nil),
		(*MutateFieldRequest_Mutation_AddIndex)(nil),
	}
}

type MutateFieldResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFieldResponse) Reset()         { *m = MutateFieldResponse{} }
func (m *MutateFieldResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFieldResponse) ProtoMessage()    {}
func (*MutateFieldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13fad1e859dad9f5, []int{7}
}

func (m *MutateFieldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFieldResponse.Unmarshal(m, b)
}
func (m *MutateFieldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFieldResponse.Marshal(b, m, deterministic)
}
func (m *MutateFieldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFieldResponse.Merge(m, src)
}
func (m *MutateFieldResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFieldResponse.Size(m)
}
func (m *MutateFieldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFieldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFieldResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("sajari.engine.v2.Field_Type", Field_Type_name, Field_Type_value)
	proto.RegisterEnum("sajari.engine.v2.Field_Mode", Field_Mode_name, Field_Mode_value)
	proto.RegisterType((*ListFieldsRequest)(nil), "sajari.engine.v2.ListFieldsRequest")
	proto.RegisterType((*ListFieldsResponse)(nil), "sajari.engine.v2.ListFieldsResponse")
	proto.RegisterType((*CreateFieldRequest)(nil), "sajari.engine.v2.CreateFieldRequest")
	proto.RegisterType((*CreateFieldResponse)(nil), "sajari.engine.v2.CreateFieldResponse")
	proto.RegisterType((*Field)(nil), "sajari.engine.v2.Field")
	proto.RegisterType((*FieldIndex)(nil), "sajari.engine.v2.FieldIndex")
	proto.RegisterType((*MutateFieldRequest)(nil), "sajari.engine.v2.MutateFieldRequest")
	proto.RegisterType((*MutateFieldRequest_Mutation)(nil), "sajari.engine.v2.MutateFieldRequest.Mutation")
	proto.RegisterType((*MutateFieldResponse)(nil), "sajari.engine.v2.MutateFieldResponse")
}

func init() { proto.RegisterFile("sajari/engine/v2/schema.proto", fileDescriptor_13fad1e859dad9f5) }

var fileDescriptor_13fad1e859dad9f5 = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4f, 0x6f, 0x9b, 0x4e,
	0x10, 0x35, 0xb6, 0x71, 0x60, 0x9c, 0xfc, 0x7e, 0x74, 0xdb, 0xaa, 0xc8, 0x4d, 0x24, 0x97, 0xfe,
	0x91, 0x2f, 0xc1, 0x15, 0x55, 0xdb, 0x43, 0x4e, 0x76, 0x42, 0x12, 0x24, 0x6c, 0x27, 0x6b, 0xac,
	0x4a, 0xb9, 0x58, 0xc4, 0x4c, 0x5d, 0xda, 0x18, 0xa8, 0x21, 0x51, 0x92, 0x6f, 0x50, 0xa9, 0xb7,
	0x9e, 0x2a, 0xf5, 0xc3, 0x56, 0xbb, 0x98, 0x98, 0xc4, 0x4e, 0xdc, 0xdb, 0x32, 0xbc, 0x79, 0xf3,
	0x76, 0xe6, 0x31, 0xc0, 0x56, 0xec, 0x7e, 0x75, 0xa7, 0x7e, 0x13, 0x83, 0xb1, 0x1f, 0x60, 0xf3,
	0xc2, 0x68, 0xc6, 0xa3, 0x2f, 0x38, 0x71, 0xf5, 0x68, 0x1a, 0x26, 0x21, 0x51, 0xd2, 0xd7, 0x7a,
	0xfa, 0x5a, 0xbf, 0x30, 0xb4, 0x1e, 0x3c, 0xb2, 0xfd, 0x38, 0xd9, 0xf7, 0xf1, 0xcc, 0x8b, 0x29,
	0x7e, 0x3f, 0xc7, 0x38, 0x21, 0xcf, 0x41, 0x8e, 0xdc, 0x31, 0x0e, 0x63, 0xff, 0x1a, 0x55, 0xa1,
	0x2e, 0x34, 0x44, 0x2a, 0xb1, 0x40, 0xdf, 0xbf, 0x46, 0xb2, 0x05, 0xc0, 0x5f, 0x26, 0xe1, 0x37,
	0x0c, 0xd4, 0x62, 0x5d, 0x68, 0xc8, 0x94, 0xc3, 0x1d, 0x16, 0xd0, 0x7e, 0x0a, 0x40, 0xf2, 0x8c,
	0x71, 0x14, 0x06, 0x31, 0x92, 0x26, 0x54, 0x3e, 0xf3, 0x88, 0x2a, 0xd4, 0x4b, 0x8d, 0xaa, 0xf1,
	0x4c, 0xbf, 0x2b, 0x45, 0xe7, 0x19, 0x74, 0x06, 0x63, 0x65, 0x92, 0x30, 0x71, 0xcf, 0x52, 0x11,
	0x45, 0x2e, 0x42, 0xe6, 0x11, 0xae, 0xe2, 0x0d, 0xfc, 0x1f, 0xe0, 0x65, 0x32, 0xcc, 0x49, 0x29,
	0x71, 0x29, 0x1b, 0x2c, 0x7c, 0x74, 0x23, 0x67, 0x17, 0xc8, 0xee, 0x14, 0xdd, 0x04, 0x53, 0xf6,
	0xd9, 0x05, 0xb7, 0x41, 0xe4, 0x65, 0xf8, 0xe5, 0x1e, 0x10, 0x93, 0xa2, 0xb4, 0xa7, 0xf0, 0xf8,
	0x16, 0x49, 0x7a, 0x27, 0xed, 0x4f, 0x09, 0x44, 0x1e, 0x21, 0x04, 0xca, 0x81, 0x3b, 0x49, 0x7b,
	0x25, 0x53, 0x7e, 0x26, 0x75, 0xa8, 0x7a, 0x18, 0x8f, 0xa6, 0x7e, 0x94, 0xf8, 0x61, 0xd6, 0xa8,
	0x7c, 0x88, 0xbc, 0x85, 0x72, 0x72, 0x15, 0x21, 0x17, 0xfe, 0x9f, 0xb1, 0x79, 0x8f, 0x08, 0xdd,
	0xb9, 0x8a, 0x90, 0x72, 0x24, 0xcb, 0x98, 0x84, 0x1e, 0xaa, 0xe5, 0x87, 0x33, 0x3a, 0xa1, 0x87,
	0x94, 0x23, 0x49, 0x0d, 0xa4, 0x29, 0x46, 0x4c, 0xbb, 0xa7, 0x8a, 0x75, 0xa1, 0x21, 0xd1, 0x9b,
	0x67, 0xf2, 0x02, 0xd6, 0xb3, 0xf3, 0xf0, 0x0c, 0x03, 0xb5, 0xc2, 0x9b, 0x5c, 0xcd, 0x62, 0x36,
	0x06, 0xe4, 0x03, 0xac, 0xf9, 0x81, 0x87, 0x97, 0x18, 0xab, 0x6b, 0x7c, 0x6e, 0xf7, 0xd5, 0xb4,
	0x18, 0x8a, 0x66, 0x60, 0x8d, 0x42, 0x99, 0xc9, 0x26, 0x00, 0x95, 0xbe, 0x43, 0xad, 0xee, 0x81,
	0x52, 0x20, 0x55, 0x58, 0xb3, 0xba, 0x8e, 0x79, 0x60, 0x52, 0x45, 0x20, 0x32, 0x88, 0xfb, 0x76,
	0xaf, 0xe5, 0x28, 0x45, 0x86, 0xd9, 0xeb, 0x0d, 0xda, 0xb6, 0xa9, 0x94, 0x18, 0xa6, 0xdd, 0xeb,
	0xd9, 0x66, 0xab, 0xab, 0x94, 0xc9, 0x06, 0xc8, 0x8e, 0xd5, 0x31, 0xfb, 0x4e, 0xab, 0x73, 0xa4,
	0x88, 0x9a, 0x0e, 0x65, 0x76, 0x31, 0xb2, 0x0e, 0x52, 0x77, 0x60, 0xdb, 0x2d, 0x96, 0x51, 0x60,
	0x4f, 0xd4, 0x3c, 0x1e, 0x58, 0xd4, 0xdc, 0x53, 0x04, 0xc6, 0x35, 0xe8, 0x5a, 0xc7, 0x03, 0x53,
	0x29, 0x6a, 0x6d, 0x80, 0xb9, 0x34, 0x36, 0xa2, 0x38, 0xc2, 0x51, 0x36, 0x22, 0x76, 0x5e, 0x3d,
	0x22, 0xed, 0x47, 0x09, 0x48, 0xe7, 0x3c, 0xb9, 0xeb, 0x9f, 0x65, 0xf3, 0xb6, 0x40, 0x9a, 0x30,
	0x64, 0xc6, 0x54, 0x35, 0xb6, 0x17, 0x7b, 0xb5, 0xc8, 0x95, 0x86, 0xfc, 0x30, 0xa0, 0x37, 0xe9,
	0xb5, 0x5f, 0x45, 0x90, 0xb2, 0x30, 0x79, 0x92, 0xaf, 0x75, 0x58, 0x98, 0x55, 0xd3, 0x96, 0x48,
	0x3f, 0x2c, 0xdc, 0xf6, 0x97, 0xf1, 0xef, 0xfe, 0x62, 0xbc, 0xdc, 0x61, 0x9b, 0x39, 0xbf, 0x30,
	0x97, 0x49, 0x87, 0x85, 0x9c, 0x63, 0x8c, 0x99, 0xff, 0xc4, 0xd5, 0xfe, 0x63, 0x8c, 0xdc, 0x81,
	0x3b, 0x20, 0xbb, 0x9e, 0x37, 0xe4, 0xce, 0xe0, 0x16, 0x5b, 0x61, 0x22, 0x56, 0xd0, 0xf5, 0xd2,
	0x73, 0x1b, 0xe6, 0x4d, 0x65, 0x5f, 0xe1, 0xad, 0xf6, 0xa5, 0x5f, 0xa1, 0xf1, 0xbb, 0x08, 0x95,
	0x3e, 0x5f, 0x72, 0xe4, 0x13, 0xc0, 0x7c, 0xf5, 0x90, 0x97, 0x8b, 0x55, 0x16, 0x56, 0x5d, 0xed,
	0xd5, 0xc3, 0xa0, 0xd9, 0xf6, 0x3a, 0x81, 0x6a, 0x6e, 0x01, 0x90, 0x25, 0x49, 0x8b, 0x4b, 0xa6,
	0xf6, 0x7a, 0x05, 0x6a, 0xce, 0x9d, 0xbb, 0xd6, 0x32, 0xee, 0x45, 0xd3, 0x2c, 0xe3, 0x5e, 0xd2,
	0x9b, 0xf6, 0xc7, 0x93, 0xf7, 0xa3, 0xd0, 0xc3, 0x0c, 0x3c, 0x0a, 0x27, 0x4d, 0xfe, 0x23, 0x18,
	0x63, 0xb0, 0x3d, 0x0e, 0x9b, 0x77, 0x7f, 0x16, 0x3b, 0xe9, 0x29, 0x3a, 0x3d, 0xad, 0x70, 0xd8,
	0xbb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x18, 0xa0, 0xfb, 0x50, 0x06, 0x00, 0x00,
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
	// ListFields returns the fields in the schema.
	ListFields(ctx context.Context, in *ListFieldsRequest, opts ...grpc.CallOption) (*ListFieldsResponse, error)
	// CreateField create a new field in the schema.
	CreateField(ctx context.Context, in *CreateFieldRequest, opts ...grpc.CallOption) (*CreateFieldResponse, error)
	// MutateField mutates a field in the schema.
	MutateField(ctx context.Context, in *MutateFieldRequest, opts ...grpc.CallOption) (*MutateFieldResponse, error)
}

type schemaClient struct {
	cc *grpc.ClientConn
}

func NewSchemaClient(cc *grpc.ClientConn) SchemaClient {
	return &schemaClient{cc}
}

func (c *schemaClient) ListFields(ctx context.Context, in *ListFieldsRequest, opts ...grpc.CallOption) (*ListFieldsResponse, error) {
	out := new(ListFieldsResponse)
	err := c.cc.Invoke(ctx, "/sajari.engine.v2.Schema/ListFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaClient) CreateField(ctx context.Context, in *CreateFieldRequest, opts ...grpc.CallOption) (*CreateFieldResponse, error) {
	out := new(CreateFieldResponse)
	err := c.cc.Invoke(ctx, "/sajari.engine.v2.Schema/CreateField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaClient) MutateField(ctx context.Context, in *MutateFieldRequest, opts ...grpc.CallOption) (*MutateFieldResponse, error) {
	out := new(MutateFieldResponse)
	err := c.cc.Invoke(ctx, "/sajari.engine.v2.Schema/MutateField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServer is the server API for Schema service.
type SchemaServer interface {
	// ListFields returns the fields in the schema.
	ListFields(context.Context, *ListFieldsRequest) (*ListFieldsResponse, error)
	// CreateField create a new field in the schema.
	CreateField(context.Context, *CreateFieldRequest) (*CreateFieldResponse, error)
	// MutateField mutates a field in the schema.
	MutateField(context.Context, *MutateFieldRequest) (*MutateFieldResponse, error)
}

// UnimplementedSchemaServer can be embedded to have forward compatible implementations.
type UnimplementedSchemaServer struct {
}

func (*UnimplementedSchemaServer) ListFields(ctx context.Context, req *ListFieldsRequest) (*ListFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFields not implemented")
}
func (*UnimplementedSchemaServer) CreateField(ctx context.Context, req *CreateFieldRequest) (*CreateFieldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateField not implemented")
}
func (*UnimplementedSchemaServer) MutateField(ctx context.Context, req *MutateFieldRequest) (*MutateFieldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateField not implemented")
}

func RegisterSchemaServer(s *grpc.Server, srv SchemaServer) {
	s.RegisterService(&_Schema_serviceDesc, srv)
}

func _Schema_ListFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServer).ListFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.engine.v2.Schema/ListFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServer).ListFields(ctx, req.(*ListFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schema_CreateField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServer).CreateField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.engine.v2.Schema/CreateField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServer).CreateField(ctx, req.(*CreateFieldRequest))
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
		FullMethod: "/sajari.engine.v2.Schema/MutateField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServer).MutateField(ctx, req.(*MutateFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Schema_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sajari.engine.v2.Schema",
	HandlerType: (*SchemaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFields",
			Handler:    _Schema_ListFields_Handler,
		},
		{
			MethodName: "CreateField",
			Handler:    _Schema_CreateField_Handler,
		},
		{
			MethodName: "MutateField",
			Handler:    _Schema_MutateField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sajari/engine/v2/schema.proto",
}