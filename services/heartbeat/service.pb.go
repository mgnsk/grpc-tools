// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/heartbeat/service.proto

package heartbeat

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PingRequest struct {
	Message              []byte   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a8dcc4dd6d17f26, []int{0}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return m.Size()
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (*PingRequest) XXX_MessageName() string {
	return "heartbeat.PingRequest"
}

type PingReply struct {
	Message              []byte   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a8dcc4dd6d17f26, []int{1}
}
func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(m, src)
}
func (m *PingReply) XXX_Size() int {
	return m.Size()
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

func (m *PingReply) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (*PingReply) XXX_MessageName() string {
	return "heartbeat.PingReply"
}

type StreamRequest struct {
	Id                   []byte     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateDate           *time.Time `protobuf:"bytes,2,opt,name=create_date,json=createDate,proto3,stdtime" json:"create_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a8dcc4dd6d17f26, []int{2}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return m.Size()
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *StreamRequest) GetCreateDate() *time.Time {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (*StreamRequest) XXX_MessageName() string {
	return "heartbeat.StreamRequest"
}

type StreamPacket struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamPacket) Reset()         { *m = StreamPacket{} }
func (m *StreamPacket) String() string { return proto.CompactTextString(m) }
func (*StreamPacket) ProtoMessage()    {}
func (*StreamPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a8dcc4dd6d17f26, []int{3}
}
func (m *StreamPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamPacket.Merge(m, src)
}
func (m *StreamPacket) XXX_Size() int {
	return m.Size()
}
func (m *StreamPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamPacket.DiscardUnknown(m)
}

var xxx_messageInfo_StreamPacket proto.InternalMessageInfo

func (m *StreamPacket) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *StreamPacket) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (*StreamPacket) XXX_MessageName() string {
	return "heartbeat.StreamPacket"
}
func init() {
	proto.RegisterType((*PingRequest)(nil), "heartbeat.PingRequest")
	golang_proto.RegisterType((*PingRequest)(nil), "heartbeat.PingRequest")
	proto.RegisterType((*PingReply)(nil), "heartbeat.PingReply")
	golang_proto.RegisterType((*PingReply)(nil), "heartbeat.PingReply")
	proto.RegisterType((*StreamRequest)(nil), "heartbeat.StreamRequest")
	golang_proto.RegisterType((*StreamRequest)(nil), "heartbeat.StreamRequest")
	proto.RegisterType((*StreamPacket)(nil), "heartbeat.StreamPacket")
	golang_proto.RegisterType((*StreamPacket)(nil), "heartbeat.StreamPacket")
}

func init() { proto.RegisterFile("protobuf/heartbeat/service.proto", fileDescriptor_1a8dcc4dd6d17f26) }
func init() {
	golang_proto.RegisterFile("protobuf/heartbeat/service.proto", fileDescriptor_1a8dcc4dd6d17f26)
}

var fileDescriptor_1a8dcc4dd6d17f26 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0x86, 0xb4, 0x51, 0x37, 0x29, 0xaa, 0x56, 0x08, 0x22, 0xab, 0x72, 0x2d, 0x4b, 0x88,
	0xa8, 0x22, 0xde, 0x12, 0x6e, 0x85, 0x4b, 0x7f, 0x90, 0x7a, 0x41, 0x0a, 0x2e, 0x27, 0x2e, 0x68,
	0xec, 0x0c, 0x1b, 0x2b, 0xb6, 0x77, 0xf1, 0x6e, 0x80, 0x1e, 0xb8, 0xf0, 0x04, 0x08, 0xde, 0x82,
	0xa7, 0xe0, 0xd8, 0x23, 0x12, 0x0f, 0x00, 0x4a, 0x79, 0x10, 0x94, 0xdd, 0x38, 0x0d, 0x0a, 0xea,
	0xc9, 0x33, 0xfb, 0x7d, 0xfb, 0xed, 0xf8, 0xfb, 0x86, 0x06, 0xaa, 0x92, 0x46, 0x26, 0xd3, 0x37,
	0x7c, 0x8c, 0x50, 0x99, 0x04, 0xc1, 0x70, 0x8d, 0xd5, 0xbb, 0x2c, 0xc5, 0xc8, 0x42, 0x6c, 0x6b,
	0x09, 0x78, 0xbb, 0x42, 0x4a, 0x91, 0x23, 0x07, 0x95, 0x71, 0x28, 0x4b, 0x69, 0xc0, 0x64, 0xb2,
	0xd4, 0x8e, 0xe8, 0xed, 0x2d, 0xd0, 0xa5, 0xa2, 0xc9, 0x0a, 0xd4, 0x06, 0x0a, 0xb5, 0x20, 0x3c,
	0xb4, 0x9f, 0xb4, 0x2f, 0xb0, 0xec, 0xeb, 0xf7, 0x20, 0x04, 0x56, 0x5c, 0x2a, 0x2b, 0xf1, 0x1f,
	0xb9, 0xbe, 0xc8, 0xcc, 0x78, 0x9a, 0x44, 0xa9, 0x2c, 0xb8, 0x90, 0x42, 0x5e, 0xeb, 0xce, 0x3b,
	0xdb, 0xd8, 0xca, 0xd1, 0xc3, 0x07, 0xb4, 0x3d, 0xcc, 0x4a, 0x11, 0xe3, 0xdb, 0x29, 0x6a, 0xc3,
	0xba, 0xb4, 0xf5, 0x1c, 0xb5, 0x06, 0x81, 0x5d, 0x12, 0x90, 0x5e, 0x27, 0xae, 0xdb, 0xf0, 0x3e,
	0xdd, 0x72, 0x44, 0x95, 0x5f, 0xdc, 0x40, 0x4b, 0xe8, 0xf6, 0xb9, 0xa9, 0x10, 0x8a, 0x5a, 0xf1,
	0x36, 0x6d, 0x64, 0xa3, 0x05, 0xab, 0x91, 0x8d, 0xd8, 0x11, 0x6d, 0xa7, 0x15, 0x82, 0xc1, 0xd7,
	0x23, 0x30, 0xd8, 0x6d, 0x04, 0xa4, 0xd7, 0x1e, 0x78, 0x91, 0x33, 0x21, 0xaa, 0x87, 0x8d, 0x5e,
	0xd6, 0x26, 0x1c, 0x37, 0x3f, 0xff, 0xda, 0x23, 0x31, 0x75, 0x97, 0x4e, 0xc1, 0x60, 0x38, 0xa0,
	0x1d, 0xf7, 0xc6, 0x10, 0xd2, 0x09, 0xae, 0x3f, 0xc1, 0x68, 0xf3, 0x14, 0x0c, 0x58, 0xed, 0x4e,
	0x6c, 0xeb, 0xc1, 0x37, 0x42, 0x77, 0xce, 0xea, 0x44, 0xce, 0x5d, 0x52, 0xec, 0x84, 0x36, 0xe7,
	0xff, 0xc4, 0xee, 0x46, 0xcb, 0xb0, 0xa2, 0x15, 0x37, 0xbc, 0x3b, 0x6b, 0xe7, 0x2a, 0xbf, 0x08,
	0xb7, 0x3f, 0xfd, 0xfc, 0xf3, 0xb5, 0xd1, 0x62, 0x1b, 0x5c, 0xcd, 0x2f, 0xbf, 0xa0, 0x9b, 0x6e,
	0x1a, 0xd6, 0x5d, 0xa1, 0xff, 0x63, 0x82, 0x77, 0x6f, 0x0d, 0x71, 0xa3, 0x87, 0xcc, 0x6a, 0x75,
	0xc2, 0x16, 0xd7, 0xf6, 0xf8, 0x90, 0xec, 0x1f, 0x90, 0xe3, 0x8f, 0x5f, 0x8e, 0xce, 0xd8, 0xc6,
	0xe0, 0xd6, 0xa3, 0xe8, 0x60, 0x9f, 0x34, 0xaa, 0xa7, 0x74, 0xe7, 0xd9, 0x07, 0x28, 0x54, 0x8e,
	0x81, 0x88, 0x87, 0x27, 0x01, 0xa8, 0x8c, 0xf5, 0xc6, 0xc6, 0x28, 0x7d, 0xc8, 0xf9, 0x4a, 0xdc,
	0x05, 0x88, 0x72, 0xaa, 0x27, 0x72, 0x32, 0xe1, 0xa2, 0x52, 0x69, 0xdf, 0x48, 0x99, 0xeb, 0xcb,
	0x99, 0x4f, 0x7e, 0xcc, 0x7c, 0xf2, 0x7b, 0xe6, 0x93, 0xef, 0x57, 0x3e, 0xb9, 0xbc, 0xf2, 0xc9,
	0xab, 0xdd, 0xc5, 0xae, 0xea, 0xeb, 0xf5, 0x7d, 0xb2, 0xac, 0x92, 0x4d, 0x9b, 0xc2, 0xe3, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x5e, 0xa2, 0x85, 0xe5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeartbeatServiceClient is the client API for HeartbeatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartbeatServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (HeartbeatService_StreamClient, error)
}

type heartbeatServiceClient struct {
	cc *grpc.ClientConn
}

func NewHeartbeatServiceClient(cc *grpc.ClientConn) HeartbeatServiceClient {
	return &heartbeatServiceClient{cc}
}

func (c *heartbeatServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/heartbeat.HeartbeatService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heartbeatServiceClient) Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (HeartbeatService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HeartbeatService_serviceDesc.Streams[0], "/heartbeat.HeartbeatService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &heartbeatServiceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HeartbeatService_StreamClient interface {
	Recv() (*StreamPacket, error)
	grpc.ClientStream
}

type heartbeatServiceStreamClient struct {
	grpc.ClientStream
}

func (x *heartbeatServiceStreamClient) Recv() (*StreamPacket, error) {
	m := new(StreamPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HeartbeatServiceServer is the server API for HeartbeatService service.
type HeartbeatServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Stream(*StreamRequest, HeartbeatService_StreamServer) error
}

func RegisterHeartbeatServiceServer(s *grpc.Server, srv HeartbeatServiceServer) {
	s.RegisterService(&_HeartbeatService_serviceDesc, srv)
}

func _HeartbeatService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heartbeat.HeartbeatService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeartbeatService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HeartbeatServiceServer).Stream(m, &heartbeatServiceStreamServer{stream})
}

type HeartbeatService_StreamServer interface {
	Send(*StreamPacket) error
	grpc.ServerStream
}

type heartbeatServiceStreamServer struct {
	grpc.ServerStream
}

func (x *heartbeatServiceStreamServer) Send(m *StreamPacket) error {
	return x.ServerStream.SendMsg(m)
}

var _HeartbeatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heartbeat.HeartbeatService",
	HandlerType: (*HeartbeatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _HeartbeatService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _HeartbeatService_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protobuf/heartbeat/service.proto",
}

func (m *PingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PingReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StreamRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.CreateDate != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintService(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreateDate)))
		n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreateDate, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StreamPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamPacket) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PingReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StreamRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.CreateDate != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreateDate)
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StreamPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message[:0], dAtA[iNdEx:postIndex]...)
			if m.Message == nil {
				m.Message = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PingReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PingReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message[:0], dAtA[iNdEx:postIndex]...)
			if m.Message == nil {
				m.Message = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreateDate == nil {
				m.CreateDate = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreateDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthService
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService   = fmt.Errorf("proto: integer overflow")
)
