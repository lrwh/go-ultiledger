// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package rpcpb is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	TxStatus
	HelloRequest
	HelloResponse
	SubmitTxRequest
	SubmitTxResponse
	QueryTxRequest
	QueryTxResponse
	CreateTestAccountRequest
	CreateTestAccountResponse
	GetAccountRequest
	GetAccountResponse
	NotifyRequest
	NotifyResponse
	QueryRequest
	QueryResponse
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

// TxStatusCode enumerates the status of a transaction in the node.
type TxStatusCode int32

const (
	TxStatusCode_NOTEXIST  TxStatusCode = 0
	TxStatusCode_REJECTED  TxStatusCode = 1
	TxStatusCode_ACCEPTED  TxStatusCode = 2
	TxStatusCode_CONFIRMED TxStatusCode = 3
	TxStatusCode_FAILED    TxStatusCode = 4
)

var TxStatusCode_name = map[int32]string{
	0: "NOTEXIST",
	1: "REJECTED",
	2: "ACCEPTED",
	3: "CONFIRMED",
	4: "FAILED",
}
var TxStatusCode_value = map[string]int32{
	"NOTEXIST":  0,
	"REJECTED":  1,
	"ACCEPTED":  2,
	"CONFIRMED": 3,
	"FAILED":    4,
}

func (x TxStatusCode) String() string {
	return proto.EnumName(TxStatusCode_name, int32(x))
}
func (TxStatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

// QueryMsgType enumerates the type of query issued by the node.
type QueryMsgType int32

const (
	// Query Quorum information.
	QueryMsgType_QUORUM QueryMsgType = 0
	// Query TxSet information.
	QueryMsgType_TXSET QueryMsgType = 1
	// Query Ledger information.
	QueryMsgType_LEDGER QueryMsgType = 2
)

var QueryMsgType_name = map[int32]string{
	0: "QUORUM",
	1: "TXSET",
	2: "LEDGER",
}
var QueryMsgType_value = map[string]int32{
	"QUORUM": 0,
	"TXSET":  1,
	"LEDGER": 2,
}

func (x QueryMsgType) String() string {
	return proto.EnumName(QueryMsgType_name, int32(x))
}
func (QueryMsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// TxStatus contains the status code of a queried transaction and
// the error message associated with a failed transaction.
type TxStatus struct {
	StatusCode   TxStatusCode `protobuf:"varint,1,opt,name=StatusCode,enum=rpcpb.TxStatusCode" json:"StatusCode,omitempty"`
	ErrorMessage string       `protobuf:"bytes,2,opt,name=ErrorMessage" json:"ErrorMessage,omitempty"`
	// The full information of the tx in pb format.
	Data []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *TxStatus) Reset()                    { *m = TxStatus{} }
func (m *TxStatus) String() string            { return proto.CompactTextString(m) }
func (*TxStatus) ProtoMessage()               {}
func (*TxStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TxStatus) GetStatusCode() TxStatusCode {
	if m != nil {
		return m.StatusCode
	}
	return TxStatusCode_NOTEXIST
}

func (m *TxStatus) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *TxStatus) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type HelloRequest struct {
	NetworkID string `protobuf:"bytes,1,opt,name=NetworkID" json:"NetworkID,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

type HelloResponse struct {
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SubmitTxRequest struct {
	// Network id of the network.
	NetworkID string `protobuf:"bytes,1,opt,name=NetworkID" json:"NetworkID,omitempty"`
	// Transaction data in pb format.
	Data []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	// Transaction ULTKey.
	TxKey string `protobuf:"bytes,3,opt,name=TxKey" json:"TxKey,omitempty"`
	// Digital signature of the data signed by the source account private key.
	Signature string `protobuf:"bytes,4,opt,name=Signature" json:"Signature,omitempty"`
}

func (m *SubmitTxRequest) Reset()                    { *m = SubmitTxRequest{} }
func (m *SubmitTxRequest) String() string            { return proto.CompactTextString(m) }
func (*SubmitTxRequest) ProtoMessage()               {}
func (*SubmitTxRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SubmitTxRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *SubmitTxRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SubmitTxRequest) GetTxKey() string {
	if m != nil {
		return m.TxKey
	}
	return ""
}

func (m *SubmitTxRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type SubmitTxResponse struct {
}

func (m *SubmitTxResponse) Reset()                    { *m = SubmitTxResponse{} }
func (m *SubmitTxResponse) String() string            { return proto.CompactTextString(m) }
func (*SubmitTxResponse) ProtoMessage()               {}
func (*SubmitTxResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type QueryTxRequest struct {
	NetworkID string `protobuf:"bytes,1,opt,name=NetworkID" json:"NetworkID,omitempty"`
	TxKey     string `protobuf:"bytes,2,opt,name=TxKey" json:"TxKey,omitempty"`
}

func (m *QueryTxRequest) Reset()                    { *m = QueryTxRequest{} }
func (m *QueryTxRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryTxRequest) ProtoMessage()               {}
func (*QueryTxRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *QueryTxRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *QueryTxRequest) GetTxKey() string {
	if m != nil {
		return m.TxKey
	}
	return ""
}

type QueryTxResponse struct {
	TxStatus *TxStatus `protobuf:"bytes,1,opt,name=TxStatus" json:"TxStatus,omitempty"`
}

func (m *QueryTxResponse) Reset()                    { *m = QueryTxResponse{} }
func (m *QueryTxResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryTxResponse) ProtoMessage()               {}
func (*QueryTxResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QueryTxResponse) GetTxStatus() *TxStatus {
	if m != nil {
		return m.TxStatus
	}
	return nil
}

type CreateTestAccountRequest struct {
	NetworkID string `protobuf:"bytes,1,opt,name=NetworkID" json:"NetworkID,omitempty"`
	AccountID string `protobuf:"bytes,2,opt,name=AccountID" json:"AccountID,omitempty"`
}

func (m *CreateTestAccountRequest) Reset()                    { *m = CreateTestAccountRequest{} }
func (m *CreateTestAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTestAccountRequest) ProtoMessage()               {}
func (*CreateTestAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateTestAccountRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *CreateTestAccountRequest) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

type CreateTestAccountResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *CreateTestAccountResponse) Reset()                    { *m = CreateTestAccountResponse{} }
func (m *CreateTestAccountResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateTestAccountResponse) ProtoMessage()               {}
func (*CreateTestAccountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreateTestAccountResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetAccountRequest struct {
	NetworkID string `protobuf:"bytes,1,opt,name=NetworkID" json:"NetworkID,omitempty"`
	AccountID string `protobuf:"bytes,2,opt,name=AccountID" json:"AccountID,omitempty"`
}

func (m *GetAccountRequest) Reset()                    { *m = GetAccountRequest{} }
func (m *GetAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAccountRequest) ProtoMessage()               {}
func (*GetAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetAccountRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *GetAccountRequest) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

type GetAccountResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *GetAccountResponse) Reset()                    { *m = GetAccountResponse{} }
func (m *GetAccountResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAccountResponse) ProtoMessage()               {}
func (*GetAccountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetAccountResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type NotifyRequest struct {
	NetworkID string `protobuf:"bytes,1,opt,name=NetworkID" json:"NetworkID,omitempty"`
	// Type of the message.
	MsgType NotifyMsgType `protobuf:"varint,2,opt,name=MsgType,enum=rpcpb.NotifyMsgType" json:"MsgType,omitempty"`
	// Message payload in pb format.
	Data []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	// Digital signature of the data signed by the private key of peer node.
	Signature string `protobuf:"bytes,4,opt,name=Signature" json:"Signature,omitempty"`
}

func (m *NotifyRequest) Reset()                    { *m = NotifyRequest{} }
func (m *NotifyRequest) String() string            { return proto.CompactTextString(m) }
func (*NotifyRequest) ProtoMessage()               {}
func (*NotifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *NotifyRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

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
func (*NotifyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type QueryRequest struct {
	NetworkID string `protobuf:"bytes,1,opt,name=NetworkID" json:"NetworkID,omitempty"`
	// Type of message.
	MsgType QueryMsgType `protobuf:"varint,2,opt,name=MsgType,enum=rpcpb.QueryMsgType" json:"MsgType,omitempty"`
	// Message payload in pb format.
	Data []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	// Digital signature of the data signed by the private key of peer node.
	Signature string `protobuf:"bytes,4,opt,name=Signature" json:"Signature,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *QueryRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *QueryRequest) GetMsgType() QueryMsgType {
	if m != nil {
		return m.MsgType
	}
	return QueryMsgType_QUORUM
}

func (m *QueryRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *QueryRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type QueryResponse struct {
	// requested data encoded in pb format
	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *QueryResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TxStatus)(nil), "rpcpb.TxStatus")
	proto.RegisterType((*HelloRequest)(nil), "rpcpb.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "rpcpb.HelloResponse")
	proto.RegisterType((*SubmitTxRequest)(nil), "rpcpb.SubmitTxRequest")
	proto.RegisterType((*SubmitTxResponse)(nil), "rpcpb.SubmitTxResponse")
	proto.RegisterType((*QueryTxRequest)(nil), "rpcpb.QueryTxRequest")
	proto.RegisterType((*QueryTxResponse)(nil), "rpcpb.QueryTxResponse")
	proto.RegisterType((*CreateTestAccountRequest)(nil), "rpcpb.CreateTestAccountRequest")
	proto.RegisterType((*CreateTestAccountResponse)(nil), "rpcpb.CreateTestAccountResponse")
	proto.RegisterType((*GetAccountRequest)(nil), "rpcpb.GetAccountRequest")
	proto.RegisterType((*GetAccountResponse)(nil), "rpcpb.GetAccountResponse")
	proto.RegisterType((*NotifyRequest)(nil), "rpcpb.NotifyRequest")
	proto.RegisterType((*NotifyResponse)(nil), "rpcpb.NotifyResponse")
	proto.RegisterType((*QueryRequest)(nil), "rpcpb.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "rpcpb.QueryResponse")
	proto.RegisterEnum("rpcpb.TxStatusCode", TxStatusCode_name, TxStatusCode_value)
	proto.RegisterEnum("rpcpb.NotifyMsgType", NotifyMsgType_name, NotifyMsgType_value)
	proto.RegisterEnum("rpcpb.QueryMsgType", QueryMsgType_name, QueryMsgType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	// CreateTestAccount allows the client to create a test account for testing.
	CreateTestAccount(ctx context.Context, in *CreateTestAccountRequest, opts ...grpc.CallOption) (*CreateTestAccountResponse, error)
	// GetAccount allows the client to get the account information of an account.
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	// SubmitTx is the main way for the client to submit a transaction.
	SubmitTx(ctx context.Context, in *SubmitTxRequest, opts ...grpc.CallOption) (*SubmitTxResponse, error)
	// QueryTx is used for querying the status of previous submitted transactions.
	QueryTx(ctx context.Context, in *QueryTxRequest, opts ...grpc.CallOption) (*QueryTxResponse, error)
	// Hello is used for initial connection between new peers.
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// Notify is used for the node to notify transactions and consensus messages to other peers.
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
	// Query is used for the node to ask for missing information about transactions.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) CreateTestAccount(ctx context.Context, in *CreateTestAccountRequest, opts ...grpc.CallOption) (*CreateTestAccountResponse, error) {
	out := new(CreateTestAccountResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Node/CreateTestAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Node/GetAccount", in, out, c.cc, opts...)
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

func (c *nodeClient) QueryTx(ctx context.Context, in *QueryTxRequest, opts ...grpc.CallOption) (*QueryTxResponse, error) {
	out := new(QueryTxResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Node/QueryTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Node/Hello", in, out, c.cc, opts...)
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

func (c *nodeClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Node/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	// CreateTestAccount allows the client to create a test account for testing.
	CreateTestAccount(context.Context, *CreateTestAccountRequest) (*CreateTestAccountResponse, error)
	// GetAccount allows the client to get the account information of an account.
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	// SubmitTx is the main way for the client to submit a transaction.
	SubmitTx(context.Context, *SubmitTxRequest) (*SubmitTxResponse, error)
	// QueryTx is used for querying the status of previous submitted transactions.
	QueryTx(context.Context, *QueryTxRequest) (*QueryTxResponse, error)
	// Hello is used for initial connection between new peers.
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	// Notify is used for the node to notify transactions and consensus messages to other peers.
	Notify(context.Context, *NotifyRequest) (*NotifyResponse, error)
	// Query is used for the node to ask for missing information about transactions.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_CreateTestAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTestAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).CreateTestAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Node/CreateTestAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).CreateTestAccount(ctx, req.(*CreateTestAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Node/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetAccount(ctx, req.(*GetAccountRequest))
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

func _Node_QueryTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).QueryTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Node/QueryTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).QueryTx(ctx, req.(*QueryTxRequest))
	}
	return interceptor(ctx, in, info, handler)
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

func _Node_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Node/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTestAccount",
			Handler:    _Node_CreateTestAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Node_GetAccount_Handler,
		},
		{
			MethodName: "SubmitTx",
			Handler:    _Node_SubmitTx_Handler,
		},
		{
			MethodName: "QueryTx",
			Handler:    _Node_QueryTx_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _Node_Hello_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Node_Notify_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Node_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x8d, 0xdd, 0x24, 0xad, 0xe7, 0x73, 0x12, 0x77, 0xbf, 0xb6, 0xb8, 0x51, 0x25, 0x22, 0x23,
	0xa1, 0xa8, 0x40, 0x10, 0x29, 0x12, 0x12, 0x12, 0x48, 0x91, 0xbd, 0x2d, 0x81, 0xc6, 0xa1, 0xeb,
	0x2d, 0xca, 0xad, 0x1b, 0x96, 0x2a, 0xa2, 0xd4, 0xc1, 0x5e, 0xd3, 0xe6, 0x21, 0x10, 0x4f, 0xc5,
	0x7b, 0x21, 0xdb, 0xeb, 0x9f, 0xb8, 0x2d, 0xe4, 0xa2, 0x77, 0xde, 0xd9, 0x33, 0x67, 0x8e, 0x66,
	0xce, 0x8e, 0xa1, 0x11, 0x30, 0xff, 0xc7, 0x6c, 0xca, 0x7a, 0x73, 0xdf, 0xe3, 0x1e, 0xaa, 0xf9,
	0xf3, 0xe9, 0xfc, 0xcc, 0xb8, 0x82, 0x0d, 0x7a, 0xed, 0x70, 0x97, 0x87, 0x01, 0x3a, 0x00, 0x48,
	0xbe, 0x4c, 0xef, 0x33, 0xd3, 0xa5, 0x8e, 0xd4, 0x6d, 0xf6, 0xff, 0xef, 0xc5, 0xb8, 0x5e, 0x0a,
	0x8a, 0xae, 0x48, 0x01, 0x86, 0x0c, 0x50, 0xb1, 0xef, 0x7b, 0xfe, 0x88, 0x05, 0x81, 0x7b, 0xce,
	0x74, 0xb9, 0x23, 0x75, 0x15, 0xb2, 0x14, 0x43, 0x08, 0xaa, 0x96, 0xcb, 0x5d, 0x7d, 0xad, 0x23,
	0x75, 0x55, 0x12, 0x7f, 0x1b, 0x4f, 0x41, 0x7d, 0xc7, 0x2e, 0x2e, 0x3c, 0xc2, 0xbe, 0x87, 0x2c,
	0xe0, 0x68, 0x0f, 0x14, 0x9b, 0xf1, 0x2b, 0xcf, 0xff, 0x3a, 0xb4, 0xe2, 0xda, 0x0a, 0xc9, 0x03,
	0x46, 0x0b, 0x1a, 0x02, 0x1d, 0xcc, 0xbd, 0xcb, 0x80, 0x19, 0x57, 0xd0, 0x72, 0xc2, 0xb3, 0x6f,
	0x33, 0x4e, 0xaf, 0x57, 0x62, 0xc8, 0x34, 0xc8, 0xb9, 0x06, 0xb4, 0x05, 0x35, 0x7a, 0xfd, 0x81,
	0x2d, 0x62, 0x61, 0x0a, 0x49, 0x0e, 0x11, 0x8f, 0x33, 0x3b, 0xbf, 0x74, 0x79, 0xe8, 0x33, 0xbd,
	0x9a, 0xf0, 0x64, 0x01, 0x03, 0x81, 0x96, 0x17, 0x16, 0x62, 0x2c, 0x68, 0x9e, 0x84, 0xcc, 0x5f,
	0xac, 0xaa, 0x25, 0xab, 0x2b, 0x17, 0xea, 0x1a, 0x6f, 0xa1, 0x95, 0xb1, 0x24, 0xc4, 0xe8, 0x49,
	0x3e, 0x9d, 0x98, 0xe5, 0xbf, 0x7e, 0xab, 0x34, 0x0f, 0x92, 0x01, 0x8c, 0x4f, 0xa0, 0x9b, 0x3e,
	0x73, 0x39, 0xa3, 0x2c, 0xe0, 0x83, 0xe9, 0xd4, 0x0b, 0x2f, 0xf9, 0x6a, 0x7a, 0xf6, 0x40, 0x11,
	0xf8, 0xa1, 0x25, 0x34, 0xe5, 0x01, 0xe3, 0x39, 0xec, 0xde, 0xc2, 0x2b, 0x14, 0xa6, 0x6d, 0x95,
	0x0a, 0xa3, 0x1d, 0xc3, 0xe6, 0x11, 0xbb, 0x4f, 0x05, 0x5d, 0x40, 0x45, 0xc2, 0xbf, 0x94, 0xfe,
	0x25, 0x41, 0xc3, 0xf6, 0xf8, 0xec, 0xcb, 0x62, 0xb5, 0xba, 0x3d, 0x58, 0x1f, 0x05, 0xe7, 0x74,
	0x31, 0x4f, 0x8c, 0xdb, 0xec, 0x6f, 0x89, 0xfe, 0x26, 0x24, 0xe2, 0x8e, 0xa4, 0xa0, 0xdb, 0x9c,
	0xfc, 0x0f, 0xbf, 0x68, 0xd0, 0x4c, 0x05, 0x09, 0xb7, 0xfc, 0x94, 0x40, 0x8d, 0x07, 0xbd, 0x9a,
	0xc4, 0x67, 0x65, 0x89, 0xe9, 0x93, 0x8c, 0x39, 0xee, 0x41, 0xe1, 0x23, 0x68, 0x08, 0x39, 0x77,
	0x37, 0x76, 0xdf, 0x01, 0xb5, 0xb8, 0x02, 0x90, 0x0a, 0x1b, 0xf6, 0x98, 0xe2, 0xc9, 0xd0, 0xa1,
	0x5a, 0x25, 0x3a, 0x11, 0xfc, 0x1e, 0x9b, 0x14, 0x5b, 0x9a, 0x14, 0x9d, 0x06, 0xa6, 0x89, 0x3f,
	0x46, 0x27, 0x19, 0x35, 0x40, 0x31, 0xc7, 0xf6, 0xe1, 0x90, 0x8c, 0xb0, 0xa5, 0xad, 0x21, 0x80,
	0xfa, 0xe1, 0x60, 0x78, 0x8c, 0x2d, 0xad, 0xba, 0xff, 0x38, 0x1d, 0x56, 0x2a, 0xbe, 0x0e, 0x32,
	0x9d, 0x68, 0x95, 0x28, 0xc7, 0xa1, 0x03, 0x8a, 0x47, 0xd8, 0xa6, 0x9a, 0xb4, 0xff, 0x42, 0x34,
	0x2c, 0x85, 0x01, 0xd4, 0x4f, 0x4e, 0xc7, 0xe4, 0x74, 0xa4, 0x55, 0x90, 0x02, 0x35, 0x3a, 0x71,
	0x30, 0xd5, 0xa4, 0x28, 0x7c, 0x8c, 0xad, 0x23, 0x4c, 0x34, 0xb9, 0xff, 0x7b, 0x0d, 0xaa, 0x76,
	0x24, 0x74, 0x02, 0x9b, 0x37, 0xdc, 0x8b, 0x1e, 0x8a, 0x16, 0xde, 0xf5, 0x5e, 0xda, 0x9d, 0xbb,
	0x01, 0x62, 0x8a, 0x15, 0x64, 0x02, 0xe4, 0xae, 0x44, 0xba, 0xc8, 0xb8, 0xe1, 0xfc, 0xf6, 0xee,
	0x2d, 0x37, 0x19, 0xc9, 0x1b, 0xd8, 0x48, 0xd7, 0x09, 0xda, 0x11, 0xc0, 0xd2, 0x62, 0x6b, 0x3f,
	0xb8, 0x11, 0xcf, 0xd2, 0x5f, 0xc3, 0xba, 0xd8, 0x19, 0x68, 0xbb, 0x68, 0x8b, 0x3c, 0x79, 0xa7,
	0x1c, 0xce, 0x72, 0x5f, 0x42, 0x2d, 0xde, 0xa9, 0x28, 0x35, 0x54, 0x71, 0x1f, 0xb7, 0xb7, 0x96,
	0x83, 0x59, 0xd6, 0x2b, 0xa8, 0x27, 0x33, 0x43, 0xcb, 0x4f, 0x25, 0xcd, 0xdb, 0x2e, 0x45, 0x8b,
	0xe5, 0x62, 0x0d, 0x68, 0xc9, 0xbf, 0xe5, 0x72, 0x4b, 0x4e, 0x34, 0x2a, 0x67, 0xf5, 0xf8, 0x6f,
	0x75, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0xad, 0x59, 0x7c, 0x69, 0xbe, 0x06, 0x00, 0x00,
}
