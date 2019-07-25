// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: idl/echo/echov1/echo_api.proto

package echov1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EchoRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d59fd4a7bca5ed, []int{0}
}
func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return m.Size()
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (*EchoRequest) XXX_MessageName() string {
	return "app.echo.v1.EchoRequest"
}

type EchoResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d59fd4a7bca5ed, []int{1}
}
func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return m.Size()
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (*EchoResponse) XXX_MessageName() string {
	return "app.echo.v1.EchoResponse"
}
func init() {
	proto.RegisterType((*EchoRequest)(nil), "app.echo.v1.EchoRequest")
	golang_proto.RegisterType((*EchoRequest)(nil), "app.echo.v1.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "app.echo.v1.EchoResponse")
	golang_proto.RegisterType((*EchoResponse)(nil), "app.echo.v1.EchoResponse")
}

func init() { proto.RegisterFile("idl/echo/echov1/echo_api.proto", fileDescriptor_94d59fd4a7bca5ed) }
func init() {
	golang_proto.RegisterFile("idl/echo/echov1/echo_api.proto", fileDescriptor_94d59fd4a7bca5ed)
}

var fileDescriptor_94d59fd4a7bca5ed = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbf, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0xf3, 0x80, 0xf6, 0x44, 0x0a, 0x0c, 0x99, 0x8e, 0x13, 0x7a, 0x8a, 0x6e, 0xe1, 0x84,
	0xb8, 0xe4, 0xae, 0x2c, 0xa8, 0x5b, 0xf9, 0x31, 0x94, 0xa9, 0x3a, 0x89, 0x85, 0x05, 0xb9, 0xe9,
	0xab, 0x13, 0x5d, 0xce, 0xcf, 0xc4, 0x4e, 0x58, 0x11, 0x7f, 0x01, 0x3f, 0xfe, 0x09, 0xfe, 0x04,
	0xc6, 0x8e, 0x1d, 0x4f, 0x62, 0xe9, 0xd8, 0x38, 0x0c, 0xb0, 0x75, 0x64, 0x44, 0xe7, 0x08, 0xa9,
	0x43, 0x59, 0xec, 0x8f, 0xdf, 0xfb, 0xda, 0x5f, 0xf9, 0x6b, 0x87, 0x58, 0x1c, 0x97, 0x29, 0x65,
	0x39, 0xfb, 0xa1, 0x99, 0xfb, 0xe9, 0xad, 0xd0, 0x45, 0xa2, 0x2b, 0xb6, 0x1c, 0xed, 0x08, 0xad,
	0x93, 0x4d, 0x2d, 0x69, 0xe6, 0xa3, 0xc7, 0xbe, 0x96, 0x4d, 0x25, 0xa9, 0xa9, 0x79, 0x2f, 0xa4,
	0xa4, 0x2a, 0x65, 0x6d, 0x0b, 0x56, 0x26, 0x15, 0x4a, 0xb1, 0x15, 0x9e, 0xfb, 0xad, 0xa3, 0x07,
	0x92, 0x59, 0x96, 0x94, 0x0a, 0x5d, 0x5c, 0xd3, 0x0d, 0x25, 0x4b, 0xee, 0x79, 0xfc, 0x30, 0xdc,
	0x79, 0x99, 0xe5, 0xbc, 0xa0, 0x77, 0x35, 0x19, 0x1b, 0x0d, 0xc3, 0xc1, 0x8a, 0x8c, 0x11, 0x92,
	0x86, 0x10, 0xc3, 0xe4, 0xf6, 0xe2, 0xdf, 0x72, 0x3c, 0x09, 0xef, 0xf4, 0x42, 0xa3, 0x59, 0x19,
	0xfa, 0xbf, 0x72, 0xf7, 0x75, 0x38, 0xd8, 0x28, 0xf7, 0x0f, 0x0f, 0xa2, 0x57, 0xe1, 0xad, 0x0d,
	0x46, 0xc3, 0xe4, 0xca, 0x5d, 0x92, 0x2b, 0x86, 0xa3, 0xfb, 0xd7, 0x74, 0x7a, 0x87, 0xf1, 0xdd,
	0x8f, 0x3f, 0x7e, 0x7e, 0xbd, 0x31, 0x88, 0xb6, 0x7c, 0x30, 0xcf, 0x7e, 0xc3, 0x97, 0xfd, 0xcf,
	0x10, 0x6d, 0xed, 0xde, 0x9c, 0x27, 0xb3, 0xd1, 0xbd, 0x92, 0x33, 0x51, 0xe6, 0x6c, 0xec, 0xde,
	0xd3, 0xd9, 0x6c, 0xf6, 0x08, 0xa0, 0xaa, 0xc3, 0x83, 0x17, 0xd4, 0x50, 0xc9, 0x7a, 0x45, 0xca,
	0xc6, 0xa4, 0x9a, 0xa2, 0x62, 0xe5, 0x59, 0xa8, 0xe3, 0xd8, 0x32, 0x97, 0x26, 0x3e, 0xe1, 0x2a,
	0x3e, 0xa9, 0xcb, 0x72, 0x6a, 0xac, 0xc8, 0x96, 0xb1, 0x5c, 0x1c, 0x3e, 0x8f, 0x0d, 0x55, 0x4d,
	0x91, 0x51, 0x2c, 0x49, 0x51, 0xe5, 0x93, 0x4a, 0xa2, 0x49, 0x6e, 0xad, 0x36, 0x7b, 0x69, 0x2a,
	0x0b, 0x9b, 0xd7, 0x47, 0x49, 0xc6, 0xab, 0x74, 0x25, 0xa4, 0xaa, 0xcd, 0x92, 0x97, 0xcb, 0x54,
	0x56, 0x3a, 0x9b, 0xfa, 0x13, 0xd7, 0x2d, 0x06, 0xe7, 0x2d, 0xc2, 0x45, 0x8b, 0x70, 0xd9, 0x22,
	0xfc, 0x69, 0x11, 0x3e, 0x38, 0x0c, 0xbe, 0x39, 0x84, 0xef, 0x0e, 0xe1, 0xd4, 0x21, 0x9c, 0x39,
	0x84, 0xb5, 0x43, 0xb8, 0x70, 0x08, 0xbf, 0x1c, 0x06, 0x97, 0x0e, 0xe1, 0x53, 0x87, 0xc1, 0x69,
	0x87, 0x70, 0xd6, 0x21, 0xac, 0x3b, 0x0c, 0xce, 0x3b, 0x0c, 0xde, 0x6c, 0xf7, 0x3f, 0xe1, 0x68,
	0xdb, 0x3f, 0xce, 0x93, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x25, 0x2c, 0xb3, 0xa5, 0x23, 0x02,
	0x00, 0x00,
}

func (this *EchoRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*EchoRequest)
	if !ok {
		that2, ok := that.(EchoRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *EchoRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *EchoRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *EchoRequest but is not nil && this == nil")
	}
	if this.Message != that1.Message {
		return fmt.Errorf("Message this(%v) Not Equal that(%v)", this.Message, that1.Message)
	}
	return nil
}
func (this *EchoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EchoRequest)
	if !ok {
		that2, ok := that.(EchoRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *EchoResponse) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*EchoResponse)
	if !ok {
		that2, ok := that.(EchoResponse)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *EchoResponse")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *EchoResponse but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *EchoResponse but is not nil && this == nil")
	}
	if this.Message != that1.Message {
		return fmt.Errorf("Message this(%v) Not Equal that(%v)", this.Message, that1.Message)
	}
	return nil
}
func (this *EchoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EchoResponse)
	if !ok {
		that2, ok := that.(EchoResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *EchoRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&echov1.EchoRequest{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *EchoResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&echov1.EchoResponse{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringEchoApi(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoAPIClient is the client API for EchoAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoAPIClient interface {
	// Echo same request.
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoAPIClient struct {
	cc *grpc.ClientConn
}

func NewEchoAPIClient(cc *grpc.ClientConn) EchoAPIClient {
	return &echoAPIClient{cc}
}

func (c *echoAPIClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/app.echo.v1.EchoAPI/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoAPIServer is the server API for EchoAPI service.
type EchoAPIServer interface {
	// Echo same request.
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

func RegisterEchoAPIServer(s *grpc.Server, srv EchoAPIServer) {
	s.RegisterService(&_EchoAPI_serviceDesc, srv)
}

func _EchoAPI_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoAPIServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.echo.v1.EchoAPI/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoAPIServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.echo.v1.EchoAPI",
	HandlerType: (*EchoAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoAPI_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/echo/echov1/echo_api.proto",
}

func (m *EchoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEchoApi(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func (m *EchoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEchoApi(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func encodeVarintEchoApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedEchoRequest(r randyEchoApi, easy bool) *EchoRequest {
	this := &EchoRequest{}
	this.Message = string(randStringEchoApi(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedEchoResponse(r randyEchoApi, easy bool) *EchoResponse {
	this := &EchoResponse{}
	this.Message = string(randStringEchoApi(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyEchoApi interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEchoApi(r randyEchoApi) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringEchoApi(r randyEchoApi) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneEchoApi(r)
	}
	return string(tmps)
}
func randUnrecognizedEchoApi(r randyEchoApi, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldEchoApi(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldEchoApi(dAtA []byte, r randyEchoApi, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateEchoApi(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateEchoApi(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateEchoApi(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateEchoApi(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateEchoApi(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateEchoApi(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateEchoApi(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *EchoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovEchoApi(uint64(l))
	}
	return n
}

func (m *EchoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovEchoApi(uint64(l))
	}
	return n
}

func sovEchoApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEchoApi(x uint64) (n int) {
	return sovEchoApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EchoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEchoApi
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
			return fmt.Errorf("proto: EchoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EchoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEchoApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEchoApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEchoApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEchoApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEchoApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEchoApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EchoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEchoApi
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
			return fmt.Errorf("proto: EchoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EchoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEchoApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEchoApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEchoApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEchoApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEchoApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEchoApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEchoApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEchoApi
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
					return 0, ErrIntOverflowEchoApi
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
					return 0, ErrIntOverflowEchoApi
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
				return 0, ErrInvalidLengthEchoApi
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthEchoApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEchoApi
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
				next, err := skipEchoApi(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthEchoApi
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
	ErrInvalidLengthEchoApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEchoApi   = fmt.Errorf("proto: integer overflow")
)
