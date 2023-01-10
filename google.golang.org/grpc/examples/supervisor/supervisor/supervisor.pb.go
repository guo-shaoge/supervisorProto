// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supervisor.proto

package supervisor

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Result struct {
	HasErr               bool     `protobuf:"varint,1,opt,name=hasErr,proto3" json:"hasErr,omitempty"`
	ErrInfo              string   `protobuf:"bytes,2,opt,name=errInfo,proto3" json:"errInfo,omitempty"`
	TenantID             string   `protobuf:"bytes,3,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	StartTime            int64    `protobuf:"varint,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetHasErr() bool {
	if m != nil {
		return m.HasErr
	}
	return false
}

func (m *Result) GetErrInfo() string {
	if m != nil {
		return m.ErrInfo
	}
	return ""
}

func (m *Result) GetTenantID() string {
	if m != nil {
		return m.TenantID
	}
	return ""
}

func (m *Result) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

type AssignRequest struct {
	TenantID             string   `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	TidbStatusAddr       string   `protobuf:"bytes,2,opt,name=tidbStatusAddr,proto3" json:"tidbStatusAddr,omitempty"`
	PdAddr               string   `protobuf:"bytes,3,opt,name=pdAddr,proto3" json:"pdAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignRequest) Reset()         { *m = AssignRequest{} }
func (m *AssignRequest) String() string { return proto.CompactTextString(m) }
func (*AssignRequest) ProtoMessage()    {}
func (*AssignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{1}
}

func (m *AssignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignRequest.Unmarshal(m, b)
}
func (m *AssignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignRequest.Marshal(b, m, deterministic)
}
func (m *AssignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignRequest.Merge(m, src)
}
func (m *AssignRequest) XXX_Size() int {
	return xxx_messageInfo_AssignRequest.Size(m)
}
func (m *AssignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssignRequest proto.InternalMessageInfo

func (m *AssignRequest) GetTenantID() string {
	if m != nil {
		return m.TenantID
	}
	return ""
}

func (m *AssignRequest) GetTidbStatusAddr() string {
	if m != nil {
		return m.TidbStatusAddr
	}
	return ""
}

func (m *AssignRequest) GetPdAddr() string {
	if m != nil {
		return m.PdAddr
	}
	return ""
}

type UnassignRequest struct {
	AssertTenantID       string   `protobuf:"bytes,1,opt,name=assertTenantID,proto3" json:"assertTenantID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnassignRequest) Reset()         { *m = UnassignRequest{} }
func (m *UnassignRequest) String() string { return proto.CompactTextString(m) }
func (*UnassignRequest) ProtoMessage()    {}
func (*UnassignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{2}
}

func (m *UnassignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnassignRequest.Unmarshal(m, b)
}
func (m *UnassignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnassignRequest.Marshal(b, m, deterministic)
}
func (m *UnassignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnassignRequest.Merge(m, src)
}
func (m *UnassignRequest) XXX_Size() int {
	return xxx_messageInfo_UnassignRequest.Size(m)
}
func (m *UnassignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnassignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnassignRequest proto.InternalMessageInfo

func (m *UnassignRequest) GetAssertTenantID() string {
	if m != nil {
		return m.AssertTenantID
	}
	return ""
}

type GetTenantResponse struct {
	TenantID             string   `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	StartTime            int64    `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTenantResponse) Reset()         { *m = GetTenantResponse{} }
func (m *GetTenantResponse) String() string { return proto.CompactTextString(m) }
func (*GetTenantResponse) ProtoMessage()    {}
func (*GetTenantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{3}
}

func (m *GetTenantResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTenantResponse.Unmarshal(m, b)
}
func (m *GetTenantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTenantResponse.Marshal(b, m, deterministic)
}
func (m *GetTenantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTenantResponse.Merge(m, src)
}
func (m *GetTenantResponse) XXX_Size() int {
	return xxx_messageInfo_GetTenantResponse.Size(m)
}
func (m *GetTenantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTenantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTenantResponse proto.InternalMessageInfo

func (m *GetTenantResponse) GetTenantID() string {
	if m != nil {
		return m.TenantID
	}
	return ""
}

func (m *GetTenantResponse) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Result)(nil), "supervisor.Result")
	proto.RegisterType((*AssignRequest)(nil), "supervisor.AssignRequest")
	proto.RegisterType((*UnassignRequest)(nil), "supervisor.UnassignRequest")
	proto.RegisterType((*GetTenantResponse)(nil), "supervisor.GetTenantResponse")
}

func init() { proto.RegisterFile("supervisor.proto", fileDescriptor_b8b9452d77b1c7d2) }

var fileDescriptor_b8b9452d77b1c7d2 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcf, 0xaf, 0xd2, 0x40,
	0x10, 0xa6, 0x60, 0x2a, 0x4c, 0x14, 0x70, 0x0f, 0xa4, 0x16, 0x4d, 0x48, 0x0f, 0x84, 0xd3, 0x36,
	0xd1, 0x78, 0xf0, 0x64, 0x00, 0x09, 0x21, 0xc6, 0x84, 0x14, 0xbc, 0x78, 0x5b, 0xec, 0xb0, 0x36,
	0x96, 0xdd, 0xba, 0xbb, 0x35, 0xbe, 0xbf, 0xf5, 0xfd, 0x33, 0x2f, 0xfd, 0x05, 0x6d, 0x5f, 0xde,
	0xbb, 0xcd, 0x37, 0xb3, 0xf3, 0xcd, 0xf7, 0xed, 0x0c, 0x8c, 0x75, 0x9a, 0xa0, 0xfa, 0x17, 0x69,
	0xa9, 0x68, 0xa2, 0xa4, 0x91, 0x04, 0x6e, 0x19, 0x77, 0xca, 0xa5, 0xe4, 0x31, 0xfa, 0x79, 0xe5,
	0x94, 0x9e, 0x7d, 0xbc, 0x24, 0xe6, 0xae, 0x78, 0xe8, 0x19, 0xb0, 0x03, 0xd4, 0x69, 0x6c, 0xc8,
	0x04, 0xec, 0xdf, 0x4c, 0x6f, 0x94, 0x72, 0xac, 0x99, 0xb5, 0xe8, 0x07, 0x25, 0x22, 0x0e, 0xbc,
	0x44, 0xa5, 0x76, 0xe2, 0x2c, 0x9d, 0xee, 0xcc, 0x5a, 0x0c, 0x82, 0x0a, 0x12, 0x17, 0xfa, 0x06,
	0x05, 0x13, 0x66, 0xf7, 0xd5, 0xe9, 0xe5, 0xa5, 0x2b, 0x26, 0xef, 0x60, 0xa0, 0x0d, 0x53, 0xe6,
	0x18, 0x5d, 0xd0, 0x79, 0x31, 0xb3, 0x16, 0xbd, 0xe0, 0x96, 0xf0, 0xfe, 0xc0, 0xeb, 0xa5, 0xd6,
	0x11, 0x17, 0x01, 0xfe, 0x4d, 0x51, 0x9b, 0x06, 0x95, 0xd5, 0xa2, 0x9a, 0xc3, 0xd0, 0x44, 0xe1,
	0xe9, 0x60, 0x98, 0x49, 0xf5, 0x32, 0x0c, 0x55, 0xa9, 0xa3, 0x95, 0xcd, 0x0c, 0x24, 0x61, 0x5e,
	0x2f, 0xc4, 0x94, 0xc8, 0xfb, 0x0c, 0xa3, 0x1f, 0x82, 0x35, 0xc6, 0xcd, 0x61, 0xc8, 0xb4, 0x46,
	0x65, 0x8e, 0xcd, 0xa1, 0xad, 0xac, 0xf7, 0x1d, 0xde, 0x6c, 0xb1, 0x84, 0x01, 0xea, 0x44, 0x0a,
	0x8d, 0xcf, 0x6a, 0x6d, 0xd8, 0xee, 0xb6, 0x6c, 0x7f, 0xb8, 0xb7, 0xc0, 0x2e, 0x7c, 0x93, 0x2f,
	0xf0, 0xaa, 0x88, 0x0a, 0x72, 0xf2, 0x96, 0xd6, 0x76, 0xd8, 0xf8, 0x1b, 0x97, 0xd4, 0x4b, 0xc5,
	0xb2, 0xbc, 0x0e, 0x59, 0xc3, 0xb0, 0x72, 0x55, 0x52, 0x4c, 0xeb, 0xef, 0x5a, 0x8e, 0x9f, 0x20,
	0xf9, 0x06, 0xe3, 0x2d, 0x9a, 0x75, 0xaa, 0x14, 0x8a, 0xd2, 0x26, 0x99, 0xd0, 0xe2, 0x5e, 0x68,
	0x75, 0x2f, 0x74, 0x93, 0xdd, 0x8b, 0xfb, 0xbe, 0xce, 0xf0, 0xe8, 0x57, 0xbc, 0xce, 0x8a, 0xc3,
	0x34, 0x92, 0x94, 0xab, 0xe4, 0x17, 0xc5, 0xff, 0xec, 0x92, 0xc4, 0xa8, 0x6b, 0x2d, 0xab, 0xd1,
	0xe1, 0x1a, 0xef, 0x33, 0xe6, 0xbd, 0xf5, 0xf3, 0x53, 0x39, 0x89, 0xcb, 0x98, 0x09, 0x4e, 0xa5,
	0xe2, 0x7e, 0xd6, 0xee, 0x57, 0xed, 0xfe, 0xad, 0xbd, 0x16, 0x9e, 0xec, 0x5c, 0xd9, 0xc7, 0x87,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0xe2, 0xae, 0x55, 0xf7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AssignClient is the client API for Assign service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssignClient interface {
	AssignTenant(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*Result, error)
	UnassignTenant(ctx context.Context, in *UnassignRequest, opts ...grpc.CallOption) (*Result, error)
	GetCurrentTenant(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTenantResponse, error)
}

type assignClient struct {
	cc *grpc.ClientConn
}

func NewAssignClient(cc *grpc.ClientConn) AssignClient {
	return &assignClient{cc}
}

func (c *assignClient) AssignTenant(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/supervisor.Assign/AssignTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignClient) UnassignTenant(ctx context.Context, in *UnassignRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/supervisor.Assign/UnassignTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignClient) GetCurrentTenant(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTenantResponse, error) {
	out := new(GetTenantResponse)
	err := c.cc.Invoke(ctx, "/supervisor.Assign/GetCurrentTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssignServer is the server API for Assign service.
type AssignServer interface {
	AssignTenant(context.Context, *AssignRequest) (*Result, error)
	UnassignTenant(context.Context, *UnassignRequest) (*Result, error)
	GetCurrentTenant(context.Context, *emptypb.Empty) (*GetTenantResponse, error)
}

func RegisterAssignServer(s *grpc.Server, srv AssignServer) {
	s.RegisterService(&_Assign_serviceDesc, srv)
}

func _Assign_AssignTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignServer).AssignTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Assign/AssignTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignServer).AssignTenant(ctx, req.(*AssignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assign_UnassignTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnassignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignServer).UnassignTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Assign/UnassignTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignServer).UnassignTenant(ctx, req.(*UnassignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assign_GetCurrentTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignServer).GetCurrentTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Assign/GetCurrentTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignServer).GetCurrentTenant(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Assign_serviceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.Assign",
	HandlerType: (*AssignServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignTenant",
			Handler:    _Assign_AssignTenant_Handler,
		},
		{
			MethodName: "UnassignTenant",
			Handler:    _Assign_UnassignTenant_Handler,
		},
		{
			MethodName: "GetCurrentTenant",
			Handler:    _Assign_GetCurrentTenant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supervisor.proto",
}
