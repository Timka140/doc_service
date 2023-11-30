// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: pkg/transport/protocol/protocol.proto

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

type CreateService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid  string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"` // id сервиса
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Pack []byte `protobuf:"bytes,3,opt,name=pack,proto3" json:"pack,omitempty"`
}

func (x *CreateService) Reset() {
	*x = CreateService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateService) ProtoMessage() {}

func (x *CreateService) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateService.ProtoReflect.Descriptor instead.
func (*CreateService) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *CreateService) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *CreateService) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateService) GetPack() []byte {
	if x != nil {
		return x.Pack
	}
	return nil
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid  string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"` // id сервиса
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Pack []byte `protobuf:"bytes,3,opt,name=pack,proto3" json:"pack,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *Info) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *Info) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Info) GetPack() []byte {
	if x != nil {
		return x.Pack
	}
	return nil
}

type ReportFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string        `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"` // Тип отчета
	Info []*ReportInfo `protobuf:"bytes,4,rep,name=info,proto3" json:"info,omitempty"` // Описание посылки
	Pack []byte        `protobuf:"bytes,5,opt,name=pack,proto3" json:"pack,omitempty"` // Готовые данные в бинарном виде
}

func (x *ReportFormat) Reset() {
	*x = ReportFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportFormat) ProtoMessage() {}

func (x *ReportFormat) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportFormat.ProtoReflect.Descriptor instead.
func (*ReportFormat) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *ReportFormat) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ReportFormat) GetInfo() []*ReportInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ReportFormat) GetPack() []byte {
	if x != nil {
		return x.Pack
	}
	return nil
}

type ReportInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size  int32  `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`  // Размер файла
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"` // Описание ошибки формирования отчета
}

func (x *ReportInfo) Reset() {
	*x = ReportInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportInfo) ProtoMessage() {}

func (x *ReportInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportInfo.ProtoReflect.Descriptor instead.
func (*ReportInfo) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{3}
}

func (x *ReportInfo) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ReportInfo) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrvAdr *ReportFormat `protobuf:"bytes,1,opt,name=SrvAdr,proto3" json:"SrvAdr,omitempty"`
}

func (x *ReportReq) Reset() {
	*x = ReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportReq) ProtoMessage() {}

func (x *ReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportReq.ProtoReflect.Descriptor instead.
func (*ReportReq) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{4}
}

func (x *ReportReq) GetSrvAdr() *ReportFormat {
	if x != nil {
		return x.SrvAdr
	}
	return nil
}

type ReportResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrvAdr *ReportFormat `protobuf:"bytes,1,opt,name=SrvAdr,proto3" json:"SrvAdr,omitempty"`
}

func (x *ReportResp) Reset() {
	*x = ReportResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResp) ProtoMessage() {}

func (x *ReportResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResp.ProtoReflect.Descriptor instead.
func (*ReportResp) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{5}
}

func (x *ReportResp) GetSrvAdr() *ReportFormat {
	if x != nil {
		return x.SrvAdr
	}
	return nil
}

// Ping
type ServerPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid  string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`   // id сервиса
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"` // Тип опроса
	Tm   int64  `protobuf:"varint,3,opt,name=tm,proto3" json:"tm,omitempty"`    // Время задержки
}

func (x *ServerPing) Reset() {
	*x = ServerPing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerPing) ProtoMessage() {}

func (x *ServerPing) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerPing.ProtoReflect.Descriptor instead.
func (*ServerPing) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{6}
}

func (x *ServerPing) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *ServerPing) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ServerPing) GetTm() int64 {
	if x != nil {
		return x.Tm
	}
	return 0
}

type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrvPing *ServerPing `protobuf:"bytes,1,opt,name=SrvPing,proto3" json:"SrvPing,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{7}
}

func (x *PingReq) GetSrvPing() *ServerPing {
	if x != nil {
		return x.SrvPing
	}
	return nil
}

type PingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrvPing *ServerPing `protobuf:"bytes,1,opt,name=SrvPing,proto3" json:"SrvPing,omitempty"`
}

func (x *PingResp) Reset() {
	*x = PingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResp) ProtoMessage() {}

func (x *PingResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResp.ProtoReflect.Descriptor instead.
func (*PingResp) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{8}
}

func (x *PingResp) GetSrvPing() *ServerPing {
	if x != nil {
		return x.SrvPing
	}
	return nil
}

// Service
type CreateSrvReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Srv *CreateService `protobuf:"bytes,1,opt,name=Srv,proto3" json:"Srv,omitempty"`
}

func (x *CreateSrvReq) Reset() {
	*x = CreateSrvReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSrvReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSrvReq) ProtoMessage() {}

func (x *CreateSrvReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSrvReq.ProtoReflect.Descriptor instead.
func (*CreateSrvReq) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{9}
}

func (x *CreateSrvReq) GetSrv() *CreateService {
	if x != nil {
		return x.Srv
	}
	return nil
}

type CreateSrvResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Srv *CreateService `protobuf:"bytes,1,opt,name=Srv,proto3" json:"Srv,omitempty"`
}

func (x *CreateSrvResp) Reset() {
	*x = CreateSrvResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSrvResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSrvResp) ProtoMessage() {}

func (x *CreateSrvResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSrvResp.ProtoReflect.Descriptor instead.
func (*CreateSrvResp) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{10}
}

func (x *CreateSrvResp) GetSrv() *CreateService {
	if x != nil {
		return x.Srv
	}
	return nil
}

// Info
type InfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrvInfo *Info `protobuf:"bytes,1,opt,name=SrvInfo,proto3" json:"SrvInfo,omitempty"`
}

func (x *InfoReq) Reset() {
	*x = InfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReq) ProtoMessage() {}

func (x *InfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoReq.ProtoReflect.Descriptor instead.
func (*InfoReq) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{11}
}

func (x *InfoReq) GetSrvInfo() *Info {
	if x != nil {
		return x.SrvInfo
	}
	return nil
}

type InfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrvInfo *Info `protobuf:"bytes,1,opt,name=SrvInfo,proto3" json:"SrvInfo,omitempty"`
}

func (x *InfoResp) Reset() {
	*x = InfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResp) ProtoMessage() {}

func (x *InfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transport_protocol_protocol_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResp.ProtoReflect.Descriptor instead.
func (*InfoResp) Descriptor() ([]byte, []int) {
	return file_pkg_transport_protocol_protocol_proto_rawDescGZIP(), []int{12}
}

func (x *InfoResp) GetSrvInfo() *Info {
	if x != nil {
		return x.SrvInfo
	}
	return nil
}

var File_pkg_transport_protocol_protocol_proto protoreflect.FileDescriptor

var file_pkg_transport_protocol_protocol_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x49, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x22, 0x40, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x22, 0x5a, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x70, 0x61, 0x63, 0x6b, 0x22, 0x36, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x09,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x06, 0x53, 0x72, 0x76,
	0x41, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x53, 0x72, 0x76,
	0x41, 0x64, 0x72, 0x22, 0x36, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x28, 0x0a, 0x06, 0x53, 0x72, 0x76, 0x41, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x52, 0x06, 0x53, 0x72, 0x76, 0x41, 0x64, 0x72, 0x22, 0x42, 0x0a, 0x0a, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x6d, 0x22,
	0x33, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x07, 0x53, 0x72,
	0x76, 0x50, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x53, 0x72, 0x76,
	0x50, 0x69, 0x6e, 0x67, 0x22, 0x34, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x28, 0x0a, 0x07, 0x53, 0x72, 0x76, 0x50, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x53, 0x72, 0x76, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x33, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x72, 0x76, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x03, 0x53, 0x72,
	0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x03, 0x53, 0x72, 0x76, 0x22,
	0x34, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x72, 0x76, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x23, 0x0a, 0x03, 0x53, 0x72, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x03, 0x53, 0x72, 0x76, 0x22, 0x2d, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x22, 0x0a, 0x07, 0x53, 0x72, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x53, 0x72, 0x76,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x22, 0x0a, 0x07, 0x53, 0x72, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x53, 0x72, 0x76,
	0x49, 0x6e, 0x66, 0x6f, 0x32, 0xb2, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2f, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x21, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x72, 0x76, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x72, 0x76, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x72, 0x76, 0x52, 0x65, 0x73, 0x70, 0x42, 0x4d, 0x42, 0x0d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x2f, 0x64, 0x6f, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_transport_protocol_protocol_proto_rawDescOnce sync.Once
	file_pkg_transport_protocol_protocol_proto_rawDescData = file_pkg_transport_protocol_protocol_proto_rawDesc
)

func file_pkg_transport_protocol_protocol_proto_rawDescGZIP() []byte {
	file_pkg_transport_protocol_protocol_proto_rawDescOnce.Do(func() {
		file_pkg_transport_protocol_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_transport_protocol_protocol_proto_rawDescData)
	})
	return file_pkg_transport_protocol_protocol_proto_rawDescData
}

var file_pkg_transport_protocol_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pkg_transport_protocol_protocol_proto_goTypes = []interface{}{
	(*CreateService)(nil), // 0: pb.CreateService
	(*Info)(nil),          // 1: pb.Info
	(*ReportFormat)(nil),  // 2: pb.ReportFormat
	(*ReportInfo)(nil),    // 3: pb.ReportInfo
	(*ReportReq)(nil),     // 4: pb.ReportReq
	(*ReportResp)(nil),    // 5: pb.ReportResp
	(*ServerPing)(nil),    // 6: pb.ServerPing
	(*PingReq)(nil),       // 7: pb.PingReq
	(*PingResp)(nil),      // 8: pb.PingResp
	(*CreateSrvReq)(nil),  // 9: pb.CreateSrvReq
	(*CreateSrvResp)(nil), // 10: pb.CreateSrvResp
	(*InfoReq)(nil),       // 11: pb.InfoReq
	(*InfoResp)(nil),      // 12: pb.InfoResp
}
var file_pkg_transport_protocol_protocol_proto_depIdxs = []int32{
	3,  // 0: pb.ReportFormat.info:type_name -> pb.ReportInfo
	2,  // 1: pb.ReportReq.SrvAdr:type_name -> pb.ReportFormat
	2,  // 2: pb.ReportResp.SrvAdr:type_name -> pb.ReportFormat
	6,  // 3: pb.PingReq.SrvPing:type_name -> pb.ServerPing
	6,  // 4: pb.PingResp.SrvPing:type_name -> pb.ServerPing
	0,  // 5: pb.CreateSrvReq.Srv:type_name -> pb.CreateService
	0,  // 6: pb.CreateSrvResp.Srv:type_name -> pb.CreateService
	1,  // 7: pb.InfoReq.SrvInfo:type_name -> pb.Info
	1,  // 8: pb.InfoResp.SrvInfo:type_name -> pb.Info
	4,  // 9: pb.Service.GenerateReport:input_type -> pb.ReportReq
	7,  // 10: pb.Service.Ping:input_type -> pb.PingReq
	11, // 11: pb.Service.Info:input_type -> pb.InfoReq
	9,  // 12: pb.Service.CreateSrv:input_type -> pb.CreateSrvReq
	5,  // 13: pb.Service.GenerateReport:output_type -> pb.ReportResp
	8,  // 14: pb.Service.Ping:output_type -> pb.PingResp
	12, // 15: pb.Service.Info:output_type -> pb.InfoResp
	10, // 16: pb.Service.CreateSrv:output_type -> pb.CreateSrvResp
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_pkg_transport_protocol_protocol_proto_init() }
func file_pkg_transport_protocol_protocol_proto_init() {
	if File_pkg_transport_protocol_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_transport_protocol_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateService); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportFormat); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportInfo); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportReq); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportResp); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerPing); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResp); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSrvReq); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSrvResp); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoReq); i {
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
		file_pkg_transport_protocol_protocol_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResp); i {
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
			RawDescriptor: file_pkg_transport_protocol_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_transport_protocol_protocol_proto_goTypes,
		DependencyIndexes: file_pkg_transport_protocol_protocol_proto_depIdxs,
		MessageInfos:      file_pkg_transport_protocol_protocol_proto_msgTypes,
	}.Build()
	File_pkg_transport_protocol_protocol_proto = out.File
	file_pkg_transport_protocol_protocol_proto_rawDesc = nil
	file_pkg_transport_protocol_protocol_proto_goTypes = nil
	file_pkg_transport_protocol_protocol_proto_depIdxs = nil
}
