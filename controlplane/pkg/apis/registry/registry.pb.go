// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registry.proto

package registry

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type NetworkServiceEndpoint struct {
	NetworkServiceName        string            `protobuf:"bytes,1,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	Payload                   string            `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	NetworkServiceManagerName string            `protobuf:"bytes,3,opt,name=network_service_manager_name,json=networkServiceManagerName,proto3" json:"network_service_manager_name,omitempty"`
	EndpointName              string            `protobuf:"bytes,4,opt,name=endpoint_name,json=endpointName,proto3" json:"endpoint_name,omitempty"`
	Labels                    map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	State                     string            `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}          `json:"-"`
	XXX_unrecognized          []byte            `json:"-"`
	XXX_sizecache             int32             `json:"-"`
}

func (m *NetworkServiceEndpoint) Reset()         { *m = NetworkServiceEndpoint{} }
func (m *NetworkServiceEndpoint) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceEndpoint) ProtoMessage()    {}
func (*NetworkServiceEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_418477a5ea7506d5, []int{0}
}
func (m *NetworkServiceEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceEndpoint.Unmarshal(m, b)
}
func (m *NetworkServiceEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceEndpoint.Marshal(b, m, deterministic)
}
func (dst *NetworkServiceEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceEndpoint.Merge(dst, src)
}
func (m *NetworkServiceEndpoint) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceEndpoint.Size(m)
}
func (m *NetworkServiceEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceEndpoint proto.InternalMessageInfo

func (m *NetworkServiceEndpoint) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetNetworkServiceManagerName() string {
	if m != nil {
		return m.NetworkServiceManagerName
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetEndpointName() string {
	if m != nil {
		return m.EndpointName
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *NetworkServiceEndpoint) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type NetworkService struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Payload              string   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Matches              []*Match `protobuf:"bytes,3,rep,name=matches,proto3" json:"matches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkService) Reset()         { *m = NetworkService{} }
func (m *NetworkService) String() string { return proto.CompactTextString(m) }
func (*NetworkService) ProtoMessage()    {}
func (*NetworkService) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_418477a5ea7506d5, []int{1}
}
func (m *NetworkService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkService.Unmarshal(m, b)
}
func (m *NetworkService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkService.Marshal(b, m, deterministic)
}
func (dst *NetworkService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkService.Merge(dst, src)
}
func (m *NetworkService) XXX_Size() int {
	return xxx_messageInfo_NetworkService.Size(m)
}
func (m *NetworkService) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkService.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkService proto.InternalMessageInfo

func (m *NetworkService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkService) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *NetworkService) GetMatches() []*Match {
	if m != nil {
		return m.Matches
	}
	return nil
}

type Match struct {
	SourceSelector       map[string]string `protobuf:"bytes,1,rep,name=source_selector,json=sourceSelector,proto3" json:"source_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Routes               []*Destination    `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_418477a5ea7506d5, []int{2}
}
func (m *Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Match.Unmarshal(m, b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Match.Marshal(b, m, deterministic)
}
func (dst *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(dst, src)
}
func (m *Match) XXX_Size() int {
	return xxx_messageInfo_Match.Size(m)
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetSourceSelector() map[string]string {
	if m != nil {
		return m.SourceSelector
	}
	return nil
}

func (m *Match) GetRoutes() []*Destination {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Destination struct {
	DestinationSelector  map[string]string `protobuf:"bytes,1,rep,name=destination_selector,json=destinationSelector,proto3" json:"destination_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Weight               uint32            `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Destination) Reset()         { *m = Destination{} }
func (m *Destination) String() string { return proto.CompactTextString(m) }
func (*Destination) ProtoMessage()    {}
func (*Destination) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_418477a5ea7506d5, []int{3}
}
func (m *Destination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Destination.Unmarshal(m, b)
}
func (m *Destination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Destination.Marshal(b, m, deterministic)
}
func (dst *Destination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Destination.Merge(dst, src)
}
func (m *Destination) XXX_Size() int {
	return xxx_messageInfo_Destination.Size(m)
}
func (m *Destination) XXX_DiscardUnknown() {
	xxx_messageInfo_Destination.DiscardUnknown(m)
}

var xxx_messageInfo_Destination proto.InternalMessageInfo

func (m *Destination) GetDestinationSelector() map[string]string {
	if m != nil {
		return m.DestinationSelector
	}
	return nil
}

func (m *Destination) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type NetworkServiceManager struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	LastSeen             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_seen,json=lastSeen,proto3" json:"last_seen,omitempty"`
	State                string               `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NetworkServiceManager) Reset()         { *m = NetworkServiceManager{} }
func (m *NetworkServiceManager) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceManager) ProtoMessage()    {}
func (*NetworkServiceManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_418477a5ea7506d5, []int{4}
}
func (m *NetworkServiceManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceManager.Unmarshal(m, b)
}
func (m *NetworkServiceManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceManager.Marshal(b, m, deterministic)
}
func (dst *NetworkServiceManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceManager.Merge(dst, src)
}
func (m *NetworkServiceManager) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceManager.Size(m)
}
func (m *NetworkServiceManager) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceManager.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceManager proto.InternalMessageInfo

func (m *NetworkServiceManager) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkServiceManager) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *NetworkServiceManager) GetLastSeen() *timestamp.Timestamp {
	if m != nil {
		return m.LastSeen
	}
	return nil
}

func (m *NetworkServiceManager) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type RemoveNSERequest struct {
	EndpointName         string   `protobuf:"bytes,1,opt,name=endpoint_name,json=endpointName,proto3" json:"endpoint_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveNSERequest) Reset()         { *m = RemoveNSERequest{} }
func (m *RemoveNSERequest) String() string { return proto.CompactTextString(m) }
func (*RemoveNSERequest) ProtoMessage()    {}
func (*RemoveNSERequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_418477a5ea7506d5, []int{5}
}
func (m *RemoveNSERequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveNSERequest.Unmarshal(m, b)
}
func (m *RemoveNSERequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveNSERequest.Marshal(b, m, deterministic)
}
func (dst *RemoveNSERequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveNSERequest.Merge(dst, src)
}
func (m *RemoveNSERequest) XXX_Size() int {
	return xxx_messageInfo_RemoveNSERequest.Size(m)
}
func (m *RemoveNSERequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveNSERequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveNSERequest proto.InternalMessageInfo

func (m *RemoveNSERequest) GetEndpointName() string {
	if m != nil {
		return m.EndpointName
	}
	return ""
}

type FindNetworkServiceRequest struct {
	NetworkServiceName   string   `protobuf:"bytes,1,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindNetworkServiceRequest) Reset()         { *m = FindNetworkServiceRequest{} }
func (m *FindNetworkServiceRequest) String() string { return proto.CompactTextString(m) }
func (*FindNetworkServiceRequest) ProtoMessage()    {}
func (*FindNetworkServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_418477a5ea7506d5, []int{6}
}
func (m *FindNetworkServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNetworkServiceRequest.Unmarshal(m, b)
}
func (m *FindNetworkServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNetworkServiceRequest.Marshal(b, m, deterministic)
}
func (dst *FindNetworkServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNetworkServiceRequest.Merge(dst, src)
}
func (m *FindNetworkServiceRequest) XXX_Size() int {
	return xxx_messageInfo_FindNetworkServiceRequest.Size(m)
}
func (m *FindNetworkServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNetworkServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindNetworkServiceRequest proto.InternalMessageInfo

func (m *FindNetworkServiceRequest) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

type FindNetworkServiceResponse struct {
	Payload                 string                            `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	NetworkService          *NetworkService                   `protobuf:"bytes,2,opt,name=network_service,json=networkService,proto3" json:"network_service,omitempty"`
	NetworkServiceManagers  map[string]*NetworkServiceManager `protobuf:"bytes,3,rep,name=network_service_managers,json=networkServiceManagers,proto3" json:"network_service_managers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NetworkServiceEndpoints []*NetworkServiceEndpoint         `protobuf:"bytes,4,rep,name=network_service_endpoints,json=networkServiceEndpoints,proto3" json:"network_service_endpoints,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                          `json:"-"`
	XXX_unrecognized        []byte                            `json:"-"`
	XXX_sizecache           int32                             `json:"-"`
}

func (m *FindNetworkServiceResponse) Reset()         { *m = FindNetworkServiceResponse{} }
func (m *FindNetworkServiceResponse) String() string { return proto.CompactTextString(m) }
func (*FindNetworkServiceResponse) ProtoMessage()    {}
func (*FindNetworkServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_418477a5ea7506d5, []int{7}
}
func (m *FindNetworkServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNetworkServiceResponse.Unmarshal(m, b)
}
func (m *FindNetworkServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNetworkServiceResponse.Marshal(b, m, deterministic)
}
func (dst *FindNetworkServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNetworkServiceResponse.Merge(dst, src)
}
func (m *FindNetworkServiceResponse) XXX_Size() int {
	return xxx_messageInfo_FindNetworkServiceResponse.Size(m)
}
func (m *FindNetworkServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNetworkServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindNetworkServiceResponse proto.InternalMessageInfo

func (m *FindNetworkServiceResponse) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *FindNetworkServiceResponse) GetNetworkService() *NetworkService {
	if m != nil {
		return m.NetworkService
	}
	return nil
}

func (m *FindNetworkServiceResponse) GetNetworkServiceManagers() map[string]*NetworkServiceManager {
	if m != nil {
		return m.NetworkServiceManagers
	}
	return nil
}

func (m *FindNetworkServiceResponse) GetNetworkServiceEndpoints() []*NetworkServiceEndpoint {
	if m != nil {
		return m.NetworkServiceEndpoints
	}
	return nil
}

type NSERegistration struct {
	NetworkService         *NetworkService         `protobuf:"bytes,1,opt,name=network_service,json=networkService,proto3" json:"network_service,omitempty"`
	NetworkServiceManager  *NetworkServiceManager  `protobuf:"bytes,2,opt,name=network_service_manager,json=networkServiceManager,proto3" json:"network_service_manager,omitempty"`
	NetworkserviceEndpoint *NetworkServiceEndpoint `protobuf:"bytes,3,opt,name=networkservice_endpoint,json=networkserviceEndpoint,proto3" json:"networkservice_endpoint,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *NSERegistration) Reset()         { *m = NSERegistration{} }
func (m *NSERegistration) String() string { return proto.CompactTextString(m) }
func (*NSERegistration) ProtoMessage()    {}
func (*NSERegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_418477a5ea7506d5, []int{8}
}
func (m *NSERegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NSERegistration.Unmarshal(m, b)
}
func (m *NSERegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NSERegistration.Marshal(b, m, deterministic)
}
func (dst *NSERegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NSERegistration.Merge(dst, src)
}
func (m *NSERegistration) XXX_Size() int {
	return xxx_messageInfo_NSERegistration.Size(m)
}
func (m *NSERegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_NSERegistration.DiscardUnknown(m)
}

var xxx_messageInfo_NSERegistration proto.InternalMessageInfo

func (m *NSERegistration) GetNetworkService() *NetworkService {
	if m != nil {
		return m.NetworkService
	}
	return nil
}

func (m *NSERegistration) GetNetworkServiceManager() *NetworkServiceManager {
	if m != nil {
		return m.NetworkServiceManager
	}
	return nil
}

func (m *NSERegistration) GetNetworkserviceEndpoint() *NetworkServiceEndpoint {
	if m != nil {
		return m.NetworkserviceEndpoint
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkServiceEndpoint)(nil), "registry.NetworkServiceEndpoint")
	proto.RegisterMapType((map[string]string)(nil), "registry.NetworkServiceEndpoint.LabelsEntry")
	proto.RegisterType((*NetworkService)(nil), "registry.NetworkService")
	proto.RegisterType((*Match)(nil), "registry.Match")
	proto.RegisterMapType((map[string]string)(nil), "registry.Match.SourceSelectorEntry")
	proto.RegisterType((*Destination)(nil), "registry.Destination")
	proto.RegisterMapType((map[string]string)(nil), "registry.Destination.DestinationSelectorEntry")
	proto.RegisterType((*NetworkServiceManager)(nil), "registry.NetworkServiceManager")
	proto.RegisterType((*RemoveNSERequest)(nil), "registry.RemoveNSERequest")
	proto.RegisterType((*FindNetworkServiceRequest)(nil), "registry.FindNetworkServiceRequest")
	proto.RegisterType((*FindNetworkServiceResponse)(nil), "registry.FindNetworkServiceResponse")
	proto.RegisterMapType((map[string]*NetworkServiceManager)(nil), "registry.FindNetworkServiceResponse.NetworkServiceManagersEntry")
	proto.RegisterType((*NSERegistration)(nil), "registry.NSERegistration")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceRegistryClient is the client API for NetworkServiceRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceRegistryClient interface {
	RegisterNSE(ctx context.Context, in *NSERegistration, opts ...grpc.CallOption) (*NSERegistration, error)
	RemoveNSE(ctx context.Context, in *RemoveNSERequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServiceRegistryClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceRegistryClient(cc *grpc.ClientConn) NetworkServiceRegistryClient {
	return &networkServiceRegistryClient{cc}
}

func (c *networkServiceRegistryClient) RegisterNSE(ctx context.Context, in *NSERegistration, opts ...grpc.CallOption) (*NSERegistration, error) {
	out := new(NSERegistration)
	err := c.cc.Invoke(ctx, "/registry.NetworkServiceRegistry/RegisterNSE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceRegistryClient) RemoveNSE(ctx context.Context, in *RemoveNSERequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/registry.NetworkServiceRegistry/RemoveNSE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceRegistryServer is the server API for NetworkServiceRegistry service.
type NetworkServiceRegistryServer interface {
	RegisterNSE(context.Context, *NSERegistration) (*NSERegistration, error)
	RemoveNSE(context.Context, *RemoveNSERequest) (*empty.Empty, error)
}

func RegisterNetworkServiceRegistryServer(s *grpc.Server, srv NetworkServiceRegistryServer) {
	s.RegisterService(&_NetworkServiceRegistry_serviceDesc, srv)
}

func _NetworkServiceRegistry_RegisterNSE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NSERegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceRegistryServer).RegisterNSE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.NetworkServiceRegistry/RegisterNSE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceRegistryServer).RegisterNSE(ctx, req.(*NSERegistration))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServiceRegistry_RemoveNSE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNSERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceRegistryServer).RemoveNSE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.NetworkServiceRegistry/RemoveNSE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceRegistryServer).RemoveNSE(ctx, req.(*RemoveNSERequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServiceRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registry.NetworkServiceRegistry",
	HandlerType: (*NetworkServiceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNSE",
			Handler:    _NetworkServiceRegistry_RegisterNSE_Handler,
		},
		{
			MethodName: "RemoveNSE",
			Handler:    _NetworkServiceRegistry_RemoveNSE_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.proto",
}

// NetworkServiceDiscoveryClient is the client API for NetworkServiceDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceDiscoveryClient interface {
	FindNetworkService(ctx context.Context, in *FindNetworkServiceRequest, opts ...grpc.CallOption) (*FindNetworkServiceResponse, error)
}

type networkServiceDiscoveryClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceDiscoveryClient(cc *grpc.ClientConn) NetworkServiceDiscoveryClient {
	return &networkServiceDiscoveryClient{cc}
}

func (c *networkServiceDiscoveryClient) FindNetworkService(ctx context.Context, in *FindNetworkServiceRequest, opts ...grpc.CallOption) (*FindNetworkServiceResponse, error) {
	out := new(FindNetworkServiceResponse)
	err := c.cc.Invoke(ctx, "/registry.NetworkServiceDiscovery/FindNetworkService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceDiscoveryServer is the server API for NetworkServiceDiscovery service.
type NetworkServiceDiscoveryServer interface {
	FindNetworkService(context.Context, *FindNetworkServiceRequest) (*FindNetworkServiceResponse, error)
}

func RegisterNetworkServiceDiscoveryServer(s *grpc.Server, srv NetworkServiceDiscoveryServer) {
	s.RegisterService(&_NetworkServiceDiscovery_serviceDesc, srv)
}

func _NetworkServiceDiscovery_FindNetworkService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNetworkServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceDiscoveryServer).FindNetworkService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.NetworkServiceDiscovery/FindNetworkService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceDiscoveryServer).FindNetworkService(ctx, req.(*FindNetworkServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServiceDiscovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registry.NetworkServiceDiscovery",
	HandlerType: (*NetworkServiceDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindNetworkService",
			Handler:    _NetworkServiceDiscovery_FindNetworkService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.proto",
}

func init() { proto.RegisterFile("registry.proto", fileDescriptor_registry_418477a5ea7506d5) }

var fileDescriptor_registry_418477a5ea7506d5 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xcf, 0x4e, 0xdb, 0x4e,
	0x10, 0x96, 0x93, 0x10, 0x60, 0xfc, 0x23, 0x41, 0x0b, 0x04, 0xc7, 0xfc, 0x24, 0x50, 0xe8, 0x81,
	0x4a, 0xad, 0xa9, 0x52, 0x55, 0xb4, 0xbd, 0x50, 0x54, 0xc2, 0x09, 0x72, 0x70, 0x2a, 0x55, 0x95,
	0x2a, 0x45, 0x4b, 0x32, 0x0d, 0x2e, 0xf6, 0x6e, 0xea, 0xdd, 0x04, 0x85, 0x27, 0xe8, 0xa1, 0xcf,
	0xd0, 0x57, 0xe9, 0xb1, 0xd7, 0xbe, 0x42, 0xdf, 0xa4, 0xca, 0xda, 0x26, 0xb6, 0x63, 0x03, 0xb9,
	0xed, 0x9f, 0x6f, 0xbe, 0x99, 0xf9, 0x3c, 0xdf, 0x26, 0x50, 0xf1, 0x71, 0xe0, 0x08, 0xe9, 0x4f,
	0xac, 0xa1, 0xcf, 0x25, 0x27, 0x2b, 0xd1, 0xde, 0xdc, 0x19, 0x70, 0x3e, 0x70, 0xf1, 0x50, 0x9d,
	0x5f, 0x8e, 0xbe, 0x1c, 0xa2, 0x37, 0x94, 0x21, 0xcc, 0xdc, 0x4d, 0x5f, 0x4a, 0xc7, 0x43, 0x21,
	0xa9, 0x37, 0x0c, 0x00, 0x8d, 0xbf, 0x05, 0xa8, 0xb5, 0x51, 0xde, 0x70, 0xff, 0xba, 0x83, 0xfe,
	0xd8, 0xe9, 0x61, 0x8b, 0xf5, 0x87, 0xdc, 0x61, 0x92, 0xbc, 0x80, 0x4d, 0x16, 0xdc, 0x74, 0x45,
	0x70, 0xd5, 0x65, 0xd4, 0x43, 0x43, 0xdb, 0xd3, 0x0e, 0x56, 0x6d, 0xc2, 0x12, 0x51, 0x6d, 0xea,
	0x21, 0x31, 0x60, 0x79, 0x48, 0x27, 0x2e, 0xa7, 0x7d, 0xa3, 0xa0, 0x40, 0xd1, 0x96, 0x1c, 0xc3,
	0xff, 0x69, 0x2e, 0x8f, 0x32, 0x3a, 0x40, 0x3f, 0xe0, 0x2c, 0x2a, 0x78, 0x3d, 0xc9, 0x79, 0x11,
	0x20, 0x14, 0xf5, 0x3e, 0xac, 0x61, 0x58, 0x58, 0x10, 0x51, 0x52, 0x11, 0xff, 0x45, 0x87, 0x0a,
	0x74, 0x0a, 0x65, 0x97, 0x5e, 0xa2, 0x2b, 0x8c, 0xa5, 0xbd, 0xe2, 0x81, 0xde, 0x7c, 0x66, 0xdd,
	0xa9, 0x96, 0xdd, 0xa3, 0x75, 0xae, 0xe0, 0x2d, 0x26, 0xfd, 0x89, 0x1d, 0xc6, 0x92, 0x4d, 0x58,
	0x12, 0x92, 0x4a, 0x34, 0xca, 0x2a, 0x45, 0xb0, 0x31, 0xdf, 0x80, 0x1e, 0x03, 0x93, 0x75, 0x28,
	0x5e, 0xe3, 0x24, 0xd4, 0x62, 0xba, 0x9c, 0x86, 0x8d, 0xa9, 0x3b, 0xc2, 0xb0, 0xf5, 0x60, 0xf3,
	0xb6, 0xf0, 0x5a, 0x6b, 0x38, 0x50, 0x49, 0xa6, 0x27, 0x04, 0x4a, 0x31, 0x29, 0xd5, 0xfa, 0x1e,
	0xf1, 0x9e, 0xc2, 0xb2, 0x47, 0x65, 0xef, 0x0a, 0x85, 0x51, 0x54, 0x7d, 0x55, 0x67, 0x7d, 0x5d,
	0x4c, 0x2f, 0xec, 0xe8, 0xbe, 0xf1, 0x5b, 0x83, 0x25, 0x75, 0x44, 0xce, 0xa1, 0x2a, 0xf8, 0xc8,
	0xef, 0x61, 0x57, 0xa0, 0x8b, 0x3d, 0xc9, 0x7d, 0x43, 0x53, 0xc1, 0xfb, 0xa9, 0x60, 0xab, 0xa3,
	0x60, 0x9d, 0x10, 0x15, 0x68, 0x51, 0x11, 0x89, 0x43, 0xf2, 0x1c, 0xca, 0x3e, 0x1f, 0x49, 0x14,
	0x46, 0x41, 0x91, 0x6c, 0xcd, 0x48, 0x4e, 0x51, 0x48, 0x87, 0x51, 0xe9, 0x70, 0x66, 0x87, 0x20,
	0xf3, 0x04, 0x36, 0x32, 0x58, 0x17, 0x12, 0xed, 0x8f, 0x06, 0x7a, 0x8c, 0x9a, 0x50, 0xd8, 0xec,
	0xcf, 0xb6, 0xe9, 0xa6, 0xac, 0xcc, 0x7a, 0xe2, 0xeb, 0x64, 0x7f, 0x1b, 0xfd, 0xf9, 0x1b, 0x52,
	0x83, 0xf2, 0x0d, 0x3a, 0x83, 0x2b, 0xa9, 0xaa, 0x59, 0xb3, 0xc3, 0x9d, 0x79, 0x06, 0x46, 0x1e,
	0xd1, 0x42, 0x2d, 0xfd, 0xd0, 0x60, 0xab, 0x9d, 0x35, 0xe1, 0x99, 0xf3, 0xb0, 0x0e, 0xc5, 0x91,
	0xef, 0x86, 0x2c, 0xd3, 0x25, 0x39, 0x82, 0x55, 0x97, 0x0a, 0xd9, 0x15, 0x88, 0x4c, 0x39, 0x46,
	0x6f, 0x9a, 0x56, 0x60, 0x70, 0x2b, 0x32, 0xb8, 0xf5, 0x21, 0x32, 0xb8, 0xbd, 0x32, 0x05, 0x77,
	0x10, 0xd9, 0x6c, 0xa2, 0x4b, 0xb1, 0x89, 0x6e, 0x1c, 0xc1, 0xba, 0x8d, 0x1e, 0x1f, 0x63, 0xbb,
	0xd3, 0xb2, 0xf1, 0xdb, 0x08, 0x85, 0x9c, 0xb7, 0x99, 0x36, 0x6f, 0xb3, 0xc6, 0x05, 0xd4, 0xcf,
	0x1c, 0xd6, 0x4f, 0xb6, 0x12, 0x31, 0x2c, 0xfc, 0x6a, 0x34, 0x7e, 0x15, 0xc1, 0xcc, 0xe2, 0x13,
	0x43, 0xce, 0x44, 0xc2, 0x17, 0x5a, 0xd2, 0x17, 0x27, 0x50, 0x4d, 0xa5, 0x52, 0x6a, 0xe9, 0x4d,
	0x23, 0xcf, 0xf7, 0x76, 0x25, 0x99, 0x9f, 0xdc, 0x82, 0x91, 0xf3, 0x2e, 0x45, 0x5e, 0x7b, 0x37,
	0xe3, 0xca, 0x2f, 0xd2, 0xca, 0xfc, 0xac, 0xe1, 0xbb, 0x52, 0xcb, 0x7c, 0xd5, 0x04, 0xf9, 0x0c,
	0xf5, 0x74, 0xee, 0x48, 0x66, 0x61, 0x94, 0x54, 0xf2, 0xbd, 0x87, 0x1e, 0x30, 0x7b, 0x9b, 0x65,
	0x9e, 0x0b, 0xf3, 0x2b, 0xec, 0xdc, 0x53, 0x54, 0xc6, 0xdc, 0xbe, 0x8a, 0xcf, 0xad, 0xde, 0xdc,
	0xcd, 0x4b, 0x1d, 0xf2, 0xc4, 0x07, 0xfb, 0x7b, 0x01, 0xaa, 0x6a, 0x88, 0x54, 0x40, 0xe0, 0xd7,
	0x8c, 0x8f, 0xa3, 0x2d, 0xf8, 0x71, 0x3e, 0xc2, 0x76, 0xce, 0xc7, 0x79, 0x6c, 0x8d, 0x5b, 0x99,
	0xd2, 0x93, 0x4f, 0x77, 0xc4, 0x69, 0xe1, 0x43, 0x5b, 0x3d, 0xac, 0x7b, 0x2d, 0x49, 0x10, 0x9d,
	0x37, 0x7f, 0x6a, 0xe9, 0xdf, 0xd3, 0x50, 0x95, 0x09, 0x79, 0x0f, 0x7a, 0xb0, 0x46, 0xbf, 0xdd,
	0x69, 0x91, 0x7a, 0x2c, 0x47, 0x52, 0x3b, 0x33, 0xff, 0x8a, 0x1c, 0xc3, 0xea, 0x9d, 0x69, 0x89,
	0x39, 0xc3, 0xa5, 0x9d, 0x6c, 0xd6, 0xe6, 0x5e, 0x86, 0xd6, 0xf4, 0x7f, 0x41, 0xf3, 0x16, 0xb6,
	0x93, 0xf5, 0x9d, 0x3a, 0xa2, 0xc7, 0xc7, 0xe8, 0x4f, 0x48, 0x17, 0xc8, 0xfc, 0x88, 0x93, 0xfd,
	0xfb, 0x0d, 0x10, 0x64, 0x7b, 0xf2, 0x18, 0x97, 0x5c, 0x96, 0x55, 0x2d, 0x2f, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x25, 0xdb, 0x5d, 0x52, 0xcd, 0x08, 0x00, 0x00,
}
