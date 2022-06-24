// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: protofiles/serverHandlingService.proto

package serverHandling

import (
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

type SERVER_TYPE int32

const (
	SERVER_TYPE_REALTIME    SERVER_TYPE = 0
	SERVER_TYPE_INTERACTION SERVER_TYPE = 1
	SERVER_TYPE_DATA        SERVER_TYPE = 2
)

// Enum value maps for SERVER_TYPE.
var (
	SERVER_TYPE_name = map[int32]string{
		0: "REALTIME",
		1: "INTERACTION",
		2: "DATA",
	}
	SERVER_TYPE_value = map[string]int32{
		"REALTIME":    0,
		"INTERACTION": 1,
		"DATA":        2,
	}
)

func (x SERVER_TYPE) Enum() *SERVER_TYPE {
	p := new(SERVER_TYPE)
	*p = x
	return p
}

func (x SERVER_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SERVER_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_protofiles_serverHandlingService_proto_enumTypes[0].Descriptor()
}

func (SERVER_TYPE) Type() protoreflect.EnumType {
	return &file_protofiles_serverHandlingService_proto_enumTypes[0]
}

func (x SERVER_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SERVER_TYPE.Descriptor instead.
func (SERVER_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_protofiles_serverHandlingService_proto_rawDescGZIP(), []int{0}
}

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       SERVER_TYPE `protobuf:"varint,1,opt,name=Type,proto3,enum=halalwedd_managing_server.SERVER_TYPE" json:"Type,omitempty"`
	Public_IP  string      `protobuf:"bytes,2,opt,name=Public_IP,json=PublicIP,proto3" json:"Public_IP,omitempty"`
	Private_IP string      `protobuf:"bytes,3,opt,name=Private_IP,json=PrivateIP,proto3" json:"Private_IP,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_serverHandlingService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_serverHandlingService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_protofiles_serverHandlingService_proto_rawDescGZIP(), []int{0}
}

func (x *ServerInfo) GetType() SERVER_TYPE {
	if x != nil {
		return x.Type
	}
	return SERVER_TYPE_REALTIME
}

func (x *ServerInfo) GetPublic_IP() string {
	if x != nil {
		return x.Public_IP
	}
	return ""
}

func (x *ServerInfo) GetPrivate_IP() string {
	if x != nil {
		return x.Private_IP
	}
	return ""
}

type StoredServerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InternalID int32  `protobuf:"varint,1,opt,name=InternalID,proto3" json:"InternalID,omitempty"`
	Public_IP  string `protobuf:"bytes,2,opt,name=Public_IP,json=PublicIP,proto3" json:"Public_IP,omitempty"`
	Private_IP string `protobuf:"bytes,3,opt,name=Private_IP,json=PrivateIP,proto3" json:"Private_IP,omitempty"`
}

func (x *StoredServerData) Reset() {
	*x = StoredServerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_serverHandlingService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredServerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredServerData) ProtoMessage() {}

func (x *StoredServerData) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_serverHandlingService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredServerData.ProtoReflect.Descriptor instead.
func (*StoredServerData) Descriptor() ([]byte, []int) {
	return file_protofiles_serverHandlingService_proto_rawDescGZIP(), []int{1}
}

func (x *StoredServerData) GetInternalID() int32 {
	if x != nil {
		return x.InternalID
	}
	return 0
}

func (x *StoredServerData) GetPublic_IP() string {
	if x != nil {
		return x.Public_IP
	}
	return ""
}

func (x *StoredServerData) GetPrivate_IP() string {
	if x != nil {
		return x.Private_IP
	}
	return ""
}

type ServerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string              `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	InternalID int32               `protobuf:"varint,2,opt,name=InternalID,proto3" json:"InternalID,omitempty"`
	ServerData []*StoredServerData `protobuf:"bytes,3,rep,name=ServerData,proto3" json:"ServerData,omitempty"`
}

func (x *ServerInfoResponse) Reset() {
	*x = ServerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_serverHandlingService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfoResponse) ProtoMessage() {}

func (x *ServerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_serverHandlingService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfoResponse.ProtoReflect.Descriptor instead.
func (*ServerInfoResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_serverHandlingService_proto_rawDescGZIP(), []int{2}
}

func (x *ServerInfoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ServerInfoResponse) GetInternalID() int32 {
	if x != nil {
		return x.InternalID
	}
	return 0
}

func (x *ServerInfoResponse) GetServerData() []*StoredServerData {
	if x != nil {
		return x.ServerData
	}
	return nil
}

var File_protofiles_serverHandlingService_proto protoreflect.FileDescriptor

var file_protofiles_serverHandlingService_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x68, 0x61, 0x6c, 0x61, 0x6c, 0x77,
	0x65, 0x64, 0x64, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x26, 0x2e, 0x68, 0x61, 0x6c, 0x61, 0x6c, 0x77, 0x65, 0x64, 0x64, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x12, 0x1d, 0x0a, 0x0a, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x50, 0x22, 0x6e, 0x0a, 0x10, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e,
	0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x1b,
	0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x12, 0x1d, 0x0a, 0x0a, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x50, 0x22, 0x9b, 0x01, 0x0a, 0x12, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x4b, 0x0a, 0x0a, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x68, 0x61, 0x6c, 0x61, 0x6c, 0x77, 0x65, 0x64, 0x64, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2a, 0x36, 0x0a, 0x0b, 0x53, 0x45, 0x52, 0x56,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x41, 0x4c, 0x54,
	0x49, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x41, 0x10, 0x02,
	0x32, 0x77, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69,
	0x6e, 0x67, 0x12, 0x65, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x68, 0x61, 0x6c, 0x61, 0x6c, 0x77, 0x65, 0x64, 0x64, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x2d, 0x2e, 0x68, 0x61, 0x6c,
	0x61, 0x6c, 0x77, 0x65, 0x64, 0x64, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofiles_serverHandlingService_proto_rawDescOnce sync.Once
	file_protofiles_serverHandlingService_proto_rawDescData = file_protofiles_serverHandlingService_proto_rawDesc
)

func file_protofiles_serverHandlingService_proto_rawDescGZIP() []byte {
	file_protofiles_serverHandlingService_proto_rawDescOnce.Do(func() {
		file_protofiles_serverHandlingService_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofiles_serverHandlingService_proto_rawDescData)
	})
	return file_protofiles_serverHandlingService_proto_rawDescData
}

var file_protofiles_serverHandlingService_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protofiles_serverHandlingService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protofiles_serverHandlingService_proto_goTypes = []interface{}{
	(SERVER_TYPE)(0),           // 0: halalwedd_managing_server.SERVER_TYPE
	(*ServerInfo)(nil),         // 1: halalwedd_managing_server.ServerInfo
	(*StoredServerData)(nil),   // 2: halalwedd_managing_server.StoredServerData
	(*ServerInfoResponse)(nil), // 3: halalwedd_managing_server.ServerInfoResponse
}
var file_protofiles_serverHandlingService_proto_depIdxs = []int32{
	0, // 0: halalwedd_managing_server.ServerInfo.Type:type_name -> halalwedd_managing_server.SERVER_TYPE
	2, // 1: halalwedd_managing_server.ServerInfoResponse.ServerData:type_name -> halalwedd_managing_server.StoredServerData
	1, // 2: halalwedd_managing_server.ServerHandling.publishServer:input_type -> halalwedd_managing_server.ServerInfo
	3, // 3: halalwedd_managing_server.ServerHandling.publishServer:output_type -> halalwedd_managing_server.ServerInfoResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protofiles_serverHandlingService_proto_init() }
func file_protofiles_serverHandlingService_proto_init() {
	if File_protofiles_serverHandlingService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofiles_serverHandlingService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
		file_protofiles_serverHandlingService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredServerData); i {
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
		file_protofiles_serverHandlingService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfoResponse); i {
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
			RawDescriptor: file_protofiles_serverHandlingService_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protofiles_serverHandlingService_proto_goTypes,
		DependencyIndexes: file_protofiles_serverHandlingService_proto_depIdxs,
		EnumInfos:         file_protofiles_serverHandlingService_proto_enumTypes,
		MessageInfos:      file_protofiles_serverHandlingService_proto_msgTypes,
	}.Build()
	File_protofiles_serverHandlingService_proto = out.File
	file_protofiles_serverHandlingService_proto_rawDesc = nil
	file_protofiles_serverHandlingService_proto_goTypes = nil
	file_protofiles_serverHandlingService_proto_depIdxs = nil
}