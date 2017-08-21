// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service.proto

/*
Package push is a generated protocol buffer package.

It is generated from these files:
	proto/service.proto

It has these top-level messages:
	FirebaseCloudMessagingApp
	Id
	Topic
	Subscription
	FirebaseCloudMessagingEndpoint
	SetFirebaseCloudMessagingEndpoint
	DeleteFirebaseCloudMessagingEndpoint
	FirebaseCloudMessagingNotification
	Log
*/
package push

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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

type FirebaseCloudMessagingApp struct {
	ApiKey string `protobuf:"bytes,1,opt,name=api_key,json=apiKey" json:"api_key,omitempty"`
}

func (m *FirebaseCloudMessagingApp) Reset()                    { *m = FirebaseCloudMessagingApp{} }
func (m *FirebaseCloudMessagingApp) String() string            { return proto.CompactTextString(m) }
func (*FirebaseCloudMessagingApp) ProtoMessage()               {}
func (*FirebaseCloudMessagingApp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FirebaseCloudMessagingApp) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

type Id struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Topic struct {
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Topic) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Subscription struct {
	Topic *Topic `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	Key   *Id    `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Subscription) GetTopic() *Topic {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *Subscription) GetKey() *Id {
	if m != nil {
		return m.Key
	}
	return nil
}

type FirebaseCloudMessagingEndpoint struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *FirebaseCloudMessagingEndpoint) Reset()                    { *m = FirebaseCloudMessagingEndpoint{} }
func (m *FirebaseCloudMessagingEndpoint) String() string            { return proto.CompactTextString(m) }
func (*FirebaseCloudMessagingEndpoint) ProtoMessage()               {}
func (*FirebaseCloudMessagingEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FirebaseCloudMessagingEndpoint) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SetFirebaseCloudMessagingEndpoint struct {
	Key   *Id                             `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value *FirebaseCloudMessagingEndpoint `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *SetFirebaseCloudMessagingEndpoint) Reset()         { *m = SetFirebaseCloudMessagingEndpoint{} }
func (m *SetFirebaseCloudMessagingEndpoint) String() string { return proto.CompactTextString(m) }
func (*SetFirebaseCloudMessagingEndpoint) ProtoMessage()    {}
func (*SetFirebaseCloudMessagingEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5}
}

func (m *SetFirebaseCloudMessagingEndpoint) GetKey() *Id {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SetFirebaseCloudMessagingEndpoint) GetValue() *FirebaseCloudMessagingEndpoint {
	if m != nil {
		return m.Value
	}
	return nil
}

type DeleteFirebaseCloudMessagingEndpoint struct {
	Value *FirebaseCloudMessagingEndpoint `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *DeleteFirebaseCloudMessagingEndpoint) Reset()         { *m = DeleteFirebaseCloudMessagingEndpoint{} }
func (m *DeleteFirebaseCloudMessagingEndpoint) String() string { return proto.CompactTextString(m) }
func (*DeleteFirebaseCloudMessagingEndpoint) ProtoMessage()    {}
func (*DeleteFirebaseCloudMessagingEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6}
}

func (m *DeleteFirebaseCloudMessagingEndpoint) GetValue() *FirebaseCloudMessagingEndpoint {
	if m != nil {
		return m.Value
	}
	return nil
}

type FirebaseCloudMessagingNotification struct {
	Topic     string                                      `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	Condition string                                      `protobuf:"bytes,2,opt,name=condition" json:"condition,omitempty"`
	Endpoint  []*FirebaseCloudMessagingEndpoint           `protobuf:"bytes,3,rep,name=endpoint" json:"endpoint,omitempty"`
	Payload   *FirebaseCloudMessagingNotification_Payload `protobuf:"bytes,4,opt,name=payload" json:"payload,omitempty"`
}

func (m *FirebaseCloudMessagingNotification) Reset()         { *m = FirebaseCloudMessagingNotification{} }
func (m *FirebaseCloudMessagingNotification) String() string { return proto.CompactTextString(m) }
func (*FirebaseCloudMessagingNotification) ProtoMessage()    {}
func (*FirebaseCloudMessagingNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7}
}

func (m *FirebaseCloudMessagingNotification) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *FirebaseCloudMessagingNotification) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

func (m *FirebaseCloudMessagingNotification) GetEndpoint() []*FirebaseCloudMessagingEndpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *FirebaseCloudMessagingNotification) GetPayload() *FirebaseCloudMessagingNotification_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type FirebaseCloudMessagingNotification_Payload struct {
	Json string `protobuf:"bytes,1,opt,name=json" json:"json,omitempty"`
}

func (m *FirebaseCloudMessagingNotification_Payload) Reset() {
	*m = FirebaseCloudMessagingNotification_Payload{}
}
func (m *FirebaseCloudMessagingNotification_Payload) String() string {
	return proto.CompactTextString(m)
}
func (*FirebaseCloudMessagingNotification_Payload) ProtoMessage() {}
func (*FirebaseCloudMessagingNotification_Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

func (m *FirebaseCloudMessagingNotification_Payload) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

type Log struct {
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Text      string                      `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Log) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Log) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*FirebaseCloudMessagingApp)(nil), "push.FirebaseCloudMessagingApp")
	proto.RegisterType((*Id)(nil), "push.Id")
	proto.RegisterType((*Topic)(nil), "push.Topic")
	proto.RegisterType((*Subscription)(nil), "push.Subscription")
	proto.RegisterType((*FirebaseCloudMessagingEndpoint)(nil), "push.FirebaseCloudMessagingEndpoint")
	proto.RegisterType((*SetFirebaseCloudMessagingEndpoint)(nil), "push.SetFirebaseCloudMessagingEndpoint")
	proto.RegisterType((*DeleteFirebaseCloudMessagingEndpoint)(nil), "push.DeleteFirebaseCloudMessagingEndpoint")
	proto.RegisterType((*FirebaseCloudMessagingNotification)(nil), "push.FirebaseCloudMessagingNotification")
	proto.RegisterType((*FirebaseCloudMessagingNotification_Payload)(nil), "push.FirebaseCloudMessagingNotification.Payload")
	proto.RegisterType((*Log)(nil), "push.Log")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AppService service

type AppServiceClient interface {
	SetFirebaseCloudMessaging(ctx context.Context, in *FirebaseCloudMessagingApp, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	GetFirebaseCloudMessaging(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*FirebaseCloudMessagingApp, error)
}

type appServiceClient struct {
	cc *grpc.ClientConn
}

func NewAppServiceClient(cc *grpc.ClientConn) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) SetFirebaseCloudMessaging(ctx context.Context, in *FirebaseCloudMessagingApp, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/push.AppService/SetFirebaseCloudMessaging", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetFirebaseCloudMessaging(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*FirebaseCloudMessagingApp, error) {
	out := new(FirebaseCloudMessagingApp)
	err := grpc.Invoke(ctx, "/push.AppService/GetFirebaseCloudMessaging", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppService service

type AppServiceServer interface {
	SetFirebaseCloudMessaging(context.Context, *FirebaseCloudMessagingApp) (*google_protobuf.Empty, error)
	GetFirebaseCloudMessaging(context.Context, *google_protobuf.Empty) (*FirebaseCloudMessagingApp, error)
}

func RegisterAppServiceServer(s *grpc.Server, srv AppServiceServer) {
	s.RegisterService(&_AppService_serviceDesc, srv)
}

func _AppService_SetFirebaseCloudMessaging_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FirebaseCloudMessagingApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).SetFirebaseCloudMessaging(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.AppService/SetFirebaseCloudMessaging",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).SetFirebaseCloudMessaging(ctx, req.(*FirebaseCloudMessagingApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetFirebaseCloudMessaging_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetFirebaseCloudMessaging(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.AppService/GetFirebaseCloudMessaging",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetFirebaseCloudMessaging(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "push.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetFirebaseCloudMessaging",
			Handler:    _AppService_SetFirebaseCloudMessaging_Handler,
		},
		{
			MethodName: "GetFirebaseCloudMessaging",
			Handler:    _AppService_GetFirebaseCloudMessaging_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}

// Client API for SubscriptionService service

type SubscriptionServiceClient interface {
	Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Unsubscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Get(ctx context.Context, in *Topic, opts ...grpc.CallOption) (SubscriptionService_GetClient, error)
}

type subscriptionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSubscriptionServiceClient(cc *grpc.ClientConn) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/push.SubscriptionService/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) Unsubscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/push.SubscriptionService/Unsubscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) Get(ctx context.Context, in *Topic, opts ...grpc.CallOption) (SubscriptionService_GetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SubscriptionService_serviceDesc.Streams[0], c.cc, "/push.SubscriptionService/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionServiceGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubscriptionService_GetClient interface {
	Recv() (*Id, error)
	grpc.ClientStream
}

type subscriptionServiceGetClient struct {
	grpc.ClientStream
}

func (x *subscriptionServiceGetClient) Recv() (*Id, error) {
	m := new(Id)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SubscriptionService service

type SubscriptionServiceServer interface {
	Subscribe(context.Context, *Subscription) (*google_protobuf.Empty, error)
	Unsubscribe(context.Context, *Subscription) (*google_protobuf.Empty, error)
	Get(*Topic, SubscriptionService_GetServer) error
}

func RegisterSubscriptionServiceServer(s *grpc.Server, srv SubscriptionServiceServer) {
	s.RegisterService(&_SubscriptionService_serviceDesc, srv)
}

func _SubscriptionService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.SubscriptionService/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).Subscribe(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.SubscriptionService/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).Unsubscribe(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Topic)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionServiceServer).Get(m, &subscriptionServiceGetServer{stream})
}

type SubscriptionService_GetServer interface {
	Send(*Id) error
	grpc.ServerStream
}

type subscriptionServiceGetServer struct {
	grpc.ServerStream
}

func (x *subscriptionServiceGetServer) Send(m *Id) error {
	return x.ServerStream.SendMsg(m)
}

var _SubscriptionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "push.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _SubscriptionService_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _SubscriptionService_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _SubscriptionService_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}

// Client API for FirebaseCloudMessagingEndpointService service

type FirebaseCloudMessagingEndpointServiceClient interface {
	Set(ctx context.Context, in *SetFirebaseCloudMessagingEndpoint, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Delete(ctx context.Context, in *DeleteFirebaseCloudMessagingEndpoint, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (FirebaseCloudMessagingEndpointService_GetClient, error)
}

type firebaseCloudMessagingEndpointServiceClient struct {
	cc *grpc.ClientConn
}

func NewFirebaseCloudMessagingEndpointServiceClient(cc *grpc.ClientConn) FirebaseCloudMessagingEndpointServiceClient {
	return &firebaseCloudMessagingEndpointServiceClient{cc}
}

func (c *firebaseCloudMessagingEndpointServiceClient) Set(ctx context.Context, in *SetFirebaseCloudMessagingEndpoint, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/push.FirebaseCloudMessagingEndpointService/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firebaseCloudMessagingEndpointServiceClient) Delete(ctx context.Context, in *DeleteFirebaseCloudMessagingEndpoint, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/push.FirebaseCloudMessagingEndpointService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firebaseCloudMessagingEndpointServiceClient) Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (FirebaseCloudMessagingEndpointService_GetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FirebaseCloudMessagingEndpointService_serviceDesc.Streams[0], c.cc, "/push.FirebaseCloudMessagingEndpointService/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &firebaseCloudMessagingEndpointServiceGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FirebaseCloudMessagingEndpointService_GetClient interface {
	Recv() (*FirebaseCloudMessagingEndpoint, error)
	grpc.ClientStream
}

type firebaseCloudMessagingEndpointServiceGetClient struct {
	grpc.ClientStream
}

func (x *firebaseCloudMessagingEndpointServiceGetClient) Recv() (*FirebaseCloudMessagingEndpoint, error) {
	m := new(FirebaseCloudMessagingEndpoint)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FirebaseCloudMessagingEndpointService service

type FirebaseCloudMessagingEndpointServiceServer interface {
	Set(context.Context, *SetFirebaseCloudMessagingEndpoint) (*google_protobuf.Empty, error)
	Delete(context.Context, *DeleteFirebaseCloudMessagingEndpoint) (*google_protobuf.Empty, error)
	Get(*Id, FirebaseCloudMessagingEndpointService_GetServer) error
}

func RegisterFirebaseCloudMessagingEndpointServiceServer(s *grpc.Server, srv FirebaseCloudMessagingEndpointServiceServer) {
	s.RegisterService(&_FirebaseCloudMessagingEndpointService_serviceDesc, srv)
}

func _FirebaseCloudMessagingEndpointService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFirebaseCloudMessagingEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirebaseCloudMessagingEndpointServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.FirebaseCloudMessagingEndpointService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirebaseCloudMessagingEndpointServiceServer).Set(ctx, req.(*SetFirebaseCloudMessagingEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirebaseCloudMessagingEndpointService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFirebaseCloudMessagingEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirebaseCloudMessagingEndpointServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.FirebaseCloudMessagingEndpointService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirebaseCloudMessagingEndpointServiceServer).Delete(ctx, req.(*DeleteFirebaseCloudMessagingEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirebaseCloudMessagingEndpointService_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Id)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FirebaseCloudMessagingEndpointServiceServer).Get(m, &firebaseCloudMessagingEndpointServiceGetServer{stream})
}

type FirebaseCloudMessagingEndpointService_GetServer interface {
	Send(*FirebaseCloudMessagingEndpoint) error
	grpc.ServerStream
}

type firebaseCloudMessagingEndpointServiceGetServer struct {
	grpc.ServerStream
}

func (x *firebaseCloudMessagingEndpointServiceGetServer) Send(m *FirebaseCloudMessagingEndpoint) error {
	return x.ServerStream.SendMsg(m)
}

var _FirebaseCloudMessagingEndpointService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "push.FirebaseCloudMessagingEndpointService",
	HandlerType: (*FirebaseCloudMessagingEndpointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _FirebaseCloudMessagingEndpointService_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FirebaseCloudMessagingEndpointService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _FirebaseCloudMessagingEndpointService_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}

// Client API for FirebaseCloudMessagingService service

type FirebaseCloudMessagingServiceClient interface {
	Send(ctx context.Context, in *FirebaseCloudMessagingNotification, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type firebaseCloudMessagingServiceClient struct {
	cc *grpc.ClientConn
}

func NewFirebaseCloudMessagingServiceClient(cc *grpc.ClientConn) FirebaseCloudMessagingServiceClient {
	return &firebaseCloudMessagingServiceClient{cc}
}

func (c *firebaseCloudMessagingServiceClient) Send(ctx context.Context, in *FirebaseCloudMessagingNotification, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/push.FirebaseCloudMessagingService/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FirebaseCloudMessagingService service

type FirebaseCloudMessagingServiceServer interface {
	Send(context.Context, *FirebaseCloudMessagingNotification) (*google_protobuf.Empty, error)
}

func RegisterFirebaseCloudMessagingServiceServer(s *grpc.Server, srv FirebaseCloudMessagingServiceServer) {
	s.RegisterService(&_FirebaseCloudMessagingService_serviceDesc, srv)
}

func _FirebaseCloudMessagingService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FirebaseCloudMessagingNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirebaseCloudMessagingServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.FirebaseCloudMessagingService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirebaseCloudMessagingServiceServer).Send(ctx, req.(*FirebaseCloudMessagingNotification))
	}
	return interceptor(ctx, in, info, handler)
}

var _FirebaseCloudMessagingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "push.FirebaseCloudMessagingService",
	HandlerType: (*FirebaseCloudMessagingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _FirebaseCloudMessagingService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}

// Client API for LogService service

type LogServiceClient interface {
	Info(ctx context.Context, in *Log, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Tail(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (LogService_TailClient, error)
}

type logServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogServiceClient(cc *grpc.ClientConn) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) Info(ctx context.Context, in *Log, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/push.LogService/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) Tail(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (LogService_TailClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_LogService_serviceDesc.Streams[0], c.cc, "/push.LogService/Tail", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceTailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_TailClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type logServiceTailClient struct {
	grpc.ClientStream
}

func (x *logServiceTailClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for LogService service

type LogServiceServer interface {
	Info(context.Context, *Log) (*google_protobuf.Empty, error)
	Tail(*google_protobuf.Empty, LogService_TailServer) error
}

func RegisterLogServiceServer(s *grpc.Server, srv LogServiceServer) {
	s.RegisterService(&_LogService_serviceDesc, srv)
}

func _LogService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.LogService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).Info(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_Tail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(google_protobuf.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).Tail(m, &logServiceTailServer{stream})
}

type LogService_TailServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type logServiceTailServer struct {
	grpc.ServerStream
}

func (x *logServiceTailServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

var _LogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "push.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _LogService_Info_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tail",
			Handler:       _LogService_Tail_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}

func init() { proto.RegisterFile("proto/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x5d, 0xfe, 0xec, 0x4f, 0xee, 0x7e, 0xfa, 0x3d, 0x78, 0x13, 0x74, 0x19, 0x63, 0x9b, 0x35,
	0xc4, 0x04, 0x22, 0x9b, 0x0a, 0x42, 0x0c, 0x5e, 0x98, 0x60, 0x8c, 0x8d, 0x6d, 0x42, 0x49, 0xe1,
	0x15, 0x25, 0x8d, 0x9b, 0x99, 0xa6, 0xb1, 0xd5, 0xb8, 0x13, 0x15, 0xcf, 0x7c, 0x1e, 0xbe, 0x00,
	0x1f, 0x8d, 0x07, 0x64, 0x3b, 0x49, 0x2b, 0x68, 0xd3, 0xc2, 0x9b, 0xe3, 0x7b, 0xee, 0xb9, 0xe7,
	0x38, 0xf7, 0xc0, 0x1a, 0xef, 0x33, 0xc1, 0x0e, 0x72, 0xd2, 0xbf, 0xa1, 0x6d, 0xe2, 0xa9, 0x2f,
	0x64, 0xf3, 0x41, 0x7e, 0xed, 0x6e, 0x26, 0x8c, 0x25, 0x29, 0x39, 0x50, 0x77, 0xd1, 0xa0, 0x73,
	0x40, 0x7a, 0x5c, 0x0c, 0x35, 0xc4, 0xdd, 0xfe, 0xbd, 0x28, 0x68, 0x8f, 0xe4, 0x22, 0xec, 0x71,
	0x0d, 0xc0, 0x4f, 0x60, 0xe3, 0x0d, 0xed, 0x93, 0x28, 0xcc, 0xc9, 0xab, 0x94, 0x0d, 0xe2, 0x4b,
	0x92, 0xe7, 0x61, 0x42, 0xb3, 0xe4, 0x98, 0x73, 0x74, 0x1b, 0x96, 0x43, 0x4e, 0x3f, 0x75, 0xc9,
	0xb0, 0x61, 0xec, 0x18, 0xfb, 0x8e, 0xbf, 0x14, 0x72, 0xfa, 0x8e, 0x0c, 0xf1, 0x3a, 0x98, 0x67,
	0x31, 0xfa, 0x1f, 0x4c, 0x1a, 0x17, 0x15, 0x93, 0xc6, 0x78, 0x13, 0x16, 0x5b, 0x8c, 0xd3, 0x36,
	0x42, 0x60, 0x67, 0x61, 0x8f, 0x34, 0x4c, 0x55, 0x52, 0x67, 0x7c, 0x09, 0xff, 0x05, 0x83, 0x28,
	0x6f, 0xf7, 0x29, 0x17, 0x94, 0x65, 0x68, 0x17, 0x16, 0x85, 0x04, 0xab, 0xfe, 0xd5, 0xe6, 0xaa,
	0x27, 0xcd, 0x78, 0xaa, 0xdf, 0xd7, 0x15, 0xe4, 0x82, 0x25, 0x47, 0x9b, 0x0a, 0xb0, 0xa2, 0x01,
	0x67, 0xb1, 0x2f, 0x2f, 0xf1, 0x53, 0xb8, 0x3b, 0x59, 0xf7, 0x49, 0x16, 0x73, 0x46, 0x33, 0x81,
	0xd6, 0xe5, 0x80, 0x2e, 0xc9, 0x0a, 0x81, 0xfa, 0x03, 0x7f, 0x85, 0xdd, 0x80, 0x88, 0x19, 0xad,
	0xc5, 0x60, 0x63, 0xc2, 0x60, 0xf4, 0x1c, 0x16, 0x6f, 0xc2, 0x74, 0x40, 0x0a, 0x59, 0x7b, 0xba,
	0x5a, 0x4f, 0xe8, 0xeb, 0x16, 0x1c, 0xc1, 0xde, 0x6b, 0x92, 0x12, 0x41, 0x66, 0xcc, 0xaf, 0x66,
	0x18, 0x7f, 0x3f, 0xe3, 0x9b, 0x09, 0x78, 0x32, 0xf2, 0x8a, 0x09, 0xda, 0xa1, 0xed, 0x50, 0x3d,
	0xff, 0xfa, 0xf8, 0xf3, 0x3b, 0xe5, 0x8b, 0xdf, 0x01, 0xa7, 0xcd, 0xb2, 0x98, 0x4a, 0x48, 0xf1,
	0xf7, 0x46, 0x17, 0xe8, 0x25, 0xac, 0x90, 0x62, 0x5a, 0xc3, 0xda, 0xb1, 0xe6, 0x56, 0x56, 0x75,
	0xa1, 0x73, 0x58, 0xe6, 0xe1, 0x30, 0x65, 0x61, 0xdc, 0xb0, 0x95, 0xb5, 0xc3, 0x3a, 0x82, 0x71,
	0xc1, 0xde, 0x7b, 0xdd, 0xe7, 0x97, 0x04, 0xee, 0x16, 0x2c, 0x17, 0x77, 0x72, 0xdf, 0x3e, 0xe7,
	0xac, 0xfc, 0xd3, 0xea, 0x8c, 0x03, 0xb0, 0x2e, 0x58, 0x82, 0x9e, 0x81, 0x53, 0xad, 0x7c, 0xf1,
	0x9c, 0xae, 0xa7, 0x43, 0xe1, 0x95, 0xa1, 0xf0, 0x5a, 0x25, 0xc2, 0x1f, 0x81, 0x25, 0xa9, 0x20,
	0x5f, 0x44, 0xb9, 0xc4, 0xf2, 0xdc, 0xfc, 0x61, 0x00, 0x1c, 0x73, 0x1e, 0xe8, 0x18, 0xa2, 0x8f,
	0xb0, 0x31, 0x75, 0x99, 0xd0, 0x76, 0x9d, 0xb5, 0x63, 0xce, 0xdd, 0x5b, 0x7f, 0xe8, 0x38, 0x91,
	0xc9, 0xc5, 0x0b, 0x92, 0xf7, 0x74, 0x2a, 0xef, 0x94, 0x36, 0x77, 0xd6, 0x3c, 0xbc, 0xd0, 0xfc,
	0x6e, 0xc0, 0xda, 0x78, 0x08, 0x4b, 0x1f, 0x47, 0xe0, 0x14, 0xd7, 0x11, 0x41, 0x48, 0xf3, 0x8c,
	0xe3, 0x6a, 0xa4, 0xbe, 0x80, 0xd5, 0x0f, 0x59, 0xfe, 0x8f, 0xcd, 0x3b, 0x60, 0x9d, 0x12, 0x81,
	0xc6, 0xb3, 0xef, 0x56, 0x71, 0xc3, 0x0b, 0x87, 0x46, 0xf3, 0xa7, 0x01, 0xf7, 0xea, 0xb7, 0xab,
	0xf4, 0xf0, 0x16, 0xac, 0x80, 0x08, 0x74, 0xbf, 0x10, 0x30, 0x2b, 0xe3, 0x35, 0xaa, 0xae, 0x60,
	0x49, 0xa7, 0x14, 0x3d, 0xd0, 0x64, 0xf3, 0x64, 0xb6, 0x86, 0xef, 0x48, 0xbb, 0xac, 0x8c, 0xb9,
	0x73, 0xa5, 0x46, 0xd9, 0xef, 0xc2, 0xd6, 0x64, 0x54, 0xe9, 0xfa, 0x1c, 0xec, 0x80, 0x64, 0x31,
	0xda, 0x9f, 0x37, 0x47, 0xd3, 0x75, 0x36, 0xaf, 0x01, 0x2e, 0x58, 0xc5, 0xfc, 0x10, 0xec, 0xb3,
	0xac, 0xc3, 0x90, 0xa3, 0x99, 0x2f, 0x58, 0x52, 0x63, 0xf1, 0x11, 0xd8, 0xad, 0x90, 0xa6, 0x53,
	0x77, 0x73, 0x44, 0x22, 0x6d, 0x45, 0x4b, 0xaa, 0xfc, 0xf8, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2b, 0xd9, 0xa8, 0xef, 0xd6, 0x06, 0x00, 0x00,
}
