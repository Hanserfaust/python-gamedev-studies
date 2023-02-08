// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: messages.proto

package gateserver

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

type MType int32

const (
	MType_UNKNOWN      MType = 0
	MType_PLAYER_LOGIN MType = 1
	MType_MOUSE_EVENT  MType = 2
	MType_PING_EVENT   MType = 3
)

// Enum value maps for MType.
var (
	MType_name = map[int32]string{
		0: "UNKNOWN",
		1: "PLAYER_LOGIN",
		2: "MOUSE_EVENT",
		3: "PING_EVENT",
	}
	MType_value = map[string]int32{
		"UNKNOWN":      0,
		"PLAYER_LOGIN": 1,
		"MOUSE_EVENT":  2,
		"PING_EVENT":   3,
	}
)

func (x MType) Enum() *MType {
	p := new(MType)
	*p = x
	return p
}

func (x MType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[0].Descriptor()
}

func (MType) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[0]
}

func (x MType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MType.Descriptor instead.
func (MType) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

type PlayerLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *PlayerLogin) Reset() {
	*x = PlayerLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLogin) ProtoMessage() {}

func (x *PlayerLogin) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLogin.ProtoReflect.Descriptor instead.
func (*PlayerLogin) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerLogin) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PlayerLogin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type MouseEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X          int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y          int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	LeftClick  bool  `protobuf:"varint,3,opt,name=left_click,json=leftClick,proto3" json:"left_click,omitempty"`
	RightClick bool  `protobuf:"varint,4,opt,name=right_click,json=rightClick,proto3" json:"right_click,omitempty"`
}

func (x *MouseEvent) Reset() {
	*x = MouseEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MouseEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MouseEvent) ProtoMessage() {}

func (x *MouseEvent) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MouseEvent.ProtoReflect.Descriptor instead.
func (*MouseEvent) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *MouseEvent) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MouseEvent) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *MouseEvent) GetLeftClick() bool {
	if x != nil {
		return x.LeftClick
	}
	return false
}

func (x *MouseEvent) GetRightClick() bool {
	if x != nil {
		return x.RightClick
	}
	return false
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SentEpoch     uint64 `protobuf:"varint,1,opt,name=sentEpoch,proto3" json:"sentEpoch,omitempty"`
	BounceEpoch   uint64 `protobuf:"varint,2,opt,name=bounceEpoch,proto3" json:"bounceEpoch,omitempty"`
	ReceivedEpoch uint64 `protobuf:"varint,3,opt,name=receivedEpoch,proto3" json:"receivedEpoch,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Ping) GetSentEpoch() uint64 {
	if x != nil {
		return x.SentEpoch
	}
	return 0
}

func (x *Ping) GetBounceEpoch() uint64 {
	if x != nil {
		return x.BounceEpoch
	}
	return 0
}

func (x *Ping) GetReceivedEpoch() uint64 {
	if x != nil {
		return x.ReceivedEpoch
	}
	return 0
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x0b, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x68, 0x0a, 0x0a, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x65, 0x66, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x6c, 0x65, 0x66, 0x74, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x72, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x22, 0x6c, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x45, 0x70, 0x6f, 0x63, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x2a, 0x47, 0x0a, 0x05, 0x4d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x4f, 0x55, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x10, 0x03, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_messages_proto_goTypes = []interface{}{
	(MType)(0),          // 0: messages.MType
	(*PlayerLogin)(nil), // 1: messages.PlayerLogin
	(*MouseEvent)(nil),  // 2: messages.MouseEvent
	(*Ping)(nil),        // 3: messages.Ping
}
var file_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLogin); i {
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
		file_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MouseEvent); i {
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
		file_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		EnumInfos:         file_messages_proto_enumTypes,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
