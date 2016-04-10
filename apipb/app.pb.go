// Code generated by protoc-gen-go.
// source: app.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	app.proto
	auth.proto
	deployment.proto

It has these top-level messages:
	App
	ConfigRequest
	ConfigResponse
	AppCreateResponse
	User
	AuthResponse
	DeploymentRequest
	DeploymentJob
	DeploymentResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion1

// App represents an app
type App struct {
	Name   string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Owner  string            `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
	Config map[string]string `protobuf:"bytes,3,rep,name=config" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *App) Reset()                    { *m = App{} }
func (m *App) String() string            { return proto.CompactTextString(m) }
func (*App) ProtoMessage()               {}
func (*App) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *App) GetConfig() map[string]string {
	if m != nil {
		return m.Config
	}
	return nil
}

// ConfigRequest represents a config request
type ConfigRequest struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *ConfigRequest) Reset()                    { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()               {}
func (*ConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// ConfigResponse represents a config response
type ConfigResponse struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ConfigResponse) Reset()                    { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()               {}
func (*ConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// AppCreateResponse represents an app creation response
type AppCreateResponse struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AppCreateResponse) Reset()                    { *m = AppCreateResponse{} }
func (m *AppCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*AppCreateResponse) ProtoMessage()               {}
func (*AppCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*App)(nil), "api.App")
	proto.RegisterType((*ConfigRequest)(nil), "api.ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "api.ConfigResponse")
	proto.RegisterType((*AppCreateResponse)(nil), "api.AppCreateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for AppService service

type AppServiceClient interface {
	CreateApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*AppCreateResponse, error)
	GetApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*App, error)
	GetAppConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	SetAppConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
}

type appServiceClient struct {
	cc *grpc.ClientConn
}

func NewAppServiceClient(cc *grpc.ClientConn) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) CreateApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*AppCreateResponse, error) {
	out := new(AppCreateResponse)
	err := grpc.Invoke(ctx, "/api.AppService/CreateApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*App, error) {
	out := new(App)
	err := grpc.Invoke(ctx, "/api.AppService/GetApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetAppConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := grpc.Invoke(ctx, "/api.AppService/GetAppConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) SetAppConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := grpc.Invoke(ctx, "/api.AppService/SetAppConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppService service

type AppServiceServer interface {
	CreateApp(context.Context, *App) (*AppCreateResponse, error)
	GetApp(context.Context, *App) (*App, error)
	GetAppConfig(context.Context, *ConfigRequest) (*ConfigResponse, error)
	SetAppConfig(context.Context, *ConfigRequest) (*ConfigResponse, error)
}

func RegisterAppServiceServer(s *grpc.Server, srv AppServiceServer) {
	s.RegisterService(&_AppService_serviceDesc, srv)
}

func _AppService_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(App)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AppServiceServer).CreateApp(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _AppService_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(App)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AppServiceServer).GetApp(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _AppService_GetAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AppServiceServer).GetAppConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _AppService_SetAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AppServiceServer).SetAppConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _AppService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApp",
			Handler:    _AppService_CreateApp_Handler,
		},
		{
			MethodName: "GetApp",
			Handler:    _AppService_GetApp_Handler,
		},
		{
			MethodName: "GetAppConfig",
			Handler:    _AppService_GetAppConfig_Handler,
		},
		{
			MethodName: "SetAppConfig",
			Handler:    _AppService_SetAppConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0x28, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x9a, 0xc1, 0xc8, 0xc5, 0xec,
	0x58, 0x50, 0x20, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0x66, 0x0b, 0x89, 0x70, 0xb1, 0xe6, 0x97, 0xe7, 0xa5, 0x16, 0x49, 0x30, 0x81, 0x05,
	0x21, 0x1c, 0x21, 0x1d, 0x2e, 0xb6, 0xe4, 0xfc, 0xbc, 0xb4, 0xcc, 0x74, 0x09, 0x66, 0x05, 0x66,
	0x0d, 0x6e, 0x23, 0x11, 0x3d, 0xa0, 0x39, 0x7a, 0x40, 0x33, 0xf4, 0x9c, 0xc1, 0xc2, 0xae, 0x79,
	0x25, 0x45, 0x95, 0x41, 0x50, 0x35, 0x52, 0x96, 0x5c, 0xdc, 0x48, 0xc2, 0x42, 0x02, 0x5c, 0xcc,
	0xd9, 0xa9, 0x95, 0x50, 0x5b, 0x40, 0x4c, 0x90, 0x25, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x30, 0x4b,
	0xc0, 0x1c, 0x2b, 0x26, 0x0b, 0x46, 0x25, 0x6f, 0x2e, 0x5e, 0x88, 0xd6, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0xac, 0x6e, 0x84, 0x1a, 0xc8, 0x84, 0xc5, 0x40, 0x66, 0x24, 0x03, 0x95, 0x2c,
	0xb8, 0xf8, 0x60, 0x86, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x12, 0xeb, 0x14, 0x25, 0x75, 0x2e,
	0x41, 0xa0, 0xe7, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0xe1, 0x9a, 0xb1, 0x38, 0xc5, 0xe8, 0x32,
	0x23, 0x17, 0x17, 0x50, 0x65, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x3e, 0x17, 0x27,
	0x44, 0x13, 0x28, 0x78, 0x39, 0x60, 0x81, 0x24, 0x25, 0x06, 0x63, 0xa1, 0x9a, 0xa8, 0xc4, 0x20,
	0x24, 0xc7, 0xc5, 0xe6, 0x9e, 0x5a, 0x82, 0xaa, 0x1a, 0xce, 0x02, 0xca, 0x5b, 0x72, 0xf1, 0x40,
	0xe4, 0x21, 0x1e, 0x11, 0x12, 0x02, 0xcb, 0xa1, 0x04, 0x91, 0x94, 0x30, 0x8a, 0x18, 0xdc, 0x68,
	0xa0, 0xd6, 0x60, 0xf2, 0xb4, 0x26, 0xb1, 0x81, 0x13, 0x8b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xc4, 0xca, 0x2b, 0x8d, 0x39, 0x02, 0x00, 0x00,
}
