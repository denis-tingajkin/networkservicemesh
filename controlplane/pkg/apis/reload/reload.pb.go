// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reload.proto

package reload

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ReloadMessage struct {
	Context              string   `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReloadMessage) Reset()         { *m = ReloadMessage{} }
func (m *ReloadMessage) String() string { return proto.CompactTextString(m) }
func (*ReloadMessage) ProtoMessage()    {}
func (*ReloadMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8276ccc67e782642, []int{0}
}

func (m *ReloadMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadMessage.Unmarshal(m, b)
}
func (m *ReloadMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadMessage.Marshal(b, m, deterministic)
}
func (m *ReloadMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadMessage.Merge(m, src)
}
func (m *ReloadMessage) XXX_Size() int {
	return xxx_messageInfo_ReloadMessage.Size(m)
}
func (m *ReloadMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadMessage proto.InternalMessageInfo

func (m *ReloadMessage) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func init() {
	proto.RegisterType((*ReloadMessage)(nil), "reload.ReloadMessage")
}

func init() { proto.RegisterFile("reload.proto", fileDescriptor_8276ccc67e782642) }

var fileDescriptor_8276ccc67e782642 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0xcd, 0xc9,
	0x4f, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x24, 0x0a, 0x4a,
	0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0x21, 0x24, 0x44, 0x85, 0x92, 0x26,
	0x17, 0x6f, 0x10, 0x58, 0x8d, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x90, 0x04, 0x17, 0x7b,
	0x72, 0x7e, 0x5e, 0x49, 0x6a, 0x45, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x6b,
	0xe4, 0x01, 0x53, 0x1a, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xce, 0xc5, 0x06, 0x11,
	0x10, 0x12, 0xd5, 0x83, 0x5a, 0x8b, 0x62, 0x96, 0x94, 0x98, 0x5e, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e,
	0x2a, 0xc4, 0xae, 0xa4, 0xd2, 0x34, 0x3d, 0x57, 0x90, 0xd5, 0x49, 0x6c, 0x60, 0xbe, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x51, 0x62, 0x14, 0x81, 0xad, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReloadServiceClient is the client API for ReloadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReloadServiceClient interface {
	Reload(ctx context.Context, in *ReloadMessage, opts ...grpc.CallOption) (*empty.Empty, error)
}

type reloadServiceClient struct {
	cc *grpc.ClientConn
}

func NewReloadServiceClient(cc *grpc.ClientConn) ReloadServiceClient {
	return &reloadServiceClient{cc}
}

func (c *reloadServiceClient) Reload(ctx context.Context, in *ReloadMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/reload.ReloadService/Reload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReloadServiceServer is the server API for ReloadService service.
type ReloadServiceServer interface {
	Reload(context.Context, *ReloadMessage) (*empty.Empty, error)
}

func RegisterReloadServiceServer(s *grpc.Server, srv ReloadServiceServer) {
	s.RegisterService(&_ReloadService_serviceDesc, srv)
}

func _ReloadService_Reload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReloadServiceServer).Reload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reload.ReloadService/Reload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReloadServiceServer).Reload(ctx, req.(*ReloadMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReloadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reload.ReloadService",
	HandlerType: (*ReloadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reload",
			Handler:    _ReloadService_Reload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reload.proto",
}