// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PushCode int32

const (
	PushCode_PC_ADD_DEFAULT         PushCode = 0
	PushCode_PC_ADD_FRIEND          PushCode = 100
	PushCode_PC_AGREE_ADD_FRIEND    PushCode = 101
	PushCode_PC_UPDATE_GROUP        PushCode = 110
	PushCode_PC_ADD_GROUP_MEMBERS   PushCode = 120
	PushCode_PC_REMOVE_GROUP_MEMBER PushCode = 121
)

var PushCode_name = map[int32]string{
	0:   "PC_ADD_DEFAULT",
	100: "PC_ADD_FRIEND",
	101: "PC_AGREE_ADD_FRIEND",
	110: "PC_UPDATE_GROUP",
	120: "PC_ADD_GROUP_MEMBERS",
	121: "PC_REMOVE_GROUP_MEMBER",
}

var PushCode_value = map[string]int32{
	"PC_ADD_DEFAULT":         0,
	"PC_ADD_FRIEND":          100,
	"PC_AGREE_ADD_FRIEND":    101,
	"PC_UPDATE_GROUP":        110,
	"PC_ADD_GROUP_MEMBERS":   120,
	"PC_REMOVE_GROUP_MEMBER": 121,
}

func (x PushCode) String() string {
	return proto.EnumName(PushCode_name, int32(x))
}

func (PushCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{0}
}

// 推送码 PC_ADD_FRIEND = 100
type AddFriendPush struct {
	FriendId             int64    `protobuf:"varint,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFriendPush) Reset()         { *m = AddFriendPush{} }
func (m *AddFriendPush) String() string { return proto.CompactTextString(m) }
func (*AddFriendPush) ProtoMessage()    {}
func (*AddFriendPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{0}
}

func (m *AddFriendPush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFriendPush.Unmarshal(m, b)
}
func (m *AddFriendPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFriendPush.Marshal(b, m, deterministic)
}
func (m *AddFriendPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFriendPush.Merge(m, src)
}
func (m *AddFriendPush) XXX_Size() int {
	return xxx_messageInfo_AddFriendPush.Size(m)
}
func (m *AddFriendPush) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFriendPush.DiscardUnknown(m)
}

var xxx_messageInfo_AddFriendPush proto.InternalMessageInfo

func (m *AddFriendPush) GetFriendId() int64 {
	if m != nil {
		return m.FriendId
	}
	return 0
}

func (m *AddFriendPush) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AddFriendPush) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *AddFriendPush) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// 推送码 PC_AGREE_ADD_FRIEND = 101
type AgreeAddFriendPush struct {
	FriendId             int64    `protobuf:"varint,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgreeAddFriendPush) Reset()         { *m = AgreeAddFriendPush{} }
func (m *AgreeAddFriendPush) String() string { return proto.CompactTextString(m) }
func (*AgreeAddFriendPush) ProtoMessage()    {}
func (*AgreeAddFriendPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{1}
}

func (m *AgreeAddFriendPush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgreeAddFriendPush.Unmarshal(m, b)
}
func (m *AgreeAddFriendPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgreeAddFriendPush.Marshal(b, m, deterministic)
}
func (m *AgreeAddFriendPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgreeAddFriendPush.Merge(m, src)
}
func (m *AgreeAddFriendPush) XXX_Size() int {
	return xxx_messageInfo_AgreeAddFriendPush.Size(m)
}
func (m *AgreeAddFriendPush) XXX_DiscardUnknown() {
	xxx_messageInfo_AgreeAddFriendPush.DiscardUnknown(m)
}

var xxx_messageInfo_AgreeAddFriendPush proto.InternalMessageInfo

func (m *AgreeAddFriendPush) GetFriendId() int64 {
	if m != nil {
		return m.FriendId
	}
	return 0
}

func (m *AgreeAddFriendPush) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AgreeAddFriendPush) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

// 更新群组 PC_UPDATE_GROUP = 110
type UpdateGroupPush struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Introduction         string   `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Extra                string   `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateGroupPush) Reset()         { *m = UpdateGroupPush{} }
func (m *UpdateGroupPush) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupPush) ProtoMessage()    {}
func (*UpdateGroupPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{2}
}

func (m *UpdateGroupPush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupPush.Unmarshal(m, b)
}
func (m *UpdateGroupPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupPush.Marshal(b, m, deterministic)
}
func (m *UpdateGroupPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupPush.Merge(m, src)
}
func (m *UpdateGroupPush) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupPush.Size(m)
}
func (m *UpdateGroupPush) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupPush.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupPush proto.InternalMessageInfo

func (m *UpdateGroupPush) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateGroupPush) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UpdateGroupPush) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateGroupPush) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *UpdateGroupPush) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

// 添加群组成员 PC_AGREE_ADD_GROUPS = 120
type AddGroupMembersPush struct {
	UserId               int64          `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Nickname             string         `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Members              []*GroupMember `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddGroupMembersPush) Reset()         { *m = AddGroupMembersPush{} }
func (m *AddGroupMembersPush) String() string { return proto.CompactTextString(m) }
func (*AddGroupMembersPush) ProtoMessage()    {}
func (*AddGroupMembersPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{3}
}

func (m *AddGroupMembersPush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddGroupMembersPush.Unmarshal(m, b)
}
func (m *AddGroupMembersPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddGroupMembersPush.Marshal(b, m, deterministic)
}
func (m *AddGroupMembersPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddGroupMembersPush.Merge(m, src)
}
func (m *AddGroupMembersPush) XXX_Size() int {
	return xxx_messageInfo_AddGroupMembersPush.Size(m)
}
func (m *AddGroupMembersPush) XXX_DiscardUnknown() {
	xxx_messageInfo_AddGroupMembersPush.DiscardUnknown(m)
}

var xxx_messageInfo_AddGroupMembersPush proto.InternalMessageInfo

func (m *AddGroupMembersPush) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AddGroupMembersPush) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AddGroupMembersPush) GetMembers() []*GroupMember {
	if m != nil {
		return m.Members
	}
	return nil
}

// 删除群组成员 PC_REMOVE_GROUP_MEMBER = 121
type RemoveGroupMemberPush struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	DeletedUserId        int64    `protobuf:"varint,3,opt,name=deleted_user_id,json=deletedUserId,proto3" json:"deleted_user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveGroupMemberPush) Reset()         { *m = RemoveGroupMemberPush{} }
func (m *RemoveGroupMemberPush) String() string { return proto.CompactTextString(m) }
func (*RemoveGroupMemberPush) ProtoMessage()    {}
func (*RemoveGroupMemberPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{4}
}

func (m *RemoveGroupMemberPush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveGroupMemberPush.Unmarshal(m, b)
}
func (m *RemoveGroupMemberPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveGroupMemberPush.Marshal(b, m, deterministic)
}
func (m *RemoveGroupMemberPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveGroupMemberPush.Merge(m, src)
}
func (m *RemoveGroupMemberPush) XXX_Size() int {
	return xxx_messageInfo_RemoveGroupMemberPush.Size(m)
}
func (m *RemoveGroupMemberPush) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveGroupMemberPush.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveGroupMemberPush proto.InternalMessageInfo

func (m *RemoveGroupMemberPush) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RemoveGroupMemberPush) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *RemoveGroupMemberPush) GetDeletedUserId() int64 {
	if m != nil {
		return m.DeletedUserId
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.PushCode", PushCode_name, PushCode_value)
	proto.RegisterType((*AddFriendPush)(nil), "pb.AddFriendPush")
	proto.RegisterType((*AgreeAddFriendPush)(nil), "pb.AgreeAddFriendPush")
	proto.RegisterType((*UpdateGroupPush)(nil), "pb.UpdateGroupPush")
	proto.RegisterType((*AddGroupMembersPush)(nil), "pb.AddGroupMembersPush")
	proto.RegisterType((*RemoveGroupMemberPush)(nil), "pb.RemoveGroupMemberPush")
}

func init() { proto.RegisterFile("push.proto", fileDescriptor_d1e4bfd2e9d102bb) }

var fileDescriptor_d1e4bfd2e9d102bb = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xc9, 0xb2, 0x3f, 0xed, 0x3b, 0x4a, 0x82, 0x3b, 0x58, 0x54, 0x84, 0x54, 0xe5, 0x80,
	0x0a, 0x87, 0x1e, 0xe0, 0x13, 0x84, 0xc6, 0xad, 0x26, 0x51, 0x16, 0x99, 0x85, 0xab, 0x95, 0xc4,
	0x2f, 0x5b, 0x44, 0x12, 0x47, 0x8e, 0x33, 0x95, 0x8f, 0xc0, 0x1d, 0xbe, 0x2f, 0x8a, 0xd3, 0x41,
	0x76, 0x42, 0xda, 0x61, 0x37, 0x3f, 0xbf, 0x47, 0xef, 0xf3, 0xbc, 0xb6, 0x64, 0x80, 0xba, 0x6d,
	0x6e, 0x96, 0xb5, 0x92, 0x5a, 0x92, 0x83, 0x3a, 0x9d, 0x39, 0x85, 0xbc, 0xce, 0xb3, 0x25, 0xee,
	0x74, 0x0f, 0xfd, 0x9f, 0x16, 0x4c, 0x02, 0x21, 0xd6, 0x2a, 0xc7, 0x4a, 0x44, 0x6d, 0x73, 0x43,
	0x5e, 0xc1, 0xf8, 0x9b, 0x51, 0x3c, 0x17, 0x9e, 0x35, 0xb7, 0x16, 0x36, 0x1b, 0xf5, 0xe0, 0x42,
	0x90, 0x19, 0x8c, 0xaa, 0x3c, 0xfb, 0x5e, 0x25, 0x25, 0x7a, 0x07, 0x73, 0x6b, 0x31, 0x66, 0x7f,
	0x35, 0x79, 0x0d, 0x90, 0xdc, 0x26, 0x3a, 0x51, 0xbc, 0x55, 0x85, 0x67, 0x1b, 0x77, 0xdc, 0x93,
	0x58, 0x15, 0x64, 0x0e, 0xa7, 0x02, 0x9b, 0x4c, 0xe5, 0xb5, 0xce, 0x65, 0xe5, 0x1d, 0x1a, 0x7f,
	0x88, 0xfc, 0x02, 0x48, 0x70, 0xad, 0x10, 0x1f, 0x65, 0x1f, 0xff, 0x97, 0x05, 0x4e, 0x5c, 0x8b,
	0x44, 0xe3, 0x46, 0xc9, 0xb6, 0x36, 0x5d, 0xe7, 0x70, 0xd2, 0x36, 0xa8, 0xfe, 0x35, 0x1d, 0x77,
	0xf2, 0x3f, 0x3d, 0x04, 0x0e, 0x0d, 0xef, 0x1b, 0xcc, 0x99, 0xf8, 0xf0, 0x34, 0xaf, 0xb4, 0x92,
	0xa2, 0xcd, 0x06, 0xb7, 0xbd, 0xc7, 0xc8, 0x19, 0x1c, 0xe1, 0x4e, 0xab, 0xc4, 0x3b, 0x32, 0x66,
	0x2f, 0xfc, 0x16, 0xa6, 0x81, 0x10, 0x66, 0xa5, 0x2d, 0x96, 0x29, 0xaa, 0xe6, 0xe1, 0x9b, 0xbd,
	0x85, 0x93, 0xb2, 0xcf, 0xf0, 0xec, 0xb9, 0xbd, 0x38, 0x7d, 0xef, 0x2c, 0xeb, 0x74, 0x39, 0xc8,
	0x66, 0x77, 0xbe, 0xaf, 0xe1, 0x05, 0xc3, 0x52, 0xde, 0xe2, 0xc0, 0x7d, 0x78, 0xf1, 0x1b, 0x70,
	0x04, 0x16, 0xa8, 0x51, 0xf0, 0xbb, 0x61, 0xdb, 0x0c, 0x4f, 0xf6, 0x38, 0x36, 0x19, 0xef, 0x7e,
	0x5b, 0x30, 0xea, 0x5a, 0x56, 0x52, 0x74, 0xef, 0xf8, 0x2c, 0x5a, 0xf1, 0x20, 0x0c, 0x79, 0x48,
	0xd7, 0x41, 0xfc, 0xe9, 0xca, 0x7d, 0x42, 0x9e, 0xc3, 0x64, 0xcf, 0xd6, 0xec, 0x82, 0x7e, 0x0e,
	0x5d, 0x41, 0xce, 0x61, 0xda, 0xa1, 0x0d, 0xa3, 0x74, 0x68, 0x20, 0x99, 0x82, 0x13, 0xad, 0x78,
	0x1c, 0x85, 0xc1, 0x15, 0xe5, 0x1b, 0x76, 0x19, 0x47, 0x6e, 0x45, 0x3c, 0x38, 0xdb, 0x07, 0x18,
	0xc2, 0xb7, 0x74, 0xfb, 0x91, 0xb2, 0x2f, 0xee, 0x8e, 0xcc, 0xe0, 0x65, 0xb4, 0xe2, 0x8c, 0x6e,
	0x2f, 0xbf, 0xd2, 0x7b, 0xa6, 0xfb, 0x23, 0x3d, 0x36, 0x9f, 0xe3, 0xc3, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0x3e, 0x0b, 0x3f, 0x3f, 0x03, 0x00, 0x00,
}
