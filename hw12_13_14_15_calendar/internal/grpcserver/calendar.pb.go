// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: calendar.proto

package grpcserver

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type Day struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day string `protobuf:"bytes,1,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *Day) Reset() {
	*x = Day{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Day) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Day) ProtoMessage() {}

func (x *Day) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Day.ProtoReflect.Descriptor instead.
func (*Day) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{1}
}

func (x *Day) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Header       string               `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Datetime     *timestamp.Timestamp `protobuf:"bytes,3,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Duration     string               `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Description  string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	UserUuid     string               `protobuf:"bytes,6,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	NotifyBefore string               `protobuf:"bytes,7,opt,name=notify_before,json=notifyBefore,proto3" json:"notify_before,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Event) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *Event) GetDatetime() *timestamp.Timestamp {
	if x != nil {
		return x.Datetime
	}
	return nil
}

func (x *Event) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *Event) GetNotifyBefore() string {
	if x != nil {
		return x.NotifyBefore
	}
	return ""
}

type EventUUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *EventUUID) Reset() {
	*x = EventUUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventUUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventUUID) ProtoMessage() {}

func (x *EventUUID) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventUUID.ProtoReflect.Descriptor instead.
func (*EventUUID) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{3}
}

func (x *EventUUID) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ListEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty"`
}

func (x *ListEvents) Reset() {
	*x = ListEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEvents) ProtoMessage() {}

func (x *ListEvents) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEvents.ProtoReflect.Descriptor instead.
func (*ListEvents) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{4}
}

func (x *ListEvents) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_calendar_proto protoreflect.FileDescriptor

var file_calendar_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x18, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x17, 0x0a, 0x03, 0x44,
	0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x64, 0x61, 0x79, 0x22, 0xeb, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x42, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x22, 0x1f, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x1e, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x32, 0xc8, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x1b,
	0x0a, 0x08, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0a, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1c, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x79, 0x12, 0x04, 0x2e, 0x44, 0x61, 0x79,
	0x1a, 0x0b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a,
	0x08, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x04, 0x2e, 0x44, 0x61, 0x79, 0x1a,
	0x0b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x04, 0x2e, 0x44, 0x61, 0x79, 0x1a,
	0x0b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x4d, 0x5a, 0x4b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x65, 0x6d, 0x61, 0x63,
	0x72, 0x2f, 0x6f, 0x74, 0x75, 0x73, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x68, 0x77,
	0x2f, 0x68, 0x77, 0x31, 0x32, 0x5f, 0x31, 0x33, 0x5f, 0x31, 0x34, 0x5f, 0x31, 0x35, 0x5f, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_calendar_proto_rawDescOnce sync.Once
	file_calendar_proto_rawDescData = file_calendar_proto_rawDesc
)

func file_calendar_proto_rawDescGZIP() []byte {
	file_calendar_proto_rawDescOnce.Do(func() {
		file_calendar_proto_rawDescData = protoimpl.X.CompressGZIP(file_calendar_proto_rawDescData)
	})
	return file_calendar_proto_rawDescData
}

var file_calendar_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_calendar_proto_goTypes = []interface{}{
	(*Result)(nil),              // 0: Result
	(*Day)(nil),                 // 1: Day
	(*Event)(nil),               // 2: Event
	(*EventUUID)(nil),           // 3: EventUUID
	(*ListEvents)(nil),          // 4: ListEvents
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_calendar_proto_depIdxs = []int32{
	5, // 0: Event.datetime:type_name -> google.protobuf.Timestamp
	2, // 1: ListEvents.Events:type_name -> Event
	2, // 2: Calendar.AddEvent:input_type -> Event
	2, // 3: Calendar.ModifyEvent:input_type -> Event
	3, // 4: Calendar.DeleteEvent:input_type -> EventUUID
	1, // 5: Calendar.ListDay:input_type -> Day
	1, // 6: Calendar.ListWeek:input_type -> Day
	1, // 7: Calendar.ListMonth:input_type -> Day
	0, // 8: Calendar.AddEvent:output_type -> Result
	0, // 9: Calendar.ModifyEvent:output_type -> Result
	0, // 10: Calendar.DeleteEvent:output_type -> Result
	4, // 11: Calendar.ListDay:output_type -> ListEvents
	4, // 12: Calendar.ListWeek:output_type -> ListEvents
	4, // 13: Calendar.ListMonth:output_type -> ListEvents
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_calendar_proto_init() }
func file_calendar_proto_init() {
	if File_calendar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calendar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_calendar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Day); i {
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
		file_calendar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_calendar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventUUID); i {
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
		file_calendar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEvents); i {
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
			RawDescriptor: file_calendar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calendar_proto_goTypes,
		DependencyIndexes: file_calendar_proto_depIdxs,
		MessageInfos:      file_calendar_proto_msgTypes,
	}.Build()
	File_calendar_proto = out.File
	file_calendar_proto_rawDesc = nil
	file_calendar_proto_goTypes = nil
	file_calendar_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Result, error)
	ModifyEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Result, error)
	DeleteEvent(ctx context.Context, in *EventUUID, opts ...grpc.CallOption) (*Result, error)
	ListDay(ctx context.Context, in *Day, opts ...grpc.CallOption) (*ListEvents, error)
	ListWeek(ctx context.Context, in *Day, opts ...grpc.CallOption) (*ListEvents, error)
	ListMonth(ctx context.Context, in *Day, opts ...grpc.CallOption) (*ListEvents, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/Calendar/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ModifyEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/Calendar/ModifyEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) DeleteEvent(ctx context.Context, in *EventUUID, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/Calendar/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ListDay(ctx context.Context, in *Day, opts ...grpc.CallOption) (*ListEvents, error) {
	out := new(ListEvents)
	err := c.cc.Invoke(ctx, "/Calendar/ListDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ListWeek(ctx context.Context, in *Day, opts ...grpc.CallOption) (*ListEvents, error) {
	out := new(ListEvents)
	err := c.cc.Invoke(ctx, "/Calendar/ListWeek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ListMonth(ctx context.Context, in *Day, opts ...grpc.CallOption) (*ListEvents, error) {
	out := new(ListEvents)
	err := c.cc.Invoke(ctx, "/Calendar/ListMonth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	AddEvent(context.Context, *Event) (*Result, error)
	ModifyEvent(context.Context, *Event) (*Result, error)
	DeleteEvent(context.Context, *EventUUID) (*Result, error)
	ListDay(context.Context, *Day) (*ListEvents, error)
	ListWeek(context.Context, *Day) (*ListEvents, error)
	ListMonth(context.Context, *Day) (*ListEvents, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) AddEvent(context.Context, *Event) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (*UnimplementedCalendarServer) ModifyEvent(context.Context, *Event) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyEvent not implemented")
}
func (*UnimplementedCalendarServer) DeleteEvent(context.Context, *EventUUID) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (*UnimplementedCalendarServer) ListDay(context.Context, *Day) (*ListEvents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDay not implemented")
}
func (*UnimplementedCalendarServer) ListWeek(context.Context, *Day) (*ListEvents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWeek not implemented")
}
func (*UnimplementedCalendarServer) ListMonth(context.Context, *Day) (*ListEvents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMonth not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).AddEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ModifyEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ModifyEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/ModifyEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ModifyEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventUUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).DeleteEvent(ctx, req.(*EventUUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ListDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Day)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ListDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/ListDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ListDay(ctx, req.(*Day))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ListWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Day)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ListWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/ListWeek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ListWeek(ctx, req.(*Day))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ListMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Day)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ListMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/ListMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ListMonth(ctx, req.(*Day))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _Calendar_AddEvent_Handler,
		},
		{
			MethodName: "ModifyEvent",
			Handler:    _Calendar_ModifyEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Calendar_DeleteEvent_Handler,
		},
		{
			MethodName: "ListDay",
			Handler:    _Calendar_ListDay_Handler,
		},
		{
			MethodName: "ListWeek",
			Handler:    _Calendar_ListWeek_Handler,
		},
		{
			MethodName: "ListMonth",
			Handler:    _Calendar_ListMonth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar.proto",
}