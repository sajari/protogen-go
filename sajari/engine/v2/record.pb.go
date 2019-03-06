// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sajari/engine/v2/record.proto

package enginepb

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

// Value is the representation for record field values.
type Value struct {
	// Types that are valid to be assigned to Value:
	//	*Value_Null
	//	*Value_Single
	//	*Value_Repeated_
	Value                isValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_17346a3c91f94d5c, []int{0}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Value interface {
	isValue_Value()
}

type Value_Null struct {
	Null bool `protobuf:"varint,1,opt,name=null,proto3,oneof"`
}

type Value_Single struct {
	Single string `protobuf:"bytes,2,opt,name=single,proto3,oneof"`
}

type Value_Repeated_ struct {
	Repeated *Value_Repeated `protobuf:"bytes,3,opt,name=repeated,proto3,oneof"`
}

func (*Value_Null) isValue_Value() {}

func (*Value_Single) isValue_Value() {}

func (*Value_Repeated_) isValue_Value() {}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetNull() bool {
	if x, ok := m.GetValue().(*Value_Null); ok {
		return x.Null
	}
	return false
}

func (m *Value) GetSingle() string {
	if x, ok := m.GetValue().(*Value_Single); ok {
		return x.Single
	}
	return ""
}

func (m *Value) GetRepeated() *Value_Repeated {
	if x, ok := m.GetValue().(*Value_Repeated_); ok {
		return x.Repeated
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_Null)(nil),
		(*Value_Single)(nil),
		(*Value_Repeated_)(nil),
	}
}

// Repeated values
type Value_Repeated struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value_Repeated) Reset()         { *m = Value_Repeated{} }
func (m *Value_Repeated) String() string { return proto.CompactTextString(m) }
func (*Value_Repeated) ProtoMessage()    {}
func (*Value_Repeated) Descriptor() ([]byte, []int) {
	return fileDescriptor_17346a3c91f94d5c, []int{0, 0}
}

func (m *Value_Repeated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value_Repeated.Unmarshal(m, b)
}
func (m *Value_Repeated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value_Repeated.Marshal(b, m, deterministic)
}
func (m *Value_Repeated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value_Repeated.Merge(m, src)
}
func (m *Value_Repeated) XXX_Size() int {
	return xxx_messageInfo_Value_Repeated.Size(m)
}
func (m *Value_Repeated) XXX_DiscardUnknown() {
	xxx_messageInfo_Value_Repeated.DiscardUnknown(m)
}

var xxx_messageInfo_Value_Repeated proto.InternalMessageInfo

func (m *Value_Repeated) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// Key is a key-value pair that uniquely determines a record in a collection.
// Any unique field in a collection can be used to create a key.
type Key struct {
	// Field is the meta field (must be a unique field).
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// Value is the identifying value.
	Value                *Value   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_17346a3c91f94d5c, []int{1}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *Key) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

// Record is a single record made up of key-value pairs.
type Record struct {
	// A map of key-value pairs.
	Values               map[string]*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_17346a3c91f94d5c, []int{2}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetValues() map[string]*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// GetRequest.
type GetRecordRequest struct {
	Key                  *Key     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRecordRequest) Reset()         { *m = GetRecordRequest{} }
func (m *GetRecordRequest) String() string { return proto.CompactTextString(m) }
func (*GetRecordRequest) ProtoMessage()    {}
func (*GetRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17346a3c91f94d5c, []int{3}
}

func (m *GetRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecordRequest.Unmarshal(m, b)
}
func (m *GetRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecordRequest.Marshal(b, m, deterministic)
}
func (m *GetRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecordRequest.Merge(m, src)
}
func (m *GetRecordRequest) XXX_Size() int {
	return xxx_messageInfo_GetRecordRequest.Size(m)
}
func (m *GetRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecordRequest proto.InternalMessageInfo

func (m *GetRecordRequest) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

// GetResponse is returned from Get.
type GetRecordResponse struct {
	Record               *Record  `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRecordResponse) Reset()         { *m = GetRecordResponse{} }
func (m *GetRecordResponse) String() string { return proto.CompactTextString(m) }
func (*GetRecordResponse) ProtoMessage()    {}
func (*GetRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17346a3c91f94d5c, []int{4}
}

func (m *GetRecordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecordResponse.Unmarshal(m, b)
}
func (m *GetRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecordResponse.Marshal(b, m, deterministic)
}
func (m *GetRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecordResponse.Merge(m, src)
}
func (m *GetRecordResponse) XXX_Size() int {
	return xxx_messageInfo_GetRecordResponse.Size(m)
}
func (m *GetRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecordResponse proto.InternalMessageInfo

func (m *GetRecordResponse) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

// DeleteRequest.
type DeleteRecordRequest struct {
	Key                  *Key     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRecordRequest) Reset()         { *m = DeleteRecordRequest{} }
func (m *DeleteRecordRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRecordRequest) ProtoMessage()    {}
func (*DeleteRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17346a3c91f94d5c, []int{5}
}

func (m *DeleteRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRecordRequest.Unmarshal(m, b)
}
func (m *DeleteRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRecordRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRecordRequest.Merge(m, src)
}
func (m *DeleteRecordRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRecordRequest.Size(m)
}
func (m *DeleteRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRecordRequest proto.InternalMessageInfo

func (m *DeleteRecordRequest) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

// DeleteResponse is returned from Delete.
type DeleteRecordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRecordResponse) Reset()         { *m = DeleteRecordResponse{} }
func (m *DeleteRecordResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteRecordResponse) ProtoMessage()    {}
func (*DeleteRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17346a3c91f94d5c, []int{6}
}

func (m *DeleteRecordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRecordResponse.Unmarshal(m, b)
}
func (m *DeleteRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRecordResponse.Marshal(b, m, deterministic)
}
func (m *DeleteRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRecordResponse.Merge(m, src)
}
func (m *DeleteRecordResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteRecordResponse.Size(m)
}
func (m *DeleteRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRecordResponse proto.InternalMessageInfo

type MutateRecordRequest struct {
	// Key which uniquely identifies record.
	Key *Key `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// List of field mutations to apply.
	FieldMutations       []*MutateRecordRequest_FieldMutation `protobuf:"bytes,2,rep,name=field_mutations,json=fieldMutations,proto3" json:"field_mutations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *MutateRecordRequest) Reset()         { *m = MutateRecordRequest{} }
func (m *MutateRecordRequest) String() string { return proto.CompactTextString(m) }
func (*MutateRecordRequest) ProtoMessage()    {}
func (*MutateRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17346a3c91f94d5c, []int{7}
}

func (m *MutateRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateRecordRequest.Unmarshal(m, b)
}
func (m *MutateRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateRecordRequest.Marshal(b, m, deterministic)
}
func (m *MutateRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateRecordRequest.Merge(m, src)
}
func (m *MutateRecordRequest) XXX_Size() int {
	return xxx_messageInfo_MutateRecordRequest.Size(m)
}
func (m *MutateRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateRecordRequest proto.InternalMessageInfo

func (m *MutateRecordRequest) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MutateRecordRequest) GetFieldMutations() []*MutateRecordRequest_FieldMutation {
	if m != nil {
		return m.FieldMutations
	}
	return nil
}

// MutateField defines a mutation of a field.
type MutateRecordRequest_FieldMutation struct {
	// Field is the name of the field to mutate.
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// Types that are valid to be assigned to Mutation:
	//	*MutateRecordRequest_FieldMutation_Set
	//	*MutateRecordRequest_FieldMutation_Increment
	//	*MutateRecordRequest_FieldMutation_Append
	Mutation             isMutateRecordRequest_FieldMutation_Mutation `protobuf_oneof:"mutation"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *MutateRecordRequest_FieldMutation) Reset()         { *m = MutateRecordRequest_FieldMutation{} }
func (m *MutateRecordRequest_FieldMutation) String() string { return proto.CompactTextString(m) }
func (*MutateRecordRequest_FieldMutation) ProtoMessage()    {}
func (*MutateRecordRequest_FieldMutation) Descriptor() ([]byte, []int) {
	return fileDescriptor_17346a3c91f94d5c, []int{7, 0}
}

func (m *MutateRecordRequest_FieldMutation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateRecordRequest_FieldMutation.Unmarshal(m, b)
}
func (m *MutateRecordRequest_FieldMutation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateRecordRequest_FieldMutation.Marshal(b, m, deterministic)
}
func (m *MutateRecordRequest_FieldMutation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateRecordRequest_FieldMutation.Merge(m, src)
}
func (m *MutateRecordRequest_FieldMutation) XXX_Size() int {
	return xxx_messageInfo_MutateRecordRequest_FieldMutation.Size(m)
}
func (m *MutateRecordRequest_FieldMutation) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateRecordRequest_FieldMutation.DiscardUnknown(m)
}

var xxx_messageInfo_MutateRecordRequest_FieldMutation proto.InternalMessageInfo

func (m *MutateRecordRequest_FieldMutation) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type isMutateRecordRequest_FieldMutation_Mutation interface {
	isMutateRecordRequest_FieldMutation_Mutation()
}

type MutateRecordRequest_FieldMutation_Set struct {
	Set *Value `protobuf:"bytes,2,opt,name=set,proto3,oneof"`
}

type MutateRecordRequest_FieldMutation_Increment struct {
	Increment *Value `protobuf:"bytes,3,opt,name=increment,proto3,oneof"`
}

type MutateRecordRequest_FieldMutation_Append struct {
	Append *Value `protobuf:"bytes,4,opt,name=append,proto3,oneof"`
}

func (*MutateRecordRequest_FieldMutation_Set) isMutateRecordRequest_FieldMutation_Mutation() {}

func (*MutateRecordRequest_FieldMutation_Increment) isMutateRecordRequest_FieldMutation_Mutation() {}

func (*MutateRecordRequest_FieldMutation_Append) isMutateRecordRequest_FieldMutation_Mutation() {}

func (m *MutateRecordRequest_FieldMutation) GetMutation() isMutateRecordRequest_FieldMutation_Mutation {
	if m != nil {
		return m.Mutation
	}
	return nil
}

func (m *MutateRecordRequest_FieldMutation) GetSet() *Value {
	if x, ok := m.GetMutation().(*MutateRecordRequest_FieldMutation_Set); ok {
		return x.Set
	}
	return nil
}

func (m *MutateRecordRequest_FieldMutation) GetIncrement() *Value {
	if x, ok := m.GetMutation().(*MutateRecordRequest_FieldMutation_Increment); ok {
		return x.Increment
	}
	return nil
}

func (m *MutateRecordRequest_FieldMutation) GetAppend() *Value {
	if x, ok := m.GetMutation().(*MutateRecordRequest_FieldMutation_Append); ok {
		return x.Append
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MutateRecordRequest_FieldMutation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MutateRecordRequest_FieldMutation_Set)(nil),
		(*MutateRecordRequest_FieldMutation_Increment)(nil),
		(*MutateRecordRequest_FieldMutation_Append)(nil),
	}
}

type MutateRecordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateRecordResponse) Reset()         { *m = MutateRecordResponse{} }
func (m *MutateRecordResponse) String() string { return proto.CompactTextString(m) }
func (*MutateRecordResponse) ProtoMessage()    {}
func (*MutateRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17346a3c91f94d5c, []int{8}
}

func (m *MutateRecordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateRecordResponse.Unmarshal(m, b)
}
func (m *MutateRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateRecordResponse.Marshal(b, m, deterministic)
}
func (m *MutateRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateRecordResponse.Merge(m, src)
}
func (m *MutateRecordResponse) XXX_Size() int {
	return xxx_messageInfo_MutateRecordResponse.Size(m)
}
func (m *MutateRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateRecordResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Value)(nil), "sajari.engine.v2.Value")
	proto.RegisterType((*Value_Repeated)(nil), "sajari.engine.v2.Value.Repeated")
	proto.RegisterType((*Key)(nil), "sajari.engine.v2.Key")
	proto.RegisterType((*Record)(nil), "sajari.engine.v2.Record")
	proto.RegisterMapType((map[string]*Value)(nil), "sajari.engine.v2.Record.ValuesEntry")
	proto.RegisterType((*GetRecordRequest)(nil), "sajari.engine.v2.GetRecordRequest")
	proto.RegisterType((*GetRecordResponse)(nil), "sajari.engine.v2.GetRecordResponse")
	proto.RegisterType((*DeleteRecordRequest)(nil), "sajari.engine.v2.DeleteRecordRequest")
	proto.RegisterType((*DeleteRecordResponse)(nil), "sajari.engine.v2.DeleteRecordResponse")
	proto.RegisterType((*MutateRecordRequest)(nil), "sajari.engine.v2.MutateRecordRequest")
	proto.RegisterType((*MutateRecordRequest_FieldMutation)(nil), "sajari.engine.v2.MutateRecordRequest.FieldMutation")
	proto.RegisterType((*MutateRecordResponse)(nil), "sajari.engine.v2.MutateRecordResponse")
}

func init() { proto.RegisterFile("sajari/engine/v2/record.proto", fileDescriptor_17346a3c91f94d5c) }

var fileDescriptor_17346a3c91f94d5c = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0xc6, 0x24, 0x13, 0x2e, 0x65, 0x1b, 0x8a, 0x15, 0x09, 0x29, 0x32, 0xb7, 0x48,
	0xa8, 0x36, 0xb8, 0x42, 0x45, 0x14, 0xf5, 0x21, 0xa2, 0x10, 0x51, 0xf5, 0x65, 0x41, 0x3c, 0x20,
	0x10, 0x72, 0xe3, 0x69, 0x64, 0x70, 0xd6, 0xc6, 0xbb, 0x89, 0x94, 0x5f, 0x81, 0x57, 0xbe, 0x86,
	0x3f, 0xe1, 0x2f, 0x90, 0x77, 0x37, 0x8d, 0x13, 0x3b, 0x8d, 0xe8, 0xdb, 0x8e, 0x7d, 0xce, 0x99,
	0xe3, 0x39, 0xe3, 0x85, 0x7b, 0xdc, 0xff, 0xe6, 0xa7, 0xa1, 0x8b, 0x6c, 0x14, 0x32, 0x74, 0xa7,
	0x9e, 0x9b, 0xe2, 0x30, 0x4e, 0x03, 0x27, 0x49, 0x63, 0x11, 0x93, 0x6d, 0xf5, 0xda, 0x51, 0xaf,
	0x9d, 0xa9, 0x67, 0xff, 0x36, 0xa0, 0xfe, 0xd1, 0x8f, 0x26, 0x48, 0xda, 0xb0, 0xc5, 0x26, 0x51,
	0x64, 0x19, 0x5d, 0xa3, 0xd7, 0x18, 0x54, 0xa8, 0xac, 0x88, 0x05, 0x26, 0x0f, 0xd9, 0x28, 0x42,
	0xab, 0xda, 0x35, 0x7a, 0xcd, 0x41, 0x85, 0xea, 0x9a, 0x1c, 0x41, 0x23, 0xc5, 0x04, 0x7d, 0x81,
	0x81, 0x55, 0xeb, 0x1a, 0xbd, 0x96, 0xd7, 0x75, 0x56, 0xe5, 0x1d, 0x29, 0xed, 0x50, 0x8d, 0x1b,
	0x54, 0xe8, 0x05, 0xa7, 0x63, 0x43, 0x63, 0xfe, 0x9c, 0xec, 0x82, 0x39, 0xcd, 0x90, 0xdc, 0x32,
	0xba, 0xb5, 0x5e, 0x93, 0xea, 0xaa, 0x7f, 0x0d, 0xea, 0xf2, 0x64, 0xbf, 0x83, 0xda, 0x09, 0xce,
	0x48, 0x1b, 0xea, 0xe7, 0x21, 0x46, 0x81, 0x34, 0xd9, 0xa4, 0xaa, 0x20, 0x7b, 0x1a, 0x25, 0x2d,
	0xb6, 0xbc, 0xbb, 0x6b, 0x6c, 0x50, 0xad, 0xf5, 0xd3, 0x00, 0x93, 0xca, 0xa9, 0x90, 0x57, 0x4b,
	0x7d, 0x5b, 0xde, 0x83, 0x22, 0x55, 0x21, 0x95, 0x02, 0x3f, 0x66, 0x22, 0x9d, 0xcd, 0xdd, 0x75,
	0x28, 0xb4, 0x72, 0x8f, 0xc9, 0x36, 0xd4, 0xbe, 0xe3, 0x4c, 0x5b, 0xcb, 0x8e, 0xff, 0x69, 0xec,
	0x65, 0xf5, 0x85, 0x61, 0x1f, 0xc2, 0xf6, 0x5b, 0x14, 0xaa, 0x29, 0xc5, 0x1f, 0x13, 0xe4, 0x82,
	0x3c, 0x5e, 0x08, 0xb7, 0xbc, 0x3b, 0x45, 0x91, 0x13, 0x9c, 0xc9, 0x7e, 0xf6, 0x31, 0xdc, 0xce,
	0x91, 0x79, 0x12, 0x33, 0x8e, 0xe4, 0x29, 0x98, 0x6a, 0x07, 0xb4, 0x80, 0xb5, 0xee, 0x1b, 0xa9,
	0xc6, 0xd9, 0x47, 0xb0, 0xf3, 0x1a, 0x23, 0x14, 0x78, 0x45, 0x1b, 0xbb, 0xd0, 0x5e, 0xe6, 0x2b,
	0x27, 0xf6, 0xdf, 0x2a, 0xec, 0x9c, 0x4e, 0x84, 0x7f, 0x55, 0x61, 0xf2, 0x19, 0x6e, 0xc9, 0xc4,
	0xbf, 0x8e, 0x33, 0x95, 0x30, 0x66, 0xdc, 0xaa, 0xca, 0xdc, 0xf6, 0x8b, 0xa4, 0x92, 0x46, 0xce,
	0x9b, 0x8c, 0x7c, 0xaa, 0xb9, 0xf4, 0xe6, 0x79, 0xbe, 0xe4, 0x9d, 0x3f, 0x06, 0xdc, 0x58, 0x42,
	0xac, 0x59, 0xb7, 0x27, 0x50, 0xe3, 0x28, 0x36, 0x64, 0x3a, 0xa8, 0xd0, 0x0c, 0x45, 0x0e, 0xa0,
	0x19, 0xb2, 0x61, 0x8a, 0x63, 0x64, 0x42, 0xff, 0x26, 0x97, 0x50, 0x16, 0x58, 0xf2, 0x0c, 0x4c,
	0x3f, 0x49, 0x90, 0x05, 0xd6, 0xd6, 0x26, 0x96, 0x06, 0xf6, 0x01, 0x1a, 0xf3, 0xc1, 0x64, 0x19,
	0x2c, 0x4f, 0x40, 0x65, 0xe0, 0xfd, 0xaa, 0x42, 0xfd, 0xbd, 0x88, 0x53, 0x24, 0x1f, 0xa0, 0x79,
	0xb1, 0x2c, 0xc4, 0x2e, 0xaa, 0xaf, 0xae, 0x61, 0xe7, 0xfe, 0xa5, 0x18, 0xbd, 0x6d, 0x5f, 0xe0,
	0x7a, 0x3e, 0x7b, 0xf2, 0xb0, 0x48, 0x2a, 0xd9, 0xad, 0xce, 0xa3, 0x4d, 0xb0, 0x85, 0x7c, 0xfe,
	0xb3, 0xca, 0xe4, 0x4b, 0x82, 0x2f, 0x93, 0x2f, 0x9b, 0x4e, 0xff, 0xe0, 0xd3, 0xf3, 0x61, 0x1c,
	0xe0, 0x1c, 0x3d, 0x8c, 0xc7, 0xae, 0xbc, 0x38, 0x47, 0xc8, 0xf6, 0x46, 0xb1, 0xbb, 0x7a, 0xb9,
	0x1e, 0xaa, 0x53, 0x72, 0x76, 0x66, 0x4a, 0xd8, 0xfe, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b,
	0xa5, 0x99, 0x5d, 0x80, 0x05, 0x00, 0x00,
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
	// GetRecord retrieves a record.
	GetRecord(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*GetRecordResponse, error)
	// DeleteRecord removes a records.
	DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error)
	// MutateRecord applies mutations to a record.
	MutateRecord(ctx context.Context, in *MutateRecordRequest, opts ...grpc.CallOption) (*MutateRecordResponse, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) GetRecord(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*GetRecordResponse, error) {
	out := new(GetRecordResponse)
	err := c.cc.Invoke(ctx, "/sajari.engine.v2.Store/GetRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error) {
	out := new(DeleteRecordResponse)
	err := c.cc.Invoke(ctx, "/sajari.engine.v2.Store/DeleteRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) MutateRecord(ctx context.Context, in *MutateRecordRequest, opts ...grpc.CallOption) (*MutateRecordResponse, error) {
	out := new(MutateRecordResponse)
	err := c.cc.Invoke(ctx, "/sajari.engine.v2.Store/MutateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	// GetRecord retrieves a record.
	GetRecord(context.Context, *GetRecordRequest) (*GetRecordResponse, error)
	// DeleteRecord removes a records.
	DeleteRecord(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error)
	// MutateRecord applies mutations to a record.
	MutateRecord(context.Context, *MutateRecordRequest) (*MutateRecordResponse, error)
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.engine.v2.Store/GetRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetRecord(ctx, req.(*GetRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.engine.v2.Store/DeleteRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).DeleteRecord(ctx, req.(*DeleteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_MutateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).MutateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sajari.engine.v2.Store/MutateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).MutateRecord(ctx, req.(*MutateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sajari.engine.v2.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecord",
			Handler:    _Store_GetRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _Store_DeleteRecord_Handler,
		},
		{
			MethodName: "MutateRecord",
			Handler:    _Store_MutateRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sajari/engine/v2/record.proto",
}
