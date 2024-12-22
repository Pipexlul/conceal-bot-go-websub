// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.24.4
// source: websub.proto

package generated

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

type SubscribeTopic int32

const (
	SubscribeTopic_TOPIC_UNSPECIFIED SubscribeTopic = 0
	SubscribeTopic_TOPIC_EARTHQUAKES SubscribeTopic = 1
)

// Enum value maps for SubscribeTopic.
var (
	SubscribeTopic_name = map[int32]string{
		0: "TOPIC_UNSPECIFIED",
		1: "TOPIC_EARTHQUAKES",
	}
	SubscribeTopic_value = map[string]int32{
		"TOPIC_UNSPECIFIED": 0,
		"TOPIC_EARTHQUAKES": 1,
	}
)

func (x SubscribeTopic) Enum() *SubscribeTopic {
	p := new(SubscribeTopic)
	*p = x
	return p
}

func (x SubscribeTopic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubscribeTopic) Descriptor() protoreflect.EnumDescriptor {
	return file_websub_proto_enumTypes[0].Descriptor()
}

func (SubscribeTopic) Type() protoreflect.EnumType {
	return &file_websub_proto_enumTypes[0]
}

func (x SubscribeTopic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubscribeTopic.Descriptor instead.
func (SubscribeTopic) EnumDescriptor() ([]byte, []int) {
	return file_websub_proto_rawDescGZIP(), []int{0}
}

type SubscribeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         SubscribeTopic         `protobuf:"varint,1,opt,name=topic,proto3,enum=websub.SubscribeTopic" json:"topic,omitempty"`
	Callback      string                 `protobuf:"bytes,2,opt,name=callback,proto3" json:"callback,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	mi := &file_websub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_websub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_websub_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetTopic() SubscribeTopic {
	if x != nil {
		return x.Topic
	}
	return SubscribeTopic_TOPIC_UNSPECIFIED
}

func (x *SubscribeRequest) GetCallback() string {
	if x != nil {
		return x.Callback
	}
	return ""
}

type SubscribeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	mi := &file_websub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_websub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_websub_proto_rawDescGZIP(), []int{1}
}

type NotificationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Payload       string                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationRequest) Reset() {
	*x = NotificationRequest{}
	mi := &file_websub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationRequest) ProtoMessage() {}

func (x *NotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_websub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationRequest.ProtoReflect.Descriptor instead.
func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return file_websub_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *NotificationRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type NotificationAck struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SubscriberId  string                 `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	Received      bool                   `protobuf:"varint,2,opt,name=received,proto3" json:"received,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationAck) Reset() {
	*x = NotificationAck{}
	mi := &file_websub_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationAck) ProtoMessage() {}

func (x *NotificationAck) ProtoReflect() protoreflect.Message {
	mi := &file_websub_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationAck.ProtoReflect.Descriptor instead.
func (*NotificationAck) Descriptor() ([]byte, []int) {
	return file_websub_proto_rawDescGZIP(), []int{3}
}

func (x *NotificationAck) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *NotificationAck) GetReceived() bool {
	if x != nil {
		return x.Received
	}
	return false
}

var File_websub_proto protoreflect.FileDescriptor

var file_websub_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x65, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x77, 0x65, 0x62, 0x73, 0x75, 0x62, 0x22, 0x5c, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x77, 0x65, 0x62, 0x73,
	0x75, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x13, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x52, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x63, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x2a, 0x3e, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x5f, 0x45, 0x41, 0x52, 0x54, 0x48, 0x51, 0x55, 0x41, 0x4b,
	0x45, 0x53, 0x10, 0x01, 0x32, 0xe7, 0x01, 0x0a, 0x09, 0x57, 0x65, 0x62, 0x53, 0x75, 0x62, 0x48,
	0x75, 0x62, 0x12, 0x42, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12,
	0x18, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77, 0x65, 0x62, 0x73,
	0x75, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x18, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x13, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1b, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x70,
	0x65, 0x78, 0x6c, 0x75, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x61, 0x6c, 0x2d, 0x62, 0x6f,
	0x74, 0x2d, 0x67, 0x6f, 0x2d, 0x77, 0x65, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_websub_proto_rawDescOnce sync.Once
	file_websub_proto_rawDescData = file_websub_proto_rawDesc
)

func file_websub_proto_rawDescGZIP() []byte {
	file_websub_proto_rawDescOnce.Do(func() {
		file_websub_proto_rawDescData = protoimpl.X.CompressGZIP(file_websub_proto_rawDescData)
	})
	return file_websub_proto_rawDescData
}

var file_websub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_websub_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_websub_proto_goTypes = []any{
	(SubscribeTopic)(0),         // 0: websub.SubscribeTopic
	(*SubscribeRequest)(nil),    // 1: websub.SubscribeRequest
	(*SubscribeResponse)(nil),   // 2: websub.SubscribeResponse
	(*NotificationRequest)(nil), // 3: websub.NotificationRequest
	(*NotificationAck)(nil),     // 4: websub.NotificationAck
}
var file_websub_proto_depIdxs = []int32{
	0, // 0: websub.SubscribeRequest.topic:type_name -> websub.SubscribeTopic
	1, // 1: websub.WebSubHub.Subscribe:input_type -> websub.SubscribeRequest
	1, // 2: websub.WebSubHub.Unsubscribe:input_type -> websub.SubscribeRequest
	3, // 3: websub.WebSubHub.StreamNotifications:input_type -> websub.NotificationRequest
	2, // 4: websub.WebSubHub.Subscribe:output_type -> websub.SubscribeResponse
	1, // 5: websub.WebSubHub.Unsubscribe:output_type -> websub.SubscribeRequest
	4, // 6: websub.WebSubHub.StreamNotifications:output_type -> websub.NotificationAck
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_websub_proto_init() }
func file_websub_proto_init() {
	if File_websub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_websub_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_websub_proto_goTypes,
		DependencyIndexes: file_websub_proto_depIdxs,
		EnumInfos:         file_websub_proto_enumTypes,
		MessageInfos:      file_websub_proto_msgTypes,
	}.Build()
	File_websub_proto = out.File
	file_websub_proto_rawDesc = nil
	file_websub_proto_goTypes = nil
	file_websub_proto_depIdxs = nil
}