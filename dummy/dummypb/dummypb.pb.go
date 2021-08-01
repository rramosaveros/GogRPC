// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: dummy/dummypb/dummypb.proto

package dummypb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestDummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RequestDummy) Reset() {
	*x = RequestDummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_dummypb_dummypb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDummy) ProtoMessage() {}

func (x *RequestDummy) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_dummypb_dummypb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDummy.ProtoReflect.Descriptor instead.
func (*RequestDummy) Descriptor() ([]byte, []int) {
	return file_dummy_dummypb_dummypb_proto_rawDescGZIP(), []int{0}
}

func (x *RequestDummy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ResponseDummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ResponseDummy) Reset() {
	*x = ResponseDummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_dummypb_dummypb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDummy) ProtoMessage() {}

func (x *ResponseDummy) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_dummypb_dummypb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDummy.ProtoReflect.Descriptor instead.
func (*ResponseDummy) Descriptor() ([]byte, []int) {
	return file_dummy_dummypb_dummypb_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseDummy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_dummy_dummypb_dummypb_proto protoreflect.FileDescriptor

var file_dummy_dummypb_dummypb_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x70, 0x62, 0x2f,
	0x64, 0x75, 0x6d, 0x6d, 0x79, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64,
	0x75, 0x6d, 0x6d, 0x79, 0x22, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x42, 0x0a,
	0x0c, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x1a, 0x14, 0x2e, 0x64, 0x75,
	0x6d, 0x6d, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x42, 0x0f, 0x5a, 0x0d, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dummy_dummypb_dummypb_proto_rawDescOnce sync.Once
	file_dummy_dummypb_dummypb_proto_rawDescData = file_dummy_dummypb_dummypb_proto_rawDesc
)

func file_dummy_dummypb_dummypb_proto_rawDescGZIP() []byte {
	file_dummy_dummypb_dummypb_proto_rawDescOnce.Do(func() {
		file_dummy_dummypb_dummypb_proto_rawDescData = protoimpl.X.CompressGZIP(file_dummy_dummypb_dummypb_proto_rawDescData)
	})
	return file_dummy_dummypb_dummypb_proto_rawDescData
}

var file_dummy_dummypb_dummypb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dummy_dummypb_dummypb_proto_goTypes = []interface{}{
	(*RequestDummy)(nil),  // 0: dummy.RequestDummy
	(*ResponseDummy)(nil), // 1: dummy.ResponseDummy
}
var file_dummy_dummypb_dummypb_proto_depIdxs = []int32{
	0, // 0: dummy.DummyService.Hello:input_type -> dummy.RequestDummy
	1, // 1: dummy.DummyService.Hello:output_type -> dummy.ResponseDummy
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dummy_dummypb_dummypb_proto_init() }
func file_dummy_dummypb_dummypb_proto_init() {
	if File_dummy_dummypb_dummypb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dummy_dummypb_dummypb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDummy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dummy_dummypb_dummypb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDummy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dummy_dummypb_dummypb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dummy_dummypb_dummypb_proto_goTypes,
		DependencyIndexes: file_dummy_dummypb_dummypb_proto_depIdxs,
		MessageInfos:      file_dummy_dummypb_dummypb_proto_msgTypes,
	}.Build()
	File_dummy_dummypb_dummypb_proto = out.File
	file_dummy_dummypb_dummypb_proto_rawDesc = nil
	file_dummy_dummypb_dummypb_proto_goTypes = nil
	file_dummy_dummypb_dummypb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DummyServiceClient is the client API for DummyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DummyServiceClient interface {
	Hello(ctx context.Context, in *RequestDummy, opts ...grpc.CallOption) (*ResponseDummy, error)
}

type dummyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDummyServiceClient(cc grpc.ClientConnInterface) DummyServiceClient {
	return &dummyServiceClient{cc}
}

func (c *dummyServiceClient) Hello(ctx context.Context, in *RequestDummy, opts ...grpc.CallOption) (*ResponseDummy, error) {
	out := new(ResponseDummy)
	err := c.cc.Invoke(ctx, "/dummy.DummyService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DummyServiceServer is the server API for DummyService service.
type DummyServiceServer interface {
	Hello(context.Context, *RequestDummy) (*ResponseDummy, error)
}

// UnimplementedDummyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDummyServiceServer struct {
}

func (*UnimplementedDummyServiceServer) Hello(context.Context, *RequestDummy) (*ResponseDummy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

func RegisterDummyServiceServer(s *grpc.Server, srv DummyServiceServer) {
	s.RegisterService(&_DummyService_serviceDesc, srv)
}

func _DummyService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDummy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dummy.DummyService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServiceServer).Hello(ctx, req.(*RequestDummy))
	}
	return interceptor(ctx, in, info, handler)
}

var _DummyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dummy.DummyService",
	HandlerType: (*DummyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _DummyService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dummy/dummypb/dummypb.proto",
}