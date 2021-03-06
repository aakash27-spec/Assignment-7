package calculator

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OperationRequest struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=firstNumber,proto3" json:"firstNumber,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=secondNumber,proto3" json:"secondNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationRequest) Reset()         { *m = OperationRequest{} }
func (m *OperationRequest) String() string { return proto.CompactTextString(m) }
func (*OperationRequest) ProtoMessage()    {}
func (*OperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

func (m *OperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationRequest.Unmarshal(m, b)
}
func (m *OperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationRequest.Marshal(b, m, deterministic)
}
func (m *OperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationRequest.Merge(m, src)
}
func (m *OperationRequest) XXX_Size() int {
	return xxx_messageInfo_OperationRequest.Size(m)
}
func (m *OperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperationRequest proto.InternalMessageInfo

func (m *OperationRequest) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *OperationRequest) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type ResultResponse struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResultResponse) Reset()         { *m = ResultResponse{} }
func (m *ResultResponse) String() string { return proto.CompactTextString(m) }
func (*ResultResponse) ProtoMessage()    {}
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{1}
}

func (m *ResultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultResponse.Unmarshal(m, b)
}
func (m *ResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultResponse.Marshal(b, m, deterministic)
}
func (m *ResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultResponse.Merge(m, src)
}
func (m *ResultResponse) XXX_Size() int {
	return xxx_messageInfo_ResultResponse.Size(m)
}
func (m *ResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResultResponse proto.InternalMessageInfo

func (m *ResultResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*OperationRequest)(nil), "calculator.OperationRequest")
	proto.RegisterType((*ResultResponse)(nil), "calculator.ResultResponse")
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor_c686ea360062a8cf) }

var fileDescriptor_c686ea360062a8cf = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xbd, 0x4a, 0xc0, 0x30,
	0x14, 0x46, 0x6d, 0xc1, 0x52, 0x6e, 0xa5, 0x94, 0x0c, 0x22, 0xc5, 0xa1, 0x64, 0xea, 0xd4, 0x41,
	0x9f, 0x40, 0x94, 0x82, 0xe0, 0x0f, 0xc4, 0xc5, 0x35, 0x4d, 0xaf, 0x10, 0x88, 0x4d, 0x4c, 0x6e,
	0x7c, 0x33, 0xdf, 0x4f, 0x0c, 0x45, 0x5b, 0xd7, 0x8e, 0x39, 0xf9, 0x38, 0x39, 0x10, 0x68, 0x94,
	0x34, 0x2a, 0x1a, 0x49, 0xd6, 0x0f, 0xce, 0x5b, 0xb2, 0x0c, 0xfe, 0x08, 0x7f, 0x85, 0xe6, 0xd9,
	0xa1, 0x97, 0xa4, 0xed, 0x22, 0xf0, 0x23, 0x62, 0x20, 0xd6, 0x41, 0xf5, 0xa6, 0x7d, 0xa0, 0xa7,
	0xf8, 0x3e, 0xa1, 0xbf, 0xc8, 0xba, 0xac, 0x3f, 0x15, 0x5b, 0xc4, 0x38, 0x9c, 0x05, 0x54, 0x76,
	0x99, 0xd7, 0x49, 0x9e, 0x26, 0x3b, 0xc6, 0x7b, 0xa8, 0x05, 0x86, 0x68, 0x48, 0x60, 0x70, 0x76,
	0x09, 0xc8, 0xce, 0xa1, 0xf0, 0x89, 0x24, 0x65, 0x2e, 0xd6, 0xd3, 0xd5, 0x57, 0x0e, 0x70, 0xfb,
	0x9b, 0xc4, 0x46, 0x28, 0x6f, 0xe6, 0x59, 0xff, 0x14, 0xb1, 0xcb, 0x61, 0x53, 0xff, 0x3f, 0xb4,
	0x6d, 0xb7, 0xb7, 0xfb, 0xc7, 0xf8, 0x09, 0xbb, 0x87, 0xea, 0x25, 0x4e, 0xe4, 0xa5, 0x3a, 0xac,
	0x7a, 0x80, 0xfa, 0x31, 0x1a, 0xd2, 0xce, 0x68, 0x25, 0x0f, 0xdb, 0x46, 0x28, 0xef, 0xf4, 0xa7,
	0x0e, 0x07, 0x3d, 0x53, 0x91, 0xbe, 0xf3, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xe0, 0x26,
	0x10, 0xe2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	Addition(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*ResultResponse, error)
	Subtraction(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*ResultResponse, error)
	Multiplication(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*ResultResponse, error)
	Division(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*ResultResponse, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Addition(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Addition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Subtraction(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Subtraction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Multiplication(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Multiplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Division(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Division", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	Addition(context.Context, *OperationRequest) (*ResultResponse, error)
	Subtraction(context.Context, *OperationRequest) (*ResultResponse, error)
	Multiplication(context.Context, *OperationRequest) (*ResultResponse, error)
	Division(context.Context, *OperationRequest) (*ResultResponse, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Addition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Addition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Addition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Addition(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Subtraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Subtraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Subtraction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Subtraction(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Multiplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Multiplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Multiplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Multiplication(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Division_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Division(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Division",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Division(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Addition",
			Handler:    _Calculator_Addition_Handler,
		},
		{
			MethodName: "Subtraction",
			Handler:    _Calculator_Subtraction_Handler,
		},
		{
			MethodName: "Multiplication",
			Handler:    _Calculator_Multiplication_Handler,
		},
		{
			MethodName: "Division",
			Handler:    _Calculator_Division_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}