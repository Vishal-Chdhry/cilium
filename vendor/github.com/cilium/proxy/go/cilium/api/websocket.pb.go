// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.21.7
// source: cilium/api/websocket.proto

package cilium

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type WebSocketClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the unix domain socket for the cilium access log, if any.
	AccessLogPath string `protobuf:"bytes,1,opt,name=access_log_path,json=accessLogPath,proto3" json:"access_log_path,omitempty"`
	// Host header value, required.
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// Path value. Defaults to "/".
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// sec-websocket-key value to use, defaults to a random key.
	Key string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	// Websocket version, defaults to "13".
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	// Origin header, if any.
	Origin string `protobuf:"bytes,6,opt,name=origin,proto3" json:"origin,omitempty"`
	// Websocket handshake timeout, default is 5 seconds.
	HandshakeTimeout *durationpb.Duration `protobuf:"bytes,7,opt,name=handshake_timeout,json=handshakeTimeout,proto3" json:"handshake_timeout,omitempty"`
	// ping interval, default is 0 (disabled).
	// Connection is assumed dead if response is not received before the next ping is to be sent.
	PingInterval *durationpb.Duration `protobuf:"bytes,8,opt,name=ping_interval,json=pingInterval,proto3" json:"ping_interval,omitempty"`
	// ping only on when idle on both directions.
	// ping_interval must be non-zero when this is true.
	PingWhenIdle bool `protobuf:"varint,9,opt,name=ping_when_idle,json=pingWhenIdle,proto3" json:"ping_when_idle,omitempty"`
}

func (x *WebSocketClient) Reset() {
	*x = WebSocketClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cilium_api_websocket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebSocketClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebSocketClient) ProtoMessage() {}

func (x *WebSocketClient) ProtoReflect() protoreflect.Message {
	mi := &file_cilium_api_websocket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebSocketClient.ProtoReflect.Descriptor instead.
func (*WebSocketClient) Descriptor() ([]byte, []int) {
	return file_cilium_api_websocket_proto_rawDescGZIP(), []int{0}
}

func (x *WebSocketClient) GetAccessLogPath() string {
	if x != nil {
		return x.AccessLogPath
	}
	return ""
}

func (x *WebSocketClient) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *WebSocketClient) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *WebSocketClient) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *WebSocketClient) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *WebSocketClient) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *WebSocketClient) GetHandshakeTimeout() *durationpb.Duration {
	if x != nil {
		return x.HandshakeTimeout
	}
	return nil
}

func (x *WebSocketClient) GetPingInterval() *durationpb.Duration {
	if x != nil {
		return x.PingInterval
	}
	return nil
}

func (x *WebSocketClient) GetPingWhenIdle() bool {
	if x != nil {
		return x.PingWhenIdle
	}
	return false
}

type WebSocketServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the unix domain socket for the cilium access log, if any.
	AccessLogPath string `protobuf:"bytes,1,opt,name=access_log_path,json=accessLogPath,proto3" json:"access_log_path,omitempty"`
	// Expected host header value, if any.
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// Expected path value, if any.
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// sec-websocket-key value to expect, if any.
	Key string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	// Websocket version, ignored if omitted.
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	// Origin header, if any. Origin header is not allowed if omitted.
	Origin string `protobuf:"bytes,6,opt,name=origin,proto3" json:"origin,omitempty"`
	// Websocket handshake timeout, default is 5 seconds.
	HandshakeTimeout *durationpb.Duration `protobuf:"bytes,7,opt,name=handshake_timeout,json=handshakeTimeout,proto3" json:"handshake_timeout,omitempty"`
	// ping interval, default is 0 (disabled).
	// Connection is assumed dead if response is not received before the next ping is to be sent.
	PingInterval *durationpb.Duration `protobuf:"bytes,8,opt,name=ping_interval,json=pingInterval,proto3" json:"ping_interval,omitempty"`
	// ping only on when idle on both directions.
	// ping_interval must be non-zero when this is true.
	PingWhenIdle bool `protobuf:"varint,9,opt,name=ping_when_idle,json=pingWhenIdle,proto3" json:"ping_when_idle,omitempty"`
}

func (x *WebSocketServer) Reset() {
	*x = WebSocketServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cilium_api_websocket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebSocketServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebSocketServer) ProtoMessage() {}

func (x *WebSocketServer) ProtoReflect() protoreflect.Message {
	mi := &file_cilium_api_websocket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebSocketServer.ProtoReflect.Descriptor instead.
func (*WebSocketServer) Descriptor() ([]byte, []int) {
	return file_cilium_api_websocket_proto_rawDescGZIP(), []int{1}
}

func (x *WebSocketServer) GetAccessLogPath() string {
	if x != nil {
		return x.AccessLogPath
	}
	return ""
}

func (x *WebSocketServer) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *WebSocketServer) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *WebSocketServer) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *WebSocketServer) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *WebSocketServer) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *WebSocketServer) GetHandshakeTimeout() *durationpb.Duration {
	if x != nil {
		return x.HandshakeTimeout
	}
	return nil
}

func (x *WebSocketServer) GetPingInterval() *durationpb.Duration {
	if x != nil {
		return x.PingInterval
	}
	return nil
}

func (x *WebSocketServer) GetPingWhenIdle() bool {
	if x != nil {
		return x.PingWhenIdle
	}
	return false
}

var File_cilium_api_websocket_proto protoreflect.FileDescriptor

var file_cilium_api_websocket_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x69,
	0x6c, 0x69, 0x75, 0x6d, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x02,
	0x0a, 0x0f, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x46,
	0x0a, 0x11, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x77,
	0x68, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x70, 0x69, 0x6e, 0x67, 0x57, 0x68, 0x65, 0x6e, 0x49, 0x64, 0x6c, 0x65, 0x22, 0xd3, 0x02, 0x0a,
	0x0f, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x12, 0x46, 0x0a, 0x11, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x68, 0x61, 0x6e, 0x64,
	0x73, 0x68, 0x61, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d,
	0x70, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x70, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0e,
	0x70, 0x69, 0x6e, 0x67, 0x5f, 0x77, 0x68, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x70, 0x69, 0x6e, 0x67, 0x57, 0x68, 0x65, 0x6e, 0x49, 0x64,
	0x6c, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x63, 0x69, 0x6c, 0x69,
	0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cilium_api_websocket_proto_rawDescOnce sync.Once
	file_cilium_api_websocket_proto_rawDescData = file_cilium_api_websocket_proto_rawDesc
)

func file_cilium_api_websocket_proto_rawDescGZIP() []byte {
	file_cilium_api_websocket_proto_rawDescOnce.Do(func() {
		file_cilium_api_websocket_proto_rawDescData = protoimpl.X.CompressGZIP(file_cilium_api_websocket_proto_rawDescData)
	})
	return file_cilium_api_websocket_proto_rawDescData
}

var file_cilium_api_websocket_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cilium_api_websocket_proto_goTypes = []interface{}{
	(*WebSocketClient)(nil),     // 0: cilium.WebSocketClient
	(*WebSocketServer)(nil),     // 1: cilium.WebSocketServer
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
}
var file_cilium_api_websocket_proto_depIdxs = []int32{
	2, // 0: cilium.WebSocketClient.handshake_timeout:type_name -> google.protobuf.Duration
	2, // 1: cilium.WebSocketClient.ping_interval:type_name -> google.protobuf.Duration
	2, // 2: cilium.WebSocketServer.handshake_timeout:type_name -> google.protobuf.Duration
	2, // 3: cilium.WebSocketServer.ping_interval:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cilium_api_websocket_proto_init() }
func file_cilium_api_websocket_proto_init() {
	if File_cilium_api_websocket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cilium_api_websocket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebSocketClient); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cilium_api_websocket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebSocketServer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cilium_api_websocket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cilium_api_websocket_proto_goTypes,
		DependencyIndexes: file_cilium_api_websocket_proto_depIdxs,
		MessageInfos:      file_cilium_api_websocket_proto_msgTypes,
	}.Build()
	File_cilium_api_websocket_proto = out.File
	file_cilium_api_websocket_proto_rawDesc = nil
	file_cilium_api_websocket_proto_goTypes = nil
	file_cilium_api_websocket_proto_depIdxs = nil
}
