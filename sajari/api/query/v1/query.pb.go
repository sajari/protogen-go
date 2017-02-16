// Code generated by protoc-gen-go.
// source: sajari/api/query/v1/query.proto
// DO NOT EDIT!

/*
Package sajari_api_query_v1 is a generated protocol buffer package.

It is generated from these files:
	sajari/api/query/v1/query.proto

It has these top-level messages:
	SearchRequest
	SearchResponse
	Token
	Transform
*/
package sajari_api_query_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sajari_engine_query_v1 "github.com/sajari/go-genproto/sajari/engine/query/v1"

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

type SearchRequest_Tracking_Type int32

const (
	// No tracking will occur.
	SearchRequest_Tracking_NONE SearchRequest_Tracking_Type = 0
	// Click tracking will be used.
	//
	// A click token will be generated for each result.  To register a click see <the other
	// API calls>.  Results which are returned and not clicked on will fall down rankings, and
	// similarly low-ranked documents which receive clicks will be moved up the rankings.
	SearchRequest_Tracking_CLICK SearchRequest_Tracking_Type = 1
	// Pos/neg tokens will be generated for each result.  Each document in the result set can be
	// marked with pos/neg value for the search.  This is then fed back into the ranking algorithm
	// to improve future results.  Unlike CLICK, if no tokens are reported for documents then
	// no action is taken.
	SearchRequest_Tracking_POS_NEG SearchRequest_Tracking_Type = 2
)

var SearchRequest_Tracking_Type_name = map[int32]string{
	0: "NONE",
	1: "CLICK",
	2: "POS_NEG",
}
var SearchRequest_Tracking_Type_value = map[string]int32{
	"NONE":    0,
	"CLICK":   1,
	"POS_NEG": 2,
}

func (x SearchRequest_Tracking_Type) String() string {
	return proto.EnumName(SearchRequest_Tracking_Type_name, int32(x))
}
func (SearchRequest_Tracking_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

type Transform_RunType int32

const (
	// Run the transform before the query.
	Transform_PRE_QUERY Transform_RunType = 0
	// Run the transform on the request.
	Transform_POST_NON_EMPTY Transform_RunType = 1
	// After an empty query this will apply the transform to the query and the re-run the query.
	Transform_POST_EMPTY_PRE_RETRY Transform_RunType = 2
)

var Transform_RunType_name = map[int32]string{
	0: "PRE_QUERY",
	1: "POST_NON_EMPTY",
	2: "POST_EMPTY_PRE_RETRY",
}
var Transform_RunType_value = map[string]int32{
	"PRE_QUERY":            0,
	"POST_NON_EMPTY":       1,
	"POST_EMPTY_PRE_RETRY": 2,
}

func (x Transform_RunType) String() string {
	return proto.EnumName(Transform_RunType_name, int32(x))
}
func (Transform_RunType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// SearchRequest
type SearchRequest struct {
	// Underlying search request to be passed to the engine
	SearchRequest *sajari_engine_query_v1.SearchRequest `protobuf:"bytes,1,opt,name=search_request,json=searchRequest" json:"search_request,omitempty"`
	// Tracking configuration for the query.
	Tracking *SearchRequest_Tracking `protobuf:"bytes,4,opt,name=tracking" json:"tracking,omitempty"`
	// Transforms to be applied to the query (can be before, after or even on failure).
	Transforms []*Transform `protobuf:"bytes,5,rep,name=transforms" json:"transforms,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SearchRequest) GetSearchRequest() *sajari_engine_query_v1.SearchRequest {
	if m != nil {
		return m.SearchRequest
	}
	return nil
}

func (m *SearchRequest) GetTracking() *SearchRequest_Tracking {
	if m != nil {
		return m.Tracking
	}
	return nil
}

func (m *SearchRequest) GetTransforms() []*Transform {
	if m != nil {
		return m.Transforms
	}
	return nil
}

type SearchRequest_Tracking struct {
	// Tracking mode for query.
	//
	// Tracking is done using tokens which are added to result sets and identify individual results.
	// Tokens are used to improve the ranking of documents by reporting clicks (i.e. positive action)
	// or pos/neg (i.e positive or negative reporting) on the position of a document in results.
	Type SearchRequest_Tracking_Type `protobuf:"varint,1,opt,name=type,enum=sajari.api.query.v1.SearchRequest_Tracking_Type" json:"type,omitempty"`
	// Query ID of the query.  If this is empty, then one is generated.
	QueryId string `protobuf:"bytes,2,opt,name=query_id,json=queryId" json:"query_id,omitempty"`
	// Sequence number of query.
	Sequence int32 `protobuf:"varint,3,opt,name=sequence" json:"sequence,omitempty"`
	// Tracking field (must be unique in the collection) used to identify documents in the collection.
	Field string `protobuf:"bytes,4,opt,name=field" json:"field,omitempty"`
	// Custom values to be included in tracking data.
	Data map[string]string `protobuf:"bytes,5,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SearchRequest_Tracking) Reset()                    { *m = SearchRequest_Tracking{} }
func (m *SearchRequest_Tracking) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest_Tracking) ProtoMessage()               {}
func (*SearchRequest_Tracking) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *SearchRequest_Tracking) GetType() SearchRequest_Tracking_Type {
	if m != nil {
		return m.Type
	}
	return SearchRequest_Tracking_NONE
}

func (m *SearchRequest_Tracking) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *SearchRequest_Tracking) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *SearchRequest_Tracking) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *SearchRequest_Tracking) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

// SearchResponse is a wrapped response from the engine with interaction tokens
// and other information which was used to construct the query from the SearchRequest.
type SearchResponse struct {
	// SearchResponse from the engine request.
	SearchResponse *sajari_engine_query_v1.SearchResponse `protobuf:"bytes,1,opt,name=search_response,json=searchResponse" json:"search_response,omitempty"`
	// SearchRequest used in the query (if this has been changed).
	SearchRequest *sajari_engine_query_v1.SearchRequest `protobuf:"bytes,2,opt,name=search_request,json=searchRequest" json:"search_request,omitempty"`
	// Tokens which correspond to the result documents.
	Tokens []*Token `protobuf:"bytes,3,rep,name=tokens" json:"tokens,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchResponse) GetSearchResponse() *sajari_engine_query_v1.SearchResponse {
	if m != nil {
		return m.SearchResponse
	}
	return nil
}

func (m *SearchResponse) GetSearchRequest() *sajari_engine_query_v1.SearchRequest {
	if m != nil {
		return m.SearchRequest
	}
	return nil
}

func (m *SearchResponse) GetTokens() []*Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

// Tokens are used to mark result documents as well/poorly ranked for a query.
// TODO(dhowden): fix this comment!
type Token struct {
	// Types that are valid to be assigned to Token:
	//	*Token_Click_
	//	*Token_PosNeg_
	Token isToken_Token `protobuf_oneof:"token"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isToken_Token interface {
	isToken_Token()
}

type Token_Click_ struct {
	Click *Token_Click `protobuf:"bytes,1,opt,name=click,oneof"`
}
type Token_PosNeg_ struct {
	PosNeg *Token_PosNeg `protobuf:"bytes,2,opt,name=pos_neg,json=posNeg,oneof"`
}

func (*Token_Click_) isToken_Token()  {}
func (*Token_PosNeg_) isToken_Token() {}

func (m *Token) GetToken() isToken_Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Token) GetClick() *Token_Click {
	if x, ok := m.GetToken().(*Token_Click_); ok {
		return x.Click
	}
	return nil
}

func (m *Token) GetPosNeg() *Token_PosNeg {
	if x, ok := m.GetToken().(*Token_PosNeg_); ok {
		return x.PosNeg
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Token) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Token_OneofMarshaler, _Token_OneofUnmarshaler, _Token_OneofSizer, []interface{}{
		(*Token_Click_)(nil),
		(*Token_PosNeg_)(nil),
	}
}

func _Token_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Token)
	// token
	switch x := m.Token.(type) {
	case *Token_Click_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Click); err != nil {
			return err
		}
	case *Token_PosNeg_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PosNeg); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Token.Token has unexpected type %T", x)
	}
	return nil
}

func _Token_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Token)
	switch tag {
	case 1: // token.click
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Token_Click)
		err := b.DecodeMessage(msg)
		m.Token = &Token_Click_{msg}
		return true, err
	case 2: // token.pos_neg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Token_PosNeg)
		err := b.DecodeMessage(msg)
		m.Token = &Token_PosNeg_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Token_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Token)
	// token
	switch x := m.Token.(type) {
	case *Token_Click_:
		s := proto.Size(x.Click)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Token_PosNeg_:
		s := proto.Size(x.PosNeg)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Click is a token that corresponds to selecting a document from a result set.
// It is taken as an indication that this document is a good match for the
// corresponding SearchRequest.
type Token_Click struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *Token_Click) Reset()                    { *m = Token_Click{} }
func (m *Token_Click) String() string            { return proto.CompactTextString(m) }
func (*Token_Click) ProtoMessage()               {}
func (*Token_Click) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *Token_Click) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// PosNeg is a pair of tokens which are used to mark a document as a good/bad
// result for a SearchRequest.
type Token_PosNeg struct {
	Pos string `protobuf:"bytes,1,opt,name=pos" json:"pos,omitempty"`
	Neg string `protobuf:"bytes,2,opt,name=neg" json:"neg,omitempty"`
}

func (m *Token_PosNeg) Reset()                    { *m = Token_PosNeg{} }
func (m *Token_PosNeg) String() string            { return proto.CompactTextString(m) }
func (*Token_PosNeg) ProtoMessage()               {}
func (*Token_PosNeg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

func (m *Token_PosNeg) GetPos() string {
	if m != nil {
		return m.Pos
	}
	return ""
}

func (m *Token_PosNeg) GetNeg() string {
	if m != nil {
		return m.Neg
	}
	return ""
}

type Transform struct {
	// When to run the transform.
	RunType Transform_RunType `protobuf:"varint,1,opt,name=run_type,json=runType,enum=sajari.api.query.v1.Transform_RunType" json:"run_type,omitempty"`
	// Identifier for the transform.
	Identifier string `protobuf:"bytes,2,opt,name=identifier" json:"identifier,omitempty"`
}

func (m *Transform) Reset()                    { *m = Transform{} }
func (m *Transform) String() string            { return proto.CompactTextString(m) }
func (*Transform) ProtoMessage()               {}
func (*Transform) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Transform) GetRunType() Transform_RunType {
	if m != nil {
		return m.RunType
	}
	return Transform_PRE_QUERY
}

func (m *Transform) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "sajari.api.query.v1.SearchRequest")
	proto.RegisterType((*SearchRequest_Tracking)(nil), "sajari.api.query.v1.SearchRequest.Tracking")
	proto.RegisterType((*SearchResponse)(nil), "sajari.api.query.v1.SearchResponse")
	proto.RegisterType((*Token)(nil), "sajari.api.query.v1.Token")
	proto.RegisterType((*Token_Click)(nil), "sajari.api.query.v1.Token.Click")
	proto.RegisterType((*Token_PosNeg)(nil), "sajari.api.query.v1.Token.PosNeg")
	proto.RegisterType((*Transform)(nil), "sajari.api.query.v1.Transform")
	proto.RegisterEnum("sajari.api.query.v1.SearchRequest_Tracking_Type", SearchRequest_Tracking_Type_name, SearchRequest_Tracking_Type_value)
	proto.RegisterEnum("sajari.api.query.v1.Transform_RunType", Transform_RunType_name, Transform_RunType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Query service

type QueryClient interface {
	// Search performs a search.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type queryClient struct {
	cc *grpc.ClientConn
}

func NewQueryClient(cc *grpc.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/sajari.api.query.v1.Query/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Query service

type QueryServer interface {
	// Search performs a search.
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
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
		FullMethod: "/sajari.api.query.v1.Query/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sajari.api.query.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Query_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sajari/api/query/v1/query.proto",
}

func init() { proto.RegisterFile("sajari/api/query/v1/query.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x8d, 0x93, 0x38, 0x89, 0xa7, 0x6a, 0x3e, 0x6b, 0xbf, 0x5e, 0x18, 0x4b, 0x94, 0x60, 0x44,
	0x15, 0x09, 0xe4, 0xd2, 0x20, 0x44, 0x85, 0x10, 0x12, 0xb4, 0xa6, 0xad, 0x28, 0x49, 0xba, 0x0d,
	0x17, 0xe5, 0xc6, 0x5a, 0x92, 0x6d, 0x30, 0x29, 0x6b, 0x77, 0xd7, 0xa9, 0x94, 0xf7, 0x42, 0xbc,
	0x04, 0xd7, 0x3c, 0x01, 0x2f, 0x82, 0xf6, 0x27, 0x56, 0x23, 0x42, 0xaa, 0x8a, 0xbb, 0x9d, 0xa3,
	0x73, 0xce, 0xcc, 0x78, 0x66, 0x0c, 0xf7, 0x04, 0xf9, 0x42, 0x78, 0xb2, 0x4d, 0xb2, 0x64, 0xfb,
	0x72, 0x4a, 0xf9, 0x6c, 0xfb, 0x6a, 0x47, 0x3f, 0xc2, 0x8c, 0xa7, 0x79, 0x8a, 0xfe, 0xd7, 0x84,
	0x90, 0x64, 0x49, 0xa8, 0xf1, 0xab, 0x1d, 0x3f, 0x30, 0x2a, 0xca, 0xc6, 0x09, 0xa3, 0x4b, 0x85,
	0xc1, 0xb7, 0x2a, 0xac, 0x9f, 0x52, 0xc2, 0x87, 0x9f, 0x31, 0xbd, 0x9c, 0x52, 0x91, 0xa3, 0x63,
	0x68, 0x0a, 0x05, 0xc4, 0x5c, 0x23, 0x9e, 0xd5, 0xb2, 0xda, 0x6b, 0x9d, 0x87, 0xa1, 0xc9, 0xa1,
	0xed, 0x8a, 0x34, 0xe1, 0x82, 0x1c, 0xaf, 0x8b, 0x05, 0xb7, 0x03, 0x68, 0xe4, 0x9c, 0x0c, 0x27,
	0x09, 0x1b, 0x7b, 0x55, 0xe5, 0xf3, 0x28, 0x5c, 0x52, 0xeb, 0xa2, 0x49, 0x38, 0x30, 0x12, 0x5c,
	0x88, 0xd1, 0x2b, 0x80, 0x9c, 0x13, 0x26, 0xce, 0x53, 0xfe, 0x55, 0x78, 0x76, 0xab, 0xd2, 0x5e,
	0xeb, 0x6c, 0x2e, 0xb5, 0x1a, 0xcc, 0x69, 0xf8, 0x9a, 0xc2, 0xff, 0x51, 0x86, 0xc6, 0xdc, 0x16,
	0xed, 0x43, 0x35, 0x9f, 0x65, 0x54, 0x75, 0xd6, 0xec, 0x3c, 0xb9, 0x45, 0x45, 0xe1, 0x60, 0x96,
	0x51, 0xac, 0xd4, 0xe8, 0x0e, 0x34, 0x14, 0x3b, 0x4e, 0x46, 0x5e, 0xb9, 0x65, 0xb5, 0x1d, 0x5c,
	0x57, 0xf1, 0xd1, 0x08, 0xf9, 0xd0, 0x10, 0x52, 0xc9, 0x86, 0xd4, 0xab, 0xb4, 0xac, 0xb6, 0x8d,
	0x8b, 0x18, 0x6d, 0x80, 0x7d, 0x9e, 0xd0, 0x8b, 0x91, 0xfa, 0x1e, 0x0e, 0xd6, 0x01, 0x3a, 0x82,
	0xea, 0x88, 0xe4, 0xc4, 0x74, 0xf6, 0xec, 0x36, 0x25, 0xed, 0x93, 0x9c, 0x44, 0x2c, 0xe7, 0x33,
	0xac, 0x2c, 0xfc, 0xe7, 0xe0, 0x14, 0x10, 0x72, 0xa1, 0x32, 0xa1, 0x33, 0xd5, 0xa9, 0x83, 0xe5,
	0x53, 0xe6, 0xbf, 0x22, 0x17, 0x53, 0x6a, 0x6a, 0xd6, 0xc1, 0x8b, 0xf2, 0xae, 0x15, 0xb4, 0xa1,
	0x2a, 0xdb, 0x43, 0x0d, 0xa8, 0x76, 0x7b, 0xdd, 0xc8, 0x2d, 0x21, 0x07, 0xec, 0xbd, 0xe3, 0xa3,
	0xbd, 0x77, 0xae, 0x85, 0xd6, 0xa0, 0xde, 0xef, 0x9d, 0xc6, 0xdd, 0xe8, 0xc0, 0x2d, 0x07, 0xbf,
	0x2c, 0x68, 0xce, 0xab, 0x11, 0x59, 0xca, 0x04, 0x45, 0x3d, 0xf8, 0xaf, 0xd8, 0x1b, 0x0d, 0x99,
	0xc5, 0xd9, 0xba, 0x69, 0x71, 0x34, 0x1b, 0x37, 0xc5, 0xa2, 0xe1, 0x9f, 0x8b, 0x58, 0xfe, 0x87,
	0x45, 0xec, 0x40, 0x2d, 0x4f, 0x27, 0x94, 0x09, 0xaf, 0xa2, 0xbe, 0xb0, 0xbf, 0x7c, 0x77, 0x24,
	0x05, 0x1b, 0x66, 0xf0, 0xd3, 0x02, 0x5b, 0x21, 0x68, 0x17, 0xec, 0xe1, 0x45, 0x32, 0x9c, 0x98,
	0x96, 0x5a, 0x7f, 0x17, 0x87, 0x7b, 0x92, 0x77, 0x58, 0xc2, 0x5a, 0x80, 0x5e, 0x42, 0x3d, 0x4b,
	0x45, 0xcc, 0xe8, 0xd8, 0x94, 0x7f, 0x7f, 0x85, 0xb6, 0x9f, 0x8a, 0x2e, 0x1d, 0x1f, 0x96, 0x70,
	0x2d, 0x53, 0x2f, 0xff, 0x2e, 0xd8, 0xca, 0x4f, 0x0e, 0x4d, 0x15, 0x65, 0x06, 0xa9, 0x03, 0xff,
	0x31, 0xd4, 0xb4, 0x44, 0x8e, 0x39, 0x4b, 0xc5, 0x7c, 0xcc, 0x59, 0x2a, 0x24, 0x32, 0x4f, 0xea,
	0x60, 0xf9, 0x7c, 0x53, 0x37, 0x1e, 0xc1, 0x77, 0x0b, 0x9c, 0xe2, 0x4a, 0xd0, 0x6b, 0x68, 0xf0,
	0x29, 0x8b, 0xaf, 0x1d, 0xc4, 0xd6, 0xea, 0xbb, 0x0a, 0xf1, 0x94, 0xa9, 0x33, 0xa8, 0x73, 0xfd,
	0x40, 0x9b, 0x00, 0xc9, 0x88, 0xb2, 0x3c, 0x39, 0x4f, 0x28, 0x37, 0x29, 0xaf, 0x21, 0xc1, 0x5b,
	0xa8, 0x1b, 0x0d, 0x5a, 0x07, 0xa7, 0x8f, 0xa3, 0xf8, 0xe4, 0x43, 0x84, 0xcf, 0xdc, 0x12, 0x42,
	0xd0, 0xec, 0xf7, 0x4e, 0x07, 0x71, 0xb7, 0xd7, 0x8d, 0xa3, 0xf7, 0xfd, 0xc1, 0x99, 0x6b, 0x21,
	0x0f, 0x36, 0x14, 0xa6, 0xe2, 0x58, 0xb2, 0x71, 0x34, 0xc0, 0x67, 0x6e, 0xb9, 0xf3, 0x11, 0xec,
	0x13, 0x59, 0x0e, 0x3a, 0x81, 0x9a, 0x9e, 0x36, 0x0a, 0x6e, 0xbe, 0x14, 0xff, 0xc1, 0x4a, 0x8e,
	0x5e, 0xb7, 0x4f, 0x35, 0xf5, 0x43, 0x7c, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x87, 0x12, 0x42,
	0x7e, 0x6c, 0x05, 0x00, 0x00,
}
