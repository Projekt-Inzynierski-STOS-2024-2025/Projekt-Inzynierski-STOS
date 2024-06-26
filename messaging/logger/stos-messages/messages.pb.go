// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: messages.proto

package stos_messages

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

//
// Errors
//
type ErrorType int32

const (
	ErrorType_FILE_UPLOAD_ERROR      ErrorType = 0
	ErrorType_FILE_DOWNLOAD_ERROR    ErrorType = 1
	ErrorType_WORKER_OPERATION_ERROR ErrorType = 2
	ErrorType_EVALUATOR_ERROR        ErrorType = 3
)

// Enum value maps for ErrorType.
var (
	ErrorType_name = map[int32]string{
		0: "FILE_UPLOAD_ERROR",
		1: "FILE_DOWNLOAD_ERROR",
		2: "WORKER_OPERATION_ERROR",
		3: "EVALUATOR_ERROR",
	}
	ErrorType_value = map[string]int32{
		"FILE_UPLOAD_ERROR":      0,
		"FILE_DOWNLOAD_ERROR":    1,
		"WORKER_OPERATION_ERROR": 2,
		"EVALUATOR_ERROR":        3,
	}
)

func (x ErrorType) Enum() *ErrorType {
	p := new(ErrorType)
	*p = x
	return p
}

func (x ErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[0].Descriptor()
}

func (ErrorType) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[0]
}

func (x ErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorType.Descriptor instead.
func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

type LogType int32

const (
	LogType_INFO    LogType = 0
	LogType_ERROR   LogType = 1
	LogType_WARNING LogType = 2
)

// Enum value maps for LogType.
var (
	LogType_name = map[int32]string{
		0: "INFO",
		1: "ERROR",
		2: "WARNING",
	}
	LogType_value = map[string]int32{
		"INFO":    0,
		"ERROR":   1,
		"WARNING": 2,
	}
)

func (x LogType) Enum() *LogType {
	p := new(LogType)
	*p = x
	return p
}

func (x LogType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[1].Descriptor()
}

func (LogType) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[1]
}

func (x LogType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogType.Descriptor instead.
func (LogType) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

//
// File server stuff
//
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Error logging
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time         string    `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Error        ErrorType `protobuf:"varint,2,opt,name=error,proto3,enum=stos.messages.ErrorType" json:"error,omitempty"`
	ErrorMessage string    `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Error) GetError() ErrorType {
	if x != nil {
		return x.Error
	}
	return ErrorType_FILE_UPLOAD_ERROR
}

func (x *Error) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// Main event exchange message
type TaskDispatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string  `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	StudentId string  `protobuf:"bytes,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	FilesHash []byte  `protobuf:"bytes,3,opt,name=files_hash,json=filesHash,proto3" json:"files_hash,omitempty"`
	Data      []*File `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *TaskDispatch) Reset() {
	*x = TaskDispatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDispatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDispatch) ProtoMessage() {}

func (x *TaskDispatch) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TaskDispatch.ProtoReflect.Descriptor instead.
func (*TaskDispatch) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *TaskDispatch) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskDispatch) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *TaskDispatch) GetFilesHash() []byte {
	if x != nil {
		return x.FilesHash
	}
	return nil
}

func (x *TaskDispatch) GetData() []*File {
	if x != nil {
		return x.Data
	}
	return nil
}

type LogEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time    string  `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Type    LogType `protobuf:"varint,2,opt,name=type,proto3,enum=stos.messages.LogType" json:"type,omitempty"`
	Content string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *LogEvent) Reset() {
	*x = LogEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEvent) ProtoMessage() {}

func (x *LogEvent) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEvent.ProtoReflect.Descriptor instead.
func (*LogEvent) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

func (x *LogEvent) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *LogEvent) GetType() LogType {
	if x != nil {
		return x.Type
	}
	return LogType_INFO
}

func (x *LogEvent) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// File transfer message
type Files struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []*File `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
	Uuid    string  `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Files) Reset() {
	*x = Files{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Files) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Files) ProtoMessage() {}

func (x *Files) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Files.ProtoReflect.Descriptor instead.
func (*Files) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{4}
}

func (x *Files) GetContent() []*File {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Files) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x73, 0x74, 0x6f, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x2e, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x70, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x74,
	0x6f, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x8e, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x73, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x64, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x2a, 0x6c, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x4c, 0x45,
	0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x01, 0x12, 0x1a, 0x0a, 0x16, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x56, 0x41, 0x4c, 0x55, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x03, 0x2a, 0x2b, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x42,
	0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74, 0x6f, 0x73, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_messages_proto_goTypes = []interface{}{
	(ErrorType)(0),       // 0: stos.messages.ErrorType
	(LogType)(0),         // 1: stos.messages.LogType
	(*File)(nil),         // 2: stos.messages.File
	(*Error)(nil),        // 3: stos.messages.Error
	(*TaskDispatch)(nil), // 4: stos.messages.TaskDispatch
	(*LogEvent)(nil),     // 5: stos.messages.LogEvent
	(*Files)(nil),        // 6: stos.messages.Files
}
var file_messages_proto_depIdxs = []int32{
	0, // 0: stos.messages.Error.error:type_name -> stos.messages.ErrorType
	2, // 1: stos.messages.TaskDispatch.data:type_name -> stos.messages.File
	1, // 2: stos.messages.LogEvent.type:type_name -> stos.messages.LogType
	2, // 3: stos.messages.Files.content:type_name -> stos.messages.File
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			switch v := v.(*Error); i {
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
			switch v := v.(*TaskDispatch); i {
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
		file_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEvent); i {
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
		file_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Files); i {
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
			NumEnums:      2,
			NumMessages:   5,
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
