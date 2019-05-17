// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: configurator/configurator.proto

package configurator

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	linux "github.com/ligato/vpp-agent/api/models/linux"
	vpp "github.com/ligato/vpp-agent/api/models/vpp"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Config groups all supported config data into single message.
type Config struct {
	VppConfig            *vpp.ConfigData   `protobuf:"bytes,1,opt,name=vpp_config,json=vppConfig,proto3" json:"vpp_config,omitempty"`
	LinuxConfig          *linux.ConfigData `protobuf:"bytes,2,opt,name=linux_config,json=linuxConfig,proto3" json:"linux_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetVppConfig() *vpp.ConfigData {
	if m != nil {
		return m.VppConfig
	}
	return nil
}

func (m *Config) GetLinuxConfig() *linux.ConfigData {
	if m != nil {
		return m.LinuxConfig
	}
	return nil
}

// Notification groups all notification data into single message.
type Notification struct {
	// Types that are valid to be assigned to Notification:
	//	*Notification_VppNotification
	//	*Notification_LinuxNotification
	Notification         isNotification_Notification `protobuf_oneof:"notification"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{1}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

type isNotification_Notification interface {
	isNotification_Notification()
}

type Notification_VppNotification struct {
	VppNotification *vpp.Notification `protobuf:"bytes,1,opt,name=vpp_notification,json=vppNotification,proto3,oneof"`
}
type Notification_LinuxNotification struct {
	LinuxNotification *linux.Notification `protobuf:"bytes,2,opt,name=linux_notification,json=linuxNotification,proto3,oneof"`
}

func (*Notification_VppNotification) isNotification_Notification()   {}
func (*Notification_LinuxNotification) isNotification_Notification() {}

func (m *Notification) GetNotification() isNotification_Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

func (m *Notification) GetVppNotification() *vpp.Notification {
	if x, ok := m.GetNotification().(*Notification_VppNotification); ok {
		return x.VppNotification
	}
	return nil
}

func (m *Notification) GetLinuxNotification() *linux.Notification {
	if x, ok := m.GetNotification().(*Notification_LinuxNotification); ok {
		return x.LinuxNotification
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Notification) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Notification_OneofMarshaler, _Notification_OneofUnmarshaler, _Notification_OneofSizer, []interface{}{
		(*Notification_VppNotification)(nil),
		(*Notification_LinuxNotification)(nil),
	}
}

func _Notification_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Notification)
	// notification
	switch x := m.Notification.(type) {
	case *Notification_VppNotification:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VppNotification); err != nil {
			return err
		}
	case *Notification_LinuxNotification:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LinuxNotification); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Notification.Notification has unexpected type %T", x)
	}
	return nil
}

func _Notification_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Notification)
	switch tag {
	case 1: // notification.vpp_notification
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(vpp.Notification)
		err := b.DecodeMessage(msg)
		m.Notification = &Notification_VppNotification{msg}
		return true, err
	case 2: // notification.linux_notification
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(linux.Notification)
		err := b.DecodeMessage(msg)
		m.Notification = &Notification_LinuxNotification{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Notification_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Notification)
	// notification
	switch x := m.Notification.(type) {
	case *Notification_VppNotification:
		s := proto.Size(x.VppNotification)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Notification_LinuxNotification:
		s := proto.Size(x.LinuxNotification)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type UpdateRequest struct {
	// Update describes config data to be updated.
	Update *Config `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
	// The full_resync option can be used
	// to overwrite all existing data.
	FullResync           bool     `protobuf:"varint,2,opt,name=full_resync,json=fullResync,proto3" json:"full_resync,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{2}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetUpdate() *Config {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *UpdateRequest) GetFullResync() bool {
	if m != nil {
		return m.FullResync
	}
	return false
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{3}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type DeleteRequest struct {
	// Delete describes config data to be deleted.
	Delete               *Config  `protobuf:"bytes,1,opt,name=delete,proto3" json:"delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{4}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetDelete() *Config {
	if m != nil {
		return m.Delete
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{5}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{6}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type GetResponse struct {
	// Config describes desired config retrieved from agent.
	Config               *Config  `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{7}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type DumpRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpRequest) Reset()         { *m = DumpRequest{} }
func (m *DumpRequest) String() string { return proto.CompactTextString(m) }
func (*DumpRequest) ProtoMessage()    {}
func (*DumpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{8}
}
func (m *DumpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpRequest.Unmarshal(m, b)
}
func (m *DumpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpRequest.Marshal(b, m, deterministic)
}
func (m *DumpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpRequest.Merge(m, src)
}
func (m *DumpRequest) XXX_Size() int {
	return xxx_messageInfo_DumpRequest.Size(m)
}
func (m *DumpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DumpRequest proto.InternalMessageInfo

type DumpResponse struct {
	// Dump describes running config dumped from southbound.
	Dump                 *Config  `protobuf:"bytes,1,opt,name=dump,proto3" json:"dump,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpResponse) Reset()         { *m = DumpResponse{} }
func (m *DumpResponse) String() string { return proto.CompactTextString(m) }
func (*DumpResponse) ProtoMessage()    {}
func (*DumpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{9}
}
func (m *DumpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpResponse.Unmarshal(m, b)
}
func (m *DumpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpResponse.Marshal(b, m, deterministic)
}
func (m *DumpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpResponse.Merge(m, src)
}
func (m *DumpResponse) XXX_Size() int {
	return xxx_messageInfo_DumpResponse.Size(m)
}
func (m *DumpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DumpResponse proto.InternalMessageInfo

func (m *DumpResponse) GetDump() *Config {
	if m != nil {
		return m.Dump
	}
	return nil
}

// NotificationRequest represent a notification request which contains
// index of next required message
type NotificationRequest struct {
	Idx                  uint32   `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationRequest) Reset()         { *m = NotificationRequest{} }
func (m *NotificationRequest) String() string { return proto.CompactTextString(m) }
func (*NotificationRequest) ProtoMessage()    {}
func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{10}
}
func (m *NotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationRequest.Unmarshal(m, b)
}
func (m *NotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationRequest.Marshal(b, m, deterministic)
}
func (m *NotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationRequest.Merge(m, src)
}
func (m *NotificationRequest) XXX_Size() int {
	return xxx_messageInfo_NotificationRequest.Size(m)
}
func (m *NotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationRequest proto.InternalMessageInfo

func (m *NotificationRequest) GetIdx() uint32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

// Response to notification request 'get'. Returns indexed notification.
type NotificationResponse struct {
	// Index of following notification
	NextIdx uint32 `protobuf:"varint,1,opt,name=next_idx,json=nextIdx,proto3" json:"next_idx,omitempty"`
	// Notification data
	Notification         *Notification `protobuf:"bytes,2,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NotificationResponse) Reset()         { *m = NotificationResponse{} }
func (m *NotificationResponse) String() string { return proto.CompactTextString(m) }
func (*NotificationResponse) ProtoMessage()    {}
func (*NotificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_150898d063c79000, []int{11}
}
func (m *NotificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationResponse.Unmarshal(m, b)
}
func (m *NotificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationResponse.Marshal(b, m, deterministic)
}
func (m *NotificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationResponse.Merge(m, src)
}
func (m *NotificationResponse) XXX_Size() int {
	return xxx_messageInfo_NotificationResponse.Size(m)
}
func (m *NotificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationResponse proto.InternalMessageInfo

func (m *NotificationResponse) GetNextIdx() uint32 {
	if m != nil {
		return m.NextIdx
	}
	return 0
}

func (m *NotificationResponse) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "configurator.Config")
	proto.RegisterType((*Notification)(nil), "configurator.Notification")
	proto.RegisterType((*UpdateRequest)(nil), "configurator.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "configurator.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "configurator.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "configurator.DeleteResponse")
	proto.RegisterType((*GetRequest)(nil), "configurator.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "configurator.GetResponse")
	proto.RegisterType((*DumpRequest)(nil), "configurator.DumpRequest")
	proto.RegisterType((*DumpResponse)(nil), "configurator.DumpResponse")
	proto.RegisterType((*NotificationRequest)(nil), "configurator.NotificationRequest")
	proto.RegisterType((*NotificationResponse)(nil), "configurator.NotificationResponse")
}

func init() { proto.RegisterFile("configurator/configurator.proto", fileDescriptor_150898d063c79000) }

var fileDescriptor_150898d063c79000 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe9, 0x36, 0x85, 0xf1, 0x92, 0x6e, 0xad, 0xd7, 0x43, 0x17, 0x90, 0x06, 0xbe, 0xb0,
	0x03, 0xa4, 0x68, 0x70, 0x00, 0xaa, 0xed, 0xb0, 0x56, 0x1a, 0x5c, 0x90, 0x88, 0xc4, 0x85, 0x03,
	0x55, 0xd6, 0xb8, 0x25, 0x52, 0x6a, 0x7b, 0x8d, 0x53, 0x75, 0xdf, 0x87, 0xcf, 0xc6, 0xe7, 0x40,
	0xf6, 0x73, 0x54, 0xbb, 0x2a, 0x3d, 0x6c, 0xea, 0xfb, 0xbf, 0xff, 0xfb, 0xbd, 0x17, 0x3f, 0xcb,
	0x70, 0x31, 0x15, 0x7c, 0x56, 0xcc, 0xeb, 0x65, 0xa6, 0xc4, 0x72, 0xe0, 0x06, 0x89, 0x5c, 0x0a,
	0x25, 0x48, 0xe4, 0x6a, 0x71, 0x6f, 0x21, 0x72, 0x56, 0x56, 0x83, 0x95, 0x94, 0xfa, 0x0f, 0x3d,
	0x71, 0xdf, 0xaa, 0x65, 0xc1, 0xeb, 0x35, 0xfe, 0xc7, 0x0c, 0xe5, 0x10, 0x8c, 0x4c, 0x3d, 0x49,
	0x00, 0x56, 0x52, 0x4e, 0x90, 0xd6, 0x6f, 0xbd, 0x6c, 0x5d, 0x86, 0x57, 0xa7, 0x89, 0x66, 0xa0,
	0x61, 0x9c, 0xa9, 0x2c, 0x7d, 0xb6, 0x92, 0xd2, 0xfa, 0x3f, 0x40, 0x64, 0x40, 0x4d, 0xc5, 0x81,
	0xa9, 0xe8, 0x26, 0x48, 0x77, 0x6a, 0x42, 0xa3, 0xa0, 0x40, 0xff, 0xb4, 0x20, 0xfa, 0x26, 0x54,
	0x31, 0x2b, 0xa6, 0x99, 0x2a, 0x04, 0x27, 0x37, 0xd0, 0xd1, 0x6d, 0xb9, 0xa3, 0xd9, 0xe6, 0x5d,
	0xd3, 0xdc, 0x35, 0x7f, 0x79, 0x92, 0x9e, 0xae, 0xa4, 0xf4, 0xea, 0xc7, 0x40, 0x70, 0x0c, 0x8f,
	0x80, 0xc3, 0x9c, 0xd9, 0x61, 0xb6, 0x18, 0x5d, 0xa3, 0xba, 0xe2, 0xed, 0x09, 0x44, 0x6e, 0x3d,
	0xfd, 0x05, 0xed, 0x1f, 0x32, 0xcf, 0x14, 0x4b, 0xd9, 0x43, 0xcd, 0x2a, 0x45, 0xde, 0x40, 0x50,
	0x1b, 0xc1, 0x0e, 0xd7, 0x4b, 0xbc, 0x55, 0xe0, 0xd7, 0xa5, 0xd6, 0x43, 0x2e, 0x20, 0x9c, 0xd5,
	0x65, 0x39, 0x59, 0xb2, 0xea, 0x91, 0x4f, 0xcd, 0x34, 0xc7, 0x29, 0x68, 0x29, 0x35, 0x0a, 0xed,
	0xc0, 0x49, 0xc3, 0xaf, 0xa4, 0xe0, 0x15, 0xa3, 0xd7, 0xd0, 0x1e, 0xb3, 0x92, 0x79, 0x1d, 0x73,
	0x23, 0xec, 0xef, 0x88, 0x1e, 0x0d, 0x6c, 0xca, 0x2d, 0x30, 0x02, 0xb8, 0x63, 0xca, 0xd2, 0xe8,
	0x10, 0x42, 0x13, 0x61, 0x52, 0xc3, 0xbd, 0x45, 0xff, 0x07, 0x8e, 0x22, 0x6d, 0x43, 0x38, 0xae,
	0x17, 0xb2, 0x61, 0x7d, 0x84, 0x08, 0x43, 0x0b, 0xbb, 0x84, 0xa3, 0xbc, 0x5e, 0xc8, 0xbd, 0x28,
	0xe3, 0xa0, 0xaf, 0xe1, 0xcc, 0x3d, 0xf6, 0xe6, 0x53, 0x3b, 0x70, 0x58, 0xe4, 0x6b, 0x53, 0xdf,
	0x4e, 0xf5, 0x4f, 0xfa, 0x00, 0x3d, 0xdf, 0x68, 0x5b, 0x9d, 0xc3, 0x31, 0x67, 0x6b, 0x35, 0xd9,
	0xd8, 0x9f, 0xea, 0xf8, 0x6b, 0xbe, 0x26, 0x37, 0xfe, 0x0a, 0xed, 0x15, 0x88, 0xfd, 0x69, 0x3c,
	0xa8, 0xe7, 0xbf, 0xfa, 0x7b, 0x00, 0xd1, 0xc8, 0xf1, 0x92, 0xcf, 0x70, 0x78, 0xc7, 0x14, 0xe9,
	0xfb, 0x84, 0xcd, 0x99, 0xc6, 0xe7, 0x3b, 0x32, 0x76, 0xce, 0x11, 0x04, 0xb8, 0x5f, 0xf2, 0xdc,
	0x37, 0x79, 0xb7, 0x2a, 0x7e, 0xb1, 0x3b, 0xb9, 0x81, 0xe0, 0x4e, 0xb7, 0x21, 0xde, 0x45, 0xd9,
	0x86, 0xf8, 0xd7, 0x80, 0x5c, 0xc3, 0x91, 0x5e, 0x16, 0xd9, 0x1a, 0xd6, 0xd9, 0x67, 0x1c, 0xef,
	0x4a, 0xd9, 0xf2, 0xef, 0x10, 0x98, 0x33, 0x7b, 0x24, 0xaf, 0xf6, 0x9c, 0xa4, 0x05, 0xd1, 0x7d,
	0x16, 0x04, 0xbe, 0x6b, 0xdd, 0x0e, 0x7f, 0x7e, 0x9a, 0x17, 0xea, 0x77, 0x7d, 0x9f, 0x4c, 0xc5,
	0x62, 0x50, 0x16, 0xf3, 0x4c, 0x09, 0xfd, 0x56, 0xbd, 0xcd, 0xe6, 0x8c, 0xab, 0x41, 0x26, 0x0b,
	0xef, 0x99, 0x1b, 0xba, 0xc1, 0x7d, 0x60, 0x9e, 0xad, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x51, 0xa8, 0x29, 0xab, 0x17, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfiguratorClient is the client API for Configurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfiguratorClient interface {
	// Get is used for listing desired config.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Update is used for updating desired config.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete is used for deleting desired config.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Dump is used for dumping running config.
	Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (*DumpResponse, error)
	// Notify is used for subscribing to notifications.
	Notify(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (Configurator_NotifyClient, error)
}

type configuratorClient struct {
	cc *grpc.ClientConn
}

func NewConfiguratorClient(cc *grpc.ClientConn) ConfiguratorClient {
	return &configuratorClient{cc}
}

func (c *configuratorClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/configurator.Configurator/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configuratorClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/configurator.Configurator/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configuratorClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/configurator.Configurator/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configuratorClient) Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (*DumpResponse, error) {
	out := new(DumpResponse)
	err := c.cc.Invoke(ctx, "/configurator.Configurator/Dump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configuratorClient) Notify(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (Configurator_NotifyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Configurator_serviceDesc.Streams[0], "/configurator.Configurator/Notify", opts...)
	if err != nil {
		return nil, err
	}
	x := &configuratorNotifyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Configurator_NotifyClient interface {
	Recv() (*NotificationResponse, error)
	grpc.ClientStream
}

type configuratorNotifyClient struct {
	grpc.ClientStream
}

func (x *configuratorNotifyClient) Recv() (*NotificationResponse, error) {
	m := new(NotificationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConfiguratorServer is the server API for Configurator service.
type ConfiguratorServer interface {
	// Get is used for listing desired config.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Update is used for updating desired config.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete is used for deleting desired config.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Dump is used for dumping running config.
	Dump(context.Context, *DumpRequest) (*DumpResponse, error)
	// Notify is used for subscribing to notifications.
	Notify(*NotificationRequest, Configurator_NotifyServer) error
}

func RegisterConfiguratorServer(s *grpc.Server, srv ConfiguratorServer) {
	s.RegisterService(&_Configurator_serviceDesc, srv)
}

func _Configurator_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configurator.Configurator/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configurator_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configurator.Configurator/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configurator_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configurator.Configurator/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configurator_Dump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServer).Dump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configurator.Configurator/Dump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServer).Dump(ctx, req.(*DumpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configurator_Notify_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotificationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfiguratorServer).Notify(m, &configuratorNotifyServer{stream})
}

type Configurator_NotifyServer interface {
	Send(*NotificationResponse) error
	grpc.ServerStream
}

type configuratorNotifyServer struct {
	grpc.ServerStream
}

func (x *configuratorNotifyServer) Send(m *NotificationResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Configurator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "configurator.Configurator",
	HandlerType: (*ConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Configurator_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Configurator_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Configurator_Delete_Handler,
		},
		{
			MethodName: "Dump",
			Handler:    _Configurator_Dump_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Notify",
			Handler:       _Configurator_Notify_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "configurator/configurator.proto",
}