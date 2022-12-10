// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: chat.proto

package pb

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

type MessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	RoomIdx uint32 `protobuf:"varint,2,opt,name=room_idx,json=roomIdx,proto3" json:"room_idx,omitempty"`
	UserIdx uint32 `protobuf:"varint,3,opt,name=user_idx,json=userIdx,proto3" json:"user_idx,omitempty"`
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageReq) Reset() {
	*x = MessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReq) ProtoMessage() {}

func (x *MessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReq.ProtoReflect.Descriptor instead.
func (*MessageReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *MessageReq) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MessageReq) GetRoomIdx() uint32 {
	if x != nil {
		return x.RoomIdx
	}
	return 0
}

func (x *MessageReq) GetUserIdx() uint32 {
	if x != nil {
		return x.UserIdx
	}
	return 0
}

func (x *MessageReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MessageRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	RoomIdx  uint32 `protobuf:"varint,2,opt,name=room_idx,json=roomIdx,proto3" json:"room_idx,omitempty"`
	UserIdx  uint32 `protobuf:"varint,3,opt,name=user_idx,json=userIdx,proto3" json:"user_idx,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Message  string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageRes) Reset() {
	*x = MessageRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRes) ProtoMessage() {}

func (x *MessageRes) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRes.ProtoReflect.Descriptor instead.
func (*MessageRes) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *MessageRes) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MessageRes) GetRoomIdx() uint32 {
	if x != nil {
		return x.RoomIdx
	}
	return 0
}

func (x *MessageRes) GetUserIdx() uint32 {
	if x != nil {
		return x.UserIdx
	}
	return 0
}

func (x *MessageRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MessageRes) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *MessageRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x70, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x78, 0x12, 0x19, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x78,
	0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x34, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x2c,
	0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chat_proto_goTypes = []interface{}{
	(*MessageReq)(nil), // 0: pb.MessageReq
	(*MessageRes)(nil), // 1: pb.MessageRes
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: pb.Chat.chat:input_type -> pb.MessageReq
	1, // 1: pb.Chat.chat:output_type -> pb.MessageRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReq); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRes); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
