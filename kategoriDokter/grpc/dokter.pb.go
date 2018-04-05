// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dokter.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	dokter.proto

It has these top-level messages:
	AddDokterReq
	ReadDokterByKodeReq
	ReadDokterByKodeResp
	ReadDokterResp
	UpdateDokterReq
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type AddDokterReq struct {
	KodeDokter         string `protobuf:"bytes,1,opt,name=kodeDokter" json:"kodeDokter,omitempty"`
	NamaDokter         string `protobuf:"bytes,2,opt,name=namaDokter" json:"namaDokter,omitempty"`
	JenisKelamin       string `protobuf:"bytes,3,opt,name=jenisKelamin" json:"jenisKelamin,omitempty"`
	Alamat             string `protobuf:"bytes,4,opt,name=alamat" json:"alamat,omitempty"`
	NomorTelepon       string `protobuf:"bytes,5,opt,name=nomorTelepon" json:"nomorTelepon,omitempty"`
	KodeKategoriDokter string `protobuf:"bytes,6,opt,name=kodeKategoriDokter" json:"kodeKategoriDokter,omitempty"`
	Biaya              int64  `protobuf:"varint,7,opt,name=biaya" json:"biaya,omitempty"`
}

func (m *AddDokterReq) Reset()                    { *m = AddDokterReq{} }
func (m *AddDokterReq) String() string            { return proto.CompactTextString(m) }
func (*AddDokterReq) ProtoMessage()               {}
func (*AddDokterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddDokterReq) GetKodeDokter() string {
	if m != nil {
		return m.KodeDokter
	}
	return ""
}

func (m *AddDokterReq) GetNamaDokter() string {
	if m != nil {
		return m.NamaDokter
	}
	return ""
}

func (m *AddDokterReq) GetJenisKelamin() string {
	if m != nil {
		return m.JenisKelamin
	}
	return ""
}

func (m *AddDokterReq) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *AddDokterReq) GetNomorTelepon() string {
	if m != nil {
		return m.NomorTelepon
	}
	return ""
}

func (m *AddDokterReq) GetKodeKategoriDokter() string {
	if m != nil {
		return m.KodeKategoriDokter
	}
	return ""
}

func (m *AddDokterReq) GetBiaya() int64 {
	if m != nil {
		return m.Biaya
	}
	return 0
}

// request adalah yang datang lalu ditampung
type ReadDokterByKodeReq struct {
	Kode string `protobuf:"bytes,1,opt,name=kode" json:"kode,omitempty"`
}

func (m *ReadDokterByKodeReq) Reset()                    { *m = ReadDokterByKodeReq{} }
func (m *ReadDokterByKodeReq) String() string            { return proto.CompactTextString(m) }
func (*ReadDokterByKodeReq) ProtoMessage()               {}
func (*ReadDokterByKodeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadDokterByKodeReq) GetKode() string {
	if m != nil {
		return m.Kode
	}
	return ""
}

// request adalah sesuatu yang return yang akan dikembalikan biasanya ada request dan respon
type ReadDokterByKodeResp struct {
	KodeDokter         string `protobuf:"bytes,1,opt,name=kodeDokter" json:"kodeDokter,omitempty"`
	NamaDokter         string `protobuf:"bytes,2,opt,name=namaDokter" json:"namaDokter,omitempty"`
	JenisKelamin       string `protobuf:"bytes,3,opt,name=jenisKelamin" json:"jenisKelamin,omitempty"`
	Alamat             string `protobuf:"bytes,4,opt,name=alamat" json:"alamat,omitempty"`
	NomorTelepon       string `protobuf:"bytes,5,opt,name=nomorTelepon" json:"nomorTelepon,omitempty"`
	KodeKategoriDokter string `protobuf:"bytes,6,opt,name=kodeKategoriDokter" json:"kodeKategoriDokter,omitempty"`
	Biaya              int64  `protobuf:"varint,7,opt,name=biaya" json:"biaya,omitempty"`
}

func (m *ReadDokterByKodeResp) Reset()                    { *m = ReadDokterByKodeResp{} }
func (m *ReadDokterByKodeResp) String() string            { return proto.CompactTextString(m) }
func (*ReadDokterByKodeResp) ProtoMessage()               {}
func (*ReadDokterByKodeResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadDokterByKodeResp) GetKodeDokter() string {
	if m != nil {
		return m.KodeDokter
	}
	return ""
}

func (m *ReadDokterByKodeResp) GetNamaDokter() string {
	if m != nil {
		return m.NamaDokter
	}
	return ""
}

func (m *ReadDokterByKodeResp) GetJenisKelamin() string {
	if m != nil {
		return m.JenisKelamin
	}
	return ""
}

func (m *ReadDokterByKodeResp) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *ReadDokterByKodeResp) GetNomorTelepon() string {
	if m != nil {
		return m.NomorTelepon
	}
	return ""
}

func (m *ReadDokterByKodeResp) GetKodeKategoriDokter() string {
	if m != nil {
		return m.KodeKategoriDokter
	}
	return ""
}

func (m *ReadDokterByKodeResp) GetBiaya() int64 {
	if m != nil {
		return m.Biaya
	}
	return 0
}

type ReadDokterResp struct {
	// repeated itu sama kayak menulis ulang
	Allkode []*ReadDokterByKodeResp `protobuf:"bytes,1,rep,name=allkode" json:"allkode,omitempty"`
}

func (m *ReadDokterResp) Reset()                    { *m = ReadDokterResp{} }
func (m *ReadDokterResp) String() string            { return proto.CompactTextString(m) }
func (*ReadDokterResp) ProtoMessage()               {}
func (*ReadDokterResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadDokterResp) GetAllkode() []*ReadDokterByKodeResp {
	if m != nil {
		return m.Allkode
	}
	return nil
}

type UpdateDokterReq struct {
	KodeDokter         string `protobuf:"bytes,1,opt,name=kodeDokter" json:"kodeDokter,omitempty"`
	NamaDokter         string `protobuf:"bytes,2,opt,name=namaDokter" json:"namaDokter,omitempty"`
	JenisKelamin       string `protobuf:"bytes,3,opt,name=jenisKelamin" json:"jenisKelamin,omitempty"`
	Alamat             string `protobuf:"bytes,4,opt,name=alamat" json:"alamat,omitempty"`
	NomorTelepon       string `protobuf:"bytes,5,opt,name=nomorTelepon" json:"nomorTelepon,omitempty"`
	KodeKategoriDokter string `protobuf:"bytes,6,opt,name=kodeKategoriDokter" json:"kodeKategoriDokter,omitempty"`
	Biaya              int64  `protobuf:"varint,7,opt,name=biaya" json:"biaya,omitempty"`
}

func (m *UpdateDokterReq) Reset()                    { *m = UpdateDokterReq{} }
func (m *UpdateDokterReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateDokterReq) ProtoMessage()               {}
func (*UpdateDokterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateDokterReq) GetKodeDokter() string {
	if m != nil {
		return m.KodeDokter
	}
	return ""
}

func (m *UpdateDokterReq) GetNamaDokter() string {
	if m != nil {
		return m.NamaDokter
	}
	return ""
}

func (m *UpdateDokterReq) GetJenisKelamin() string {
	if m != nil {
		return m.JenisKelamin
	}
	return ""
}

func (m *UpdateDokterReq) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *UpdateDokterReq) GetNomorTelepon() string {
	if m != nil {
		return m.NomorTelepon
	}
	return ""
}

func (m *UpdateDokterReq) GetKodeKategoriDokter() string {
	if m != nil {
		return m.KodeKategoriDokter
	}
	return ""
}

func (m *UpdateDokterReq) GetBiaya() int64 {
	if m != nil {
		return m.Biaya
	}
	return 0
}

func init() {
	proto.RegisterType((*AddDokterReq)(nil), "grpc.AddDokterReq")
	proto.RegisterType((*ReadDokterByKodeReq)(nil), "grpc.ReadDokterByKodeReq")
	proto.RegisterType((*ReadDokterByKodeResp)(nil), "grpc.ReadDokterByKodeResp")
	proto.RegisterType((*ReadDokterResp)(nil), "grpc.ReadDokterResp")
	proto.RegisterType((*UpdateDokterReq)(nil), "grpc.UpdateDokterReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for DokterService service

type DokterServiceClient interface {
	AddDokter(ctx context.Context, in *AddDokterReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadDokterByKode(ctx context.Context, in *ReadDokterByKodeReq, opts ...grpc1.CallOption) (*ReadDokterByKodeResp, error)
	ReadDokter(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadDokterResp, error)
	UpdateDokter(ctx context.Context, in *UpdateDokterReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
}

type dokterServiceClient struct {
	cc *grpc1.ClientConn
}

func NewDokterServiceClient(cc *grpc1.ClientConn) DokterServiceClient {
	return &dokterServiceClient{cc}
}

func (c *dokterServiceClient) AddDokter(ctx context.Context, in *AddDokterReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.DokterService/AddDokter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokterServiceClient) ReadDokterByKode(ctx context.Context, in *ReadDokterByKodeReq, opts ...grpc1.CallOption) (*ReadDokterByKodeResp, error) {
	out := new(ReadDokterByKodeResp)
	err := grpc1.Invoke(ctx, "/grpc.DokterService/ReadDokterByKode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokterServiceClient) ReadDokter(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadDokterResp, error) {
	out := new(ReadDokterResp)
	err := grpc1.Invoke(ctx, "/grpc.DokterService/ReadDokter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokterServiceClient) UpdateDokter(ctx context.Context, in *UpdateDokterReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.DokterService/UpdateDokter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DokterService service

type DokterServiceServer interface {
	AddDokter(context.Context, *AddDokterReq) (*google_protobuf.Empty, error)
	ReadDokterByKode(context.Context, *ReadDokterByKodeReq) (*ReadDokterByKodeResp, error)
	ReadDokter(context.Context, *google_protobuf.Empty) (*ReadDokterResp, error)
	UpdateDokter(context.Context, *UpdateDokterReq) (*google_protobuf.Empty, error)
}

func RegisterDokterServiceServer(s *grpc1.Server, srv DokterServiceServer) {
	s.RegisterService(&_DokterService_serviceDesc, srv)
}

func _DokterService_AddDokter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDokterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokterServiceServer).AddDokter(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DokterService/AddDokter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokterServiceServer).AddDokter(ctx, req.(*AddDokterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DokterService_ReadDokterByKode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadDokterByKodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokterServiceServer).ReadDokterByKode(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DokterService/ReadDokterByKode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokterServiceServer).ReadDokterByKode(ctx, req.(*ReadDokterByKodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DokterService_ReadDokter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokterServiceServer).ReadDokter(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DokterService/ReadDokter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokterServiceServer).ReadDokter(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DokterService_UpdateDokter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDokterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokterServiceServer).UpdateDokter(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DokterService/UpdateDokter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokterServiceServer).UpdateDokter(ctx, req.(*UpdateDokterReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DokterService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.DokterService",
	HandlerType: (*DokterServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddDokter",
			Handler:    _DokterService_AddDokter_Handler,
		},
		{
			MethodName: "ReadDokterByKode",
			Handler:    _DokterService_ReadDokterByKode_Handler,
		},
		{
			MethodName: "ReadDokter",
			Handler:    _DokterService_ReadDokter_Handler,
		},
		{
			MethodName: "UpdateDokter",
			Handler:    _DokterService_UpdateDokter_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "dokter.proto",
}

func init() { proto.RegisterFile("dokter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x53, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0xa5, 0xfc, 0x86, 0xfb, 0xf5, 0x53, 0x73, 0x45, 0x52, 0x6b, 0x62, 0x48, 0x57, 0xb8, 0x29,
	0x09, 0xba, 0x31, 0x31, 0x31, 0x1a, 0x75, 0xd3, 0x5d, 0xd5, 0x07, 0x18, 0xe8, 0xb5, 0xa9, 0xb4,
	0x9d, 0x52, 0xaa, 0x09, 0x2f, 0xe2, 0x93, 0x9a, 0xb8, 0x71, 0x61, 0xe6, 0x07, 0xa9, 0x08, 0xbe,
	0x00, 0xbb, 0xce, 0x39, 0xe7, 0x4e, 0xe7, 0x9c, 0x99, 0x03, 0x66, 0xc0, 0x27, 0x05, 0xe5, 0x6e,
	0x96, 0xf3, 0x82, 0x63, 0x3d, 0xcc, 0xb3, 0xb1, 0x7d, 0x14, 0x72, 0x1e, 0xc6, 0x34, 0x90, 0xd8,
	0xe8, 0xe5, 0x69, 0x40, 0x49, 0x56, 0xcc, 0x95, 0xc4, 0x79, 0x37, 0xc0, 0xbc, 0x0a, 0x82, 0x1b,
	0x39, 0xe6, 0xd3, 0x14, 0x8f, 0x01, 0x26, 0x3c, 0x20, 0x05, 0x58, 0x46, 0xcf, 0xe8, 0xb7, 0xfd,
	0x12, 0x22, 0xf8, 0x94, 0x25, 0x4c, 0xf3, 0x55, 0xc5, 0x2f, 0x11, 0x74, 0xc0, 0x7c, 0xa6, 0x34,
	0x9a, 0x79, 0x14, 0xb3, 0x24, 0x4a, 0xad, 0x9a, 0x54, 0xfc, 0xc0, 0xb0, 0x0b, 0x4d, 0x16, 0xb3,
	0x84, 0x15, 0x56, 0x5d, 0xb2, 0x7a, 0x25, 0x66, 0x53, 0x9e, 0xf0, 0xfc, 0x81, 0x62, 0xca, 0x78,
	0x6a, 0x35, 0xd4, 0x6c, 0x19, 0x43, 0x17, 0x50, 0x9c, 0xc6, 0x63, 0x05, 0x85, 0x3c, 0x8f, 0xf4,
	0x39, 0x9a, 0x52, 0xb9, 0x86, 0xc1, 0x0e, 0x34, 0x46, 0x11, 0x9b, 0x33, 0xab, 0xd5, 0x33, 0xfa,
	0x35, 0x5f, 0x2d, 0x9c, 0x13, 0xd8, 0xf7, 0x89, 0x69, 0xdb, 0xd7, 0x73, 0x8f, 0x07, 0x24, 0xcc,
	0x23, 0xd4, 0xc5, 0x16, 0xda, 0xb6, 0xfc, 0x76, 0x3e, 0x0d, 0xe8, 0xfc, 0xd6, 0xce, 0xb2, 0x2d,
	0x49, 0xea, 0x0e, 0x76, 0x96, 0xee, 0xa5, 0xef, 0x33, 0x68, 0xb1, 0x38, 0xd6, 0x39, 0xd5, 0xfa,
	0xff, 0x86, 0xb6, 0x2b, 0xde, 0x99, 0xbb, 0x2e, 0x24, 0x7f, 0x21, 0x75, 0x3e, 0x0c, 0xd8, 0x7d,
	0xcc, 0x02, 0x56, 0xd0, 0x96, 0xbd, 0xb5, 0xe1, 0x5b, 0x15, 0xfe, 0x2b, 0xc1, 0x3d, 0xe5, 0xaf,
	0xd1, 0x98, 0xf0, 0x1c, 0xda, 0xdf, 0x9d, 0x43, 0x54, 0xe9, 0x95, 0x4b, 0x68, 0x77, 0x5d, 0xd5,
	0x59, 0x77, 0xd1, 0x59, 0xf7, 0x56, 0x74, 0xd6, 0xa9, 0xa0, 0x07, 0x7b, 0xab, 0x39, 0xe3, 0xe1,
	0xa6, 0xfc, 0xa7, 0xf6, 0x1f, 0x57, 0xe3, 0x54, 0xf0, 0x02, 0x60, 0xc9, 0xe0, 0x86, 0x9f, 0xda,
	0x9d, 0xd5, 0x3d, 0xf4, 0xf4, 0x25, 0x98, 0xe5, 0x0b, 0xc5, 0x03, 0xa5, 0x5b, 0xb9, 0xe4, 0xcd,
	0x5e, 0x46, 0x4d, 0x89, 0x9c, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x21, 0xd9, 0x77, 0xb0, 0xb5,
	0x04, 0x00, 0x00,
}
