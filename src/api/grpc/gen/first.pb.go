// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: first.proto

package timergrpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CallbackType int32

const (
	CallbackType_HTTP     CallbackType = 0
	CallbackType_RABBITMQ CallbackType = 1
)

// Enum value maps for CallbackType.
var (
	CallbackType_name = map[int32]string{
		0: "HTTP",
		1: "RABBITMQ",
	}
	CallbackType_value = map[string]int32{
		"HTTP":     0,
		"RABBITMQ": 1,
	}
)

func (x CallbackType) Enum() *CallbackType {
	p := new(CallbackType)
	*p = x
	return p
}

func (x CallbackType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CallbackType) Descriptor() protoreflect.EnumDescriptor {
	return file_first_proto_enumTypes[0].Descriptor()
}

func (CallbackType) Type() protoreflect.EnumType {
	return &file_first_proto_enumTypes[0]
}

func (x CallbackType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CallbackType.Descriptor instead.
func (CallbackType) EnumDescriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{0}
}

type TimerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpiredAt        *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	Payload          *structpb.Struct       `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	CallbackConfigId string                 `protobuf:"bytes,3,opt,name=callback_config_id,json=callbackConfigId,proto3" json:"callback_config_id,omitempty"`
}

func (x *TimerRequest) Reset() {
	*x = TimerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerRequest) ProtoMessage() {}

func (x *TimerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerRequest.ProtoReflect.Descriptor instead.
func (*TimerRequest) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{0}
}

func (x *TimerRequest) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

func (x *TimerRequest) GetPayload() *structpb.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *TimerRequest) GetCallbackConfigId() string {
	if x != nil {
		return x.CallbackConfigId
	}
	return ""
}

type TimerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExpiredAt      *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	CreatedAt      *timestamppb.Timestamp  `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	SendAt         *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=send_at,json=sendAt,proto3" json:"send_at,omitempty"`
	Status         string                  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	RetryCount     int32                   `protobuf:"varint,6,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	ErrorMessage   string                  `protobuf:"bytes,7,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	LastRetried    *timestamppb.Timestamp  `protobuf:"bytes,8,opt,name=last_retried,json=lastRetried,proto3" json:"last_retried,omitempty"`
	Payload        *structpb.Struct        `protobuf:"bytes,9,opt,name=payload,proto3" json:"payload,omitempty"`
	CallbackConfig *CallbackConfigResponse `protobuf:"bytes,10,opt,name=callback_config,json=callbackConfig,proto3" json:"callback_config,omitempty"`
	Url            string                  `protobuf:"bytes,11,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *TimerResponse) Reset() {
	*x = TimerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerResponse) ProtoMessage() {}

func (x *TimerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerResponse.ProtoReflect.Descriptor instead.
func (*TimerResponse) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{1}
}

func (x *TimerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TimerResponse) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

func (x *TimerResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TimerResponse) GetSendAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SendAt
	}
	return nil
}

func (x *TimerResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TimerResponse) GetRetryCount() int32 {
	if x != nil {
		return x.RetryCount
	}
	return 0
}

func (x *TimerResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *TimerResponse) GetLastRetried() *timestamppb.Timestamp {
	if x != nil {
		return x.LastRetried
	}
	return nil
}

func (x *TimerResponse) GetPayload() *structpb.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *TimerResponse) GetCallbackConfig() *CallbackConfigResponse {
	if x != nil {
		return x.CallbackConfig
	}
	return nil
}

func (x *TimerResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CallbackConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    *string                     `protobuf:"bytes,1,opt,name=url,proto3,oneof" json:"url,omitempty"`
	Type   CallbackType                `protobuf:"varint,2,opt,name=type,proto3,enum=api.CallbackType" json:"type,omitempty"`
	Params map[string]*structpb.Struct `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CallbackConfigRequest) Reset() {
	*x = CallbackConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackConfigRequest) ProtoMessage() {}

func (x *CallbackConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackConfigRequest.ProtoReflect.Descriptor instead.
func (*CallbackConfigRequest) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{2}
}

func (x *CallbackConfigRequest) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *CallbackConfigRequest) GetType() CallbackType {
	if x != nil {
		return x.Type
	}
	return CallbackType_HTTP
}

func (x *CallbackConfigRequest) GetParams() map[string]*structpb.Struct {
	if x != nil {
		return x.Params
	}
	return nil
}

type CallbackConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type   CallbackType                `protobuf:"varint,2,opt,name=type,proto3,enum=api.CallbackType" json:"type,omitempty"`
	Url    string                      `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Params map[string]*structpb.Struct `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CallbackConfigResponse) Reset() {
	*x = CallbackConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackConfigResponse) ProtoMessage() {}

func (x *CallbackConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackConfigResponse.ProtoReflect.Descriptor instead.
func (*CallbackConfigResponse) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{3}
}

func (x *CallbackConfigResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CallbackConfigResponse) GetType() CallbackType {
	if x != nil {
		return x.Type
	}
	return CallbackType_HTTP
}

func (x *CallbackConfigResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CallbackConfigResponse) GetParams() map[string]*structpb.Struct {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_first_proto protoreflect.FileDescriptor

var file_first_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x22, 0xf2,
	0x03, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x44, 0x0a, 0x0f, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0xf1, 0x01, 0x0a, 0x15, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x52, 0x0a, 0x0b, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0xf6, 0x01, 0x0a, 0x16, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x3f, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x52, 0x0a, 0x0b,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x2a, 0x26, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x41,
	0x42, 0x42, 0x49, 0x54, 0x4d, 0x51, 0x10, 0x01, 0x32, 0x89, 0x01, 0x0a, 0x08, 0x41, 0x50, 0x49,
	0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x05, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x61, 0x6c, 0x65, 0x78, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x3b, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_first_proto_rawDescOnce sync.Once
	file_first_proto_rawDescData = file_first_proto_rawDesc
)

func file_first_proto_rawDescGZIP() []byte {
	file_first_proto_rawDescOnce.Do(func() {
		file_first_proto_rawDescData = protoimpl.X.CompressGZIP(file_first_proto_rawDescData)
	})
	return file_first_proto_rawDescData
}

var file_first_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_first_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_first_proto_goTypes = []interface{}{
	(CallbackType)(0),              // 0: api.CallbackType
	(*TimerRequest)(nil),           // 1: api.TimerRequest
	(*TimerResponse)(nil),          // 2: api.TimerResponse
	(*CallbackConfigRequest)(nil),  // 3: api.CallbackConfigRequest
	(*CallbackConfigResponse)(nil), // 4: api.CallbackConfigResponse
	nil,                            // 5: api.CallbackConfigRequest.ParamsEntry
	nil,                            // 6: api.CallbackConfigResponse.ParamsEntry
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*structpb.Struct)(nil),        // 8: google.protobuf.Struct
}
var file_first_proto_depIdxs = []int32{
	7,  // 0: api.TimerRequest.expired_at:type_name -> google.protobuf.Timestamp
	8,  // 1: api.TimerRequest.payload:type_name -> google.protobuf.Struct
	7,  // 2: api.TimerResponse.expired_at:type_name -> google.protobuf.Timestamp
	7,  // 3: api.TimerResponse.created_at:type_name -> google.protobuf.Timestamp
	7,  // 4: api.TimerResponse.send_at:type_name -> google.protobuf.Timestamp
	7,  // 5: api.TimerResponse.last_retried:type_name -> google.protobuf.Timestamp
	8,  // 6: api.TimerResponse.payload:type_name -> google.protobuf.Struct
	4,  // 7: api.TimerResponse.callback_config:type_name -> api.CallbackConfigResponse
	0,  // 8: api.CallbackConfigRequest.type:type_name -> api.CallbackType
	5,  // 9: api.CallbackConfigRequest.params:type_name -> api.CallbackConfigRequest.ParamsEntry
	0,  // 10: api.CallbackConfigResponse.type:type_name -> api.CallbackType
	6,  // 11: api.CallbackConfigResponse.params:type_name -> api.CallbackConfigResponse.ParamsEntry
	8,  // 12: api.CallbackConfigRequest.ParamsEntry.value:type_name -> google.protobuf.Struct
	8,  // 13: api.CallbackConfigResponse.ParamsEntry.value:type_name -> google.protobuf.Struct
	1,  // 14: api.APITimer.Timer:input_type -> api.TimerRequest
	3,  // 15: api.APITimer.CallbackConfig:input_type -> api.CallbackConfigRequest
	2,  // 16: api.APITimer.Timer:output_type -> api.TimerResponse
	4,  // 17: api.APITimer.CallbackConfig:output_type -> api.CallbackConfigResponse
	16, // [16:18] is the sub-list for method output_type
	14, // [14:16] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_first_proto_init() }
func file_first_proto_init() {
	if File_first_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_first_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerRequest); i {
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
		file_first_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerResponse); i {
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
		file_first_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackConfigRequest); i {
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
		file_first_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackConfigResponse); i {
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
	file_first_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_first_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_first_proto_goTypes,
		DependencyIndexes: file_first_proto_depIdxs,
		EnumInfos:         file_first_proto_enumTypes,
		MessageInfos:      file_first_proto_msgTypes,
	}.Build()
	File_first_proto = out.File
	file_first_proto_rawDesc = nil
	file_first_proto_goTypes = nil
	file_first_proto_depIdxs = nil
}
