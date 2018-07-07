// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/datastore/v1/datastore.proto

/*
Package datastore is a generated protocol buffer package.

It is generated from these files:
	github.com/ehazlett/stellar/api/services/datastore/v1/datastore.proto

It has these top-level messages:
	AcquireLockRequest
	ReleaseLockRequest
	SetRequest
	KeyValue
	GetRequest
	SearchRequest
	SearchResponse
	GetResponse
	DeleteRequest
	BackupRequest
	BackupResponse
	RestoreRequest
*/
package datastore

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AcquireLockRequest struct {
	Timeout *google_protobuf1.Duration `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *AcquireLockRequest) Reset()                    { *m = AcquireLockRequest{} }
func (m *AcquireLockRequest) String() string            { return proto.CompactTextString(m) }
func (*AcquireLockRequest) ProtoMessage()               {}
func (*AcquireLockRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{0} }

func (m *AcquireLockRequest) GetTimeout() *google_protobuf1.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

type ReleaseLockRequest struct {
}

func (m *ReleaseLockRequest) Reset()                    { *m = ReleaseLockRequest{} }
func (m *ReleaseLockRequest) String() string            { return proto.CompactTextString(m) }
func (*ReleaseLockRequest) ProtoMessage()               {}
func (*ReleaseLockRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{1} }

type SetRequest struct {
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Sync   bool   `protobuf:"varint,4,opt,name=sync,proto3" json:"sync,omitempty"`
}

func (m *SetRequest) Reset()                    { *m = SetRequest{} }
func (m *SetRequest) String() string            { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()               {}
func (*SetRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{2} }

func (m *SetRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *SetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SetRequest) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

type KeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{3} }

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetRequest struct {
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{4} }

func (m *GetRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type SearchRequest struct {
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{5} }

func (m *SearchRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *SearchRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type SearchResponse struct {
	Bucket string      `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Data   []*KeyValue `protobuf:"bytes,2,rep,name=data" json:"data,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{6} }

func (m *SearchResponse) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *SearchResponse) GetData() []*KeyValue {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetResponse struct {
	Bucket string    `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Data   *KeyValue `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{7} }

func (m *GetResponse) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *GetResponse) GetData() *KeyValue {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteRequest struct {
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Sync   bool   `protobuf:"varint,3,opt,name=sync,proto3" json:"sync,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{8} }

func (m *DeleteRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *DeleteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DeleteRequest) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

type BackupRequest struct {
}

func (m *BackupRequest) Reset()                    { *m = BackupRequest{} }
func (m *BackupRequest) String() string            { return proto.CompactTextString(m) }
func (*BackupRequest) ProtoMessage()               {}
func (*BackupRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{9} }

type BackupResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BackupResponse) Reset()                    { *m = BackupResponse{} }
func (m *BackupResponse) String() string            { return proto.CompactTextString(m) }
func (*BackupResponse) ProtoMessage()               {}
func (*BackupResponse) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{10} }

func (m *BackupResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RestoreRequest struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *RestoreRequest) Reset()                    { *m = RestoreRequest{} }
func (m *RestoreRequest) String() string            { return proto.CompactTextString(m) }
func (*RestoreRequest) ProtoMessage()               {}
func (*RestoreRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{11} }

func (m *RestoreRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*AcquireLockRequest)(nil), "stellar.services.datastore.v1.AcquireLockRequest")
	proto.RegisterType((*ReleaseLockRequest)(nil), "stellar.services.datastore.v1.ReleaseLockRequest")
	proto.RegisterType((*SetRequest)(nil), "stellar.services.datastore.v1.SetRequest")
	proto.RegisterType((*KeyValue)(nil), "stellar.services.datastore.v1.KeyValue")
	proto.RegisterType((*GetRequest)(nil), "stellar.services.datastore.v1.GetRequest")
	proto.RegisterType((*SearchRequest)(nil), "stellar.services.datastore.v1.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "stellar.services.datastore.v1.SearchResponse")
	proto.RegisterType((*GetResponse)(nil), "stellar.services.datastore.v1.GetResponse")
	proto.RegisterType((*DeleteRequest)(nil), "stellar.services.datastore.v1.DeleteRequest")
	proto.RegisterType((*BackupRequest)(nil), "stellar.services.datastore.v1.BackupRequest")
	proto.RegisterType((*BackupResponse)(nil), "stellar.services.datastore.v1.BackupResponse")
	proto.RegisterType((*RestoreRequest)(nil), "stellar.services.datastore.v1.RestoreRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Datastore service

type DatastoreClient interface {
	AcquireLock(ctx context.Context, in *AcquireLockRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	ReleaseLock(ctx context.Context, in *ReleaseLockRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResponse, error)
	Restore(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type datastoreClient struct {
	cc *grpc.ClientConn
}

func NewDatastoreClient(cc *grpc.ClientConn) DatastoreClient {
	return &datastoreClient{cc}
}

func (c *datastoreClient) AcquireLock(ctx context.Context, in *AcquireLockRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.datastore.v1.Datastore/AcquireLock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreClient) ReleaseLock(ctx context.Context, in *ReleaseLockRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.datastore.v1.Datastore/ReleaseLock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.datastore.v1.Datastore/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/stellar.services.datastore.v1.Datastore/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/stellar.services.datastore.v1.Datastore/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.datastore.v1.Datastore/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreClient) Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResponse, error) {
	out := new(BackupResponse)
	err := grpc.Invoke(ctx, "/stellar.services.datastore.v1.Datastore/Backup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreClient) Restore(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.datastore.v1.Datastore/Restore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Datastore service

type DatastoreServer interface {
	AcquireLock(context.Context, *AcquireLockRequest) (*google_protobuf.Empty, error)
	ReleaseLock(context.Context, *ReleaseLockRequest) (*google_protobuf.Empty, error)
	Set(context.Context, *SetRequest) (*google_protobuf.Empty, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Delete(context.Context, *DeleteRequest) (*google_protobuf.Empty, error)
	Backup(context.Context, *BackupRequest) (*BackupResponse, error)
	Restore(context.Context, *RestoreRequest) (*google_protobuf.Empty, error)
}

func RegisterDatastoreServer(s *grpc.Server, srv DatastoreServer) {
	s.RegisterService(&_Datastore_serviceDesc, srv)
}

func _Datastore_AcquireLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcquireLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).AcquireLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.datastore.v1.Datastore/AcquireLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).AcquireLock(ctx, req.(*AcquireLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datastore_ReleaseLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).ReleaseLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.datastore.v1.Datastore/ReleaseLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).ReleaseLock(ctx, req.(*ReleaseLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datastore_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.datastore.v1.Datastore/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datastore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.datastore.v1.Datastore/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datastore_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.datastore.v1.Datastore/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datastore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.datastore.v1.Datastore/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datastore_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.datastore.v1.Datastore/Backup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).Backup(ctx, req.(*BackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datastore_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.datastore.v1.Datastore/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).Restore(ctx, req.(*RestoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Datastore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.datastore.v1.Datastore",
	HandlerType: (*DatastoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcquireLock",
			Handler:    _Datastore_AcquireLock_Handler,
		},
		{
			MethodName: "ReleaseLock",
			Handler:    _Datastore_ReleaseLock_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Datastore_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Datastore_Get_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Datastore_Search_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Datastore_Delete_Handler,
		},
		{
			MethodName: "Backup",
			Handler:    _Datastore_Backup_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _Datastore_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/datastore/v1/datastore.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/datastore/v1/datastore.proto", fileDescriptorDatastore)
}

var fileDescriptorDatastore = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xdf, 0x6b, 0xd3, 0x50,
	0x14, 0xc7, 0x49, 0xd3, 0x65, 0xdb, 0xa9, 0xad, 0x72, 0x19, 0x25, 0x46, 0x94, 0x12, 0x06, 0x56,
	0x71, 0x09, 0xed, 0xc0, 0x97, 0x3d, 0x8c, 0x8d, 0x8e, 0x29, 0xfe, 0x40, 0x52, 0x10, 0x11, 0x1f,
	0xbc, 0xc9, 0xce, 0xda, 0xd0, 0xb4, 0x37, 0x4b, 0x6e, 0x8a, 0xf5, 0x1f, 0xf5, 0xdf, 0x91, 0x26,
	0xf7, 0xda, 0x6c, 0x5d, 0xbc, 0x5d, 0xdf, 0x72, 0x72, 0xcf, 0xf9, 0xdc, 0x73, 0x72, 0xbe, 0x5f,
	0x02, 0x17, 0xa3, 0x90, 0x8f, 0x33, 0xdf, 0x09, 0xd8, 0xd4, 0xc5, 0x31, 0xfd, 0x1d, 0x21, 0xe7,
	0x6e, 0xca, 0x31, 0x8a, 0x68, 0xe2, 0xd2, 0x38, 0x74, 0x53, 0x4c, 0xe6, 0x61, 0x80, 0xa9, 0x7b,
	0x45, 0x39, 0x4d, 0x39, 0x4b, 0xd0, 0x9d, 0xf7, 0x56, 0x81, 0x13, 0x27, 0x8c, 0x33, 0xf2, 0x5c,
	0x94, 0x38, 0x32, 0xdd, 0x59, 0x65, 0xcc, 0x7b, 0xd6, 0xb3, 0x11, 0x63, 0xa3, 0x08, 0xdd, 0x3c,
	0xd9, 0xcf, 0xae, 0x5d, 0x9c, 0xc6, 0x7c, 0x51, 0xd4, 0x5a, 0x2f, 0xee, 0x1e, 0x5e, 0x65, 0x09,
	0xe5, 0x21, 0x9b, 0x15, 0xe7, 0xf6, 0x7b, 0x20, 0x67, 0xc1, 0x4d, 0x16, 0x26, 0xf8, 0x91, 0x05,
	0x13, 0x0f, 0x6f, 0x32, 0x4c, 0x39, 0x39, 0x86, 0x5d, 0x1e, 0x4e, 0x91, 0x65, 0xdc, 0xd4, 0x3a,
	0x5a, 0xb7, 0xd1, 0x7f, 0xea, 0x14, 0x1c, 0x47, 0x72, 0x9c, 0x81, 0xe0, 0x78, 0x32, 0xd3, 0x3e,
	0x00, 0xe2, 0x61, 0x84, 0x34, 0x2d, 0xa3, 0xec, 0x9f, 0x00, 0x43, 0xe4, 0x12, 0xdc, 0x06, 0xc3,
	0xcf, 0x82, 0x09, 0x16, 0xdc, 0x7d, 0x4f, 0x44, 0xe4, 0x09, 0xe8, 0x13, 0x5c, 0x98, 0xb5, 0xfc,
	0xe5, 0xf2, 0x91, 0x1c, 0xc0, 0xce, 0x9c, 0x46, 0x19, 0x9a, 0x7a, 0x47, 0xeb, 0x3e, 0xf2, 0x8a,
	0x80, 0x10, 0xa8, 0xa7, 0x8b, 0x59, 0x60, 0xd6, 0x3b, 0x5a, 0x77, 0xcf, 0xcb, 0x9f, 0xed, 0x3e,
	0xec, 0x7d, 0xc0, 0xc5, 0xd7, 0xfc, 0x5c, 0x70, 0xb4, 0x7b, 0x38, 0xb5, 0x12, 0xc7, 0x7e, 0x0b,
	0x70, 0xb9, 0x45, 0x57, 0xf6, 0x29, 0x34, 0x87, 0x48, 0x93, 0x60, 0xac, 0x2a, 0x6d, 0x83, 0x11,
	0x27, 0x78, 0x1d, 0xfe, 0x12, 0xd5, 0x22, 0xb2, 0x11, 0x5a, 0x12, 0x90, 0xc6, 0x6c, 0x96, 0x62,
	0x25, 0xe1, 0x04, 0xea, 0xcb, 0x35, 0x9b, 0xb5, 0x8e, 0xde, 0x6d, 0xf4, 0x5f, 0x3a, 0xff, 0x15,
	0x81, 0x23, 0xbf, 0x80, 0x97, 0x17, 0xd9, 0x3e, 0x34, 0xf2, 0xf9, 0x36, 0xbe, 0x43, 0x7b, 0xf8,
	0x1d, 0x9f, 0xa0, 0x39, 0xc0, 0x08, 0x39, 0x3e, 0x7c, 0xb9, 0x72, 0x8d, 0x7a, 0x69, 0x8d, 0x8f,
	0xa1, 0x79, 0x4e, 0x83, 0x49, 0x16, 0x4b, 0xe5, 0x1c, 0x42, 0x4b, 0xbe, 0x10, 0x63, 0x10, 0xd1,
	0xae, 0x96, 0xaf, 0xb2, 0xe8, 0xe2, 0x10, 0x5a, 0x1e, 0xe6, 0x2d, 0xca, 0x36, 0xee, 0xc9, 0xea,
	0xff, 0xd9, 0x81, 0xfd, 0x81, 0x9c, 0x85, 0x7c, 0x83, 0x46, 0x49, 0xf4, 0xa4, 0xa7, 0x98, 0x7b,
	0xdd, 0x20, 0x56, 0x7b, 0xcd, 0x0f, 0x17, 0x4b, 0xd3, 0x2d, 0xc9, 0x25, 0x0f, 0x28, 0xc9, 0xeb,
	0x7e, 0xa9, 0x24, 0xbf, 0x03, 0x7d, 0x88, 0x9c, 0xbc, 0x52, 0x10, 0x57, 0x5e, 0xab, 0x24, 0xfd,
	0x00, 0xfd, 0x72, 0x03, 0xd2, 0xca, 0x1f, 0xd6, 0xeb, 0x4d, 0x52, 0xc5, 0x8e, 0x10, 0x8c, 0x42,
	0xe0, 0xe4, 0x8d, 0xb2, 0xd5, 0x92, 0x91, 0xac, 0xa3, 0x0d, 0xb3, 0xc5, 0x35, 0x9f, 0xc1, 0x28,
	0xc4, 0xa7, 0xbc, 0xe6, 0x96, 0x46, 0x2b, 0x3f, 0x0a, 0x82, 0x51, 0x88, 0x4d, 0xc9, 0xbb, 0x25,
	0x52, 0x65, 0xdb, 0x77, 0x14, 0xfc, 0x05, 0x76, 0x85, 0x5a, 0xc9, 0x91, 0x52, 0x1b, 0x65, 0x55,
	0x57, 0x35, 0x7e, 0x7e, 0xf6, 0xfd, 0x74, 0xab, 0xbf, 0xcc, 0xc9, 0xbf, 0xc0, 0x37, 0x72, 0xe4,
	0xf1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x15, 0x21, 0x44, 0xaf, 0x06, 0x00, 0x00,
}