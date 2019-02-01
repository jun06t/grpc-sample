// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uploader.proto

/*
Package upload is a generated protocol buffer package.

It is generated from these files:
	uploader.proto

It has these top-level messages:
	UploadRequest
	Meta
	Chunk
	UploadResponse
*/
package upload

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

type UploadRequest struct {
	// Types that are valid to be assigned to Value:
	//	*UploadRequest_Meta
	//	*UploadRequest_Chunk
	Value isUploadRequest_Value `protobuf_oneof:"value"`
}

func (m *UploadRequest) Reset()                    { *m = UploadRequest{} }
func (m *UploadRequest) String() string            { return proto.CompactTextString(m) }
func (*UploadRequest) ProtoMessage()               {}
func (*UploadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isUploadRequest_Value interface {
	isUploadRequest_Value()
}

type UploadRequest_Meta struct {
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,oneof"`
}
type UploadRequest_Chunk struct {
	Chunk *Chunk `protobuf:"bytes,2,opt,name=chunk,oneof"`
}

func (*UploadRequest_Meta) isUploadRequest_Value()  {}
func (*UploadRequest_Chunk) isUploadRequest_Value() {}

func (m *UploadRequest) GetValue() isUploadRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *UploadRequest) GetMeta() *Meta {
	if x, ok := m.GetValue().(*UploadRequest_Meta); ok {
		return x.Meta
	}
	return nil
}

func (m *UploadRequest) GetChunk() *Chunk {
	if x, ok := m.GetValue().(*UploadRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UploadRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UploadRequest_OneofMarshaler, _UploadRequest_OneofUnmarshaler, _UploadRequest_OneofSizer, []interface{}{
		(*UploadRequest_Meta)(nil),
		(*UploadRequest_Chunk)(nil),
	}
}

func _UploadRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UploadRequest)
	// value
	switch x := m.Value.(type) {
	case *UploadRequest_Meta:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Meta); err != nil {
			return err
		}
	case *UploadRequest_Chunk:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Chunk); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UploadRequest.Value has unexpected type %T", x)
	}
	return nil
}

func _UploadRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UploadRequest)
	switch tag {
	case 1: // value.meta
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Meta)
		err := b.DecodeMessage(msg)
		m.Value = &UploadRequest_Meta{msg}
		return true, err
	case 2: // value.chunk
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Chunk)
		err := b.DecodeMessage(msg)
		m.Value = &UploadRequest_Chunk{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UploadRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UploadRequest)
	// value
	switch x := m.Value.(type) {
	case *UploadRequest_Meta:
		s := proto.Size(x.Meta)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UploadRequest_Chunk:
		s := proto.Size(x.Chunk)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Meta struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Quality string `protobuf:"bytes,3,opt,name=quality" json:"quality,omitempty"`
}

func (m *Meta) Reset()                    { *m = Meta{} }
func (m *Meta) String() string            { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()               {}
func (*Meta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Meta) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Meta) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Meta) GetQuality() string {
	if m != nil {
		return m.Quality
	}
	return ""
}

type Chunk struct {
	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Position int64  `protobuf:"varint,2,opt,name=position" json:"position,omitempty"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Chunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Chunk) GetPosition() int64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type UploadResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *UploadResponse) Reset()                    { *m = UploadResponse{} }
func (m *UploadResponse) String() string            { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()               {}
func (*UploadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UploadResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*UploadRequest)(nil), "upload.UploadRequest")
	proto.RegisterType((*Meta)(nil), "upload.Meta")
	proto.RegisterType((*Chunk)(nil), "upload.Chunk")
	proto.RegisterType((*UploadResponse)(nil), "upload.UploadResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Uploader service

type UploaderClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (Uploader_UploadClient, error)
}

type uploaderClient struct {
	cc *grpc.ClientConn
}

func NewUploaderClient(cc *grpc.ClientConn) UploaderClient {
	return &uploaderClient{cc}
}

func (c *uploaderClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Uploader_UploadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Uploader_serviceDesc.Streams[0], c.cc, "/upload.Uploader/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploaderUploadClient{stream}
	return x, nil
}

type Uploader_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type uploaderUploadClient struct {
	grpc.ClientStream
}

func (x *uploaderUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploaderUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Uploader service

type UploaderServer interface {
	Upload(Uploader_UploadServer) error
}

func RegisterUploaderServer(s *grpc.Server, srv UploaderServer) {
	s.RegisterService(&_Uploader_serviceDesc, srv)
}

func _Uploader_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploaderServer).Upload(&uploaderUploadServer{stream})
}

type Uploader_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type uploaderUploadServer struct {
	grpc.ServerStream
}

func (x *uploaderUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploaderUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Uploader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "upload.Uploader",
	HandlerType: (*UploaderServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Uploader_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "uploader.proto",
}

func init() { proto.RegisterFile("uploader.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0xb5, 0x1d, 0xdb, 0x49, 0xae, 0x89, 0x87, 0x83, 0x06, 0x93, 0xa9, 0x08, 0x0a, 0x9e, 0x3c,
	0xa4, 0x43, 0x87, 0x6e, 0x6d, 0xa1, 0x59, 0xba, 0x08, 0xf2, 0x01, 0x6a, 0x2c, 0xa8, 0xa9, 0x6b,
	0x29, 0xd1, 0xa9, 0x90, 0xbf, 0x2f, 0x3e, 0xd9, 0x85, 0x66, 0x7b, 0xef, 0xe9, 0xdd, 0xd3, 0xdd,
	0x83, 0xc2, 0xdb, 0xce, 0xa8, 0x46, 0x9f, 0x6b, 0x7b, 0x36, 0x64, 0x30, 0x0f, 0x5c, 0x1c, 0x61,
	0x7d, 0x60, 0x24, 0xf5, 0xc9, 0x6b, 0x47, 0x28, 0x20, 0xfd, 0xd6, 0xa4, 0xca, 0xf8, 0x2e, 0xae,
	0x6e, 0x76, 0xab, 0x3a, 0xf8, 0xea, 0x77, 0x4d, 0x6a, 0x1f, 0x49, 0x7e, 0xc3, 0x7b, 0xc8, 0x8e,
	0x9f, 0xbe, 0xff, 0x2a, 0x13, 0x36, 0xad, 0x27, 0xd3, 0xcb, 0x20, 0xee, 0x23, 0x19, 0x5e, 0x9f,
	0xe7, 0x90, 0xfd, 0xa8, 0xce, 0x6b, 0xf1, 0x0a, 0xe9, 0x30, 0x8f, 0x05, 0x24, 0x6d, 0xc3, 0xc9,
	0x4b, 0x99, 0xb4, 0x0d, 0x22, 0xa4, 0x74, 0xb1, 0x9a, 0x63, 0x96, 0x92, 0x31, 0x96, 0x30, 0x3f,
	0x79, 0xd5, 0xb5, 0x74, 0x29, 0x67, 0x2c, 0x4f, 0x54, 0x3c, 0x42, 0xc6, 0x1f, 0x0c, 0x63, 0x8d,
	0x1a, 0x57, 0x5c, 0x49, 0xc6, 0xb8, 0x85, 0x85, 0x35, 0xae, 0xa5, 0xd6, 0xf4, 0x1c, 0x37, 0x93,
	0x7f, 0x5c, 0x54, 0x50, 0x4c, 0x37, 0x3a, 0x6b, 0x7a, 0xa7, 0x71, 0x03, 0xb9, 0x23, 0x45, 0xde,
	0x8d, 0xcb, 0x8c, 0x6c, 0xf7, 0x06, 0x8b, 0xc3, 0xd8, 0x13, 0x3e, 0x41, 0x1e, 0x30, 0xde, 0x4e,
	0xf7, 0xfd, 0x6b, 0x6a, 0xbb, 0xb9, 0x96, 0x43, 0xb8, 0x88, 0xaa, 0xf8, 0x23, 0xe7, 0x96, 0x1f,
	0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x16, 0xa1, 0xc5, 0x5f, 0x77, 0x01, 0x00, 0x00,
}