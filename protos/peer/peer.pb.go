
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:35</date>
//</624456121196744704>

//由Protoc Gen Go生成的代码。不要编辑。
//来源：peer/peer.proto

package peer //导入“github.com/hyperledger/fabric/protos/peer”

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

//引用导入以禁止错误（如果未使用）。
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//这是一个编译时断言，以确保生成的文件
//与正在编译的proto包兼容。
//此行的编译错误可能意味着您的
//需要更新proto包。
const _ = proto.ProtoPackageIsVersion2 //请升级proto包

type PeerID struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerID) Reset()         { *m = PeerID{} }
func (m *PeerID) String() string { return proto.CompactTextString(m) }
func (*PeerID) ProtoMessage()    {}
func (*PeerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_peer_f225b6e437338532, []int{0}
}
func (m *PeerID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerID.Unmarshal(m, b)
}
func (m *PeerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerID.Marshal(b, m, deterministic)
}
func (dst *PeerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerID.Merge(dst, src)
}
func (m *PeerID) XXX_Size() int {
	return xxx_messageInfo_PeerID.Size(m)
}
func (m *PeerID) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerID.DiscardUnknown(m)
}

var xxx_messageInfo_PeerID proto.InternalMessageInfo

func (m *PeerID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PeerEndpoint struct {
	Id                   *PeerID  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerEndpoint) Reset()         { *m = PeerEndpoint{} }
func (m *PeerEndpoint) String() string { return proto.CompactTextString(m) }
func (*PeerEndpoint) ProtoMessage()    {}
func (*PeerEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_peer_f225b6e437338532, []int{1}
}
func (m *PeerEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerEndpoint.Unmarshal(m, b)
}
func (m *PeerEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerEndpoint.Marshal(b, m, deterministic)
}
func (dst *PeerEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerEndpoint.Merge(dst, src)
}
func (m *PeerEndpoint) XXX_Size() int {
	return xxx_messageInfo_PeerEndpoint.Size(m)
}
func (m *PeerEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_PeerEndpoint proto.InternalMessageInfo

func (m *PeerEndpoint) GetId() *PeerID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PeerEndpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*PeerID)(nil), "protos.PeerID")
	proto.RegisterType((*PeerEndpoint)(nil), "protos.PeerEndpoint")
}

//引用导入以禁止错误（如果未使用）。
var _ context.Context
var _ grpc.ClientConn

//这是一个编译时断言，以确保生成的文件
//与正在编译的GRPC包兼容。
const _ = grpc.SupportPackageIsVersion4

//背书人客户端是背书人服务的客户端API。
//
//有关CTX使用和关闭/结束流式RPC的语义，请参阅https://godoc.org/google.golang.org/grpc clientconn.newstream。
type EndorserClient interface {
	ProcessProposal(ctx context.Context, in *SignedProposal, opts ...grpc.CallOption) (*ProposalResponse, error)
}

type endorserClient struct {
	cc *grpc.ClientConn
}

func NewEndorserClient(cc *grpc.ClientConn) EndorserClient {
	return &endorserClient{cc}
}

func (c *endorserClient) ProcessProposal(ctx context.Context, in *SignedProposal, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := c.cc.Invoke(ctx, "/protos.Endorser/ProcessProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//背书服务器是背书服务的服务器API。
type EndorserServer interface {
	ProcessProposal(context.Context, *SignedProposal) (*ProposalResponse, error)
}

func RegisterEndorserServer(s *grpc.Server, srv EndorserServer) {
	s.RegisterService(&_Endorser_serviceDesc, srv)
}

func _Endorser_ProcessProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndorserServer).ProcessProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Endorser/ProcessProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndorserServer).ProcessProposal(ctx, req.(*SignedProposal))
	}
	return interceptor(ctx, in, info, handler)
}

var _Endorser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Endorser",
	HandlerType: (*EndorserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessProposal",
			Handler:    _Endorser_ProcessProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer/peer.proto",
}

func init() { proto.RegisterFile("peer/peer.proto", fileDescriptor_peer_f225b6e437338532) }

var fileDescriptor_peer_f225b6e437338532 = []byte{
//gzip文件描述符或协议的246字节
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xdd, 0x22, 0xab, 0x8e, 0xe2, 0x42, 0x04, 0x29, 0x65, 0x11, 0xe9, 0x49, 0x2f, 0x29,
	0xac, 0xdf, 0x40, 0x2c, 0xe8, 0xc9, 0x5a, 0x6f, 0x5e, 0xa4, 0x6d, 0xc6, 0x6e, 0x60, 0x37, 0x13,
	0x66, 0xea, 0xc1, 0x6f, 0x2f, 0x4d, 0x5a, 0x71, 0x2f, 0xf9, 0xf3, 0xde, 0x6f, 0x5e, 0x26, 0x03,
	0x2b, 0x8f, 0xc8, 0xc5, 0xb8, 0x68, 0xcf, 0x34, 0x90, 0x5a, 0x86, 0x4d, 0xb2, 0xab, 0x68, 0x30,
	0x79, 0x92, 0x66, 0x17, 0xcd, 0x6c, 0x7d, 0x20, 0x7e, 0x32, 0x8a, 0x27, 0x27, 0x18, 0xdd, 0x7c,
	0x0d, 0xcb, 0x0a, 0x91, 0x5f, 0x9e, 0x94, 0x82, 0x63, 0xd7, 0xec, 0x31, 0x5d, 0xdc, 0x2e, 0xee,
	0xce, 0xea, 0x70, 0xce, 0x9f, 0xe1, 0x62, 0x74, 0x4b, 0x67, 0x3c, 0x59, 0x37, 0xa8, 0x1b, 0x48,
	0xac, 0x09, 0xc4, 0xf9, 0xe6, 0x32, 0x26, 0x88, 0x8e, 0xf5, 0x75, 0x62, 0x8d, 0x4a, 0xe1, 0xa4,
	0x31, 0x86, 0x51, 0x24, 0x4d, 0x42, 0xcc, 0x7c, 0xdd, 0xbc, 0xc1, 0x69, 0xe9, 0x0c, 0xb1, 0x20,
	0xab, 0x12, 0x56, 0x15, 0x53, 0x87, 0x22, 0xd5, 0xd4, 0x95, 0xba, 0x9e, 0xc3, 0xde, 0x6d, 0xef,
	0xd0, 0xcc, 0x7a, 0x96, 0xfe, 0x3d, 0x32, 0x29, 0xf5, 0xd4, 0x7e, 0x7e, 0xf4, 0xf8, 0x0a, 0x39,
	0x71, 0xaf, 0xb7, 0x3f, 0x1e, 0x79, 0x87, 0xa6, 0x47, 0xd6, 0x5f, 0x4d, 0xcb, 0xb6, 0x9b, 0x6b,
	0xc6, 0x8f, 0x7f, 0xdc, 0xf7, 0x76, 0xd8, 0x7e, 0xb7, 0xba, 0xa3, 0x7d, 0xf1, 0x0f, 0x2d, 0x22,
	0x5a, 0x44, 0x34, 0x0c, 0xb3, 0x8d, 0x63, 0x7c, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xef, 0x32,
	0xf2, 0x1f, 0x60, 0x01, 0x00, 0x00,
}
