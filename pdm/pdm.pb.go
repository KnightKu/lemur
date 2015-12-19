// Code generated by protoc-gen-go.
// source: pdm/pdm.proto
// DO NOT EDIT!

/*
Package pdm is a generated protocol buffer package.

It is generated from these files:
	pdm/pdm.proto

It has these top-level messages:
	Endpoint
	Handle
	ActionItem
	ActionStatus
	Empty
*/
package pdm

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

type Endpoint struct {
	Archive uint32 `protobuf:"varint,1,opt,name=archive" json:"archive,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Handle struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Handle) Reset()                    { *m = Handle{} }
func (m *Handle) String() string            { return proto.CompactTextString(m) }
func (*Handle) ProtoMessage()               {}
func (*Handle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ActionItem struct {
	Cookie     uint64 `protobuf:"varint,1,opt,name=cookie" json:"cookie,omitempty"`
	Op         uint32 `protobuf:"varint,2,opt,name=op" json:"op,omitempty"`
	PrimaryFid string `protobuf:"bytes,3,opt,name=primary_fid" json:"primary_fid,omitempty"`
	WriteFid   string `protobuf:"bytes,4,opt,name=write_fid" json:"write_fid,omitempty"`
	Offset     uint64 `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Length     uint64 `protobuf:"varint,6,opt,name=length" json:"length,omitempty"`
	Data       []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ActionItem) Reset()                    { *m = ActionItem{} }
func (m *ActionItem) String() string            { return proto.CompactTextString(m) }
func (*ActionItem) ProtoMessage()               {}
func (*ActionItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ActionStatus struct {
	Cookie    uint64  `protobuf:"varint,1,opt,name=cookie" json:"cookie,omitempty"`
	Completed bool    `protobuf:"varint,2,opt,name=completed" json:"completed,omitempty"`
	Error     int32   `protobuf:"varint,3,opt,name=error" json:"error,omitempty"`
	Offset    uint64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	Length    uint64  `protobuf:"varint,5,opt,name=length" json:"length,omitempty"`
	Handle    *Handle `protobuf:"bytes,6,opt,name=handle" json:"handle,omitempty"`
}

func (m *ActionStatus) Reset()                    { *m = ActionStatus{} }
func (m *ActionStatus) String() string            { return proto.CompactTextString(m) }
func (*ActionStatus) ProtoMessage()               {}
func (*ActionStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ActionStatus) GetHandle() *Handle {
	if m != nil {
		return m.Handle
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Endpoint)(nil), "pdm.Endpoint")
	proto.RegisterType((*Handle)(nil), "pdm.Handle")
	proto.RegisterType((*ActionItem)(nil), "pdm.ActionItem")
	proto.RegisterType((*ActionStatus)(nil), "pdm.ActionStatus")
	proto.RegisterType((*Empty)(nil), "pdm.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for DataMover service

type DataMoverClient interface {
	Register(ctx context.Context, in *Endpoint, opts ...grpc.CallOption) (*Handle, error)
	GetActions(ctx context.Context, in *Handle, opts ...grpc.CallOption) (DataMover_GetActionsClient, error)
	StatusStream(ctx context.Context, opts ...grpc.CallOption) (DataMover_StatusStreamClient, error)
}

type dataMoverClient struct {
	cc *grpc.ClientConn
}

func NewDataMoverClient(cc *grpc.ClientConn) DataMoverClient {
	return &dataMoverClient{cc}
}

func (c *dataMoverClient) Register(ctx context.Context, in *Endpoint, opts ...grpc.CallOption) (*Handle, error) {
	out := new(Handle)
	err := grpc.Invoke(ctx, "/pdm.DataMover/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataMoverClient) GetActions(ctx context.Context, in *Handle, opts ...grpc.CallOption) (DataMover_GetActionsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DataMover_serviceDesc.Streams[0], c.cc, "/pdm.DataMover/GetActions", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataMoverGetActionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataMover_GetActionsClient interface {
	Recv() (*ActionItem, error)
	grpc.ClientStream
}

type dataMoverGetActionsClient struct {
	grpc.ClientStream
}

func (x *dataMoverGetActionsClient) Recv() (*ActionItem, error) {
	m := new(ActionItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataMoverClient) StatusStream(ctx context.Context, opts ...grpc.CallOption) (DataMover_StatusStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DataMover_serviceDesc.Streams[1], c.cc, "/pdm.DataMover/StatusStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataMoverStatusStreamClient{stream}
	return x, nil
}

type DataMover_StatusStreamClient interface {
	Send(*ActionStatus) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type dataMoverStatusStreamClient struct {
	grpc.ClientStream
}

func (x *dataMoverStatusStreamClient) Send(m *ActionStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataMoverStatusStreamClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DataMover service

type DataMoverServer interface {
	Register(context.Context, *Endpoint) (*Handle, error)
	GetActions(*Handle, DataMover_GetActionsServer) error
	StatusStream(DataMover_StatusStreamServer) error
}

func RegisterDataMoverServer(s *grpc.Server, srv DataMoverServer) {
	s.RegisterService(&_DataMover_serviceDesc, srv)
}

func _DataMover_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Endpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DataMoverServer).Register(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _DataMover_GetActions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Handle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataMoverServer).GetActions(m, &dataMoverGetActionsServer{stream})
}

type DataMover_GetActionsServer interface {
	Send(*ActionItem) error
	grpc.ServerStream
}

type dataMoverGetActionsServer struct {
	grpc.ServerStream
}

func (x *dataMoverGetActionsServer) Send(m *ActionItem) error {
	return x.ServerStream.SendMsg(m)
}

func _DataMover_StatusStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataMoverServer).StatusStream(&dataMoverStatusStreamServer{stream})
}

type DataMover_StatusStreamServer interface {
	SendAndClose(*Empty) error
	Recv() (*ActionStatus, error)
	grpc.ServerStream
}

type dataMoverStatusStreamServer struct {
	grpc.ServerStream
}

func (x *dataMoverStatusStreamServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataMoverStatusStreamServer) Recv() (*ActionStatus, error) {
	m := new(ActionStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataMover_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pdm.DataMover",
	HandlerType: (*DataMoverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _DataMover_Register_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetActions",
			Handler:       _DataMover_GetActions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StatusStream",
			Handler:       _DataMover_StatusStream_Handler,
			ClientStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x51, 0xc1, 0x4e, 0xeb, 0x30,
	0x10, 0x6c, 0xda, 0x24, 0x6d, 0xb7, 0xc9, 0xab, 0xea, 0xc7, 0x21, 0x6a, 0x2f, 0xc8, 0xa7, 0x9e,
	0x0a, 0x94, 0x2f, 0x40, 0xa2, 0x02, 0x0e, 0x5c, 0xe8, 0x07, 0x20, 0x93, 0x6c, 0x5b, 0x8b, 0x26,
	0xb6, 0x9c, 0xa5, 0xa8, 0x07, 0xf8, 0x0e, 0x3e, 0x17, 0xc7, 0x16, 0x22, 0xc0, 0x21, 0x92, 0x67,
	0xbc, 0xe3, 0x99, 0x9d, 0x40, 0xaa, 0x8b, 0xf2, 0xcc, 0x7e, 0x0b, 0x6d, 0x14, 0x29, 0xd6, 0xb3,
	0x47, 0x3e, 0x83, 0xc1, 0xaa, 0x2a, 0xb4, 0x92, 0x15, 0xb1, 0x31, 0xf4, 0x85, 0xc9, 0x77, 0xf2,
	0x80, 0x59, 0x70, 0x1a, 0xcc, 0x53, 0x7e, 0x02, 0xf1, 0xad, 0xa8, 0x8a, 0x3d, 0x32, 0x80, 0xae,
	0x2c, 0x1c, 0x1b, 0xf2, 0x77, 0x80, 0xab, 0x9c, 0xa4, 0xaa, 0xee, 0x08, 0x4b, 0xf6, 0x0f, 0xe2,
	0x5c, 0xa9, 0x67, 0xe9, 0x35, 0x61, 0x33, 0xa9, 0x74, 0xd6, 0x6d, 0xf4, 0xec, 0x3f, 0x8c, 0xb4,
	0x91, 0xa5, 0x30, 0xc7, 0xc7, 0x8d, 0x95, 0xf7, 0x2c, 0x39, 0x64, 0x13, 0x18, 0xbe, 0x1a, 0x49,
	0xe8, 0xa8, 0xd0, 0x51, 0xf6, 0x0d, 0xb5, 0xd9, 0xd4, 0x48, 0x59, 0xe4, 0xde, 0xb0, 0x78, 0x8f,
	0xd5, 0x96, 0x76, 0x59, 0xec, 0x70, 0x02, 0x61, 0x21, 0x48, 0x64, 0x7d, 0x8b, 0x12, 0xfe, 0x06,
	0x89, 0xf7, 0x5f, 0x93, 0xa0, 0x97, 0xfa, 0x4f, 0x02, 0x6b, 0x90, 0xab, 0x52, 0xef, 0x91, 0xb0,
	0x70, 0x41, 0x06, 0x2c, 0x85, 0x08, 0x8d, 0x51, 0xc6, 0x45, 0x88, 0x5a, 0x7e, 0xe1, 0x2f, 0x3f,
	0xef, 0x3f, 0x83, 0x78, 0xe7, 0xf6, 0x76, 0xfe, 0xa3, 0xe5, 0x68, 0xd1, 0xb4, 0xe6, 0xab, 0xe0,
	0x7d, 0x88, 0x56, 0xa5, 0xa6, 0xe3, 0xf2, 0x23, 0x80, 0xe1, 0xb5, 0x8d, 0x75, 0xaf, 0x0e, 0x68,
	0xd8, 0x1c, 0x06, 0x0f, 0xb8, 0x95, 0x35, 0xd9, 0x73, 0xea, 0xe6, 0xbf, 0x7a, 0x9d, 0xfe, 0x90,
	0x77, 0xd8, 0x02, 0xe0, 0x06, 0xc9, 0xaf, 0x50, 0xb3, 0xf6, 0xe5, 0x74, 0xec, 0xc0, 0x77, 0xbb,
	0xbc, 0x73, 0x1e, 0xb0, 0x0b, 0x48, 0xfc, 0xa6, 0x6b, 0x32, 0x28, 0x4a, 0x36, 0x69, 0x0d, 0xf9,
	0x8b, 0x29, 0x78, 0xc3, 0x26, 0x16, 0xef, 0xcc, 0x83, 0xa7, 0xd8, 0xfd, 0xe1, 0xcb, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0xb1, 0x81, 0x3a, 0xf2, 0x01, 0x00, 0x00,
}
