// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package rpcpb is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
	SubmitTxRequest
	SubmitTxResponse
	NotifyRequest
	NotifyResponse
*/
package rpcpb

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TxStatusEnum int32

const (
	TxStatusEnum_NOTEXIST  TxStatusEnum = 0
	TxStatusEnum_REJECTED  TxStatusEnum = 1
	TxStatusEnum_ACCEPTED  TxStatusEnum = 2
	TxStatusEnum_CONFIRMED TxStatusEnum = 3
	TxStatusEnum_FAILED    TxStatusEnum = 4
)

var TxStatusEnum_name = map[int32]string{
	0: "NOTEXIST",
	1: "REJECTED",
	2: "ACCEPTED",
	3: "CONFIRMED",
	4: "FAILED",
}
var TxStatusEnum_value = map[string]int32{
	"NOTEXIST":  0,
	"REJECTED":  1,
	"ACCEPTED":  2,
	"CONFIRMED": 3,
	"FAILED":    4,
}

func (x TxStatusEnum) String() string {
	return proto.EnumName(TxStatusEnum_name, int32(x))
}
func (TxStatusEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NotifyMsgType int32

const (
	NotifyMsgType_TX        NotifyMsgType = 0
	NotifyMsgType_STATEMENT NotifyMsgType = 1
)

var NotifyMsgType_name = map[int32]string{
	0: "TX",
	1: "STATEMENT",
}
var NotifyMsgType_value = map[string]int32{
	"TX":        0,
	"STATEMENT": 1,
}

func (x NotifyMsgType) String() string {
	return proto.EnumName(NotifyMsgType_name, int32(x))
}
func (NotifyMsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type HelloRequest struct {
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HelloResponse struct {
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SubmitTxRequest struct {
	// transaction data in pb format
	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	// digital signature of the data signed by
	// the source account private key
	Signature string `protobuf:"bytes,2,opt,name=Signature" json:"Signature,omitempty"`
}

func (m *SubmitTxRequest) Reset()                    { *m = SubmitTxRequest{} }
func (m *SubmitTxRequest) String() string            { return proto.CompactTextString(m) }
func (*SubmitTxRequest) ProtoMessage()               {}
func (*SubmitTxRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SubmitTxRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SubmitTxRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type SubmitTxResponse struct {
	TxStatus TxStatusEnum `protobuf:"varint,1,opt,name=TxStatus,enum=rpcpb.TxStatusEnum" json:"TxStatus,omitempty"`
	// the transaction hash is only valid when the response
	// status is ACCEPTED or CONFIRMED
	TxHash string `protobuf:"bytes,2,opt,name=TxHash" json:"TxHash,omitempty"`
	// error message for REJECTED transaction
	ErrorMessage string `protobuf:"bytes,3,opt,name=ErrorMessage" json:"ErrorMessage,omitempty"`
}

func (m *SubmitTxResponse) Reset()                    { *m = SubmitTxResponse{} }
func (m *SubmitTxResponse) String() string            { return proto.CompactTextString(m) }
func (*SubmitTxResponse) ProtoMessage()               {}
func (*SubmitTxResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SubmitTxResponse) GetTxStatus() TxStatusEnum {
	if m != nil {
		return m.TxStatus
	}
	return TxStatusEnum_NOTEXIST
}

func (m *SubmitTxResponse) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *SubmitTxResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type NotifyRequest struct {
	// type of message
	MsgType NotifyMsgType `protobuf:"varint,1,opt,name=MsgType,enum=rpcpb.NotifyMsgType" json:"MsgType,omitempty"`
	// message payload in pb format
	Data []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	// digital signature of the data signed by
	// the private key of peer node
	Signature string `protobuf:"bytes,3,opt,name=Signature" json:"Signature,omitempty"`
}

func (m *NotifyRequest) Reset()                    { *m = NotifyRequest{} }
func (m *NotifyRequest) String() string            { return proto.CompactTextString(m) }
func (*NotifyRequest) ProtoMessage()               {}
func (*NotifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NotifyRequest) GetMsgType() NotifyMsgType {
	if m != nil {
		return m.MsgType
	}
	return NotifyMsgType_TX
}

func (m *NotifyRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *NotifyRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type NotifyResponse struct {
}

func (m *NotifyResponse) Reset()                    { *m = NotifyResponse{} }
func (m *NotifyResponse) String() string            { return proto.CompactTextString(m) }
func (*NotifyResponse) ProtoMessage()               {}
func (*NotifyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "rpcpb.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "rpcpb.HelloResponse")
	proto.RegisterType((*SubmitTxRequest)(nil), "rpcpb.SubmitTxRequest")
	proto.RegisterType((*SubmitTxResponse)(nil), "rpcpb.SubmitTxResponse")
	proto.RegisterType((*NotifyRequest)(nil), "rpcpb.NotifyRequest")
	proto.RegisterType((*NotifyResponse)(nil), "rpcpb.NotifyResponse")
	proto.RegisterEnum("rpcpb.TxStatusEnum", TxStatusEnum_name, TxStatusEnum_value)
	proto.RegisterEnum("rpcpb.NotifyMsgType", NotifyMsgType_name, NotifyMsgType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	SubmitTx(ctx context.Context, in *SubmitTxRequest, opts ...grpc.CallOption) (*SubmitTxResponse, error)
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Node/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) SubmitTx(ctx context.Context, in *SubmitTxRequest, opts ...grpc.CallOption) (*SubmitTxResponse, error) {
	out := new(SubmitTxResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Node/SubmitTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Node/Notify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	SubmitTx(context.Context, *SubmitTxRequest) (*SubmitTxResponse, error)
	Notify(context.Context, *NotifyRequest) (*NotifyResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Node/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_SubmitTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).SubmitTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Node/SubmitTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).SubmitTx(ctx, req.(*SubmitTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Node/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Node_Hello_Handler,
		},
		{
			MethodName: "SubmitTx",
			Handler:    _Node_SubmitTx_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Node_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0xaf, 0xd2, 0x40,
	0x14, 0xa5, 0x85, 0x57, 0xe1, 0xa6, 0xe5, 0x4d, 0xc6, 0xe7, 0x93, 0x10, 0x17, 0xa4, 0x0b, 0x43,
	0x58, 0xd4, 0x04, 0x4d, 0x5c, 0xb9, 0x20, 0xed, 0x10, 0x30, 0xb6, 0x98, 0x76, 0x16, 0x6c, 0x0b,
	0x8e, 0xd8, 0x04, 0x68, 0xe9, 0x4c, 0x0d, 0xac, 0xfc, 0x53, 0xfe, 0x40, 0xd3, 0x76, 0x86, 0x2f,
	0xf3, 0x76, 0x73, 0xce, 0x9c, 0xfb, 0x71, 0xee, 0xbd, 0x60, 0x71, 0x96, 0xff, 0x4e, 0xd6, 0xcc,
	0xc9, 0xf2, 0x54, 0xa4, 0xf8, 0x21, 0xcf, 0xd6, 0xd9, 0xca, 0xee, 0x82, 0x39, 0x63, 0xdb, 0x6d,
	0x1a, 0xb2, 0x43, 0xc1, 0xb8, 0xb0, 0x1f, 0xc1, 0x92, 0x98, 0x67, 0xe9, 0x9e, 0x33, 0xdb, 0x85,
	0xc7, 0xa8, 0x58, 0xed, 0x12, 0x41, 0x8f, 0x52, 0x83, 0x31, 0xb4, 0xbc, 0x58, 0xc4, 0x3d, 0x6d,
	0xa0, 0x0d, 0xcd, 0xb0, 0x7a, 0xe3, 0x77, 0xd0, 0x89, 0x92, 0xcd, 0x3e, 0x16, 0x45, 0xce, 0x7a,
	0xfa, 0x40, 0x1b, 0x76, 0xc2, 0x0b, 0x61, 0xff, 0x01, 0x74, 0x49, 0x52, 0x27, 0xc6, 0x1f, 0xa0,
	0x4d, 0x8f, 0x91, 0x88, 0x45, 0xc1, 0xab, 0x4c, 0xdd, 0xf1, 0x6b, 0xa7, 0xea, 0xc9, 0x51, 0x34,
	0xd9, 0x17, 0xbb, 0xf0, 0x2c, 0xc2, 0xcf, 0x60, 0xd0, 0xe3, 0x2c, 0xe6, 0xbf, 0x64, 0x7e, 0x89,
	0xb0, 0x0d, 0x26, 0xc9, 0xf3, 0x34, 0xf7, 0x19, 0xe7, 0xf1, 0x86, 0xf5, 0x9a, 0xd5, 0xef, 0x0d,
	0x67, 0x1f, 0xc0, 0x0a, 0x52, 0x91, 0xfc, 0x3c, 0x29, 0x0f, 0x0e, 0xbc, 0xf2, 0xf9, 0x86, 0x9e,
	0x32, 0x26, 0x8b, 0x3f, 0xc9, 0xe2, 0xb5, 0x4c, 0xfe, 0x85, 0x4a, 0x74, 0xf6, 0xac, 0xbf, 0xe4,
	0xb9, 0x79, 0xef, 0x19, 0x41, 0x57, 0x95, 0xac, 0x1d, 0x8f, 0x22, 0x30, 0xaf, 0xad, 0x61, 0x13,
	0xda, 0xc1, 0x82, 0x92, 0xe5, 0x3c, 0xa2, 0xa8, 0x51, 0xa2, 0x90, 0x7c, 0x25, 0x2e, 0x25, 0x1e,
	0xd2, 0x4a, 0x34, 0x71, 0x5d, 0xf2, 0xbd, 0x44, 0x3a, 0xb6, 0xa0, 0xe3, 0x2e, 0x82, 0xe9, 0x3c,
	0xf4, 0x89, 0x87, 0x9a, 0x18, 0xc0, 0x98, 0x4e, 0xe6, 0xdf, 0x88, 0x87, 0x5a, 0xa3, 0xf7, 0xca,
	0x99, 0xea, 0xd4, 0x00, 0x9d, 0x2e, 0x51, 0xa3, 0x8c, 0x89, 0xe8, 0x84, 0x12, 0x9f, 0x04, 0x14,
	0x69, 0xe3, 0xbf, 0x1a, 0xb4, 0x82, 0xf4, 0x07, 0xc3, 0x9f, 0xe0, 0xa1, 0xda, 0x30, 0x56, 0xe3,
	0xbe, 0xde, 0x7f, 0xff, 0xe9, 0x96, 0x94, 0x47, 0xd0, 0xc0, 0x5f, 0xa0, 0xad, 0x36, 0x88, 0x9f,
	0xa5, 0xe6, 0xee, 0x2e, 0xfa, 0x6f, 0xff, 0xe3, 0xcf, 0xe1, 0x9f, 0xc1, 0xa8, 0xbb, 0xc4, 0xb7,
	0x73, 0x56, 0xa1, 0x6f, 0xee, 0x58, 0x15, 0xb8, 0x32, 0xaa, 0x6b, 0xfd, 0xf8, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x3b, 0x67, 0xa1, 0x34, 0xbe, 0x02, 0x00, 0x00,
}