// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.18.1
// source: logstream.proto

package heads

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_logstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_logstream_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

var File_logstream_proto protoreflect.FileDescriptor

var file_logstream_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x68, 0x65, 0x61, 0x64, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x32,
	0x36, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x29, 0x0a, 0x0b,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x0c, 0x2e, 0x68, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x30, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x63, 0x6b, 0x74, 0x6f, 0x70, 0x75, 0x73, 0x2f,
	0x74, 0x68, 0x65, 0x68, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logstream_proto_rawDescOnce sync.Once
	file_logstream_proto_rawDescData = file_logstream_proto_rawDesc
)

func file_logstream_proto_rawDescGZIP() []byte {
	file_logstream_proto_rawDescOnce.Do(func() {
		file_logstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_logstream_proto_rawDescData)
	})
	return file_logstream_proto_rawDescData
}

var file_logstream_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_logstream_proto_goTypes = []interface{}{
	(*Log)(nil),   // 0: heads.Log
	(*Empty)(nil), // 1: heads.Empty
}
var file_logstream_proto_depIdxs = []int32{
	1, // 0: heads.logstream.stream_logs:input_type -> heads.Empty
	0, // 1: heads.logstream.stream_logs:output_type -> heads.Log
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logstream_proto_init() }
func file_logstream_proto_init() {
	if File_logstream_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_logstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_logstream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logstream_proto_goTypes,
		DependencyIndexes: file_logstream_proto_depIdxs,
		MessageInfos:      file_logstream_proto_msgTypes,
	}.Build()
	File_logstream_proto = out.File
	file_logstream_proto_rawDesc = nil
	file_logstream_proto_goTypes = nil
	file_logstream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogstreamClient is the client API for Logstream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogstreamClient interface {
	StreamLogs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Logstream_StreamLogsClient, error)
}

type logstreamClient struct {
	cc grpc.ClientConnInterface
}

func NewLogstreamClient(cc grpc.ClientConnInterface) LogstreamClient {
	return &logstreamClient{cc}
}

func (c *logstreamClient) StreamLogs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Logstream_StreamLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Logstream_serviceDesc.Streams[0], "/heads.logstream/stream_logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &logstreamStreamLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Logstream_StreamLogsClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type logstreamStreamLogsClient struct {
	grpc.ClientStream
}

func (x *logstreamStreamLogsClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogstreamServer is the server API for Logstream service.
type LogstreamServer interface {
	StreamLogs(*Empty, Logstream_StreamLogsServer) error
}

// UnimplementedLogstreamServer can be embedded to have forward compatible implementations.
type UnimplementedLogstreamServer struct {
}

func (*UnimplementedLogstreamServer) StreamLogs(*Empty, Logstream_StreamLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLogs not implemented")
}

func RegisterLogstreamServer(s *grpc.Server, srv LogstreamServer) {
	s.RegisterService(&_Logstream_serviceDesc, srv)
}

func _Logstream_StreamLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogstreamServer).StreamLogs(m, &logstreamStreamLogsServer{stream})
}

type Logstream_StreamLogsServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type logstreamStreamLogsServer struct {
	grpc.ServerStream
}

func (x *logstreamStreamLogsServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

var _Logstream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heads.logstream",
	HandlerType: (*LogstreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "stream_logs",
			Handler:       _Logstream_StreamLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "logstream.proto",
}
